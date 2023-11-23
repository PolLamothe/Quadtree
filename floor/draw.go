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
func (f Floor) Draw(screen *ebiten.Image) {

	for y := range f.Content {
		for x := range f.Content[y] {
			if f.Content[y][x] >= 0 && f.Content[y][x] <= 5 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize), float64(y*configuration.Global.TileSize))

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
