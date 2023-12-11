package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}
	g.Character.Update(g.floor.Blocking(g.Character.X, g.Character.Y, g.camera.X, g.camera.Y), &(g.floor))
	g.camera.Update(g.Character.X, g.Character.Y, &(g.floor), g.floor.QuadtreeContent)
	g.floor.Update(&g.camera.X, &g.camera.Y)
	return nil
}
