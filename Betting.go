// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BettingABI is the input ABI used to generate the binding from.
const BettingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Betting is an auto generated Go binding around an Ethereum contract.
type Betting struct {
	BettingCaller     // Read-only binding to the contract
	BettingTransactor // Write-only binding to the contract
	BettingFilterer   // Log filterer for contract events
}

// BettingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BettingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BettingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BettingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BettingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BettingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BettingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BettingSession struct {
	Contract     *Betting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BettingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BettingCallerSession struct {
	Contract *BettingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BettingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BettingTransactorSession struct {
	Contract     *BettingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BettingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BettingRaw struct {
	Contract *Betting // Generic contract binding to access the raw methods on
}

// BettingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BettingCallerRaw struct {
	Contract *BettingCaller // Generic read-only contract binding to access the raw methods on
}

// BettingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BettingTransactorRaw struct {
	Contract *BettingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBetting creates a new instance of Betting, bound to a specific deployed contract.
func NewBetting(address common.Address, backend bind.ContractBackend) (*Betting, error) {
	contract, err := bindBetting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Betting{BettingCaller: BettingCaller{contract: contract}, BettingTransactor: BettingTransactor{contract: contract}, BettingFilterer: BettingFilterer{contract: contract}}, nil
}

// NewBettingCaller creates a new read-only instance of Betting, bound to a specific deployed contract.
func NewBettingCaller(address common.Address, caller bind.ContractCaller) (*BettingCaller, error) {
	contract, err := bindBetting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BettingCaller{contract: contract}, nil
}

// NewBettingTransactor creates a new write-only instance of Betting, bound to a specific deployed contract.
func NewBettingTransactor(address common.Address, transactor bind.ContractTransactor) (*BettingTransactor, error) {
	contract, err := bindBetting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BettingTransactor{contract: contract}, nil
}

// NewBettingFilterer creates a new log filterer instance of Betting, bound to a specific deployed contract.
func NewBettingFilterer(address common.Address, filterer bind.ContractFilterer) (*BettingFilterer, error) {
	contract, err := bindBetting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BettingFilterer{contract: contract}, nil
}

// bindBetting binds a generic wrapper to an already deployed contract.
func bindBetting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BettingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting *BettingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Betting.Contract.BettingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting *BettingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.Contract.BettingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting *BettingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting.Contract.BettingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting *BettingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Betting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting *BettingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting *BettingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting *BettingCaller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Betting.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting *BettingSession) Version() (string, error) {
	return _Betting.Contract.Version(&_Betting.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting *BettingCallerSession) Version() (string, error) {
	return _Betting.Contract.Version(&_Betting.CallOpts)
}
