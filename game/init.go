package game

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	if configuration.Global.GenerationInfinie {
		configuration.Global.FloorKind = 2
	}
	g.Character.Init()
	g.floor.Init()
	g.camera.Init(g.floor.QuadtreeContent.Width, g.floor.QuadtreeContent.Height, &g.floor.AllBlockDisplayed)
}
