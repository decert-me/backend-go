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

// IBadgeQuestData is an auto generated low-level Go binding around an user-defined struct.
type IBadgeQuestData struct {
	Creator common.Address
	StartTs uint32
	EndTs   uint32
	Title   string
	Uri     string
}

// BadgeV2MetaData contains all meta data concerning the BadgeV2 contract.
var BadgeV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyHoldsBadge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonexistentQuest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotClaimedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuestIdAlreadyExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Donation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"MinterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIBadge.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"}],\"name\":\"QuestInit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIBadge.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"}],\"name\":\"QuestUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"URIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"addrToQuestToBadge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"badgeToQuest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIBadge.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"claimWithInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"}],\"name\":\"getBadgeNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"}],\"name\":\"getQuest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIBadge.QuestData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"questUri\",\"type\":\"string\"}],\"name\":\"updateQuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BadgeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use BadgeV2MetaData.ABI instead.
var BadgeV2ABI = BadgeV2MetaData.ABI

// BadgeV2 is an auto generated Go binding around an Ethereum contract.
type BadgeV2 struct {
	BadgeV2Caller     // Read-only binding to the contract
	BadgeV2Transactor // Write-only binding to the contract
	BadgeV2Filterer   // Log filterer for contract events
}

// BadgeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type BadgeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BadgeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BadgeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BadgeV2Session struct {
	Contract     *BadgeV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BadgeV2CallerSession struct {
	Contract *BadgeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BadgeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BadgeV2TransactorSession struct {
	Contract     *BadgeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BadgeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type BadgeV2Raw struct {
	Contract *BadgeV2 // Generic contract binding to access the raw methods on
}

// BadgeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BadgeV2CallerRaw struct {
	Contract *BadgeV2Caller // Generic read-only contract binding to access the raw methods on
}

// BadgeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BadgeV2TransactorRaw struct {
	Contract *BadgeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBadgeV2 creates a new instance of BadgeV2, bound to a specific deployed contract.
func NewBadgeV2(address common.Address, backend bind.ContractBackend) (*BadgeV2, error) {
	contract, err := bindBadgeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BadgeV2{BadgeV2Caller: BadgeV2Caller{contract: contract}, BadgeV2Transactor: BadgeV2Transactor{contract: contract}, BadgeV2Filterer: BadgeV2Filterer{contract: contract}}, nil
}

// NewBadgeV2Caller creates a new read-only instance of BadgeV2, bound to a specific deployed contract.
func NewBadgeV2Caller(address common.Address, caller bind.ContractCaller) (*BadgeV2Caller, error) {
	contract, err := bindBadgeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeV2Caller{contract: contract}, nil
}

// NewBadgeV2Transactor creates a new write-only instance of BadgeV2, bound to a specific deployed contract.
func NewBadgeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*BadgeV2Transactor, error) {
	contract, err := bindBadgeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeV2Transactor{contract: contract}, nil
}

// NewBadgeV2Filterer creates a new log filterer instance of BadgeV2, bound to a specific deployed contract.
func NewBadgeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*BadgeV2Filterer, error) {
	contract, err := bindBadgeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BadgeV2Filterer{contract: contract}, nil
}

// bindBadgeV2 binds a generic wrapper to an already deployed contract.
func bindBadgeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BadgeV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BadgeV2 *BadgeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BadgeV2.Contract.BadgeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BadgeV2 *BadgeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BadgeV2.Contract.BadgeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BadgeV2 *BadgeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BadgeV2.Contract.BadgeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BadgeV2 *BadgeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BadgeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BadgeV2 *BadgeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BadgeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BadgeV2 *BadgeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BadgeV2.Contract.contract.Transact(opts, method, params...)
}

// AddrToQuestToBadge is a free data retrieval call binding the contract method 0x3392470a.
//
// Solidity: function addrToQuestToBadge(address , uint256 ) view returns(uint256)
func (_BadgeV2 *BadgeV2Caller) AddrToQuestToBadge(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "addrToQuestToBadge", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddrToQuestToBadge is a free data retrieval call binding the contract method 0x3392470a.
//
// Solidity: function addrToQuestToBadge(address , uint256 ) view returns(uint256)
func (_BadgeV2 *BadgeV2Session) AddrToQuestToBadge(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BadgeV2.Contract.AddrToQuestToBadge(&_BadgeV2.CallOpts, arg0, arg1)
}

// AddrToQuestToBadge is a free data retrieval call binding the contract method 0x3392470a.
//
// Solidity: function addrToQuestToBadge(address , uint256 ) view returns(uint256)
func (_BadgeV2 *BadgeV2CallerSession) AddrToQuestToBadge(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _BadgeV2.Contract.AddrToQuestToBadge(&_BadgeV2.CallOpts, arg0, arg1)
}

// BadgeToQuest is a free data retrieval call binding the contract method 0xb5006862.
//
// Solidity: function badgeToQuest(uint256 ) view returns(uint256)
func (_BadgeV2 *BadgeV2Caller) BadgeToQuest(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "badgeToQuest", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BadgeToQuest is a free data retrieval call binding the contract method 0xb5006862.
//
// Solidity: function badgeToQuest(uint256 ) view returns(uint256)
func (_BadgeV2 *BadgeV2Session) BadgeToQuest(arg0 *big.Int) (*big.Int, error) {
	return _BadgeV2.Contract.BadgeToQuest(&_BadgeV2.CallOpts, arg0)
}

// BadgeToQuest is a free data retrieval call binding the contract method 0xb5006862.
//
// Solidity: function badgeToQuest(uint256 ) view returns(uint256)
func (_BadgeV2 *BadgeV2CallerSession) BadgeToQuest(arg0 *big.Int) (*big.Int, error) {
	return _BadgeV2.Contract.BadgeToQuest(&_BadgeV2.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BadgeV2 *BadgeV2Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BadgeV2 *BadgeV2Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BadgeV2.Contract.BalanceOf(&_BadgeV2.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BadgeV2 *BadgeV2CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BadgeV2.Contract.BalanceOf(&_BadgeV2.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_BadgeV2 *BadgeV2Caller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_BadgeV2 *BadgeV2Session) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _BadgeV2.Contract.GetApproved(&_BadgeV2.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_BadgeV2 *BadgeV2CallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _BadgeV2.Contract.GetApproved(&_BadgeV2.CallOpts, arg0)
}

// GetBadgeNum is a free data retrieval call binding the contract method 0x938a00d6.
//
// Solidity: function getBadgeNum(uint256 questId) view returns(uint256)
func (_BadgeV2 *BadgeV2Caller) GetBadgeNum(opts *bind.CallOpts, questId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "getBadgeNum", questId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBadgeNum is a free data retrieval call binding the contract method 0x938a00d6.
//
// Solidity: function getBadgeNum(uint256 questId) view returns(uint256)
func (_BadgeV2 *BadgeV2Session) GetBadgeNum(questId *big.Int) (*big.Int, error) {
	return _BadgeV2.Contract.GetBadgeNum(&_BadgeV2.CallOpts, questId)
}

// GetBadgeNum is a free data retrieval call binding the contract method 0x938a00d6.
//
// Solidity: function getBadgeNum(uint256 questId) view returns(uint256)
func (_BadgeV2 *BadgeV2CallerSession) GetBadgeNum(questId *big.Int) (*big.Int, error) {
	return _BadgeV2.Contract.GetBadgeNum(&_BadgeV2.CallOpts, questId)
}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 questId) view returns((address,uint32,uint32,string,string))
func (_BadgeV2 *BadgeV2Caller) GetQuest(opts *bind.CallOpts, questId *big.Int) (IBadgeQuestData, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "getQuest", questId)

	if err != nil {
		return *new(IBadgeQuestData), err
	}

	out0 := *abi.ConvertType(out[0], new(IBadgeQuestData)).(*IBadgeQuestData)

	return out0, err

}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 questId) view returns((address,uint32,uint32,string,string))
func (_BadgeV2 *BadgeV2Session) GetQuest(questId *big.Int) (IBadgeQuestData, error) {
	return _BadgeV2.Contract.GetQuest(&_BadgeV2.CallOpts, questId)
}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 questId) view returns((address,uint32,uint32,string,string))
func (_BadgeV2 *BadgeV2CallerSession) GetQuest(questId *big.Int) (IBadgeQuestData, error) {
	return _BadgeV2.Contract.GetQuest(&_BadgeV2.CallOpts, questId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_BadgeV2 *BadgeV2Caller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_BadgeV2 *BadgeV2Session) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _BadgeV2.Contract.IsApprovedForAll(&_BadgeV2.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_BadgeV2 *BadgeV2CallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _BadgeV2.Contract.IsApprovedForAll(&_BadgeV2.CallOpts, arg0, arg1)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_BadgeV2 *BadgeV2Caller) Locked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "locked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_BadgeV2 *BadgeV2Session) Locked(tokenId *big.Int) (bool, error) {
	return _BadgeV2.Contract.Locked(&_BadgeV2.CallOpts, tokenId)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_BadgeV2 *BadgeV2CallerSession) Locked(tokenId *big.Int) (bool, error) {
	return _BadgeV2.Contract.Locked(&_BadgeV2.CallOpts, tokenId)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_BadgeV2 *BadgeV2Caller) Minters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_BadgeV2 *BadgeV2Session) Minters(arg0 common.Address) (bool, error) {
	return _BadgeV2.Contract.Minters(&_BadgeV2.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_BadgeV2 *BadgeV2CallerSession) Minters(arg0 common.Address) (bool, error) {
	return _BadgeV2.Contract.Minters(&_BadgeV2.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BadgeV2 *BadgeV2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BadgeV2 *BadgeV2Session) Name() (string, error) {
	return _BadgeV2.Contract.Name(&_BadgeV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BadgeV2 *BadgeV2CallerSession) Name() (string, error) {
	return _BadgeV2.Contract.Name(&_BadgeV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BadgeV2 *BadgeV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BadgeV2 *BadgeV2Session) Owner() (common.Address, error) {
	return _BadgeV2.Contract.Owner(&_BadgeV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BadgeV2 *BadgeV2CallerSession) Owner() (common.Address, error) {
	return _BadgeV2.Contract.Owner(&_BadgeV2.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BadgeV2 *BadgeV2Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BadgeV2 *BadgeV2Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BadgeV2.Contract.OwnerOf(&_BadgeV2.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BadgeV2 *BadgeV2CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BadgeV2.Contract.OwnerOf(&_BadgeV2.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BadgeV2 *BadgeV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BadgeV2 *BadgeV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BadgeV2.Contract.SupportsInterface(&_BadgeV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BadgeV2 *BadgeV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BadgeV2.Contract.SupportsInterface(&_BadgeV2.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BadgeV2 *BadgeV2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BadgeV2 *BadgeV2Session) Symbol() (string, error) {
	return _BadgeV2.Contract.Symbol(&_BadgeV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BadgeV2 *BadgeV2CallerSession) Symbol() (string, error) {
	return _BadgeV2.Contract.Symbol(&_BadgeV2.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BadgeV2 *BadgeV2Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BadgeV2 *BadgeV2Session) TokenURI(tokenId *big.Int) (string, error) {
	return _BadgeV2.Contract.TokenURI(&_BadgeV2.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BadgeV2 *BadgeV2CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BadgeV2.Contract.TokenURI(&_BadgeV2.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BadgeV2 *BadgeV2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BadgeV2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BadgeV2 *BadgeV2Session) TotalSupply() (*big.Int, error) {
	return _BadgeV2.Contract.TotalSupply(&_BadgeV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BadgeV2 *BadgeV2CallerSession) TotalSupply() (*big.Int, error) {
	return _BadgeV2.Contract.TotalSupply(&_BadgeV2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_BadgeV2 *BadgeV2Transactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_BadgeV2 *BadgeV2Session) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _BadgeV2.Contract.Approve(&_BadgeV2.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_BadgeV2 *BadgeV2TransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _BadgeV2.Contract.Approve(&_BadgeV2.TransactOpts, arg0, arg1)
}

// Claim is a paid mutator transaction binding the contract method 0xb80ad9a5.
//
// Solidity: function claim(address to, uint256 questId, string uri) returns()
func (_BadgeV2 *BadgeV2Transactor) Claim(opts *bind.TransactOpts, to common.Address, questId *big.Int, uri string) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "claim", to, questId, uri)
}

// Claim is a paid mutator transaction binding the contract method 0xb80ad9a5.
//
// Solidity: function claim(address to, uint256 questId, string uri) returns()
func (_BadgeV2 *BadgeV2Session) Claim(to common.Address, questId *big.Int, uri string) (*types.Transaction, error) {
	return _BadgeV2.Contract.Claim(&_BadgeV2.TransactOpts, to, questId, uri)
}

// Claim is a paid mutator transaction binding the contract method 0xb80ad9a5.
//
// Solidity: function claim(address to, uint256 questId, string uri) returns()
func (_BadgeV2 *BadgeV2TransactorSession) Claim(to common.Address, questId *big.Int, uri string) (*types.Transaction, error) {
	return _BadgeV2.Contract.Claim(&_BadgeV2.TransactOpts, to, questId, uri)
}

// ClaimWithInit is a paid mutator transaction binding the contract method 0x9d5ea0ac.
//
// Solidity: function claimWithInit((address,uint32,uint32,string,string) questData, uint256 questId, address to, string uri) returns()
func (_BadgeV2 *BadgeV2Transactor) ClaimWithInit(opts *bind.TransactOpts, questData IBadgeQuestData, questId *big.Int, to common.Address, uri string) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "claimWithInit", questData, questId, to, uri)
}

// ClaimWithInit is a paid mutator transaction binding the contract method 0x9d5ea0ac.
//
// Solidity: function claimWithInit((address,uint32,uint32,string,string) questData, uint256 questId, address to, string uri) returns()
func (_BadgeV2 *BadgeV2Session) ClaimWithInit(questData IBadgeQuestData, questId *big.Int, to common.Address, uri string) (*types.Transaction, error) {
	return _BadgeV2.Contract.ClaimWithInit(&_BadgeV2.TransactOpts, questData, questId, to, uri)
}

// ClaimWithInit is a paid mutator transaction binding the contract method 0x9d5ea0ac.
//
// Solidity: function claimWithInit((address,uint32,uint32,string,string) questData, uint256 questId, address to, string uri) returns()
func (_BadgeV2 *BadgeV2TransactorSession) ClaimWithInit(questData IBadgeQuestData, questId *big.Int, to common.Address, uri string) (*types.Transaction, error) {
	return _BadgeV2.Contract.ClaimWithInit(&_BadgeV2.TransactOpts, questData, questId, to, uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BadgeV2 *BadgeV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BadgeV2 *BadgeV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _BadgeV2.Contract.RenounceOwnership(&_BadgeV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BadgeV2 *BadgeV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BadgeV2.Contract.RenounceOwnership(&_BadgeV2.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_BadgeV2 *BadgeV2Transactor) SafeTransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "safeTransferFrom", arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_BadgeV2 *BadgeV2Session) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _BadgeV2.Contract.SafeTransferFrom(&_BadgeV2.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_BadgeV2 *BadgeV2TransactorSession) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _BadgeV2.Contract.SafeTransferFrom(&_BadgeV2.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_BadgeV2 *BadgeV2Transactor) SafeTransferFrom0(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "safeTransferFrom0", arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_BadgeV2 *BadgeV2Session) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BadgeV2.Contract.SafeTransferFrom0(&_BadgeV2.TransactOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_BadgeV2 *BadgeV2TransactorSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _BadgeV2.Contract.SafeTransferFrom0(&_BadgeV2.TransactOpts, arg0, arg1, arg2, arg3)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_BadgeV2 *BadgeV2Transactor) SetApprovalForAll(opts *bind.TransactOpts, arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "setApprovalForAll", arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_BadgeV2 *BadgeV2Session) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _BadgeV2.Contract.SetApprovalForAll(&_BadgeV2.TransactOpts, arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_BadgeV2 *BadgeV2TransactorSession) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _BadgeV2.Contract.SetApprovalForAll(&_BadgeV2.TransactOpts, arg0, arg1)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_BadgeV2 *BadgeV2Transactor) SetMinter(opts *bind.TransactOpts, minter common.Address, enabled bool) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "setMinter", minter, enabled)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_BadgeV2 *BadgeV2Session) SetMinter(minter common.Address, enabled bool) (*types.Transaction, error) {
	return _BadgeV2.Contract.SetMinter(&_BadgeV2.TransactOpts, minter, enabled)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_BadgeV2 *BadgeV2TransactorSession) SetMinter(minter common.Address, enabled bool) (*types.Transaction, error) {
	return _BadgeV2.Contract.SetMinter(&_BadgeV2.TransactOpts, minter, enabled)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_BadgeV2 *BadgeV2Transactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_BadgeV2 *BadgeV2Session) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _BadgeV2.Contract.TransferFrom(&_BadgeV2.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_BadgeV2 *BadgeV2TransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _BadgeV2.Contract.TransferFrom(&_BadgeV2.TransactOpts, arg0, arg1, arg2)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BadgeV2 *BadgeV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BadgeV2 *BadgeV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BadgeV2.Contract.TransferOwnership(&_BadgeV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BadgeV2 *BadgeV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BadgeV2.Contract.TransferOwnership(&_BadgeV2.TransactOpts, newOwner)
}

// UpdateQuest is a paid mutator transaction binding the contract method 0x843ef06f.
//
// Solidity: function updateQuest(uint256 questId, uint32 startTs, uint32 endTs, string title, string questUri) returns()
func (_BadgeV2 *BadgeV2Transactor) UpdateQuest(opts *bind.TransactOpts, questId *big.Int, startTs uint32, endTs uint32, title string, questUri string) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "updateQuest", questId, startTs, endTs, title, questUri)
}

// UpdateQuest is a paid mutator transaction binding the contract method 0x843ef06f.
//
// Solidity: function updateQuest(uint256 questId, uint32 startTs, uint32 endTs, string title, string questUri) returns()
func (_BadgeV2 *BadgeV2Session) UpdateQuest(questId *big.Int, startTs uint32, endTs uint32, title string, questUri string) (*types.Transaction, error) {
	return _BadgeV2.Contract.UpdateQuest(&_BadgeV2.TransactOpts, questId, startTs, endTs, title, questUri)
}

// UpdateQuest is a paid mutator transaction binding the contract method 0x843ef06f.
//
// Solidity: function updateQuest(uint256 questId, uint32 startTs, uint32 endTs, string title, string questUri) returns()
func (_BadgeV2 *BadgeV2TransactorSession) UpdateQuest(questId *big.Int, startTs uint32, endTs uint32, title string, questUri string) (*types.Transaction, error) {
	return _BadgeV2.Contract.UpdateQuest(&_BadgeV2.TransactOpts, questId, startTs, endTs, title, questUri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x31d41c69.
//
// Solidity: function updateURI(uint256 tokenId, string uri) returns()
func (_BadgeV2 *BadgeV2Transactor) UpdateURI(opts *bind.TransactOpts, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _BadgeV2.contract.Transact(opts, "updateURI", tokenId, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x31d41c69.
//
// Solidity: function updateURI(uint256 tokenId, string uri) returns()
func (_BadgeV2 *BadgeV2Session) UpdateURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _BadgeV2.Contract.UpdateURI(&_BadgeV2.TransactOpts, tokenId, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x31d41c69.
//
// Solidity: function updateURI(uint256 tokenId, string uri) returns()
func (_BadgeV2 *BadgeV2TransactorSession) UpdateURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _BadgeV2.Contract.UpdateURI(&_BadgeV2.TransactOpts, tokenId, uri)
}

// BadgeV2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BadgeV2 contract.
type BadgeV2ApprovalIterator struct {
	Event *BadgeV2Approval // Event containing the contract specifics and raw log

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
func (it *BadgeV2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2Approval)
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
		it.Event = new(BadgeV2Approval)
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
func (it *BadgeV2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2Approval represents a Approval event raised by the BadgeV2 contract.
type BadgeV2Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BadgeV2 *BadgeV2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BadgeV2ApprovalIterator, error) {

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

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeV2ApprovalIterator{contract: _BadgeV2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BadgeV2 *BadgeV2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BadgeV2Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2Approval)
				if err := _BadgeV2.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BadgeV2 *BadgeV2Filterer) ParseApproval(log types.Log) (*BadgeV2Approval, error) {
	event := new(BadgeV2Approval)
	if err := _BadgeV2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the BadgeV2 contract.
type BadgeV2ApprovalForAllIterator struct {
	Event *BadgeV2ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BadgeV2ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2ApprovalForAll)
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
		it.Event = new(BadgeV2ApprovalForAll)
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
func (it *BadgeV2ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2ApprovalForAll represents a ApprovalForAll event raised by the BadgeV2 contract.
type BadgeV2ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BadgeV2 *BadgeV2Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BadgeV2ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BadgeV2ApprovalForAllIterator{contract: _BadgeV2.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BadgeV2 *BadgeV2Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BadgeV2ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2ApprovalForAll)
				if err := _BadgeV2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_BadgeV2 *BadgeV2Filterer) ParseApprovalForAll(log types.Log) (*BadgeV2ApprovalForAll, error) {
	event := new(BadgeV2ApprovalForAll)
	if err := _BadgeV2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2ClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the BadgeV2 contract.
type BadgeV2ClaimedIterator struct {
	Event *BadgeV2Claimed // Event containing the contract specifics and raw log

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
func (it *BadgeV2ClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2Claimed)
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
		it.Event = new(BadgeV2Claimed)
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
func (it *BadgeV2ClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2ClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2Claimed represents a Claimed event raised by the BadgeV2 contract.
type BadgeV2Claimed struct {
	TokenId *big.Int
	QuestId *big.Int
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xc09f7b3c1c5b70bf6d30e485a4525d625d4723aed5d319376825446ef135c7ca.
//
// Solidity: event Claimed(uint256 indexed tokenId, uint256 questId, address sender)
func (_BadgeV2 *BadgeV2Filterer) FilterClaimed(opts *bind.FilterOpts, tokenId []*big.Int) (*BadgeV2ClaimedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "Claimed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeV2ClaimedIterator{contract: _BadgeV2.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xc09f7b3c1c5b70bf6d30e485a4525d625d4723aed5d319376825446ef135c7ca.
//
// Solidity: event Claimed(uint256 indexed tokenId, uint256 questId, address sender)
func (_BadgeV2 *BadgeV2Filterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *BadgeV2Claimed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "Claimed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2Claimed)
				if err := _BadgeV2.contract.UnpackLog(event, "Claimed", log); err != nil {
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
// Solidity: event Claimed(uint256 indexed tokenId, uint256 questId, address sender)
func (_BadgeV2 *BadgeV2Filterer) ParseClaimed(log types.Log) (*BadgeV2Claimed, error) {
	event := new(BadgeV2Claimed)
	if err := _BadgeV2.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2DonationIterator is returned from FilterDonation and is used to iterate over the raw logs and unpacked data for Donation events raised by the BadgeV2 contract.
type BadgeV2DonationIterator struct {
	Event *BadgeV2Donation // Event containing the contract specifics and raw log

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
func (it *BadgeV2DonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2Donation)
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
		it.Event = new(BadgeV2Donation)
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
func (it *BadgeV2DonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2DonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2Donation represents a Donation event raised by the BadgeV2 contract.
type BadgeV2Donation struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonation is a free log retrieval operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_BadgeV2 *BadgeV2Filterer) FilterDonation(opts *bind.FilterOpts) (*BadgeV2DonationIterator, error) {

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return &BadgeV2DonationIterator{contract: _BadgeV2.contract, event: "Donation", logs: logs, sub: sub}, nil
}

// WatchDonation is a free log subscription operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_BadgeV2 *BadgeV2Filterer) WatchDonation(opts *bind.WatchOpts, sink chan<- *BadgeV2Donation) (event.Subscription, error) {

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2Donation)
				if err := _BadgeV2.contract.UnpackLog(event, "Donation", log); err != nil {
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
func (_BadgeV2 *BadgeV2Filterer) ParseDonation(log types.Log) (*BadgeV2Donation, error) {
	event := new(BadgeV2Donation)
	if err := _BadgeV2.contract.UnpackLog(event, "Donation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2LockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the BadgeV2 contract.
type BadgeV2LockedIterator struct {
	Event *BadgeV2Locked // Event containing the contract specifics and raw log

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
func (it *BadgeV2LockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2Locked)
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
		it.Event = new(BadgeV2Locked)
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
func (it *BadgeV2LockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2LockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2Locked represents a Locked event raised by the BadgeV2 contract.
type BadgeV2Locked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_BadgeV2 *BadgeV2Filterer) FilterLocked(opts *bind.FilterOpts) (*BadgeV2LockedIterator, error) {

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return &BadgeV2LockedIterator{contract: _BadgeV2.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_BadgeV2 *BadgeV2Filterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *BadgeV2Locked) (event.Subscription, error) {

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2Locked)
				if err := _BadgeV2.contract.UnpackLog(event, "Locked", log); err != nil {
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
func (_BadgeV2 *BadgeV2Filterer) ParseLocked(log types.Log) (*BadgeV2Locked, error) {
	event := new(BadgeV2Locked)
	if err := _BadgeV2.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2MinterSetIterator is returned from FilterMinterSet and is used to iterate over the raw logs and unpacked data for MinterSet events raised by the BadgeV2 contract.
type BadgeV2MinterSetIterator struct {
	Event *BadgeV2MinterSet // Event containing the contract specifics and raw log

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
func (it *BadgeV2MinterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2MinterSet)
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
		it.Event = new(BadgeV2MinterSet)
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
func (it *BadgeV2MinterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2MinterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2MinterSet represents a MinterSet event raised by the BadgeV2 contract.
type BadgeV2MinterSet struct {
	Minter  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterSet is a free log retrieval operation binding the contract event 0x583b0aa0e528532caf4b907c11d7a8158a122fe2a6fb80cd9b09776ebea8d92d.
//
// Solidity: event MinterSet(address minter, bool enabled)
func (_BadgeV2 *BadgeV2Filterer) FilterMinterSet(opts *bind.FilterOpts) (*BadgeV2MinterSetIterator, error) {

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "MinterSet")
	if err != nil {
		return nil, err
	}
	return &BadgeV2MinterSetIterator{contract: _BadgeV2.contract, event: "MinterSet", logs: logs, sub: sub}, nil
}

// WatchMinterSet is a free log subscription operation binding the contract event 0x583b0aa0e528532caf4b907c11d7a8158a122fe2a6fb80cd9b09776ebea8d92d.
//
// Solidity: event MinterSet(address minter, bool enabled)
func (_BadgeV2 *BadgeV2Filterer) WatchMinterSet(opts *bind.WatchOpts, sink chan<- *BadgeV2MinterSet) (event.Subscription, error) {

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "MinterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2MinterSet)
				if err := _BadgeV2.contract.UnpackLog(event, "MinterSet", log); err != nil {
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
func (_BadgeV2 *BadgeV2Filterer) ParseMinterSet(log types.Log) (*BadgeV2MinterSet, error) {
	event := new(BadgeV2MinterSet)
	if err := _BadgeV2.contract.UnpackLog(event, "MinterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BadgeV2 contract.
type BadgeV2OwnershipTransferredIterator struct {
	Event *BadgeV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BadgeV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2OwnershipTransferred)
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
		it.Event = new(BadgeV2OwnershipTransferred)
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
func (it *BadgeV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2OwnershipTransferred represents a OwnershipTransferred event raised by the BadgeV2 contract.
type BadgeV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BadgeV2 *BadgeV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BadgeV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BadgeV2OwnershipTransferredIterator{contract: _BadgeV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BadgeV2 *BadgeV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BadgeV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2OwnershipTransferred)
				if err := _BadgeV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BadgeV2 *BadgeV2Filterer) ParseOwnershipTransferred(log types.Log) (*BadgeV2OwnershipTransferred, error) {
	event := new(BadgeV2OwnershipTransferred)
	if err := _BadgeV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2QuestInitIterator is returned from FilterQuestInit and is used to iterate over the raw logs and unpacked data for QuestInit events raised by the BadgeV2 contract.
type BadgeV2QuestInitIterator struct {
	Event *BadgeV2QuestInit // Event containing the contract specifics and raw log

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
func (it *BadgeV2QuestInitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2QuestInit)
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
		it.Event = new(BadgeV2QuestInit)
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
func (it *BadgeV2QuestInitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2QuestInitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2QuestInit represents a QuestInit event raised by the BadgeV2 contract.
type BadgeV2QuestInit struct {
	QuestId   *big.Int
	QuestData IBadgeQuestData
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQuestInit is a free log retrieval operation binding the contract event 0x850df14eacc76f9e3cf422c9f7399a39336693bf27b7def779fb7a72ff67b03a.
//
// Solidity: event QuestInit(uint256 indexed questId, (address,uint32,uint32,string,string) questData)
func (_BadgeV2 *BadgeV2Filterer) FilterQuestInit(opts *bind.FilterOpts, questId []*big.Int) (*BadgeV2QuestInitIterator, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "QuestInit", questIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeV2QuestInitIterator{contract: _BadgeV2.contract, event: "QuestInit", logs: logs, sub: sub}, nil
}

// WatchQuestInit is a free log subscription operation binding the contract event 0x850df14eacc76f9e3cf422c9f7399a39336693bf27b7def779fb7a72ff67b03a.
//
// Solidity: event QuestInit(uint256 indexed questId, (address,uint32,uint32,string,string) questData)
func (_BadgeV2 *BadgeV2Filterer) WatchQuestInit(opts *bind.WatchOpts, sink chan<- *BadgeV2QuestInit, questId []*big.Int) (event.Subscription, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "QuestInit", questIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2QuestInit)
				if err := _BadgeV2.contract.UnpackLog(event, "QuestInit", log); err != nil {
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

// ParseQuestInit is a log parse operation binding the contract event 0x850df14eacc76f9e3cf422c9f7399a39336693bf27b7def779fb7a72ff67b03a.
//
// Solidity: event QuestInit(uint256 indexed questId, (address,uint32,uint32,string,string) questData)
func (_BadgeV2 *BadgeV2Filterer) ParseQuestInit(log types.Log) (*BadgeV2QuestInit, error) {
	event := new(BadgeV2QuestInit)
	if err := _BadgeV2.contract.UnpackLog(event, "QuestInit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2QuestUpdatedIterator is returned from FilterQuestUpdated and is used to iterate over the raw logs and unpacked data for QuestUpdated events raised by the BadgeV2 contract.
type BadgeV2QuestUpdatedIterator struct {
	Event *BadgeV2QuestUpdated // Event containing the contract specifics and raw log

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
func (it *BadgeV2QuestUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2QuestUpdated)
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
		it.Event = new(BadgeV2QuestUpdated)
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
func (it *BadgeV2QuestUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2QuestUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2QuestUpdated represents a QuestUpdated event raised by the BadgeV2 contract.
type BadgeV2QuestUpdated struct {
	QuestId   *big.Int
	QuestData IBadgeQuestData
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQuestUpdated is a free log retrieval operation binding the contract event 0x0e710ba4fe722bed44e55cf8a7fbf3a59f2673584b8551e7356f2b11700ac3e7.
//
// Solidity: event QuestUpdated(uint256 indexed questId, (address,uint32,uint32,string,string) questData)
func (_BadgeV2 *BadgeV2Filterer) FilterQuestUpdated(opts *bind.FilterOpts, questId []*big.Int) (*BadgeV2QuestUpdatedIterator, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "QuestUpdated", questIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeV2QuestUpdatedIterator{contract: _BadgeV2.contract, event: "QuestUpdated", logs: logs, sub: sub}, nil
}

// WatchQuestUpdated is a free log subscription operation binding the contract event 0x0e710ba4fe722bed44e55cf8a7fbf3a59f2673584b8551e7356f2b11700ac3e7.
//
// Solidity: event QuestUpdated(uint256 indexed questId, (address,uint32,uint32,string,string) questData)
func (_BadgeV2 *BadgeV2Filterer) WatchQuestUpdated(opts *bind.WatchOpts, sink chan<- *BadgeV2QuestUpdated, questId []*big.Int) (event.Subscription, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "QuestUpdated", questIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2QuestUpdated)
				if err := _BadgeV2.contract.UnpackLog(event, "QuestUpdated", log); err != nil {
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

// ParseQuestUpdated is a log parse operation binding the contract event 0x0e710ba4fe722bed44e55cf8a7fbf3a59f2673584b8551e7356f2b11700ac3e7.
//
// Solidity: event QuestUpdated(uint256 indexed questId, (address,uint32,uint32,string,string) questData)
func (_BadgeV2 *BadgeV2Filterer) ParseQuestUpdated(log types.Log) (*BadgeV2QuestUpdated, error) {
	event := new(BadgeV2QuestUpdated)
	if err := _BadgeV2.contract.UnpackLog(event, "QuestUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BadgeV2 contract.
type BadgeV2TransferIterator struct {
	Event *BadgeV2Transfer // Event containing the contract specifics and raw log

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
func (it *BadgeV2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2Transfer)
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
		it.Event = new(BadgeV2Transfer)
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
func (it *BadgeV2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2Transfer represents a Transfer event raised by the BadgeV2 contract.
type BadgeV2Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BadgeV2 *BadgeV2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BadgeV2TransferIterator, error) {

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

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeV2TransferIterator{contract: _BadgeV2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BadgeV2 *BadgeV2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BadgeV2Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2Transfer)
				if err := _BadgeV2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BadgeV2 *BadgeV2Filterer) ParseTransfer(log types.Log) (*BadgeV2Transfer, error) {
	event := new(BadgeV2Transfer)
	if err := _BadgeV2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2URIUpdatedIterator is returned from FilterURIUpdated and is used to iterate over the raw logs and unpacked data for URIUpdated events raised by the BadgeV2 contract.
type BadgeV2URIUpdatedIterator struct {
	Event *BadgeV2URIUpdated // Event containing the contract specifics and raw log

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
func (it *BadgeV2URIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2URIUpdated)
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
		it.Event = new(BadgeV2URIUpdated)
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
func (it *BadgeV2URIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2URIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2URIUpdated represents a URIUpdated event raised by the BadgeV2 contract.
type BadgeV2URIUpdated struct {
	TokenId *big.Int
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterURIUpdated is a free log retrieval operation binding the contract event 0xbf032ae7dcb484f99f9f3977ba9cde222756e7073de74244ceb89de862569bdb.
//
// Solidity: event URIUpdated(uint256 indexed tokenId, string uri)
func (_BadgeV2 *BadgeV2Filterer) FilterURIUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*BadgeV2URIUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "URIUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BadgeV2URIUpdatedIterator{contract: _BadgeV2.contract, event: "URIUpdated", logs: logs, sub: sub}, nil
}

// WatchURIUpdated is a free log subscription operation binding the contract event 0xbf032ae7dcb484f99f9f3977ba9cde222756e7073de74244ceb89de862569bdb.
//
// Solidity: event URIUpdated(uint256 indexed tokenId, string uri)
func (_BadgeV2 *BadgeV2Filterer) WatchURIUpdated(opts *bind.WatchOpts, sink chan<- *BadgeV2URIUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "URIUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2URIUpdated)
				if err := _BadgeV2.contract.UnpackLog(event, "URIUpdated", log); err != nil {
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
func (_BadgeV2 *BadgeV2Filterer) ParseURIUpdated(log types.Log) (*BadgeV2URIUpdated, error) {
	event := new(BadgeV2URIUpdated)
	if err := _BadgeV2.contract.UnpackLog(event, "URIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeV2UnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the BadgeV2 contract.
type BadgeV2UnlockedIterator struct {
	Event *BadgeV2Unlocked // Event containing the contract specifics and raw log

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
func (it *BadgeV2UnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeV2Unlocked)
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
		it.Event = new(BadgeV2Unlocked)
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
func (it *BadgeV2UnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeV2UnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeV2Unlocked represents a Unlocked event raised by the BadgeV2 contract.
type BadgeV2Unlocked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 tokenId)
func (_BadgeV2 *BadgeV2Filterer) FilterUnlocked(opts *bind.FilterOpts) (*BadgeV2UnlockedIterator, error) {

	logs, sub, err := _BadgeV2.contract.FilterLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return &BadgeV2UnlockedIterator{contract: _BadgeV2.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 tokenId)
func (_BadgeV2 *BadgeV2Filterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *BadgeV2Unlocked) (event.Subscription, error) {

	logs, sub, err := _BadgeV2.contract.WatchLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeV2Unlocked)
				if err := _BadgeV2.contract.UnpackLog(event, "Unlocked", log); err != nil {
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
func (_BadgeV2 *BadgeV2Filterer) ParseUnlocked(log types.Log) (*BadgeV2Unlocked, error) {
	event := new(BadgeV2Unlocked)
	if err := _BadgeV2.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
