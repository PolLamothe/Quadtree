package floor

import (
	"io/ioutil"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case fromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case quadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	var data []byte
	var err error
	data, err = ioutil.ReadFile(fileName)
	if err == nil {
		var data2 string = string(data)
		var result [][]int = [][]int{{}}
		for i := 0; i < len(data2); i++ {
			if string(data2[i]) != "\n" {
				value, err := strconv.Atoi(string(data2[i]))
				if err == nil {
					result[len(result)-1] = append(result[len(result)-1], value)
				}
			} else {
				result = append(result, []int{})
			}
		}
		for i := 0; i < configuration.Global.NumTileY/2; i++ {
			result = append(result)
		}
		return result
	}
	return
}
