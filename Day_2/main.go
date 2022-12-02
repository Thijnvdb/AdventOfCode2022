package main

import (
	"day2/puzzles/puzzle1"
	"day2/puzzles/puzzle2"
	"os"
)

func main() {
	// this is made asuming I am not dumb and do not need error handling on my command line args...
	switch os.Args[1] {
	case "1":
		puzzle1.Run(os.Args[2])
	case "2":
		puzzle2.Run(os.Args[2])
	}
}
