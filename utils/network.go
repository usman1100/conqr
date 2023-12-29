package utils

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func CheckIfRangeHeaderIsSupported(url string) {

	head, err := http.Head(url)

	if err != nil {
		color.Red("Could not send head request")
	}

	rangeHeader := head.Request.Header.Get("range")

	rangeHeaderSupported := len(rangeHeader) == 0

	if rangeHeaderSupported {
		color.Green("Multi part download is suported on this file")
	} else {
		color.Red("Server does not support multipart downloads")
	}

}

func DownloadFile(fileUrl string) {

	fileName := GetFileNameFromUrl(fileUrl)

	CheckIfRangeHeaderIsSupported(fileUrl)

	fmt.Println("Downloading", fileName)

	response, err := http.Get(fileUrl)

	if err != nil {
		color.Red("Could not download file, %s", err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	fmt.Println("Size:", response.ContentLength, "bytes")

}
