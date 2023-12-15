package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.

func Recur(position string, flootContent [][]int, parent node) node {
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
	if position == "Root" { // On définit le node actuelle en fonction du sa position dans le node Parent
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
	if !(laNode.width == 1 && laNode.height == 1) { //on vérifie si tout les blocs dans le node actuelle sont les mêmes
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
	if (laNode.width == 1 && laNode.height == 1) || (state) { // si le node actuelle ne contient que un bloc ou si tout les bloc du node sont les mêmes
		laNode.content = flootContent[laNode.TopLeftY][laNode.TopLeftX]
	} else {
		var topLeftNode node = Recur("topLeft", flootContent, laNode) //il y'aura toujours un topLeftNode
		laNode.topLeftNode = &topLeftNode
		if laNode.width == 1 || laNode.height == 1 { // si on a l'une des deux dimensions qui est égale a 1
			if laNode.width <= 1 { // si le node ne fait que 1 de largeur il n'y aura pas de node a droite
				var bottomLeftNode node = Recur("bottomLeft", flootContent, laNode)
				laNode.bottomLeftNode = &bottomLeftNode
			} else { // si le node ne fait que 1 de hauteur il n'y aura pas de node en bas
				var topRightNode node = Recur("topRight", flootContent, laNode)
				laNode.topRightNode = &topRightNode
			}
		} else { //sinon on le node actuelle aura 4 nodes enfants
			var topRightNode node = Recur("topRight", flootContent, laNode)
			var bottomLeftNode node = Recur("bottomLeft", flootContent, laNode)
			var bottomRightNode node = Recur("bottomRight", flootContent, laNode)
			laNode.topRightNode = &topRightNode
			laNode.bottomLeftNode = &bottomLeftNode
			laNode.bottomRightNode = &bottomRightNode
		}
	}
	return laNode
}

func MakeFromArray(floorContent [][]int) (q Quadtree) {
	var Quad Quadtree = Quadtree{
		width: len(floorContent[0]), height: len(floorContent)}
	var Laroot node = Recur("Root", floorContent, node{width: Quad.width, height: Quad.height, TopLeftX: 0, TopLeftY: 0})
	Quad.root = &Laroot
	return Quad
}
