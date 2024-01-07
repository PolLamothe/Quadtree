package quadtree

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/multiplayer"
	"math/rand"
)

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.

func pop(slice []map[string]int, index int) []map[string]int {
	var result []map[string]int
	for i := 0; i < len(slice); i++ {
		if i != index {
			result = append(result, slice[i])
		}
	}
	return result
}

func Recur(position string, flootContent [][]int, parent node, initTopLeftX, initTopLeftY int) node {
	var laNode node
	var parentWidth int = parent.width
	var parentHeight int = parent.height
	var parentWidthHalf int = parentWidth
	var parentHeightHalf int = parentHeight
	if (parentWidth)%2 != 0 { //si la largeur est impaire on augmente la largeur des node a Gauche et augmente le TopLeftX de ce de Droite
		if position == "topLeft" || position == "bottomLeft" {
			parentWidth += 1
		} else {
			parentWidthHalf += 1
		}
	}
	if (parentHeight)%2 != 0 { //si la hauteur est impaire on augmente la hauteur des node en Haut et augmente le TopLeftY de ce en Bas
		if position == "topLeft" || position == "topRight" {
			parentHeight += 1
		} else {
			parentHeightHalf += 1
		}
	}
	if position == "Root" {
		laNode = node{TopLeftX: parent.TopLeftX, TopLeftY: parent.TopLeftY, content: -1, height: parentHeight, width: parentWidth}
	} else if position == "topLeft" {
		laNode = node{TopLeftX: parent.TopLeftX, TopLeftY: parent.TopLeftY, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	} else if position == "topRight" {
		laNode = node{TopLeftX: parent.TopLeftX + parentWidthHalf/2, TopLeftY: parent.TopLeftY, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	} else if position == "bottomLeft" {
		laNode = node{TopLeftX: parent.TopLeftX, TopLeftY: parent.TopLeftY + parentHeightHalf/2, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	} else if position == "bottomRight" {
		laNode = node{TopLeftX: parent.TopLeftX + parentWidthHalf/2, TopLeftY: parent.TopLeftY + parentHeightHalf/2, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	}
	var state bool = false
	if !(laNode.width == 1 && laNode.height == 1) && !configuration.Global.GenerationInfinie { //on vérifie si tout les blocs dans le node actuelle sont les mêmes
		var origin int
		state = true
		var flootContentX, flootContentY int = laNode.TopLeftX, laNode.TopLeftY
		var XRange, YRange int = laNode.TopLeftX, laNode.TopLeftY
		origin = flootContent[flootContentY][flootContentX]
		for i := YRange; i < YRange+laNode.height; i++ {
			for x := XRange; x < XRange+laNode.width; x++ {
				if flootContent[i][x] != origin { //si jamais il y'a un des bloc dont le node actuelle qui n'est pas égale aux autres
					state = false
				}
			}
		}
	}
	if (laNode.width == 1 && laNode.height == 1) || (state) {
		if !configuration.Global.GenerationInfinie { //on ne définie pas de valeur au bloc si on est en génération infinie car on les définiras lorsque ils seront chargés
			laNode.content = flootContent[laNode.TopLeftY][laNode.TopLeftX]
		} else {
			laNode.content = rand.Intn(5)
			if configuration.Global.GenerationInfinie && configuration.Global.MultiplayerKind != 0 {
				for i := 0; i < laNode.width; i++ {
					for x := 0; x < laNode.height; x++ {
						var state bool = false
						for y := 0; y < len(multiplayer.BlockReceived); y++ { // on s'assure que le bloc n'a pas déja été reçu
							if laNode.TopLeftX+i == multiplayer.BlockReceived[y]["X"] && laNode.TopLeftY+x == multiplayer.BlockReceived[y]["Y"] {
								laNode.content = multiplayer.BlockReceived[y]["Value"]
								multiplayer.BlockReceived = pop(multiplayer.BlockReceived, y)
								state = true
							}
						}
						if !state {
							multiplayer.BlockToSend = append(multiplayer.BlockToSend, map[string]int{"X": laNode.TopLeftX + i, "Y": laNode.TopLeftY + x, "Value": laNode.content})
							multiplayer.SendBlock()
						}
					}
				}
			}
		}
	} else if len(flootContent) > 0 {
		var topLeftNode node = Recur("topLeft", flootContent, laNode, initTopLeftX, initTopLeftY)
		laNode.topLeftNode = &topLeftNode
		if (laNode.width == 1 || laNode.height == 1) && (laNode.width != 1 || laNode.height != 1) { // si on a l'une des deux dimensions qui est égale a 1 mais pas les deux
			if laNode.width <= 1 {
				var bottomLeftNode node = Recur("bottomLeft", flootContent, laNode, initTopLeftX, initTopLeftY)
				laNode.bottomLeftNode = &bottomLeftNode
			} else {
				var topRightNode node = Recur("topRight", flootContent, laNode, initTopLeftX, initTopLeftY)
				laNode.topRightNode = &topRightNode
			}
		} else {
			var topRightNode node = Recur("topRight", flootContent, laNode, initTopLeftX, initTopLeftY)
			var bottomLeftNode node = Recur("bottomLeft", flootContent, laNode, initTopLeftX, initTopLeftY)
			var bottomRightNode node = Recur("bottomRight", flootContent, laNode, initTopLeftX, initTopLeftY)
			laNode.topRightNode = &topRightNode
			laNode.bottomLeftNode = &bottomLeftNode
			laNode.bottomRightNode = &bottomRightNode
		}
	}
	return laNode
}
func MakeFromArray(floorContent [][]int, width, height, TopLeftX, TopLeftY int) (q Quadtree) {
	var Quad Quadtree = Quadtree{
		Width: width, Height: height}
	var Laroot node = Recur("Root", floorContent, node{width: Quad.Width, height: Quad.Height, TopLeftX: TopLeftX, TopLeftY: TopLeftY}, TopLeftX, TopLeftY)
	Quad.Root = &Laroot
	return Quad
}
