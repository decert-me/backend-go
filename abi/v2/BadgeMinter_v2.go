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

// BadgeMinterV2MetaData contains all meta data concerning the BadgeMinterV2 contract.
var BadgeMinterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badge_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCreator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"Airdroped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Donation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"questIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"airdropBadge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"badge\",\"outputs\":[{\"internalType\":\"contractIBadge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIBadge.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claimWithInit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIBadge.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"initQuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"questUri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateQuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BadgeMinterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use BadgeMinterV2MetaData.ABI instead.
var BadgeMinterV2ABI = BadgeMinterV2MetaData.ABI

// BadgeMinterV2 is an auto generated Go binding around an Ethereum contract.
type BadgeMinterV2 struct {
	BadgeMinterV2Caller     // Read-only binding to the contract
	BadgeMinterV2Transactor // Write-only binding to the contract
	BadgeMinterV2Filterer   // Log filterer for contract events
}

// BadgeMinterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type BadgeMinterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeMinterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BadgeMinterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeMinterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BadgeMinterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeMinterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BadgeMinterV2Session struct {
	Contract     *BadgeMinterV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgeMinterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BadgeMinterV2CallerSession struct {
	Contract *BadgeMinterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BadgeMinterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BadgeMinterV2TransactorSession struct {
	Contract     *BadgeMinterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BadgeMinterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type BadgeMinterV2Raw struct {
	Contract *BadgeMinterV2 // Generic contract binding to access the raw methods on
}

// BadgeMinterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BadgeMinterV2CallerRaw struct {
	Contract *BadgeMinterV2Caller // Generic read-only contract binding to access the raw methods on
}

// BadgeMinterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BadgeMinterV2TransactorRaw struct {
	Contract *BadgeMinterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBadgeMinterV2 creates a new instance of BadgeMinterV2, bound to a specific deployed contract.
func NewBadgeMinterV2(address common.Address, backend bind.ContractBackend) (*BadgeMinterV2, error) {
	contract, err := bindBadgeMinterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterV2{BadgeMinterV2Caller: BadgeMinterV2Caller{contract: contract}, BadgeMinterV2Transactor: BadgeMinterV2Transactor{contract: contract}, BadgeMinterV2Filterer: BadgeMinterV2Filterer{contract: contract}}, nil
}

// NewBadgeMinterV2Caller creates a new read-only instance of BadgeMinterV2, bound to a specific deployed contract.
func NewBadgeMinterV2Caller(address common.Address, caller bind.ContractCaller) (*BadgeMinterV2Caller, error) {
	contract, err := bindBadgeMinterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterV2Caller{contract: contract}, nil
}

// NewBadgeMinterV2Transactor creates a new write-only instance of BadgeMinterV2, bound to a specific deployed contract.
func NewBadgeMinterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*BadgeMinterV2Transactor, error) {
	contract, err := bindBadgeMinterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterV2Transactor{contract: contract}, nil
}

// NewBadgeMinterV2Filterer creates a new log filterer instance of BadgeMinterV2, bound to a specific deployed contract.
func NewBadgeMinterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*BadgeMinterV2Filterer, error) {
	contract, err := bindBadgeMinterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterV2Filterer{contract: contract}, nil
}

// bindBadgeMinterV2 binds a generic wrapper to an already deployed contract.
func bindBadgeMinterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BadgeMinterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BadgeMinterV2 *BadgeMinterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BadgeMinterV2.Contract.BadgeMinterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BadgeMinterV2 *BadgeMinterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.BadgeMinterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BadgeMinterV2 *BadgeMinterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.BadgeMinterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BadgeMinterV2 *BadgeMinterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BadgeMinterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BadgeMinterV2 *BadgeMinterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BadgeMinterV2 *BadgeMinterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.contract.Transact(opts, method, params...)
}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_BadgeMinterV2 *BadgeMinterV2Caller) Badge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BadgeMinterV2.contract.Call(opts, &out, "badge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_BadgeMinterV2 *BadgeMinterV2Session) Badge() (common.Address, error) {
	return _BadgeMinterV2.Contract.Badge(&_BadgeMinterV2.CallOpts)
}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_BadgeMinterV2 *BadgeMinterV2CallerSession) Badge() (common.Address, error) {
	return _BadgeMinterV2.Contract.Badge(&_BadgeMinterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BadgeMinterV2 *BadgeMinterV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BadgeMinterV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BadgeMinterV2 *BadgeMinterV2Session) Owner() (common.Address, error) {
	return _BadgeMinterV2.Contract.Owner(&_BadgeMinterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BadgeMinterV2 *BadgeMinterV2CallerSession) Owner() (common.Address, error) {
	return _BadgeMinterV2.Contract.Owner(&_BadgeMinterV2.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_BadgeMinterV2 *BadgeMinterV2Caller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BadgeMinterV2.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_BadgeMinterV2 *BadgeMinterV2Session) Signer() (common.Address, error) {
	return _BadgeMinterV2.Contract.Signer(&_BadgeMinterV2.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_BadgeMinterV2 *BadgeMinterV2CallerSession) Signer() (common.Address, error) {
	return _BadgeMinterV2.Contract.Signer(&_BadgeMinterV2.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_BadgeMinterV2 *BadgeMinterV2Caller) StartTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BadgeMinterV2.contract.Call(opts, &out, "startTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_BadgeMinterV2 *BadgeMinterV2Session) StartTokenId() (*big.Int, error) {
	return _BadgeMinterV2.Contract.StartTokenId(&_BadgeMinterV2.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_BadgeMinterV2 *BadgeMinterV2CallerSession) StartTokenId() (*big.Int, error) {
	return _BadgeMinterV2.Contract.StartTokenId(&_BadgeMinterV2.CallOpts)
}

// AirdropBadge is a paid mutator transaction binding the contract method 0x613fde91.
//
// Solidity: function airdropBadge(uint256[] questIds, address[] receivers, string[] uris, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2Transactor) AirdropBadge(opts *bind.TransactOpts, questIds []*big.Int, receivers []common.Address, uris []string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.contract.Transact(opts, "airdropBadge", questIds, receivers, uris, signature)
}

// AirdropBadge is a paid mutator transaction binding the contract method 0x613fde91.
//
// Solidity: function airdropBadge(uint256[] questIds, address[] receivers, string[] uris, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2Session) AirdropBadge(questIds []*big.Int, receivers []common.Address, uris []string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.AirdropBadge(&_BadgeMinterV2.TransactOpts, questIds, receivers, uris, signature)
}

// AirdropBadge is a paid mutator transaction binding the contract method 0x613fde91.
//
// Solidity: function airdropBadge(uint256[] questIds, address[] receivers, string[] uris, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2TransactorSession) AirdropBadge(questIds []*big.Int, receivers []common.Address, uris []string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.AirdropBadge(&_BadgeMinterV2.TransactOpts, questIds, receivers, uris, signature)
}

// Claim is a paid mutator transaction binding the contract method 0xdf751012.
//
// Solidity: function claim(address to, uint256 questId, string uri, bytes signature) payable returns()
func (_BadgeMinterV2 *BadgeMinterV2Transactor) Claim(opts *bind.TransactOpts, to common.Address, questId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.contract.Transact(opts, "claim", to, questId, uri, signature)
}

// Claim is a paid mutator transaction binding the contract method 0xdf751012.
//
// Solidity: function claim(address to, uint256 questId, string uri, bytes signature) payable returns()
func (_BadgeMinterV2 *BadgeMinterV2Session) Claim(to common.Address, questId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.Claim(&_BadgeMinterV2.TransactOpts, to, questId, uri, signature)
}

// Claim is a paid mutator transaction binding the contract method 0xdf751012.
//
// Solidity: function claim(address to, uint256 questId, string uri, bytes signature) payable returns()
func (_BadgeMinterV2 *BadgeMinterV2TransactorSession) Claim(to common.Address, questId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.Claim(&_BadgeMinterV2.TransactOpts, to, questId, uri, signature)
}

// ClaimWithInit is a paid mutator transaction binding the contract method 0xed72aaec.
//
// Solidity: function claimWithInit((address,uint32,uint32,string,string) questData, uint256 questId, address to, string uri, bytes signature) payable returns()
func (_BadgeMinterV2 *BadgeMinterV2Transactor) ClaimWithInit(opts *bind.TransactOpts, questData IBadgeQuestData, questId *big.Int, to common.Address, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.contract.Transact(opts, "claimWithInit", questData, questId, to, uri, signature)
}

// ClaimWithInit is a paid mutator transaction binding the contract method 0xed72aaec.
//
// Solidity: function claimWithInit((address,uint32,uint32,string,string) questData, uint256 questId, address to, string uri, bytes signature) payable returns()
func (_BadgeMinterV2 *BadgeMinterV2Session) ClaimWithInit(questData IBadgeQuestData, questId *big.Int, to common.Address, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.ClaimWithInit(&_BadgeMinterV2.TransactOpts, questData, questId, to, uri, signature)
}

// ClaimWithInit is a paid mutator transaction binding the contract method 0xed72aaec.
//
// Solidity: function claimWithInit((address,uint32,uint32,string,string) questData, uint256 questId, address to, string uri, bytes signature) payable returns()
func (_BadgeMinterV2 *BadgeMinterV2TransactorSession) ClaimWithInit(questData IBadgeQuestData, questId *big.Int, to common.Address, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.ClaimWithInit(&_BadgeMinterV2.TransactOpts, questData, questId, to, uri, signature)
}

// InitQuest is a paid mutator transaction binding the contract method 0xd31e6b9a.
//
// Solidity: function initQuest(uint256 questId, (address,uint32,uint32,string,string) questData, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2Transactor) InitQuest(opts *bind.TransactOpts, questId *big.Int, questData IBadgeQuestData, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.contract.Transact(opts, "initQuest", questId, questData, signature)
}

// InitQuest is a paid mutator transaction binding the contract method 0xd31e6b9a.
//
// Solidity: function initQuest(uint256 questId, (address,uint32,uint32,string,string) questData, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2Session) InitQuest(questId *big.Int, questData IBadgeQuestData, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.InitQuest(&_BadgeMinterV2.TransactOpts, questId, questData, signature)
}

// InitQuest is a paid mutator transaction binding the contract method 0xd31e6b9a.
//
// Solidity: function initQuest(uint256 questId, (address,uint32,uint32,string,string) questData, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2TransactorSession) InitQuest(questId *big.Int, questData IBadgeQuestData, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.InitQuest(&_BadgeMinterV2.TransactOpts, questId, questData, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BadgeMinterV2 *BadgeMinterV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BadgeMinterV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BadgeMinterV2 *BadgeMinterV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.RenounceOwnership(&_BadgeMinterV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BadgeMinterV2 *BadgeMinterV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.RenounceOwnership(&_BadgeMinterV2.TransactOpts)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_BadgeMinterV2 *BadgeMinterV2Transactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _BadgeMinterV2.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_BadgeMinterV2 *BadgeMinterV2Session) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.SetSigner(&_BadgeMinterV2.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_BadgeMinterV2 *BadgeMinterV2TransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.SetSigner(&_BadgeMinterV2.TransactOpts, signer_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BadgeMinterV2 *BadgeMinterV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BadgeMinterV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BadgeMinterV2 *BadgeMinterV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.TransferOwnership(&_BadgeMinterV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BadgeMinterV2 *BadgeMinterV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.TransferOwnership(&_BadgeMinterV2.TransactOpts, newOwner)
}

// UpdateQuest is a paid mutator transaction binding the contract method 0xce869d70.
//
// Solidity: function updateQuest(uint256 questId, uint32 startTs, uint32 endTs, string title, string questUri, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2Transactor) UpdateQuest(opts *bind.TransactOpts, questId *big.Int, startTs uint32, endTs uint32, title string, questUri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.contract.Transact(opts, "updateQuest", questId, startTs, endTs, title, questUri, signature)
}

// UpdateQuest is a paid mutator transaction binding the contract method 0xce869d70.
//
// Solidity: function updateQuest(uint256 questId, uint32 startTs, uint32 endTs, string title, string questUri, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2Session) UpdateQuest(questId *big.Int, startTs uint32, endTs uint32, title string, questUri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.UpdateQuest(&_BadgeMinterV2.TransactOpts, questId, startTs, endTs, title, questUri, signature)
}

// UpdateQuest is a paid mutator transaction binding the contract method 0xce869d70.
//
// Solidity: function updateQuest(uint256 questId, uint32 startTs, uint32 endTs, string title, string questUri, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2TransactorSession) UpdateQuest(questId *big.Int, startTs uint32, endTs uint32, title string, questUri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.UpdateQuest(&_BadgeMinterV2.TransactOpts, questId, startTs, endTs, title, questUri, signature)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x45a2c33b.
//
// Solidity: function updateURI(uint256 tokenId, string uri, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2Transactor) UpdateURI(opts *bind.TransactOpts, tokenId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.contract.Transact(opts, "updateURI", tokenId, uri, signature)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x45a2c33b.
//
// Solidity: function updateURI(uint256 tokenId, string uri, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2Session) UpdateURI(tokenId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.UpdateURI(&_BadgeMinterV2.TransactOpts, tokenId, uri, signature)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x45a2c33b.
//
// Solidity: function updateURI(uint256 tokenId, string uri, bytes signature) returns()
func (_BadgeMinterV2 *BadgeMinterV2TransactorSession) UpdateURI(tokenId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinterV2.Contract.UpdateURI(&_BadgeMinterV2.TransactOpts, tokenId, uri, signature)
}

// BadgeMinterV2AirdropedIterator is returned from FilterAirdroped and is used to iterate over the raw logs and unpacked data for Airdroped events raised by the BadgeMinterV2 contract.
type BadgeMinterV2AirdropedIterator struct {
	Event *BadgeMinterV2Airdroped // Event containing the contract specifics and raw log

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
func (it *BadgeMinterV2AirdropedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeMinterV2Airdroped)
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
		it.Event = new(BadgeMinterV2Airdroped)
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
func (it *BadgeMinterV2AirdropedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeMinterV2AirdropedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeMinterV2Airdroped represents a Airdroped event raised by the BadgeMinterV2 contract.
type BadgeMinterV2Airdroped struct {
	QuestId *big.Int
	To      common.Address
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAirdroped is a free log retrieval operation binding the contract event 0x62487ca50a3368fcdd206510d5e64e56d008589216ded84509c941cfeef54d2b.
//
// Solidity: event Airdroped(uint256 indexed questId, address indexed to, string uri)
func (_BadgeMinterV2 *BadgeMinterV2Filterer) FilterAirdroped(opts *bind.FilterOpts, questId []*big.Int, to []common.Address) (*BadgeMinterV2AirdropedIterator, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BadgeMinterV2.contract.FilterLogs(opts, "Airdroped", questIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterV2AirdropedIterator{contract: _BadgeMinterV2.contract, event: "Airdroped", logs: logs, sub: sub}, nil
}

// WatchAirdroped is a free log subscription operation binding the contract event 0x62487ca50a3368fcdd206510d5e64e56d008589216ded84509c941cfeef54d2b.
//
// Solidity: event Airdroped(uint256 indexed questId, address indexed to, string uri)
func (_BadgeMinterV2 *BadgeMinterV2Filterer) WatchAirdroped(opts *bind.WatchOpts, sink chan<- *BadgeMinterV2Airdroped, questId []*big.Int, to []common.Address) (event.Subscription, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BadgeMinterV2.contract.WatchLogs(opts, "Airdroped", questIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeMinterV2Airdroped)
				if err := _BadgeMinterV2.contract.UnpackLog(event, "Airdroped", log); err != nil {
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

// ParseAirdroped is a log parse operation binding the contract event 0x62487ca50a3368fcdd206510d5e64e56d008589216ded84509c941cfeef54d2b.
//
// Solidity: event Airdroped(uint256 indexed questId, address indexed to, string uri)
func (_BadgeMinterV2 *BadgeMinterV2Filterer) ParseAirdroped(log types.Log) (*BadgeMinterV2Airdroped, error) {
	event := new(BadgeMinterV2Airdroped)
	if err := _BadgeMinterV2.contract.UnpackLog(event, "Airdroped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeMinterV2DonationIterator is returned from FilterDonation and is used to iterate over the raw logs and unpacked data for Donation events raised by the BadgeMinterV2 contract.
type BadgeMinterV2DonationIterator struct {
	Event *BadgeMinterV2Donation // Event containing the contract specifics and raw log

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
func (it *BadgeMinterV2DonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeMinterV2Donation)
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
		it.Event = new(BadgeMinterV2Donation)
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
func (it *BadgeMinterV2DonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeMinterV2DonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeMinterV2Donation represents a Donation event raised by the BadgeMinterV2 contract.
type BadgeMinterV2Donation struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonation is a free log retrieval operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_BadgeMinterV2 *BadgeMinterV2Filterer) FilterDonation(opts *bind.FilterOpts) (*BadgeMinterV2DonationIterator, error) {

	logs, sub, err := _BadgeMinterV2.contract.FilterLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return &BadgeMinterV2DonationIterator{contract: _BadgeMinterV2.contract, event: "Donation", logs: logs, sub: sub}, nil
}

// WatchDonation is a free log subscription operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_BadgeMinterV2 *BadgeMinterV2Filterer) WatchDonation(opts *bind.WatchOpts, sink chan<- *BadgeMinterV2Donation) (event.Subscription, error) {

	logs, sub, err := _BadgeMinterV2.contract.WatchLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeMinterV2Donation)
				if err := _BadgeMinterV2.contract.UnpackLog(event, "Donation", log); err != nil {
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
func (_BadgeMinterV2 *BadgeMinterV2Filterer) ParseDonation(log types.Log) (*BadgeMinterV2Donation, error) {
	event := new(BadgeMinterV2Donation)
	if err := _BadgeMinterV2.contract.UnpackLog(event, "Donation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeMinterV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BadgeMinterV2 contract.
type BadgeMinterV2OwnershipTransferredIterator struct {
	Event *BadgeMinterV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BadgeMinterV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeMinterV2OwnershipTransferred)
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
		it.Event = new(BadgeMinterV2OwnershipTransferred)
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
func (it *BadgeMinterV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeMinterV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeMinterV2OwnershipTransferred represents a OwnershipTransferred event raised by the BadgeMinterV2 contract.
type BadgeMinterV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BadgeMinterV2 *BadgeMinterV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BadgeMinterV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BadgeMinterV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterV2OwnershipTransferredIterator{contract: _BadgeMinterV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BadgeMinterV2 *BadgeMinterV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BadgeMinterV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BadgeMinterV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeMinterV2OwnershipTransferred)
				if err := _BadgeMinterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BadgeMinterV2 *BadgeMinterV2Filterer) ParseOwnershipTransferred(log types.Log) (*BadgeMinterV2OwnershipTransferred, error) {
	event := new(BadgeMinterV2OwnershipTransferred)
	if err := _BadgeMinterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeMinterV2SignerChangedIterator is returned from FilterSignerChanged and is used to iterate over the raw logs and unpacked data for SignerChanged events raised by the BadgeMinterV2 contract.
type BadgeMinterV2SignerChangedIterator struct {
	Event *BadgeMinterV2SignerChanged // Event containing the contract specifics and raw log

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
func (it *BadgeMinterV2SignerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeMinterV2SignerChanged)
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
		it.Event = new(BadgeMinterV2SignerChanged)
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
func (it *BadgeMinterV2SignerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeMinterV2SignerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeMinterV2SignerChanged represents a SignerChanged event raised by the BadgeMinterV2 contract.
type BadgeMinterV2SignerChanged struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerChanged is a free log retrieval operation binding the contract event 0x5719a5656c5cfdaafa148ecf366fd3b0a7fae06449ce2a46225977fb7417e29d.
//
// Solidity: event SignerChanged(address signer)
func (_BadgeMinterV2 *BadgeMinterV2Filterer) FilterSignerChanged(opts *bind.FilterOpts) (*BadgeMinterV2SignerChangedIterator, error) {

	logs, sub, err := _BadgeMinterV2.contract.FilterLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return &BadgeMinterV2SignerChangedIterator{contract: _BadgeMinterV2.contract, event: "SignerChanged", logs: logs, sub: sub}, nil
}

// WatchSignerChanged is a free log subscription operation binding the contract event 0x5719a5656c5cfdaafa148ecf366fd3b0a7fae06449ce2a46225977fb7417e29d.
//
// Solidity: event SignerChanged(address signer)
func (_BadgeMinterV2 *BadgeMinterV2Filterer) WatchSignerChanged(opts *bind.WatchOpts, sink chan<- *BadgeMinterV2SignerChanged) (event.Subscription, error) {

	logs, sub, err := _BadgeMinterV2.contract.WatchLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeMinterV2SignerChanged)
				if err := _BadgeMinterV2.contract.UnpackLog(event, "SignerChanged", log); err != nil {
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
func (_BadgeMinterV2 *BadgeMinterV2Filterer) ParseSignerChanged(log types.Log) (*BadgeMinterV2SignerChanged, error) {
	event := new(BadgeMinterV2SignerChanged)
	if err := _BadgeMinterV2.contract.UnpackLog(event, "SignerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
