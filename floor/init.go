package floor

import (
	"io/ioutil"
	"math/rand"
	"strconv"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.AllBlockDisplayed = false
	f.Content = make([][]int, configuration.Global.NumTileY)
	f.XChange, f.YChange = 0, 0
	for y := 0; y < len(f.Content); y++ {
		f.Content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case fromFileFloor:
		if configuration.Global.RandomGeneration {
			var RandomFloor [][]int
			for i := 0; i < configuration.Global.RandomTileY; i++ {
				RandomFloor = append(RandomFloor, []int{})
				for x := 0; x < configuration.Global.RandomTileX; x++ {
					var random int = rand.Intn(5)
					RandomFloor[i] = append(RandomFloor[i], random)
				}
			}
			f.FullContent = RandomFloor
			f.QuadtreeContent = quadtree.MakeFromArray(f.FullContent, len(f.FullContent[0]), len(f.FullContent), 0, 0)
		} else {
			f.FullContent = readFloorFromFile(configuration.Global.FloorFile)
			f.QuadtreeContent = quadtree.MakeFromArray(f.FullContent, len(f.FullContent[0]), len(f.FullContent), 0, 0)
		}
	case quadTreeFloor:
		f.FullContent = readFloorFromFile(configuration.Global.FloorFile)
		if configuration.Global.RandomGeneration && !configuration.Global.GenerationInfinie {
			var RandomFloor [][]int
			for i := 0; i < configuration.Global.RandomTileY; i++ {
				RandomFloor = append(RandomFloor, []int{})
				for x := 0; x < configuration.Global.RandomTileX; x++ {
					var random int = rand.Intn(5)
					RandomFloor[i] = append(RandomFloor[i], random)
				}
			}
			f.FullContent = RandomFloor
			f.QuadtreeContent = quadtree.MakeFromArray(RandomFloor, len(RandomFloor[0]), len(RandomFloor), 0, 0)
		} else {
			if configuration.Global.GenerationInfinie {
				f.QuadtreeContent = quadtree.MakeFromArray([][]int{}, configuration.Global.NumTileX*4, configuration.Global.NumTileY*4, -configuration.Global.RandomTileX, -configuration.Global.RandomTileY)
			} else {
				f.FullContent = readFloorFromFile(configuration.Global.FloorFile)
				f.QuadtreeContent = quadtree.MakeFromArray(f.FullContent, len(f.FullContent[0]), len(f.FullContent), 0, 0)
			}
		}
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
		for i := 0; i < len(data2); i++ { //pour chaque caractère du fichier
			if string(data2[i]) != "\n" { //si le caractere n'es pas une retour a la ligne
				value, err := strconv.Atoi(string(data2[i]))
				if err == nil {
					result[len(result)-1] = append(result[len(result)-1], value) //on ajoute la valeur au tableau
				}
			} else {
				result = append(result, []int{}) //sinon on ajoute une ligne au tableau
			}
		}
		return result
	}
	return
}
