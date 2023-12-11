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
	if configuration.Global.CameraBlockEdge && !configuration.Global.GenerationInfinie {
		if (f.X == configuration.Global.NumTileX/2 && XCharacter+futureX <= configuration.Global.NumTileX/2) || XCam != XCharacter || (f.X == f.QuadtreeContent.Width-configuration.Global.NumTileX/2 && XCharacter+futureX >= f.QuadtreeContent.Width-configuration.Global.NumTileX/2) {
			XShift = 0
		}
		if (f.Y == configuration.Global.NumTileY/2 && YCharacter+futureY <= configuration.Global.NumTileY/2) || YCam != YCharacter || (f.Y == f.QuadtreeContent.Height-configuration.Global.NumTileY/2 && YCharacter+futureY >= f.QuadtreeContent.Height-configuration.Global.NumTileY/2) {
			YShift = 0
		}
	}
	if !configuration.Global.CameraFluide {
		YShift = 0
		XShift = 0
	}
	for y := range f.Content {
		for x := range f.Content[y] {
			if f.Content[y][x] >= 0 && f.Content[y][x] <= 5 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize)-float64(XShift), float64(y*configuration.Global.TileSize)-float64(YShift)) //fonction qui permet de deplacer l'image du sol
				shiftX := f.Content[y][x] * configuration.Global.TileSize
				screen.DrawImage(assets.FloorImage.SubImage(
					image.Rect(shiftX, 0, shiftX+configuration.Global.TileSize, configuration.Global.TileSize),
				).(*ebiten.Image), op)
				if portal.IsPortalHere(f.X-len(f.Content[0])/2+x, f.Y-len(f.Content)/2+y) {
					screen.DrawImage(assets.PortalIMG, op)
				}
			}
		}
	}

}
