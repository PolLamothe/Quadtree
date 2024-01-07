package quadtree

import (
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/multiplayer"
)

/*
GenerateInfinite permet d'étendre le quadtree actuelle en placant la map actuelle dans un des coins du nouveau quadtree
*/

func (q *Quadtree) GenerateInfinite(postion string) {
	fmt.Println(postion)
	var Root *node = &node{width: q.Width * 2, height: q.Height * 2, content: -1} //on définie les dimensions de la Root du nouveau quadtree
	if postion == "TopLeft" {                                                     //on définie les coordonnées de la map déja générée en fonction de la position dans la Root
		(*Root).TopLeftX = q.Root.TopLeftX
		(*Root).TopLeftY = q.Root.TopLeftY
	}
	if postion == "TopRight" {
		(*Root).TopLeftY = q.Root.TopLeftY
		(*Root).TopLeftX = q.Root.TopLeftX - q.Width
	}
	if postion == "BottomLeft" {
		(*Root).TopLeftY = q.Root.TopLeftY - q.Height
		(*Root).TopLeftX = q.Root.TopLeftX
	}
	if postion == "BottomRight" {
		(*Root).TopLeftX = q.Root.TopLeftX - q.Width
		(*Root).TopLeftY = q.Root.TopLeftY - q.Height
	}
	(*Root).topRightNode = MakeFromArray([][]int{}, Root.width, Root.height, Root.TopLeftX+q.Width, Root.TopLeftY).Root
	(*Root).bottomLeftNode = MakeFromArray([][]int{}, Root.width, Root.height, Root.TopLeftX, Root.TopLeftY+q.Height).Root
	(*Root).bottomRightNode = MakeFromArray([][]int{}, Root.width, Root.height, Root.TopLeftX+q.Width, Root.TopLeftY+q.Height).Root
	(*Root).topLeftNode = MakeFromArray([][]int{}, Root.width, Root.height, Root.TopLeftX, Root.TopLeftY).Root
	if postion == "TopLeft" {
		(*Root).topLeftNode = q.Root
	}
	if postion == "TopRight" {
		(*Root).topRightNode = q.Root
	}
	if postion == "BottomLeft" {
		(*Root).bottomLeftNode = q.Root
	}
	if postion == "BottomRight" {
		(*Root).bottomRightNode = q.Root
	}
	(*q).Width *= 2
	(*q).Height *= 2
	(*q).Root = Root
	if multiplayer.Conn != nil {
		multiplayer.SendBlock()
	}
}
