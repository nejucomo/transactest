package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/nejucomo/transactest/assert"
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
	Code    *CodeSource
	Storage state.Storage
}

type CodeSourceType int

const (
	HEX     CodeSourceType = 1
	COMPILE CodeSourceType = 2
)

type CodeSource struct {
	Type CodeSourceType
	Info string
}

func (cs *CodeSource) CheckType() {
	assert.Assert(
		cs.Type == HEX || cs.Type == COMPILE,
		"Invalid CodeSourceType: %#v",
		cs.Type)
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

func (self Ether) Format(f fmt.State, c rune) { fmt.Fprintf(f, "%+v", self.i) }
func (self *Ether) AsBigInt() *big.Int        { return self.i }

type SeqNum uint64
type AccountId string

type Assertions struct {
	Accounts map[AccountId]AccountAssertion
}

type AccountAssertion struct {
	Balance Ether
	Code    *CodeSource
	Nonce   SeqNum
	Storage state.Storage
}
