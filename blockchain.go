package main

import (
	"os"
)

var blockchain = new(Blockchain)

func (blockchain *Blockchain) AddBlock(block Block) []Block {
	block.PreviousHash = latestBlock().Hash
	blockchain.Blocks = append(blockchain.Blocks, mineBlock(blockchain.Difficulty))
	return blockchain.Blocks
}

func latestBlock() Block{
	return blockchain.Blocks[len(blockchain.Blocks)-1]
}

func checkValid(){
	var len = len(blockchain.Blocks)

	for i := 1; i < len; i++ {
		var currentBlock = blockchain.Blocks[i]
		var previousBlock = blockchain.Blocks[i-1]
		if(currentBlock.Hash != calculateHash(currentBlock.Index, currentBlock.Timestamp, currentBlock.Data,currentBlock.PreviousHash, currentBlock.Nonce)){
			print("Invalid block @ ")
			os.Exit(1)
		}

		if(currentBlock.PreviousHash != previousBlock.Hash){
			print("Invalid block @ ")
			os.Exit(1)
		}
	}

	print("\nValid blockchain!!!")
}