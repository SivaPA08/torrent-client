package main

import (
	"fmt"
	"os"
)

func main(){
	res,err:=os.ReadFile("torrent/arch.torrent")
	if err!=nil {
		panic(err)
	}
	fmt.Println(string(res))
}
