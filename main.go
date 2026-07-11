package main

import (
	"fmt"
	"os"
)

func (t TorrentInfo) Print() {

	fmt.Println(
		"Name:",
		t.Name,
	)

	fmt.Printf(
		"Info Hash: %x\n",
		t.InfoHash,
	)

	fmt.Println(
		"Piece Length:",
		t.PieceLength,
	)

	fmt.Println(
		"Total Length:",
		t.TotalLength,
	)

	fmt.Println(
		"Web Seeds:",
		len(t.WebSeedList),
	)

	// for i, seed := range t.WebSeedList {

	// 	fmt.Printf(
	// 		"  [%d] %s\n",
	// 		i,
	// 		seed,
	// 	)
	// }
}

func main() {
	data, err := os.ReadFile("torrent/arch.torrent")
	if err != nil {
		panic(err)
	}

	root, infoHash, err := ExtractTorrent(data)
	if err != nil {
		panic(err)
	}

	for k, v := range root {
		fmt.Printf("%q -> %T\n", k, v)
	}
	info := root["info"].(map[string]any)

	// Extract web seeds (url-list)
	var webSeeds []string

	if v, ok := root["url-list"]; ok {

		for _, u := range v.([]any) {

			webSeeds = append(
				webSeeds,
				string(
					u.([]byte),
				),
			)
		}
	}

	torrentInfo := TorrentInfo{

		Name: string(
			info["name"].([]byte),
		),

		WebSeedList: webSeeds,

		PieceLength: info["piece length"].(int64),

		TotalLength: info["length"].(int64),

		InfoHash: ComputeInfoHash(data, infoHash.InfoStart, infoHash.InfoEnd), // SHA1 you computed from InfoStart/InfoEnd
	}

	torrentInfo.Print()

	PieceDownload(torrentInfo)
}
