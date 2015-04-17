package main

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/tests/helper"
	"math/big"
)

type TestSim struct { // implements vm.Enviroment
	statedb    *state.StateDB
	blocknum   *big.Int
	blockhash  common.Hash
	time       int64
	difficulty *big.Int
	gasLimit   *big.Int
	gas        *big.Int
	addrmap    map[AccountId]*common.Address
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
		statedb:    statedb,
		blocknum:   big.NewInt(0),
		blockhash:  common.BytesToHash(nil),
		time:       0,
		difficulty: big.NewInt(0),
		gasLimit:   big.NewInt(0),
		gas:        big.NewInt(0),
		addrmap:    map[AccountId]*common.Address{},
	}
	return
}

func (sim *TestSim) applyTransaction(txn *Transaction) (ret []byte, logs state.Logs, gasLeft *big.Int, err error) {
	var (
		keyPair *crypto.KeyPair
	)

	keyPair, err = crypto.NewKeyPairFromSec([]byte("FIXME stub secret"))
	if err != nil {
		return
	}

	snapshot := sim.statedb.Copy()

	// Note these need to update per-block when we add multiblock tests:
	coinbase := sim.statedb.GetOrNewStateObject(*sim.getAddress(COINBASE))
	coinbase.SetGasPool(sim.gasLimit)

	origin := common.BytesToAddress(keyPair.Address())

	message := helper.NewMessage(
		origin,
		sim.getAddress(txn.To),
		txn.Data,
		txn.Value.AsBigInt(),
		txn.GasLimit.AsBigInt(),
		txn.GasPrice.AsBigInt())

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
	addr, ok := sim.addrmap[acct]
	if !ok {
		t := common.BytesToAddress([]byte(acct))
		addr = &t
		sim.addrmap[acct] = addr
	}
	return addr
}