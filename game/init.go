package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/multiplayer"
	"os"
	"strconv"
)

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	v, _ := strconv.Atoi(os.Args[1])
	configuration.Global.MultiplayerKind = v
	if configuration.Global.GenerationInfinie {
		configuration.Global.FloorKind = 2
	}
	if configuration.Global.TerreRonde {
		configuration.Global.CameraBlockEdge = false
		configuration.Global.GenerationInfinie = false
	}
	if configuration.Global.GenerationInfinie {
		configuration.Global.CameraBlockEdge = false
	}
	if configuration.Global.MultiplayerKind != 2 {
		g.floor.Init()
		g.Character.CharacterNumber = 1
		g.Character.Init()
		g.Character2.CharacterNumber = 2
		g.Character2.Init()
	}
	g.camera.Init()
	if configuration.Global.MultiplayerKind == 2 {
		go multiplayer.InitAsClient()
		for {
			if multiplayer.MapReceived {
				g.floor.Init()
				g.Character.CharacterNumber = 1
				g.Character.Init()
				g.Character2.CharacterNumber = 2
				g.Character2.Init()
				break
			}
		}
	}
	if configuration.Global.MultiplayerKind == 1 {
		multiplayer.Map = g.floor.FullContent
		go multiplayer.ConnectAsServer()
	}
}
