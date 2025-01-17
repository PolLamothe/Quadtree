package camera

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"

// Camera définit les caractéristiques de la
// caméra. Pour le moment il s'agit simplement
// des coordonnées absolues de l'endroit où elle
// se trouve mais vous pourrez ajouter des choses
// au besoin lors de votre développement.
type Camera struct {
	X, Y float64
	f    floor.Floor
}

// types de caméra disponibles
const (
	Static int = iota
	FollowCharacter
)
