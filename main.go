package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		color.Red("Usage: conqr link1 link2 ...")
		os.Exit(1)
	}

	fmt.Println("Hello World")
}
