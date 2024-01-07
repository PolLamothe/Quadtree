package camera

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init met en place une caméra.
func (c *Camera) Init() {
	if configuration.Global.CameraMode == Static {
		c.X = float64(configuration.Global.ScreenCenterTileX)
		c.Y = float64(configuration.Global.ScreenCenterTileY)
	}
}
