package main

import (
	"flag"
	"os"

	"github.com/fatih/color"
	"github.com/usman1100/conqr/utils"
)

func main() {

	chunksPtr := flag.Int("c", 10, "Number of concurrent chunks")
	urlPtr := flag.String("u", "", "Url of the file to download")

	flag.Parse()

	url, _ := *urlPtr, *chunksPtr

	if len(url) == 0 {
		color.Red("Invliad URL provided")
		os.Exit(1)
	}

	utils.DownloadFile(url)
}
