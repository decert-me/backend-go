// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// IQuestQuestData is an auto generated low-level Go binding around an user-defined struct.
type IQuestQuestData struct {
	StartTs uint32
	EndTs   uint32
	Supply  *big.Int
	Title   string
	Uri     string
}

// QuestMetaData contains all meta data concerning the Quest contract.
var QuestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badge_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ClaimedCannotModify\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"supply\",\"type\":\"uint192\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"}],\"name\":\"QuestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SetMinter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"badge\",\"outputs\":[{\"internalType\":\"contractIBadge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getQuest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"supply\",\"type\":\"uint192\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"meta\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"supply\",\"type\":\"uint192\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"supply\",\"type\":\"uint192\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"}],\"name\":\"modifyQuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"quests\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"supply\",\"type\":\"uint192\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_meta\",\"type\":\"address\"}],\"name\":\"setMetaContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QuestABI is the input ABI used to generate the binding from.
// Deprecated: Use QuestMetaData.ABI instead.
var QuestABI = QuestMetaData.ABI

// Quest is an auto generated Go binding around an Ethereum contract.
type Quest struct {
	QuestCaller     // Read-only binding to the contract
	QuestTransactor // Write-only binding to the contract
	QuestFilterer   // Log filterer for contract events
}

// QuestCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuestSession struct {
	Contract     *Quest            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuestCallerSession struct {
	Contract *QuestCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuestTransactorSession struct {
	Contract     *QuestTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuestRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuestRaw struct {
	Contract *Quest // Generic contract binding to access the raw methods on
}

// QuestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuestCallerRaw struct {
	Contract *QuestCaller // Generic read-only contract binding to access the raw methods on
}

// QuestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuestTransactorRaw struct {
	Contract *QuestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuest creates a new instance of Quest, bound to a specific deployed contract.
func NewQuest(address common.Address, backend bind.ContractBackend) (*Quest, error) {
	contract, err := bindQuest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quest{QuestCaller: QuestCaller{contract: contract}, QuestTransactor: QuestTransactor{contract: contract}, QuestFilterer: QuestFilterer{contract: contract}}, nil
}

// NewQuestCaller creates a new read-only instance of Quest, bound to a specific deployed contract.
func NewQuestCaller(address common.Address, caller bind.ContractCaller) (*QuestCaller, error) {
	contract, err := bindQuest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuestCaller{contract: contract}, nil
}

// NewQuestTransactor creates a new write-only instance of Quest, bound to a specific deployed contract.
func NewQuestTransactor(address common.Address, transactor bind.ContractTransactor) (*QuestTransactor, error) {
	contract, err := bindQuest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuestTransactor{contract: contract}, nil
}

// NewQuestFilterer creates a new log filterer instance of Quest, bound to a specific deployed contract.
func NewQuestFilterer(address common.Address, filterer bind.ContractFilterer) (*QuestFilterer, error) {
	contract, err := bindQuest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuestFilterer{contract: contract}, nil
}

// bindQuest binds a generic wrapper to an already deployed contract.
func bindQuest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quest *QuestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quest.Contract.QuestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quest *QuestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quest.Contract.QuestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quest *QuestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quest.Contract.QuestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quest *QuestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quest *QuestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quest *QuestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quest.Contract.contract.Transact(opts, method, params...)
}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_Quest *QuestCaller) Badge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "badge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_Quest *QuestSession) Badge() (common.Address, error) {
	return _Quest.Contract.Badge(&_Quest.CallOpts)
}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_Quest *QuestCallerSession) Badge() (common.Address, error) {
	return _Quest.Contract.Badge(&_Quest.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Quest *QuestCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Quest *QuestSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Quest.Contract.BalanceOf(&_Quest.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Quest *QuestCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Quest.Contract.BalanceOf(&_Quest.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Quest *QuestCaller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Quest *QuestSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Quest.Contract.GetApproved(&_Quest.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Quest *QuestCallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Quest.Contract.GetApproved(&_Quest.CallOpts, arg0)
}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 tokenId) view returns((uint32,uint32,uint192,string,string) questData)
func (_Quest *QuestCaller) GetQuest(opts *bind.CallOpts, tokenId *big.Int) (IQuestQuestData, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "getQuest", tokenId)

	if err != nil {
		return *new(IQuestQuestData), err
	}

	out0 := *abi.ConvertType(out[0], new(IQuestQuestData)).(*IQuestQuestData)

	return out0, err

}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 tokenId) view returns((uint32,uint32,uint192,string,string) questData)
func (_Quest *QuestSession) GetQuest(tokenId *big.Int) (IQuestQuestData, error) {
	return _Quest.Contract.GetQuest(&_Quest.CallOpts, tokenId)
}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 tokenId) view returns((uint32,uint32,uint192,string,string) questData)
func (_Quest *QuestCallerSession) GetQuest(tokenId *big.Int) (IQuestQuestData, error) {
	return _Quest.Contract.GetQuest(&_Quest.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Quest *QuestCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Quest *QuestSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Quest.Contract.IsApprovedForAll(&_Quest.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Quest *QuestCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Quest.Contract.IsApprovedForAll(&_Quest.CallOpts, arg0, arg1)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_Quest *QuestCaller) Locked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "locked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_Quest *QuestSession) Locked(tokenId *big.Int) (bool, error) {
	return _Quest.Contract.Locked(&_Quest.CallOpts, tokenId)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_Quest *QuestCallerSession) Locked(tokenId *big.Int) (bool, error) {
	return _Quest.Contract.Locked(&_Quest.CallOpts, tokenId)
}

// Meta is a free data retrieval call binding the contract method 0xc885044e.
//
// Solidity: function meta() view returns(address)
func (_Quest *QuestCaller) Meta(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "meta")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Meta is a free data retrieval call binding the contract method 0xc885044e.
//
// Solidity: function meta() view returns(address)
func (_Quest *QuestSession) Meta() (common.Address, error) {
	return _Quest.Contract.Meta(&_Quest.CallOpts)
}

// Meta is a free data retrieval call binding the contract method 0xc885044e.
//
// Solidity: function meta() view returns(address)
func (_Quest *QuestCallerSession) Meta() (common.Address, error) {
	return _Quest.Contract.Meta(&_Quest.CallOpts)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_Quest *QuestCaller) Minters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_Quest *QuestSession) Minters(arg0 common.Address) (bool, error) {
	return _Quest.Contract.Minters(&_Quest.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_Quest *QuestCallerSession) Minters(arg0 common.Address) (bool, error) {
	return _Quest.Contract.Minters(&_Quest.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Quest *QuestCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Quest *QuestSession) Name() (string, error) {
	return _Quest.Contract.Name(&_Quest.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Quest *QuestCallerSession) Name() (string, error) {
	return _Quest.Contract.Name(&_Quest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Quest *QuestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Quest *QuestSession) Owner() (common.Address, error) {
	return _Quest.Contract.Owner(&_Quest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Quest *QuestCallerSession) Owner() (common.Address, error) {
	return _Quest.Contract.Owner(&_Quest.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Quest *QuestCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Quest *QuestSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Quest.Contract.OwnerOf(&_Quest.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Quest *QuestCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Quest.Contract.OwnerOf(&_Quest.CallOpts, tokenId)
}

// Quests is a free data retrieval call binding the contract method 0xe085f980.
//
// Solidity: function quests(uint256 ) view returns(uint32 startTs, uint32 endTs, uint192 supply, string title, string uri)
func (_Quest *QuestCaller) Quests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartTs uint32
	EndTs   uint32
	Supply  *big.Int
	Title   string
	Uri     string
}, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "quests", arg0)

	outstruct := new(struct {
		StartTs uint32
		EndTs   uint32
		Supply  *big.Int
		Title   string
		Uri     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTs = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.EndTs = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Supply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Title = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Uri = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// Quests is a free data retrieval call binding the contract method 0xe085f980.
//
// Solidity: function quests(uint256 ) view returns(uint32 startTs, uint32 endTs, uint192 supply, string title, string uri)
func (_Quest *QuestSession) Quests(arg0 *big.Int) (struct {
	StartTs uint32
	EndTs   uint32
	Supply  *big.Int
	Title   string
	Uri     string
}, error) {
	return _Quest.Contract.Quests(&_Quest.CallOpts, arg0)
}

// Quests is a free data retrieval call binding the contract method 0xe085f980.
//
// Solidity: function quests(uint256 ) view returns(uint32 startTs, uint32 endTs, uint192 supply, string title, string uri)
func (_Quest *QuestCallerSession) Quests(arg0 *big.Int) (struct {
	StartTs uint32
	EndTs   uint32
	Supply  *big.Int
	Title   string
	Uri     string
}, error) {
	return _Quest.Contract.Quests(&_Quest.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Quest *QuestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Quest *QuestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Quest.Contract.SupportsInterface(&_Quest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Quest *QuestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Quest.Contract.SupportsInterface(&_Quest.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Quest *QuestCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Quest *QuestSession) Symbol() (string, error) {
	return _Quest.Contract.Symbol(&_Quest.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Quest *QuestCallerSession) Symbol() (string, error) {
	return _Quest.Contract.Symbol(&_Quest.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Quest *QuestCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Quest *QuestSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Quest.Contract.TokenURI(&_Quest.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Quest *QuestCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Quest.Contract.TokenURI(&_Quest.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Quest *QuestCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Quest.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Quest *QuestSession) TotalSupply() (*big.Int, error) {
	return _Quest.Contract.TotalSupply(&_Quest.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Quest *QuestCallerSession) TotalSupply() (*big.Int, error) {
	return _Quest.Contract.TotalSupply(&_Quest.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_Quest *QuestTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_Quest *QuestSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Quest.Contract.Approve(&_Quest.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_Quest *QuestTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Quest.Contract.Approve(&_Quest.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x8ddf7d3b.
//
// Solidity: function mint(address to, uint256 tokenId, (uint32,uint32,uint192,string,string) questData, bytes data) returns()
func (_Quest *QuestTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, questData IQuestQuestData, data []byte) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "mint", to, tokenId, questData, data)
}

// Mint is a paid mutator transaction binding the contract method 0x8ddf7d3b.
//
// Solidity: function mint(address to, uint256 tokenId, (uint32,uint32,uint192,string,string) questData, bytes data) returns()
func (_Quest *QuestSession) Mint(to common.Address, tokenId *big.Int, questData IQuestQuestData, data []byte) (*types.Transaction, error) {
	return _Quest.Contract.Mint(&_Quest.TransactOpts, to, tokenId, questData, data)
}

// Mint is a paid mutator transaction binding the contract method 0x8ddf7d3b.
//
// Solidity: function mint(address to, uint256 tokenId, (uint32,uint32,uint192,string,string) questData, bytes data) returns()
func (_Quest *QuestTransactorSession) Mint(to common.Address, tokenId *big.Int, questData IQuestQuestData, data []byte) (*types.Transaction, error) {
	return _Quest.Contract.Mint(&_Quest.TransactOpts, to, tokenId, questData, data)
}

// ModifyQuest is a paid mutator transaction binding the contract method 0x7cee9c6c.
//
// Solidity: function modifyQuest(uint256 tokenId, (uint32,uint32,uint192,string,string) questData) returns()
func (_Quest *QuestTransactor) ModifyQuest(opts *bind.TransactOpts, tokenId *big.Int, questData IQuestQuestData) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "modifyQuest", tokenId, questData)
}

// ModifyQuest is a paid mutator transaction binding the contract method 0x7cee9c6c.
//
// Solidity: function modifyQuest(uint256 tokenId, (uint32,uint32,uint192,string,string) questData) returns()
func (_Quest *QuestSession) ModifyQuest(tokenId *big.Int, questData IQuestQuestData) (*types.Transaction, error) {
	return _Quest.Contract.ModifyQuest(&_Quest.TransactOpts, tokenId, questData)
}

// ModifyQuest is a paid mutator transaction binding the contract method 0x7cee9c6c.
//
// Solidity: function modifyQuest(uint256 tokenId, (uint32,uint32,uint192,string,string) questData) returns()
func (_Quest *QuestTransactorSession) ModifyQuest(tokenId *big.Int, questData IQuestQuestData) (*types.Transaction, error) {
	return _Quest.Contract.ModifyQuest(&_Quest.TransactOpts, tokenId, questData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Quest *QuestTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Quest *QuestSession) RenounceOwnership() (*types.Transaction, error) {
	return _Quest.Contract.RenounceOwnership(&_Quest.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Quest *QuestTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Quest.Contract.RenounceOwnership(&_Quest.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_Quest *QuestTransactor) SafeTransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "safeTransferFrom", arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_Quest *QuestSession) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Quest.Contract.SafeTransferFrom(&_Quest.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_Quest *QuestTransactorSession) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Quest.Contract.SafeTransferFrom(&_Quest.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_Quest *QuestTransactor) SafeTransferFrom0(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "safeTransferFrom0", arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_Quest *QuestSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Quest.Contract.SafeTransferFrom0(&_Quest.TransactOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_Quest *QuestTransactorSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Quest.Contract.SafeTransferFrom0(&_Quest.TransactOpts, arg0, arg1, arg2, arg3)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_Quest *QuestTransactor) SetApprovalForAll(opts *bind.TransactOpts, arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "setApprovalForAll", arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_Quest *QuestSession) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _Quest.Contract.SetApprovalForAll(&_Quest.TransactOpts, arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_Quest *QuestTransactorSession) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _Quest.Contract.SetApprovalForAll(&_Quest.TransactOpts, arg0, arg1)
}

// SetMetaContract is a paid mutator transaction binding the contract method 0x0957a4b6.
//
// Solidity: function setMetaContract(address _meta) returns()
func (_Quest *QuestTransactor) SetMetaContract(opts *bind.TransactOpts, _meta common.Address) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "setMetaContract", _meta)
}

// SetMetaContract is a paid mutator transaction binding the contract method 0x0957a4b6.
//
// Solidity: function setMetaContract(address _meta) returns()
func (_Quest *QuestSession) SetMetaContract(_meta common.Address) (*types.Transaction, error) {
	return _Quest.Contract.SetMetaContract(&_Quest.TransactOpts, _meta)
}

// SetMetaContract is a paid mutator transaction binding the contract method 0x0957a4b6.
//
// Solidity: function setMetaContract(address _meta) returns()
func (_Quest *QuestTransactorSession) SetMetaContract(_meta common.Address) (*types.Transaction, error) {
	return _Quest.Contract.SetMetaContract(&_Quest.TransactOpts, _meta)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_Quest *QuestTransactor) SetMinter(opts *bind.TransactOpts, minter common.Address, enabled bool) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "setMinter", minter, enabled)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_Quest *QuestSession) SetMinter(minter common.Address, enabled bool) (*types.Transaction, error) {
	return _Quest.Contract.SetMinter(&_Quest.TransactOpts, minter, enabled)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_Quest *QuestTransactorSession) SetMinter(minter common.Address, enabled bool) (*types.Transaction, error) {
	return _Quest.Contract.SetMinter(&_Quest.TransactOpts, minter, enabled)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_Quest *QuestTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_Quest *QuestSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Quest.Contract.TransferFrom(&_Quest.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_Quest *QuestTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Quest.Contract.TransferFrom(&_Quest.TransactOpts, arg0, arg1, arg2)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Quest *QuestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Quest *QuestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Quest.Contract.TransferOwnership(&_Quest.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Quest *QuestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Quest.Contract.TransferOwnership(&_Quest.TransactOpts, newOwner)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x31d41c69.
//
// Solidity: function updateURI(uint256 tokenId, string uri) returns()
func (_Quest *QuestTransactor) UpdateURI(opts *bind.TransactOpts, tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Quest.contract.Transact(opts, "updateURI", tokenId, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x31d41c69.
//
// Solidity: function updateURI(uint256 tokenId, string uri) returns()
func (_Quest *QuestSession) UpdateURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Quest.Contract.UpdateURI(&_Quest.TransactOpts, tokenId, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x31d41c69.
//
// Solidity: function updateURI(uint256 tokenId, string uri) returns()
func (_Quest *QuestTransactorSession) UpdateURI(tokenId *big.Int, uri string) (*types.Transaction, error) {
	return _Quest.Contract.UpdateURI(&_Quest.TransactOpts, tokenId, uri)
}

// QuestApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Quest contract.
type QuestApprovalIterator struct {
	Event *QuestApproval // Event containing the contract specifics and raw log

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
func (it *QuestApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestApproval)
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
		it.Event = new(QuestApproval)
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
func (it *QuestApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestApproval represents a Approval event raised by the Quest contract.
type QuestApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Quest *QuestFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*QuestApprovalIterator, error) {

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

	logs, sub, err := _Quest.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuestApprovalIterator{contract: _Quest.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Quest *QuestFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *QuestApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Quest.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestApproval)
				if err := _Quest.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Quest *QuestFilterer) ParseApproval(log types.Log) (*QuestApproval, error) {
	event := new(QuestApproval)
	if err := _Quest.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Quest contract.
type QuestApprovalForAllIterator struct {
	Event *QuestApprovalForAll // Event containing the contract specifics and raw log

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
func (it *QuestApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestApprovalForAll)
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
		it.Event = new(QuestApprovalForAll)
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
func (it *QuestApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestApprovalForAll represents a ApprovalForAll event raised by the Quest contract.
type QuestApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Quest *QuestFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*QuestApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Quest.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &QuestApprovalForAllIterator{contract: _Quest.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Quest *QuestFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *QuestApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Quest.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestApprovalForAll)
				if err := _Quest.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Quest *QuestFilterer) ParseApprovalForAll(log types.Log) (*QuestApprovalForAll, error) {
	event := new(QuestApprovalForAll)
	if err := _Quest.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the Quest contract.
type QuestLockedIterator struct {
	Event *QuestLocked // Event containing the contract specifics and raw log

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
func (it *QuestLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestLocked)
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
		it.Event = new(QuestLocked)
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
func (it *QuestLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestLocked represents a Locked event raised by the Quest contract.
type QuestLocked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_Quest *QuestFilterer) FilterLocked(opts *bind.FilterOpts) (*QuestLockedIterator, error) {

	logs, sub, err := _Quest.contract.FilterLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return &QuestLockedIterator{contract: _Quest.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_Quest *QuestFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *QuestLocked) (event.Subscription, error) {

	logs, sub, err := _Quest.contract.WatchLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestLocked)
				if err := _Quest.contract.UnpackLog(event, "Locked", log); err != nil {
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
func (_Quest *QuestFilterer) ParseLocked(log types.Log) (*QuestLocked, error) {
	event := new(QuestLocked)
	if err := _Quest.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Quest contract.
type QuestOwnershipTransferredIterator struct {
	Event *QuestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QuestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestOwnershipTransferred)
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
		it.Event = new(QuestOwnershipTransferred)
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
func (it *QuestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestOwnershipTransferred represents a OwnershipTransferred event raised by the Quest contract.
type QuestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Quest *QuestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QuestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Quest.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuestOwnershipTransferredIterator{contract: _Quest.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Quest *QuestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QuestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Quest.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestOwnershipTransferred)
				if err := _Quest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Quest *QuestFilterer) ParseOwnershipTransferred(log types.Log) (*QuestOwnershipTransferred, error) {
	event := new(QuestOwnershipTransferred)
	if err := _Quest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestQuestCreatedIterator is returned from FilterQuestCreated and is used to iterate over the raw logs and unpacked data for QuestCreated events raised by the Quest contract.
type QuestQuestCreatedIterator struct {
	Event *QuestQuestCreated // Event containing the contract specifics and raw log

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
func (it *QuestQuestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestQuestCreated)
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
		it.Event = new(QuestQuestCreated)
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
func (it *QuestQuestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestQuestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestQuestCreated represents a QuestCreated event raised by the Quest contract.
type QuestQuestCreated struct {
	Creator   common.Address
	TokenId   *big.Int
	QuestData IQuestQuestData
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQuestCreated is a free log retrieval operation binding the contract event 0x7dcb53be1d5e10c26d363e67d5c06d05da75f184de604abd3f9e41efdb7ee576.
//
// Solidity: event QuestCreated(address indexed creator, uint256 indexed tokenId, (uint32,uint32,uint192,string,string) questData)
func (_Quest *QuestFilterer) FilterQuestCreated(opts *bind.FilterOpts, creator []common.Address, tokenId []*big.Int) (*QuestQuestCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Quest.contract.FilterLogs(opts, "QuestCreated", creatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuestQuestCreatedIterator{contract: _Quest.contract, event: "QuestCreated", logs: logs, sub: sub}, nil
}

// WatchQuestCreated is a free log subscription operation binding the contract event 0x7dcb53be1d5e10c26d363e67d5c06d05da75f184de604abd3f9e41efdb7ee576.
//
// Solidity: event QuestCreated(address indexed creator, uint256 indexed tokenId, (uint32,uint32,uint192,string,string) questData)
func (_Quest *QuestFilterer) WatchQuestCreated(opts *bind.WatchOpts, sink chan<- *QuestQuestCreated, creator []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Quest.contract.WatchLogs(opts, "QuestCreated", creatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestQuestCreated)
				if err := _Quest.contract.UnpackLog(event, "QuestCreated", log); err != nil {
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

// ParseQuestCreated is a log parse operation binding the contract event 0x7dcb53be1d5e10c26d363e67d5c06d05da75f184de604abd3f9e41efdb7ee576.
//
// Solidity: event QuestCreated(address indexed creator, uint256 indexed tokenId, (uint32,uint32,uint192,string,string) questData)
func (_Quest *QuestFilterer) ParseQuestCreated(log types.Log) (*QuestQuestCreated, error) {
	event := new(QuestQuestCreated)
	if err := _Quest.contract.UnpackLog(event, "QuestCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestSetMinterIterator is returned from FilterSetMinter and is used to iterate over the raw logs and unpacked data for SetMinter events raised by the Quest contract.
type QuestSetMinterIterator struct {
	Event *QuestSetMinter // Event containing the contract specifics and raw log

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
func (it *QuestSetMinterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestSetMinter)
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
		it.Event = new(QuestSetMinter)
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
func (it *QuestSetMinterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestSetMinterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestSetMinter represents a SetMinter event raised by the Quest contract.
type QuestSetMinter struct {
	Minter  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetMinter is a free log retrieval operation binding the contract event 0x1f96bc657d385fd83da973a43f2ad969e6d96b6779b779571a7306db7ca1cd00.
//
// Solidity: event SetMinter(address minter, bool enabled)
func (_Quest *QuestFilterer) FilterSetMinter(opts *bind.FilterOpts) (*QuestSetMinterIterator, error) {

	logs, sub, err := _Quest.contract.FilterLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return &QuestSetMinterIterator{contract: _Quest.contract, event: "SetMinter", logs: logs, sub: sub}, nil
}

// WatchSetMinter is a free log subscription operation binding the contract event 0x1f96bc657d385fd83da973a43f2ad969e6d96b6779b779571a7306db7ca1cd00.
//
// Solidity: event SetMinter(address minter, bool enabled)
func (_Quest *QuestFilterer) WatchSetMinter(opts *bind.WatchOpts, sink chan<- *QuestSetMinter) (event.Subscription, error) {

	logs, sub, err := _Quest.contract.WatchLogs(opts, "SetMinter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestSetMinter)
				if err := _Quest.contract.UnpackLog(event, "SetMinter", log); err != nil {
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

// ParseSetMinter is a log parse operation binding the contract event 0x1f96bc657d385fd83da973a43f2ad969e6d96b6779b779571a7306db7ca1cd00.
//
// Solidity: event SetMinter(address minter, bool enabled)
func (_Quest *QuestFilterer) ParseSetMinter(log types.Log) (*QuestSetMinter, error) {
	event := new(QuestSetMinter)
	if err := _Quest.contract.UnpackLog(event, "SetMinter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Quest contract.
type QuestTransferIterator struct {
	Event *QuestTransfer // Event containing the contract specifics and raw log

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
func (it *QuestTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestTransfer)
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
		it.Event = new(QuestTransfer)
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
func (it *QuestTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestTransfer represents a Transfer event raised by the Quest contract.
type QuestTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Quest *QuestFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*QuestTransferIterator, error) {

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

	logs, sub, err := _Quest.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuestTransferIterator{contract: _Quest.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Quest *QuestFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *QuestTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Quest.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestTransfer)
				if err := _Quest.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Quest *QuestFilterer) ParseTransfer(log types.Log) (*QuestTransfer, error) {
	event := new(QuestTransfer)
	if err := _Quest.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the Quest contract.
type QuestUnlockedIterator struct {
	Event *QuestUnlocked // Event containing the contract specifics and raw log

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
func (it *QuestUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestUnlocked)
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
		it.Event = new(QuestUnlocked)
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
func (it *QuestUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestUnlocked represents a Unlocked event raised by the Quest contract.
type QuestUnlocked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 tokenId)
func (_Quest *QuestFilterer) FilterUnlocked(opts *bind.FilterOpts) (*QuestUnlockedIterator, error) {

	logs, sub, err := _Quest.contract.FilterLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return &QuestUnlockedIterator{contract: _Quest.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 tokenId)
func (_Quest *QuestFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *QuestUnlocked) (event.Subscription, error) {

	logs, sub, err := _Quest.contract.WatchLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestUnlocked)
				if err := _Quest.contract.UnpackLog(event, "Unlocked", log); err != nil {
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
func (_Quest *QuestFilterer) ParseUnlocked(log types.Log) (*QuestUnlocked, error) {
	event := new(QuestUnlocked)
	if err := _Quest.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
