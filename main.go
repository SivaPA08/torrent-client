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

	root, _, err := ExtractTorrent(data)
	if err != nil {
		panic(err)
	}

	for k, v := range root {
		fmt.Printf("%q -> %T\n", k, v)
	}
	torrentInfo := TorrentInfo{
		Name:        root["info"].(map[string]any)["name"].(string),
		WebSeedList: root["info"].(map[string]any)["url-list"].([]string),
	}

}
