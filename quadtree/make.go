package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.

func recur(position string, flootContent [][]int, parent node) node {
	var laNode node
	var parentWidth int = parent.width
	var parentHeight int = parent.height
	var parentWidthHalf int = parentWidth
	var parentHeightHalf int = parentHeight
	var parentTopY int = parent.topLeftY
	var parentTopx int = parent.topLeftX
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
	if position == "root" {
		laNode = node{topLeftX: 0, topLeftY: 0, content: -1, height: len(flootContent), width: len(flootContent[0])}
	} else if position == "topLeft" {
		laNode = node{topLeftX: parentTopx, topLeftY: parentTopY, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	} else if position == "topRight" {
		laNode = node{topLeftX: parentTopx + parentWidthHalf/2, topLeftY: parentTopY, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	} else if position == "bottomLeft" {
		laNode = node{topLeftX: parentTopx, topLeftY: parentTopY + parentHeightHalf/2, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	} else if position == "bottomRight" {
		laNode = node{topLeftX: parentTopx + parentWidthHalf/2, topLeftY: parentTopY + parentHeightHalf/2, content: -1, width: parentWidth / 2, height: parentHeight / 2}
	}
	var state bool = false
	if !(laNode.width == 1 && laNode.height == 1) {
		var origin int = flootContent[laNode.topLeftY][laNode.topLeftX]
		state = true
		for i := laNode.topLeftY; i < laNode.topLeftY+laNode.height; i++ {
			for x := laNode.topLeftX; x < laNode.topLeftX+laNode.width; x++ {
				if flootContent[i][x] != origin {
					state = false
				}
			}
		}
	}
	if (laNode.width == 1 && laNode.height == 1) || (state) {
		laNode.content = flootContent[laNode.topLeftY][laNode.topLeftX]
		laNode.topLeftNode = &laNode
		laNode.topRightNode = &laNode
		laNode.bottomLeftNode = &laNode
		laNode.bottomRightNode = &laNode
	} else {
		var topLeftNode node = recur("topLeft", flootContent, laNode)
		laNode.topLeftNode = &topLeftNode
		if (laNode.width == 1 || laNode.height == 1) && (laNode.width != 1 || laNode.height != 1) {
			if laNode.width <= 1 {
				var bottomLeftNode node = recur("bottomLeft", flootContent, laNode)
				laNode.topRightNode = &topLeftNode
				laNode.bottomRightNode = &bottomLeftNode
			} else {
				var topRightNode node = recur("topRight", flootContent, laNode)
				laNode.topRightNode = &topRightNode
				if laNode.height > 1 {
					var bottomRightNode node = recur("bottomRight", flootContent, laNode)
					laNode.bottomRightNode = &bottomRightNode
				} else {
					laNode.bottomRightNode = &topRightNode
				}
			}
			if laNode.height <= 1 {
				var topRightNode node = recur("topRight", flootContent, laNode)
				laNode.bottomLeftNode = &topLeftNode
				laNode.bottomRightNode = &topRightNode
			} else {
				var bottomLeftNode node = recur("bottomLeft", flootContent, laNode)
				laNode.bottomLeftNode = &bottomLeftNode
				if laNode.width > 1 {
					var bottomRightNode node = recur("bottomRight", flootContent, laNode)
					laNode.bottomRightNode = &bottomRightNode
				} else {
					laNode.bottomRightNode = &bottomLeftNode
				}
			}
		} else {
			var topRightNode node = recur("topRight", flootContent, laNode)
			var bottomLeftNode node = recur("bottomLeft", flootContent, laNode)
			var bottomRightNode node = recur("bottomRight", flootContent, laNode)
			laNode.topRightNode = &topRightNode
			laNode.bottomLeftNode = &bottomLeftNode
			laNode.bottomRightNode = &bottomRightNode
		}
	}
	return laNode
}

func MakeFromArray(floorContent [][]int) (q Quadtree) {
	var Quad Quadtree = Quadtree{
		Width: len(floorContent[0]), Height: len(floorContent)}
	var Laroot node = recur("root", floorContent, node{width: len(floorContent[0]), height: len(floorContent), topLeftX: 0, topLeftY: 0})
	Quad.root = &Laroot
	return Quad
}
