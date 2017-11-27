package main

import (
	"os"
	"fmt"
	"goenv/lib"
)

func main() {
	switch {
	case len(os.Args) >= 2:
		lib.CreateGoEnv(os.Args[1])
	default:
		fmt.Println("You must provide a destination folder")
		lib.Usage()
	}
}