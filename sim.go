package main

import (
	"errors"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/state"
	"github.com/ethereum/go-ethereum/tests/helper"
	"github.com/ethereum/go-ethereum/vm"
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

// Implementation of vm.Environment for *TestSim:
func (sim *TestSim) State() *state.StateDB   { return sim.statedb }
func (sim *TestSim) Origin() []byte          { return make([]byte, 0) } // FIXME stub
func (sim *TestSim) BlockNumber() *big.Int   { return big.NewInt(0) }   // FIXME stub
func (sim *TestSim) GetHash(n uint64) []byte { return make([]byte, 0) } // FIXME stub
func (sim *TestSim) Coinbase() []byte        { return make([]byte, 0) } // FIXME stub
func (sim *TestSim) Time() int64             { return 0 }               // FIXME stub
func (sim *TestSim) Difficulty() *big.Int    { return big.NewInt(0) }   // FIXME stub
func (sim *TestSim) GasLimit() *big.Int      { return big.NewInt(0) }   // FIXME stub

func (sim *TestSim) Transfer(from, to vm.Account, amount *big.Int) error {
	// FIXME stub
	return nil
}

func (sim *TestSim) AddLog(state.Log) {
	// FIXME stub
}

func (sim *TestSim) VmType() vm.Type { return vm.StdVmTy } // FIXME stub
func (sim *TestSim) Depth() int      { return 0 }          // FIXME stub
func (sim *TestSim) SetDepth(i int)  { return }            // FIXME stub

func (sim *TestSim) Call(caller vm.ContextRef, addr, data []byte, gas, price, value *big.Int) ([]byte, error) {
	exec := core.NewExecution(sim, addr, data, gas, price, value)
	ret, err := exec.Call(addr, caller)
	sim.gas = exec.Gas
	return ret, err
}

func (sim *TestSim) CallCode(caller vm.ContextRef, addr, data []byte, gas, price, value *big.Int) ([]byte, error) {
	// FIXME stub
	return nil, errors.New("Not Implemented.")
}

func (sim *TestSim) Create(caller vm.ContextRef, addr, data []byte, gas, price, value *big.Int) ([]byte, error, vm.ContextRef) {
	// FIXME stub
	return nil, errors.New("Not Implemented."), nil
}
