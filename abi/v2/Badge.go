// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v2

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

// BadgeMetaData contains all meta data concerning the Badge contract.
var BadgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyHoldsBadge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCreator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonexistentQuest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuestIdAlreadyExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Donation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"MinterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"URIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addrHoldQuestBadge\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"questBadgeNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BadgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BadgeMetaData.ABI instead.
var BadgeABI = BadgeMetaData.ABI

// Badge is an auto generated Go binding around an Ethereum contract.
type Badge struct {
	BadgeCaller     // Read-only binding to the contract
	BadgeTransactor // Write-only binding to the contract
	BadgeFilterer   // Log filterer for contract events
}

// BadgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BadgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BadgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BadgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BadgeSession struct {
	Contract     *Badge            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BadgeCallerSession struct {
	Contract *BadgeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BadgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BadgeTransactorSession struct {
	Contract     *BadgeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BadgeRaw struct {
	Contract *Badge // Generic contract binding to access the raw methods on
}

// BadgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BadgeCallerRaw struct {
	Contract *BadgeCaller // Generic read-only contract binding to access the raw methods on
}

// BadgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BadgeTransactorRaw struct {
	Contract *BadgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBadge creates a new instance of Badge, bound to a specific deployed contract.
func NewBadge(address common.Address, backend bind.ContractBackend) (*Badge, error) {
	contract, err := bindBadge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Badge{BadgeCaller: BadgeCaller{contract: contract}, BadgeTransactor: BadgeTransactor{contract: contract}, BadgeFilterer: BadgeFilterer{contract: contract}}, nil
}

// NewBadgeCaller creates a new read-only instance of Badge, bound to a specific deployed contract.
func NewBadgeCaller(address common.Address, caller bind.ContractCaller) (*BadgeCaller, error) {
	contract, err := bindBadge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeCaller{contract: contract}, nil
}

// NewBadgeTransactor creates a new write-only instance of Badge, bound to a specific deployed contract.
func NewBadgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BadgeTransactor, error) {
	contract, err := bindBadge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeTransactor{contract: contract}, nil
}

// NewBadgeFilterer creates a new log filterer instance of Badge, bound to a specific deployed contract.
func NewBadgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BadgeFilterer, error) {
	contract, err := bindBadge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BadgeFilterer{contract: contract}, nil
}

// bindBadge binds a generic wrapper to an already deployed contract.
func bindBadge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BadgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Badge *BadgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Badge.Contract.BadgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Badge *BadgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badge.Contract.BadgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Badge *BadgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Badge.Contract.BadgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Badge *BadgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Badge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Badge *BadgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Badge *BadgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Badge.Contract.contract.Transact(opts, method, params...)
}

// AddrHoldQuestBadge is a free data retrieval call binding the contract method 0x7c7f07ff.
//
// Solidity: function addrHoldQuestBadge(address , uint256 ) view returns(bool)
func (_Badge *BadgeCaller) AddrHoldQuestBadge(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "addrHoldQuestBadge", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AddrHoldQuestBadge is a free data retrieval call binding the contract method 0x7c7f07ff.
//
// Solidity: function addrHoldQuestBadge(address , uint256 ) view returns(bool)
func (_Badge *BadgeSession) AddrHoldQuestBadge(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Badge.Contract.AddrHoldQuestBadge(&_Badge.CallOpts, arg0, arg1)
}

// AddrHoldQuestBadge is a free data retrieval call binding the contract method 0x7c7f07ff.
//
// Solidity: function addrHoldQuestBadge(address , uint256 ) view returns(bool)
func (_Badge *BadgeCallerSession) AddrHoldQuestBadge(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Badge.Contract.AddrHoldQuestBadge(&_Badge.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Badge *BadgeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Badge *BadgeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Badge.Contract.BalanceOf(&_Badge.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Badge *BadgeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Badge.Contract.BalanceOf(&_Badge.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Badge *BadgeCaller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Badge *BadgeSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Badge.Contract.GetApproved(&_Badge.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Badge *BadgeCallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Badge.Contract.GetApproved(&_Badge.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Badge *BadgeCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Badge *BadgeSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Badge.Contract.IsApprovedForAll(&_Badge.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Badge *BadgeCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Badge.Contract.IsApprovedForAll(&_Badge.CallOpts, arg0, arg1)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_Badge *BadgeCaller) Locked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "locked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_Badge *BadgeSession) Locked(tokenId *big.Int) (bool, error) {
	return _Badge.Contract.Locked(&_Badge.CallOpts, tokenId)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_Badge *BadgeCallerSession) Locked(tokenId *big.Int) (bool, error) {
	return _Badge.Contract.Locked(&_Badge.CallOpts, tokenId)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_Badge *BadgeCaller) Minters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_Badge *BadgeSession) Minters(arg0 common.Address) (bool, error) {
	return _Badge.Contract.Minters(&_Badge.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_Badge *BadgeCallerSession) Minters(arg0 common.Address) (bool, error) {
	return _Badge.Contract.Minters(&_Badge.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badge *BadgeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badge *BadgeSession) Name() (string, error) {
	return _Badge.Contract.Name(&_Badge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Badge *BadgeCallerSession) Name() (string, error) {
	return _Badge.Contract.Name(&_Badge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Badge *BadgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Badge *BadgeSession) Owner() (common.Address, error) {
	return _Badge.Contract.Owner(&_Badge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Badge *BadgeCallerSession) Owner() (common.Address, error) {
	return _Badge.Contract.Owner(&_Badge.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Badge *BadgeCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Badge *BadgeSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Badge.Contract.OwnerOf(&_Badge.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Badge *BadgeCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Badge.Contract.OwnerOf(&_Badge.CallOpts, tokenId)
}

// QuestBadgeNum is a free data retrieval call binding the contract method 0xad113eb3.
//
// Solidity: function questBadgeNum(uint256 ) view returns(uint256)
func (_Badge *BadgeCaller) QuestBadgeNum(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "questBadgeNum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuestBadgeNum is a free data retrieval call binding the contract method 0xad113eb3.
//
// Solidity: function questBadgeNum(uint256 ) view returns(uint256)
func (_Badge *BadgeSession) QuestBadgeNum(arg0 *big.Int) (*big.Int, error) {
	return _Badge.Contract.QuestBadgeNum(&_Badge.CallOpts, arg0)
}

// QuestBadgeNum is a free data retrieval call binding the contract method 0xad113eb3.
//
// Solidity: function questBadgeNum(uint256 ) view returns(uint256)
func (_Badge *BadgeCallerSession) QuestBadgeNum(arg0 *big.Int) (*big.Int, error) {
	return _Badge.Contract.QuestBadgeNum(&_Badge.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Badge *BadgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Badge *BadgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Badge.Contract.SupportsInterface(&_Badge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Badge *BadgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Badge.Contract.SupportsInterface(&_Badge.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badge *BadgeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badge *BadgeSession) Symbol() (string, error) {
	return _Badge.Contract.Symbol(&_Badge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Badge *BadgeCallerSession) Symbol() (string, error) {
	return _Badge.Contract.Symbol(&_Badge.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Badge *BadgeCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Badge *BadgeSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Badge.Contract.TokenURI(&_Badge.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Badge *BadgeCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Badge.Contract.TokenURI(&_Badge.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badge *BadgeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Badge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badge *BadgeSession) TotalSupply() (*big.Int, error) {
	return _Badge.Contract.TotalSupply(&_Badge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Badge *BadgeCallerSession) TotalSupply() (*big.Int, error) {
	return _Badge.Contract.TotalSupply(&_Badge.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_Badge *BadgeTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_Badge *BadgeSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.Approve(&_Badge.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_Badge *BadgeTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.Approve(&_Badge.TransactOpts, arg0, arg1)
}

// Claim is a paid mutator transaction binding the contract method 0xb80ad9a5.
//
// Solidity: function claim(address to, uint256 questId, string uri) returns()
func (_Badge *BadgeTransactor) Claim(opts *bind.TransactOpts, to common.Address, questId *big.Int, uri string) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "claim", to, questId, uri)
}

// Claim is a paid mutator transaction binding the contract method 0xb80ad9a5.
//
// Solidity: function claim(address to, uint256 questId, string uri) returns()
func (_Badge *BadgeSession) Claim(to common.Address, questId *big.Int, uri string) (*types.Transaction, error) {
	return _Badge.Contract.Claim(&_Badge.TransactOpts, to, questId, uri)
}

// Claim is a paid mutator transaction binding the contract method 0xb80ad9a5.
//
// Solidity: function claim(address to, uint256 questId, string uri) returns()
func (_Badge *BadgeTransactorSession) Claim(to common.Address, questId *big.Int, uri string) (*types.Transaction, error) {
	return _Badge.Contract.Claim(&_Badge.TransactOpts, to, questId, uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Badge *BadgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Badge *BadgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Badge.Contract.RenounceOwnership(&_Badge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Badge *BadgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Badge.Contract.RenounceOwnership(&_Badge.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_Badge *BadgeTransactor) SafeTransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "safeTransferFrom", arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_Badge *BadgeSession) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.SafeTransferFrom(&_Badge.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_Badge *BadgeTransactorSession) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.SafeTransferFrom(&_Badge.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_Badge *BadgeTransactor) SafeTransferFrom0(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "safeTransferFrom0", arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_Badge *BadgeSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Badge.Contract.SafeTransferFrom0(&_Badge.TransactOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_Badge *BadgeTransactorSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Badge.Contract.SafeTransferFrom0(&_Badge.TransactOpts, arg0, arg1, arg2, arg3)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_Badge *BadgeTransactor) SetApprovalForAll(opts *bind.TransactOpts, arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "setApprovalForAll", arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_Badge *BadgeSession) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _Badge.Contract.SetApprovalForAll(&_Badge.TransactOpts, arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_Badge *BadgeTransactorSession) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _Badge.Contract.SetApprovalForAll(&_Badge.TransactOpts, arg0, arg1)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_Badge *BadgeTransactor) SetMinter(opts *bind.TransactOpts, minter common.Address, enabled bool) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "setMinter", minter, enabled)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_Badge *BadgeSession) SetMinter(minter common.Address, enabled bool) (*types.Transaction, error) {
	return _Badge.Contract.SetMinter(&_Badge.TransactOpts, minter, enabled)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_Badge *BadgeTransactorSession) SetMinter(minter common.Address, enabled bool) (*types.Transaction, error) {
	return _Badge.Contract.SetMinter(&_Badge.TransactOpts, minter, enabled)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_Badge *BadgeTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_Badge *BadgeSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.TransferFrom(&_Badge.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_Badge *BadgeTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Badge.Contract.TransferFrom(&_Badge.TransactOpts, arg0, arg1, arg2)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Badge *BadgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Badge *BadgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Badge.Contract.TransferOwnership(&_Badge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Badge *BadgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Badge.Contract.TransferOwnership(&_Badge.TransactOpts, newOwner)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x31d41c69.
//
// Solidity: function updateURI(uint256 tokenId, string uri) returns()
func (_Badge *BadgeTransactor) UpdateURI(opts *bind.TransactOpts, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Badge.contract.Transact(opts, "updateURI", tokenId, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x31d41c69.
//
// Solidity: function updateURI(uint256 tokenId, string uri) returns()
func (_Badge *BadgeSession) UpdateURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Badge.Contract.UpdateURI(&_Badge.TransactOpts, tokenId, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x31d41c69.
//
// Solidity: function updateURI(uint256 tokenId, string uri) returns()
func (_Badge *BadgeTransactorSession) UpdateURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Badge.Contract.UpdateURI(&_Badge.TransactOpts, tokenId, uri)
}

// BadgeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Badge contract.
type BadgeApprovalIterator struct {
	Event *BadgeApproval // Event containing the contract specifics and raw log

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
func (it *BadgeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeApproval)
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
		it.Event = new(BadgeApproval)
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
func (it *BadgeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeApproval represents a Approval event raised by the Badge contract.
type BadgeApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BadgeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeApprovalIterator{contract: _Badge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BadgeApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeApproval)
				if err := _Badge.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) ParseApproval(log types.Log) (*BadgeApproval, error) {
	event := new(BadgeApproval)
	if err := _Badge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Badge contract.
type BadgeApprovalForAllIterator struct {
	Event *BadgeApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BadgeApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeApprovalForAll)
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
		it.Event = new(BadgeApprovalForAll)
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
func (it *BadgeApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeApprovalForAll represents a ApprovalForAll event raised by the Badge contract.
type BadgeApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Badge *BadgeFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BadgeApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BadgeApprovalForAllIterator{contract: _Badge.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Badge *BadgeFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BadgeApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeApprovalForAll)
				if err := _Badge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Badge *BadgeFilterer) ParseApprovalForAll(log types.Log) (*BadgeApprovalForAll, error) {
	event := new(BadgeApprovalForAll)
	if err := _Badge.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the Badge contract.
type BadgeClaimedIterator struct {
	Event *BadgeClaimed // Event containing the contract specifics and raw log

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
func (it *BadgeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeClaimed)
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
		it.Event = new(BadgeClaimed)
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
func (it *BadgeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeClaimed represents a Claimed event raised by the Badge contract.
type BadgeClaimed struct {
	TokenId  *big.Int
	QuestId  *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xc09f7b3c1c5b70bf6d30e485a4525d625d4723aed5d319376825446ef135c7ca.
//
// Solidity: event Claimed(uint256 indexed tokenId, uint256 questId, address receiver)
func (_Badge *BadgeFilterer) FilterClaimed(opts *bind.FilterOpts, tokenId []*big.Int) (*BadgeClaimedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "Claimed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeClaimedIterator{contract: _Badge.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xc09f7b3c1c5b70bf6d30e485a4525d625d4723aed5d319376825446ef135c7ca.
//
// Solidity: event Claimed(uint256 indexed tokenId, uint256 questId, address receiver)
func (_Badge *BadgeFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *BadgeClaimed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "Claimed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeClaimed)
				if err := _Badge.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xc09f7b3c1c5b70bf6d30e485a4525d625d4723aed5d319376825446ef135c7ca.
//
// Solidity: event Claimed(uint256 indexed tokenId, uint256 questId, address receiver)
func (_Badge *BadgeFilterer) ParseClaimed(log types.Log) (*BadgeClaimed, error) {
	event := new(BadgeClaimed)
	if err := _Badge.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeDonationIterator is returned from FilterDonation and is used to iterate over the raw logs and unpacked data for Donation events raised by the Badge contract.
type BadgeDonationIterator struct {
	Event *BadgeDonation // Event containing the contract specifics and raw log

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
func (it *BadgeDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeDonation)
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
		it.Event = new(BadgeDonation)
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
func (it *BadgeDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeDonation represents a Donation event raised by the Badge contract.
type BadgeDonation struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonation is a free log retrieval operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_Badge *BadgeFilterer) FilterDonation(opts *bind.FilterOpts) (*BadgeDonationIterator, error) {

	logs, sub, err := _Badge.contract.FilterLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return &BadgeDonationIterator{contract: _Badge.contract, event: "Donation", logs: logs, sub: sub}, nil
}

// WatchDonation is a free log subscription operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_Badge *BadgeFilterer) WatchDonation(opts *bind.WatchOpts, sink chan<- *BadgeDonation) (event.Subscription, error) {

	logs, sub, err := _Badge.contract.WatchLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeDonation)
				if err := _Badge.contract.UnpackLog(event, "Donation", log); err != nil {
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

// ParseDonation is a log parse operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_Badge *BadgeFilterer) ParseDonation(log types.Log) (*BadgeDonation, error) {
	event := new(BadgeDonation)
	if err := _Badge.contract.UnpackLog(event, "Donation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the Badge contract.
type BadgeLockedIterator struct {
	Event *BadgeLocked // Event containing the contract specifics and raw log

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
func (it *BadgeLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeLocked)
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
		it.Event = new(BadgeLocked)
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
func (it *BadgeLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeLocked represents a Locked event raised by the Badge contract.
type BadgeLocked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_Badge *BadgeFilterer) FilterLocked(opts *bind.FilterOpts) (*BadgeLockedIterator, error) {

	logs, sub, err := _Badge.contract.FilterLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return &BadgeLockedIterator{contract: _Badge.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_Badge *BadgeFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BadgeLocked) (event.Subscription, error) {

	logs, sub, err := _Badge.contract.WatchLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeLocked)
				if err := _Badge.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_Badge *BadgeFilterer) ParseLocked(log types.Log) (*BadgeLocked, error) {
	event := new(BadgeLocked)
	if err := _Badge.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeMinterSetIterator is returned from FilterMinterSet and is used to iterate over the raw logs and unpacked data for MinterSet events raised by the Badge contract.
type BadgeMinterSetIterator struct {
	Event *BadgeMinterSet // Event containing the contract specifics and raw log

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
func (it *BadgeMinterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeMinterSet)
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
		it.Event = new(BadgeMinterSet)
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
func (it *BadgeMinterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeMinterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeMinterSet represents a MinterSet event raised by the Badge contract.
type BadgeMinterSet struct {
	Minter  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterSet is a free log retrieval operation binding the contract event 0x583b0aa0e528532caf4b907c11d7a8158a122fe2a6fb80cd9b09776ebea8d92d.
//
// Solidity: event MinterSet(address minter, bool enabled)
func (_Badge *BadgeFilterer) FilterMinterSet(opts *bind.FilterOpts) (*BadgeMinterSetIterator, error) {

	logs, sub, err := _Badge.contract.FilterLogs(opts, "MinterSet")
	if err != nil {
		return nil, err
	}
	return &BadgeMinterSetIterator{contract: _Badge.contract, event: "MinterSet", logs: logs, sub: sub}, nil
}

// WatchMinterSet is a free log subscription operation binding the contract event 0x583b0aa0e528532caf4b907c11d7a8158a122fe2a6fb80cd9b09776ebea8d92d.
//
// Solidity: event MinterSet(address minter, bool enabled)
func (_Badge *BadgeFilterer) WatchMinterSet(opts *bind.WatchOpts, sink chan<- *BadgeMinterSet) (event.Subscription, error) {

	logs, sub, err := _Badge.contract.WatchLogs(opts, "MinterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeMinterSet)
				if err := _Badge.contract.UnpackLog(event, "MinterSet", log); err != nil {
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

// ParseMinterSet is a log parse operation binding the contract event 0x583b0aa0e528532caf4b907c11d7a8158a122fe2a6fb80cd9b09776ebea8d92d.
//
// Solidity: event MinterSet(address minter, bool enabled)
func (_Badge *BadgeFilterer) ParseMinterSet(log types.Log) (*BadgeMinterSet, error) {
	event := new(BadgeMinterSet)
	if err := _Badge.contract.UnpackLog(event, "MinterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Badge contract.
type BadgeOwnershipTransferredIterator struct {
	Event *BadgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BadgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeOwnershipTransferred)
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
		it.Event = new(BadgeOwnershipTransferred)
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
func (it *BadgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeOwnershipTransferred represents a OwnershipTransferred event raised by the Badge contract.
type BadgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Badge *BadgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BadgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BadgeOwnershipTransferredIterator{contract: _Badge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Badge *BadgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BadgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeOwnershipTransferred)
				if err := _Badge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Badge *BadgeFilterer) ParseOwnershipTransferred(log types.Log) (*BadgeOwnershipTransferred, error) {
	event := new(BadgeOwnershipTransferred)
	if err := _Badge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Badge contract.
type BadgeTransferIterator struct {
	Event *BadgeTransfer // Event containing the contract specifics and raw log

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
func (it *BadgeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeTransfer)
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
		it.Event = new(BadgeTransfer)
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
func (it *BadgeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeTransfer represents a Transfer event raised by the Badge contract.
type BadgeTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BadgeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeTransferIterator{contract: _Badge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BadgeTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeTransfer)
				if err := _Badge.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Badge *BadgeFilterer) ParseTransfer(log types.Log) (*BadgeTransfer, error) {
	event := new(BadgeTransfer)
	if err := _Badge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeURIUpdatedIterator is returned from FilterURIUpdated and is used to iterate over the raw logs and unpacked data for URIUpdated events raised by the Badge contract.
type BadgeURIUpdatedIterator struct {
	Event *BadgeURIUpdated // Event containing the contract specifics and raw log

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
func (it *BadgeURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeURIUpdated)
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
		it.Event = new(BadgeURIUpdated)
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
func (it *BadgeURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeURIUpdated represents a URIUpdated event raised by the Badge contract.
type BadgeURIUpdated struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterURIUpdated is a free log retrieval operation binding the contract event 0xbf032ae7dcb484f99f9f3977ba9cde222756e7073de74244ceb89de862569bdb.
//
// Solidity: event URIUpdated(uint256 indexed tokenId, string uri)
func (_Badge *BadgeFilterer) FilterURIUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*BadgeURIUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.FilterLogs(opts, "URIUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeURIUpdatedIterator{contract: _Badge.contract, event: "URIUpdated", logs: logs, sub: sub}, nil
}

// WatchURIUpdated is a free log subscription operation binding the contract event 0xbf032ae7dcb484f99f9f3977ba9cde222756e7073de74244ceb89de862569bdb.
//
// Solidity: event URIUpdated(uint256 indexed tokenId, string uri)
func (_Badge *BadgeFilterer) WatchURIUpdated(opts *bind.WatchOpts, sink chan<- *BadgeURIUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Badge.contract.WatchLogs(opts, "URIUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeURIUpdated)
				if err := _Badge.contract.UnpackLog(event, "URIUpdated", log); err != nil {
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

// ParseURIUpdated is a log parse operation binding the contract event 0xbf032ae7dcb484f99f9f3977ba9cde222756e7073de74244ceb89de862569bdb.
//
// Solidity: event URIUpdated(uint256 indexed tokenId, string uri)
func (_Badge *BadgeFilterer) ParseURIUpdated(log types.Log) (*BadgeURIUpdated, error) {
	event := new(BadgeURIUpdated)
	if err := _Badge.contract.UnpackLog(event, "URIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the Badge contract.
type BadgeUnlockedIterator struct {
	Event *BadgeUnlocked // Event containing the contract specifics and raw log

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
func (it *BadgeUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeUnlocked)
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
		it.Event = new(BadgeUnlocked)
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
func (it *BadgeUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeUnlocked represents a Unlocked event raised by the Badge contract.
type BadgeUnlocked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 tokenId)
func (_Badge *BadgeFilterer) FilterUnlocked(opts *bind.FilterOpts) (*BadgeUnlockedIterator, error) {

	logs, sub, err := _Badge.contract.FilterLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return &BadgeUnlockedIterator{contract: _Badge.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 tokenId)
func (_Badge *BadgeFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *BadgeUnlocked) (event.Subscription, error) {

	logs, sub, err := _Badge.contract.WatchLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeUnlocked)
				if err := _Badge.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 tokenId)
func (_Badge *BadgeFilterer) ParseUnlocked(log types.Log) (*BadgeUnlocked, error) {
	event := new(BadgeUnlocked)
	if err := _Badge.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
