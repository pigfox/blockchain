package main

import (
	"fmt"
)

func main() {
	b := block.genesisBlock()
	chain := blockchain.factory(*b)
	chain.createTransaction(transaction.factory(testUserAddress2, testUserAddress, 200))
	chain.createTransaction(transaction.factory(testUserAddress, testUserAddress2, 500))

	for i := 1; i < numBlocks; i++ {
		newBlock := block.factory(block.timeStamp(), block.Transactions, block.PreviousHash)
		chain.AddBlock(*newBlock)
	}

	blockchain.inspect(chain)
	blockchain.checkValid()
	println("\n================================================\n")

	print("\nStarting the miner...\n")
	chain.minePendingTransactions(testUserAddress)
	balance := chain.getAddressBalance(testUserAddress)
	print("\nBalace of " + testUserAddress + " is:" + fmt.Sprintf("%f", balance))

	print("\nStarting the miner again...\n")
	chain.minePendingTransactions(testUserAddress)
	balance2 := chain.getAddressBalance(testUserAddress)
	print("\nBalace of " + testUserAddress + " is:" + fmt.Sprintf("%f", balance2))
}
