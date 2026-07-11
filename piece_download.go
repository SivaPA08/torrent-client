package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func goRoutine(torrent TorrentInfo, start int64, end int64, file *os.File, wg *sync.WaitGroup) {
	defer wg.Done()
	var data []byte
	var err error

	for _, seed := range torrent.WebSeedList {
		url := strings.TrimRight(seed, "/") + "/" + torrent.Name

		fmt.Printf("Trying %s for %d-%d\n", seed, start, end)

		data, err = DownloadPiece(url, start, end)
		if err == nil {
			fmt.Println("Success:", seed)
			break
		}

		fmt.Println("Failed:", seed, err)
	}

	if err != nil {
		fmt.Printf("All web seeds failed for %d-%d\n", start, end)
		return
	}

	_, err = file.WriteAt(data, start)
	if err != nil {
		fmt.Println("Error writing file:", err)
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
	var wg sync.WaitGroup
	no_of_pieces := (torrent.TotalLength + torrent.PieceLength - 1) / torrent.PieceLength
	fmt.Println("Pieces:", len(torrent.Pieces))
	for i := int64(0); i < no_of_pieces; i++ {
		wg.Add(1)
		start := i * torrent.PieceLength
		end := start + torrent.PieceLength - 1
		if end >= torrent.TotalLength {
			end = torrent.TotalLength - 1
		}

		go goRoutine(torrent, start, end, file, &wg)
	}
	wg.Wait()

}
