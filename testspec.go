package main

import (
	"math/big"
)

type TestSpec struct {
	Transactions []TransactionAssertions
}

type TransactionAssertions struct {
	Transaction Transaction
	Assertions  Assertions
}

type Transaction struct {
	Data     []byte
	GasLimit Ether
	GasPrice Ether
	Nonce    SeqNum
	Sender   AccountId
	To       AccountId
	Value    Ether
}

type Ether big.Int

func (self *Ether) AsBigInt() *big.Int { return (*big.Int)(self) }

type SeqNum uint
type AccountId string

type Assertions struct {
	Accounts map[AccountId]AccountAssertion
}

type AccountAssertion struct {
	Balance Ether
	Code    ByteCode
	Nonce   SeqNum
	Storage map[string]string
}

type ByteCode []byte
