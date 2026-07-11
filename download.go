package main

import (
	"fmt"
	"io"
	"net/http"
)

func DownloadPiece(url string, start, end int64) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Range",
		fmt.Sprintf("bytes=%d-%d", start, end),
	)

	//@dbg
	fmt.Println("Request URL:", url)
	fmt.Println("Range Header:", req.Header.Get("Range"))
	//@dbg

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//@dbg
	fmt.Println("Status:", resp.Status)
	fmt.Println("Content-Type:", resp.Header.Get("Content-Type"))
	fmt.Println("Content-Length:", resp.Header.Get("Content-Length"))
	fmt.Println("Accept-Ranges:", resp.Header.Get("Accept-Ranges"))
	fmt.Println("Content-Range:", resp.Header.Get("Content-Range"))
	//@dbg

	if resp.StatusCode != http.StatusPartialContent &&
		resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Received", len(data), "bytes")
	return data, nil
}
