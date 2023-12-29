package utils

import (
	"fmt"
	"net/http"

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

func DownloadChunk(rangeStart int, rangeEnd int, url string) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		color.Red("Error in making request at chunk", rangeStart, rangeEnd)
		return
	}
	rangeValue := fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd)
	req.Header.Set("Range", rangeValue)

	res, err := httpClient.Do(req)

	if err != nil {
		color.Red("Request failed for chunk", rangeStart, rangeEnd)
		return
	}

	defer res.Body.Close()

	fmt.Println(res.ContentLength)
}

func DownloadFile(fileUrl string) {

	fileName := GetFileNameFromUrl(fileUrl)

	CheckIfRangeHeaderIsSupported(fileUrl)

	fmt.Println("Downloading '", fileName, "'")

	DownloadChunk(0, 1, fileUrl)

}
