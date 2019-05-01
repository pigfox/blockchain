package main

import (
	"encoding/json"
	"fmt"
	"github.com/jimlawless/whereami"
	"os"
	"regexp"
	"strconv"
)

var blockchain = &Blockchain{}

func (blockchain *Blockchain) factory(block Block) *Blockchain {
	blockchain = new(Blockchain)
	blockchain.Chain = append(blockchain.Chain, block)
	blockchain.Difficulty = difficulty
	blockchain.MiningReward = miningReward
	blockchain.PendingTransactions = []Transaction{}
	return blockchain
}

func (blockchain *Blockchain) AddBlock(block Block) []Block {
	block.PreviousHash = blockchain.latestBlock().Hash
	blockchain.Chain = append(blockchain.Chain, block.mineBlock(blockchain.Difficulty))
	return blockchain.Chain
}

func (blockchain *Blockchain) createTransaction(t *Transaction) {
	valid := transaction.validate(t)

	if valid {
		blockchain.PendingTransactions = append(blockchain.PendingTransactions, *t)
	} else {
		println("Invalid transaction")
		os.Exit(1)
	}
}

func (blockchain *Blockchain) latestBlock() Block {
	return blockchain.Chain[len(blockchain.Chain)-1]
}

func (blockchain *Blockchain) minePendingTransactions(miningRewardAddress string) {
	// Create new block with all pending transactions and mine it..
	b := block.factory(block.timeStamp(), blockchain.PendingTransactions, blockchain.latestBlock().Hash)
	b.mineBlock(difficulty)

	// Add the newly mined block to the chain
	blockchain.Chain = append(blockchain.Chain, *b)

	// Reset the pending transactions and send the mining reward
	trns := new(Transaction)
	trns.factory("", miningRewardAddress, miningReward)
	blockchain.PendingTransactions = append(blockchain.PendingTransactions, *trns)
}

func (blockchain *Blockchain) checkValid() {
	var len = len(blockchain.Chain)

	for i := 1; i < len; i++ {
		var currentBlock = blockchain.Chain[i]
		var previousBlock = blockchain.Chain[i-1]
		if currentBlock.Hash != block.calculateHash() {
			fmt.Printf("%s\n", whereami.WhereAmI())
			os.Exit(1)
		}

		if currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Printf("%s\n", whereami.WhereAmI())
			os.Exit(1)
		}
	}

	print("\nValid blockchain!!!")
}

func (blockchain *Blockchain) getAddressBalance(address string) float64 {
	balance := float64(0.0) // you start at zero!

	// Loop over each block and each transaction inside the block
	for _, b := range blockchain.Chain {
		for _, t := range b.Transactions {
			matchFromAddress, err := regexp.MatchString(bitcoinAddressRegex, t.FromAddress)

			if err != nil {
				println(err.Error())
			}

			matchToAddress, err := regexp.MatchString(bitcoinAddressRegex, t.ToAddress)

			if err != nil {
				println(err.Error())
			}

			matchAmount, err := regexp.MatchString(floatRegex, fmt.Sprintf("%f", t.Amount))

			if err != nil {
				println(err.Error())
			}

			if matchFromAddress && matchToAddress && matchAmount {
				println("--------------------------")
				println("address:" + address)
				println("t.FromAddress:" + t.FromAddress)
				println("t.ToAddress:" + t.ToAddress)
				println("t.Amount:" + fmt.Sprintf("%f", t.Amount))
				println("--------------------------")

				// If the given address is the sender -> reduce the balance
				if t.FromAddress == address {
					balance = balance - t.Amount
				}

				// If the given address is the receiver -> increase the balance
				if t.ToAddress == address {
					balance = balance + t.Amount
				}
			}
		}
	}

	return balance
}

func (blockchain *Blockchain) inspect(bc *Blockchain) {
	fmt.Printf("Number of blocks: " + strconv.Itoa(len(bc.Chain)) + "\n")
	b, err := json.Marshal(*bc)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Printf("%s\n", b)
}
