package network

import (
	"io"
	"net/http"
)

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
