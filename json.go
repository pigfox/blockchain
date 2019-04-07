package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func inspect() {
	fmt.Printf("Number of blocks: " + strconv.Itoa(len(blockchain.Blocks)) + "\n")
	b, err := json.Marshal(*blockchain)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%s\n", b)
}