package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func DownloadPiece(url string, start int64, end int64) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, errors.New("New error has came while downloading")
	}
	req.Header.Set(
		"Range",
		fmt.Sprintf("bytes=%d-%d", start, end),
	)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
		return nil, errors.New("IDK some error came in download")
	}
	return io.ReadAll(resp.Body)
}
