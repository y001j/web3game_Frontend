// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MetaCoinMetaData contains all meta data concerning the MetaCoin contract.
var MetaCoinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TransEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalanceInEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEthBalanceInContract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"sufficient\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferEth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MetaCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use MetaCoinMetaData.ABI instead.
var MetaCoinABI = MetaCoinMetaData.ABI

// MetaCoin is an auto generated Go binding around an Ethereum contract.
type MetaCoin struct {
	MetaCoinCaller     // Read-only binding to the contract
	MetaCoinTransactor // Write-only binding to the contract
	MetaCoinFilterer   // Log filterer for contract events
}

// MetaCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaCoinSession struct {
	Contract     *MetaCoin         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaCoinCallerSession struct {
	Contract *MetaCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MetaCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaCoinTransactorSession struct {
	Contract     *MetaCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MetaCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaCoinRaw struct {
	Contract *MetaCoin // Generic contract binding to access the raw methods on
}

// MetaCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaCoinCallerRaw struct {
	Contract *MetaCoinCaller // Generic read-only contract binding to access the raw methods on
}

// MetaCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaCoinTransactorRaw struct {
	Contract *MetaCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaCoin creates a new instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoin(address common.Address, backend bind.ContractBackend) (*MetaCoin, error) {
	contract, err := bindMetaCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaCoin{MetaCoinCaller: MetaCoinCaller{contract: contract}, MetaCoinTransactor: MetaCoinTransactor{contract: contract}, MetaCoinFilterer: MetaCoinFilterer{contract: contract}}, nil
}

// NewMetaCoinCaller creates a new read-only instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoinCaller(address common.Address, caller bind.ContractCaller) (*MetaCoinCaller, error) {
	contract, err := bindMetaCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaCoinCaller{contract: contract}, nil
}

// NewMetaCoinTransactor creates a new write-only instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaCoinTransactor, error) {
	contract, err := bindMetaCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaCoinTransactor{contract: contract}, nil
}

// NewMetaCoinFilterer creates a new log filterer instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaCoinFilterer, error) {
	contract, err := bindMetaCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaCoinFilterer{contract: contract}, nil
}

// bindMetaCoin binds a generic wrapper to an already deployed contract.
func bindMetaCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaCoin *MetaCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaCoin.Contract.MetaCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaCoin *MetaCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaCoin.Contract.MetaCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaCoin *MetaCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaCoin.Contract.MetaCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaCoin *MetaCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetaCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaCoin *MetaCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaCoin *MetaCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaCoin.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_MetaCoin *MetaCoinCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaCoin.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_MetaCoin *MetaCoinSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.GetBalance(&_MetaCoin.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_MetaCoin *MetaCoinCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.GetBalance(&_MetaCoin.CallOpts, addr)
}

// GetBalanceInEth is a free data retrieval call binding the contract method 0x7bd703e8.
//
// Solidity: function getBalanceInEth(address addr) view returns(uint256)
func (_MetaCoin *MetaCoinCaller) GetBalanceInEth(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaCoin.contract.Call(opts, &out, "getBalanceInEth", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceInEth is a free data retrieval call binding the contract method 0x7bd703e8.
//
// Solidity: function getBalanceInEth(address addr) view returns(uint256)
func (_MetaCoin *MetaCoinSession) GetBalanceInEth(addr common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.GetBalanceInEth(&_MetaCoin.CallOpts, addr)
}

// GetBalanceInEth is a free data retrieval call binding the contract method 0x7bd703e8.
//
// Solidity: function getBalanceInEth(address addr) view returns(uint256)
func (_MetaCoin *MetaCoinCallerSession) GetBalanceInEth(addr common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.GetBalanceInEth(&_MetaCoin.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256)
func (_MetaCoin *MetaCoinCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MetaCoin.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256)
func (_MetaCoin *MetaCoinSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.GetEthBalance(&_MetaCoin.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256)
func (_MetaCoin *MetaCoinCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.GetEthBalance(&_MetaCoin.CallOpts, addr)
}

// GetEthBalanceInContract is a free data retrieval call binding the contract method 0xba08e975.
//
// Solidity: function getEthBalanceInContract() view returns(uint256)
func (_MetaCoin *MetaCoinCaller) GetEthBalanceInContract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MetaCoin.contract.Call(opts, &out, "getEthBalanceInContract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalanceInContract is a free data retrieval call binding the contract method 0xba08e975.
//
// Solidity: function getEthBalanceInContract() view returns(uint256)
func (_MetaCoin *MetaCoinSession) GetEthBalanceInContract() (*big.Int, error) {
	return _MetaCoin.Contract.GetEthBalanceInContract(&_MetaCoin.CallOpts)
}

// GetEthBalanceInContract is a free data retrieval call binding the contract method 0xba08e975.
//
// Solidity: function getEthBalanceInContract() view returns(uint256)
func (_MetaCoin *MetaCoinCallerSession) GetEthBalanceInContract() (*big.Int, error) {
	return _MetaCoin.Contract.GetEthBalanceInContract(&_MetaCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaCoin *MetaCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MetaCoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaCoin *MetaCoinSession) Owner() (common.Address, error) {
	return _MetaCoin.Contract.Owner(&_MetaCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MetaCoin *MetaCoinCallerSession) Owner() (common.Address, error) {
	return _MetaCoin.Contract.Owner(&_MetaCoin.CallOpts)
}

// SendCoin is a paid mutator transaction binding the contract method 0x90b98a11.
//
// Solidity: function sendCoin(address receiver, uint256 amount) returns(bool sufficient)
func (_MetaCoin *MetaCoinTransactor) SendCoin(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.contract.Transact(opts, "sendCoin", receiver, amount)
}

// SendCoin is a paid mutator transaction binding the contract method 0x90b98a11.
//
// Solidity: function sendCoin(address receiver, uint256 amount) returns(bool sufficient)
func (_MetaCoin *MetaCoinSession) SendCoin(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.Contract.SendCoin(&_MetaCoin.TransactOpts, receiver, amount)
}

// SendCoin is a paid mutator transaction binding the contract method 0x90b98a11.
//
// Solidity: function sendCoin(address receiver, uint256 amount) returns(bool sufficient)
func (_MetaCoin *MetaCoinTransactorSession) SendCoin(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.Contract.SendCoin(&_MetaCoin.TransactOpts, receiver, amount)
}

// SendEth is a paid mutator transaction binding the contract method 0x06e99fef.
//
// Solidity: function sendEth() payable returns()
func (_MetaCoin *MetaCoinTransactor) SendEth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaCoin.contract.Transact(opts, "sendEth")
}

// SendEth is a paid mutator transaction binding the contract method 0x06e99fef.
//
// Solidity: function sendEth() payable returns()
func (_MetaCoin *MetaCoinSession) SendEth() (*types.Transaction, error) {
	return _MetaCoin.Contract.SendEth(&_MetaCoin.TransactOpts)
}

// SendEth is a paid mutator transaction binding the contract method 0x06e99fef.
//
// Solidity: function sendEth() payable returns()
func (_MetaCoin *MetaCoinTransactorSession) SendEth() (*types.Transaction, error) {
	return _MetaCoin.Contract.SendEth(&_MetaCoin.TransactOpts)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(address recipient, uint256 amount) returns(bool)
func (_MetaCoin *MetaCoinTransactor) TransferEth(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.contract.Transact(opts, "transferEth", recipient, amount)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(address recipient, uint256 amount) returns(bool)
func (_MetaCoin *MetaCoinSession) TransferEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.Contract.TransferEth(&_MetaCoin.TransactOpts, recipient, amount)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(address recipient, uint256 amount) returns(bool)
func (_MetaCoin *MetaCoinTransactorSession) TransferEth(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.Contract.TransferEth(&_MetaCoin.TransactOpts, recipient, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MetaCoin *MetaCoinTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MetaCoin.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MetaCoin *MetaCoinSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MetaCoin.Contract.Fallback(&_MetaCoin.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MetaCoin *MetaCoinTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MetaCoin.Contract.Fallback(&_MetaCoin.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetaCoin *MetaCoinTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaCoin.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetaCoin *MetaCoinSession) Receive() (*types.Transaction, error) {
	return _MetaCoin.Contract.Receive(&_MetaCoin.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MetaCoin *MetaCoinTransactorSession) Receive() (*types.Transaction, error) {
	return _MetaCoin.Contract.Receive(&_MetaCoin.TransactOpts)
}

// MetaCoinTransEventIterator is returned from FilterTransEvent and is used to iterate over the raw logs and unpacked data for TransEvent events raised by the MetaCoin contract.
type MetaCoinTransEventIterator struct {
	Event *MetaCoinTransEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaCoinTransEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaCoinTransEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaCoinTransEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaCoinTransEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaCoinTransEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaCoinTransEvent represents a TransEvent event raised by the MetaCoin contract.
type MetaCoinTransEvent struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransEvent is a free log retrieval operation binding the contract event 0x5e3e3ac48142223910fa7d0c0fb4e15e6aac9c9cbc7a8dc20aa5dc9227d28683.
//
// Solidity: event TransEvent(address arg0, uint256 arg1)
func (_MetaCoin *MetaCoinFilterer) FilterTransEvent(opts *bind.FilterOpts) (*MetaCoinTransEventIterator, error) {

	logs, sub, err := _MetaCoin.contract.FilterLogs(opts, "TransEvent")
	if err != nil {
		return nil, err
	}
	return &MetaCoinTransEventIterator{contract: _MetaCoin.contract, event: "TransEvent", logs: logs, sub: sub}, nil
}

// WatchTransEvent is a free log subscription operation binding the contract event 0x5e3e3ac48142223910fa7d0c0fb4e15e6aac9c9cbc7a8dc20aa5dc9227d28683.
//
// Solidity: event TransEvent(address arg0, uint256 arg1)
func (_MetaCoin *MetaCoinFilterer) WatchTransEvent(opts *bind.WatchOpts, sink chan<- *MetaCoinTransEvent) (event.Subscription, error) {

	logs, sub, err := _MetaCoin.contract.WatchLogs(opts, "TransEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaCoinTransEvent)
				if err := _MetaCoin.contract.UnpackLog(event, "TransEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransEvent is a log parse operation binding the contract event 0x5e3e3ac48142223910fa7d0c0fb4e15e6aac9c9cbc7a8dc20aa5dc9227d28683.
//
// Solidity: event TransEvent(address arg0, uint256 arg1)
func (_MetaCoin *MetaCoinFilterer) ParseTransEvent(log types.Log) (*MetaCoinTransEvent, error) {
	event := new(MetaCoinTransEvent)
	if err := _MetaCoin.contract.UnpackLog(event, "TransEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MetaCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MetaCoin contract.
type MetaCoinTransferIterator struct {
	Event *MetaCoinTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MetaCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetaCoinTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MetaCoinTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MetaCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetaCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetaCoinTransfer represents a Transfer event raised by the MetaCoin contract.
type MetaCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_MetaCoin *MetaCoinFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*MetaCoinTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _MetaCoin.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &MetaCoinTransferIterator{contract: _MetaCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_MetaCoin *MetaCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MetaCoinTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _MetaCoin.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetaCoinTransfer)
				if err := _MetaCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_MetaCoin *MetaCoinFilterer) ParseTransfer(log types.Log) (*MetaCoinTransfer, error) {
	event := new(MetaCoinTransfer)
	if err := _MetaCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
