package main

import "crypto/rand"

func PeerId() ([20]byte, error) {
	var peerId [20]byte
	//first 8 bytes are mine bro remaining are random idk
	copy(peerId[:8], []byte("-SPA008-"))
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	random := make([]byte, 12)
	if _, err := rand.Read(random); err != nil {
		return peerId, err
	}
	for i, val := range random {
		peerId[8+i] = chars[val%byte(len(chars))]
	}
	return peerId, nil
}
