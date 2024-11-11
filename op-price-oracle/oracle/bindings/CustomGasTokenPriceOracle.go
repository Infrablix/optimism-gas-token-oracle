// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
	_ = abi.ConvertType
)

// CustomGasTokenPriceOracleMetaData contains all meta data concerning the CustomGasTokenPriceOracle contract.
var CustomGasTokenPriceOracleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"_priceInEth\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_lastUpdateTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastUpdateTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceInEth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"update\",\"inputs\":[{\"name\":\"_priceInEth\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// CustomGasTokenPriceOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use CustomGasTokenPriceOracleMetaData.ABI instead.
var CustomGasTokenPriceOracleABI = CustomGasTokenPriceOracleMetaData.ABI

// CustomGasTokenPriceOracle is an auto generated Go binding around an Ethereum contract.
type CustomGasTokenPriceOracle struct {
	CustomGasTokenPriceOracleCaller     // Read-only binding to the contract
	CustomGasTokenPriceOracleTransactor // Write-only binding to the contract
	CustomGasTokenPriceOracleFilterer   // Log filterer for contract events
}

// CustomGasTokenPriceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CustomGasTokenPriceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomGasTokenPriceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CustomGasTokenPriceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomGasTokenPriceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CustomGasTokenPriceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CustomGasTokenPriceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CustomGasTokenPriceOracleSession struct {
	Contract     *CustomGasTokenPriceOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CustomGasTokenPriceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CustomGasTokenPriceOracleCallerSession struct {
	Contract *CustomGasTokenPriceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// CustomGasTokenPriceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CustomGasTokenPriceOracleTransactorSession struct {
	Contract     *CustomGasTokenPriceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// CustomGasTokenPriceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CustomGasTokenPriceOracleRaw struct {
	Contract *CustomGasTokenPriceOracle // Generic contract binding to access the raw methods on
}

// CustomGasTokenPriceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CustomGasTokenPriceOracleCallerRaw struct {
	Contract *CustomGasTokenPriceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// CustomGasTokenPriceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CustomGasTokenPriceOracleTransactorRaw struct {
	Contract *CustomGasTokenPriceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCustomGasTokenPriceOracle creates a new instance of CustomGasTokenPriceOracle, bound to a specific deployed contract.
func NewCustomGasTokenPriceOracle(address common.Address, backend bind.ContractBackend) (*CustomGasTokenPriceOracle, error) {
	contract, err := bindCustomGasTokenPriceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CustomGasTokenPriceOracle{CustomGasTokenPriceOracleCaller: CustomGasTokenPriceOracleCaller{contract: contract}, CustomGasTokenPriceOracleTransactor: CustomGasTokenPriceOracleTransactor{contract: contract}, CustomGasTokenPriceOracleFilterer: CustomGasTokenPriceOracleFilterer{contract: contract}}, nil
}

// NewCustomGasTokenPriceOracleCaller creates a new read-only instance of CustomGasTokenPriceOracle, bound to a specific deployed contract.
func NewCustomGasTokenPriceOracleCaller(address common.Address, caller bind.ContractCaller) (*CustomGasTokenPriceOracleCaller, error) {
	contract, err := bindCustomGasTokenPriceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CustomGasTokenPriceOracleCaller{contract: contract}, nil
}

// NewCustomGasTokenPriceOracleTransactor creates a new write-only instance of CustomGasTokenPriceOracle, bound to a specific deployed contract.
func NewCustomGasTokenPriceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*CustomGasTokenPriceOracleTransactor, error) {
	contract, err := bindCustomGasTokenPriceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CustomGasTokenPriceOracleTransactor{contract: contract}, nil
}

// NewCustomGasTokenPriceOracleFilterer creates a new log filterer instance of CustomGasTokenPriceOracle, bound to a specific deployed contract.
func NewCustomGasTokenPriceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*CustomGasTokenPriceOracleFilterer, error) {
	contract, err := bindCustomGasTokenPriceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CustomGasTokenPriceOracleFilterer{contract: contract}, nil
}

// bindCustomGasTokenPriceOracle binds a generic wrapper to an already deployed contract.
func bindCustomGasTokenPriceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CustomGasTokenPriceOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomGasTokenPriceOracle.Contract.CustomGasTokenPriceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.CustomGasTokenPriceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.CustomGasTokenPriceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CustomGasTokenPriceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256 _priceInEth, uint256 _lastUpdateTimestamp)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCaller) GetPrice(opts *bind.CallOpts) (struct {
	PriceInEth          *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _CustomGasTokenPriceOracle.contract.Call(opts, &out, "getPrice")

	outstruct := new(struct {
		PriceInEth          *big.Int
		LastUpdateTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PriceInEth = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256 _priceInEth, uint256 _lastUpdateTimestamp)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleSession) GetPrice() (struct {
	PriceInEth          *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _CustomGasTokenPriceOracle.Contract.GetPrice(&_CustomGasTokenPriceOracle.CallOpts)
}

// GetPrice is a free data retrieval call binding the contract method 0x98d5fdca.
//
// Solidity: function getPrice() view returns(uint256 _priceInEth, uint256 _lastUpdateTimestamp)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCallerSession) GetPrice() (struct {
	PriceInEth          *big.Int
	LastUpdateTimestamp *big.Int
}, error) {
	return _CustomGasTokenPriceOracle.Contract.GetPrice(&_CustomGasTokenPriceOracle.CallOpts)
}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCaller) LastUpdateTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustomGasTokenPriceOracle.contract.Call(opts, &out, "lastUpdateTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleSession) LastUpdateTimestamp() (*big.Int, error) {
	return _CustomGasTokenPriceOracle.Contract.LastUpdateTimestamp(&_CustomGasTokenPriceOracle.CallOpts)
}

// LastUpdateTimestamp is a free data retrieval call binding the contract method 0x14bcec9f.
//
// Solidity: function lastUpdateTimestamp() view returns(uint256)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCallerSession) LastUpdateTimestamp() (*big.Int, error) {
	return _CustomGasTokenPriceOracle.Contract.LastUpdateTimestamp(&_CustomGasTokenPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CustomGasTokenPriceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleSession) Owner() (common.Address, error) {
	return _CustomGasTokenPriceOracle.Contract.Owner(&_CustomGasTokenPriceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCallerSession) Owner() (common.Address, error) {
	return _CustomGasTokenPriceOracle.Contract.Owner(&_CustomGasTokenPriceOracle.CallOpts)
}

// PriceInEth is a free data retrieval call binding the contract method 0xfdd760ea.
//
// Solidity: function priceInEth() view returns(uint256)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCaller) PriceInEth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CustomGasTokenPriceOracle.contract.Call(opts, &out, "priceInEth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceInEth is a free data retrieval call binding the contract method 0xfdd760ea.
//
// Solidity: function priceInEth() view returns(uint256)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleSession) PriceInEth() (*big.Int, error) {
	return _CustomGasTokenPriceOracle.Contract.PriceInEth(&_CustomGasTokenPriceOracle.CallOpts)
}

// PriceInEth is a free data retrieval call binding the contract method 0xfdd760ea.
//
// Solidity: function priceInEth() view returns(uint256)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCallerSession) PriceInEth() (*big.Int, error) {
	return _CustomGasTokenPriceOracle.Contract.PriceInEth(&_CustomGasTokenPriceOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CustomGasTokenPriceOracle.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleSession) Version() (string, error) {
	return _CustomGasTokenPriceOracle.Contract.Version(&_CustomGasTokenPriceOracle.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleCallerSession) Version() (string, error) {
	return _CustomGasTokenPriceOracle.Contract.Version(&_CustomGasTokenPriceOracle.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.RenounceOwnership(&_CustomGasTokenPriceOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.RenounceOwnership(&_CustomGasTokenPriceOracle.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.TransferOwnership(&_CustomGasTokenPriceOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.TransferOwnership(&_CustomGasTokenPriceOracle.TransactOpts, newOwner)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 _priceInEth) returns()
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleTransactor) Update(opts *bind.TransactOpts, _priceInEth *big.Int) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.contract.Transact(opts, "update", _priceInEth)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 _priceInEth) returns()
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleSession) Update(_priceInEth *big.Int) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.Update(&_CustomGasTokenPriceOracle.TransactOpts, _priceInEth)
}

// Update is a paid mutator transaction binding the contract method 0x82ab890a.
//
// Solidity: function update(uint256 _priceInEth) returns()
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleTransactorSession) Update(_priceInEth *big.Int) (*types.Transaction, error) {
	return _CustomGasTokenPriceOracle.Contract.Update(&_CustomGasTokenPriceOracle.TransactOpts, _priceInEth)
}

// CustomGasTokenPriceOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CustomGasTokenPriceOracle contract.
type CustomGasTokenPriceOracleOwnershipTransferredIterator struct {
	Event *CustomGasTokenPriceOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CustomGasTokenPriceOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CustomGasTokenPriceOracleOwnershipTransferred)
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
		it.Event = new(CustomGasTokenPriceOracleOwnershipTransferred)
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
func (it *CustomGasTokenPriceOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CustomGasTokenPriceOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CustomGasTokenPriceOracleOwnershipTransferred represents a OwnershipTransferred event raised by the CustomGasTokenPriceOracle contract.
type CustomGasTokenPriceOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CustomGasTokenPriceOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustomGasTokenPriceOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CustomGasTokenPriceOracleOwnershipTransferredIterator{contract: _CustomGasTokenPriceOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CustomGasTokenPriceOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CustomGasTokenPriceOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CustomGasTokenPriceOracleOwnershipTransferred)
				if err := _CustomGasTokenPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CustomGasTokenPriceOracle *CustomGasTokenPriceOracleFilterer) ParseOwnershipTransferred(log types.Log) (*CustomGasTokenPriceOracleOwnershipTransferred, error) {
	event := new(CustomGasTokenPriceOracleOwnershipTransferred)
	if err := _CustomGasTokenPriceOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
