package main

import (
	"github.com/ethereum/go-ethereum/core/state"
	"math/big"
)

type TestSpec struct {
	Accounts     []Account
	Transactions []TransactionAssertions
}

type Account struct {
	Id            AccountId
	Balance       Ether
	ContractState *ContractState
	Nonce         SeqNum
}

type ContractState struct {
	Code    string
	Storage state.Storage
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

type Ether struct {
	i *big.Int
}

func (self *Ether) AsBigInt() *big.Int { return self.i }

type SeqNum uint64
type AccountId string

type Assertions struct {
	Accounts map[AccountId]AccountAssertion
}

type AccountAssertion struct {
	Balance Ether
	Code    ByteCode
	Nonce   SeqNum
	Storage state.Storage
}

type ByteCode []byte
