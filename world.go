package main



type World struct {
	width int
	height int
	buffer []Cell
}

func NewWorld(width int, height int) (*World) {

	newBufferLen := width*height

	newBuffer := make([]Cell, newBufferLen)
	for i:=0; i<len(newBuffer);i++ {
		newBuffer[i].Next    = ' '
		newBuffer[i].Current = ' '
	}

	me := World{
		width      : width,
		height     : height,
		buffer     : newBuffer,
	}

	return &me
}



func (me *World) Width() int {
	return me.width
}
func (me *World) Height() int {
	return me.height
}



// Converts (x,y) into their canonical form.
// The canonical form is such that 0 ≤ x < me.width and 0 ≤ y < me.height;
// we use modular arithmetic to put it in that range.
func (me *World) canonicalXY(x, y int) (int, int) {

	xx := x % me.width
	yy := y % me.height

	if 0 > xx {
		xx = xx + me.width
	}
	if 0 > yy {
		yy = yy + me.height
	}

	return xx,yy
}



func (me *World) xyToIndex(x,y int) int {

	x,y = me.canonicalXY(x,y)

	index := y*me.width + x

	return index
}
func (me *World) indexToXY(index int) (x, y int) {
	x = index % me.width
	y = ((index - x) / me.width) % me.height

	// Just in case. (Shouldn't be needed though.)
	x,y = me.canonicalXY(x,y)

	return x,y
}



func (me *World) setNext(x, y int, value rune) {

	index := me.xyToIndex(x,y)

	me.buffer[index].Next = value
}
func (me *World) Set(x, y int, value rune) {

	index := me.xyToIndex(x,y)

	me.buffer[index].Current = value
}
func (me *World) Get(x, y int) rune {

	x,y = me.canonicalXY(x,y)

	index := me.xyToIndex(x,y)

	r := me.buffer[index].Current

	return r
}
