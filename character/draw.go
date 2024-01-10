package character

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw permet d'afficher le personnage dans une *ebiten.Image
// (en pratique, celle qui représente la fenêtre de jeu) en
// fonction des charactéristiques du personnage (position, orientation,
// étape d'animation, etc) et de la position de la caméra (le personnage
// est affiché relativement à la caméra).
func (c *Character) Draw(screen *ebiten.Image, MapWidth, MapHeight int, camX, camY float64, XShift, YShift int) {
	xShift := 0
	yShift := 0
	var orientation string
	switch c.orientation {
	case orientedDown:
		yShift = c.shift
		orientation = "Y"
	case orientedUp:
		yShift = -c.shift
		orientation = "Y"
	case orientedLeft:
		xShift = -c.shift
		orientation = "X"
	case orientedRight:
		xShift = c.shift
		orientation = "X"
	}
	var camX2, camY2 int = int(camX), int(camY)

	var numTileXHalf, numTileYHalf int = configuration.Global.NumTileX / 2, configuration.Global.NumTileY / 2

	if configuration.Global.NumTileX%2 != 0 {
		numTileXHalf++
	}
	if configuration.Global.NumTileY%2 != 0 {
		numTileYHalf++
	}

	if configuration.Global.CameraBlockEdge && configuration.Global.CameraFluide {
		if xShift < 0 && camX2+1 == c.X && camX2+numTileXHalf < MapWidth {
			camX2++
		}
		if yShift < 0 && camY2+1 == c.Y && camY2+numTileYHalf < MapHeight {
			camY2++
		}
	}

	xTileForDisplay := c.X - camX2 + configuration.Global.ScreenCenterTileX
	yTileForDisplay := c.Y - camY2 + configuration.Global.ScreenCenterTileY
	xPos := (xTileForDisplay)*configuration.Global.TileSize + XShift
	yPos := (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + YShift

	var futureX, futureY int = 0, 0
	if xShift > 0 {
		futureX = 1
	} else if xShift < 0 {
		futureX = -1
	}
	if yShift > 0 {
		futureY = 1
	} else if yShift < 0 {
		futureY = -1
	}
	if configuration.Global.CameraFluide && (configuration.Global.MultiplayerKind == 0 || configuration.Global.MultiplayerKind == c.CharacterNumber) {
		var camXExtern, camYExtern int = camX2, camY2
		if configuration.Global.NumTileX%2 != 0 { //si jamais le nombre de tuile est pair, il y'aura plus de tuile d'un coté que de l'autre donc il faut rééquilibrer cela
			camXExtern++
		}
		if configuration.Global.NumTileY%2 != 0 {
			camYExtern++
		}
		if configuration.Global.CameraBlockEdge && ((((camX2 == configuration.Global.NumTileX/2 && c.X+futureX <= camX2) || (camXExtern == MapWidth-configuration.Global.NumTileX/2 && c.X+futureX >= camX2)) && orientation == "X") || (((camY2 == configuration.Global.NumTileY/2 && c.Y+futureY <= camY2) || (camYExtern == MapHeight-configuration.Global.NumTileY/2 && c.Y+futureY >= camY2)) && orientation == "Y")) && MapWidth >= configuration.Global.NumTileX && MapHeight >= configuration.Global.NumTileY { //condition a remplir pour que le personnage bouge visuellement
			xPos = (xTileForDisplay)*configuration.Global.TileSize + xShift
			yPos = (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + yShift
		} else {
			if configuration.Global.CameraBlockEdge {
				if configuration.Global.NumTileX > MapWidth {
					if xShift < 0 {
						xTileForDisplay--
					}
				}
				if configuration.Global.NumTileY > MapHeight {
					if yShift < 0 {
						yTileForDisplay--
					}
				}
				if camX2 == c.X {
					xPos = ((configuration.Global.ScreenCenterTileX) * configuration.Global.TileSize)
				} else {
					xPos = xTileForDisplay * configuration.Global.TileSize
				}
				if camY2 == c.Y {
					yPos = ((configuration.Global.ScreenCenterTileY)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2)
				} else {
					yPos = yTileForDisplay*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2
				}
			} else {
				if yShift < 0 && camY >= 0 {
					yTileForDisplay--
				}
				if xShift < 0 && camX >= 0 {
					xTileForDisplay--
				}
				if configuration.Global.GenerationInfinie {
					if xShift > 0 && camX < 0 {
						xTileForDisplay++
					}
					if yShift > 0 && camY < 0 {
						yTileForDisplay++
					}
				}
				xPos = xTileForDisplay * configuration.Global.TileSize
				yPos = yTileForDisplay*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2
			}
		}
	}
	if configuration.Global.MultiplayerKind != 0 && configuration.Global.MultiplayerKind != c.CharacterNumber { // on dessine l'autre joueur
		var XDiff, YDiff float64 = camX - float64(int(camX)), camY - float64(int(camY))
		if xShift < 0 && !(c.XShift < 0 && XShift < 0) {
			XDiff = 1 - XDiff
		}
		if yShift < 0 && !(c.YShift < 0 && YShift < 0) {
			YDiff = 1 - YDiff
		}
		if !configuration.Global.CameraFluide {
			xPos = (xTileForDisplay)*configuration.Global.TileSize + c.XShift
			yPos = (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 + c.YShift
		} else {
			if c.XShift < 0 {
				xTileForDisplay += 2
				if configuration.Global.CameraBlockEdge && ((c.X+c.xInc != int(camX) || c.orientation != orientedLeft || c.X == camX2+1) || MapWidth < configuration.Global.NumTileX) {
					xTileForDisplay--
				}
				if XShift < 0 && c.XShift < 0 {
					xTileForDisplay--
				}
				if !configuration.Global.CameraBlockEdge {
					xTileForDisplay--
				}
			}
			if c.YShift < 0 {
				yTileForDisplay += 2
				if configuration.Global.CameraBlockEdge && ((c.Y+c.yInc != int(camY) || c.orientation != orientedUp || c.Y == MapHeight-2 || c.Y == camY2+1) || MapHeight < configuration.Global.NumTileY) {
					yTileForDisplay--
				}
				if YShift < 0 && c.YShift < 0 {
					yTileForDisplay--
				}
				if !configuration.Global.CameraBlockEdge {
					yTileForDisplay--
				}
			}
			xPos = (xTileForDisplay)*configuration.Global.TileSize - int(XDiff*float64(configuration.Global.TileSize)) + c.XShift
			yPos = (yTileForDisplay)*configuration.Global.TileSize - configuration.Global.TileSize/2 + 2 - int(YDiff*float64(configuration.Global.TileSize)) + c.YShift
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xPos), float64(yPos))

	shiftX := configuration.Global.TileSize
	if c.moving {
		shiftX += c.animationStep * configuration.Global.TileSize
	}
	shiftY := c.orientation * configuration.Global.TileSize

	screen.DrawImage(assets.CharacterImage.SubImage(
		image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
	).(*ebiten.Image), op)
}
