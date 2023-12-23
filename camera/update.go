package camera

import "C"
import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
	"math"
)

// Update met à jour la position de la caméra à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(characterPosX, characterPosY int, f *floor.Floor, q quadtree.Quadtree, XShift, YShift int) {
	switch configuration.Global.CameraMode {
	case Static:
		c.updateStatic()
	case FollowCharacter:
		c.updateFollowCharacter(characterPosX, characterPosY, f, q, XShift, YShift)
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
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY int, f *floor.Floor, q quadtree.Quadtree, XShift, YShift int) {
	var MapWidth, MapHeight int
	MapHeight = q.Height
	MapWidth = q.Width
	if configuration.Global.CameraBlockEdge && !configuration.Global.TerreRonde && !configuration.Global.GenerationInfinie {
		if MapWidth >= configuration.Global.NumTileX && MapHeight >= configuration.Global.NumTileY && characterPosX < configuration.Global.NumTileX/2 && characterPosY < configuration.Global.NumTileY/2 {
			c.X = float64(configuration.Global.NumTileX / 2)
			c.Y = float64(configuration.Global.NumTileY / 2)
			(*f).AllBlockDisplayed = true
		}
	}
	var XDir, YDir int = 0, 0
	var XState, YState bool = false, false
	if XShift > 0 {
		XDir = 1
	} else if XShift < 0 {
		XDir = -1
	}
	if YShift > 0 {
		YDir = 1
	} else if YShift < 0 {
		YDir = -1
	}
	if configuration.Global.NumTileX%2 == 0 && XDir == 1 {
		XState = true
		XDir--
	}
	if configuration.Global.NumTileY%2 == 0 && YDir == 1 {
		YState = true
		YDir--
	}

	if configuration.Global.CameraBlockEdge && (*f).AllBlockDisplayed && !configuration.Global.GenerationInfinie {
		var cameraX, cameraY int = int(c.X), int(c.Y)
		if configuration.Global.CameraFluide {
			if XDir == -1 && (c.X)-math.Floor(c.X) != 0 {
				cameraX = int(c.X) + 1
			}
			if YDir == -1 && (c.Y)-math.Floor(c.Y) != 0 {
				cameraY = int(c.Y) + 1
			}
		}
		if (cameraX) == characterPosX || cameraY == characterPosY || (c.X)-math.Floor(c.X) != 0 || (c.Y)-math.Floor(c.Y) != 0 {
			if cameraX-configuration.Global.NumTileX/2+XDir >= 0 && cameraX+configuration.Global.NumTileX/2+XDir < MapWidth && (cameraX == characterPosX || (c.X)-math.Floor(c.X) != 0) {
				c.X = float64(characterPosX) + float64(XDir)
				if configuration.Global.CameraFluide {
					c.X = float64(characterPosX) + float64(float64(XShift)/float64(configuration.Global.TileSize))
				} else if configuration.Global.NumTileX%2 == 0 && XState {
					c.X++
				}
			}
			if cameraY-configuration.Global.NumTileY/2+YDir >= 0 && cameraY+configuration.Global.NumTileY/2+YDir < MapHeight && (cameraY == characterPosY || (c.Y)-math.Floor(c.Y) != 0) {
				c.Y = float64(characterPosY) + float64(YDir)
				if configuration.Global.CameraFluide {
					c.Y = float64(characterPosY) + float64(float64(YShift)/float64(configuration.Global.TileSize))
				} else if configuration.Global.NumTileY%2 == 0 && YState {
					c.Y++
				}
			}
		}

	} else {
		if configuration.Global.CameraFluide {
			c.X = float64(characterPosX) + float64(float64(XShift)/float64(configuration.Global.TileSize))
			c.Y = float64(characterPosY) + float64(float64(YShift)/float64(configuration.Global.TileSize))
		} else {
			c.X = float64(characterPosX)
			c.Y = float64(characterPosY)
		}
	}
}
