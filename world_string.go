package main



import "bytes"



func (me *World) String() string {
	var buffer bytes.Buffer

	for h:=0; h<me.height; h++ {
		for w:=0; w<me.width; w++ {

			index := me.xyToIndex(w,h)

			buffer.WriteRune(me.buffer[index].Current)
		}
		buffer.WriteRune('\n')
	}

	return buffer.String()
}

