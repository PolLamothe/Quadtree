package camera

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init met en place une caméra.
func (c *Camera) Init(MapWidth, MapHeight int, AllBlockDisplayed *bool) {
	if configuration.Global.CameraMode == Static {
		c.X = configuration.Global.ScreenCenterTileX
		c.Y = configuration.Global.ScreenCenterTileY
	}
}
