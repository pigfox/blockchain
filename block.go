package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var block = &Block{}

func (block *Block) factory(timestamp string, transactions []Transaction, previousHash string) *Block {
	block = new(Block)
	block.PreviousHash = previousHash
	block.Timestamp = timestamp
	block.Transactions = transactions
	block.Nonce = 0
	block.Hash = block.calculateHash()
	return block
}

func (block *Block) genesisBlock() *Block {
	t := *new(Transaction)
	var tArray []Transaction
	tArray = append(tArray, t)
	return block.factory(block.timeStamp(), tArray, "")
}

func (block *Block) timeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (block *Block) calculateHash() string {
	trans, err := json.Marshal(block.Transactions)
	if err != nil {
		println(err.Error())
	}

	seed := block.PreviousHash + block.Timestamp + string(trans) + strconv.Itoa(block.Nonce)
	bytes := sha256.Sum256([]byte(seed))
	return fmt.Sprintf("%x", bytes)
}

func (block *Block) mineBlock(difficulty int) Block {
	var zeros string
	for i := 0; i < difficulty; i++ {
		zeros += "0"
	}

	for true {
		if block.Hash[:difficulty] != zeros {
			block.Nonce++
			block.Hash = block.calculateHash()

			if block.Hash[:difficulty] == zeros {
				break
			}
		}
	}

	return *block
}
