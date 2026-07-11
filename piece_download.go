package main

import "os"

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
	data, err := DownloadPiece(torrent.WebSeedList[0], 0, torrent.PieceLenght)
	if err != nil {
		return
	}
	_, err = file.WriteAt(data, torrent.PieceLenght)
	if err != nil {
		return
	}

}
