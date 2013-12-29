package main

import "flag"
import "fmt"
import "math/rand"
import "os"
import "time"



func main() {

	// Figure out the width and height of the world.
	//
	// We do this by first trying to get the width and height of the terminal.
	//
	// (Note that, if for whatever reason, we cannot get the width and height
	// of the terminal, then we just default the terminal width and height to
	// the well known (metaphorical) "least common denominator" of 80x24.
	//
	// Then, after that we check to see if the user has specified a width and
	// height.
	//
	// So, we use the terminal width and height as default values (for the
	// width and height) and let the user override those values, if they
	// wish.
	//
	// Although, it is expected that most of the time, the user will probably
	// just want the terminal width and height, and won't override them.
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
	//
	// When we start of, all the cells are dead. To see anything interestng happen,
	// we need to make some of these cells alive.
	//
	// So, we do through each cell, and then use a pseudo-random number generator to
	// help us decide whether to make the cell alive or not.
	//
	// Essentially, each cell has a 1/unrandom chance of being made alive. Note that
	// "unrandom" is a variable. So, for example, if unrandin is 3, then each cell
	// has a probability of 1/3 of being made alive -- has a 1 in 3 chance of being
	// made alive.
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

