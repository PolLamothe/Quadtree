package portal

var PortalStore [][]int = [][]int{}

func IsPortalHere(x, y int) bool {
	for i := 0; i < len(PortalStore); i++ {
		if PortalStore[i][0] == x && PortalStore[i][1] == y {
			return true
		}
	}
	return false
}

func GetOtherCoordonate(x0, y0 int) []int {
	for i := 0; i < len(PortalStore); i++ {
		if PortalStore[i][0] != x0 || PortalStore[i][1] != y0 {
			return []int{PortalStore[i][0], PortalStore[i][1]}
		}
	}
	return nil
}
