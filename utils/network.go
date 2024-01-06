package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"

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

func DownloadChunk(rangeStart int, rangeEnd int, url string) (io.ReadCloser, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		color.Red("Error in making request at chunk", rangeStart, rangeEnd)
		return nil, err
	}
	rangeValue := fmt.Sprintf("bytes=%d-%d", rangeStart, rangeEnd)
	req.Header.Set("Range", rangeValue)

	res, err := httpClient.Do(req)

	if err != nil {
		color.Red("Request failed for chunk", rangeStart, rangeEnd)
		return nil, err
	}

	// defer res.Body.Close()

	return res.Body, nil
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

func DownloadInChunks(fileUrl string, numberOfChunks int) error {

	var wg sync.WaitGroup

	folderName := GetFileNameFromUrl(fileUrl)

	uniqueFolderName := GenerateFileName(folderName)

	fileLength, err := GetRequestContentLength(fileUrl)

	if err != nil {
		color.Red(err.Error())
	}

	chunkSize := fileLength / numberOfChunks

	// create folder
	err = os.Mkdir(uniqueFolderName, os.ModePerm)

	if err != nil {
		color.Red("Error in creating folder", uniqueFolderName)
		return errors.New("error in creating folder")
	}

	color.Yellow(fmt.Sprintf("File size %d", fileLength))
	color.Yellow(fmt.Sprintf("Chunk size %d", chunkSize))

	for i := 0; i < numberOfChunks; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := ((i + 1) * chunkSize) - 1
			color.Green(fmt.Sprintf("Downloading chunk %d from %d to %d", i, start, end))
			bodyReader, err := DownloadChunk(start, end, fileUrl)

			if err != nil {
				color.Red("Error in downloading chunk", start, end)
			}
			chunkName := uniqueFolderName + "/" + strconv.Itoa(i) + ".chunk"
			err = WriteDataToFile(&bodyReader, chunkName)

			if err != nil {
				color.Red("Error in writing chunk", start, end)
				fmt.Println("error in writing chunk" + err.Error())
			}
		}(i)
	}

	// download left over

	leftOverBytes := fileLength % numberOfChunks

	if leftOverBytes != 0 {

		lastChunkStart := fileLength - leftOverBytes
		lastChunkEnd := fileLength - 1

		lastChunkName := uniqueFolderName + "/" + strconv.Itoa(numberOfChunks) + ".chunk"
		color.Green(fmt.Sprintf("Downloading chunk %d from %d to %d", numberOfChunks, lastChunkStart, lastChunkEnd))
		bodyReader, err := DownloadChunk(lastChunkStart, lastChunkEnd, fileUrl)

		if err != nil {
			color.Red("Error in downloading chunk", lastChunkStart, lastChunkEnd)
		}
		err = WriteDataToFile(&bodyReader, lastChunkName)

		if err != nil {
			color.Red("Error in writing chunk", lastChunkStart, lastChunkEnd)
			fmt.Println("error in writing chunk" + err.Error())
		}
	}

	wg.Wait()

	StitchChunksIntoFile(uniqueFolderName, numberOfChunks)

	return nil

}
