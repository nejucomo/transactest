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
	GasLimit Gas
	GasPrice Gas
	Nonce    SeqNum
	Sender   AccountId
	To       AccountId
	Value    Ether
}

type Gas *big.Int
type SeqNum uint
type AccountId string
type Ether *big.Int

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
