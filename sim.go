package main

import (
	"errors"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/state"
	"github.com/ethereum/go-ethereum/tests/helper"
	"math/big"
)

type TestSim struct { // implements vm.Enviroment
	statedb  *state.StateDB
	coinbase []byte
	gasLimit *big.Int
	gas      *big.Int
}

func NewTestSim() (sim TestSim, err error) {
	var memdb *ethdb.MemDatabase

	memdb, err = ethdb.NewMemDatabase()
	if err != nil {
		return
	}

	statedb := state.New(nil, memdb)
	coinbase := make([]byte, 0)
	sim = TestSim{statedb, coinbase, big.NewInt(0), big.NewInt(0)}
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
	coinbase := sim.statedb.GetOrNewStateObject(sim.coinbase)
	coinbase.SetGasPool(sim.gasLimit)

	origin := keyPair.Address()

	message := helper.NewMessage(
		origin,
		[]byte(txn.To),
		txn.Data,
		txn.Value.AsBigInt(),
		txn.GasLimit.AsBigInt(),
		txn.GasPrice.AsBigInt())

	ret, gasLeft, err = core.ApplyMessage(sim, message, coinbase)
	if core.IsNonceErr(err) || core.IsInvalidTxErr(err) {
		sim.statedb.Set(snapshot)
	}
	sim.statedb.Update(sim.gas) // FIXME: What is this?

	return
}

func (sim *TestSim) checkAssertions(as *Assertions, result []byte, logs state.Logs, gasleft *big.Int) (successes uint, failures uint, err error) {
	err = errors.New("Not Implemented.")
	return
}
