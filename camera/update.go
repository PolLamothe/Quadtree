package camera

import "C"
import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Update met à jour la position de la caméra à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(characterPosX, characterPosY int, f *floor.Floor, q quadtree.Quadtree) {
	switch configuration.Global.CameraMode {
	case Static:
		c.updateStatic()
	case FollowCharacter:
		c.updateFollowCharacter(characterPosX, characterPosY, f, q)
	}
}

// updateStatic est la mise-à-jour d'une caméra qui reste
// toujours à la position (0,0). Cette fonction ne fait donc
// rien.
func (c *Camera) updateStatic() {}

// updateFollowCharacter est la mise-à-jour d'une caméra qui
// suit toujours le personnage. Elle prend en paramètres deux
// entiers qui indiquent les coordonnées du personnage et place
// la caméra au même endroit.
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY int, f *floor.Floor, q quadtree.Quadtree) {
	var MapWidth, MapHeight int
	MapHeight = q.Height
	MapWidth = q.Width
	if configuration.Global.CameraBlockEdge && !configuration.Global.TerreRonde && !configuration.Global.GenerationInfinie {
		if MapWidth >= configuration.Global.NumTileX && MapHeight >= configuration.Global.NumTileY {
			c.X = configuration.Global.NumTileX / 2
			c.Y = configuration.Global.NumTileY / 2
			(*f).AllBlockDisplayed = true
		}
	}
	if configuration.Global.CameraBlockEdge && (*f).AllBlockDisplayed && !configuration.Global.GenerationInfinie {
		if characterPosX-configuration.Global.NumTileX/2 >= 0 && characterPosX+configuration.Global.NumTileX/2 <= MapWidth {
			c.X = characterPosX
		}
		if characterPosY-configuration.Global.NumTileY/2 >= 0 && characterPosY+configuration.Global.NumTileY/2 <= MapHeight {
			c.Y = characterPosY
		}
	} else {
		if (characterPosX-configuration.Global.NumTileX/2 >= 0) && (characterPosY-configuration.Global.NumTileY/2 >= 0) {
			if characterPosX+configuration.Global.NumTileX/2 < MapWidth && characterPosY+configuration.Global.NumTileY/2 < MapHeight {
				(*f).AllBlockDisplayed = true
			}
		}
		c.X = characterPosX
		c.Y = characterPosY
	}

}
