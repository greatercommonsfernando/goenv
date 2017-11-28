package main

import (
	"os"
	"fmt"
	"github.com/greatercommonsfernando/goenv/lib"
)

func main() {
	switch {
	case len(os.Args) >= 2:
		if err := lib.CreateGoEnv(os.Args[1]); err != nil {
			panic(err)
		}
	default:
		fmt.Println("You must provide a destination folder")
		lib.Usage()
	}
}