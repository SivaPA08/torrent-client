package main

type Parser struct {
	data     []byte
	pos      int
	infoHash InfoHashPos
}

type InfoHashPos struct {
	InfoStart int
	InfoEnd   int
}

type Tracker struct {
	URL string
}

type AnnounceRequest struct {
	InfoHash   [20]byte
	PeerId     [20]byte
	Port       uint16
	Uploaded   int64
	Downloaded int64
	Left       int64
	Compact    bool
	Event      string
}
