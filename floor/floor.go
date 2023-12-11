package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - Content : partie du terrain qui doit être affichée à l'écran
//   - FullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
type Floor struct {
	Content           [][]int
	FullContent       [][]int
	QuadtreeContent   quadtree.Quadtree
	AllBlockDisplayed bool //variable utilisée pour le vérouillage de la camera au bord de la map
	X, Y              int
	XChange, YChange  int
} //variable utilisée pour modifier la position du joueur lors de la génération au fur et a mesure

// types d'affichage du terrain disponibles
const (
	gridFloor int = iota
	fromFileFloor
	quadTreeFloor
)
