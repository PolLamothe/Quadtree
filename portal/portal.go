package portal

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/multiplayer"
)

var PortalStore [][]int = [][]int{}

func IsPortalHere(x, y int) bool {
	for i := 0; i < len(PortalStore); i++ {
		if PortalStore[i][0] == x && PortalStore[i][1] == y {
			return true
		}
	}
	if configuration.Global.MultiplayerKind != 0 {
		for i := 0; i < len(multiplayer.MultiplayerPortal); i++ {
			if multiplayer.MultiplayerPortal[i][0] == x && multiplayer.MultiplayerPortal[i][1] == y {
				return true
			}
		}
	}
	return false
}

func GetOtherCoordonate(x0, y0 int) []int {
	if IsInLocalPortalStore(x0, y0) {
		for i := 0; i < len(PortalStore); i++ {
			if PortalStore[i][0] != x0 || PortalStore[i][1] != y0 {
				return []int{PortalStore[i][0], PortalStore[i][1]}
			}
		}
	} else if configuration.Global.MultiplayerKind != 0 {
		for i := 0; i < len(multiplayer.MultiplayerPortal); i++ {
			if multiplayer.MultiplayerPortal[i][0] != x0 || multiplayer.MultiplayerPortal[i][1] != y0 {
				return []int{multiplayer.MultiplayerPortal[i][0], multiplayer.MultiplayerPortal[i][1]}
			}
		}
	}
	return nil
}

func IsInLocalPortalStore(x, y int) bool {
	for i := 0; i < len(PortalStore); i++ {
		if PortalStore[i][0] == x && PortalStore[i][1] == y {
			return true
		}
	}
	return false
}
