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

// StrategyMetaData contains all meta data concerning the Strategy contract.
var StrategyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIStrategyManager\",\"name\":\"_strategyManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MaxPerDepositUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MaxTotalDepositsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"pauserRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"PauserRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newShares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"explanation\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTVLLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPerDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalDeposits\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"_underlyingToken\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_underlyingToken\",\"type\":\"address\"},{\"internalType\":\"contractIPauserRegistry\",\"name\":\"_pauserRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPerDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauserRegistry\",\"outputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPauserRegistry\",\"name\":\"newPauserRegistry\",\"type\":\"address\"}],\"name\":\"setPauserRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxPerDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newMaxTotalDeposits\",\"type\":\"uint256\"}],\"name\":\"setTVLLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"shares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountShares\",\"type\":\"uint256\"}],\"name\":\"sharesToUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountShares\",\"type\":\"uint256\"}],\"name\":\"sharesToUnderlyingView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strategyManager\",\"outputs\":[{\"internalType\":\"contractIStrategyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"}],\"name\":\"underlyingToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"}],\"name\":\"underlyingToSharesView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"underlyingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPausedStatus\",\"type\":\"uint256\"}],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userUnderlyingView\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountShares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use StrategyMetaData.ABI instead.
var StrategyABI = StrategyMetaData.ABI

// Strategy is an auto generated Go binding around an Ethereum contract.
type Strategy struct {
	StrategyCaller     // Read-only binding to the contract
	StrategyTransactor // Write-only binding to the contract
	StrategyFilterer   // Log filterer for contract events
}

// StrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrategySession struct {
	Contract     *Strategy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrategyCallerSession struct {
	Contract *StrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrategyTransactorSession struct {
	Contract     *StrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrategyRaw struct {
	Contract *Strategy // Generic contract binding to access the raw methods on
}

// StrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrategyCallerRaw struct {
	Contract *StrategyCaller // Generic read-only contract binding to access the raw methods on
}

// StrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrategyTransactorRaw struct {
	Contract *StrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrategy creates a new instance of Strategy, bound to a specific deployed contract.
func NewStrategy(address common.Address, backend bind.ContractBackend) (*Strategy, error) {
	contract, err := bindStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strategy{StrategyCaller: StrategyCaller{contract: contract}, StrategyTransactor: StrategyTransactor{contract: contract}, StrategyFilterer: StrategyFilterer{contract: contract}}, nil
}

// NewStrategyCaller creates a new read-only instance of Strategy, bound to a specific deployed contract.
func NewStrategyCaller(address common.Address, caller bind.ContractCaller) (*StrategyCaller, error) {
	contract, err := bindStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyCaller{contract: contract}, nil
}

// NewStrategyTransactor creates a new write-only instance of Strategy, bound to a specific deployed contract.
func NewStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*StrategyTransactor, error) {
	contract, err := bindStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrategyTransactor{contract: contract}, nil
}

// NewStrategyFilterer creates a new log filterer instance of Strategy, bound to a specific deployed contract.
func NewStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*StrategyFilterer, error) {
	contract, err := bindStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrategyFilterer{contract: contract}, nil
}

// bindStrategy binds a generic wrapper to an already deployed contract.
func bindStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StrategyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strategy *StrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strategy.Contract.StrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strategy *StrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.Contract.StrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strategy *StrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strategy.Contract.StrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strategy *StrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strategy *StrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strategy *StrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strategy.Contract.contract.Transact(opts, method, params...)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_Strategy *StrategyCaller) Explanation(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "explanation")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_Strategy *StrategySession) Explanation() (string, error) {
	return _Strategy.Contract.Explanation(&_Strategy.CallOpts)
}

// Explanation is a free data retrieval call binding the contract method 0xab5921e1.
//
// Solidity: function explanation() pure returns(string)
func (_Strategy *StrategyCallerSession) Explanation() (string, error) {
	return _Strategy.Contract.Explanation(&_Strategy.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_Strategy *StrategyCaller) GetTVLLimits(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "getTVLLimits")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_Strategy *StrategySession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _Strategy.Contract.GetTVLLimits(&_Strategy.CallOpts)
}

// GetTVLLimits is a free data retrieval call binding the contract method 0xdf6fadc1.
//
// Solidity: function getTVLLimits() view returns(uint256, uint256)
func (_Strategy *StrategyCallerSession) GetTVLLimits() (*big.Int, *big.Int, error) {
	return _Strategy.Contract.GetTVLLimits(&_Strategy.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_Strategy *StrategyCaller) MaxPerDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "maxPerDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_Strategy *StrategySession) MaxPerDeposit() (*big.Int, error) {
	return _Strategy.Contract.MaxPerDeposit(&_Strategy.CallOpts)
}

// MaxPerDeposit is a free data retrieval call binding the contract method 0x43fe08b0.
//
// Solidity: function maxPerDeposit() view returns(uint256)
func (_Strategy *StrategyCallerSession) MaxPerDeposit() (*big.Int, error) {
	return _Strategy.Contract.MaxPerDeposit(&_Strategy.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_Strategy *StrategyCaller) MaxTotalDeposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "maxTotalDeposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_Strategy *StrategySession) MaxTotalDeposits() (*big.Int, error) {
	return _Strategy.Contract.MaxTotalDeposits(&_Strategy.CallOpts)
}

// MaxTotalDeposits is a free data retrieval call binding the contract method 0x61b01b5d.
//
// Solidity: function maxTotalDeposits() view returns(uint256)
func (_Strategy *StrategyCallerSession) MaxTotalDeposits() (*big.Int, error) {
	return _Strategy.Contract.MaxTotalDeposits(&_Strategy.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Strategy *StrategyCaller) Paused(opts *bind.CallOpts, index uint8) (bool, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "paused", index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Strategy *StrategySession) Paused(index uint8) (bool, error) {
	return _Strategy.Contract.Paused(&_Strategy.CallOpts, index)
}

// Paused is a free data retrieval call binding the contract method 0x5ac86ab7.
//
// Solidity: function paused(uint8 index) view returns(bool)
func (_Strategy *StrategyCallerSession) Paused(index uint8) (bool, error) {
	return _Strategy.Contract.Paused(&_Strategy.CallOpts, index)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Strategy *StrategyCaller) Paused0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "paused0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Strategy *StrategySession) Paused0() (*big.Int, error) {
	return _Strategy.Contract.Paused0(&_Strategy.CallOpts)
}

// Paused0 is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(uint256)
func (_Strategy *StrategyCallerSession) Paused0() (*big.Int, error) {
	return _Strategy.Contract.Paused0(&_Strategy.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Strategy *StrategyCaller) PauserRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "pauserRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Strategy *StrategySession) PauserRegistry() (common.Address, error) {
	return _Strategy.Contract.PauserRegistry(&_Strategy.CallOpts)
}

// PauserRegistry is a free data retrieval call binding the contract method 0x886f1195.
//
// Solidity: function pauserRegistry() view returns(address)
func (_Strategy *StrategyCallerSession) PauserRegistry() (common.Address, error) {
	return _Strategy.Contract.PauserRegistry(&_Strategy.CallOpts)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_Strategy *StrategyCaller) Shares(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "shares", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_Strategy *StrategySession) Shares(user common.Address) (*big.Int, error) {
	return _Strategy.Contract.Shares(&_Strategy.CallOpts, user)
}

// Shares is a free data retrieval call binding the contract method 0xce7c2ac2.
//
// Solidity: function shares(address user) view returns(uint256)
func (_Strategy *StrategyCallerSession) Shares(user common.Address) (*big.Int, error) {
	return _Strategy.Contract.Shares(&_Strategy.CallOpts, user)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_Strategy *StrategyCaller) SharesToUnderlying(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "sharesToUnderlying", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_Strategy *StrategySession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _Strategy.Contract.SharesToUnderlying(&_Strategy.CallOpts, amountShares)
}

// SharesToUnderlying is a free data retrieval call binding the contract method 0xf3e73875.
//
// Solidity: function sharesToUnderlying(uint256 amountShares) view returns(uint256)
func (_Strategy *StrategyCallerSession) SharesToUnderlying(amountShares *big.Int) (*big.Int, error) {
	return _Strategy.Contract.SharesToUnderlying(&_Strategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_Strategy *StrategyCaller) SharesToUnderlyingView(opts *bind.CallOpts, amountShares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "sharesToUnderlyingView", amountShares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_Strategy *StrategySession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _Strategy.Contract.SharesToUnderlyingView(&_Strategy.CallOpts, amountShares)
}

// SharesToUnderlyingView is a free data retrieval call binding the contract method 0x7a8b2637.
//
// Solidity: function sharesToUnderlyingView(uint256 amountShares) view returns(uint256)
func (_Strategy *StrategyCallerSession) SharesToUnderlyingView(amountShares *big.Int) (*big.Int, error) {
	return _Strategy.Contract.SharesToUnderlyingView(&_Strategy.CallOpts, amountShares)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Strategy *StrategyCaller) StrategyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "strategyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Strategy *StrategySession) StrategyManager() (common.Address, error) {
	return _Strategy.Contract.StrategyManager(&_Strategy.CallOpts)
}

// StrategyManager is a free data retrieval call binding the contract method 0x39b70e38.
//
// Solidity: function strategyManager() view returns(address)
func (_Strategy *StrategyCallerSession) StrategyManager() (common.Address, error) {
	return _Strategy.Contract.StrategyManager(&_Strategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Strategy *StrategyCaller) TotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "totalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Strategy *StrategySession) TotalShares() (*big.Int, error) {
	return _Strategy.Contract.TotalShares(&_Strategy.CallOpts)
}

// TotalShares is a free data retrieval call binding the contract method 0x3a98ef39.
//
// Solidity: function totalShares() view returns(uint256)
func (_Strategy *StrategyCallerSession) TotalShares() (*big.Int, error) {
	return _Strategy.Contract.TotalShares(&_Strategy.CallOpts)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_Strategy *StrategyCaller) UnderlyingToShares(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "underlyingToShares", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_Strategy *StrategySession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _Strategy.Contract.UnderlyingToShares(&_Strategy.CallOpts, amountUnderlying)
}

// UnderlyingToShares is a free data retrieval call binding the contract method 0x8c871019.
//
// Solidity: function underlyingToShares(uint256 amountUnderlying) view returns(uint256)
func (_Strategy *StrategyCallerSession) UnderlyingToShares(amountUnderlying *big.Int) (*big.Int, error) {
	return _Strategy.Contract.UnderlyingToShares(&_Strategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_Strategy *StrategyCaller) UnderlyingToSharesView(opts *bind.CallOpts, amountUnderlying *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "underlyingToSharesView", amountUnderlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_Strategy *StrategySession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _Strategy.Contract.UnderlyingToSharesView(&_Strategy.CallOpts, amountUnderlying)
}

// UnderlyingToSharesView is a free data retrieval call binding the contract method 0xe3dae51c.
//
// Solidity: function underlyingToSharesView(uint256 amountUnderlying) view returns(uint256)
func (_Strategy *StrategyCallerSession) UnderlyingToSharesView(amountUnderlying *big.Int) (*big.Int, error) {
	return _Strategy.Contract.UnderlyingToSharesView(&_Strategy.CallOpts, amountUnderlying)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Strategy *StrategyCaller) UnderlyingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "underlyingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Strategy *StrategySession) UnderlyingToken() (common.Address, error) {
	return _Strategy.Contract.UnderlyingToken(&_Strategy.CallOpts)
}

// UnderlyingToken is a free data retrieval call binding the contract method 0x2495a599.
//
// Solidity: function underlyingToken() view returns(address)
func (_Strategy *StrategyCallerSession) UnderlyingToken() (common.Address, error) {
	return _Strategy.Contract.UnderlyingToken(&_Strategy.CallOpts)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_Strategy *StrategyCaller) UserUnderlyingView(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Strategy.contract.Call(opts, &out, "userUnderlyingView", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_Strategy *StrategySession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _Strategy.Contract.UserUnderlyingView(&_Strategy.CallOpts, user)
}

// UserUnderlyingView is a free data retrieval call binding the contract method 0x553ca5f8.
//
// Solidity: function userUnderlyingView(address user) view returns(uint256)
func (_Strategy *StrategyCallerSession) UserUnderlyingView(user common.Address) (*big.Int, error) {
	return _Strategy.Contract.UserUnderlyingView(&_Strategy.CallOpts, user)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_Strategy *StrategyTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_Strategy *StrategySession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Deposit(&_Strategy.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) returns(uint256 newShares)
func (_Strategy *StrategyTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Deposit(&_Strategy.TransactOpts, token, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken, address _pauserRegistry) returns()
func (_Strategy *StrategyTransactor) Initialize(opts *bind.TransactOpts, _maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "initialize", _maxPerDeposit, _maxTotalDeposits, _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken, address _pauserRegistry) returns()
func (_Strategy *StrategySession) Initialize(_maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.Initialize(&_Strategy.TransactOpts, _maxPerDeposit, _maxTotalDeposits, _underlyingToken, _pauserRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x019e2729.
//
// Solidity: function initialize(uint256 _maxPerDeposit, uint256 _maxTotalDeposits, address _underlyingToken, address _pauserRegistry) returns()
func (_Strategy *StrategyTransactorSession) Initialize(_maxPerDeposit *big.Int, _maxTotalDeposits *big.Int, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.Initialize(&_Strategy.TransactOpts, _maxPerDeposit, _maxTotalDeposits, _underlyingToken, _pauserRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_Strategy *StrategyTransactor) Initialize0(opts *bind.TransactOpts, _underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "initialize0", _underlyingToken, _pauserRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_Strategy *StrategySession) Initialize0(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.Initialize0(&_Strategy.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _underlyingToken, address _pauserRegistry) returns()
func (_Strategy *StrategyTransactorSession) Initialize0(_underlyingToken common.Address, _pauserRegistry common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.Initialize0(&_Strategy.TransactOpts, _underlyingToken, _pauserRegistry)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Strategy *StrategyTransactor) Pause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "pause", newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Strategy *StrategySession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Pause(&_Strategy.TransactOpts, newPausedStatus)
}

// Pause is a paid mutator transaction binding the contract method 0x136439dd.
//
// Solidity: function pause(uint256 newPausedStatus) returns()
func (_Strategy *StrategyTransactorSession) Pause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Pause(&_Strategy.TransactOpts, newPausedStatus)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Strategy *StrategyTransactor) PauseAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "pauseAll")
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Strategy *StrategySession) PauseAll() (*types.Transaction, error) {
	return _Strategy.Contract.PauseAll(&_Strategy.TransactOpts)
}

// PauseAll is a paid mutator transaction binding the contract method 0x595c6a67.
//
// Solidity: function pauseAll() returns()
func (_Strategy *StrategyTransactorSession) PauseAll() (*types.Transaction, error) {
	return _Strategy.Contract.PauseAll(&_Strategy.TransactOpts)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Strategy *StrategyTransactor) SetPauserRegistry(opts *bind.TransactOpts, newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "setPauserRegistry", newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Strategy *StrategySession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.SetPauserRegistry(&_Strategy.TransactOpts, newPauserRegistry)
}

// SetPauserRegistry is a paid mutator transaction binding the contract method 0x10d67a2f.
//
// Solidity: function setPauserRegistry(address newPauserRegistry) returns()
func (_Strategy *StrategyTransactorSession) SetPauserRegistry(newPauserRegistry common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.SetPauserRegistry(&_Strategy.TransactOpts, newPauserRegistry)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_Strategy *StrategyTransactor) SetTVLLimits(opts *bind.TransactOpts, newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "setTVLLimits", newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_Strategy *StrategySession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.SetTVLLimits(&_Strategy.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// SetTVLLimits is a paid mutator transaction binding the contract method 0x11c70c9d.
//
// Solidity: function setTVLLimits(uint256 newMaxPerDeposit, uint256 newMaxTotalDeposits) returns()
func (_Strategy *StrategyTransactorSession) SetTVLLimits(newMaxPerDeposit *big.Int, newMaxTotalDeposits *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.SetTVLLimits(&_Strategy.TransactOpts, newMaxPerDeposit, newMaxTotalDeposits)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Strategy *StrategyTransactor) Unpause(opts *bind.TransactOpts, newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "unpause", newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Strategy *StrategySession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Unpause(&_Strategy.TransactOpts, newPausedStatus)
}

// Unpause is a paid mutator transaction binding the contract method 0xfabc1cbc.
//
// Solidity: function unpause(uint256 newPausedStatus) returns()
func (_Strategy *StrategyTransactorSession) Unpause(newPausedStatus *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Unpause(&_Strategy.TransactOpts, newPausedStatus)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_Strategy *StrategyTransactor) UserUnderlying(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "userUnderlying", user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_Strategy *StrategySession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.UserUnderlying(&_Strategy.TransactOpts, user)
}

// UserUnderlying is a paid mutator transaction binding the contract method 0x8f6a6240.
//
// Solidity: function userUnderlying(address user) returns(uint256)
func (_Strategy *StrategyTransactorSession) UserUnderlying(user common.Address) (*types.Transaction, error) {
	return _Strategy.Contract.UserUnderlying(&_Strategy.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address depositor, address token, uint256 amountShares) returns()
func (_Strategy *StrategyTransactor) Withdraw(opts *bind.TransactOpts, depositor common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _Strategy.contract.Transact(opts, "withdraw", depositor, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address depositor, address token, uint256 amountShares) returns()
func (_Strategy *StrategySession) Withdraw(depositor common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Withdraw(&_Strategy.TransactOpts, depositor, token, amountShares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address depositor, address token, uint256 amountShares) returns()
func (_Strategy *StrategyTransactorSession) Withdraw(depositor common.Address, token common.Address, amountShares *big.Int) (*types.Transaction, error) {
	return _Strategy.Contract.Withdraw(&_Strategy.TransactOpts, depositor, token, amountShares)
}

// StrategyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Strategy contract.
type StrategyInitializedIterator struct {
	Event *StrategyInitialized // Event containing the contract specifics and raw log

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
func (it *StrategyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyInitialized)
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
		it.Event = new(StrategyInitialized)
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
func (it *StrategyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyInitialized represents a Initialized event raised by the Strategy contract.
type StrategyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Strategy *StrategyFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrategyInitializedIterator, error) {

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrategyInitializedIterator{contract: _Strategy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Strategy *StrategyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrategyInitialized) (event.Subscription, error) {

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyInitialized)
				if err := _Strategy.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Strategy *StrategyFilterer) ParseInitialized(log types.Log) (*StrategyInitialized, error) {
	event := new(StrategyInitialized)
	if err := _Strategy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyMaxPerDepositUpdatedIterator is returned from FilterMaxPerDepositUpdated and is used to iterate over the raw logs and unpacked data for MaxPerDepositUpdated events raised by the Strategy contract.
type StrategyMaxPerDepositUpdatedIterator struct {
	Event *StrategyMaxPerDepositUpdated // Event containing the contract specifics and raw log

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
func (it *StrategyMaxPerDepositUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyMaxPerDepositUpdated)
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
		it.Event = new(StrategyMaxPerDepositUpdated)
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
func (it *StrategyMaxPerDepositUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyMaxPerDepositUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyMaxPerDepositUpdated represents a MaxPerDepositUpdated event raised by the Strategy contract.
type StrategyMaxPerDepositUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxPerDepositUpdated is a free log retrieval operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_Strategy *StrategyFilterer) FilterMaxPerDepositUpdated(opts *bind.FilterOpts) (*StrategyMaxPerDepositUpdatedIterator, error) {

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return &StrategyMaxPerDepositUpdatedIterator{contract: _Strategy.contract, event: "MaxPerDepositUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPerDepositUpdated is a free log subscription operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_Strategy *StrategyFilterer) WatchMaxPerDepositUpdated(opts *bind.WatchOpts, sink chan<- *StrategyMaxPerDepositUpdated) (event.Subscription, error) {

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "MaxPerDepositUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyMaxPerDepositUpdated)
				if err := _Strategy.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
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

// ParseMaxPerDepositUpdated is a log parse operation binding the contract event 0xf97ed4e083acac67830025ecbc756d8fe847cdbdca4cee3fe1e128e98b54ecb5.
//
// Solidity: event MaxPerDepositUpdated(uint256 previousValue, uint256 newValue)
func (_Strategy *StrategyFilterer) ParseMaxPerDepositUpdated(log types.Log) (*StrategyMaxPerDepositUpdated, error) {
	event := new(StrategyMaxPerDepositUpdated)
	if err := _Strategy.contract.UnpackLog(event, "MaxPerDepositUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyMaxTotalDepositsUpdatedIterator is returned from FilterMaxTotalDepositsUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalDepositsUpdated events raised by the Strategy contract.
type StrategyMaxTotalDepositsUpdatedIterator struct {
	Event *StrategyMaxTotalDepositsUpdated // Event containing the contract specifics and raw log

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
func (it *StrategyMaxTotalDepositsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyMaxTotalDepositsUpdated)
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
		it.Event = new(StrategyMaxTotalDepositsUpdated)
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
func (it *StrategyMaxTotalDepositsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyMaxTotalDepositsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyMaxTotalDepositsUpdated represents a MaxTotalDepositsUpdated event raised by the Strategy contract.
type StrategyMaxTotalDepositsUpdated struct {
	PreviousValue *big.Int
	NewValue      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalDepositsUpdated is a free log retrieval operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_Strategy *StrategyFilterer) FilterMaxTotalDepositsUpdated(opts *bind.FilterOpts) (*StrategyMaxTotalDepositsUpdatedIterator, error) {

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return &StrategyMaxTotalDepositsUpdatedIterator{contract: _Strategy.contract, event: "MaxTotalDepositsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalDepositsUpdated is a free log subscription operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_Strategy *StrategyFilterer) WatchMaxTotalDepositsUpdated(opts *bind.WatchOpts, sink chan<- *StrategyMaxTotalDepositsUpdated) (event.Subscription, error) {

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "MaxTotalDepositsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyMaxTotalDepositsUpdated)
				if err := _Strategy.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
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

// ParseMaxTotalDepositsUpdated is a log parse operation binding the contract event 0x6ab181e0440bfbf4bacdf2e99674735ce6638005490688c5f994f5399353e452.
//
// Solidity: event MaxTotalDepositsUpdated(uint256 previousValue, uint256 newValue)
func (_Strategy *StrategyFilterer) ParseMaxTotalDepositsUpdated(log types.Log) (*StrategyMaxTotalDepositsUpdated, error) {
	event := new(StrategyMaxTotalDepositsUpdated)
	if err := _Strategy.contract.UnpackLog(event, "MaxTotalDepositsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Strategy contract.
type StrategyPausedIterator struct {
	Event *StrategyPaused // Event containing the contract specifics and raw log

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
func (it *StrategyPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyPaused)
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
		it.Event = new(StrategyPaused)
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
func (it *StrategyPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyPaused represents a Paused event raised by the Strategy contract.
type StrategyPaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Strategy *StrategyFilterer) FilterPaused(opts *bind.FilterOpts, account []common.Address) (*StrategyPausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyPausedIterator{contract: _Strategy.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Strategy *StrategyFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StrategyPaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "Paused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyPaused)
				if err := _Strategy.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address indexed account, uint256 newPausedStatus)
func (_Strategy *StrategyFilterer) ParsePaused(log types.Log) (*StrategyPaused, error) {
	event := new(StrategyPaused)
	if err := _Strategy.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyPauserRegistrySetIterator is returned from FilterPauserRegistrySet and is used to iterate over the raw logs and unpacked data for PauserRegistrySet events raised by the Strategy contract.
type StrategyPauserRegistrySetIterator struct {
	Event *StrategyPauserRegistrySet // Event containing the contract specifics and raw log

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
func (it *StrategyPauserRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyPauserRegistrySet)
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
		it.Event = new(StrategyPauserRegistrySet)
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
func (it *StrategyPauserRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyPauserRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyPauserRegistrySet represents a PauserRegistrySet event raised by the Strategy contract.
type StrategyPauserRegistrySet struct {
	PauserRegistry    common.Address
	NewPauserRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPauserRegistrySet is a free log retrieval operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Strategy *StrategyFilterer) FilterPauserRegistrySet(opts *bind.FilterOpts) (*StrategyPauserRegistrySetIterator, error) {

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return &StrategyPauserRegistrySetIterator{contract: _Strategy.contract, event: "PauserRegistrySet", logs: logs, sub: sub}, nil
}

// WatchPauserRegistrySet is a free log subscription operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Strategy *StrategyFilterer) WatchPauserRegistrySet(opts *bind.WatchOpts, sink chan<- *StrategyPauserRegistrySet) (event.Subscription, error) {

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "PauserRegistrySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyPauserRegistrySet)
				if err := _Strategy.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
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

// ParsePauserRegistrySet is a log parse operation binding the contract event 0x6e9fcd539896fca60e8b0f01dd580233e48a6b0f7df013b89ba7f565869acdb6.
//
// Solidity: event PauserRegistrySet(address pauserRegistry, address newPauserRegistry)
func (_Strategy *StrategyFilterer) ParsePauserRegistrySet(log types.Log) (*StrategyPauserRegistrySet, error) {
	event := new(StrategyPauserRegistrySet)
	if err := _Strategy.contract.UnpackLog(event, "PauserRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrategyUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Strategy contract.
type StrategyUnpausedIterator struct {
	Event *StrategyUnpaused // Event containing the contract specifics and raw log

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
func (it *StrategyUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrategyUnpaused)
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
		it.Event = new(StrategyUnpaused)
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
func (it *StrategyUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrategyUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrategyUnpaused represents a Unpaused event raised by the Strategy contract.
type StrategyUnpaused struct {
	Account         common.Address
	NewPausedStatus *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Strategy *StrategyFilterer) FilterUnpaused(opts *bind.FilterOpts, account []common.Address) (*StrategyUnpausedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Strategy.contract.FilterLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return &StrategyUnpausedIterator{contract: _Strategy.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Strategy *StrategyFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StrategyUnpaused, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Strategy.contract.WatchLogs(opts, "Unpaused", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrategyUnpaused)
				if err := _Strategy.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x3582d1828e26bf56bd801502bc021ac0bc8afb57c826e4986b45593c8fad389c.
//
// Solidity: event Unpaused(address indexed account, uint256 newPausedStatus)
func (_Strategy *StrategyFilterer) ParseUnpaused(log types.Log) (*StrategyUnpaused, error) {
	event := new(StrategyUnpaused)
	if err := _Strategy.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
