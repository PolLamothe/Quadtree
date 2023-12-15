package floor

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/portal"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (f Floor) Draw(screen *ebiten.Image, XShift, YShift, XCam, YCam, XCharacter, YCharacter int) {
	var futureX, futureY int = 0, 0
	if XShift > 0 {
		futureX = 1
	} else if XShift < 0 {
		futureX = -1
	}
	if YShift > 0 {
		futureY = 1
	} else if YShift < 0 {
		futureY = -1
	}
	if configuration.Global.CameraBlockEdge && !configuration.Global.GenerationInfinie && f.AllBlockDisplayed {
		if !configuration.Global.CameraFluide {
			if (f.X == configuration.Global.NumTileX/2 && XCharacter+futureX <= configuration.Global.NumTileX/2) || XCam != XCharacter || (f.X == f.QuadtreeContent.Width-configuration.Global.NumTileX/2 && XCharacter+futureX >= f.QuadtreeContent.Width-configuration.Global.NumTileX/2) {
				XShift = 0
			}
			if (f.Y == configuration.Global.NumTileY/2 && YCharacter+futureY <= configuration.Global.NumTileY/2) || YCam != YCharacter || (f.Y == f.QuadtreeContent.Height-configuration.Global.NumTileY/2 && YCharacter+futureY >= f.QuadtreeContent.Height-configuration.Global.NumTileY/2) {
				YShift = 0
			}
		} else {
			if !configuration.Global.CameraBlockEdge {
				if (f.X+1 == configuration.Global.NumTileX/2 && XCharacter+futureX <= configuration.Global.NumTileX/2) || XCam != XCharacter || (f.X+1 == f.QuadtreeContent.Width-configuration.Global.NumTileX/2 && XCharacter+futureX >= f.QuadtreeContent.Width-configuration.Global.NumTileX/2) {
					XShift = 0
				}
				if (f.Y+1 == configuration.Global.NumTileY/2 && YCharacter+futureY <= configuration.Global.NumTileY/2) || YCam != YCharacter || (f.Y+1 == f.QuadtreeContent.Height-configuration.Global.NumTileY/2 && YCharacter+futureY >= f.QuadtreeContent.Height-configuration.Global.NumTileY/2) {
					YShift = 0
				}
				if (f.X == configuration.Global.NumTileX/2 && XCharacter+futureX <= configuration.Global.NumTileX/2) || XCam != XCharacter || (f.X == f.QuadtreeContent.Width-configuration.Global.NumTileX/2 && XCharacter+futureX >= f.QuadtreeContent.Width-configuration.Global.NumTileX/2) {
					XShift = 0
				}
				if (f.Y == configuration.Global.NumTileY/2 && YCharacter+futureY <= configuration.Global.NumTileY/2) || YCam != YCharacter || (f.Y == f.QuadtreeContent.Height-configuration.Global.NumTileY/2 && YCharacter+futureY >= f.QuadtreeContent.Height-configuration.Global.NumTileY/2) {
					YShift = 0
				}
			} else {
				var XFutureCharacter, YFutureCharacter int = XCharacter + futureX, YCharacter + futureY
				if configuration.Global.NumTileX%2 != 0 {
					XFutureCharacter++
				}
				if configuration.Global.NumTileY%2 != 0 {
					YFutureCharacter++
				}
				var XFloor, YFloor int = f.X, f.Y
				if configuration.Global.NumTileX%2 != 0 {
					XFloor++
				}
				if configuration.Global.NumTileY%2 != 0 {
					YFloor++
				}
				if (XFloor-1 == configuration.Global.NumTileX/2 && XFutureCharacter <= configuration.Global.NumTileX/2) || XCam != XCharacter || (XFloor == f.QuadtreeContent.Width-configuration.Global.NumTileX/2 && XCharacter+futureX >= f.QuadtreeContent.Width-configuration.Global.NumTileX/2) {
					XShift = 0
				}
				if (YFloor-1 == configuration.Global.NumTileY/2 && YFutureCharacter <= configuration.Global.NumTileY/2) || YCam != YCharacter || (YFloor == f.QuadtreeContent.Height-configuration.Global.NumTileY/2 && YCharacter+futureY >= f.QuadtreeContent.Height-configuration.Global.NumTileY/2) {
					YShift = 0
				}
				if (XFloor-1 == configuration.Global.NumTileX/2 && XFutureCharacter <= configuration.Global.NumTileX/2) || XCam != XCharacter || (XFloor == f.QuadtreeContent.Width-configuration.Global.NumTileX/2 && XCharacter+futureX >= f.QuadtreeContent.Width-configuration.Global.NumTileX/2) {
					XShift = 0
				}
				if (YFloor-1 == configuration.Global.NumTileY/2 && YFutureCharacter <= configuration.Global.NumTileY/2) || YCam != YCharacter || (YFloor == f.QuadtreeContent.Height-configuration.Global.NumTileY/2 && YCharacter+futureY >= f.QuadtreeContent.Height-configuration.Global.NumTileY/2) {
					YShift = 0
				}
			}
		}
	}
	if !configuration.Global.CameraFluide {
		YShift = 0
		XShift = 0
	}
	for y := 0; y < len(f.Content); y++ {
		for x := 0; x < len(f.Content[y]); x++ {
			if f.Content[y][x] >= 0 && f.Content[y][x] <= 5 {
				op := &ebiten.DrawImageOptions{}
				if !configuration.Global.CameraFluide {
					op.GeoM.Translate(float64(x*configuration.Global.TileSize)-float64(XShift), float64(y*configuration.Global.TileSize)-float64(YShift)) //fonction qui permet de deplacer l'image du sol
				} else {
					op.GeoM.Translate(float64((x-1)*configuration.Global.TileSize)-float64(XShift), float64((y-1)*configuration.Global.TileSize)-float64(YShift)) //fonction qui permet de deplacer l'image du sol
				}
				shiftX := f.Content[y][x] * configuration.Global.TileSize
				screen.DrawImage(assets.FloorImage.SubImage( // fonction qui sert a afficher la case du sol correspondante au numéro
					image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
				).(*ebiten.Image), op)

				if portal.IsPortalHere(f.X-len(f.Content[0])/2+x, f.Y-len(f.Content)/2+y) {
					screen.DrawImage(assets.PortalIMG, op)
				}
				if configuration.Global.TerreRonde && configuration.Global.Portal {
					var XIndex, YIndex int = f.X - len(f.Content[0])/2 + x, f.Y - len(f.Content)/2 + y // position de la case qui est affichée
					//on detecte la vrai position des case hors-champ
					if XIndex >= 0 {
						XIndex = XIndex % f.QuadtreeContent.Width
					} else {
						XIndex = (f.QuadtreeContent.Width + (XIndex % f.QuadtreeContent.Width)) % f.QuadtreeContent.Width
					}
					if YIndex >= 0 {
						YIndex = YIndex % f.QuadtreeContent.Height
					} else {
						YIndex = (f.QuadtreeContent.Height + (YIndex % f.QuadtreeContent.Height)) % f.QuadtreeContent.Height
					}
					if portal.IsPortalHere(XIndex, YIndex) {
						screen.DrawImage(assets.PortalIMG, op)
					}
				}
			}
		}
	}

}
