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

// IQuestQuestData is an auto generated low-level Go binding around an user-defined struct.
type IQuestQuestData struct {
	StartTs uint32
	EndTs   uint32
	Title   string
	Uri     string
}

// QuestV2MetaData contains all meta data concerning the QuestV2 contract.
var QuestV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ClaimedCannotModify\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyMinter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"badgeNum\",\"type\":\"uint256\"}],\"name\":\"BadgeNumUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Donation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"MinterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"}],\"name\":\"QuestCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"}],\"name\":\"QuestModified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"}],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"}],\"name\":\"getBadgeNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getQuest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"locked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"meta\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"}],\"name\":\"modifyQuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"questBadgeNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"quests\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_meta\",\"type\":\"address\"}],\"name\":\"setMetaContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"badgeNum\",\"type\":\"uint256\"}],\"name\":\"updateBadgeNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QuestV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use QuestV2MetaData.ABI instead.
var QuestV2ABI = QuestV2MetaData.ABI

// QuestV2 is an auto generated Go binding around an Ethereum contract.
type QuestV2 struct {
	QuestV2Caller     // Read-only binding to the contract
	QuestV2Transactor // Write-only binding to the contract
	QuestV2Filterer   // Log filterer for contract events
}

// QuestV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type QuestV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type QuestV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuestV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuestV2Session struct {
	Contract     *QuestV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuestV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuestV2CallerSession struct {
	Contract *QuestV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// QuestV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuestV2TransactorSession struct {
	Contract     *QuestV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// QuestV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type QuestV2Raw struct {
	Contract *QuestV2 // Generic contract binding to access the raw methods on
}

// QuestV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuestV2CallerRaw struct {
	Contract *QuestV2Caller // Generic read-only contract binding to access the raw methods on
}

// QuestV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuestV2TransactorRaw struct {
	Contract *QuestV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewQuestV2 creates a new instance of QuestV2, bound to a specific deployed contract.
func NewQuestV2(address common.Address, backend bind.ContractBackend) (*QuestV2, error) {
	contract, err := bindQuestV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuestV2{QuestV2Caller: QuestV2Caller{contract: contract}, QuestV2Transactor: QuestV2Transactor{contract: contract}, QuestV2Filterer: QuestV2Filterer{contract: contract}}, nil
}

// NewQuestV2Caller creates a new read-only instance of QuestV2, bound to a specific deployed contract.
func NewQuestV2Caller(address common.Address, caller bind.ContractCaller) (*QuestV2Caller, error) {
	contract, err := bindQuestV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuestV2Caller{contract: contract}, nil
}

// NewQuestV2Transactor creates a new write-only instance of QuestV2, bound to a specific deployed contract.
func NewQuestV2Transactor(address common.Address, transactor bind.ContractTransactor) (*QuestV2Transactor, error) {
	contract, err := bindQuestV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuestV2Transactor{contract: contract}, nil
}

// NewQuestV2Filterer creates a new log filterer instance of QuestV2, bound to a specific deployed contract.
func NewQuestV2Filterer(address common.Address, filterer bind.ContractFilterer) (*QuestV2Filterer, error) {
	contract, err := bindQuestV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuestV2Filterer{contract: contract}, nil
}

// bindQuestV2 binds a generic wrapper to an already deployed contract.
func bindQuestV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuestV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestV2 *QuestV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestV2.Contract.QuestV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestV2 *QuestV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestV2.Contract.QuestV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestV2 *QuestV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestV2.Contract.QuestV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestV2 *QuestV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestV2 *QuestV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestV2 *QuestV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestV2.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_QuestV2 *QuestV2Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_QuestV2 *QuestV2Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _QuestV2.Contract.BalanceOf(&_QuestV2.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_QuestV2 *QuestV2CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _QuestV2.Contract.BalanceOf(&_QuestV2.CallOpts, owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_QuestV2 *QuestV2Caller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_QuestV2 *QuestV2Session) Exists(tokenId *big.Int) (bool, error) {
	return _QuestV2.Contract.Exists(&_QuestV2.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_QuestV2 *QuestV2CallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _QuestV2.Contract.Exists(&_QuestV2.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_QuestV2 *QuestV2Caller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_QuestV2 *QuestV2Session) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _QuestV2.Contract.GetApproved(&_QuestV2.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_QuestV2 *QuestV2CallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _QuestV2.Contract.GetApproved(&_QuestV2.CallOpts, arg0)
}

// GetBadgeNum is a free data retrieval call binding the contract method 0x938a00d6.
//
// Solidity: function getBadgeNum(uint256 questId) view returns(uint256)
func (_QuestV2 *QuestV2Caller) GetBadgeNum(opts *bind.CallOpts, questId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "getBadgeNum", questId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBadgeNum is a free data retrieval call binding the contract method 0x938a00d6.
//
// Solidity: function getBadgeNum(uint256 questId) view returns(uint256)
func (_QuestV2 *QuestV2Session) GetBadgeNum(questId *big.Int) (*big.Int, error) {
	return _QuestV2.Contract.GetBadgeNum(&_QuestV2.CallOpts, questId)
}

// GetBadgeNum is a free data retrieval call binding the contract method 0x938a00d6.
//
// Solidity: function getBadgeNum(uint256 questId) view returns(uint256)
func (_QuestV2 *QuestV2CallerSession) GetBadgeNum(questId *big.Int) (*big.Int, error) {
	return _QuestV2.Contract.GetBadgeNum(&_QuestV2.CallOpts, questId)
}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 tokenId) view returns((uint32,uint32,string,string) questData)
func (_QuestV2 *QuestV2Caller) GetQuest(opts *bind.CallOpts, tokenId *big.Int) (IQuestQuestData, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "getQuest", tokenId)

	if err != nil {
		return *new(IQuestQuestData), err
	}

	out0 := *abi.ConvertType(out[0], new(IQuestQuestData)).(*IQuestQuestData)

	return out0, err

}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 tokenId) view returns((uint32,uint32,string,string) questData)
func (_QuestV2 *QuestV2Session) GetQuest(tokenId *big.Int) (IQuestQuestData, error) {
	return _QuestV2.Contract.GetQuest(&_QuestV2.CallOpts, tokenId)
}

// GetQuest is a free data retrieval call binding the contract method 0x49f86f34.
//
// Solidity: function getQuest(uint256 tokenId) view returns((uint32,uint32,string,string) questData)
func (_QuestV2 *QuestV2CallerSession) GetQuest(tokenId *big.Int) (IQuestQuestData, error) {
	return _QuestV2.Contract.GetQuest(&_QuestV2.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_QuestV2 *QuestV2Caller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_QuestV2 *QuestV2Session) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _QuestV2.Contract.IsApprovedForAll(&_QuestV2.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_QuestV2 *QuestV2CallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _QuestV2.Contract.IsApprovedForAll(&_QuestV2.CallOpts, arg0, arg1)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_QuestV2 *QuestV2Caller) Locked(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "locked", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_QuestV2 *QuestV2Session) Locked(tokenId *big.Int) (bool, error) {
	return _QuestV2.Contract.Locked(&_QuestV2.CallOpts, tokenId)
}

// Locked is a free data retrieval call binding the contract method 0xb45a3c0e.
//
// Solidity: function locked(uint256 tokenId) view returns(bool)
func (_QuestV2 *QuestV2CallerSession) Locked(tokenId *big.Int) (bool, error) {
	return _QuestV2.Contract.Locked(&_QuestV2.CallOpts, tokenId)
}

// Meta is a free data retrieval call binding the contract method 0xc885044e.
//
// Solidity: function meta() view returns(address)
func (_QuestV2 *QuestV2Caller) Meta(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "meta")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Meta is a free data retrieval call binding the contract method 0xc885044e.
//
// Solidity: function meta() view returns(address)
func (_QuestV2 *QuestV2Session) Meta() (common.Address, error) {
	return _QuestV2.Contract.Meta(&_QuestV2.CallOpts)
}

// Meta is a free data retrieval call binding the contract method 0xc885044e.
//
// Solidity: function meta() view returns(address)
func (_QuestV2 *QuestV2CallerSession) Meta() (common.Address, error) {
	return _QuestV2.Contract.Meta(&_QuestV2.CallOpts)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_QuestV2 *QuestV2Caller) Minters(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "minters", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_QuestV2 *QuestV2Session) Minters(arg0 common.Address) (bool, error) {
	return _QuestV2.Contract.Minters(&_QuestV2.CallOpts, arg0)
}

// Minters is a free data retrieval call binding the contract method 0xf46eccc4.
//
// Solidity: function minters(address ) view returns(bool)
func (_QuestV2 *QuestV2CallerSession) Minters(arg0 common.Address) (bool, error) {
	return _QuestV2.Contract.Minters(&_QuestV2.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QuestV2 *QuestV2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QuestV2 *QuestV2Session) Name() (string, error) {
	return _QuestV2.Contract.Name(&_QuestV2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_QuestV2 *QuestV2CallerSession) Name() (string, error) {
	return _QuestV2.Contract.Name(&_QuestV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestV2 *QuestV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestV2 *QuestV2Session) Owner() (common.Address, error) {
	return _QuestV2.Contract.Owner(&_QuestV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestV2 *QuestV2CallerSession) Owner() (common.Address, error) {
	return _QuestV2.Contract.Owner(&_QuestV2.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_QuestV2 *QuestV2Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_QuestV2 *QuestV2Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _QuestV2.Contract.OwnerOf(&_QuestV2.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_QuestV2 *QuestV2CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _QuestV2.Contract.OwnerOf(&_QuestV2.CallOpts, tokenId)
}

// QuestBadgeNum is a free data retrieval call binding the contract method 0xad113eb3.
//
// Solidity: function questBadgeNum(uint256 ) view returns(uint256)
func (_QuestV2 *QuestV2Caller) QuestBadgeNum(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "questBadgeNum", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuestBadgeNum is a free data retrieval call binding the contract method 0xad113eb3.
//
// Solidity: function questBadgeNum(uint256 ) view returns(uint256)
func (_QuestV2 *QuestV2Session) QuestBadgeNum(arg0 *big.Int) (*big.Int, error) {
	return _QuestV2.Contract.QuestBadgeNum(&_QuestV2.CallOpts, arg0)
}

// QuestBadgeNum is a free data retrieval call binding the contract method 0xad113eb3.
//
// Solidity: function questBadgeNum(uint256 ) view returns(uint256)
func (_QuestV2 *QuestV2CallerSession) QuestBadgeNum(arg0 *big.Int) (*big.Int, error) {
	return _QuestV2.Contract.QuestBadgeNum(&_QuestV2.CallOpts, arg0)
}

// Quests is a free data retrieval call binding the contract method 0xe085f980.
//
// Solidity: function quests(uint256 ) view returns(uint32 startTs, uint32 endTs, string title, string uri)
func (_QuestV2 *QuestV2Caller) Quests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartTs uint32
	EndTs   uint32
	Title   string
	Uri     string
}, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "quests", arg0)

	outstruct := new(struct {
		StartTs uint32
		EndTs   uint32
		Title   string
		Uri     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTs = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.EndTs = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Title = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Uri = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Quests is a free data retrieval call binding the contract method 0xe085f980.
//
// Solidity: function quests(uint256 ) view returns(uint32 startTs, uint32 endTs, string title, string uri)
func (_QuestV2 *QuestV2Session) Quests(arg0 *big.Int) (struct {
	StartTs uint32
	EndTs   uint32
	Title   string
	Uri     string
}, error) {
	return _QuestV2.Contract.Quests(&_QuestV2.CallOpts, arg0)
}

// Quests is a free data retrieval call binding the contract method 0xe085f980.
//
// Solidity: function quests(uint256 ) view returns(uint32 startTs, uint32 endTs, string title, string uri)
func (_QuestV2 *QuestV2CallerSession) Quests(arg0 *big.Int) (struct {
	StartTs uint32
	EndTs   uint32
	Title   string
	Uri     string
}, error) {
	return _QuestV2.Contract.Quests(&_QuestV2.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_QuestV2 *QuestV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_QuestV2 *QuestV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _QuestV2.Contract.SupportsInterface(&_QuestV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_QuestV2 *QuestV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _QuestV2.Contract.SupportsInterface(&_QuestV2.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_QuestV2 *QuestV2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_QuestV2 *QuestV2Session) Symbol() (string, error) {
	return _QuestV2.Contract.Symbol(&_QuestV2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_QuestV2 *QuestV2CallerSession) Symbol() (string, error) {
	return _QuestV2.Contract.Symbol(&_QuestV2.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_QuestV2 *QuestV2Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_QuestV2 *QuestV2Session) TokenURI(tokenId *big.Int) (string, error) {
	return _QuestV2.Contract.TokenURI(&_QuestV2.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_QuestV2 *QuestV2CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _QuestV2.Contract.TokenURI(&_QuestV2.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_QuestV2 *QuestV2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuestV2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_QuestV2 *QuestV2Session) TotalSupply() (*big.Int, error) {
	return _QuestV2.Contract.TotalSupply(&_QuestV2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_QuestV2 *QuestV2CallerSession) TotalSupply() (*big.Int, error) {
	return _QuestV2.Contract.TotalSupply(&_QuestV2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_QuestV2 *QuestV2Transactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_QuestV2 *QuestV2Session) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.Approve(&_QuestV2.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns()
func (_QuestV2 *QuestV2TransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.Approve(&_QuestV2.TransactOpts, arg0, arg1)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 questId) payable returns()
func (_QuestV2 *QuestV2Transactor) Donate(opts *bind.TransactOpts, questId *big.Int) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "donate", questId)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 questId) payable returns()
func (_QuestV2 *QuestV2Session) Donate(questId *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.Donate(&_QuestV2.TransactOpts, questId)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 questId) payable returns()
func (_QuestV2 *QuestV2TransactorSession) Donate(questId *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.Donate(&_QuestV2.TransactOpts, questId)
}

// Mint is a paid mutator transaction binding the contract method 0x78a82014.
//
// Solidity: function mint(address to, uint256 tokenId, (uint32,uint32,string,string) questData, bytes data) returns()
func (_QuestV2 *QuestV2Transactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, questData IQuestQuestData, data []byte) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "mint", to, tokenId, questData, data)
}

// Mint is a paid mutator transaction binding the contract method 0x78a82014.
//
// Solidity: function mint(address to, uint256 tokenId, (uint32,uint32,string,string) questData, bytes data) returns()
func (_QuestV2 *QuestV2Session) Mint(to common.Address, tokenId *big.Int, questData IQuestQuestData, data []byte) (*types.Transaction, error) {
	return _QuestV2.Contract.Mint(&_QuestV2.TransactOpts, to, tokenId, questData, data)
}

// Mint is a paid mutator transaction binding the contract method 0x78a82014.
//
// Solidity: function mint(address to, uint256 tokenId, (uint32,uint32,string,string) questData, bytes data) returns()
func (_QuestV2 *QuestV2TransactorSession) Mint(to common.Address, tokenId *big.Int, questData IQuestQuestData, data []byte) (*types.Transaction, error) {
	return _QuestV2.Contract.Mint(&_QuestV2.TransactOpts, to, tokenId, questData, data)
}

// ModifyQuest is a paid mutator transaction binding the contract method 0xa5999349.
//
// Solidity: function modifyQuest(uint256 tokenId, (uint32,uint32,string,string) questData) returns()
func (_QuestV2 *QuestV2Transactor) ModifyQuest(opts *bind.TransactOpts, tokenId *big.Int, questData IQuestQuestData) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "modifyQuest", tokenId, questData)
}

// ModifyQuest is a paid mutator transaction binding the contract method 0xa5999349.
//
// Solidity: function modifyQuest(uint256 tokenId, (uint32,uint32,string,string) questData) returns()
func (_QuestV2 *QuestV2Session) ModifyQuest(tokenId *big.Int, questData IQuestQuestData) (*types.Transaction, error) {
	return _QuestV2.Contract.ModifyQuest(&_QuestV2.TransactOpts, tokenId, questData)
}

// ModifyQuest is a paid mutator transaction binding the contract method 0xa5999349.
//
// Solidity: function modifyQuest(uint256 tokenId, (uint32,uint32,string,string) questData) returns()
func (_QuestV2 *QuestV2TransactorSession) ModifyQuest(tokenId *big.Int, questData IQuestQuestData) (*types.Transaction, error) {
	return _QuestV2.Contract.ModifyQuest(&_QuestV2.TransactOpts, tokenId, questData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestV2 *QuestV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestV2 *QuestV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _QuestV2.Contract.RenounceOwnership(&_QuestV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestV2 *QuestV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuestV2.Contract.RenounceOwnership(&_QuestV2.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_QuestV2 *QuestV2Transactor) SafeTransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "safeTransferFrom", arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_QuestV2 *QuestV2Session) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.SafeTransferFrom(&_QuestV2.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address , address , uint256 ) returns()
func (_QuestV2 *QuestV2TransactorSession) SafeTransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.SafeTransferFrom(&_QuestV2.TransactOpts, arg0, arg1, arg2)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_QuestV2 *QuestV2Transactor) SafeTransferFrom0(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "safeTransferFrom0", arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_QuestV2 *QuestV2Session) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _QuestV2.Contract.SafeTransferFrom0(&_QuestV2.TransactOpts, arg0, arg1, arg2, arg3)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address , address , uint256 , bytes ) returns()
func (_QuestV2 *QuestV2TransactorSession) SafeTransferFrom0(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _QuestV2.Contract.SafeTransferFrom0(&_QuestV2.TransactOpts, arg0, arg1, arg2, arg3)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_QuestV2 *QuestV2Transactor) SetApprovalForAll(opts *bind.TransactOpts, arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "setApprovalForAll", arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_QuestV2 *QuestV2Session) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _QuestV2.Contract.SetApprovalForAll(&_QuestV2.TransactOpts, arg0, arg1)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address , bool ) returns()
func (_QuestV2 *QuestV2TransactorSession) SetApprovalForAll(arg0 common.Address, arg1 bool) (*types.Transaction, error) {
	return _QuestV2.Contract.SetApprovalForAll(&_QuestV2.TransactOpts, arg0, arg1)
}

// SetMetaContract is a paid mutator transaction binding the contract method 0x0957a4b6.
//
// Solidity: function setMetaContract(address _meta) returns()
func (_QuestV2 *QuestV2Transactor) SetMetaContract(opts *bind.TransactOpts, _meta common.Address) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "setMetaContract", _meta)
}

// SetMetaContract is a paid mutator transaction binding the contract method 0x0957a4b6.
//
// Solidity: function setMetaContract(address _meta) returns()
func (_QuestV2 *QuestV2Session) SetMetaContract(_meta common.Address) (*types.Transaction, error) {
	return _QuestV2.Contract.SetMetaContract(&_QuestV2.TransactOpts, _meta)
}

// SetMetaContract is a paid mutator transaction binding the contract method 0x0957a4b6.
//
// Solidity: function setMetaContract(address _meta) returns()
func (_QuestV2 *QuestV2TransactorSession) SetMetaContract(_meta common.Address) (*types.Transaction, error) {
	return _QuestV2.Contract.SetMetaContract(&_QuestV2.TransactOpts, _meta)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_QuestV2 *QuestV2Transactor) SetMinter(opts *bind.TransactOpts, minter common.Address, enabled bool) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "setMinter", minter, enabled)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_QuestV2 *QuestV2Session) SetMinter(minter common.Address, enabled bool) (*types.Transaction, error) {
	return _QuestV2.Contract.SetMinter(&_QuestV2.TransactOpts, minter, enabled)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address minter, bool enabled) returns()
func (_QuestV2 *QuestV2TransactorSession) SetMinter(minter common.Address, enabled bool) (*types.Transaction, error) {
	return _QuestV2.Contract.SetMinter(&_QuestV2.TransactOpts, minter, enabled)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_QuestV2 *QuestV2Transactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_QuestV2 *QuestV2Session) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.TransferFrom(&_QuestV2.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns()
func (_QuestV2 *QuestV2TransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.TransferFrom(&_QuestV2.TransactOpts, arg0, arg1, arg2)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestV2 *QuestV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestV2 *QuestV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuestV2.Contract.TransferOwnership(&_QuestV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestV2 *QuestV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuestV2.Contract.TransferOwnership(&_QuestV2.TransactOpts, newOwner)
}

// UpdateBadgeNum is a paid mutator transaction binding the contract method 0x9edcf1c3.
//
// Solidity: function updateBadgeNum(uint256 questId, uint256 badgeNum) returns()
func (_QuestV2 *QuestV2Transactor) UpdateBadgeNum(opts *bind.TransactOpts, questId *big.Int, badgeNum *big.Int) (*types.Transaction, error) {
	return _QuestV2.contract.Transact(opts, "updateBadgeNum", questId, badgeNum)
}

// UpdateBadgeNum is a paid mutator transaction binding the contract method 0x9edcf1c3.
//
// Solidity: function updateBadgeNum(uint256 questId, uint256 badgeNum) returns()
func (_QuestV2 *QuestV2Session) UpdateBadgeNum(questId *big.Int, badgeNum *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.UpdateBadgeNum(&_QuestV2.TransactOpts, questId, badgeNum)
}

// UpdateBadgeNum is a paid mutator transaction binding the contract method 0x9edcf1c3.
//
// Solidity: function updateBadgeNum(uint256 questId, uint256 badgeNum) returns()
func (_QuestV2 *QuestV2TransactorSession) UpdateBadgeNum(questId *big.Int, badgeNum *big.Int) (*types.Transaction, error) {
	return _QuestV2.Contract.UpdateBadgeNum(&_QuestV2.TransactOpts, questId, badgeNum)
}

// QuestV2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the QuestV2 contract.
type QuestV2ApprovalIterator struct {
	Event *QuestV2Approval // Event containing the contract specifics and raw log

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
func (it *QuestV2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2Approval)
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
		it.Event = new(QuestV2Approval)
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
func (it *QuestV2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2Approval represents a Approval event raised by the QuestV2 contract.
type QuestV2Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_QuestV2 *QuestV2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*QuestV2ApprovalIterator, error) {

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

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuestV2ApprovalIterator{contract: _QuestV2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_QuestV2 *QuestV2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *QuestV2Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2Approval)
				if err := _QuestV2.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_QuestV2 *QuestV2Filterer) ParseApproval(log types.Log) (*QuestV2Approval, error) {
	event := new(QuestV2Approval)
	if err := _QuestV2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the QuestV2 contract.
type QuestV2ApprovalForAllIterator struct {
	Event *QuestV2ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *QuestV2ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2ApprovalForAll)
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
		it.Event = new(QuestV2ApprovalForAll)
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
func (it *QuestV2ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2ApprovalForAll represents a ApprovalForAll event raised by the QuestV2 contract.
type QuestV2ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_QuestV2 *QuestV2Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*QuestV2ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &QuestV2ApprovalForAllIterator{contract: _QuestV2.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_QuestV2 *QuestV2Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *QuestV2ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2ApprovalForAll)
				if err := _QuestV2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_QuestV2 *QuestV2Filterer) ParseApprovalForAll(log types.Log) (*QuestV2ApprovalForAll, error) {
	event := new(QuestV2ApprovalForAll)
	if err := _QuestV2.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2BadgeNumUpdatedIterator is returned from FilterBadgeNumUpdated and is used to iterate over the raw logs and unpacked data for BadgeNumUpdated events raised by the QuestV2 contract.
type QuestV2BadgeNumUpdatedIterator struct {
	Event *QuestV2BadgeNumUpdated // Event containing the contract specifics and raw log

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
func (it *QuestV2BadgeNumUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2BadgeNumUpdated)
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
		it.Event = new(QuestV2BadgeNumUpdated)
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
func (it *QuestV2BadgeNumUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2BadgeNumUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2BadgeNumUpdated represents a BadgeNumUpdated event raised by the QuestV2 contract.
type QuestV2BadgeNumUpdated struct {
	QuestId  *big.Int
	BadgeNum *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBadgeNumUpdated is a free log retrieval operation binding the contract event 0x359e70231d793fa1d2bdbe036ae9239433156b1cb828cf2c8af77b7c48bc9257.
//
// Solidity: event BadgeNumUpdated(uint256 indexed questId, uint256 badgeNum)
func (_QuestV2 *QuestV2Filterer) FilterBadgeNumUpdated(opts *bind.FilterOpts, questId []*big.Int) (*QuestV2BadgeNumUpdatedIterator, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "BadgeNumUpdated", questIdRule)
	if err != nil {
		return nil, err
	}
	return &QuestV2BadgeNumUpdatedIterator{contract: _QuestV2.contract, event: "BadgeNumUpdated", logs: logs, sub: sub}, nil
}

// WatchBadgeNumUpdated is a free log subscription operation binding the contract event 0x359e70231d793fa1d2bdbe036ae9239433156b1cb828cf2c8af77b7c48bc9257.
//
// Solidity: event BadgeNumUpdated(uint256 indexed questId, uint256 badgeNum)
func (_QuestV2 *QuestV2Filterer) WatchBadgeNumUpdated(opts *bind.WatchOpts, sink chan<- *QuestV2BadgeNumUpdated, questId []*big.Int) (event.Subscription, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "BadgeNumUpdated", questIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2BadgeNumUpdated)
				if err := _QuestV2.contract.UnpackLog(event, "BadgeNumUpdated", log); err != nil {
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

// ParseBadgeNumUpdated is a log parse operation binding the contract event 0x359e70231d793fa1d2bdbe036ae9239433156b1cb828cf2c8af77b7c48bc9257.
//
// Solidity: event BadgeNumUpdated(uint256 indexed questId, uint256 badgeNum)
func (_QuestV2 *QuestV2Filterer) ParseBadgeNumUpdated(log types.Log) (*QuestV2BadgeNumUpdated, error) {
	event := new(QuestV2BadgeNumUpdated)
	if err := _QuestV2.contract.UnpackLog(event, "BadgeNumUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2DonationIterator is returned from FilterDonation and is used to iterate over the raw logs and unpacked data for Donation events raised by the QuestV2 contract.
type QuestV2DonationIterator struct {
	Event *QuestV2Donation // Event containing the contract specifics and raw log

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
func (it *QuestV2DonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2Donation)
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
		it.Event = new(QuestV2Donation)
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
func (it *QuestV2DonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2DonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2Donation represents a Donation event raised by the QuestV2 contract.
type QuestV2Donation struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonation is a free log retrieval operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_QuestV2 *QuestV2Filterer) FilterDonation(opts *bind.FilterOpts) (*QuestV2DonationIterator, error) {

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return &QuestV2DonationIterator{contract: _QuestV2.contract, event: "Donation", logs: logs, sub: sub}, nil
}

// WatchDonation is a free log subscription operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_QuestV2 *QuestV2Filterer) WatchDonation(opts *bind.WatchOpts, sink chan<- *QuestV2Donation) (event.Subscription, error) {

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2Donation)
				if err := _QuestV2.contract.UnpackLog(event, "Donation", log); err != nil {
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
func (_QuestV2 *QuestV2Filterer) ParseDonation(log types.Log) (*QuestV2Donation, error) {
	event := new(QuestV2Donation)
	if err := _QuestV2.contract.UnpackLog(event, "Donation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2LockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the QuestV2 contract.
type QuestV2LockedIterator struct {
	Event *QuestV2Locked // Event containing the contract specifics and raw log

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
func (it *QuestV2LockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2Locked)
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
		it.Event = new(QuestV2Locked)
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
func (it *QuestV2LockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2LockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2Locked represents a Locked event raised by the QuestV2 contract.
type QuestV2Locked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_QuestV2 *QuestV2Filterer) FilterLocked(opts *bind.FilterOpts) (*QuestV2LockedIterator, error) {

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return &QuestV2LockedIterator{contract: _QuestV2.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x032bc66be43dbccb7487781d168eb7bda224628a3b2c3388bdf69b532a3a1611.
//
// Solidity: event Locked(uint256 tokenId)
func (_QuestV2 *QuestV2Filterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *QuestV2Locked) (event.Subscription, error) {

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "Locked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2Locked)
				if err := _QuestV2.contract.UnpackLog(event, "Locked", log); err != nil {
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
func (_QuestV2 *QuestV2Filterer) ParseLocked(log types.Log) (*QuestV2Locked, error) {
	event := new(QuestV2Locked)
	if err := _QuestV2.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2MinterSetIterator is returned from FilterMinterSet and is used to iterate over the raw logs and unpacked data for MinterSet events raised by the QuestV2 contract.
type QuestV2MinterSetIterator struct {
	Event *QuestV2MinterSet // Event containing the contract specifics and raw log

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
func (it *QuestV2MinterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2MinterSet)
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
		it.Event = new(QuestV2MinterSet)
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
func (it *QuestV2MinterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2MinterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2MinterSet represents a MinterSet event raised by the QuestV2 contract.
type QuestV2MinterSet struct {
	Minter  common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterSet is a free log retrieval operation binding the contract event 0x583b0aa0e528532caf4b907c11d7a8158a122fe2a6fb80cd9b09776ebea8d92d.
//
// Solidity: event MinterSet(address minter, bool enabled)
func (_QuestV2 *QuestV2Filterer) FilterMinterSet(opts *bind.FilterOpts) (*QuestV2MinterSetIterator, error) {

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "MinterSet")
	if err != nil {
		return nil, err
	}
	return &QuestV2MinterSetIterator{contract: _QuestV2.contract, event: "MinterSet", logs: logs, sub: sub}, nil
}

// WatchMinterSet is a free log subscription operation binding the contract event 0x583b0aa0e528532caf4b907c11d7a8158a122fe2a6fb80cd9b09776ebea8d92d.
//
// Solidity: event MinterSet(address minter, bool enabled)
func (_QuestV2 *QuestV2Filterer) WatchMinterSet(opts *bind.WatchOpts, sink chan<- *QuestV2MinterSet) (event.Subscription, error) {

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "MinterSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2MinterSet)
				if err := _QuestV2.contract.UnpackLog(event, "MinterSet", log); err != nil {
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
func (_QuestV2 *QuestV2Filterer) ParseMinterSet(log types.Log) (*QuestV2MinterSet, error) {
	event := new(QuestV2MinterSet)
	if err := _QuestV2.contract.UnpackLog(event, "MinterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QuestV2 contract.
type QuestV2OwnershipTransferredIterator struct {
	Event *QuestV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QuestV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2OwnershipTransferred)
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
		it.Event = new(QuestV2OwnershipTransferred)
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
func (it *QuestV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2OwnershipTransferred represents a OwnershipTransferred event raised by the QuestV2 contract.
type QuestV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuestV2 *QuestV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QuestV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuestV2OwnershipTransferredIterator{contract: _QuestV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuestV2 *QuestV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QuestV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2OwnershipTransferred)
				if err := _QuestV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QuestV2 *QuestV2Filterer) ParseOwnershipTransferred(log types.Log) (*QuestV2OwnershipTransferred, error) {
	event := new(QuestV2OwnershipTransferred)
	if err := _QuestV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2QuestCreatedIterator is returned from FilterQuestCreated and is used to iterate over the raw logs and unpacked data for QuestCreated events raised by the QuestV2 contract.
type QuestV2QuestCreatedIterator struct {
	Event *QuestV2QuestCreated // Event containing the contract specifics and raw log

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
func (it *QuestV2QuestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2QuestCreated)
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
		it.Event = new(QuestV2QuestCreated)
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
func (it *QuestV2QuestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2QuestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2QuestCreated represents a QuestCreated event raised by the QuestV2 contract.
type QuestV2QuestCreated struct {
	Creator   common.Address
	TokenId   *big.Int
	QuestData IQuestQuestData
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQuestCreated is a free log retrieval operation binding the contract event 0x9b5add8a56030992fa29ad875df889fdfafcce0ffca1926ae3f6443b9213dca7.
//
// Solidity: event QuestCreated(address indexed creator, uint256 indexed tokenId, (uint32,uint32,string,string) questData)
func (_QuestV2 *QuestV2Filterer) FilterQuestCreated(opts *bind.FilterOpts, creator []common.Address, tokenId []*big.Int) (*QuestV2QuestCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "QuestCreated", creatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuestV2QuestCreatedIterator{contract: _QuestV2.contract, event: "QuestCreated", logs: logs, sub: sub}, nil
}

// WatchQuestCreated is a free log subscription operation binding the contract event 0x9b5add8a56030992fa29ad875df889fdfafcce0ffca1926ae3f6443b9213dca7.
//
// Solidity: event QuestCreated(address indexed creator, uint256 indexed tokenId, (uint32,uint32,string,string) questData)
func (_QuestV2 *QuestV2Filterer) WatchQuestCreated(opts *bind.WatchOpts, sink chan<- *QuestV2QuestCreated, creator []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "QuestCreated", creatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2QuestCreated)
				if err := _QuestV2.contract.UnpackLog(event, "QuestCreated", log); err != nil {
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

// ParseQuestCreated is a log parse operation binding the contract event 0x9b5add8a56030992fa29ad875df889fdfafcce0ffca1926ae3f6443b9213dca7.
//
// Solidity: event QuestCreated(address indexed creator, uint256 indexed tokenId, (uint32,uint32,string,string) questData)
func (_QuestV2 *QuestV2Filterer) ParseQuestCreated(log types.Log) (*QuestV2QuestCreated, error) {
	event := new(QuestV2QuestCreated)
	if err := _QuestV2.contract.UnpackLog(event, "QuestCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2QuestModifiedIterator is returned from FilterQuestModified and is used to iterate over the raw logs and unpacked data for QuestModified events raised by the QuestV2 contract.
type QuestV2QuestModifiedIterator struct {
	Event *QuestV2QuestModified // Event containing the contract specifics and raw log

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
func (it *QuestV2QuestModifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2QuestModified)
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
		it.Event = new(QuestV2QuestModified)
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
func (it *QuestV2QuestModifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2QuestModifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2QuestModified represents a QuestModified event raised by the QuestV2 contract.
type QuestV2QuestModified struct {
	Creator   common.Address
	TokenId   *big.Int
	QuestData IQuestQuestData
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterQuestModified is a free log retrieval operation binding the contract event 0xe92e573a75c7a5704b449c144ebf2fd928bf6a33142daeaefacac81b7d8862e5.
//
// Solidity: event QuestModified(address indexed creator, uint256 indexed tokenId, (uint32,uint32,string,string) questData)
func (_QuestV2 *QuestV2Filterer) FilterQuestModified(opts *bind.FilterOpts, creator []common.Address, tokenId []*big.Int) (*QuestV2QuestModifiedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "QuestModified", creatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuestV2QuestModifiedIterator{contract: _QuestV2.contract, event: "QuestModified", logs: logs, sub: sub}, nil
}

// WatchQuestModified is a free log subscription operation binding the contract event 0xe92e573a75c7a5704b449c144ebf2fd928bf6a33142daeaefacac81b7d8862e5.
//
// Solidity: event QuestModified(address indexed creator, uint256 indexed tokenId, (uint32,uint32,string,string) questData)
func (_QuestV2 *QuestV2Filterer) WatchQuestModified(opts *bind.WatchOpts, sink chan<- *QuestV2QuestModified, creator []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "QuestModified", creatorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2QuestModified)
				if err := _QuestV2.contract.UnpackLog(event, "QuestModified", log); err != nil {
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

// ParseQuestModified is a log parse operation binding the contract event 0xe92e573a75c7a5704b449c144ebf2fd928bf6a33142daeaefacac81b7d8862e5.
//
// Solidity: event QuestModified(address indexed creator, uint256 indexed tokenId, (uint32,uint32,string,string) questData)
func (_QuestV2 *QuestV2Filterer) ParseQuestModified(log types.Log) (*QuestV2QuestModified, error) {
	event := new(QuestV2QuestModified)
	if err := _QuestV2.contract.UnpackLog(event, "QuestModified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the QuestV2 contract.
type QuestV2TransferIterator struct {
	Event *QuestV2Transfer // Event containing the contract specifics and raw log

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
func (it *QuestV2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2Transfer)
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
		it.Event = new(QuestV2Transfer)
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
func (it *QuestV2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2Transfer represents a Transfer event raised by the QuestV2 contract.
type QuestV2Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_QuestV2 *QuestV2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*QuestV2TransferIterator, error) {

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

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuestV2TransferIterator{contract: _QuestV2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_QuestV2 *QuestV2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *QuestV2Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2Transfer)
				if err := _QuestV2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_QuestV2 *QuestV2Filterer) ParseTransfer(log types.Log) (*QuestV2Transfer, error) {
	event := new(QuestV2Transfer)
	if err := _QuestV2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestV2UnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the QuestV2 contract.
type QuestV2UnlockedIterator struct {
	Event *QuestV2Unlocked // Event containing the contract specifics and raw log

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
func (it *QuestV2UnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestV2Unlocked)
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
		it.Event = new(QuestV2Unlocked)
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
func (it *QuestV2UnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestV2UnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestV2Unlocked represents a Unlocked event raised by the QuestV2 contract.
type QuestV2Unlocked struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 tokenId)
func (_QuestV2 *QuestV2Filterer) FilterUnlocked(opts *bind.FilterOpts) (*QuestV2UnlockedIterator, error) {

	logs, sub, err := _QuestV2.contract.FilterLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return &QuestV2UnlockedIterator{contract: _QuestV2.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0xf27b6ce5b2f5e68ddb2fd95a8a909d4ecf1daaac270935fff052feacb24f1842.
//
// Solidity: event Unlocked(uint256 tokenId)
func (_QuestV2 *QuestV2Filterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *QuestV2Unlocked) (event.Subscription, error) {

	logs, sub, err := _QuestV2.contract.WatchLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestV2Unlocked)
				if err := _QuestV2.contract.UnpackLog(event, "Unlocked", log); err != nil {
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
func (_QuestV2 *QuestV2Filterer) ParseUnlocked(log types.Log) (*QuestV2Unlocked, error) {
	event := new(QuestV2Unlocked)
	if err := _QuestV2.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
