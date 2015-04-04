package main

import (
	"errors"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/state"
	"github.com/ethereum/go-ethereum/vm"
	"math/big"
)

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
