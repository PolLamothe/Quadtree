package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.

func Recur(position string, flootContent [][]int, parent Node) Node {
	var laNode Node
	var parentWidth int = parent.Width
	var parentHeight int = parent.Height
	var parentWidthHalf int = parentWidth
	var parentHeightHalf int = parentHeight
	if (parentWidth)%2 != 0 { //si la largeur est impaire on augmente la largeur des Node a Gauche et augmente le TopLeftX de ce de Droite
		if position == "topLeft" || position == "bottomLeft" {
			parentWidth += 1
		} else {
			parentWidthHalf += 1
		}
	}
	if (parentHeight)%2 != 0 { //si la hauteur est impaire on augmente la hauteur des Node en Haut et augmente le TopLeftY de ce en Bas
		if position == "topLeft" || position == "topRight" {
			parentHeight += 1
		} else {
			parentHeightHalf += 1
		}
	}
	if position == "Root" { // On définit le Node actuelle en fonction du sa position dans le Node Parent
		laNode = Node{TopLeftX: parent.TopLeftX, TopLeftY: parent.TopLeftY, Content: -1, Height: parentHeight, Width: parentWidth}
	} else if position == "topLeft" {
		laNode = Node{TopLeftX: parent.TopLeftX, TopLeftY: parent.TopLeftY, Content: -1, Width: parentWidth / 2, Height: parentHeight / 2}
	} else if position == "topRight" {
		laNode = Node{TopLeftX: parent.TopLeftX + parentWidthHalf/2, TopLeftY: parent.TopLeftY, Content: -1, Width: parentWidth / 2, Height: parentHeight / 2}
	} else if position == "bottomLeft" {
		laNode = Node{TopLeftX: parent.TopLeftX, TopLeftY: parent.TopLeftY + parentHeightHalf/2, Content: -1, Width: parentWidth / 2, Height: parentHeight / 2}
	} else if position == "bottomRight" {
		laNode = Node{TopLeftX: parent.TopLeftX + parentWidthHalf/2, TopLeftY: parent.TopLeftY + parentHeightHalf/2, Content: -1, Width: parentWidth / 2, Height: parentHeight / 2}
	}
	var state bool = false
	if !(laNode.Width == 1 && laNode.Height == 1) { //on vérifie si tout les blocs dans le Node actuelle sont les mêmes
		var origin int
		state = true
		var flootContentX, flootContentY int = laNode.TopLeftX, laNode.TopLeftY
		var XRange, YRange int = laNode.TopLeftX, laNode.TopLeftY
		origin = flootContent[flootContentY][flootContentX]
		for i := YRange; i < YRange+laNode.Height; i++ {
			for x := XRange; x < XRange+laNode.Width; x++ {
				if flootContent[i][x] != origin { //si jamais il y'a un des bloc dont le Node actuelle qui n'est pas égale aux autres
					state = false
				}
			}
		}
	}
	if (laNode.Width == 1 && laNode.Height == 1) || (state) { // si le Node actuelle ne contient que un bloc ou si tout les bloc du Node sont les mêmes
		laNode.Content = flootContent[laNode.TopLeftY][laNode.TopLeftX]
	} else {
		var topLeftNode Node = Recur("topLeft", flootContent, laNode) //il y'aura toujours un TopLeftNode
		laNode.TopLeftNode = &topLeftNode
		if laNode.Width == 1 || laNode.Height == 1 { // si on a l'une des deux dimensions qui est égale a 1
			if laNode.Width <= 1 { // si le Node ne fait que 1 de largeur il n'y aura pas de Node a droite
				var bottomLeftNode Node = Recur("bottomLeft", flootContent, laNode)
				laNode.BottomLeftNode = &bottomLeftNode
			} else { // si le Node ne fait que 1 de hauteur il n'y aura pas de Node en bas
				var topRightNode Node = Recur("topRight", flootContent, laNode)
				laNode.TopRightNode = &topRightNode
			}
		} else { //sinon on le Node actuelle aura 4 nodes enfants
			var topRightNode Node = Recur("topRight", flootContent, laNode)
			var bottomLeftNode Node = Recur("bottomLeft", flootContent, laNode)
			var bottomRightNode Node = Recur("bottomRight", flootContent, laNode)
			laNode.TopRightNode = &topRightNode
			laNode.BottomLeftNode = &bottomLeftNode
			laNode.BottomRightNode = &bottomRightNode
		}
	}
	return laNode
}

func MakeFromArray(floorContent [][]int) (q Quadtree) {
	var Quad Quadtree = Quadtree{
		width: len(floorContent[0]), height: len(floorContent)}
	var Laroot Node = Recur("Root", floorContent, Node{Width: Quad.width, Height: Quad.height, TopLeftX: 0, TopLeftY: 0})
	Quad.Root = &Laroot
	return Quad
}
