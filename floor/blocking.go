package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {
	var relativeXPos, relativeYPos int
	if configuration.Global.MultiplayerKind == 0 {
		if configuration.Global.CameraFluide {
			relativeXPos = (characterXPos - camXPos) + configuration.Global.ScreenCenterTileX + 1
			relativeYPos = (characterYPos - camYPos) + configuration.Global.ScreenCenterTileY + 1
			blocking[0] = relativeYPos <= 0 || f.Content[relativeYPos-1][relativeXPos] == -1
			blocking[1] = relativeXPos-1 >= configuration.Global.NumTileX-1 || f.Content[relativeYPos][relativeXPos+1] == -1
			blocking[2] = relativeYPos-1 >= configuration.Global.NumTileY-1 || f.Content[relativeYPos+1][relativeXPos] == -1
			blocking[3] = relativeXPos <= 0 || f.Content[relativeYPos][relativeXPos-1] == -1
		} else {
			relativeXPos = characterXPos - camXPos + configuration.Global.ScreenCenterTileX
			relativeYPos = characterYPos - camYPos + configuration.Global.ScreenCenterTileY
			blocking[0] = relativeYPos <= 0 || f.Content[relativeYPos-1][relativeXPos] == -1
			blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.Content[relativeYPos][relativeXPos+1] == -1
			blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.Content[relativeYPos+1][relativeXPos] == -1
			blocking[3] = relativeXPos <= 0 || f.Content[relativeYPos][relativeXPos-1] == -1
		}
	} else {
		var topLeftY, topLeftX int = characterYPos - configuration.Global.NumTileY/2, characterXPos - configuration.Global.NumTileX/2
		var mapContent [][]int = f.QuadtreeContent.GetContent(topLeftX, topLeftY, f.Content)
		relativeXPos = configuration.Global.ScreenCenterTileX
		relativeYPos = configuration.Global.ScreenCenterTileY
		if configuration.Global.CameraFluide {
			relativeXPos++
			relativeYPos++
		}
		blocking[0] = mapContent[relativeYPos-1][relativeXPos] == -1
		blocking[1] = mapContent[relativeYPos][relativeXPos+1] == -1
		blocking[2] = mapContent[relativeYPos+1][relativeXPos] == -1
		blocking[3] = mapContent[relativeYPos][relativeXPos-1] == -1
	}
	return blocking
}
