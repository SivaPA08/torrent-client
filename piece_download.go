package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func goRoutine(url string, start int64, end int64, file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := DownloadPiece(url, start, end)
	if err != nil {
		fmt.Println("Download Failed")
		return
	}
	_, err = file.WriteAt(data, 0)
	if err != nil {
		fmt.Println("Error while writting file")
	}

}

func PieceDownload(torrent TorrentInfo) {
	file, err := os.OpenFile(
		"output.iso",
		os.O_CREATE|os.O_RDWR,
		0644,
	)
	if err != nil {
		return
	}
	defer file.Close()
	url := strings.TrimRight(torrent.WebSeedList[0], "/") + "/" + torrent.Name
	// start := int64(0)
	// end := torrent.PieceLength - 1
	//@dbg
	// fmt.Println("Piece Length:", torrent.PieceLength)
	// fmt.Println("start:", start)
	// fmt.Println("end:", end)
	//@dbg

	var wg sync.WaitGroup
	no_of_pieces := int64(len(torrent.Pieces))
	for i := int64(0); i < no_of_pieces; i++ {
		wg.Add(1)
		start := i * torrent.PieceLength
		end := start + torrent.PieceLength - 1
		if end >= torrent.TotalLength {
			end = torrent.TotalLength - 1
		}

		go goRoutine(url, start, end, file, &wg)
	}
	wg.Wait()

}
