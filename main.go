package main

import "flag"
import "fmt"
import "math/rand"
import "os"
import "time"



func main() {

	// Figure out the width and height of the world.
		stdinFd := int(os.Stdout.Fd())

		terminalWidth, terminalHeight, err := TerminalSize(stdinFd)
		if nil != err {
			// Could not get the terminal width and height. So just default to 80x24.
			terminalWidth, terminalHeight = 80, 24
		}

		var worldWidth  int
		var worldHeight int

		flag.IntVar(&worldWidth  , "width"  , terminalWidth  , "the width of the world")
		flag.IntVar(&worldHeight , "height" , terminalHeight , "the height of the world")


	// Initialize the world.
		world := NewWorld(worldWidth, worldHeight)


	// Set up world.
		var unrandom int

		flag.IntVar(&unrandom, "unrandom", 3, "each cell is initially alive, with a probability of 1/unrandom")

		randomness := rand.New(rand.NewSource( time.Now().UnixNano() ))

		for y:=0; y<terminalHeight; y++ {
			for x:=0; x<terminalWidth; x++ {

				if 0 == randomness.Intn(unrandom) {
					world.Set(x,y, '\u2588')
				}

			}
		}


	// Make time progress.
		fmt.Print("\x0c") // form feed character -- ^L
		fmt.Print(world)
		for {
			time.Sleep(350 * time.Millisecond)

			world.Step()
			fmt.Print("\x0c") // form feed character -- ^L
			fmt.Print(world)
		}
}

