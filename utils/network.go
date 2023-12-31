package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

func CheckIfRangeSupported(url string) bool {
	head, err := http.Head(url)

	if err != nil {
		return false
	}

	rangeHeader := head.Header.Get("Accept-Ranges")

	return rangeHeader == "bytes"
}

func GetRequestContentLength(url string) (int, error) {

	head, err := http.Head(url)

	if err != nil {
		return 0, err
	}

	rangeHeader := head.Header.Get("Content-Length")

	if len(rangeHeader) == 0 {
		return 0, errors.New("cannot get content length")
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

// download the whole file, without range
func DownloadFullFile(url string) (io.ReadCloser, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func DownloadFile(fileUrl string, numberOfChunks int) {

	fileName := GetFileNameFromUrl(fileUrl)

	fileLength, err := GetRequestContentLength(fileUrl)

	if err != nil {
		color.Red(err.Error())
	}

	chunkSize := fileLength / numberOfChunks

	fmt.Println("Downloading '", fileName, "'")
	fmt.Println("File size '", fileLength, "'")

	for i := 0; i < numberOfChunks; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		DownloadChunk(start, end, fileUrl)
	}

}
