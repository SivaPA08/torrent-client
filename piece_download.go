package main

import (
	"fmt"
	"os"
	"strings"
)

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
	start := int64(0)
	end := torrent.PieceLength - 1
	//@dbg
	fmt.Println("Piece Length:", torrent.PieceLength)
	fmt.Println("start:", start)
	fmt.Println("end:", end)
	//@dbg

	data, err := DownloadPiece(url, start, end)
	if err != nil {
		return
	}

	_, err = file.WriteAt(data, 0)
	if err != nil {
		return
	}

}
