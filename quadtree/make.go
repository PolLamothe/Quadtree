package quadtree

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.

func recur(position string, flootContent [][]int, parent node, initTopLeftX, initTopLeftY int) node {
	var laNode node
	var parentWidth int = parent.width
	var parentHeight int = parent.height
	var parentWidthHalf int = parentWidth
	var parentHeightHalf int = parentHeight
	var parentTopY int = parent.TopLeftY
	var parentTopx int = parent.TopLeftX
	if (parentWidth)%2 != 0 {
		if position == "topLeft" || position == "bottomLeft" {
			parentWidth += 1
		} else {
			parentWidthHalf += 1
		}
	}
	if (parentHeight)%2 != 0 {
		if position == "topLeft" || position == "topRight" {
			parentHeight += 1
		} else {
			parentHeightHalf += 1
		}
	}
	if position == "Root" {
		laNode = node{TopLeftX: parentTopx, TopLeftY: parentTopY, content: -1, height: parentHeight, width: parentWidth}
	} else if position == "topLeft" {
		laNode = node{TopLeftX: parentTopx, TopLeftY: parentTopY, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	} else if position == "topRight" {
		laNode = node{TopLeftX: parentTopx + parentWidthHalf/2, TopLeftY: parentTopY, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	} else if position == "bottomLeft" {
		laNode = node{TopLeftX: parentTopx, TopLeftY: parentTopY + parentHeightHalf/2, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	} else if position == "bottomRight" {
		laNode = node{TopLeftX: parentTopx + parentWidthHalf/2, TopLeftY: parentTopY + parentHeightHalf/2, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	}
	var state bool = false
	if !(laNode.width == 1 && laNode.height == 1) && len(flootContent) > 0 {
		var origin int
		state = true
		var flootContentX, flootContentY int = laNode.TopLeftY, laNode.TopLeftX
		var XRange, YRange int = laNode.TopLeftX, laNode.TopLeftY
		origin = flootContent[flootContentY][flootContentX]
		for i := YRange; i < YRange+laNode.height; i++ {
			for x := XRange; x < XRange+laNode.width; x++ {
				if flootContent[i][x] != origin {
					state = false
				}
			}
		}
	}
	if (laNode.width == 1 && laNode.height == 1) || (state) {
		if configuration.Global.GenerationInfinie {

		} else {
			laNode.content = flootContent[laNode.TopLeftY][laNode.TopLeftX]
		}
		laNode.topLeftNode = nil
		laNode.topRightNode = nil
		laNode.bottomLeftNode = nil
		laNode.bottomRightNode = nil
	} else {
		var topLeftNode node = recur("topLeft", flootContent, laNode, initTopLeftX, initTopLeftY)
		laNode.topLeftNode = &topLeftNode
		if (laNode.width == 1 || laNode.height == 1) && (laNode.width != 1 || laNode.height != 1) {
			if laNode.width <= 1 {
				var bottomLeftNode node = recur("bottomLeft", flootContent, laNode, initTopLeftX, initTopLeftY)
				laNode.topRightNode = &topLeftNode
				laNode.bottomRightNode = &bottomLeftNode
			} else {
				var topRightNode node = recur("topRight", flootContent, laNode, initTopLeftX, initTopLeftY)
				laNode.topRightNode = &topRightNode
				if laNode.height > 1 {
					var bottomRightNode node = recur("bottomRight", flootContent, laNode, initTopLeftX, initTopLeftY)
					laNode.bottomRightNode = &bottomRightNode
				} else {
					laNode.bottomRightNode = &topRightNode
				}
			}
			if laNode.height <= 1 {
				var topRightNode node = recur("topRight", flootContent, laNode, initTopLeftX, initTopLeftY)
				laNode.bottomLeftNode = &topLeftNode
				laNode.bottomRightNode = &topRightNode
			} else {
				var bottomLeftNode node = recur("bottomLeft", flootContent, laNode, initTopLeftX, initTopLeftY)
				laNode.bottomLeftNode = &bottomLeftNode
				if laNode.width > 1 {
					var bottomRightNode node = recur("bottomRight", flootContent, laNode, initTopLeftX, initTopLeftY)
					laNode.bottomRightNode = &bottomRightNode
				} else {
					laNode.bottomRightNode = &bottomLeftNode
				}
			}
		} else {
			var topRightNode node = recur("topRight", flootContent, laNode, initTopLeftX, initTopLeftY)
			var bottomLeftNode node = recur("bottomLeft", flootContent, laNode, initTopLeftX, initTopLeftY)
			var bottomRightNode node = recur("bottomRight", flootContent, laNode, initTopLeftX, initTopLeftY)
			laNode.topRightNode = &topRightNode
			laNode.bottomLeftNode = &bottomLeftNode
			laNode.bottomRightNode = &bottomRightNode
		}
	}
	return laNode
}
func MakeFromArray(floorContent [][]int, width, height, TopLeftX, TopLeftY int) (q Quadtree) {
	var Quad Quadtree = Quadtree{
		Width: width, Height: height}
	var Laroot node = recur("Root", floorContent, node{width: Quad.Width, height: Quad.Height, TopLeftX: TopLeftX, TopLeftY: TopLeftY}, TopLeftX, TopLeftY)
	Quad.Root = &Laroot
	return Quad
}
