package quadtree

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.

func (q Quadtree) getNumberFromQuad(indexX, indexY, height, width int, current node) int {
	var currentWidth int = current.width
	var currentHeight int = current.height
	if indexX < 0 || indexY < 0 || indexX > width || indexY > height {
		return -1
	}
	if current.content != -1 {
		return current.content
	}
	if (currentWidth)%2 != 0 {
		currentWidth += 1
	}
	if (currentHeight)%2 != 0 {
		currentHeight += 1
	}
	var xChoice int = 0
	if current.topLeftX+currentWidth/2 > indexX {
		xChoice = 1
	} else if current.topLeftX+current.width > indexX {
		xChoice = 2
	}
	var yChoice int = 0
	if current.topLeftY+currentHeight/2 > indexY {
		yChoice = 1
	} else if current.topLeftY+current.height > indexY {
		yChoice = 2
	}
	var suivant node
	if xChoice == 1 && yChoice == 1 {
		suivant = *current.topLeftNode
	} else if xChoice == 2 && yChoice == 1 {
		suivant = *current.topRightNode
	} else if xChoice == 1 && yChoice == 2 {
		suivant = *current.bottomLeftNode
	} else if xChoice == 2 && yChoice == 2 {
		suivant = *current.bottomRightNode
	} else {
		return -1
	}
	return q.getNumberFromQuad(indexX, indexY, height, width, suivant)
}

func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	for i := 0; i < len(contentHolder); i++ {
		for x := 0; x < len(contentHolder[i]); x++ {
			contentHolder[i][x] = q.getNumberFromQuad(topLeftX+x, topLeftY+i, q.Height, q.Width, *q.root)
		}
	}
}
