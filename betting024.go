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

// Betting024ABI is the input ABI used to generate the binding from.
const Betting024ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"horse\",\"type\":\"bytes32\"}],\"name\":\"placeBet\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claim_reward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"winner_horse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coin_pointer\",\"type\":\"bytes32\"},{\"name\":\"result\",\"type\":\"uint256\"},{\"name\":\"isPrePrice\",\"type\":\"bool\"}],\"name\":\"priceCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"winnerPoolTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"horses\",\"outputs\":[{\"name\":\"BTC_delta\",\"type\":\"int64\"},{\"name\":\"ETH_delta\",\"type\":\"int64\"},{\"name\":\"LTC_delta\",\"type\":\"int64\"},{\"name\":\"BTC\",\"type\":\"bytes32\"},{\"name\":\"ETH\",\"type\":\"bytes32\"},{\"name\":\"LTC\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"forceVoidExternal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getChronus\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"bytes32\"},{\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"getCoinIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"chronus\",\"outputs\":[{\"name\":\"betting_open\",\"type\":\"bool\"},{\"name\":\"race_start\",\"type\":\"bool\"},{\"name\":\"race_end\",\"type\":\"bool\"},{\"name\":\"voided_bet\",\"type\":\"bool\"},{\"name\":\"starting_time\",\"type\":\"uint32\"},{\"name\":\"betting_duration\",\"type\":\"uint32\"},{\"name\":\"race_duration\",\"type\":\"uint32\"},{\"name\":\"voided_timestamp\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bettingDuration\",\"type\":\"uint32\"},{\"name\":\"_raceDuration\",\"type\":\"uint32\"}],\"name\":\"setupRace\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reward_total\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkReward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"coinIndex\",\"outputs\":[{\"name\":\"pre\",\"type\":\"uint256\"},{\"name\":\"post\",\"type\":\"uint256\"},{\"name\":\"total\",\"type\":\"uint160\"},{\"name\":\"count\",\"type\":\"uint32\"},{\"name\":\"price_check\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"total_reward\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"recovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_horse\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_date\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"coin_pointer\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"result\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"isPrePrice\",\"type\":\"bool\"}],\"name\":\"PriceCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"RefundEnabled\",\"type\":\"event\"}]"

// Betting024 is an auto generated Go binding around an Ethereum contract.
type Betting024 struct {
	Betting024Caller     // Read-only binding to the contract
	Betting024Transactor // Write-only binding to the contract
	Betting024Filterer   // Log filterer for contract events
}

// Betting024Caller is an auto generated read-only Go binding around an Ethereum contract.
type Betting024Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Betting024Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Betting024Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Betting024Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Betting024Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Betting024Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Betting024Session struct {
	Contract     *Betting024       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Betting024CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Betting024CallerSession struct {
	Contract *Betting024Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Betting024TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Betting024TransactorSession struct {
	Contract     *Betting024Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Betting024Raw is an auto generated low-level Go binding around an Ethereum contract.
type Betting024Raw struct {
	Contract *Betting024 // Generic contract binding to access the raw methods on
}

// Betting024CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Betting024CallerRaw struct {
	Contract *Betting024Caller // Generic read-only contract binding to access the raw methods on
}

// Betting024TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Betting024TransactorRaw struct {
	Contract *Betting024Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBetting024 creates a new instance of Betting024, bound to a specific deployed contract.
func NewBetting024(address common.Address, backend bind.ContractBackend) (*Betting024, error) {
	contract, err := bindBetting024(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Betting024{Betting024Caller: Betting024Caller{contract: contract}, Betting024Transactor: Betting024Transactor{contract: contract}, Betting024Filterer: Betting024Filterer{contract: contract}}, nil
}

// NewBetting024Caller creates a new read-only instance of Betting024, bound to a specific deployed contract.
func NewBetting024Caller(address common.Address, caller bind.ContractCaller) (*Betting024Caller, error) {
	contract, err := bindBetting024(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Betting024Caller{contract: contract}, nil
}

// NewBetting024Transactor creates a new write-only instance of Betting024, bound to a specific deployed contract.
func NewBetting024Transactor(address common.Address, transactor bind.ContractTransactor) (*Betting024Transactor, error) {
	contract, err := bindBetting024(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Betting024Transactor{contract: contract}, nil
}

// NewBetting024Filterer creates a new log filterer instance of Betting024, bound to a specific deployed contract.
func NewBetting024Filterer(address common.Address, filterer bind.ContractFilterer) (*Betting024Filterer, error) {
	contract, err := bindBetting024(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Betting024Filterer{contract: contract}, nil
}

// bindBetting024 binds a generic wrapper to an already deployed contract.
func bindBetting024(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Betting024ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting024 *Betting024Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Betting024.Contract.Betting024Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting024 *Betting024Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting024.Contract.Betting024Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting024 *Betting024Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting024.Contract.Betting024Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Betting024 *Betting024CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Betting024.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Betting024 *Betting024TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting024.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Betting024 *Betting024TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Betting024.Contract.contract.Transact(opts, method, params...)
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting024 *Betting024Caller) CheckReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting024.contract.Call(opts, out, "checkReward")
	return *ret0, err
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting024 *Betting024Session) CheckReward() (*big.Int, error) {
	return _Betting024.Contract.CheckReward(&_Betting024.CallOpts)
}

// CheckReward is a free data retrieval call binding the contract method 0xc4b24a46.
//
// Solidity: function checkReward() constant returns(uint256)
func (_Betting024 *Betting024CallerSession) CheckReward() (*big.Int, error) {
	return _Betting024.Contract.CheckReward(&_Betting024.CallOpts)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting024 *Betting024Caller) Chronus(opts *bind.CallOpts) (struct {
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
	err := _Betting024.contract.Call(opts, out, "chronus")
	return *ret, err
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting024 *Betting024Session) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _Betting024.Contract.Chronus(&_Betting024.CallOpts)
}

// Chronus is a free data retrieval call binding the contract method 0x84304ee5.
//
// Solidity: function chronus() constant returns(betting_open bool, race_start bool, race_end bool, voided_bet bool, starting_time uint32, betting_duration uint32, race_duration uint32, voided_timestamp uint32)
func (_Betting024 *Betting024CallerSession) Chronus() (struct {
	BettingOpen     bool
	RaceStart       bool
	RaceEnd         bool
	VoidedBet       bool
	StartingTime    uint32
	BettingDuration uint32
	RaceDuration    uint32
	VoidedTimestamp uint32
}, error) {
	return _Betting024.Contract.Chronus(&_Betting024.CallOpts)
}

// CoinIndex is a free data retrieval call binding the contract method 0xd2aed6d7.
//
// Solidity: function coinIndex( bytes32) constant returns(pre uint256, post uint256, total uint160, count uint32, price_check bool)
func (_Betting024 *Betting024Caller) CoinIndex(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _Betting024.contract.Call(opts, out, "coinIndex", arg0)
	return *ret, err
}

// CoinIndex is a free data retrieval call binding the contract method 0xd2aed6d7.
//
// Solidity: function coinIndex( bytes32) constant returns(pre uint256, post uint256, total uint160, count uint32, price_check bool)
func (_Betting024 *Betting024Session) CoinIndex(arg0 [32]byte) (struct {
	Pre        *big.Int
	Post       *big.Int
	Total      *big.Int
	Count      uint32
	PriceCheck bool
}, error) {
	return _Betting024.Contract.CoinIndex(&_Betting024.CallOpts, arg0)
}

// CoinIndex is a free data retrieval call binding the contract method 0xd2aed6d7.
//
// Solidity: function coinIndex( bytes32) constant returns(pre uint256, post uint256, total uint160, count uint32, price_check bool)
func (_Betting024 *Betting024CallerSession) CoinIndex(arg0 [32]byte) (struct {
	Pre        *big.Int
	Post       *big.Int
	Total      *big.Int
	Count      uint32
	PriceCheck bool
}, error) {
	return _Betting024.Contract.CoinIndex(&_Betting024.CallOpts, arg0)
}

// GetChronus is a free data retrieval call binding the contract method 0x5ad6ba47.
//
// Solidity: function getChronus() constant returns(uint32[])
func (_Betting024 *Betting024Caller) GetChronus(opts *bind.CallOpts) ([]uint32, error) {
	var (
		ret0 = new([]uint32)
	)
	out := ret0
	err := _Betting024.contract.Call(opts, out, "getChronus")
	return *ret0, err
}

// GetChronus is a free data retrieval call binding the contract method 0x5ad6ba47.
//
// Solidity: function getChronus() constant returns(uint32[])
func (_Betting024 *Betting024Session) GetChronus() ([]uint32, error) {
	return _Betting024.Contract.GetChronus(&_Betting024.CallOpts)
}

// GetChronus is a free data retrieval call binding the contract method 0x5ad6ba47.
//
// Solidity: function getChronus() constant returns(uint32[])
func (_Betting024 *Betting024CallerSession) GetChronus() ([]uint32, error) {
	return _Betting024.Contract.GetChronus(&_Betting024.CallOpts)
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting024 *Betting024Caller) GetCoinIndex(opts *bind.CallOpts, index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
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
	err := _Betting024.contract.Call(opts, out, "getCoinIndex", index, candidate)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting024 *Betting024Session) GetCoinIndex(index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Betting024.Contract.GetCoinIndex(&_Betting024.CallOpts, index, candidate)
}

// GetCoinIndex is a free data retrieval call binding the contract method 0x7274f35b.
//
// Solidity: function getCoinIndex(index bytes32, candidate address) constant returns(uint256, uint256, uint256, bool, uint256)
func (_Betting024 *Betting024CallerSession) GetCoinIndex(index [32]byte, candidate common.Address) (*big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Betting024.Contract.GetCoinIndex(&_Betting024.CallOpts, index, candidate)
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int64, ETH_delta int64, LTC_delta int64, BTC bytes32, ETH bytes32, LTC bytes32)
func (_Betting024 *Betting024Caller) Horses(opts *bind.CallOpts) (struct {
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
	err := _Betting024.contract.Call(opts, out, "horses")
	return *ret, err
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int64, ETH_delta int64, LTC_delta int64, BTC bytes32, ETH bytes32, LTC bytes32)
func (_Betting024 *Betting024Session) Horses() (struct {
	BTCDelta int64
	ETHDelta int64
	LTCDelta int64
	BTC      [32]byte
	ETH      [32]byte
	LTC      [32]byte
}, error) {
	return _Betting024.Contract.Horses(&_Betting024.CallOpts)
}

// Horses is a free data retrieval call binding the contract method 0x43bddf40.
//
// Solidity: function horses() constant returns(BTC_delta int64, ETH_delta int64, LTC_delta int64, BTC bytes32, ETH bytes32, LTC bytes32)
func (_Betting024 *Betting024CallerSession) Horses() (struct {
	BTCDelta int64
	ETHDelta int64
	LTCDelta int64
	BTC      [32]byte
	ETH      [32]byte
	LTC      [32]byte
}, error) {
	return _Betting024.Contract.Horses(&_Betting024.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting024 *Betting024Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Betting024.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting024 *Betting024Session) Owner() (common.Address, error) {
	return _Betting024.Contract.Owner(&_Betting024.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Betting024 *Betting024CallerSession) Owner() (common.Address, error) {
	return _Betting024.Contract.Owner(&_Betting024.CallOpts)
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting024 *Betting024Caller) RewardTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting024.contract.Call(opts, out, "reward_total")
	return *ret0, err
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting024 *Betting024Session) RewardTotal() (*big.Int, error) {
	return _Betting024.Contract.RewardTotal(&_Betting024.CallOpts)
}

// RewardTotal is a free data retrieval call binding the contract method 0xaa93038b.
//
// Solidity: function reward_total() constant returns(uint256)
func (_Betting024 *Betting024CallerSession) RewardTotal() (*big.Int, error) {
	return _Betting024.Contract.RewardTotal(&_Betting024.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting024 *Betting024Caller) TotalReward(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting024.contract.Call(opts, out, "total_reward")
	return *ret0, err
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting024 *Betting024Session) TotalReward() (*big.Int, error) {
	return _Betting024.Contract.TotalReward(&_Betting024.CallOpts)
}

// TotalReward is a free data retrieval call binding the contract method 0xd3d2172e.
//
// Solidity: function total_reward() constant returns(uint256)
func (_Betting024 *Betting024CallerSession) TotalReward() (*big.Int, error) {
	return _Betting024.Contract.TotalReward(&_Betting024.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting024 *Betting024Caller) Version(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Betting024.contract.Call(opts, out, "version")
	return *ret0, err
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting024 *Betting024Session) Version() (string, error) {
	return _Betting024.Contract.Version(&_Betting024.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() constant returns(string)
func (_Betting024 *Betting024CallerSession) Version() (string, error) {
	return _Betting024.Contract.Version(&_Betting024.CallOpts)
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting024 *Betting024Caller) WinnerPoolTotal(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Betting024.contract.Call(opts, out, "winnerPoolTotal")
	return *ret0, err
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting024 *Betting024Session) WinnerPoolTotal() (*big.Int, error) {
	return _Betting024.Contract.WinnerPoolTotal(&_Betting024.CallOpts)
}

// WinnerPoolTotal is a free data retrieval call binding the contract method 0x29114d65.
//
// Solidity: function winnerPoolTotal() constant returns(uint256)
func (_Betting024 *Betting024CallerSession) WinnerPoolTotal() (*big.Int, error) {
	return _Betting024.Contract.WinnerPoolTotal(&_Betting024.CallOpts)
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting024 *Betting024Caller) WinnerHorse(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Betting024.contract.Call(opts, out, "winner_horse", arg0)
	return *ret0, err
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting024 *Betting024Session) WinnerHorse(arg0 [32]byte) (bool, error) {
	return _Betting024.Contract.WinnerHorse(&_Betting024.CallOpts, arg0)
}

// WinnerHorse is a free data retrieval call binding the contract method 0x0f769644.
//
// Solidity: function winner_horse( bytes32) constant returns(bool)
func (_Betting024 *Betting024CallerSession) WinnerHorse(arg0 [32]byte) (bool, error) {
	return _Betting024.Contract.WinnerHorse(&_Betting024.CallOpts, arg0)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting024 *Betting024Transactor) ChangeOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Betting024.contract.Transact(opts, "changeOwnership", _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting024 *Betting024Session) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Betting024.Contract.ChangeOwnership(&_Betting024.TransactOpts, _newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x2af4c31e.
//
// Solidity: function changeOwnership(_newOwner address) returns()
func (_Betting024 *Betting024TransactorSession) ChangeOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Betting024.Contract.ChangeOwnership(&_Betting024.TransactOpts, _newOwner)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting024 *Betting024Transactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting024.contract.Transact(opts, "claim_reward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting024 *Betting024Session) ClaimReward() (*types.Transaction, error) {
	return _Betting024.Contract.ClaimReward(&_Betting024.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0x055ee253.
//
// Solidity: function claim_reward() returns()
func (_Betting024 *Betting024TransactorSession) ClaimReward() (*types.Transaction, error) {
	return _Betting024.Contract.ClaimReward(&_Betting024.TransactOpts)
}

// ForceVoidExternal is a paid mutator transaction binding the contract method 0x4564ea36.
//
// Solidity: function forceVoidExternal() returns()
func (_Betting024 *Betting024Transactor) ForceVoidExternal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting024.contract.Transact(opts, "forceVoidExternal")
}

// ForceVoidExternal is a paid mutator transaction binding the contract method 0x4564ea36.
//
// Solidity: function forceVoidExternal() returns()
func (_Betting024 *Betting024Session) ForceVoidExternal() (*types.Transaction, error) {
	return _Betting024.Contract.ForceVoidExternal(&_Betting024.TransactOpts)
}

// ForceVoidExternal is a paid mutator transaction binding the contract method 0x4564ea36.
//
// Solidity: function forceVoidExternal() returns()
func (_Betting024 *Betting024TransactorSession) ForceVoidExternal() (*types.Transaction, error) {
	return _Betting024.Contract.ForceVoidExternal(&_Betting024.TransactOpts)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting024 *Betting024Transactor) PlaceBet(opts *bind.TransactOpts, horse [32]byte) (*types.Transaction, error) {
	return _Betting024.contract.Transact(opts, "placeBet", horse)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting024 *Betting024Session) PlaceBet(horse [32]byte) (*types.Transaction, error) {
	return _Betting024.Contract.PlaceBet(&_Betting024.TransactOpts, horse)
}

// PlaceBet is a paid mutator transaction binding the contract method 0x042b5fed.
//
// Solidity: function placeBet(horse bytes32) returns()
func (_Betting024 *Betting024TransactorSession) PlaceBet(horse [32]byte) (*types.Transaction, error) {
	return _Betting024.Contract.PlaceBet(&_Betting024.TransactOpts, horse)
}

// PriceCallback is a paid mutator transaction binding the contract method 0x11dcee2f.
//
// Solidity: function priceCallback(coin_pointer bytes32, result uint256, isPrePrice bool) returns()
func (_Betting024 *Betting024Transactor) PriceCallback(opts *bind.TransactOpts, coin_pointer [32]byte, result *big.Int, isPrePrice bool) (*types.Transaction, error) {
	return _Betting024.contract.Transact(opts, "priceCallback", coin_pointer, result, isPrePrice)
}

// PriceCallback is a paid mutator transaction binding the contract method 0x11dcee2f.
//
// Solidity: function priceCallback(coin_pointer bytes32, result uint256, isPrePrice bool) returns()
func (_Betting024 *Betting024Session) PriceCallback(coin_pointer [32]byte, result *big.Int, isPrePrice bool) (*types.Transaction, error) {
	return _Betting024.Contract.PriceCallback(&_Betting024.TransactOpts, coin_pointer, result, isPrePrice)
}

// PriceCallback is a paid mutator transaction binding the contract method 0x11dcee2f.
//
// Solidity: function priceCallback(coin_pointer bytes32, result uint256, isPrePrice bool) returns()
func (_Betting024 *Betting024TransactorSession) PriceCallback(coin_pointer [32]byte, result *big.Int, isPrePrice bool) (*types.Transaction, error) {
	return _Betting024.Contract.PriceCallback(&_Betting024.TransactOpts, coin_pointer, result, isPrePrice)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting024 *Betting024Transactor) Recovery(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting024.contract.Transact(opts, "recovery")
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting024 *Betting024Session) Recovery() (*types.Transaction, error) {
	return _Betting024.Contract.Recovery(&_Betting024.TransactOpts)
}

// Recovery is a paid mutator transaction binding the contract method 0xddceafa9.
//
// Solidity: function recovery() returns()
func (_Betting024 *Betting024TransactorSession) Recovery() (*types.Transaction, error) {
	return _Betting024.Contract.Recovery(&_Betting024.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting024 *Betting024Transactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Betting024.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting024 *Betting024Session) Refund() (*types.Transaction, error) {
	return _Betting024.Contract.Refund(&_Betting024.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_Betting024 *Betting024TransactorSession) Refund() (*types.Transaction, error) {
	return _Betting024.Contract.Refund(&_Betting024.TransactOpts)
}

// SetupRace is a paid mutator transaction binding the contract method 0x8b63c86f.
//
// Solidity: function setupRace(_bettingDuration uint32, _raceDuration uint32) returns()
func (_Betting024 *Betting024Transactor) SetupRace(opts *bind.TransactOpts, _bettingDuration uint32, _raceDuration uint32) (*types.Transaction, error) {
	return _Betting024.contract.Transact(opts, "setupRace", _bettingDuration, _raceDuration)
}

// SetupRace is a paid mutator transaction binding the contract method 0x8b63c86f.
//
// Solidity: function setupRace(_bettingDuration uint32, _raceDuration uint32) returns()
func (_Betting024 *Betting024Session) SetupRace(_bettingDuration uint32, _raceDuration uint32) (*types.Transaction, error) {
	return _Betting024.Contract.SetupRace(&_Betting024.TransactOpts, _bettingDuration, _raceDuration)
}

// SetupRace is a paid mutator transaction binding the contract method 0x8b63c86f.
//
// Solidity: function setupRace(_bettingDuration uint32, _raceDuration uint32) returns()
func (_Betting024 *Betting024TransactorSession) SetupRace(_bettingDuration uint32, _raceDuration uint32) (*types.Transaction, error) {
	return _Betting024.Contract.SetupRace(&_Betting024.TransactOpts, _bettingDuration, _raceDuration)
}

// Betting024DepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Betting024 contract.
type Betting024DepositIterator struct {
	Event *Betting024Deposit // Event containing the contract specifics and raw log

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
func (it *Betting024DepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting024Deposit)
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
		it.Event = new(Betting024Deposit)
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
func (it *Betting024DepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting024DepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting024Deposit represents a Deposit event raised by the Betting024 contract.
type Betting024Deposit struct {
	From  common.Address
	Value *big.Int
	Horse [32]byte
	Date  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x60452eb7177e8d41c9d9fbc4c6e9ccf55a4d44d412355fbf2f02668e0d1a0ce1.
//
// Solidity: e Deposit(_from address, _value uint256, _horse bytes32, _date uint256)
func (_Betting024 *Betting024Filterer) FilterDeposit(opts *bind.FilterOpts) (*Betting024DepositIterator, error) {

	logs, sub, err := _Betting024.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &Betting024DepositIterator{contract: _Betting024.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x60452eb7177e8d41c9d9fbc4c6e9ccf55a4d44d412355fbf2f02668e0d1a0ce1.
//
// Solidity: e Deposit(_from address, _value uint256, _horse bytes32, _date uint256)
func (_Betting024 *Betting024Filterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *Betting024Deposit) (event.Subscription, error) {

	logs, sub, err := _Betting024.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting024Deposit)
				if err := _Betting024.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// Betting024PriceCallbackIterator is returned from FilterPriceCallback and is used to iterate over the raw logs and unpacked data for PriceCallback events raised by the Betting024 contract.
type Betting024PriceCallbackIterator struct {
	Event *Betting024PriceCallback // Event containing the contract specifics and raw log

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
func (it *Betting024PriceCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting024PriceCallback)
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
		it.Event = new(Betting024PriceCallback)
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
func (it *Betting024PriceCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting024PriceCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting024PriceCallback represents a PriceCallback event raised by the Betting024 contract.
type Betting024PriceCallback struct {
	CoinPointer [32]byte
	Result      *big.Int
	IsPrePrice  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPriceCallback is a free log retrieval operation binding the contract event 0xde16ef9c49ad256644606beb97130511ba3d64bbd230380f8edd107527e5a9da.
//
// Solidity: e PriceCallback(coin_pointer bytes32, result uint256, isPrePrice bool)
func (_Betting024 *Betting024Filterer) FilterPriceCallback(opts *bind.FilterOpts) (*Betting024PriceCallbackIterator, error) {

	logs, sub, err := _Betting024.contract.FilterLogs(opts, "PriceCallback")
	if err != nil {
		return nil, err
	}
	return &Betting024PriceCallbackIterator{contract: _Betting024.contract, event: "PriceCallback", logs: logs, sub: sub}, nil
}

// WatchPriceCallback is a free log subscription operation binding the contract event 0xde16ef9c49ad256644606beb97130511ba3d64bbd230380f8edd107527e5a9da.
//
// Solidity: e PriceCallback(coin_pointer bytes32, result uint256, isPrePrice bool)
func (_Betting024 *Betting024Filterer) WatchPriceCallback(opts *bind.WatchOpts, sink chan<- *Betting024PriceCallback) (event.Subscription, error) {

	logs, sub, err := _Betting024.contract.WatchLogs(opts, "PriceCallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting024PriceCallback)
				if err := _Betting024.contract.UnpackLog(event, "PriceCallback", log); err != nil {
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

// Betting024RefundEnabledIterator is returned from FilterRefundEnabled and is used to iterate over the raw logs and unpacked data for RefundEnabled events raised by the Betting024 contract.
type Betting024RefundEnabledIterator struct {
	Event *Betting024RefundEnabled // Event containing the contract specifics and raw log

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
func (it *Betting024RefundEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting024RefundEnabled)
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
		it.Event = new(Betting024RefundEnabled)
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
func (it *Betting024RefundEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting024RefundEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting024RefundEnabled represents a RefundEnabled event raised by the Betting024 contract.
type Betting024RefundEnabled struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundEnabled is a free log retrieval operation binding the contract event 0x9267bd1e840f8c032ec399dab88550ddacce435477212b384a3d761f395efa7f.
//
// Solidity: e RefundEnabled(reason string)
func (_Betting024 *Betting024Filterer) FilterRefundEnabled(opts *bind.FilterOpts) (*Betting024RefundEnabledIterator, error) {

	logs, sub, err := _Betting024.contract.FilterLogs(opts, "RefundEnabled")
	if err != nil {
		return nil, err
	}
	return &Betting024RefundEnabledIterator{contract: _Betting024.contract, event: "RefundEnabled", logs: logs, sub: sub}, nil
}

// WatchRefundEnabled is a free log subscription operation binding the contract event 0x9267bd1e840f8c032ec399dab88550ddacce435477212b384a3d761f395efa7f.
//
// Solidity: e RefundEnabled(reason string)
func (_Betting024 *Betting024Filterer) WatchRefundEnabled(opts *bind.WatchOpts, sink chan<- *Betting024RefundEnabled) (event.Subscription, error) {

	logs, sub, err := _Betting024.contract.WatchLogs(opts, "RefundEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting024RefundEnabled)
				if err := _Betting024.contract.UnpackLog(event, "RefundEnabled", log); err != nil {
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

// Betting024WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Betting024 contract.
type Betting024WithdrawIterator struct {
	Event *Betting024Withdraw // Event containing the contract specifics and raw log

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
func (it *Betting024WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Betting024Withdraw)
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
		it.Event = new(Betting024Withdraw)
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
func (it *Betting024WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Betting024WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Betting024Withdraw represents a Withdraw event raised by the Betting024 contract.
type Betting024Withdraw struct {
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_to address, _value uint256)
func (_Betting024 *Betting024Filterer) FilterWithdraw(opts *bind.FilterOpts) (*Betting024WithdrawIterator, error) {

	logs, sub, err := _Betting024.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &Betting024WithdrawIterator{contract: _Betting024.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_to address, _value uint256)
func (_Betting024 *Betting024Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *Betting024Withdraw) (event.Subscription, error) {

	logs, sub, err := _Betting024.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Betting024Withdraw)
				if err := _Betting024.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
