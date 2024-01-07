package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/multiplayer"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/portal"
)

//variable qui sert a temporiser le mouvement a la sortie d'un portail pour empecher le personnage d'aller sur une case -1

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool, f *floor.Floor) {
	if c.CharacterNumber == 1 {
		multiplayer.ServerPos = map[string]int{"X": c.X, "Y": c.Y}
	}
	if c.CharacterNumber == 2 {
		multiplayer.ClientPos = map[string]int{"X": c.X, "Y": c.Y}
	}
	if !c.moving && !c.PortalSecure {
		if (ebiten.IsKeyPressed(ebiten.KeyRight) && (configuration.Global.MultiplayerKind == 0 || configuration.Global.MultiplayerKind == c.CharacterNumber)) || (multiplayer.KeyPressed == "right" && configuration.Global.MultiplayerKind != c.CharacterNumber) {
			c.orientation = orientedRight
			if !blocking[1] {
				if configuration.Global.MultiplayerKind != c.CharacterNumber {
					multiplayer.KeyPressed = ""
				}
				c.xInc = 1
				c.moving = true
				if configuration.Global.MultiplayerKind != 0 && c.CharacterNumber == configuration.Global.MultiplayerKind {
					multiplayer.SendKeyPressed("right")
				}
			}
		} else if (ebiten.IsKeyPressed(ebiten.KeyLeft) && (configuration.Global.MultiplayerKind == 0 || configuration.Global.MultiplayerKind == c.CharacterNumber)) || (multiplayer.KeyPressed == "left" && configuration.Global.MultiplayerKind != c.CharacterNumber) {
			c.orientation = orientedLeft
			if !blocking[3] {
				if configuration.Global.MultiplayerKind != c.CharacterNumber {
					multiplayer.KeyPressed = ""
				}
				c.xInc = -1
				c.moving = true
				if configuration.Global.MultiplayerKind != 0 && c.CharacterNumber == configuration.Global.MultiplayerKind {
					multiplayer.SendKeyPressed("left")
				}
			}
		} else if (ebiten.IsKeyPressed(ebiten.KeyUp) && (configuration.Global.MultiplayerKind == 0 || configuration.Global.MultiplayerKind == c.CharacterNumber)) || (multiplayer.KeyPressed == "up" && configuration.Global.MultiplayerKind != c.CharacterNumber) {
			c.orientation = orientedUp
			if !blocking[0] {
				if configuration.Global.MultiplayerKind != c.CharacterNumber {
					multiplayer.KeyPressed = ""
				}
				c.yInc = -1
				c.moving = true
				if configuration.Global.MultiplayerKind != 0 && c.CharacterNumber == configuration.Global.MultiplayerKind {
					multiplayer.SendKeyPressed("up")
				}
			}
		} else if (ebiten.IsKeyPressed(ebiten.KeyDown) && (configuration.Global.MultiplayerKind == 0 || configuration.Global.MultiplayerKind == c.CharacterNumber)) || (multiplayer.KeyPressed == "down" && configuration.Global.MultiplayerKind != c.CharacterNumber) {
			c.orientation = orientedDown
			if !blocking[2] {
				if configuration.Global.MultiplayerKind != c.CharacterNumber {
					multiplayer.KeyPressed = ""
				}
				c.yInc = 1
				c.moving = true
				if configuration.Global.MultiplayerKind != 0 && c.CharacterNumber == configuration.Global.MultiplayerKind {
					multiplayer.SendKeyPressed("down")
				}
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyTab) && !portal.IsPortalHere(c.X, c.Y) && configuration.Global.Portal && (configuration.Global.MultiplayerKind == 0 || configuration.Global.MultiplayerKind == c.CharacterNumber) {
			if len(portal.PortalStore) == 2 {
				portal.PortalStore = portal.PortalStore[1:]
			}
			portal.PortalStore = append(portal.PortalStore, []int{c.X, c.Y})
			if configuration.Global.MultiplayerKind != 0 {
				multiplayer.SendKeyPressed("tab")
			}
		} else if configuration.Global.MultiplayerKind != 0 && c.CharacterNumber != configuration.Global.MultiplayerKind && multiplayer.KeyPressed == "tab" {
			if len(multiplayer.MultiplayerPortal) == 2 {
				multiplayer.MultiplayerPortal = multiplayer.MultiplayerPortal[1:]
			}
			multiplayer.MultiplayerPortal = append(multiplayer.MultiplayerPortal, []int{c.X, c.Y})
		}
	} else {
		c.animationFrameCount++
		c.PortalSecure = false
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
				if configuration.Global.TerreRonde && !configuration.Global.GenerationInfinie {
					if c.X < 0 {
						c.X = f.QuadtreeContent.Width + c.X
					}
					if c.Y < 0 {
						c.Y = f.QuadtreeContent.Height + c.Y
					}
					if c.X >= f.QuadtreeContent.Width {
						c.X = c.X % f.QuadtreeContent.Width
					}
					if c.Y >= f.QuadtreeContent.Height {
						c.Y = c.Y % f.QuadtreeContent.Height
					}
				}
				if portal.IsPortalHere(c.X, c.Y) {
					if portal.IsInLocalPortalStore(c.X, c.Y) {
						if len(portal.PortalStore) == 2 {
							var newCoord []int = portal.GetOtherCoordonate(c.X, c.Y)
							c.X = newCoord[0]
							c.Y = newCoord[1]
							c.xInc, c.yInc = 0, 0
							if configuration.Global.SingleUsagePortal {
								portal.PortalStore = [][]int{}
							}
							c.PortalSecure = true
							c.Update(blocking, f)
						}
					} else {
						if len(multiplayer.MultiplayerPortal) == 2 {
							var newCoord []int = portal.GetOtherCoordonate(c.X, c.Y)
							c.X = newCoord[0]
							c.Y = newCoord[1]
							c.xInc, c.yInc = 0, 0
							if configuration.Global.SingleUsagePortal {
								multiplayer.MultiplayerPortal = [][]int{}
							}
							c.PortalSecure = true
							c.Update(blocking, f)
						}
					}
				}
			}
		}
	}
}
