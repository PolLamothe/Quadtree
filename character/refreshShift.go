package character

func (c *Character) RefreshShift() {

	xShift := 0
	yShift := 0
	switch c.orientation {
	case orientedDown:
		yShift = c.shift
	case orientedUp:
		yShift = -c.shift
	case orientedLeft:
		xShift = -c.shift
	case orientedRight:
		xShift = c.shift
	}
	c.XShift = xShift
	c.YShift = yShift
}
