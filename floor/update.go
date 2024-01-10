package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/multiplayer"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos, XShift, YShift int) {
	f.X = camXPos
	f.Y = camYPos
	switch configuration.Global.FloorKind {
	case gridFloor:
		f.updateGridFloor(camXPos, camYPos)
	case fromFileFloor:
		f.updateFromFileFloor(camXPos, camYPos)
	case quadTreeFloor:
		f.updateQuadtreeFloor(camXPos, camYPos, XShift, YShift)
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
	if !configuration.Global.TerreRonde {
		var result [][]int
		var lenY, lenX int = configuration.Global.NumTileY, configuration.Global.NumTileX
		if configuration.Global.CameraFluide {
			lenY += 2
			lenX += 2
		}
		for i := 0; i < lenY; i++ {
			result = append(result, []int{})
			for x := 0; x < lenX; x++ {
				var indexX, indexY int = camXPos - configuration.Global.NumTileX/2 + x, camYPos - configuration.Global.NumTileY/2 + i
				if configuration.Global.CameraFluide {
					indexX -= 1
					indexY -= 1
				}
				if indexX < 0 || indexY < 0 || indexX >= len(f.FullContent[0]) || indexY >= len(f.FullContent) {
					result[i] = append(result[i], -1)
				} else {
					result[i] = append(result[i], f.FullContent[indexY][indexX])
				}
			}
		}
		f.Content = result
	} else {
		var result [][]int
		var lenY, lenX int = configuration.Global.NumTileY, configuration.Global.NumTileX
		if configuration.Global.CameraFluide {
			lenY += 2
			lenX += 2
		}
		for i := 0; i < lenY; i++ {
			result = append(result, []int{})
			for x := 0; x < lenX; x++ {
				var indexX, indexY int = camXPos - configuration.Global.NumTileX/2 + x, camYPos - configuration.Global.NumTileY/2 + i
				if configuration.Global.CameraFluide {
					indexX -= 1
					indexY -= 1
				}
				if indexX < 0 || indexY < 0 || indexX >= len(f.FullContent[0]) || indexY >= len(f.FullContent) {
					if indexY < 0 {
						indexY = (len(f.FullContent) - (-(indexY))%len(f.FullContent))
					}
					if indexX < 0 {
						indexX = (indexX) + len(f.FullContent[0])
					}
					if indexY >= len(f.FullContent) {
						indexY = ((indexY) % len(f.FullContent))
					}
					if indexX >= len(f.FullContent[0]) {
						indexX = (indexX) % len(f.FullContent[0])
					}
				}
				result[i] = append(result[i], f.FullContent[indexY][indexX])
			}
		}
		f.Content = result
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(camXPos, camYPos, XShift, YShift int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	if configuration.Global.GenerationInfinie { //si la génération infinie est activé on détecte si le joueur demande a charger des terrains pas enore généré, et si oui, on les génère et on detecte quelle node devra être le terrain actuelle
		// (ex: si le joueur demande a générer du terrain citué a gauche, on va cituer le terrain actuelle a droite, et on va générer le terrain demandé a gauche)
		if configuration.Global.CameraFluide {
			if XShift < 0 {
				topLeftX--
			}
			if YShift < 0 {
				topLeftY--
			}
		}
		if topLeftX < f.QuadtreeContent.Root.TopLeftX { //si le joueur demande a charger du terrain a gauche
			if camYPos <= 0 { //droite-Haut
				f.QuadtreeContent.GenerateInfinite("TopRight")
			} else { //droite-bas
				f.QuadtreeContent.GenerateInfinite("BottomRight")
			}
		} else if topLeftY < f.QuadtreeContent.Root.TopLeftY { //si le joueur demande a charger du terrain en haut
			if camXPos > 0 { //bas-Gauche
				f.QuadtreeContent.GenerateInfinite("BottomLeft")
			} else { //bas-Droite
				f.QuadtreeContent.GenerateInfinite("BottomRight")
			}
		} else if topLeftX+configuration.Global.NumTileX >= f.QuadtreeContent.Width/2 { //si le joueur demande a charger du terrain a droite
			if camYPos >= 0 { //haut-gauche
				f.QuadtreeContent.GenerateInfinite("TopLeft")
			} else { //Bas-Gauche
				f.QuadtreeContent.GenerateInfinite("BottomLeft")
			}
		} else if topLeftY+configuration.Global.NumTileY >= f.QuadtreeContent.Height/2 { //si le joueur demande a charger du terrain en bas
			if camXPos >= 0 { //Haut-Gauche
				f.QuadtreeContent.GenerateInfinite("TopLeft")
			} else { //haut-Droite
				f.QuadtreeContent.GenerateInfinite("TopRight")
			}
		}
		if configuration.Global.CameraFluide {
			if XShift < 0 {
				topLeftX++
			}
			if YShift < 0 {
				topLeftY++
			}
		}
	}
	f.Content = f.QuadtreeContent.GetContent(topLeftX, topLeftY, f.Content, true)
	if multiplayer.RoutineFinished && len(multiplayer.BlockToSend) > 0 {
		multiplayer.SendBlock()
	}
}
