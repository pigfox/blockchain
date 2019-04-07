package main

import "fmt"

type Blockchain struct{
	Blocks []Block `json:"blocks"`
	Difficulty int `json:"difficulty"`
}

type Block struct{
	Index int `json:"index"`
	Timestamp string `json:"timestamp"`
	Data string `json:"data"`
	PreviousHash string `json:"previousHash"`
	Hash string `json:"hash"`
	Nonce int `json:"nonce"`
}

func (b Blockchain) String() string {
	return fmt.Sprintf("[%d, %#v\n]", b.Difficulty, b.Blocks)
}

func (b Block) String() string {
	return fmt.Sprintf("[%#v, %#v, %#v, %#v, %#v, %#v]", b.Index, b.Timestamp, b.Data, b.PreviousHash, b.Hash, b.Nonce)
}