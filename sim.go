package main

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/tests/helper"
	"log"
	"math/big"
)

type TestSim struct { // implements vm.Enviroment
	memdb      *ethdb.MemDatabase
	statedb    *state.StateDB
	blocknum   *big.Int
	blockhash  common.Hash
	time       int64
	difficulty *big.Int
	gasLimit   *big.Int
	gas        *big.Int
	noncemap   map[AccountId]SeqNum
}

const (
	COINBASE = "COINBASE"
	ORIGIN   = "ORIGIN"
)

func NewTestSim() (sim TestSim, err error) {
	var memdb *ethdb.MemDatabase

	memdb, err = ethdb.NewMemDatabase()
	if err != nil {
		return
	}

	roothash := common.BytesToHash(nil) // FIXME
	statedb := state.New(roothash, memdb)

	// FIXME: Handle many dummy values better:
	sim = TestSim{
		memdb:      memdb,
		statedb:    statedb,
		blocknum:   big.NewInt(0),
		blockhash:  common.BytesToHash(nil),
		time:       0,
		difficulty: params.GenesisDifficulty, // FIXME: allow tests to specify.
		gasLimit:   params.GenesisGasLimit,   // FIXME: allow tests to specify.
		gas:        big.NewInt(9876),         // FIXME: determine purpose of this and fix it.
		noncemap:   map[AccountId]SeqNum{},
	}
	return
}

func (sim *TestSim) initAccount(acct *Account) {
	addr := sim.getAddress(acct.Id)
	log.Printf("Initializing account %+v: %+v\n", addr.Hex(), acct)

	obj := state.NewStateObject(*addr, sim.memdb)
	obj.SetBalance(acct.Balance.AsBigInt())

	if acct.ContractState != nil {
		// BUG: What about Storage?
		obj.SetCode([]byte(acct.ContractState.Code))
	}
	sim.statedb.SetStateObject(obj)

	sim.noncemap[acct.Id] = acct.Nonce
}

func (sim *TestSim) applyTransaction(txn *Transaction) (ret []byte, logs state.Logs, gasLeft *big.Int, err error) {
	snapshot := sim.statedb.Copy()

	// Note these need to update per-block when we add multiblock tests:
	coinbase := sim.statedb.GetOrNewStateObject(*sim.getAddress(COINBASE))
	coinbase.SetGasPool(sim.gasLimit)

	message := helper.NewMessage(
		*sim.getAddress(txn.Sender),
		sim.getAddress(txn.To),
		txn.Data,
		txn.Value.AsBigInt(),
		txn.GasLimit.AsBigInt(),
		txn.GasPrice.AsBigInt(),
		uint64(sim.getSenderNonce(txn.Sender)))

	ret, gasLeft, err = core.ApplyMessage(sim, message, coinbase)
	if core.IsNonceErr(err) || core.IsInvalidTxErr(err) {
		sim.statedb.Set(snapshot)
	}
	sim.statedb.Update() // FIXME: What is this?

	return
}

func (sim *TestSim) checkAssertions(as *Assertions, result []byte, logs state.Logs, gasleft *big.Int) (successes uint, failures uint, err error) {
	err = errors.New("Not Implemented.")
	return
}

func (sim *TestSim) getAddress(acct AccountId) *common.Address {
	addr := common.BytesToAddress(sim.getKeyPair(acct).Address())
	return &addr
}

func (sim *TestSim) getKeyPair(acct AccountId) *crypto.KeyPair {
	kp, err := crypto.NewKeyPairFromSec(sim.getSecretKey(acct))
	if err != nil {
		panic(err)
	}
	return kp
}

func (sim *TestSim) getSecretKey(acct AccountId) []byte {
	return crypto.Sha3([]byte(acct))
}

func (sim *TestSim) getSenderNonce(acct AccountId) SeqNum {
	nonce, ok := sim.noncemap[acct]
	if !ok {
		panic(fmt.Sprintf("Unknown acct %+v in TestSim.GetSenderNonce", acct))
	}
	return nonce
}
