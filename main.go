package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {

	chunksPtr := flag.Int("c", 10, "Number of concurrent chunks")
	urlPtr := flag.String("u", "", "Url of the file to download")

	flag.Parse()

	url, chunks := *urlPtr, *chunksPtr

	if len(url) == 0 {
		color.Red("Invliad URL provided")
		os.Exit(1)
	}

	fmt.Println(url, chunks)
}
