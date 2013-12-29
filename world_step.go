package main


func (me *World) Step() {

	// For each cell on in the world.
		lenBuffer := me.width*me.height

		for i:=0; i<lenBuffer; i++ {

			// Convert the buffer index (i) to x-y-coordinates.
				x,y := me.indexToXY(i)


			// Figure out how many neighbords the cell has.
				crowdSize := 0

				if ' ' != me.Get(x-1 , y-1) {
					crowdSize++
				}
				if ' ' != me.Get(x   , y-1) {
					crowdSize++
				}
				if ' ' != me.Get(x+1 , y-1) {
					crowdSize++
				}

				if ' ' != me.Get(x-1 , y) {
					crowdSize++
				}
				if ' ' != me.Get(x+1 , y) {
					crowdSize++
				}

				if ' ' != me.Get(x-1 , y+1) {
					crowdSize++
				}
				if ' ' != me.Get(x   , y+1) {
					crowdSize++
				}
				if ' ' != me.Get(x+1 , y+1) {
					crowdSize++
				}


			// If the cell is dead but the cell has (exactly) 3 neighbords, then make the cell live.
			// Else (if the cell is alive then) if the live cell has either 2 or 3 neighbords, then it stays alive, else make it dead.
			//
			// We store these (new) calculated values in the "Next" field in the cell, because we need the (old) "Current" values to stay
			// there until we have completed all the calculations.
				if ' ' == me.Get(x,y) {

					if 3 == crowdSize {
						me.setNext(x,y, '\u2593')
					} else {
						me.setNext(x,y, ' ')
					}

				} else {

					if 2 > crowdSize || 3 < crowdSize {
						me.setNext(x,y, ' ')
					} else {
						me.setNext(x,y, '\u2588')
					}
				}

		} // for


	// Move the (new) calculated values in the "Next" field of the cell to the "Current" field.
		me.IncrementTimeIndex()
}
