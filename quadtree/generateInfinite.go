package quadtree

import (
	"fmt"
)

func (q *Quadtree) GenerateInfinite(postion string) {
	var Root *node = &node{width: q.Width * 2, height: q.Height * 2, content: -1}
	fmt.Println(postion)
	if postion == "TopLeft" { // marche
		(*Root).TopLeftX = q.Root.TopLeftX
		(*Root).TopLeftY = q.Root.TopLeftY
	}
	if postion == "TopRight" { // marche
		(*Root).TopLeftY = q.Root.TopLeftY
		(*Root).TopLeftX = q.Root.TopLeftX - q.Width
	}
	if postion == "BottomLeft" { // marche
		(*Root).TopLeftY = q.Root.TopLeftY - q.Height
		(*Root).TopLeftX = q.Root.TopLeftX
	}
	if postion == "BottomRight" { // marche
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
}
