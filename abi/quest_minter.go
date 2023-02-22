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

// QuestMinterMetaData contains all meta data concerning the QuestMinter contract.
var QuestMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Airdroped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Donation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"airdropBadge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"badge\",\"outputs\":[{\"internalType\":\"contractIBadge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"uint192\",\"name\":\"supply\",\"type\":\"uint192\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"createQuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badge_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quest_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quest\",\"outputs\":[{\"internalType\":\"contractIQuest\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"scores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"setBadgeURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QuestMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use QuestMinterMetaData.ABI instead.
var QuestMinterABI = QuestMinterMetaData.ABI

// QuestMinter is an auto generated Go binding around an Ethereum contract.
type QuestMinter struct {
	QuestMinterCaller     // Read-only binding to the contract
	QuestMinterTransactor // Write-only binding to the contract
	QuestMinterFilterer   // Log filterer for contract events
}

// QuestMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuestMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuestMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuestMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuestMinterSession struct {
	Contract     *QuestMinter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuestMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuestMinterCallerSession struct {
	Contract *QuestMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// QuestMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuestMinterTransactorSession struct {
	Contract     *QuestMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// QuestMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuestMinterRaw struct {
	Contract *QuestMinter // Generic contract binding to access the raw methods on
}

// QuestMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuestMinterCallerRaw struct {
	Contract *QuestMinterCaller // Generic read-only contract binding to access the raw methods on
}

// QuestMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuestMinterTransactorRaw struct {
	Contract *QuestMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuestMinter creates a new instance of QuestMinter, bound to a specific deployed contract.
func NewQuestMinter(address common.Address, backend bind.ContractBackend) (*QuestMinter, error) {
	contract, err := bindQuestMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuestMinter{QuestMinterCaller: QuestMinterCaller{contract: contract}, QuestMinterTransactor: QuestMinterTransactor{contract: contract}, QuestMinterFilterer: QuestMinterFilterer{contract: contract}}, nil
}

// NewQuestMinterCaller creates a new read-only instance of QuestMinter, bound to a specific deployed contract.
func NewQuestMinterCaller(address common.Address, caller bind.ContractCaller) (*QuestMinterCaller, error) {
	contract, err := bindQuestMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuestMinterCaller{contract: contract}, nil
}

// NewQuestMinterTransactor creates a new write-only instance of QuestMinter, bound to a specific deployed contract.
func NewQuestMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*QuestMinterTransactor, error) {
	contract, err := bindQuestMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuestMinterTransactor{contract: contract}, nil
}

// NewQuestMinterFilterer creates a new log filterer instance of QuestMinter, bound to a specific deployed contract.
func NewQuestMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*QuestMinterFilterer, error) {
	contract, err := bindQuestMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuestMinterFilterer{contract: contract}, nil
}

// bindQuestMinter binds a generic wrapper to an already deployed contract.
func bindQuestMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuestMinterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestMinter *QuestMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestMinter.Contract.QuestMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestMinter *QuestMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestMinter.Contract.QuestMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestMinter *QuestMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestMinter.Contract.QuestMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestMinter *QuestMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestMinter *QuestMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestMinter *QuestMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestMinter.Contract.contract.Transact(opts, method, params...)
}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_QuestMinter *QuestMinterCaller) Badge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestMinter.contract.Call(opts, &out, "badge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_QuestMinter *QuestMinterSession) Badge() (common.Address, error) {
	return _QuestMinter.Contract.Badge(&_QuestMinter.CallOpts)
}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_QuestMinter *QuestMinterCallerSession) Badge() (common.Address, error) {
	return _QuestMinter.Contract.Badge(&_QuestMinter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestMinter *QuestMinterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestMinter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestMinter *QuestMinterSession) Owner() (common.Address, error) {
	return _QuestMinter.Contract.Owner(&_QuestMinter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestMinter *QuestMinterCallerSession) Owner() (common.Address, error) {
	return _QuestMinter.Contract.Owner(&_QuestMinter.CallOpts)
}

// Quest is a free data retrieval call binding the contract method 0x42cccf08.
//
// Solidity: function quest() view returns(address)
func (_QuestMinter *QuestMinterCaller) Quest(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestMinter.contract.Call(opts, &out, "quest")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Quest is a free data retrieval call binding the contract method 0x42cccf08.
//
// Solidity: function quest() view returns(address)
func (_QuestMinter *QuestMinterSession) Quest() (common.Address, error) {
	return _QuestMinter.Contract.Quest(&_QuestMinter.CallOpts)
}

// Quest is a free data retrieval call binding the contract method 0x42cccf08.
//
// Solidity: function quest() view returns(address)
func (_QuestMinter *QuestMinterCallerSession) Quest() (common.Address, error) {
	return _QuestMinter.Contract.Quest(&_QuestMinter.CallOpts)
}

// Scores is a free data retrieval call binding the contract method 0x7a07b674.
//
// Solidity: function scores(uint256 , address ) view returns(uint256)
func (_QuestMinter *QuestMinterCaller) Scores(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _QuestMinter.contract.Call(opts, &out, "scores", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Scores is a free data retrieval call binding the contract method 0x7a07b674.
//
// Solidity: function scores(uint256 , address ) view returns(uint256)
func (_QuestMinter *QuestMinterSession) Scores(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _QuestMinter.Contract.Scores(&_QuestMinter.CallOpts, arg0, arg1)
}

// Scores is a free data retrieval call binding the contract method 0x7a07b674.
//
// Solidity: function scores(uint256 , address ) view returns(uint256)
func (_QuestMinter *QuestMinterCallerSession) Scores(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _QuestMinter.Contract.Scores(&_QuestMinter.CallOpts, arg0, arg1)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_QuestMinter *QuestMinterCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestMinter.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_QuestMinter *QuestMinterSession) Signer() (common.Address, error) {
	return _QuestMinter.Contract.Signer(&_QuestMinter.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_QuestMinter *QuestMinterCallerSession) Signer() (common.Address, error) {
	return _QuestMinter.Contract.Signer(&_QuestMinter.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_QuestMinter *QuestMinterCaller) StartTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuestMinter.contract.Call(opts, &out, "startTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_QuestMinter *QuestMinterSession) StartTokenId() (*big.Int, error) {
	return _QuestMinter.Contract.StartTokenId(&_QuestMinter.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_QuestMinter *QuestMinterCallerSession) StartTokenId() (*big.Int, error) {
	return _QuestMinter.Contract.StartTokenId(&_QuestMinter.CallOpts)
}

// AirdropBadge is a paid mutator transaction binding the contract method 0x308a2099.
//
// Solidity: function airdropBadge(uint256 tokenId, address[] receivers, bytes signature) returns()
func (_QuestMinter *QuestMinterTransactor) AirdropBadge(opts *bind.TransactOpts, tokenId *big.Int, receivers []common.Address, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.contract.Transact(opts, "airdropBadge", tokenId, receivers, signature)
}

// AirdropBadge is a paid mutator transaction binding the contract method 0x308a2099.
//
// Solidity: function airdropBadge(uint256 tokenId, address[] receivers, bytes signature) returns()
func (_QuestMinter *QuestMinterSession) AirdropBadge(tokenId *big.Int, receivers []common.Address, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.AirdropBadge(&_QuestMinter.TransactOpts, tokenId, receivers, signature)
}

// AirdropBadge is a paid mutator transaction binding the contract method 0x308a2099.
//
// Solidity: function airdropBadge(uint256 tokenId, address[] receivers, bytes signature) returns()
func (_QuestMinter *QuestMinterTransactorSession) AirdropBadge(tokenId *big.Int, receivers []common.Address, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.AirdropBadge(&_QuestMinter.TransactOpts, tokenId, receivers, signature)
}

// Claim is a paid mutator transaction binding the contract method 0x5eddd157.
//
// Solidity: function claim(uint256 tokenId, uint256 score, bytes signature) payable returns()
func (_QuestMinter *QuestMinterTransactor) Claim(opts *bind.TransactOpts, tokenId *big.Int, score *big.Int, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.contract.Transact(opts, "claim", tokenId, score, signature)
}

// Claim is a paid mutator transaction binding the contract method 0x5eddd157.
//
// Solidity: function claim(uint256 tokenId, uint256 score, bytes signature) payable returns()
func (_QuestMinter *QuestMinterSession) Claim(tokenId *big.Int, score *big.Int, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.Claim(&_QuestMinter.TransactOpts, tokenId, score, signature)
}

// Claim is a paid mutator transaction binding the contract method 0x5eddd157.
//
// Solidity: function claim(uint256 tokenId, uint256 score, bytes signature) payable returns()
func (_QuestMinter *QuestMinterTransactorSession) Claim(tokenId *big.Int, score *big.Int, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.Claim(&_QuestMinter.TransactOpts, tokenId, score, signature)
}

// CreateQuest is a paid mutator transaction binding the contract method 0x7fac8ce4.
//
// Solidity: function createQuest((uint32,uint32,uint192,string,string) questData, bytes signature) returns()
func (_QuestMinter *QuestMinterTransactor) CreateQuest(opts *bind.TransactOpts, questData IQuestQuestData, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.contract.Transact(opts, "createQuest", questData, signature)
}

// CreateQuest is a paid mutator transaction binding the contract method 0x7fac8ce4.
//
// Solidity: function createQuest((uint32,uint32,uint192,string,string) questData, bytes signature) returns()
func (_QuestMinter *QuestMinterSession) CreateQuest(questData IQuestQuestData, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.CreateQuest(&_QuestMinter.TransactOpts, questData, signature)
}

// CreateQuest is a paid mutator transaction binding the contract method 0x7fac8ce4.
//
// Solidity: function createQuest((uint32,uint32,uint192,string,string) questData, bytes signature) returns()
func (_QuestMinter *QuestMinterTransactorSession) CreateQuest(questData IQuestQuestData, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.CreateQuest(&_QuestMinter.TransactOpts, questData, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address badge_, address quest_) returns()
func (_QuestMinter *QuestMinterTransactor) Initialize(opts *bind.TransactOpts, badge_ common.Address, quest_ common.Address) (*types.Transaction, error) {
	return _QuestMinter.contract.Transact(opts, "initialize", badge_, quest_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address badge_, address quest_) returns()
func (_QuestMinter *QuestMinterSession) Initialize(badge_ common.Address, quest_ common.Address) (*types.Transaction, error) {
	return _QuestMinter.Contract.Initialize(&_QuestMinter.TransactOpts, badge_, quest_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address badge_, address quest_) returns()
func (_QuestMinter *QuestMinterTransactorSession) Initialize(badge_ common.Address, quest_ common.Address) (*types.Transaction, error) {
	return _QuestMinter.Contract.Initialize(&_QuestMinter.TransactOpts, badge_, quest_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestMinter *QuestMinterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestMinter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestMinter *QuestMinterSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuestMinter.Contract.RenounceOwnership(&_QuestMinter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestMinter *QuestMinterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuestMinter.Contract.RenounceOwnership(&_QuestMinter.TransactOpts)
}

// SetBadgeURI is a paid mutator transaction binding the contract method 0x70069010.
//
// Solidity: function setBadgeURI(uint256 tokenId, string uri, bytes signature) returns()
func (_QuestMinter *QuestMinterTransactor) SetBadgeURI(opts *bind.TransactOpts, tokenId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.contract.Transact(opts, "setBadgeURI", tokenId, uri, signature)
}

// SetBadgeURI is a paid mutator transaction binding the contract method 0x70069010.
//
// Solidity: function setBadgeURI(uint256 tokenId, string uri, bytes signature) returns()
func (_QuestMinter *QuestMinterSession) SetBadgeURI(tokenId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.SetBadgeURI(&_QuestMinter.TransactOpts, tokenId, uri, signature)
}

// SetBadgeURI is a paid mutator transaction binding the contract method 0x70069010.
//
// Solidity: function setBadgeURI(uint256 tokenId, string uri, bytes signature) returns()
func (_QuestMinter *QuestMinterTransactorSession) SetBadgeURI(tokenId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.SetBadgeURI(&_QuestMinter.TransactOpts, tokenId, uri, signature)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_QuestMinter *QuestMinterTransactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _QuestMinter.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_QuestMinter *QuestMinterSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _QuestMinter.Contract.SetSigner(&_QuestMinter.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_QuestMinter *QuestMinterTransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _QuestMinter.Contract.SetSigner(&_QuestMinter.TransactOpts, signer_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestMinter *QuestMinterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QuestMinter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestMinter *QuestMinterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuestMinter.Contract.TransferOwnership(&_QuestMinter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestMinter *QuestMinterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuestMinter.Contract.TransferOwnership(&_QuestMinter.TransactOpts, newOwner)
}

// UpdateScore is a paid mutator transaction binding the contract method 0x385aa0d7.
//
// Solidity: function updateScore(uint256 tokenId, uint256 score, bytes signature) returns()
func (_QuestMinter *QuestMinterTransactor) UpdateScore(opts *bind.TransactOpts, tokenId *big.Int, score *big.Int, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.contract.Transact(opts, "updateScore", tokenId, score, signature)
}

// UpdateScore is a paid mutator transaction binding the contract method 0x385aa0d7.
//
// Solidity: function updateScore(uint256 tokenId, uint256 score, bytes signature) returns()
func (_QuestMinter *QuestMinterSession) UpdateScore(tokenId *big.Int, score *big.Int, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.UpdateScore(&_QuestMinter.TransactOpts, tokenId, score, signature)
}

// UpdateScore is a paid mutator transaction binding the contract method 0x385aa0d7.
//
// Solidity: function updateScore(uint256 tokenId, uint256 score, bytes signature) returns()
func (_QuestMinter *QuestMinterTransactorSession) UpdateScore(tokenId *big.Int, score *big.Int, signature []byte) (*types.Transaction, error) {
	return _QuestMinter.Contract.UpdateScore(&_QuestMinter.TransactOpts, tokenId, score, signature)
}

// QuestMinterAirdropedIterator is returned from FilterAirdroped and is used to iterate over the raw logs and unpacked data for Airdroped events raised by the QuestMinter contract.
type QuestMinterAirdropedIterator struct {
	Event *QuestMinterAirdroped // Event containing the contract specifics and raw log

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
func (it *QuestMinterAirdropedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestMinterAirdroped)
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
		it.Event = new(QuestMinterAirdroped)
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
func (it *QuestMinterAirdropedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestMinterAirdropedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestMinterAirdroped represents a Airdroped event raised by the QuestMinter contract.
type QuestMinterAirdroped struct {
	TokenId *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAirdroped is a free log retrieval operation binding the contract event 0x9842c965759b69b9e1a1576443934aa11f303635dd908a3176f22df019fa6092.
//
// Solidity: event Airdroped(uint256 indexed tokenId, address indexed to)
func (_QuestMinter *QuestMinterFilterer) FilterAirdroped(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address) (*QuestMinterAirdropedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QuestMinter.contract.FilterLogs(opts, "Airdroped", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &QuestMinterAirdropedIterator{contract: _QuestMinter.contract, event: "Airdroped", logs: logs, sub: sub}, nil
}

// WatchAirdroped is a free log subscription operation binding the contract event 0x9842c965759b69b9e1a1576443934aa11f303635dd908a3176f22df019fa6092.
//
// Solidity: event Airdroped(uint256 indexed tokenId, address indexed to)
func (_QuestMinter *QuestMinterFilterer) WatchAirdroped(opts *bind.WatchOpts, sink chan<- *QuestMinterAirdroped, tokenId []*big.Int, to []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _QuestMinter.contract.WatchLogs(opts, "Airdroped", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestMinterAirdroped)
				if err := _QuestMinter.contract.UnpackLog(event, "Airdroped", log); err != nil {
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

// ParseAirdroped is a log parse operation binding the contract event 0x9842c965759b69b9e1a1576443934aa11f303635dd908a3176f22df019fa6092.
//
// Solidity: event Airdroped(uint256 indexed tokenId, address indexed to)
func (_QuestMinter *QuestMinterFilterer) ParseAirdroped(log types.Log) (*QuestMinterAirdroped, error) {
	event := new(QuestMinterAirdroped)
	if err := _QuestMinter.contract.UnpackLog(event, "Airdroped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestMinterClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the QuestMinter contract.
type QuestMinterClaimedIterator struct {
	Event *QuestMinterClaimed // Event containing the contract specifics and raw log

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
func (it *QuestMinterClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestMinterClaimed)
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
		it.Event = new(QuestMinterClaimed)
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
func (it *QuestMinterClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestMinterClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestMinterClaimed represents a Claimed event raised by the QuestMinter contract.
type QuestMinterClaimed struct {
	TokenId *big.Int
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x6aa3eac93d079e5e100b1029be716caa33586c96aa4baac390669fb5c2a21212.
//
// Solidity: event Claimed(uint256 indexed tokenId, address indexed sender)
func (_QuestMinter *QuestMinterFilterer) FilterClaimed(opts *bind.FilterOpts, tokenId []*big.Int, sender []common.Address) (*QuestMinterClaimedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _QuestMinter.contract.FilterLogs(opts, "Claimed", tokenIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &QuestMinterClaimedIterator{contract: _QuestMinter.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x6aa3eac93d079e5e100b1029be716caa33586c96aa4baac390669fb5c2a21212.
//
// Solidity: event Claimed(uint256 indexed tokenId, address indexed sender)
func (_QuestMinter *QuestMinterFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *QuestMinterClaimed, tokenId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _QuestMinter.contract.WatchLogs(opts, "Claimed", tokenIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestMinterClaimed)
				if err := _QuestMinter.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x6aa3eac93d079e5e100b1029be716caa33586c96aa4baac390669fb5c2a21212.
//
// Solidity: event Claimed(uint256 indexed tokenId, address indexed sender)
func (_QuestMinter *QuestMinterFilterer) ParseClaimed(log types.Log) (*QuestMinterClaimed, error) {
	event := new(QuestMinterClaimed)
	if err := _QuestMinter.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestMinterDonationIterator is returned from FilterDonation and is used to iterate over the raw logs and unpacked data for Donation events raised by the QuestMinter contract.
type QuestMinterDonationIterator struct {
	Event *QuestMinterDonation // Event containing the contract specifics and raw log

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
func (it *QuestMinterDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestMinterDonation)
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
		it.Event = new(QuestMinterDonation)
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
func (it *QuestMinterDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestMinterDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestMinterDonation represents a Donation event raised by the QuestMinter contract.
type QuestMinterDonation struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonation is a free log retrieval operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_QuestMinter *QuestMinterFilterer) FilterDonation(opts *bind.FilterOpts) (*QuestMinterDonationIterator, error) {

	logs, sub, err := _QuestMinter.contract.FilterLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return &QuestMinterDonationIterator{contract: _QuestMinter.contract, event: "Donation", logs: logs, sub: sub}, nil
}

// WatchDonation is a free log subscription operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_QuestMinter *QuestMinterFilterer) WatchDonation(opts *bind.WatchOpts, sink chan<- *QuestMinterDonation) (event.Subscription, error) {

	logs, sub, err := _QuestMinter.contract.WatchLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestMinterDonation)
				if err := _QuestMinter.contract.UnpackLog(event, "Donation", log); err != nil {
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
func (_QuestMinter *QuestMinterFilterer) ParseDonation(log types.Log) (*QuestMinterDonation, error) {
	event := new(QuestMinterDonation)
	if err := _QuestMinter.contract.UnpackLog(event, "Donation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestMinterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the QuestMinter contract.
type QuestMinterInitializedIterator struct {
	Event *QuestMinterInitialized // Event containing the contract specifics and raw log

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
func (it *QuestMinterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestMinterInitialized)
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
		it.Event = new(QuestMinterInitialized)
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
func (it *QuestMinterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestMinterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestMinterInitialized represents a Initialized event raised by the QuestMinter contract.
type QuestMinterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QuestMinter *QuestMinterFilterer) FilterInitialized(opts *bind.FilterOpts) (*QuestMinterInitializedIterator, error) {

	logs, sub, err := _QuestMinter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &QuestMinterInitializedIterator{contract: _QuestMinter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_QuestMinter *QuestMinterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *QuestMinterInitialized) (event.Subscription, error) {

	logs, sub, err := _QuestMinter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestMinterInitialized)
				if err := _QuestMinter.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_QuestMinter *QuestMinterFilterer) ParseInitialized(log types.Log) (*QuestMinterInitialized, error) {
	event := new(QuestMinterInitialized)
	if err := _QuestMinter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestMinterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QuestMinter contract.
type QuestMinterOwnershipTransferredIterator struct {
	Event *QuestMinterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QuestMinterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestMinterOwnershipTransferred)
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
		it.Event = new(QuestMinterOwnershipTransferred)
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
func (it *QuestMinterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestMinterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestMinterOwnershipTransferred represents a OwnershipTransferred event raised by the QuestMinter contract.
type QuestMinterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuestMinter *QuestMinterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QuestMinterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuestMinter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuestMinterOwnershipTransferredIterator{contract: _QuestMinter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuestMinter *QuestMinterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QuestMinterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuestMinter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestMinterOwnershipTransferred)
				if err := _QuestMinter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QuestMinter *QuestMinterFilterer) ParseOwnershipTransferred(log types.Log) (*QuestMinterOwnershipTransferred, error) {
	event := new(QuestMinterOwnershipTransferred)
	if err := _QuestMinter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestMinterSignerChangedIterator is returned from FilterSignerChanged and is used to iterate over the raw logs and unpacked data for SignerChanged events raised by the QuestMinter contract.
type QuestMinterSignerChangedIterator struct {
	Event *QuestMinterSignerChanged // Event containing the contract specifics and raw log

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
func (it *QuestMinterSignerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestMinterSignerChanged)
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
		it.Event = new(QuestMinterSignerChanged)
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
func (it *QuestMinterSignerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestMinterSignerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestMinterSignerChanged represents a SignerChanged event raised by the QuestMinter contract.
type QuestMinterSignerChanged struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerChanged is a free log retrieval operation binding the contract event 0x5719a5656c5cfdaafa148ecf366fd3b0a7fae06449ce2a46225977fb7417e29d.
//
// Solidity: event SignerChanged(address signer)
func (_QuestMinter *QuestMinterFilterer) FilterSignerChanged(opts *bind.FilterOpts) (*QuestMinterSignerChangedIterator, error) {

	logs, sub, err := _QuestMinter.contract.FilterLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return &QuestMinterSignerChangedIterator{contract: _QuestMinter.contract, event: "SignerChanged", logs: logs, sub: sub}, nil
}

// WatchSignerChanged is a free log subscription operation binding the contract event 0x5719a5656c5cfdaafa148ecf366fd3b0a7fae06449ce2a46225977fb7417e29d.
//
// Solidity: event SignerChanged(address signer)
func (_QuestMinter *QuestMinterFilterer) WatchSignerChanged(opts *bind.WatchOpts, sink chan<- *QuestMinterSignerChanged) (event.Subscription, error) {

	logs, sub, err := _QuestMinter.contract.WatchLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestMinterSignerChanged)
				if err := _QuestMinter.contract.UnpackLog(event, "SignerChanged", log); err != nil {
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

// ParseSignerChanged is a log parse operation binding the contract event 0x5719a5656c5cfdaafa148ecf366fd3b0a7fae06449ce2a46225977fb7417e29d.
//
// Solidity: event SignerChanged(address signer)
func (_QuestMinter *QuestMinterFilterer) ParseSignerChanged(log types.Log) (*QuestMinterSignerChanged, error) {
	event := new(QuestMinterSignerChanged)
	if err := _QuestMinter.contract.UnpackLog(event, "SignerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
