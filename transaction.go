package main

import (
	"fmt"
	"reflect"
	"regexp"
)

var transaction = &Transaction{}

func (transaction *Transaction) factory(from string, to string, amount float64) *Transaction {
	transaction = new(Transaction)
	transaction.FromAddress = from
	transaction.ToAddress = to
	transaction.Amount = amount
	return transaction
}

func (transaction *Transaction) validate(t *Transaction) bool {
	rv := true

	if reflect.TypeOf(*t) != reflect.TypeOf(Transaction{}) {
		println("t is NOT of type: " + reflect.TypeOf(Transaction{}).String())
		rv = false
	}

	fromAddress := fmt.Sprintf("%s", t.FromAddress)
	match, err := regexp.MatchString(bitcoinAddressRegex, fromAddress)

	if err != nil {
		println(err.Error())
	}

	if !match {
		println("Invalid FromAddress")
		rv = false
	}

	toAddress := fmt.Sprintf("%s", t.ToAddress)
	match2, err := regexp.MatchString(bitcoinAddressRegex, toAddress)

	if err != nil {
		println(err.Error())
	}

	if !match2 {
		println("Invalid ToAddress")
		rv = false
	}

	amount := fmt.Sprintf("%f", t.Amount)
	match3, err := regexp.MatchString(floatRegex, amount)

	if err != nil {
		println(err.Error())
	}

	if !match3 {
		println("Invalid Amount")
		rv = false
	}

	return rv
}
