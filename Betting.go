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

// BettingABI is the input ABI used to generate the binding from.
const BettingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"horse\",\"type\":\"bytes32\"}],\"name\":\"placeBet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim_reward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"winner_horse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coin_pointer\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"uint256\"},{\"name\":\"isPrePrice\",\"type\":\"bool\"}],\"name\":\"priceCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winnerPoolTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"horses\",\"outputs\":[{\"name\":\"BTC_delta\",\"type\":\"int64\"},{\"name\":\"ETH_delta\",\"type\":\"int64\"},{\"name\":\"LTC_delta\",\"type\":\"int64\"},{\"name\":\"BTC\",\"type\":\"bytes32\"},{\"name\":\"ETH\",\"type\":\"bytes32\"},{\"name\":\"LTC\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"bytes32\"},{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"getCoinIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chronus\",\"outputs\":[{\"name\":\"betting_open\",\"type\":\"bool\"},{\"name\":\"race_start\",\"type\":\"bool\"},{\"name\":\"race_end\",\"type\":\"bool\"},{\"name\":\"voided_bet\",\"type\":\"bool\"},{\"name\":\"starting_time\",\"type\":\"uint32\"},{\"name\":\"betting_duration\",\"type\":\"uint32\"},{\"name\":\"race_duration\",\"type\":\"uint32\"},{\"name\":\"voided_timestamp\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bettingDuration\",\"type\":\"uint32\"},{\"name\":\"_raceDuration\",\"type\":\"uint32\"}],\"name\":\"setupRace\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reward_total\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"coinIndex\",\"outputs\":[{\"name\":\"pre\",\"type\":\"uint256\"},{\"name\":\"post\",\"type\":\"uint256\"},{\"name\":\"total\",\"type\":\"uint160\"},{\"name\":\"count\",\"type\":\"uint32\"},{\"name\":\"price_check\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total_reward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"recovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_horse\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_date\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"coin_pointer\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isPrePrice\",\"type\":\"bool\"}],\"name\":\"PriceCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"RefundEnabled\",\"type\":\"event\"}]"

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

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting *BettingCaller) CheckReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting.contract.Call(opts, out, "checkReward")
	return *ret0, err
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting *BettingSession) CheckReward() (*big.Int, error) {
	return _Betting.Contract.CheckReward(&_Betting.CallOpts)
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting *BettingCallerSession) CheckReward() (*big.Int, error) {
	return _Betting.Contract.CheckReward(&_Betting.CallOpts)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting *BettingCaller) Chronus(opts *bind.CallOpts) (struct {
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
	err := _Betting.contract.Call(opts, out, "chronus")
	return *ret, err
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting *BettingSession) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _Betting.Contract.Chronus(&_Betting.CallOpts)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting *BettingCallerSession) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _Betting.Contract.Chronus(&_Betting.CallOpts)
}

// CoinIndex is a free data retrieval call binding the contract method 0xd2aed6d7.
//
// Solidity: function coinIndex( bytes32) constant returns(pre uint256, post uint256, total uint160, count uint32, price_check bool)
func (_Betting *BettingCaller) CoinIndex(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Pre        *big.Int
	Post       *big.Int
	Total      *big.Int
	Count      uint32
	PriceCheck bool
}, error) {
	ret := new(struct {
		Pre        *big.Int
		Post       *big.Int
		Total      *big.Int
		Count      uint32
		PriceCheck bool
	})
	out := ret
	err := _Betting.contract.Call(opts, out, "coinIndex", arg0)
	return *ret, err
}

// CoinIndex is a free data retrieval call binding the contract method 0xd2aed6d7.
//
// Solidity: function coinIndex( bytes32) constant returns(pre uint256, post uint256, total uint160, count uint32, price_check bool)
func (_Betting *BettingSession) CoinIndex(arg0 [32]byte) (struct {
	Pre        *big.Int
	Post       *big.Int
	Total      *big.Int
	Count      uint32
	PriceCheck bool
}, error) {
	return _Betting.Contract.CoinIndex(&_Betting.CallOpts, arg0)
}

// CoinIndex is a free data retrieval call binding the contract method 0xd2aed6d7.
//
// Solidity: function coinIndex( bytes32) constant returns(pre uint256, post uint256, total uint160, count uint32, price_check bool)
func (_Betting *BettingCallerSession) CoinIndex(arg0 [32]byte) (struct {
	Pre        *big.Int
	Post       *big.Int
	Total      *big.Int
	Count      uint32
	PriceCheck bool
}, error) {
	return _Betting.Contract.CoinIndex(&_Betting.CallOpts, arg0)
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting *BettingCaller) GetCoinIndex(opts *bind.CallOpts, index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
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
	err := _Betting.contract.Call(opts, out, "getCoinIndex", index, candidate)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting *BettingSession) GetCoinIndex(index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Betting.Contract.GetCoinIndex(&_Betting.CallOpts, index, candidate)
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting *BettingCallerSession) GetCoinIndex(index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Betting.Contract.GetCoinIndex(&_Betting.CallOpts, index, candidate)
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int64, ETH_delta int64, LTC_delta int64, BTC bytes32, ETH bytes32, LTC bytes32)
func (_Betting *BettingCaller) Horses(opts *bind.CallOpts) (struct {
	BTCDelta int64
	ETHDelta int64
	LTCDelta int64
	BTC      [32]byte
	ETH      [32]byte
	LTC      [32]byte
}, error) {
	ret := new(struct {
		BTCDelta int64
		ETHDelta int64
		LTCDelta int64
		BTC      [32]byte
		ETH      [32]byte
		LTC      [32]byte
	})
	out := ret
	err := _Betting.contract.Call(opts, out, "horses")
	return *ret, err
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int64, ETH_delta int64, LTC_delta int64, BTC bytes32, ETH bytes32, LTC bytes32)
func (_Betting *BettingSession) Horses() (struct {
	BTCDelta int64
	ETHDelta int64
	LTCDelta int64
	BTC      [32]byte
	ETH      [32]byte
	LTC      [32]byte
}, error) {
	return _Betting.Contract.Horses(&_Betting.CallOpts)
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int64, ETH_delta int64, LTC_delta int64, BTC bytes32, ETH bytes32, LTC bytes32)
func (_Betting *BettingCallerSession) Horses() (struct {
	BTCDelta int64
	ETHDelta int64
	LTCDelta int64
	BTC      [32]byte
	ETH      [32]byte
	LTC      [32]byte
}, error) {
	return _Betting.Contract.Horses(&_Betting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting *BettingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Betting.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting *BettingSession) Owner() (common.Address, error) {
	return _Betting.Contract.Owner(&_Betting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting *BettingCallerSession) Owner() (common.Address, error) {
	return _Betting.Contract.Owner(&_Betting.CallOpts)
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting *BettingCaller) RewardTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting.contract.Call(opts, out, "reward_total")
	return *ret0, err
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting *BettingSession) RewardTotal() (*big.Int, error) {
	return _Betting.Contract.RewardTotal(&_Betting.CallOpts)
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting *BettingCallerSession) RewardTotal() (*big.Int, error) {
	return _Betting.Contract.RewardTotal(&_Betting.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting *BettingCaller) TotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting.contract.Call(opts, out, "total_reward")
	return *ret0, err
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting *BettingSession) TotalReward() (*big.Int, error) {
	return _Betting.Contract.TotalReward(&_Betting.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting *BettingCallerSession) TotalReward() (*big.Int, error) {
	return _Betting.Contract.TotalReward(&_Betting.CallOpts)
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

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting *BettingCaller) WinnerPoolTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting.contract.Call(opts, out, "winnerPoolTotal")
	return *ret0, err
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting *BettingSession) WinnerPoolTotal() (*big.Int, error) {
	return _Betting.Contract.WinnerPoolTotal(&_Betting.CallOpts)
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting *BettingCallerSession) WinnerPoolTotal() (*big.Int, error) {
	return _Betting.Contract.WinnerPoolTotal(&_Betting.CallOpts)
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting *BettingCaller) WinnerHorse(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Betting.contract.Call(opts, out, "winner_horse", arg0)
	return *ret0, err
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting *BettingSession) WinnerHorse(arg0 [32]byte) (bool, error) {
	return _Betting.Contract.WinnerHorse(&_Betting.CallOpts, arg0)
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting *BettingCallerSession) WinnerHorse(arg0 [32]byte) (bool, error) {
	return _Betting.Contract.WinnerHorse(&_Betting.CallOpts, arg0)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting *BettingTransactor) ChangeOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "changeOwnership", _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting *BettingSession) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Betting.Contract.ChangeOwnership(&_Betting.TransactOpts, _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting *BettingTransactorSession) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Betting.Contract.ChangeOwnership(&_Betting.TransactOpts, _newOwner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting *BettingTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "claim_reward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting *BettingSession) ClaimReward() (*types.Transaction, error) {
	return _Betting.Contract.ClaimReward(&_Betting.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting *BettingTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Betting.Contract.ClaimReward(&_Betting.TransactOpts)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting *BettingTransactor) PlaceBet(opts *bind.TransactOpts, horse [32]byte) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "placeBet", horse)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting *BettingSession) PlaceBet(horse [32]byte) (*types.Transaction, error) {
	return _Betting.Contract.PlaceBet(&_Betting.TransactOpts, horse)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting *BettingTransactorSession) PlaceBet(horse [32]byte) (*types.Transaction, error) {
	return _Betting.Contract.PlaceBet(&_Betting.TransactOpts, horse)
}

// PriceCallback is a paid mutator transaction binding the contract method 0x11dcee2f.
//
// Solidity: function priceCallback(coin_pointer bytes32, result uint256, isPrePrice bool) returns()
func (_Betting *BettingTransactor) PriceCallback(opts *bind.TransactOpts, coin_pointer [32]byte, result *big.Int, isPrePrice bool) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "priceCallback", coin_pointer, result, isPrePrice)
}

// PriceCallback is a paid mutator transaction binding the contract method 0x11dcee2f.
//
// Solidity: function priceCallback(coin_pointer bytes32, result uint256, isPrePrice bool) returns()
func (_Betting *BettingSession) PriceCallback(coin_pointer [32]byte, result *big.Int, isPrePrice bool) (*types.Transaction, error) {
	return _Betting.Contract.PriceCallback(&_Betting.TransactOpts, coin_pointer, result, isPrePrice)
}

// PriceCallback is a paid mutator transaction binding the contract method 0x11dcee2f.
//
// Solidity: function priceCallback(coin_pointer bytes32, result uint256, isPrePrice bool) returns()
func (_Betting *BettingTransactorSession) PriceCallback(coin_pointer [32]byte, result *big.Int, isPrePrice bool) (*types.Transaction, error) {
	return _Betting.Contract.PriceCallback(&_Betting.TransactOpts, coin_pointer, result, isPrePrice)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting *BettingTransactor) Recovery(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "recovery")
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting *BettingSession) Recovery() (*types.Transaction, error) {
	return _Betting.Contract.Recovery(&_Betting.TransactOpts)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting *BettingTransactorSession) Recovery() (*types.Transaction, error) {
	return _Betting.Contract.Recovery(&_Betting.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting *BettingTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting *BettingSession) Refund() (*types.Transaction, error) {
	return _Betting.Contract.Refund(&_Betting.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting *BettingTransactorSession) Refund() (*types.Transaction, error) {
	return _Betting.Contract.Refund(&_Betting.TransactOpts)
}

// SetupRace is a paid mutator transaction binding the contract method 0x8b63c86f.
//
// Solidity: function setupRace(_bettingDuration uint32, _raceDuration uint32) returns()
func (_Betting *BettingTransactor) SetupRace(opts *bind.TransactOpts, _bettingDuration uint32, _raceDuration uint32) (*types.Transaction, error) {
	return _Betting.contract.Transact(opts, "setupRace", _bettingDuration, _raceDuration)
}

// SetupRace is a paid mutator transaction binding the contract method 0x8b63c86f.
//
// Solidity: function setupRace(_bettingDuration uint32, _raceDuration uint32) returns()
func (_Betting *BettingSession) SetupRace(_bettingDuration uint32, _raceDuration uint32) (*types.Transaction, error) {
	return _Betting.Contract.SetupRace(&_Betting.TransactOpts, _bettingDuration, _raceDuration)
}

// SetupRace is a paid mutator transaction binding the contract method 0x8b63c86f.
//
// Solidity: function setupRace(_bettingDuration uint32, _raceDuration uint32) returns()
func (_Betting *BettingTransactorSession) SetupRace(_bettingDuration uint32, _raceDuration uint32) (*types.Transaction, error) {
	return _Betting.Contract.SetupRace(&_Betting.TransactOpts, _bettingDuration, _raceDuration)
}

// BettingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Betting contract.
type BettingDepositIterator struct {
	Event *BettingDeposit // Event containing the contract specifics and raw log

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
func (it *BettingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingDeposit)
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
		it.Event = new(BettingDeposit)
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
func (it *BettingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingDeposit represents a Deposit event raised by the Betting contract.
type BettingDeposit struct {
	From  common.Address
	Value *big.Int
	Horse [32]byte
	Date  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x60452eb7177e8d41c9d9fbc4c6e9ccf55a4d44d412355fbf2f02668e0d1a0ce1.
//
// Solidity: e Deposit(_from address, _value uint256, _horse bytes32, _date uint256)
func (_Betting *BettingFilterer) FilterDeposit(opts *bind.FilterOpts) (*BettingDepositIterator, error) {

	logs, sub, err := _Betting.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BettingDepositIterator{contract: _Betting.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x60452eb7177e8d41c9d9fbc4c6e9ccf55a4d44d412355fbf2f02668e0d1a0ce1.
//
// Solidity: e Deposit(_from address, _value uint256, _horse bytes32, _date uint256)
func (_Betting *BettingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BettingDeposit) (event.Subscription, error) {

	logs, sub, err := _Betting.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingDeposit)
				if err := _Betting.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// BettingPriceCallbackIterator is returned from FilterPriceCallback and is used to iterate over the raw logs and unpacked data for PriceCallback events raised by the Betting contract.
type BettingPriceCallbackIterator struct {
	Event *BettingPriceCallback // Event containing the contract specifics and raw log

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
func (it *BettingPriceCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingPriceCallback)
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
		it.Event = new(BettingPriceCallback)
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
func (it *BettingPriceCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingPriceCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingPriceCallback represents a PriceCallback event raised by the Betting contract.
type BettingPriceCallback struct {
	CoinPointer [32]byte
	Result      *big.Int
	IsPrePrice  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPriceCallback is a free log retrieval operation binding the contract event 0xde16ef9c49ad256644606beb97130511ba3d64bbd230380f8edd107527e5a9da.
//
// Solidity: e PriceCallback(coin_pointer bytes32, result uint256, isPrePrice bool)
func (_Betting *BettingFilterer) FilterPriceCallback(opts *bind.FilterOpts) (*BettingPriceCallbackIterator, error) {

	logs, sub, err := _Betting.contract.FilterLogs(opts, "PriceCallback")
	if err != nil {
		return nil, err
	}
	return &BettingPriceCallbackIterator{contract: _Betting.contract, event: "PriceCallback", logs: logs, sub: sub}, nil
}

// WatchPriceCallback is a free log subscription operation binding the contract event 0xde16ef9c49ad256644606beb97130511ba3d64bbd230380f8edd107527e5a9da.
//
// Solidity: e PriceCallback(coin_pointer bytes32, result uint256, isPrePrice bool)
func (_Betting *BettingFilterer) WatchPriceCallback(opts *bind.WatchOpts, sink chan<- *BettingPriceCallback) (event.Subscription, error) {

	logs, sub, err := _Betting.contract.WatchLogs(opts, "PriceCallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingPriceCallback)
				if err := _Betting.contract.UnpackLog(event, "PriceCallback", log); err != nil {
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

// BettingRefundEnabledIterator is returned from FilterRefundEnabled and is used to iterate over the raw logs and unpacked data for RefundEnabled events raised by the Betting contract.
type BettingRefundEnabledIterator struct {
	Event *BettingRefundEnabled // Event containing the contract specifics and raw log

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
func (it *BettingRefundEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingRefundEnabled)
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
		it.Event = new(BettingRefundEnabled)
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
func (it *BettingRefundEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingRefundEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingRefundEnabled represents a RefundEnabled event raised by the Betting contract.
type BettingRefundEnabled struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundEnabled is a free log retrieval operation binding the contract event 0x9267bd1e840f8c032ec399dab88550ddacce435477212b384a3d761f395efa7f.
//
// Solidity: e RefundEnabled(reason string)
func (_Betting *BettingFilterer) FilterRefundEnabled(opts *bind.FilterOpts) (*BettingRefundEnabledIterator, error) {

	logs, sub, err := _Betting.contract.FilterLogs(opts, "RefundEnabled")
	if err != nil {
		return nil, err
	}
	return &BettingRefundEnabledIterator{contract: _Betting.contract, event: "RefundEnabled", logs: logs, sub: sub}, nil
}

// WatchRefundEnabled is a free log subscription operation binding the contract event 0x9267bd1e840f8c032ec399dab88550ddacce435477212b384a3d761f395efa7f.
//
// Solidity: e RefundEnabled(reason string)
func (_Betting *BettingFilterer) WatchRefundEnabled(opts *bind.WatchOpts, sink chan<- *BettingRefundEnabled) (event.Subscription, error) {

	logs, sub, err := _Betting.contract.WatchLogs(opts, "RefundEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingRefundEnabled)
				if err := _Betting.contract.UnpackLog(event, "RefundEnabled", log); err != nil {
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

// BettingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Betting contract.
type BettingWithdrawIterator struct {
	Event *BettingWithdraw // Event containing the contract specifics and raw log

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
func (it *BettingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BettingWithdraw)
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
		it.Event = new(BettingWithdraw)
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
func (it *BettingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BettingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BettingWithdraw represents a Withdraw event raised by the Betting contract.
type BettingWithdraw struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_to address, _value uint256)
func (_Betting *BettingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*BettingWithdrawIterator, error) {

	logs, sub, err := _Betting.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &BettingWithdrawIterator{contract: _Betting.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_to address, _value uint256)
func (_Betting *BettingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BettingWithdraw) (event.Subscription, error) {

	logs, sub, err := _Betting.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BettingWithdraw)
				if err := _Betting.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
