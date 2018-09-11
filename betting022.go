// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Betting022ABI is the input ABI used to generate the binding from.
const Betting022ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"horse\",\"type\":\"bytes32\"}],\"name\":\"placeBet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim_reward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"winner_horse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"myid\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winnerPoolTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"myid\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"string\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"__callback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"horses\",\"outputs\":[{\"name\":\"BTC_delta\",\"type\":\"int32\"},{\"name\":\"ETH_delta\",\"type\":\"int32\"},{\"name\":\"LTC_delta\",\"type\":\"int32\"},{\"name\":\"BTC\",\"type\":\"bytes32\"},{\"name\":\"ETH\",\"type\":\"bytes32\"},{\"name\":\"LTC\",\"type\":\"bytes32\"},{\"name\":\"customGasLimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"bytes32\"},{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"getCoinIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chronus\",\"outputs\":[{\"name\":\"betting_open\",\"type\":\"bool\"},{\"name\":\"race_start\",\"type\":\"bool\"},{\"name\":\"race_end\",\"type\":\"bool\"},{\"name\":\"voided_bet\",\"type\":\"bool\"},{\"name\":\"starting_time\",\"type\":\"uint32\"},{\"name\":\"betting_duration\",\"type\":\"uint32\"},{\"name\":\"race_duration\",\"type\":\"uint32\"},{\"name\":\"voided_timestamp\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reward_total\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delay\",\"type\":\"uint256\"},{\"name\":\"locking_duration\",\"type\":\"uint256\"}],\"name\":\"setupRace\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total_reward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"recovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"description\",\"type\":\"string\"}],\"name\":\"newOraclizeQuery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"newPriceTicker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_horse\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_date\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// Betting022 is an auto generated Go binding around an Ethereum contract.
type Betting022 struct {
	Betting022Caller     // Read-only binding to the contract
	Betting022Transactor // Write-only binding to the contract
	Betting022Filterer   // Log filterer for contract events
}

// Betting022Caller is an auto generated read-only Go binding around an Ethereum contract.
type Betting022Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Betting022Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Betting022Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Betting022Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Betting022Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Betting022Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Betting022Session struct {
	Contract     *Betting022       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Betting022CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Betting022CallerSession struct {
	Contract *Betting022Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Betting022TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Betting022TransactorSession struct {
	Contract     *Betting022Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Betting022Raw is an auto generated low-level Go binding around an Ethereum contract.
type Betting022Raw struct {
	Contract *Betting022 // Generic contract binding to access the raw methods on
}

// Betting022CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Betting022CallerRaw struct {
	Contract *Betting022Caller // Generic read-only contract binding to access the raw methods on
}

// Betting022TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Betting022TransactorRaw struct {
	Contract *Betting022Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBetting022 creates a new instance of Betting022, bound to a specific deployed contract.
func NewBetting022(address common.Address, backend bind.ContractBackend) (*Betting022, error) {
	contract, err := bindBetting022(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Betting022{Betting022Caller: Betting022Caller{contract: contract}, Betting022Transactor: Betting022Transactor{contract: contract}, Betting022Filterer: Betting022Filterer{contract: contract}}, nil
}

// NewBetting022Caller creates a new read-only instance of Betting022, bound to a specific deployed contract.
func NewBetting022Caller(address common.Address, caller bind.ContractCaller) (*Betting022Caller, error) {
	contract, err := bindBetting022(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Betting022Caller{contract: contract}, nil
}

// NewBetting022Transactor creates a new write-only instance of Betting022, bound to a specific deployed contract.
func NewBetting022Transactor(address common.Address, transactor bind.ContractTransactor) (*Betting022Transactor, error) {
	contract, err := bindBetting022(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Betting022Transactor{contract: contract}, nil
}

// NewBetting022Filterer creates a new log filterer instance of Betting022, bound to a specific deployed contract.
func NewBetting022Filterer(address common.Address, filterer bind.ContractFilterer) (*Betting022Filterer, error) {
	contract, err := bindBetting022(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Betting022Filterer{contract: contract}, nil
}

// bindBetting022 binds a generic wrapper to an already deployed contract.
func bindBetting022(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Betting022ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting022 *Betting022Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Betting022.Contract.Betting022Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting022 *Betting022Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting022.Contract.Betting022Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting022 *Betting022Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting022.Contract.Betting022Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting022 *Betting022CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Betting022.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting022 *Betting022TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting022.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting022 *Betting022TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting022.Contract.contract.Transact(opts, method, params...)
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting022 *Betting022Caller) CheckReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting022.contract.Call(opts, out, "checkReward")
	return *ret0, err
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting022 *Betting022Session) CheckReward() (*big.Int, error) {
	return _Betting022.Contract.CheckReward(&_Betting022.CallOpts)
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting022 *Betting022CallerSession) CheckReward() (*big.Int, error) {
	return _Betting022.Contract.CheckReward(&_Betting022.CallOpts)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting022 *Betting022Caller) Chronus(opts *bind.CallOpts) (struct {
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
	err := _Betting022.contract.Call(opts, out, "chronus")
	return *ret, err
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting022 *Betting022Session) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _Betting022.Contract.Chronus(&_Betting022.CallOpts)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting022 *Betting022CallerSession) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _Betting022.Contract.Chronus(&_Betting022.CallOpts)
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting022 *Betting022Caller) GetCoinIndex(opts *bind.CallOpts, index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(bool)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _Betting022.contract.Call(opts, out, "getCoinIndex", index, candidate)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting022 *Betting022Session) GetCoinIndex(index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Betting022.Contract.GetCoinIndex(&_Betting022.CallOpts, index, candidate)
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting022 *Betting022CallerSession) GetCoinIndex(index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Betting022.Contract.GetCoinIndex(&_Betting022.CallOpts, index, candidate)
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int32, ETH_delta int32, LTC_delta int32, BTC bytes32, ETH bytes32, LTC bytes32, customGasLimit uint256)
func (_Betting022 *Betting022Caller) Horses(opts *bind.CallOpts) (struct {
	BTCDelta       int32
	ETHDelta       int32
	LTCDelta       int32
	BTC            [32]byte
	ETH            [32]byte
	LTC            [32]byte
	CustomGasLimit *big.Int
}, error) {
	ret := new(struct {
		BTCDelta       int32
		ETHDelta       int32
		LTCDelta       int32
		BTC            [32]byte
		ETH            [32]byte
		LTC            [32]byte
		CustomGasLimit *big.Int
	})
	out := ret
	err := _Betting022.contract.Call(opts, out, "horses")
	return *ret, err
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int32, ETH_delta int32, LTC_delta int32, BTC bytes32, ETH bytes32, LTC bytes32, customGasLimit uint256)
func (_Betting022 *Betting022Session) Horses() (struct {
	BTCDelta       int32
	ETHDelta       int32
	LTCDelta       int32
	BTC            [32]byte
	ETH            [32]byte
	LTC            [32]byte
	CustomGasLimit *big.Int
}, error) {
	return _Betting022.Contract.Horses(&_Betting022.CallOpts)
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int32, ETH_delta int32, LTC_delta int32, BTC bytes32, ETH bytes32, LTC bytes32, customGasLimit uint256)
func (_Betting022 *Betting022CallerSession) Horses() (struct {
	BTCDelta       int32
	ETHDelta       int32
	LTCDelta       int32
	BTC            [32]byte
	ETH            [32]byte
	LTC            [32]byte
	CustomGasLimit *big.Int
}, error) {
	return _Betting022.Contract.Horses(&_Betting022.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting022 *Betting022Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Betting022.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting022 *Betting022Session) Owner() (common.Address, error) {
	return _Betting022.Contract.Owner(&_Betting022.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting022 *Betting022CallerSession) Owner() (common.Address, error) {
	return _Betting022.Contract.Owner(&_Betting022.CallOpts)
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting022 *Betting022Caller) RewardTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting022.contract.Call(opts, out, "reward_total")
	return *ret0, err
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting022 *Betting022Session) RewardTotal() (*big.Int, error) {
	return _Betting022.Contract.RewardTotal(&_Betting022.CallOpts)
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting022 *Betting022CallerSession) RewardTotal() (*big.Int, error) {
	return _Betting022.Contract.RewardTotal(&_Betting022.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting022 *Betting022Caller) TotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting022.contract.Call(opts, out, "total_reward")
	return *ret0, err
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting022 *Betting022Session) TotalReward() (*big.Int, error) {
	return _Betting022.Contract.TotalReward(&_Betting022.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting022 *Betting022CallerSession) TotalReward() (*big.Int, error) {
	return _Betting022.Contract.TotalReward(&_Betting022.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting022 *Betting022Caller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Betting022.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting022 *Betting022Session) Version() (string, error) {
	return _Betting022.Contract.Version(&_Betting022.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting022 *Betting022CallerSession) Version() (string, error) {
	return _Betting022.Contract.Version(&_Betting022.CallOpts)
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting022 *Betting022Caller) WinnerPoolTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting022.contract.Call(opts, out, "winnerPoolTotal")
	return *ret0, err
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting022 *Betting022Session) WinnerPoolTotal() (*big.Int, error) {
	return _Betting022.Contract.WinnerPoolTotal(&_Betting022.CallOpts)
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting022 *Betting022CallerSession) WinnerPoolTotal() (*big.Int, error) {
	return _Betting022.Contract.WinnerPoolTotal(&_Betting022.CallOpts)
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting022 *Betting022Caller) WinnerHorse(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Betting022.contract.Call(opts, out, "winner_horse", arg0)
	return *ret0, err
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting022 *Betting022Session) WinnerHorse(arg0 [32]byte) (bool, error) {
	return _Betting022.Contract.WinnerHorse(&_Betting022.CallOpts, arg0)
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting022 *Betting022CallerSession) WinnerHorse(arg0 [32]byte) (bool, error) {
	return _Betting022.Contract.WinnerHorse(&_Betting022.CallOpts, arg0)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_Betting022 *Betting022Transactor) Callback(opts *bind.TransactOpts, myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _Betting022.contract.Transact(opts, "__callback", myid, result, proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_Betting022 *Betting022Session) Callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _Betting022.Contract.Callback(&_Betting022.TransactOpts, myid, result, proof)
}

// Callback is a paid mutator transaction binding the contract method 0x38bbfa50.
//
// Solidity: function __callback(myid bytes32, result string, proof bytes) returns()
func (_Betting022 *Betting022TransactorSession) Callback(myid [32]byte, result string, proof []byte) (*types.Transaction, error) {
	return _Betting022.Contract.Callback(&_Betting022.TransactOpts, myid, result, proof)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting022 *Betting022Transactor) ChangeOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Betting022.contract.Transact(opts, "changeOwnership", _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting022 *Betting022Session) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Betting022.Contract.ChangeOwnership(&_Betting022.TransactOpts, _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting022 *Betting022TransactorSession) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Betting022.Contract.ChangeOwnership(&_Betting022.TransactOpts, _newOwner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting022 *Betting022Transactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting022.contract.Transact(opts, "claim_reward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting022 *Betting022Session) ClaimReward() (*types.Transaction, error) {
	return _Betting022.Contract.ClaimReward(&_Betting022.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting022 *Betting022TransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Betting022.Contract.ClaimReward(&_Betting022.TransactOpts)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting022 *Betting022Transactor) PlaceBet(opts *bind.TransactOpts, horse [32]byte) (*types.Transaction, error) {
	return _Betting022.contract.Transact(opts, "placeBet", horse)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting022 *Betting022Session) PlaceBet(horse [32]byte) (*types.Transaction, error) {
	return _Betting022.Contract.PlaceBet(&_Betting022.TransactOpts, horse)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting022 *Betting022TransactorSession) PlaceBet(horse [32]byte) (*types.Transaction, error) {
	return _Betting022.Contract.PlaceBet(&_Betting022.TransactOpts, horse)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting022 *Betting022Transactor) Recovery(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting022.contract.Transact(opts, "recovery")
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting022 *Betting022Session) Recovery() (*types.Transaction, error) {
	return _Betting022.Contract.Recovery(&_Betting022.TransactOpts)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting022 *Betting022TransactorSession) Recovery() (*types.Transaction, error) {
	return _Betting022.Contract.Recovery(&_Betting022.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting022 *Betting022Transactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting022.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting022 *Betting022Session) Refund() (*types.Transaction, error) {
	return _Betting022.Contract.Refund(&_Betting022.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting022 *Betting022TransactorSession) Refund() (*types.Transaction, error) {
	return _Betting022.Contract.Refund(&_Betting022.TransactOpts)
}

// SetupRace is a paid mutator transaction binding the contract method 0xaff6b3e8.
//
// Solidity: function setupRace(delay uint256, locking_duration uint256) returns(bool)
func (_Betting022 *Betting022Transactor) SetupRace(opts *bind.TransactOpts, delay *big.Int, locking_duration *big.Int) (*types.Transaction, error) {
	return _Betting022.contract.Transact(opts, "setupRace", delay, locking_duration)
}

// SetupRace is a paid mutator transaction binding the contract method 0xaff6b3e8.
//
// Solidity: function setupRace(delay uint256, locking_duration uint256) returns(bool)
func (_Betting022 *Betting022Session) SetupRace(delay *big.Int, locking_duration *big.Int) (*types.Transaction, error) {
	return _Betting022.Contract.SetupRace(&_Betting022.TransactOpts, delay, locking_duration)
}

// SetupRace is a paid mutator transaction binding the contract method 0xaff6b3e8.
//
// Solidity: function setupRace(delay uint256, locking_duration uint256) returns(bool)
func (_Betting022 *Betting022TransactorSession) SetupRace(delay *big.Int, locking_duration *big.Int) (*types.Transaction, error) {
	return _Betting022.Contract.SetupRace(&_Betting022.TransactOpts, delay, locking_duration)
}

// Betting022DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Betting022 contract.
type Betting022DepositIterator struct {
	Event *Betting022Deposit // Event containing the contract specifics and raw log

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
func (it *Betting022DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting022Deposit)
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
		it.Event = new(Betting022Deposit)
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
func (it *Betting022DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting022DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting022Deposit represents a Deposit event raised by the Betting022 contract.
type Betting022Deposit struct {
	From  common.Address
	Value *big.Int
	Horse [32]byte
	Date  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x60452eb7177e8d41c9d9fbc4c6e9ccf55a4d44d412355fbf2f02668e0d1a0ce1.
//
// Solidity: e Deposit(_from address, _value uint256, _horse bytes32, _date uint256)
func (_Betting022 *Betting022Filterer) FilterDeposit(opts *bind.FilterOpts) (*Betting022DepositIterator, error) {

	logs, sub, err := _Betting022.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &Betting022DepositIterator{contract: _Betting022.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x60452eb7177e8d41c9d9fbc4c6e9ccf55a4d44d412355fbf2f02668e0d1a0ce1.
//
// Solidity: e Deposit(_from address, _value uint256, _horse bytes32, _date uint256)
func (_Betting022 *Betting022Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Betting022Deposit) (event.Subscription, error) {

	logs, sub, err := _Betting022.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting022Deposit)
				if err := _Betting022.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// Betting022WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Betting022 contract.
type Betting022WithdrawIterator struct {
	Event *Betting022Withdraw // Event containing the contract specifics and raw log

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
func (it *Betting022WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting022Withdraw)
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
		it.Event = new(Betting022Withdraw)
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
func (it *Betting022WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting022WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting022Withdraw represents a Withdraw event raised by the Betting022 contract.
type Betting022Withdraw struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_to address, _value uint256)
func (_Betting022 *Betting022Filterer) FilterWithdraw(opts *bind.FilterOpts) (*Betting022WithdrawIterator, error) {

	logs, sub, err := _Betting022.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &Betting022WithdrawIterator{contract: _Betting022.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_to address, _value uint256)
func (_Betting022 *Betting022Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Betting022Withdraw) (event.Subscription, error) {

	logs, sub, err := _Betting022.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting022Withdraw)
				if err := _Betting022.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// Betting022NewOraclizeQueryIterator is returned from FilterNewOraclizeQuery and is used to iterate over the raw logs and unpacked data for NewOraclizeQuery events raised by the Betting022 contract.
type Betting022NewOraclizeQueryIterator struct {
	Event *Betting022NewOraclizeQuery // Event containing the contract specifics and raw log

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
func (it *Betting022NewOraclizeQueryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting022NewOraclizeQuery)
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
		it.Event = new(Betting022NewOraclizeQuery)
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
func (it *Betting022NewOraclizeQueryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting022NewOraclizeQueryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting022NewOraclizeQuery represents a NewOraclizeQuery event raised by the Betting022 contract.
type Betting022NewOraclizeQuery struct {
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewOraclizeQuery is a free log retrieval operation binding the contract event 0x46cb989ef9cef13e930e3b7f286225a086e716a90d63e0b7da85d310a9db0c9a.
//
// Solidity: e newOraclizeQuery(description string)
func (_Betting022 *Betting022Filterer) FilterNewOraclizeQuery(opts *bind.FilterOpts) (*Betting022NewOraclizeQueryIterator, error) {

	logs, sub, err := _Betting022.contract.FilterLogs(opts, "newOraclizeQuery")
	if err != nil {
		return nil, err
	}
	return &Betting022NewOraclizeQueryIterator{contract: _Betting022.contract, event: "newOraclizeQuery", logs: logs, sub: sub}, nil
}

// WatchNewOraclizeQuery is a free log subscription operation binding the contract event 0x46cb989ef9cef13e930e3b7f286225a086e716a90d63e0b7da85d310a9db0c9a.
//
// Solidity: e newOraclizeQuery(description string)
func (_Betting022 *Betting022Filterer) WatchNewOraclizeQuery(opts *bind.WatchOpts, sink chan<- *Betting022NewOraclizeQuery) (event.Subscription, error) {

	logs, sub, err := _Betting022.contract.WatchLogs(opts, "newOraclizeQuery")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting022NewOraclizeQuery)
				if err := _Betting022.contract.UnpackLog(event, "newOraclizeQuery", log); err != nil {
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

// Betting022NewPriceTickerIterator is returned from FilterNewPriceTicker and is used to iterate over the raw logs and unpacked data for NewPriceTicker events raised by the Betting022 contract.
type Betting022NewPriceTickerIterator struct {
	Event *Betting022NewPriceTicker // Event containing the contract specifics and raw log

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
func (it *Betting022NewPriceTickerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting022NewPriceTicker)
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
		it.Event = new(Betting022NewPriceTicker)
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
func (it *Betting022NewPriceTickerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting022NewPriceTickerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting022NewPriceTicker represents a NewPriceTicker event raised by the Betting022 contract.
type Betting022NewPriceTicker struct {
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewPriceTicker is a free log retrieval operation binding the contract event 0xc4db019ddd899ab8c4169d2877afae028fd941bb9c752cf08f14d21924c06734.
//
// Solidity: e newPriceTicker(price uint256)
func (_Betting022 *Betting022Filterer) FilterNewPriceTicker(opts *bind.FilterOpts) (*Betting022NewPriceTickerIterator, error) {

	logs, sub, err := _Betting022.contract.FilterLogs(opts, "newPriceTicker")
	if err != nil {
		return nil, err
	}
	return &Betting022NewPriceTickerIterator{contract: _Betting022.contract, event: "newPriceTicker", logs: logs, sub: sub}, nil
}

// WatchNewPriceTicker is a free log subscription operation binding the contract event 0xc4db019ddd899ab8c4169d2877afae028fd941bb9c752cf08f14d21924c06734.
//
// Solidity: e newPriceTicker(price uint256)
func (_Betting022 *Betting022Filterer) WatchNewPriceTicker(opts *bind.WatchOpts, sink chan<- *Betting022NewPriceTicker) (event.Subscription, error) {

	logs, sub, err := _Betting022.contract.WatchLogs(opts, "newPriceTicker")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting022NewPriceTicker)
				if err := _Betting022.contract.UnpackLog(event, "newPriceTicker", log); err != nil {
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
