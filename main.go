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

	root, infoHashPos, err := ExtractTorrent(data)
	if err != nil {
		panic(err)
	}

	for k, v := range root {
		fmt.Printf("%q -> %T\n", k, v)
	}
	fmt.Printf("%+v\n", infoHashPos)
	hash := ComputeInfoHash(
		data,
		infoHashPos.InfoStart,
		infoHashPos.InfoEnd,
	)

	fmt.Printf("%x\n", hash)
	val, err := PeerId()
	fmt.Println(string(val[:]))
}
