package quadtree

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.

func (q Quadtree) getNumberFromQuad(indexX, indexY, height, width int, current *node) int {
	var currentWidth int = (*current).width
	var currentHeight int = (*current).height
	if indexX < 0 || indexY < 0 || indexX >= width || indexY >= height { // si les position sont en dehors du terrain
		return -1
	}
	if (*current).content != -1 { //si le node actuelle possede une valeur
		return (*current).content
	}
	if (currentWidth)%2 != 0 { //si la largeur du node actuelle est impair la largeur des nodes enfants a Gauche sera augmenté
		currentWidth += 1
	}
	if (currentHeight)%2 != 0 { //si la hauteur du node actuelle est impair la hauteur des nodes enfants en Haut sera augmenté
		currentHeight += 1
	}
	var xChoice int = 0
	if (*current).TopLeftX+currentWidth/2 > indexX {
		xChoice = 1
	} else if (*current).TopLeftX+current.width > indexX {
		xChoice = 2
	}
	var yChoice int = 0
	if (*current).TopLeftY+currentHeight/2 > indexY {
		yChoice = 1
	} else if (*current).TopLeftY+(*current).height > indexY {
		yChoice = 2
	}
	var suivant *node
	if xChoice == 1 && yChoice == 1 {
		suivant = (*current).topLeftNode
	} else if xChoice == 2 && yChoice == 1 {
		suivant = (*current).topRightNode
	} else if xChoice == 1 && yChoice == 2 {
		suivant = (*current).bottomLeftNode
	} else if xChoice == 2 && yChoice == 2 {
		suivant = (*current).bottomRightNode
	} else {
		return -1
	}
	return q.getNumberFromQuad(indexX, indexY, height, width, suivant)
}

func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	for i := 0; i < len(contentHolder); i++ {
		for x := 0; x < len(contentHolder[i]); x++ {
			contentHolder[i][x] = q.getNumberFromQuad(topLeftX+x, topLeftY+i, q.height, q.width, q.root) // pour chaque bloc du terrain a afficher on recupere la valeur du bloc avec ses coordonnées
		}
	}
}
