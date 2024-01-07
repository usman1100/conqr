package network

import (
	"errors"
	"net/http"
	"strconv"
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
