package main

type Parser struct {
	data     []byte
	pos      int
	infoHash InfoHashPos
}

type TorrentFile struct {
	Path   []string //it stores like ["src","main.sh"] like that
	Length int64
}

type TorrentInfo struct {
	Name         string
	Announce     string   //just added for future updates
	AnnounceList []string //just added for future updates comming soon :)
	WebSeedList  []string
	PieceLength  int64
	Length       int64
	InfoHash     [20]byte
	PieceLenght  int64
	TotalLengthi int64
	Pieces       [][]byte
	Files        []TorrentFile
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
