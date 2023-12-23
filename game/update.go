package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/multiplayer"
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
	if configuration.Global.MultiplayerKind == 1 {
		g.Character2.X = multiplayer.ClientPos["X"]
		g.Character2.Y = multiplayer.ClientPos["Y"]
	}
	if configuration.Global.MultiplayerKind == 2 {
		g.Character.X = multiplayer.ServerPos["X"]
		g.Character.Y = multiplayer.ServerPos["Y"]
	}
	g.Character.Update(g.floor.Blocking(g.Character.X, g.Character.Y, int(g.camera.X), int(g.camera.Y)), &(g.floor))
	g.Character2.Update(g.floor.Blocking(g.Character2.X, g.Character2.Y, int(g.camera.X), int(g.camera.Y)), &(g.floor))
	g.Character.RefreshShift()
	if configuration.Global.MultiplayerKind != 2 { // si on est pas en mode client la caméra suit le personnage 1
		g.camera.Update(g.Character.X, g.Character.Y, &(g.floor), g.floor.QuadtreeContent, g.Character.XShift, g.Character.YShift)
	} else { // si on est en mode server la caméra suit le personnage 2
		g.camera.Update(g.Character2.X, g.Character2.Y, &(g.floor), g.floor.QuadtreeContent, g.Character2.XShift, g.Character2.YShift)
	}
	g.floor.Update(int(g.camera.X), int(g.camera.Y), g.Character.XShift, g.Character.YShift)
	return nil
}
