package main

import "crypto/sha1"

func ComputeInfoHash(data []byte, start int, end int) [20]byte {
	return sha1.Sum(data[start:end])
}
