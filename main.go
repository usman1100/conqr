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

	if !utils.IsStringUrl(url) {
		color.Red("Invliad URL provided")
		os.Exit(1)
	}

	if !utils.CheckIfRangeSupported(url) {
		bodyReader, err := utils.DownloadFullFile(url)
		if err != nil {
			color.Red("Error in downloading file", err.Error())
			os.Exit(1)
		}

		fileName := utils.GetFileNameFromUrl(url)

		err = utils.WriteDataToFile(&bodyReader, fileName)

		if err != nil {
			color.Red("Error in writing file", err.Error())
			os.Exit(1)
		}

		color.Green("File downloaded successfully")

	} else {
		err := utils.DownloadInChunks(url, *chunksPtr)
		if err != nil {
			color.Red("Error in downloading file", err.Error())
			os.Exit(1)
		}
	}

}
