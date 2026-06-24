package main

import (
	"errors"
	"strconv"
)



type Parser struct{
	data []byte
	pos int
	infoHash InfoHash
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
		if key=="info" {
			p.infoHash.InfoStart=p.pos
		}
		value,err:=p.parse()

		if err!=nil {
			return nil,err
		}
		if key=="info" {
			p.infoHash.InfoEnd=p.pos
		}
		res[key]=value
	}
	p.pos++ //skipping e
	return res,nil
}


func ExtractTorrent(data []byte)(map[string]any,InfoHash,error){
	p:=new(data)
	rootAny,err:=p.parse()
	if err!=nil {
		return nil,InfoHash{},err
	}
	root,ok:=rootAny.(map[string]any)
	if !ok {
		return nil,InfoHash{},errors.New("top level myst be a dict")
	}
	if p.pos!=len(p.data){
		return nil,InfoHash{},errors.New("leftover data found")
	}
	return root,p.infoHash,nil 

}
