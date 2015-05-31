package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"math/big"
)

// Implementation of vm.Environment for *TestSim:
func (sim *TestSim) State() *state.StateDB        { return sim.statedb }
func (sim *TestSim) Origin() common.Address       { return *sim.getAddress(ORIGIN) }
func (sim *TestSim) BlockNumber() *big.Int        { return sim.blocknum }
func (sim *TestSim) GetHash(n uint64) common.Hash { return sim.blockhash }
func (sim *TestSim) Coinbase() common.Address     { return *sim.getAddress(COINBASE) }
func (sim *TestSim) Time() int64                  { return sim.time }
func (sim *TestSim) Difficulty() *big.Int         { return sim.difficulty }
func (sim *TestSim) GasLimit() *big.Int           { return sim.gasLimit }

func (sim *TestSim) Transfer(from, to vm.Account, amount *big.Int) error {
	not_implemented("TestSim<%#v>.Transfer(\n  from %#v,\n  to %#v,\n  amount %#v\n  )", from, to, amount)
	return nil
}

func (sim *TestSim) AddLog(log *state.Log) {
	not_implemented("TestSim<%#v>.AddLog(log %#v)", sim, log)
}

func (sim *TestSim) VmType() vm.Type {
	not_implemented("TestSim<%#v>.VmType()", sim)
	return vm.StdVmTy
}

func (sim *TestSim) Depth() int {
	not_implemented("TestSim<%#v>.Depth()", sim)
	return 0
}

func (sim *TestSim) SetDepth(i int) {
	not_implemented("TestSim<%#v>.SetDepth(i %#v)", sim, i)
	return
}

func (sim *TestSim) Call(caller vm.ContextRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	exec := core.NewExecution(sim, &addr, data, gas, price, value)
	ret, err := exec.Call(addr, caller)
	sim.gas = exec.Gas
	return ret, err
}

func (sim *TestSim) CallCode(caller vm.ContextRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	not_implemented("TestSim<%#v>.CallCode(\n  caller %#v,\n  addr %#v,\n  data %#v,\n  gas %#v,\n  price %#v,\n  value %#v,\n  )", sim, caller, addr, data, gas, price, value)
	return nil, nil
}

func (sim *TestSim) Create(caller vm.ContextRef, data []byte, gas, price, value *big.Int) ([]byte, error, vm.ContextRef) {
	not_implemented("TestSim<%#v>.Create(\n  caller %#v,\n  data %#v,\n  gas %#v,\n  price %#v,\n  value %#v,\n  )", sim, caller, data, gas, price, value)
	return nil, nil, nil
}
