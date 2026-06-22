package main

import (
	"errors"
	"strconv"
)




type Torrent struct{
	Announce string
	AnnouceList [][]string
	Name string
	InfoHash []byte
	PieceLength int64
	TotalLenght int64
	Pieces [][]byte
	Files []TorrentFile
}
type TorrentFile struct{
	Path []string
	Lenght int64
}



type Parser struct{
	data []byte
	pos int
}



//constructor def
func New(data []byte) *Parser{
	return &Parser{
		data: data,
		pos: 0,
	}
}


//generic parser 
func (p *Parser)Parse()(any,error){
	if p.pos>=len(p.data){
		return nil,errors.New("End of data")
	}
	switch p.data[p.pos]{
	case 'i':
	case 'l':
	case 'd':
	default:
		b:=p.data[p.pos]
		if b>='0' && b<='9'{
			return nil,nil
		}
		return nil,errors.New("Invalid bin code")

	}
	return nil,nil
}



func (p *Parser)parseInt()(int64,error){
	//parses i<intvalue>e
	// eg: i69e
	p.pos++ //skipping i
	start:=p.pos
	for {
		if p.pos>=len(p.data){
			return 0,errors.New("early EOF wtf")
		}
		if p.data[p.pos]=='e' {
			break
		}
		p.pos++;
	}
	raw:=string(p.data[start:p.pos])
	value,err:=strconv.ParseInt(raw,10,64)
	if err!=nil {
		return 0,err
	}
	p.pos++ //skipping e
	return value,nil
}
func (p *Parser)parseString()([]byte,error){
	start:=p.pos
	for{
		if p.pos>=len(p.data){
			return nil,errors.New("early EOF wtf")
		}
		if p.data[p.pos]==':' {
			break
		}
		p.pos++
	}
	//to get 32:<string> before colen
	length,err:=strconv.Atoi(string(p.data[start:p.pos]))
	if err!=nil {
		return nil,err
	}
	p.pos++ //skipping :
	end:=p.pos+length
	if end>=len(p.data){
		return nil,errors.New("early EOF wtf inside parseString")
	}
	key:=p.data[p.pos:end] //get the key
	p.pos=end
	return key,nil
}


func ExtractTorrent(root map[string]any)(Torrent,error){
	torrent:=Torrent{}
	return torrent,nil
}
