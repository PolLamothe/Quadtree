package camera

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
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
	var camXExtern, camYExtern int = configuration.Global.NumTileX / 2, configuration.Global.NumTileY / 2
	if configuration.Global.NumTileX%2 == 0 {
		camXExtern--
	}
	if configuration.Global.NumTileY%2 == 0 {
		camYExtern--
	}
	if configuration.Global.CameraBlockEdge && !configuration.Global.TerreRonde && !configuration.Global.GenerationInfinie {
		if MapWidth >= configuration.Global.NumTileX {
			if characterPosX-configuration.Global.NumTileX/2 < 0 {
				c.X = float64(configuration.Global.NumTileX / 2)
			}
			if characterPosX+camXExtern >= MapWidth {
				c.X = float64(MapWidth-camXExtern) - 1
			}
		} else {
			configuration.Global.CameraBlockEdge = false
		}
		if MapHeight >= configuration.Global.NumTileY {
			if characterPosY-configuration.Global.NumTileY/2 < 0 {
				c.Y = float64(configuration.Global.NumTileY / 2)
			}
			if characterPosY+camYExtern >= MapHeight {
				c.Y = float64(MapHeight-camYExtern) - 1
			}
		} else {
			configuration.Global.CameraBlockEdge = false
		}
	}
	var XDir, YDir int = 0, 0
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
	if XDir == 1 {
		XDir--
	}
	if YDir == 1 {
		YDir--
	}
	if configuration.Global.CameraBlockEdge && MapWidth >= configuration.Global.NumTileX && MapHeight >= configuration.Global.NumTileY && !configuration.Global.GenerationInfinie {
		if float64(characterPosX-configuration.Global.NumTileX/2)+(float64(XShift)/float64(configuration.Global.TileSize)) >= 0 && float64(characterPosX+camXExtern)+(float64(XShift)/float64(configuration.Global.TileSize)) < float64(MapWidth) {
			if XShift == 0 {
				c.X = float64(characterPosX) + float64(XDir)
			}
			if configuration.Global.CameraFluide && float64(characterPosX+camXExtern+1)+float64(XDir) < float64(MapWidth) && float64(characterPosX-configuration.Global.NumTileX/2)+float64(XDir) >= 0 {
				c.X = float64(characterPosX) + float64(float64(XShift)/float64(configuration.Global.TileSize))
			}
		}
		if float64(characterPosY-configuration.Global.NumTileY/2)+(float64(YShift)/float64(configuration.Global.TileSize)) >= 0 && float64(characterPosY+camYExtern)+(float64(YShift)/float64(configuration.Global.TileSize)) < float64(MapHeight) {
			if YShift == 0 {
				c.Y = float64(characterPosY) + float64(YDir)
			}
			if configuration.Global.CameraFluide && float64(characterPosY-configuration.Global.NumTileY/2)+float64(YDir) >= 0 && float64(characterPosY+camYExtern+1)+float64(YDir) < float64(MapHeight) {
				c.Y = float64(characterPosY) + float64(float64(YShift)/float64(configuration.Global.TileSize))
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
