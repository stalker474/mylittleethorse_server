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

// Betting023ABI is the input ABI used to generate the binding from.
const Betting023ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"horse\",\"type\":\"bytes32\"}],\"name\":\"placeBet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim_reward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"winner_horse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coin_pointer\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"uint256\"},{\"name\":\"isPrePrice\",\"type\":\"bool\"}],\"name\":\"priceCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winnerPoolTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"horses\",\"outputs\":[{\"name\":\"BTC_delta\",\"type\":\"int64\"},{\"name\":\"ETH_delta\",\"type\":\"int64\"},{\"name\":\"LTC_delta\",\"type\":\"int64\"},{\"name\":\"BTC\",\"type\":\"bytes32\"},{\"name\":\"ETH\",\"type\":\"bytes32\"},{\"name\":\"LTC\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"bytes32\"},{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"getCoinIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chronus\",\"outputs\":[{\"name\":\"betting_open\",\"type\":\"bool\"},{\"name\":\"race_start\",\"type\":\"bool\"},{\"name\":\"race_end\",\"type\":\"bool\"},{\"name\":\"voided_bet\",\"type\":\"bool\"},{\"name\":\"starting_time\",\"type\":\"uint32\"},{\"name\":\"betting_duration\",\"type\":\"uint32\"},{\"name\":\"race_duration\",\"type\":\"uint32\"},{\"name\":\"voided_timestamp\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bettingDuration\",\"type\":\"uint32\"},{\"name\":\"_raceDuration\",\"type\":\"uint32\"}],\"name\":\"setupRace\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reward_total\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"coinIndex\",\"outputs\":[{\"name\":\"pre\",\"type\":\"uint256\"},{\"name\":\"post\",\"type\":\"uint256\"},{\"name\":\"total\",\"type\":\"uint160\"},{\"name\":\"count\",\"type\":\"uint32\"},{\"name\":\"price_check\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total_reward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"recovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_horse\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_date\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"coin_pointer\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isPrePrice\",\"type\":\"bool\"}],\"name\":\"PriceCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"RefundEnabled\",\"type\":\"event\"}]"

// Betting023 is an auto generated Go binding around an Ethereum contract.
type Betting023 struct {
	Betting023Caller     // Read-only binding to the contract
	Betting023Transactor // Write-only binding to the contract
	Betting023Filterer   // Log filterer for contract events
}

// Betting023Caller is an auto generated read-only Go binding around an Ethereum contract.
type Betting023Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Betting023Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Betting023Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Betting023Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Betting023Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Betting023Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Betting023Session struct {
	Contract     *Betting023       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Betting023CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Betting023CallerSession struct {
	Contract *Betting023Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Betting023TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Betting023TransactorSession struct {
	Contract     *Betting023Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Betting023Raw is an auto generated low-level Go binding around an Ethereum contract.
type Betting023Raw struct {
	Contract *Betting023 // Generic contract binding to access the raw methods on
}

// Betting023CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Betting023CallerRaw struct {
	Contract *Betting023Caller // Generic read-only contract binding to access the raw methods on
}

// Betting023TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Betting023TransactorRaw struct {
	Contract *Betting023Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBetting023 creates a new instance of Betting023, bound to a specific deployed contract.
func NewBetting023(address common.Address, backend bind.ContractBackend) (*Betting023, error) {
	contract, err := bindBetting023(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Betting023{Betting023Caller: Betting023Caller{contract: contract}, Betting023Transactor: Betting023Transactor{contract: contract}, Betting023Filterer: Betting023Filterer{contract: contract}}, nil
}

// NewBetting023Caller creates a new read-only instance of Betting023, bound to a specific deployed contract.
func NewBetting023Caller(address common.Address, caller bind.ContractCaller) (*Betting023Caller, error) {
	contract, err := bindBetting023(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Betting023Caller{contract: contract}, nil
}

// NewBetting023Transactor creates a new write-only instance of Betting023, bound to a specific deployed contract.
func NewBetting023Transactor(address common.Address, transactor bind.ContractTransactor) (*Betting023Transactor, error) {
	contract, err := bindBetting023(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Betting023Transactor{contract: contract}, nil
}

// NewBetting023Filterer creates a new log filterer instance of Betting023, bound to a specific deployed contract.
func NewBetting023Filterer(address common.Address, filterer bind.ContractFilterer) (*Betting023Filterer, error) {
	contract, err := bindBetting023(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Betting023Filterer{contract: contract}, nil
}

// bindBetting023 binds a generic wrapper to an already deployed contract.
func bindBetting023(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Betting023ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting023 *Betting023Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Betting023.Contract.Betting023Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting023 *Betting023Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting023.Contract.Betting023Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting023 *Betting023Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting023.Contract.Betting023Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting023 *Betting023CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Betting023.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting023 *Betting023TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting023.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting023 *Betting023TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting023.Contract.contract.Transact(opts, method, params...)
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting023 *Betting023Caller) CheckReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting023.contract.Call(opts, out, "checkReward")
	return *ret0, err
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting023 *Betting023Session) CheckReward() (*big.Int, error) {
	return _Betting023.Contract.CheckReward(&_Betting023.CallOpts)
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting023 *Betting023CallerSession) CheckReward() (*big.Int, error) {
	return _Betting023.Contract.CheckReward(&_Betting023.CallOpts)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting023 *Betting023Caller) Chronus(opts *bind.CallOpts) (struct {
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
	err := _Betting023.contract.Call(opts, out, "chronus")
	return *ret, err
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting023 *Betting023Session) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _Betting023.Contract.Chronus(&_Betting023.CallOpts)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting023 *Betting023CallerSession) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _Betting023.Contract.Chronus(&_Betting023.CallOpts)
}

// CoinIndex is a free data retrieval call binding the contract method 0xd2aed6d7.
//
// Solidity: function coinIndex( bytes32) constant returns(pre uint256, post uint256, total uint160, count uint32, price_check bool)
func (_Betting023 *Betting023Caller) CoinIndex(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _Betting023.contract.Call(opts, out, "coinIndex", arg0)
	return *ret, err
}

// CoinIndex is a free data retrieval call binding the contract method 0xd2aed6d7.
//
// Solidity: function coinIndex( bytes32) constant returns(pre uint256, post uint256, total uint160, count uint32, price_check bool)
func (_Betting023 *Betting023Session) CoinIndex(arg0 [32]byte) (struct {
	Pre        *big.Int
	Post       *big.Int
	Total      *big.Int
	Count      uint32
	PriceCheck bool
}, error) {
	return _Betting023.Contract.CoinIndex(&_Betting023.CallOpts, arg0)
}

// CoinIndex is a free data retrieval call binding the contract method 0xd2aed6d7.
//
// Solidity: function coinIndex( bytes32) constant returns(pre uint256, post uint256, total uint160, count uint32, price_check bool)
func (_Betting023 *Betting023CallerSession) CoinIndex(arg0 [32]byte) (struct {
	Pre        *big.Int
	Post       *big.Int
	Total      *big.Int
	Count      uint32
	PriceCheck bool
}, error) {
	return _Betting023.Contract.CoinIndex(&_Betting023.CallOpts, arg0)
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting023 *Betting023Caller) GetCoinIndex(opts *bind.CallOpts, index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
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
	err := _Betting023.contract.Call(opts, out, "getCoinIndex", index, candidate)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting023 *Betting023Session) GetCoinIndex(index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Betting023.Contract.GetCoinIndex(&_Betting023.CallOpts, index, candidate)
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting023 *Betting023CallerSession) GetCoinIndex(index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Betting023.Contract.GetCoinIndex(&_Betting023.CallOpts, index, candidate)
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int64, ETH_delta int64, LTC_delta int64, BTC bytes32, ETH bytes32, LTC bytes32)
func (_Betting023 *Betting023Caller) Horses(opts *bind.CallOpts) (struct {
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
	err := _Betting023.contract.Call(opts, out, "horses")
	return *ret, err
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int64, ETH_delta int64, LTC_delta int64, BTC bytes32, ETH bytes32, LTC bytes32)
func (_Betting023 *Betting023Session) Horses() (struct {
	BTCDelta int64
	ETHDelta int64
	LTCDelta int64
	BTC      [32]byte
	ETH      [32]byte
	LTC      [32]byte
}, error) {
	return _Betting023.Contract.Horses(&_Betting023.CallOpts)
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int64, ETH_delta int64, LTC_delta int64, BTC bytes32, ETH bytes32, LTC bytes32)
func (_Betting023 *Betting023CallerSession) Horses() (struct {
	BTCDelta int64
	ETHDelta int64
	LTCDelta int64
	BTC      [32]byte
	ETH      [32]byte
	LTC      [32]byte
}, error) {
	return _Betting023.Contract.Horses(&_Betting023.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting023 *Betting023Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Betting023.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting023 *Betting023Session) Owner() (common.Address, error) {
	return _Betting023.Contract.Owner(&_Betting023.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting023 *Betting023CallerSession) Owner() (common.Address, error) {
	return _Betting023.Contract.Owner(&_Betting023.CallOpts)
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting023 *Betting023Caller) RewardTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting023.contract.Call(opts, out, "reward_total")
	return *ret0, err
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting023 *Betting023Session) RewardTotal() (*big.Int, error) {
	return _Betting023.Contract.RewardTotal(&_Betting023.CallOpts)
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting023 *Betting023CallerSession) RewardTotal() (*big.Int, error) {
	return _Betting023.Contract.RewardTotal(&_Betting023.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting023 *Betting023Caller) TotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting023.contract.Call(opts, out, "total_reward")
	return *ret0, err
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting023 *Betting023Session) TotalReward() (*big.Int, error) {
	return _Betting023.Contract.TotalReward(&_Betting023.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting023 *Betting023CallerSession) TotalReward() (*big.Int, error) {
	return _Betting023.Contract.TotalReward(&_Betting023.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting023 *Betting023Caller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Betting023.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting023 *Betting023Session) Version() (string, error) {
	return _Betting023.Contract.Version(&_Betting023.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting023 *Betting023CallerSession) Version() (string, error) {
	return _Betting023.Contract.Version(&_Betting023.CallOpts)
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting023 *Betting023Caller) WinnerPoolTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting023.contract.Call(opts, out, "winnerPoolTotal")
	return *ret0, err
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting023 *Betting023Session) WinnerPoolTotal() (*big.Int, error) {
	return _Betting023.Contract.WinnerPoolTotal(&_Betting023.CallOpts)
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting023 *Betting023CallerSession) WinnerPoolTotal() (*big.Int, error) {
	return _Betting023.Contract.WinnerPoolTotal(&_Betting023.CallOpts)
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting023 *Betting023Caller) WinnerHorse(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Betting023.contract.Call(opts, out, "winner_horse", arg0)
	return *ret0, err
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting023 *Betting023Session) WinnerHorse(arg0 [32]byte) (bool, error) {
	return _Betting023.Contract.WinnerHorse(&_Betting023.CallOpts, arg0)
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting023 *Betting023CallerSession) WinnerHorse(arg0 [32]byte) (bool, error) {
	return _Betting023.Contract.WinnerHorse(&_Betting023.CallOpts, arg0)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting023 *Betting023Transactor) ChangeOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Betting023.contract.Transact(opts, "changeOwnership", _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting023 *Betting023Session) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Betting023.Contract.ChangeOwnership(&_Betting023.TransactOpts, _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting023 *Betting023TransactorSession) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Betting023.Contract.ChangeOwnership(&_Betting023.TransactOpts, _newOwner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting023 *Betting023Transactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting023.contract.Transact(opts, "claim_reward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting023 *Betting023Session) ClaimReward() (*types.Transaction, error) {
	return _Betting023.Contract.ClaimReward(&_Betting023.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting023 *Betting023TransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Betting023.Contract.ClaimReward(&_Betting023.TransactOpts)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting023 *Betting023Transactor) PlaceBet(opts *bind.TransactOpts, horse [32]byte) (*types.Transaction, error) {
	return _Betting023.contract.Transact(opts, "placeBet", horse)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting023 *Betting023Session) PlaceBet(horse [32]byte) (*types.Transaction, error) {
	return _Betting023.Contract.PlaceBet(&_Betting023.TransactOpts, horse)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting023 *Betting023TransactorSession) PlaceBet(horse [32]byte) (*types.Transaction, error) {
	return _Betting023.Contract.PlaceBet(&_Betting023.TransactOpts, horse)
}

// PriceCallback is a paid mutator transaction binding the contract method 0x11dcee2f.
//
// Solidity: function priceCallback(coin_pointer bytes32, result uint256, isPrePrice bool) returns()
func (_Betting023 *Betting023Transactor) PriceCallback(opts *bind.TransactOpts, coin_pointer [32]byte, result *big.Int, isPrePrice bool) (*types.Transaction, error) {
	return _Betting023.contract.Transact(opts, "priceCallback", coin_pointer, result, isPrePrice)
}

// PriceCallback is a paid mutator transaction binding the contract method 0x11dcee2f.
//
// Solidity: function priceCallback(coin_pointer bytes32, result uint256, isPrePrice bool) returns()
func (_Betting023 *Betting023Session) PriceCallback(coin_pointer [32]byte, result *big.Int, isPrePrice bool) (*types.Transaction, error) {
	return _Betting023.Contract.PriceCallback(&_Betting023.TransactOpts, coin_pointer, result, isPrePrice)
}

// PriceCallback is a paid mutator transaction binding the contract method 0x11dcee2f.
//
// Solidity: function priceCallback(coin_pointer bytes32, result uint256, isPrePrice bool) returns()
func (_Betting023 *Betting023TransactorSession) PriceCallback(coin_pointer [32]byte, result *big.Int, isPrePrice bool) (*types.Transaction, error) {
	return _Betting023.Contract.PriceCallback(&_Betting023.TransactOpts, coin_pointer, result, isPrePrice)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting023 *Betting023Transactor) Recovery(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting023.contract.Transact(opts, "recovery")
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting023 *Betting023Session) Recovery() (*types.Transaction, error) {
	return _Betting023.Contract.Recovery(&_Betting023.TransactOpts)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting023 *Betting023TransactorSession) Recovery() (*types.Transaction, error) {
	return _Betting023.Contract.Recovery(&_Betting023.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting023 *Betting023Transactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting023.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting023 *Betting023Session) Refund() (*types.Transaction, error) {
	return _Betting023.Contract.Refund(&_Betting023.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting023 *Betting023TransactorSession) Refund() (*types.Transaction, error) {
	return _Betting023.Contract.Refund(&_Betting023.TransactOpts)
}

// SetupRace is a paid mutator transaction binding the contract method 0x8b63c86f.
//
// Solidity: function setupRace(_bettingDuration uint32, _raceDuration uint32) returns()
func (_Betting023 *Betting023Transactor) SetupRace(opts *bind.TransactOpts, _bettingDuration uint32, _raceDuration uint32) (*types.Transaction, error) {
	return _Betting023.contract.Transact(opts, "setupRace", _bettingDuration, _raceDuration)
}

// SetupRace is a paid mutator transaction binding the contract method 0x8b63c86f.
//
// Solidity: function setupRace(_bettingDuration uint32, _raceDuration uint32) returns()
func (_Betting023 *Betting023Session) SetupRace(_bettingDuration uint32, _raceDuration uint32) (*types.Transaction, error) {
	return _Betting023.Contract.SetupRace(&_Betting023.TransactOpts, _bettingDuration, _raceDuration)
}

// SetupRace is a paid mutator transaction binding the contract method 0x8b63c86f.
//
// Solidity: function setupRace(_bettingDuration uint32, _raceDuration uint32) returns()
func (_Betting023 *Betting023TransactorSession) SetupRace(_bettingDuration uint32, _raceDuration uint32) (*types.Transaction, error) {
	return _Betting023.Contract.SetupRace(&_Betting023.TransactOpts, _bettingDuration, _raceDuration)
}

// Betting023DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Betting023 contract.
type Betting023DepositIterator struct {
	Event *Betting023Deposit // Event containing the contract specifics and raw log

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
func (it *Betting023DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting023Deposit)
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
		it.Event = new(Betting023Deposit)
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
func (it *Betting023DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting023DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting023Deposit represents a Deposit event raised by the Betting023 contract.
type Betting023Deposit struct {
	From  common.Address
	Value *big.Int
	Horse [32]byte
	Date  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x60452eb7177e8d41c9d9fbc4c6e9ccf55a4d44d412355fbf2f02668e0d1a0ce1.
//
// Solidity: e Deposit(_from address, _value uint256, _horse bytes32, _date uint256)
func (_Betting023 *Betting023Filterer) FilterDeposit(opts *bind.FilterOpts) (*Betting023DepositIterator, error) {

	logs, sub, err := _Betting023.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &Betting023DepositIterator{contract: _Betting023.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x60452eb7177e8d41c9d9fbc4c6e9ccf55a4d44d412355fbf2f02668e0d1a0ce1.
//
// Solidity: e Deposit(_from address, _value uint256, _horse bytes32, _date uint256)
func (_Betting023 *Betting023Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Betting023Deposit) (event.Subscription, error) {

	logs, sub, err := _Betting023.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting023Deposit)
				if err := _Betting023.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// Betting023PriceCallbackIterator is returned from FilterPriceCallback and is used to iterate over the raw logs and unpacked data for PriceCallback events raised by the Betting023 contract.
type Betting023PriceCallbackIterator struct {
	Event *Betting023PriceCallback // Event containing the contract specifics and raw log

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
func (it *Betting023PriceCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting023PriceCallback)
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
		it.Event = new(Betting023PriceCallback)
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
func (it *Betting023PriceCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting023PriceCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting023PriceCallback represents a PriceCallback event raised by the Betting023 contract.
type Betting023PriceCallback struct {
	CoinPointer [32]byte
	Result      *big.Int
	IsPrePrice  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPriceCallback is a free log retrieval operation binding the contract event 0xde16ef9c49ad256644606beb97130511ba3d64bbd230380f8edd107527e5a9da.
//
// Solidity: e PriceCallback(coin_pointer bytes32, result uint256, isPrePrice bool)
func (_Betting023 *Betting023Filterer) FilterPriceCallback(opts *bind.FilterOpts) (*Betting023PriceCallbackIterator, error) {

	logs, sub, err := _Betting023.contract.FilterLogs(opts, "PriceCallback")
	if err != nil {
		return nil, err
	}
	return &Betting023PriceCallbackIterator{contract: _Betting023.contract, event: "PriceCallback", logs: logs, sub: sub}, nil
}

// WatchPriceCallback is a free log subscription operation binding the contract event 0xde16ef9c49ad256644606beb97130511ba3d64bbd230380f8edd107527e5a9da.
//
// Solidity: e PriceCallback(coin_pointer bytes32, result uint256, isPrePrice bool)
func (_Betting023 *Betting023Filterer) WatchPriceCallback(opts *bind.WatchOpts, sink chan<- *Betting023PriceCallback) (event.Subscription, error) {

	logs, sub, err := _Betting023.contract.WatchLogs(opts, "PriceCallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting023PriceCallback)
				if err := _Betting023.contract.UnpackLog(event, "PriceCallback", log); err != nil {
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

// Betting023RefundEnabledIterator is returned from FilterRefundEnabled and is used to iterate over the raw logs and unpacked data for RefundEnabled events raised by the Betting023 contract.
type Betting023RefundEnabledIterator struct {
	Event *Betting023RefundEnabled // Event containing the contract specifics and raw log

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
func (it *Betting023RefundEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting023RefundEnabled)
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
		it.Event = new(Betting023RefundEnabled)
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
func (it *Betting023RefundEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting023RefundEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting023RefundEnabled represents a RefundEnabled event raised by the Betting023 contract.
type Betting023RefundEnabled struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundEnabled is a free log retrieval operation binding the contract event 0x9267bd1e840f8c032ec399dab88550ddacce435477212b384a3d761f395efa7f.
//
// Solidity: e RefundEnabled(reason string)
func (_Betting023 *Betting023Filterer) FilterRefundEnabled(opts *bind.FilterOpts) (*Betting023RefundEnabledIterator, error) {

	logs, sub, err := _Betting023.contract.FilterLogs(opts, "RefundEnabled")
	if err != nil {
		return nil, err
	}
	return &Betting023RefundEnabledIterator{contract: _Betting023.contract, event: "RefundEnabled", logs: logs, sub: sub}, nil
}

// WatchRefundEnabled is a free log subscription operation binding the contract event 0x9267bd1e840f8c032ec399dab88550ddacce435477212b384a3d761f395efa7f.
//
// Solidity: e RefundEnabled(reason string)
func (_Betting023 *Betting023Filterer) WatchRefundEnabled(opts *bind.WatchOpts, sink chan<- *Betting023RefundEnabled) (event.Subscription, error) {

	logs, sub, err := _Betting023.contract.WatchLogs(opts, "RefundEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting023RefundEnabled)
				if err := _Betting023.contract.UnpackLog(event, "RefundEnabled", log); err != nil {
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

// Betting023WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Betting023 contract.
type Betting023WithdrawIterator struct {
	Event *Betting023Withdraw // Event containing the contract specifics and raw log

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
func (it *Betting023WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting023Withdraw)
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
		it.Event = new(Betting023Withdraw)
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
func (it *Betting023WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting023WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting023Withdraw represents a Withdraw event raised by the Betting023 contract.
type Betting023Withdraw struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_to address, _value uint256)
func (_Betting023 *Betting023Filterer) FilterWithdraw(opts *bind.FilterOpts) (*Betting023WithdrawIterator, error) {

	logs, sub, err := _Betting023.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &Betting023WithdrawIterator{contract: _Betting023.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_to address, _value uint256)
func (_Betting023 *Betting023Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Betting023Withdraw) (event.Subscription, error) {

	logs, sub, err := _Betting023.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting023Withdraw)
				if err := _Betting023.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
