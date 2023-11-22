package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/portal"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool) {

	if !c.moving {
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			c.orientation = orientedRight
			if !blocking[1] {
				c.xInc = 1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			c.orientation = orientedLeft
			if !blocking[3] {
				c.xInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			c.orientation = orientedUp
			if !blocking[0] {
				c.yInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			c.orientation = orientedDown
			if !blocking[2] {
				c.yInc = 1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyTab) && !portal.IsPortalHere(c.X, c.Y) {
			if len(portal.PortalStore) == 2 {
				portal.PortalStore = portal.PortalStore[1:]
			}
			portal.PortalStore = append(portal.PortalStore, []int{c.X, c.Y})
		}
	} else {
		c.animationFrameCount++
		if c.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			c.shift += shiftStep
			c.animationStep = -c.animationStep
			if c.shift > configuration.Global.TileSize-shiftStep {
				c.shift = 0
				c.moving = false
				c.X += c.xInc
				c.Y += c.yInc
				c.xInc = 0
				c.yInc = 0
				if portal.IsPortalHere(c.X, c.Y) && len(portal.PortalStore) == 2 {
					var newCoord []int = portal.GetOtherCoordonate(c.X, c.Y)
					c.X = newCoord[0]
					c.Y = newCoord[1]
					if configuration.Global.SingleUsagePortal {
						portal.PortalStore = [][]int{}
					}
					c.Update(blocking)
				}
			}
		}
	}

}
