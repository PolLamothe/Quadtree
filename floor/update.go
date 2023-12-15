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

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
func (f *Floor) updateFromFileFloor(camXPos, camYPos int) {
	var result [][]int
	var lenY, lenX int = configuration.Global.NumTileY, configuration.Global.NumTileX
	for i := 0; i < lenY; i++ {
		result = append(result, []int{})
		for x := 0; x < lenX; x++ {
			var indexX, indexY int = camXPos - configuration.Global.NumTileX/2 + x, camYPos - configuration.Global.NumTileY/2 + i
			if indexX < 0 || indexY < 0 || indexX >= len(f.FullContent[0]) || indexY >= len(f.FullContent) { // si les coordonnées sont en dehors du terrain
				result[i] = append(result[i], -1)
			} else {
				result[i] = append(result[i], f.FullContent[indexY][indexX])
			}
		}
	}
	f.Content = result

}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	f.QuadtreeContent.GetContent(topLeftX, topLeftY, f.Content)
}
