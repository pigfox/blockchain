package main

import "fmt"

type Blockchain struct {
	Chain               []Block       `json:"blocks"`
	Difficulty          int           `json:"difficulty"`
	MiningReward        int           `json:"miningreward"`
	PendingTransactions []Transaction `json:"pendingtransactions"`
}

type Block struct {
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash string        `json:"previousHash"`
	Hash         string        `json:"hash"`
	Nonce        int           `json:"nonce"`
}

type Transaction struct {
	FromAddress string  `json:"fromAddress"`
	ToAddress   string  `json:"toAddress"`
	Amount      float64 `json:"amount"`
}

func (b Blockchain) string() string {
	return fmt.Sprintf("[%d, %d,%#v\n]", b.Difficulty, b.MiningReward, b.Chain)
}

func (b Block) string() string {
	return fmt.Sprintf("[%#v, %#v, %#v, %#v, %#v, %#v]", b.Timestamp, b.Transactions, b.PreviousHash, b.Hash, b.Nonce)
}

func (t Transaction) string() string {
	return fmt.Sprintf("[%#v, %#v, %#v]", t.FromAddress, t.ToAddress, t.Amount)
}
