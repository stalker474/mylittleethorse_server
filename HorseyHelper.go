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

// EthorseRaceABI is the input ABI used to generate the binding from.
const EthorseRaceABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"chronus\",\"outputs\":[{\"name\":\"betting_open\",\"type\":\"bool\"},{\"name\":\"race_start\",\"type\":\"bool\"},{\"name\":\"race_end\",\"type\":\"bool\"},{\"name\":\"voided_bet\",\"type\":\"bool\"},{\"name\":\"starting_time\",\"type\":\"uint32\"},{\"name\":\"betting_duration\",\"type\":\"uint32\"},{\"name\":\"race_duration\",\"type\":\"uint32\"},{\"name\":\"voided_timestamp\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EthorseRaceBin is the compiled bytecode used for deploying new contracts.
const EthorseRaceBin = `0x`

// DeployEthorseRace deploys a new Ethereum contract, binding an instance of EthorseRace to it.
func DeployEthorseRace(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthorseRace, error) {
	parsed, err := abi.JSON(strings.NewReader(EthorseRaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EthorseRaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthorseRace{EthorseRaceCaller: EthorseRaceCaller{contract: contract}, EthorseRaceTransactor: EthorseRaceTransactor{contract: contract}, EthorseRaceFilterer: EthorseRaceFilterer{contract: contract}}, nil
}

// EthorseRace is an auto generated Go binding around an Ethereum contract.
type EthorseRace struct {
	EthorseRaceCaller     // Read-only binding to the contract
	EthorseRaceTransactor // Write-only binding to the contract
	EthorseRaceFilterer   // Log filterer for contract events
}

// EthorseRaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthorseRaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthorseRaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthorseRaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthorseRaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthorseRaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthorseRaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthorseRaceSession struct {
	Contract     *EthorseRace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthorseRaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthorseRaceCallerSession struct {
	Contract *EthorseRaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EthorseRaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthorseRaceTransactorSession struct {
	Contract     *EthorseRaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EthorseRaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthorseRaceRaw struct {
	Contract *EthorseRace // Generic contract binding to access the raw methods on
}

// EthorseRaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthorseRaceCallerRaw struct {
	Contract *EthorseRaceCaller // Generic read-only contract binding to access the raw methods on
}

// EthorseRaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthorseRaceTransactorRaw struct {
	Contract *EthorseRaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthorseRace creates a new instance of EthorseRace, bound to a specific deployed contract.
func NewEthorseRace(address common.Address, backend bind.ContractBackend) (*EthorseRace, error) {
	contract, err := bindEthorseRace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthorseRace{EthorseRaceCaller: EthorseRaceCaller{contract: contract}, EthorseRaceTransactor: EthorseRaceTransactor{contract: contract}, EthorseRaceFilterer: EthorseRaceFilterer{contract: contract}}, nil
}

// NewEthorseRaceCaller creates a new read-only instance of EthorseRace, bound to a specific deployed contract.
func NewEthorseRaceCaller(address common.Address, caller bind.ContractCaller) (*EthorseRaceCaller, error) {
	contract, err := bindEthorseRace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthorseRaceCaller{contract: contract}, nil
}

// NewEthorseRaceTransactor creates a new write-only instance of EthorseRace, bound to a specific deployed contract.
func NewEthorseRaceTransactor(address common.Address, transactor bind.ContractTransactor) (*EthorseRaceTransactor, error) {
	contract, err := bindEthorseRace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthorseRaceTransactor{contract: contract}, nil
}

// NewEthorseRaceFilterer creates a new log filterer instance of EthorseRace, bound to a specific deployed contract.
func NewEthorseRaceFilterer(address common.Address, filterer bind.ContractFilterer) (*EthorseRaceFilterer, error) {
	contract, err := bindEthorseRace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthorseRaceFilterer{contract: contract}, nil
}

// bindEthorseRace binds a generic wrapper to an already deployed contract.
func bindEthorseRace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthorseRaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthorseRace *EthorseRaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthorseRace.Contract.EthorseRaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthorseRace *EthorseRaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthorseRace.Contract.EthorseRaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthorseRace *EthorseRaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthorseRace.Contract.EthorseRaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthorseRace *EthorseRaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EthorseRace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthorseRace *EthorseRaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthorseRace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthorseRace *EthorseRaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthorseRace.Contract.contract.Transact(opts, method, params...)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_EthorseRace *EthorseRaceCaller) Chronus(opts *bind.CallOpts) (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	ret := new(struct {
		BettingOpen     bool
		RaceStart       bool
		RaceEnd         bool
		VoidedBet       bool
		StartingTime    uint32
		BettingDuration uint32
		RaceDuration    uint32
		VoidedTimestamp uint32
	})
	out := ret
	err := _EthorseRace.contract.Call(opts, out, "chronus")
	return *ret, err
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_EthorseRace *EthorseRaceSession) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _EthorseRace.Contract.Chronus(&_EthorseRace.CallOpts)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_EthorseRace *EthorseRaceCallerSession) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _EthorseRace.Contract.Chronus(&_EthorseRace.CallOpts)
}

// HorseyHelperABI is the input ABI used to generate the binding from.
const HorseyHelperABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"newRace\",\"type\":\"address\"}],\"name\":\"isRefunded\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// HorseyHelperBin is the compiled bytecode used for deploying new contracts.
const HorseyHelperBin = `0x608060405234801561001057600080fd5b50610152806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663ffc9152e8114610045575b600080fd5b34801561005157600080fd5b5061007373ffffffffffffffffffffffffffffffffffffffff60043516610087565b604080519115158252519081900360200190f35b6000808273ffffffffffffffffffffffffffffffffffffffff166384304ee56040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161010060405180830381600087803b1580156100ef57600080fd5b505af1158015610103573d6000803e3d6000fd5b505050506040513d61010081101561011a57600080fd5b506060015193925050505600a165627a7a72305820ffd18e2ca546b851aa434122ad2916284907c886f729090e4d80917df9e7566b0029`

// DeployHorseyHelper deploys a new Ethereum contract, binding an instance of HorseyHelper to it.
func DeployHorseyHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HorseyHelper, error) {
	parsed, err := abi.JSON(strings.NewReader(HorseyHelperABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HorseyHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HorseyHelper{HorseyHelperCaller: HorseyHelperCaller{contract: contract}, HorseyHelperTransactor: HorseyHelperTransactor{contract: contract}, HorseyHelperFilterer: HorseyHelperFilterer{contract: contract}}, nil
}

// HorseyHelper is an auto generated Go binding around an Ethereum contract.
type HorseyHelper struct {
	HorseyHelperCaller     // Read-only binding to the contract
	HorseyHelperTransactor // Write-only binding to the contract
	HorseyHelperFilterer   // Log filterer for contract events
}

// HorseyHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type HorseyHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HorseyHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HorseyHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HorseyHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HorseyHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HorseyHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HorseyHelperSession struct {
	Contract     *HorseyHelper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HorseyHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HorseyHelperCallerSession struct {
	Contract *HorseyHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HorseyHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HorseyHelperTransactorSession struct {
	Contract     *HorseyHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HorseyHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type HorseyHelperRaw struct {
	Contract *HorseyHelper // Generic contract binding to access the raw methods on
}

// HorseyHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HorseyHelperCallerRaw struct {
	Contract *HorseyHelperCaller // Generic read-only contract binding to access the raw methods on
}

// HorseyHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HorseyHelperTransactorRaw struct {
	Contract *HorseyHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHorseyHelper creates a new instance of HorseyHelper, bound to a specific deployed contract.
func NewHorseyHelper(address common.Address, backend bind.ContractBackend) (*HorseyHelper, error) {
	contract, err := bindHorseyHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HorseyHelper{HorseyHelperCaller: HorseyHelperCaller{contract: contract}, HorseyHelperTransactor: HorseyHelperTransactor{contract: contract}, HorseyHelperFilterer: HorseyHelperFilterer{contract: contract}}, nil
}

// NewHorseyHelperCaller creates a new read-only instance of HorseyHelper, bound to a specific deployed contract.
func NewHorseyHelperCaller(address common.Address, caller bind.ContractCaller) (*HorseyHelperCaller, error) {
	contract, err := bindHorseyHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HorseyHelperCaller{contract: contract}, nil
}

// NewHorseyHelperTransactor creates a new write-only instance of HorseyHelper, bound to a specific deployed contract.
func NewHorseyHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*HorseyHelperTransactor, error) {
	contract, err := bindHorseyHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HorseyHelperTransactor{contract: contract}, nil
}

// NewHorseyHelperFilterer creates a new log filterer instance of HorseyHelper, bound to a specific deployed contract.
func NewHorseyHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*HorseyHelperFilterer, error) {
	contract, err := bindHorseyHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HorseyHelperFilterer{contract: contract}, nil
}

// bindHorseyHelper binds a generic wrapper to an already deployed contract.
func bindHorseyHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HorseyHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HorseyHelper *HorseyHelperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HorseyHelper.Contract.HorseyHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HorseyHelper *HorseyHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HorseyHelper.Contract.HorseyHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HorseyHelper *HorseyHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HorseyHelper.Contract.HorseyHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HorseyHelper *HorseyHelperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HorseyHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HorseyHelper *HorseyHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HorseyHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HorseyHelper *HorseyHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HorseyHelper.Contract.contract.Transact(opts, method, params...)
}

// IsRefunded is a free data retrieval call binding the contract method 0xffc9152e.
//
// Solidity: function isRefunded(newRace address) constant returns(bool)
func (_HorseyHelper *HorseyHelperCaller) IsRefunded(opts *bind.CallOpts, newRace common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _HorseyHelper.contract.Call(opts, out, "isRefunded", newRace)
	return *ret0, err
}

// IsRefunded is a free data retrieval call binding the contract method 0xffc9152e.
//
// Solidity: function isRefunded(newRace address) constant returns(bool)
func (_HorseyHelper *HorseyHelperSession) IsRefunded(newRace common.Address) (bool, error) {
	return _HorseyHelper.Contract.IsRefunded(&_HorseyHelper.CallOpts, newRace)
}

// IsRefunded is a free data retrieval call binding the contract method 0xffc9152e.
//
// Solidity: function isRefunded(newRace address) constant returns(bool)
func (_HorseyHelper *HorseyHelperCallerSession) IsRefunded(newRace common.Address) (bool, error) {
	return _HorseyHelper.Contract.IsRefunded(&_HorseyHelper.CallOpts, newRace)
}
