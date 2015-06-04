package simulator

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/nejucomo/transactest/assert"
	"math/big"
)

// Implementation of vm.Environment for *testsim:
func (sim *testsim) State() *state.StateDB        { return sim.statedb }
func (sim *testsim) Origin() common.Address       { return *sim.getAddress(ORIGIN) }
func (sim *testsim) BlockNumber() *big.Int        { return sim.blocknum }
func (sim *testsim) GetHash(n uint64) common.Hash { return sim.blockhash }
func (sim *testsim) Coinbase() common.Address     { return *sim.getAddress(COINBASE) }
func (sim *testsim) Time() int64                  { return sim.time }
func (sim *testsim) Difficulty() *big.Int         { return sim.difficulty }
func (sim *testsim) GasLimit() *big.Int           { return sim.gasLimit }

func (sim *testsim) Transfer(from, to vm.Account, amount *big.Int) error {
	return vm.Transfer(from, to, amount)
}

func (sim *testsim) AddLog(log *state.Log) {
	assert.NotImplemented("<\n  %#v\n>.AddLog(log %#v)", sim, log)
}

func (sim *testsim) VmType() vm.Type { return vm.StdVmTy }
func (sim *testsim) Depth() int      { return sim.depth }
func (sim *testsim) SetDepth(i int)  { sim.depth = i }

func (sim *testsim) Call(caller vm.ContextRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	exec := core.NewExecution(sim, &addr, data, gas, price, value)
	ret, err := exec.Call(addr, caller)
	sim.gas = exec.Gas
	return ret, err
}

func (sim *testsim) CallCode(caller vm.ContextRef, addr common.Address, data []byte, gas, price, value *big.Int) ([]byte, error) {
	assert.NotImplemented("<\n  %#v\n>.CallCode(\n  caller %#v,\n  addr %#v,\n  data %#v,\n  gas %#v,\n  price %#v,\n  value %#v,\n  )", sim, caller, addr, data, gas, price, value)
	return nil, nil
}

func (sim *testsim) Create(caller vm.ContextRef, data []byte, gas, price, value *big.Int) ([]byte, error, vm.ContextRef) {
	assert.NotImplemented("<\n  %#v\n>.Create(\n  caller %#v,\n  data %#v,\n  gas %#v,\n  price %#v,\n  value %#v,\n  )", sim, caller, data, gas, price, value)
	return nil, nil, nil
}
