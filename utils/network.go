package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

func GetRequestContentLength(url string) (int, error) {

	head, err := http.Head(url)

	if err != nil {
		return 0, err
	}

	rangeHeader := head.Header.Get("Content-Length")

	if len(rangeHeader) == 0 {
		return 0, errors.New("multi part download is not suported on this file")
	}

	convertedRange, err := strconv.Atoi(rangeHeader)

	if err != nil {
		return 0, err
	}

	return convertedRange, nil

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

	fileLength, err := GetRequestContentLength(fileUrl)

	if err != nil {
		color.Red(err.Error())
	}

	fmt.Println("Downloading '", fileName, "'")
	fmt.Println("File size '", fileLength, "'")

	DownloadChunk(0, 1, fileUrl)

}
