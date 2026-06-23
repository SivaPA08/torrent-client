package main

import (
	"errors"
	"strconv"
)




type Torrent struct{
	Announce string
	AnnounceList [][]string
	Name string
	InfoHash []byte
	PieceLength int64
	TotalLength int64
	Pieces [][]byte
	Files []TorrentFile
}
type TorrentFile struct{
	Path []string
	Length int64
}



type Parser struct{
	data []byte
	pos int
}



//constructor def
func new(data []byte) *Parser{
	return &Parser{
		data: data,
		pos: 0,
	}
}


//generic parser 
func (p *Parser)parse()(any,error){
	if p.pos>=len(p.data){
		return nil,errors.New("End of data")
	}
	switch p.data[p.pos]{
	case 'i':
		return p.parseInt()
	case 'l':
		return p.parseList()
	case 'd':
		return p.parseDict()
	default:
		b:=p.data[p.pos]
		if b>='0' && b<='9'{
			return p.parseString()
		}
		return nil,errors.New("Invalid bin code")

	}
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
	if end>len(p.data){
		return nil,errors.New("early EOF wtf inside parseString")
	}
	key:=p.data[p.pos:end] //get the key
	p.pos=end
	return key,nil
}

func (p *Parser) parseList()([]any,error){
	p.pos++ //skipping l
	var res []any
	for {
		if p.pos>=len(p.data){
			return nil,errors.New("early EOF in ParseList wtf")
		}
		if p.data[p.pos]=='e' {
			break
		}
		value,err:=p.parse()
		if err!=nil {
			return nil,err
		}
		res=append(res,value)
	}
	p.pos++ //skipping e
	return res,nil
}

func (p *Parser)parseDict()(map[string]any,error){
	p.pos++ //skipping d
	res:=make(map[string]any)
	for {
		if p.pos>=len(p.data){
			return nil,errors.New("early EOF in ParseDict wtf")
		}
		if p.data[p.pos]=='e' {
			break
		}
		keyBytes,err:=p.parseString()
		if err!=nil {
			return nil,err
		}
		key:=string(keyBytes)
		value,err:=p.parse()
		if err!=nil {
			return nil,err
		}
		res[key]=value
	}
	p.pos++ //skipping e
	return res,nil
}


func extract(root map[string]any)(Torrent,error) {
	torrent:=Torrent{}
	//fetching announce
	if v,ok:=root["announce"]; ok {
		torrent.Announce=string(v.([]byte))
	}
	//fetching announce-list
	if v,ok:=root["announce-list"];ok {
		tiers := v.([]any)
		for _,tier :=range tiers {
			trackers :=tier.([]any)
			row:=[]string{}


			for _,t :=range trackers {
				row=append(row,string(t.([]byte)))
			}
			torrent.AnnounceList=append(torrent.AnnounceList,row)
		}
	}



	//fetching info 
	info,ok:=root["info"].(map[string]any)
	if !ok {
		return torrent,errors.New("missing info")
	}



	//takng name
	torrent.Name=string(info["name"].([]byte))



	//piece len
	pieces:=info["pieces"].([]byte)
	for i:=0;i<len(pieces);i+=20 { //20bytes for sha1
		hash:=pieces[i:i+20]
		torrent.Pieces=append(torrent.Pieces,hash) //appending to pices
	}



	//single file torrent
	if v,ok:=info["length"]; ok {
		torrent.TotalLength=v.(int64)
	}



	//multi file torrent
	if v,ok:=info["files"]; ok {
		files:=v.([]any)
		for _,f:= range files {
			file:=f.(map[string]any)
			len:=file["length"].(int64)
			path:=file["path"].([]any)
			var parts []string
			for _,p:=range path {
				parts=append(parts,string(p.([]byte)))
			}
			torrent.Files=append(torrent.Files,TorrentFile{Path:parts,Length: len})
			torrent.TotalLength+=len
		}
	}
	return torrent,nil
}
func ExtractTorrent(data []byte)(Torrent,error){
	p:=new(data)
	rootAny,err:=p.parse()
	if err!=nil {
		return Torrent{},err
	}
	root,ok:=rootAny.(map[string]any)
	if !ok {
		return Torrent{},errors.New("top level myst be a dict")
	}
	if p.pos!=len(p.data){
		return Torrent{},errors.New("leftover data found")
	}
	return extract(root) //this return (Torrent,error) type

}
