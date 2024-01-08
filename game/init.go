package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/multiplayer"
	"os"
	"path/filepath"
	"strconv"
)

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	if len(os.Args) > 1 {
		v, _ := strconv.Atoi(os.Args[1])
		configuration.Global.MultiplayerKind = v
	}
	if configuration.Global.MultiplayerKind != 0 {
		configuration.Global.FloorKind = 2
	}
	if configuration.Global.GenerationInfinie {
		configuration.Global.FloorKind = 2
	}
	if configuration.Global.TerreRonde {
		configuration.Global.CameraBlockEdge = false
	}
	if configuration.Global.GenerationInfinie {
		configuration.Global.CameraBlockEdge = false
		configuration.Global.TerreRonde = false
	}
	if configuration.Global.MultiplayerKind != 2 {
		g.floor.Init()
		g.Character.CharacterNumber = 1
		g.Character.Init()
	}
	g.camera.Init()
	if configuration.Global.MultiplayerKind == 2 {
		go multiplayer.InitAsClient()
		path, err := filepath.Abs("../multiplayer/BlockGeneratedClient")
		if err != nil {
			os.Exit(1)
		}
		os.Truncate(path, 0)
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
		path, err := filepath.Abs("../multiplayer/BlockGeneratedServer")
		if err != nil {
			os.Exit(1)
		}
		os.Truncate(path, 0)
		multiplayer.Map = g.floor.FullContent
		go multiplayer.ConnectAsServer()
	}
}
