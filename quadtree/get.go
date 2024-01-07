package quadtree

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (TopLeftX, TopLeftY)) à partir du qadtree q.

func (q Quadtree) getNumberFromQuad(indexX, indexY, height, width int, current *node) int {
	var currentWidth int = (*current).width
	var currentHeight int = (*current).height
	if !configuration.Global.GenerationInfinie && !configuration.Global.TerreRonde {
		if indexX < 0 || indexY < 0 || indexX > width || indexY > height {
			return -1
		}
	}
	if configuration.Global.TerreRonde {
		if indexX < 0 || indexY < 0 || indexX >= width || indexY >= height {
			if indexY < 0 {
				indexY = (height - (-(indexY))%height)
			}
			if indexX < 0 {
				indexX = (indexX) + width
			}
			if indexY >= height {
				indexY = ((indexY) % height)
			}
			if indexX >= width {
				indexX = (indexX) % width
			}
			return q.getNumberFromQuad(indexX, indexY, height, width, current)
		}
	}
	if (*current).content != -1 {
		return (*current).content
	}
	if (currentWidth)%2 != 0 {
		currentWidth += 1
	}
	if (currentHeight)%2 != 0 {
		currentHeight += 1
	}
	var xChoice int = 0
	if (*current).TopLeftX+currentWidth/2 > indexX {
		xChoice = 1
	} else if (*current).TopLeftX+current.width > indexX || configuration.Global.GenerationInfinie {
		xChoice = 2
	}
	var yChoice int = 0
	if (*current).TopLeftY+currentHeight/2 > indexY {
		yChoice = 1
	} else if (*current).TopLeftY+(*current).height > indexY || configuration.Global.GenerationInfinie {
		yChoice = 2
	}
	var suivant *node
	var suivant1 node
	if xChoice == 1 && yChoice == 1 {
		if configuration.Global.GenerationInfinie && (*current).topLeftNode == nil {
			suivant1 = Recur("topLeft", [][]int{}, *current, (*current).TopLeftX, (*current).TopLeftY)
			(*current).topLeftNode = &suivant1
			suivant = &suivant1
		} else {
			suivant = (*current).topLeftNode
		}
	} else if xChoice == 2 && yChoice == 1 {
		if configuration.Global.GenerationInfinie && (*current).topRightNode == nil {
			suivant1 = Recur("topRight", [][]int{}, *current, (*current).TopLeftX, (*current).TopLeftY)
			(*current).topRightNode = &suivant1
			suivant = &suivant1
		} else {
			suivant = (*current).topRightNode
		}
	} else if xChoice == 1 && yChoice == 2 {
		if configuration.Global.GenerationInfinie && (*current).bottomLeftNode == nil {
			suivant1 = Recur("bottomLeft", [][]int{}, *current, (*current).TopLeftX, (*current).TopLeftY)
			(*current).bottomLeftNode = &suivant1
			suivant = &suivant1
		} else {
			suivant = (*current).bottomLeftNode
		}
	} else if xChoice == 2 && yChoice == 2 {
		if configuration.Global.GenerationInfinie && (*current).bottomRightNode == nil {
			suivant1 = Recur("bottomRight", [][]int{}, *current, (*current).TopLeftX, (*current).TopLeftY)
			(*current).bottomRightNode = &suivant1
			suivant = &suivant1
		} else {
			suivant = (*current).bottomRightNode
		}
	} else {
		return -1
	}
	return q.getNumberFromQuad(indexX, indexY, height, width, suivant)
}

func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) [][]int {
	contentHolder = [][]int{}
	var lenY, lenX int = configuration.Global.NumTileY, configuration.Global.NumTileX
	if configuration.Global.CameraFluide {
		lenY += 2
		lenX += 2
	}
	for i := 0; i < lenY; i++ {
		contentHolder = append(contentHolder, []int{})
		for x := 0; x < lenX; x++ {
			if configuration.Global.CameraFluide {
				contentHolder[i] = append(contentHolder[i], q.getNumberFromQuad(topLeftX+x-1, topLeftY+i-1, q.Height, q.Width, q.Root))
			} else {
				contentHolder[i] = append(contentHolder[i], q.getNumberFromQuad(topLeftX+x, topLeftY+i, q.Height, q.Width, q.Root))
			}
		}
	}
	return contentHolder
}
