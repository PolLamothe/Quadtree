package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	f.X = camXPos
	f.Y = camYPos
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(camXPos, camYPos)
	case fromFileFloor:
		f.updateFromFileFloor(camXPos, camYPos)
	case quadTreeFloor:
		f.updateQuadtreeFloor(camXPos, camYPos)
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(camXPos, camYPos int) {
	for y := 0; y < len(f.Content); y++ {
		for x := 0; x < len(f.Content[y]); x++ {
			absCamX := camXPos
			if absCamX < 0 {
				absCamX = -absCamX
			}
			absCamY := camYPos
			if absCamY < 0 {
				absCamY = -absCamY
			}
			f.Content[y][x] = ((x + absCamX%2) + (y + absCamY%2)) % 2
		}
	}
}

func decalRight(array *[][]int) {
	var result [][]int = [][]int{{}}
	for i := 0; i < len(*array); i++ {
		result = append(result, (*array)[i])
	}
	*array = result
}

func decalRight2(array *[]int) {
	var result []int = []int{0}
	for i := 0; i < len(*array); i++ {
		result = append(result, (*array)[i])
	}
	*array = result
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
func (f *Floor) updateFromFileFloor(camXPos, camYPos int) {
	var Ytile [][]int = f.FullContent
	var Yfill []int
	for i := 0; i < len(Ytile[0])+configuration.Global.NumTileX; i++ {
		Yfill = append(Yfill, -1)
	}
	for i := 0; i < configuration.Global.NumTileY; i++ {
		Ytile = append(Ytile, Yfill)
		decalRight(&Ytile)
		Ytile[0] = Yfill
	}
	for i := 0; i < len(Ytile); i++ {
		for x := 0; x < configuration.Global.NumTileX; x++ {
			Ytile[i] = append(Ytile[i], -1)
			decalRight2(&Ytile[i])
			Ytile[i][0] = -1
		}
	}
	var Xtile [][]int
	Xtile = Ytile[configuration.Global.NumTileY/2+camYPos+1 : configuration.Global.NumTileY/2+camYPos+configuration.Global.NumTileY+1]
	for i := 0; i < len(Xtile); i++ {
		Xtile[i] = Xtile[i][configuration.Global.NumTileX/2+camXPos+1 : configuration.Global.NumTileX/2+camXPos+configuration.Global.NumTileX+1]
	}
	f.Content = Xtile

}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	f.QuadtreeContent.GetContent(topLeftX, topLeftY, f.Content)
}
