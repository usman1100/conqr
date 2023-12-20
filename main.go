package main

import (
	"net/http"
	"os"

	"github.com/fatih/color"
)

func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		color.Red("Usage: conqr link1 link2 ...")
		os.Exit(1)
	}

	firstLink := args[0]

	color.Green("Downloading %s", firstLink)
	response, err := http.Get(firstLink)

	if err != nil {
		color.Red("Error downloading %s", firstLink)
		os.Exit(1)
	}

	color.Green("Downloaded")

	responseSize := response.Request.ContentLength

	color.Yellow("Size: %d", responseSize)
}
