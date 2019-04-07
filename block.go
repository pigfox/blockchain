package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

var difficulty = 4
var nonce = 0
var hash = ""
var block = new (Block)

func timeStamp() string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func calculateHash(index int, time string, data string, ph string, nonce int) string{
	seed := strconv.Itoa(index) + ph + time + data + strconv.Itoa(nonce)
	bytes := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%x", bytes)
}

func setBlock(i int, d string, ph string, n int) {
	block.Index = i
	block.Timestamp = timeStamp()
	block.Data = d
	block.PreviousHash = ph
	block.Nonce = n
	block.Hash = calculateHash(block.Index,block.Timestamp, block.Data, block.PreviousHash, block.Nonce)
}

func mineBlock(difficulty int) Block{
	var zeros string
	for i := 0; i < difficulty; i++ {
		zeros += "0"
	}

	for true {
		if block.Hash[:difficulty] != zeros {
			block.Nonce++
			block.Hash = calculateHash(block.Index, block.Timestamp, block.Data, block.PreviousHash, block.Nonce)

			if(block.Hash[:difficulty] == zeros){
				break
			}
		}
	}

	return *block
}

