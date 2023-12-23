package character

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/camera"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/multiplayer"
)

// Init met en place un personnage. Pour le moment
// cela consiste simplement à initialiser une variable
// responsable de définir l'étape d'animation courante.
func (c *Character) Init() {
	c.animationStep = 1

	if configuration.Global.CameraMode == camera.Static {
		if configuration.Global.MultiplayerKind == 0 {
			c.X = configuration.Global.ScreenCenterTileX
			c.Y = configuration.Global.ScreenCenterTileY
		} else if configuration.Global.MultiplayerKind == 1 {
			if c.CharacterNumber == 1 {
				c.X = configuration.Global.ScreenCenterTileX
				c.Y = configuration.Global.ScreenCenterTileY
				multiplayer.ServerPos["X"] = 0
				multiplayer.ServerPos["Y"] = 0
			}
			if c.CharacterNumber == 2 {
				if multiplayer.ServerPos["X"] == 0 {
					c.X = configuration.Global.ScreenCenterTileX + 1
				} else {
					c.X = configuration.Global.ScreenCenterTileX
				}
				if multiplayer.ServerPos["Y"] == 0 {
					c.Y = configuration.Global.ScreenCenterTileY + 1
				} else {
					c.Y = configuration.Global.ScreenCenterTileY
				}
				multiplayer.ClientPos["X"] = c.X
				multiplayer.ClientPos["Y"] = c.Y
			}
		} else {
			if c.CharacterNumber == 1 {
				c.X = configuration.Global.ScreenCenterTileX - multiplayer.ServerPos["X"]
				c.Y = configuration.Global.ScreenCenterTileY + multiplayer.ServerPos["Y"]
			} else {
				c.X = configuration.Global.ScreenCenterTileX - multiplayer.ClientPos["X"]
				c.Y = configuration.Global.ScreenCenterTileY + multiplayer.ClientPos["Y"]
			}
		}
	}
}
