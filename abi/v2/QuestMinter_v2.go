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


// QuestMinterV2MetaData contains all meta data concerning the QuestMinterV2 contract.
var QuestMinterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"quest_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCreator\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerChanged\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"createQuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"startTs\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endTs\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"internalType\":\"structIQuest.QuestData\",\"name\":\"questData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"modifyQuest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quest\",\"outputs\":[{\"internalType\":\"contractIQuest\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"badgeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateBadgeNum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QuestMinterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use QuestMinterV2MetaData.ABI instead.
var QuestMinterV2ABI = QuestMinterV2MetaData.ABI

// QuestMinterV2 is an auto generated Go binding around an Ethereum contract.
type QuestMinterV2 struct {
	QuestMinterV2Caller     // Read-only binding to the contract
	QuestMinterV2Transactor // Write-only binding to the contract
	QuestMinterV2Filterer   // Log filterer for contract events
}

// QuestMinterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type QuestMinterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestMinterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type QuestMinterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestMinterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuestMinterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuestMinterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuestMinterV2Session struct {
	Contract     *QuestMinterV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuestMinterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuestMinterV2CallerSession struct {
	Contract *QuestMinterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// QuestMinterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuestMinterV2TransactorSession struct {
	Contract     *QuestMinterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// QuestMinterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type QuestMinterV2Raw struct {
	Contract *QuestMinterV2 // Generic contract binding to access the raw methods on
}

// QuestMinterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuestMinterV2CallerRaw struct {
	Contract *QuestMinterV2Caller // Generic read-only contract binding to access the raw methods on
}

// QuestMinterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuestMinterV2TransactorRaw struct {
	Contract *QuestMinterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewQuestMinterV2 creates a new instance of QuestMinterV2, bound to a specific deployed contract.
func NewQuestMinterV2(address common.Address, backend bind.ContractBackend) (*QuestMinterV2, error) {
	contract, err := bindQuestMinterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuestMinterV2{QuestMinterV2Caller: QuestMinterV2Caller{contract: contract}, QuestMinterV2Transactor: QuestMinterV2Transactor{contract: contract}, QuestMinterV2Filterer: QuestMinterV2Filterer{contract: contract}}, nil
}

// NewQuestMinterV2Caller creates a new read-only instance of QuestMinterV2, bound to a specific deployed contract.
func NewQuestMinterV2Caller(address common.Address, caller bind.ContractCaller) (*QuestMinterV2Caller, error) {
	contract, err := bindQuestMinterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuestMinterV2Caller{contract: contract}, nil
}

// NewQuestMinterV2Transactor creates a new write-only instance of QuestMinterV2, bound to a specific deployed contract.
func NewQuestMinterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*QuestMinterV2Transactor, error) {
	contract, err := bindQuestMinterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuestMinterV2Transactor{contract: contract}, nil
}

// NewQuestMinterV2Filterer creates a new log filterer instance of QuestMinterV2, bound to a specific deployed contract.
func NewQuestMinterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*QuestMinterV2Filterer, error) {
	contract, err := bindQuestMinterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuestMinterV2Filterer{contract: contract}, nil
}

// bindQuestMinterV2 binds a generic wrapper to an already deployed contract.
func bindQuestMinterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuestMinterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestMinterV2 *QuestMinterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestMinterV2.Contract.QuestMinterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestMinterV2 *QuestMinterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.QuestMinterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestMinterV2 *QuestMinterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.QuestMinterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuestMinterV2 *QuestMinterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuestMinterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuestMinterV2 *QuestMinterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuestMinterV2 *QuestMinterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestMinterV2 *QuestMinterV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestMinterV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestMinterV2 *QuestMinterV2Session) Owner() (common.Address, error) {
	return _QuestMinterV2.Contract.Owner(&_QuestMinterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_QuestMinterV2 *QuestMinterV2CallerSession) Owner() (common.Address, error) {
	return _QuestMinterV2.Contract.Owner(&_QuestMinterV2.CallOpts)
}

// Quest is a free data retrieval call binding the contract method 0x42cccf08.
//
// Solidity: function quest() view returns(address)
func (_QuestMinterV2 *QuestMinterV2Caller) Quest(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestMinterV2.contract.Call(opts, &out, "quest")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Quest is a free data retrieval call binding the contract method 0x42cccf08.
//
// Solidity: function quest() view returns(address)
func (_QuestMinterV2 *QuestMinterV2Session) Quest() (common.Address, error) {
	return _QuestMinterV2.Contract.Quest(&_QuestMinterV2.CallOpts)
}

// Quest is a free data retrieval call binding the contract method 0x42cccf08.
//
// Solidity: function quest() view returns(address)
func (_QuestMinterV2 *QuestMinterV2CallerSession) Quest() (common.Address, error) {
	return _QuestMinterV2.Contract.Quest(&_QuestMinterV2.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_QuestMinterV2 *QuestMinterV2Caller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _QuestMinterV2.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_QuestMinterV2 *QuestMinterV2Session) Signer() (common.Address, error) {
	return _QuestMinterV2.Contract.Signer(&_QuestMinterV2.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_QuestMinterV2 *QuestMinterV2CallerSession) Signer() (common.Address, error) {
	return _QuestMinterV2.Contract.Signer(&_QuestMinterV2.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_QuestMinterV2 *QuestMinterV2Caller) StartTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _QuestMinterV2.contract.Call(opts, &out, "startTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_QuestMinterV2 *QuestMinterV2Session) StartTokenId() (*big.Int, error) {
	return _QuestMinterV2.Contract.StartTokenId(&_QuestMinterV2.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_QuestMinterV2 *QuestMinterV2CallerSession) StartTokenId() (*big.Int, error) {
	return _QuestMinterV2.Contract.StartTokenId(&_QuestMinterV2.CallOpts)
}

// CreateQuest is a paid mutator transaction binding the contract method 0x7a1d3f44.
//
// Solidity: function createQuest((uint32,uint32,string,string) questData, bytes signature) returns()
func (_QuestMinterV2 *QuestMinterV2Transactor) CreateQuest(opts *bind.TransactOpts, questData IQuestQuestData, signature []byte) (*types.Transaction, error) {
	return _QuestMinterV2.contract.Transact(opts, "createQuest", questData, signature)
}

// CreateQuest is a paid mutator transaction binding the contract method 0x7a1d3f44.
//
// Solidity: function createQuest((uint32,uint32,string,string) questData, bytes signature) returns()
func (_QuestMinterV2 *QuestMinterV2Session) CreateQuest(questData IQuestQuestData, signature []byte) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.CreateQuest(&_QuestMinterV2.TransactOpts, questData, signature)
}

// CreateQuest is a paid mutator transaction binding the contract method 0x7a1d3f44.
//
// Solidity: function createQuest((uint32,uint32,string,string) questData, bytes signature) returns()
func (_QuestMinterV2 *QuestMinterV2TransactorSession) CreateQuest(questData IQuestQuestData, signature []byte) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.CreateQuest(&_QuestMinterV2.TransactOpts, questData, signature)
}

// ModifyQuest is a paid mutator transaction binding the contract method 0xec093d64.
//
// Solidity: function modifyQuest(uint256 tokenId, (uint32,uint32,string,string) questData, bytes signature) returns()
func (_QuestMinterV2 *QuestMinterV2Transactor) ModifyQuest(opts *bind.TransactOpts, tokenId *big.Int, questData IQuestQuestData, signature []byte) (*types.Transaction, error) {
	return _QuestMinterV2.contract.Transact(opts, "modifyQuest", tokenId, questData, signature)
}

// ModifyQuest is a paid mutator transaction binding the contract method 0xec093d64.
//
// Solidity: function modifyQuest(uint256 tokenId, (uint32,uint32,string,string) questData, bytes signature) returns()
func (_QuestMinterV2 *QuestMinterV2Session) ModifyQuest(tokenId *big.Int, questData IQuestQuestData, signature []byte) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.ModifyQuest(&_QuestMinterV2.TransactOpts, tokenId, questData, signature)
}

// ModifyQuest is a paid mutator transaction binding the contract method 0xec093d64.
//
// Solidity: function modifyQuest(uint256 tokenId, (uint32,uint32,string,string) questData, bytes signature) returns()
func (_QuestMinterV2 *QuestMinterV2TransactorSession) ModifyQuest(tokenId *big.Int, questData IQuestQuestData, signature []byte) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.ModifyQuest(&_QuestMinterV2.TransactOpts, tokenId, questData, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestMinterV2 *QuestMinterV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuestMinterV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestMinterV2 *QuestMinterV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _QuestMinterV2.Contract.RenounceOwnership(&_QuestMinterV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_QuestMinterV2 *QuestMinterV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _QuestMinterV2.Contract.RenounceOwnership(&_QuestMinterV2.TransactOpts)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_QuestMinterV2 *QuestMinterV2Transactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _QuestMinterV2.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_QuestMinterV2 *QuestMinterV2Session) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.SetSigner(&_QuestMinterV2.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_QuestMinterV2 *QuestMinterV2TransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.SetSigner(&_QuestMinterV2.TransactOpts, signer_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestMinterV2 *QuestMinterV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _QuestMinterV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestMinterV2 *QuestMinterV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.TransferOwnership(&_QuestMinterV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_QuestMinterV2 *QuestMinterV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.TransferOwnership(&_QuestMinterV2.TransactOpts, newOwner)
}

// UpdateBadgeNum is a paid mutator transaction binding the contract method 0x1f043e22.
//
// Solidity: function updateBadgeNum(uint256 questId, uint256 badgeNum, bytes signature) returns()
func (_QuestMinterV2 *QuestMinterV2Transactor) UpdateBadgeNum(opts *bind.TransactOpts, questId *big.Int, badgeNum *big.Int, signature []byte) (*types.Transaction, error) {
	return _QuestMinterV2.contract.Transact(opts, "updateBadgeNum", questId, badgeNum, signature)
}

// UpdateBadgeNum is a paid mutator transaction binding the contract method 0x1f043e22.
//
// Solidity: function updateBadgeNum(uint256 questId, uint256 badgeNum, bytes signature) returns()
func (_QuestMinterV2 *QuestMinterV2Session) UpdateBadgeNum(questId *big.Int, badgeNum *big.Int, signature []byte) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.UpdateBadgeNum(&_QuestMinterV2.TransactOpts, questId, badgeNum, signature)
}

// UpdateBadgeNum is a paid mutator transaction binding the contract method 0x1f043e22.
//
// Solidity: function updateBadgeNum(uint256 questId, uint256 badgeNum, bytes signature) returns()
func (_QuestMinterV2 *QuestMinterV2TransactorSession) UpdateBadgeNum(questId *big.Int, badgeNum *big.Int, signature []byte) (*types.Transaction, error) {
	return _QuestMinterV2.Contract.UpdateBadgeNum(&_QuestMinterV2.TransactOpts, questId, badgeNum, signature)
}

// QuestMinterV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the QuestMinterV2 contract.
type QuestMinterV2OwnershipTransferredIterator struct {
	Event *QuestMinterV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QuestMinterV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestMinterV2OwnershipTransferred)
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
		it.Event = new(QuestMinterV2OwnershipTransferred)
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
func (it *QuestMinterV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestMinterV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestMinterV2OwnershipTransferred represents a OwnershipTransferred event raised by the QuestMinterV2 contract.
type QuestMinterV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuestMinterV2 *QuestMinterV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QuestMinterV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuestMinterV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuestMinterV2OwnershipTransferredIterator{contract: _QuestMinterV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_QuestMinterV2 *QuestMinterV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QuestMinterV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _QuestMinterV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestMinterV2OwnershipTransferred)
				if err := _QuestMinterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_QuestMinterV2 *QuestMinterV2Filterer) ParseOwnershipTransferred(log types.Log) (*QuestMinterV2OwnershipTransferred, error) {
	event := new(QuestMinterV2OwnershipTransferred)
	if err := _QuestMinterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuestMinterV2SignerChangedIterator is returned from FilterSignerChanged and is used to iterate over the raw logs and unpacked data for SignerChanged events raised by the QuestMinterV2 contract.
type QuestMinterV2SignerChangedIterator struct {
	Event *QuestMinterV2SignerChanged // Event containing the contract specifics and raw log

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
func (it *QuestMinterV2SignerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuestMinterV2SignerChanged)
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
		it.Event = new(QuestMinterV2SignerChanged)
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
func (it *QuestMinterV2SignerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuestMinterV2SignerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuestMinterV2SignerChanged represents a SignerChanged event raised by the QuestMinterV2 contract.
type QuestMinterV2SignerChanged struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerChanged is a free log retrieval operation binding the contract event 0x5719a5656c5cfdaafa148ecf366fd3b0a7fae06449ce2a46225977fb7417e29d.
//
// Solidity: event SignerChanged(address signer)
func (_QuestMinterV2 *QuestMinterV2Filterer) FilterSignerChanged(opts *bind.FilterOpts) (*QuestMinterV2SignerChangedIterator, error) {

	logs, sub, err := _QuestMinterV2.contract.FilterLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return &QuestMinterV2SignerChangedIterator{contract: _QuestMinterV2.contract, event: "SignerChanged", logs: logs, sub: sub}, nil
}

// WatchSignerChanged is a free log subscription operation binding the contract event 0x5719a5656c5cfdaafa148ecf366fd3b0a7fae06449ce2a46225977fb7417e29d.
//
// Solidity: event SignerChanged(address signer)
func (_QuestMinterV2 *QuestMinterV2Filterer) WatchSignerChanged(opts *bind.WatchOpts, sink chan<- *QuestMinterV2SignerChanged) (event.Subscription, error) {

	logs, sub, err := _QuestMinterV2.contract.WatchLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuestMinterV2SignerChanged)
				if err := _QuestMinterV2.contract.UnpackLog(event, "SignerChanged", log); err != nil {
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
func (_QuestMinterV2 *QuestMinterV2Filterer) ParseSignerChanged(log types.Log) (*QuestMinterV2SignerChanged, error) {
	event := new(QuestMinterV2SignerChanged)
	if err := _QuestMinterV2.contract.UnpackLog(event, "SignerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
