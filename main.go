package main

import (
	"fmt"
)

func main(){
	setBlock(1, "Genesis block", "", 0)
	blockchain.Blocks = append(blockchain.Blocks, *block)
	blockchain.Difficulty = 4
	values := []string {"d.duck", "dumbo", "clown", "cod", "omaha, omaha", "double", "fake", "reverse"}

	for _, value := range values {
		hash := fmt.Sprintf("%s", string(blockchain.Blocks[len(blockchain.Blocks)-1].Hash))
		index := blockchain.Blocks[len(blockchain.Blocks)-1].Index + 1
		setBlock(index, value, hash, block.Nonce)
		blockchain.AddBlock(*block)
	}
	inspect()
	checkValid()
}
