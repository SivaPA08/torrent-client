package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("torrent/arch.torrent")

	if err != nil {
		panic(err)
	}

	torrent, err :=ExtractTorrent(data)

	if err != nil {
		panic(err)
	}

	fmt.Println("Name:",torrent.Name,)
	fmt.Println("Announce:",torrent.Announce,)
	fmt.Println("Piece Length:",torrent.PieceLength,)
	fmt.Println("Total Length:",torrent.TotalLength,)
	fmt.Println("Pieces:",len(torrent.Pieces),)
}
