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

// BadgeMinterMetaData contains all meta data concerning the BadgeMinter contract.
var BadgeMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"badge_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCreator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"Airdroped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Donation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"questIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"uris\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"airdropBadge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"badge\",\"outputs\":[{\"internalType\":\"contractIBadge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"questId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer_\",\"type\":\"address\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BadgeMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use BadgeMinterMetaData.ABI instead.
var BadgeMinterABI = BadgeMinterMetaData.ABI

// BadgeMinter is an auto generated Go binding around an Ethereum contract.
type BadgeMinter struct {
	BadgeMinterCaller     // Read-only binding to the contract
	BadgeMinterTransactor // Write-only binding to the contract
	BadgeMinterFilterer   // Log filterer for contract events
}

// BadgeMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BadgeMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BadgeMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BadgeMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BadgeMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BadgeMinterSession struct {
	Contract     *BadgeMinter      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BadgeMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BadgeMinterCallerSession struct {
	Contract *BadgeMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BadgeMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BadgeMinterTransactorSession struct {
	Contract     *BadgeMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BadgeMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BadgeMinterRaw struct {
	Contract *BadgeMinter // Generic contract binding to access the raw methods on
}

// BadgeMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BadgeMinterCallerRaw struct {
	Contract *BadgeMinterCaller // Generic read-only contract binding to access the raw methods on
}

// BadgeMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BadgeMinterTransactorRaw struct {
	Contract *BadgeMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBadgeMinter creates a new instance of BadgeMinter, bound to a specific deployed contract.
func NewBadgeMinter(address common.Address, backend bind.ContractBackend) (*BadgeMinter, error) {
	contract, err := bindBadgeMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BadgeMinter{BadgeMinterCaller: BadgeMinterCaller{contract: contract}, BadgeMinterTransactor: BadgeMinterTransactor{contract: contract}, BadgeMinterFilterer: BadgeMinterFilterer{contract: contract}}, nil
}

// NewBadgeMinterCaller creates a new read-only instance of BadgeMinter, bound to a specific deployed contract.
func NewBadgeMinterCaller(address common.Address, caller bind.ContractCaller) (*BadgeMinterCaller, error) {
	contract, err := bindBadgeMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterCaller{contract: contract}, nil
}

// NewBadgeMinterTransactor creates a new write-only instance of BadgeMinter, bound to a specific deployed contract.
func NewBadgeMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*BadgeMinterTransactor, error) {
	contract, err := bindBadgeMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterTransactor{contract: contract}, nil
}

// NewBadgeMinterFilterer creates a new log filterer instance of BadgeMinter, bound to a specific deployed contract.
func NewBadgeMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*BadgeMinterFilterer, error) {
	contract, err := bindBadgeMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterFilterer{contract: contract}, nil
}

// bindBadgeMinter binds a generic wrapper to an already deployed contract.
func bindBadgeMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BadgeMinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BadgeMinter *BadgeMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BadgeMinter.Contract.BadgeMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BadgeMinter *BadgeMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BadgeMinter.Contract.BadgeMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BadgeMinter *BadgeMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BadgeMinter.Contract.BadgeMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BadgeMinter *BadgeMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BadgeMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BadgeMinter *BadgeMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BadgeMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BadgeMinter *BadgeMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BadgeMinter.Contract.contract.Transact(opts, method, params...)
}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_BadgeMinter *BadgeMinterCaller) Badge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BadgeMinter.contract.Call(opts, &out, "badge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_BadgeMinter *BadgeMinterSession) Badge() (common.Address, error) {
	return _BadgeMinter.Contract.Badge(&_BadgeMinter.CallOpts)
}

// Badge is a free data retrieval call binding the contract method 0x91d768de.
//
// Solidity: function badge() view returns(address)
func (_BadgeMinter *BadgeMinterCallerSession) Badge() (common.Address, error) {
	return _BadgeMinter.Contract.Badge(&_BadgeMinter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BadgeMinter *BadgeMinterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BadgeMinter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BadgeMinter *BadgeMinterSession) Owner() (common.Address, error) {
	return _BadgeMinter.Contract.Owner(&_BadgeMinter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BadgeMinter *BadgeMinterCallerSession) Owner() (common.Address, error) {
	return _BadgeMinter.Contract.Owner(&_BadgeMinter.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_BadgeMinter *BadgeMinterCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BadgeMinter.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_BadgeMinter *BadgeMinterSession) Signer() (common.Address, error) {
	return _BadgeMinter.Contract.Signer(&_BadgeMinter.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_BadgeMinter *BadgeMinterCallerSession) Signer() (common.Address, error) {
	return _BadgeMinter.Contract.Signer(&_BadgeMinter.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_BadgeMinter *BadgeMinterCaller) StartTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BadgeMinter.contract.Call(opts, &out, "startTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_BadgeMinter *BadgeMinterSession) StartTokenId() (*big.Int, error) {
	return _BadgeMinter.Contract.StartTokenId(&_BadgeMinter.CallOpts)
}

// StartTokenId is a free data retrieval call binding the contract method 0xe6798baa.
//
// Solidity: function startTokenId() view returns(uint256)
func (_BadgeMinter *BadgeMinterCallerSession) StartTokenId() (*big.Int, error) {
	return _BadgeMinter.Contract.StartTokenId(&_BadgeMinter.CallOpts)
}

// AirdropBadge is a paid mutator transaction binding the contract method 0x613fde91.
//
// Solidity: function airdropBadge(uint256[] questIds, address[] receivers, string[] uris, bytes signature) returns()
func (_BadgeMinter *BadgeMinterTransactor) AirdropBadge(opts *bind.TransactOpts, questIds []*big.Int, receivers []common.Address, uris []string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinter.contract.Transact(opts, "airdropBadge", questIds, receivers, uris, signature)
}

// AirdropBadge is a paid mutator transaction binding the contract method 0x613fde91.
//
// Solidity: function airdropBadge(uint256[] questIds, address[] receivers, string[] uris, bytes signature) returns()
func (_BadgeMinter *BadgeMinterSession) AirdropBadge(questIds []*big.Int, receivers []common.Address, uris []string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinter.Contract.AirdropBadge(&_BadgeMinter.TransactOpts, questIds, receivers, uris, signature)
}

// AirdropBadge is a paid mutator transaction binding the contract method 0x613fde91.
//
// Solidity: function airdropBadge(uint256[] questIds, address[] receivers, string[] uris, bytes signature) returns()
func (_BadgeMinter *BadgeMinterTransactorSession) AirdropBadge(questIds []*big.Int, receivers []common.Address, uris []string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinter.Contract.AirdropBadge(&_BadgeMinter.TransactOpts, questIds, receivers, uris, signature)
}

// Claim is a paid mutator transaction binding the contract method 0xdf751012.
//
// Solidity: function claim(address to, uint256 questId, string uri, bytes signature) payable returns()
func (_BadgeMinter *BadgeMinterTransactor) Claim(opts *bind.TransactOpts, to common.Address, questId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinter.contract.Transact(opts, "claim", to, questId, uri, signature)
}

// Claim is a paid mutator transaction binding the contract method 0xdf751012.
//
// Solidity: function claim(address to, uint256 questId, string uri, bytes signature) payable returns()
func (_BadgeMinter *BadgeMinterSession) Claim(to common.Address, questId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinter.Contract.Claim(&_BadgeMinter.TransactOpts, to, questId, uri, signature)
}

// Claim is a paid mutator transaction binding the contract method 0xdf751012.
//
// Solidity: function claim(address to, uint256 questId, string uri, bytes signature) payable returns()
func (_BadgeMinter *BadgeMinterTransactorSession) Claim(to common.Address, questId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinter.Contract.Claim(&_BadgeMinter.TransactOpts, to, questId, uri, signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BadgeMinter *BadgeMinterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BadgeMinter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BadgeMinter *BadgeMinterSession) RenounceOwnership() (*types.Transaction, error) {
	return _BadgeMinter.Contract.RenounceOwnership(&_BadgeMinter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BadgeMinter *BadgeMinterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BadgeMinter.Contract.RenounceOwnership(&_BadgeMinter.TransactOpts)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_BadgeMinter *BadgeMinterTransactor) SetSigner(opts *bind.TransactOpts, signer_ common.Address) (*types.Transaction, error) {
	return _BadgeMinter.contract.Transact(opts, "setSigner", signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_BadgeMinter *BadgeMinterSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _BadgeMinter.Contract.SetSigner(&_BadgeMinter.TransactOpts, signer_)
}

// SetSigner is a paid mutator transaction binding the contract method 0x6c19e783.
//
// Solidity: function setSigner(address signer_) returns()
func (_BadgeMinter *BadgeMinterTransactorSession) SetSigner(signer_ common.Address) (*types.Transaction, error) {
	return _BadgeMinter.Contract.SetSigner(&_BadgeMinter.TransactOpts, signer_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BadgeMinter *BadgeMinterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BadgeMinter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BadgeMinter *BadgeMinterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BadgeMinter.Contract.TransferOwnership(&_BadgeMinter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BadgeMinter *BadgeMinterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BadgeMinter.Contract.TransferOwnership(&_BadgeMinter.TransactOpts, newOwner)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x45a2c33b.
//
// Solidity: function updateURI(uint256 tokenId, string uri, bytes signature) returns()
func (_BadgeMinter *BadgeMinterTransactor) UpdateURI(opts *bind.TransactOpts, tokenId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinter.contract.Transact(opts, "updateURI", tokenId, uri, signature)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x45a2c33b.
//
// Solidity: function updateURI(uint256 tokenId, string uri, bytes signature) returns()
func (_BadgeMinter *BadgeMinterSession) UpdateURI(tokenId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinter.Contract.UpdateURI(&_BadgeMinter.TransactOpts, tokenId, uri, signature)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x45a2c33b.
//
// Solidity: function updateURI(uint256 tokenId, string uri, bytes signature) returns()
func (_BadgeMinter *BadgeMinterTransactorSession) UpdateURI(tokenId *big.Int, uri string, signature []byte) (*types.Transaction, error) {
	return _BadgeMinter.Contract.UpdateURI(&_BadgeMinter.TransactOpts, tokenId, uri, signature)
}

// BadgeMinterAirdropedIterator is returned from FilterAirdroped and is used to iterate over the raw logs and unpacked data for Airdroped events raised by the BadgeMinter contract.
type BadgeMinterAirdropedIterator struct {
	Event *BadgeMinterAirdroped // Event containing the contract specifics and raw log

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
func (it *BadgeMinterAirdropedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeMinterAirdroped)
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
		it.Event = new(BadgeMinterAirdroped)
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
func (it *BadgeMinterAirdropedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeMinterAirdropedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeMinterAirdroped represents a Airdroped event raised by the BadgeMinter contract.
type BadgeMinterAirdroped struct {
	QuestId *big.Int
	To      common.Address
	Uri     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAirdroped is a free log retrieval operation binding the contract event 0x62487ca50a3368fcdd206510d5e64e56d008589216ded84509c941cfeef54d2b.
//
// Solidity: event Airdroped(uint256 indexed questId, address indexed to, string uri)
func (_BadgeMinter *BadgeMinterFilterer) FilterAirdroped(opts *bind.FilterOpts, questId []*big.Int, to []common.Address) (*BadgeMinterAirdropedIterator, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BadgeMinter.contract.FilterLogs(opts, "Airdroped", questIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterAirdropedIterator{contract: _BadgeMinter.contract, event: "Airdroped", logs: logs, sub: sub}, nil
}

// WatchAirdroped is a free log subscription operation binding the contract event 0x62487ca50a3368fcdd206510d5e64e56d008589216ded84509c941cfeef54d2b.
//
// Solidity: event Airdroped(uint256 indexed questId, address indexed to, string uri)
func (_BadgeMinter *BadgeMinterFilterer) WatchAirdroped(opts *bind.WatchOpts, sink chan<- *BadgeMinterAirdroped, questId []*big.Int, to []common.Address) (event.Subscription, error) {

	var questIdRule []interface{}
	for _, questIdItem := range questId {
		questIdRule = append(questIdRule, questIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BadgeMinter.contract.WatchLogs(opts, "Airdroped", questIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeMinterAirdroped)
				if err := _BadgeMinter.contract.UnpackLog(event, "Airdroped", log); err != nil {
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
func (_BadgeMinter *BadgeMinterFilterer) ParseAirdroped(log types.Log) (*BadgeMinterAirdroped, error) {
	event := new(BadgeMinterAirdroped)
	if err := _BadgeMinter.contract.UnpackLog(event, "Airdroped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeMinterDonationIterator is returned from FilterDonation and is used to iterate over the raw logs and unpacked data for Donation events raised by the BadgeMinter contract.
type BadgeMinterDonationIterator struct {
	Event *BadgeMinterDonation // Event containing the contract specifics and raw log

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
func (it *BadgeMinterDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeMinterDonation)
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
		it.Event = new(BadgeMinterDonation)
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
func (it *BadgeMinterDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeMinterDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeMinterDonation represents a Donation event raised by the BadgeMinter contract.
type BadgeMinterDonation struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonation is a free log retrieval operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_BadgeMinter *BadgeMinterFilterer) FilterDonation(opts *bind.FilterOpts) (*BadgeMinterDonationIterator, error) {

	logs, sub, err := _BadgeMinter.contract.FilterLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return &BadgeMinterDonationIterator{contract: _BadgeMinter.contract, event: "Donation", logs: logs, sub: sub}, nil
}

// WatchDonation is a free log subscription operation binding the contract event 0x597440e65d8cdea1298e19df50e115bda25508180452d87a7f4f99195fb843a4.
//
// Solidity: event Donation(address from, address to, uint256 amount)
func (_BadgeMinter *BadgeMinterFilterer) WatchDonation(opts *bind.WatchOpts, sink chan<- *BadgeMinterDonation) (event.Subscription, error) {

	logs, sub, err := _BadgeMinter.contract.WatchLogs(opts, "Donation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeMinterDonation)
				if err := _BadgeMinter.contract.UnpackLog(event, "Donation", log); err != nil {
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
func (_BadgeMinter *BadgeMinterFilterer) ParseDonation(log types.Log) (*BadgeMinterDonation, error) {
	event := new(BadgeMinterDonation)
	if err := _BadgeMinter.contract.UnpackLog(event, "Donation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeMinterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BadgeMinter contract.
type BadgeMinterOwnershipTransferredIterator struct {
	Event *BadgeMinterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BadgeMinterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeMinterOwnershipTransferred)
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
		it.Event = new(BadgeMinterOwnershipTransferred)
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
func (it *BadgeMinterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeMinterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeMinterOwnershipTransferred represents a OwnershipTransferred event raised by the BadgeMinter contract.
type BadgeMinterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BadgeMinter *BadgeMinterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BadgeMinterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BadgeMinter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BadgeMinterOwnershipTransferredIterator{contract: _BadgeMinter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BadgeMinter *BadgeMinterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BadgeMinterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BadgeMinter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeMinterOwnershipTransferred)
				if err := _BadgeMinter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BadgeMinter *BadgeMinterFilterer) ParseOwnershipTransferred(log types.Log) (*BadgeMinterOwnershipTransferred, error) {
	event := new(BadgeMinterOwnershipTransferred)
	if err := _BadgeMinter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BadgeMinterSignerChangedIterator is returned from FilterSignerChanged and is used to iterate over the raw logs and unpacked data for SignerChanged events raised by the BadgeMinter contract.
type BadgeMinterSignerChangedIterator struct {
	Event *BadgeMinterSignerChanged // Event containing the contract specifics and raw log

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
func (it *BadgeMinterSignerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BadgeMinterSignerChanged)
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
		it.Event = new(BadgeMinterSignerChanged)
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
func (it *BadgeMinterSignerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BadgeMinterSignerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BadgeMinterSignerChanged represents a SignerChanged event raised by the BadgeMinter contract.
type BadgeMinterSignerChanged struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerChanged is a free log retrieval operation binding the contract event 0x5719a5656c5cfdaafa148ecf366fd3b0a7fae06449ce2a46225977fb7417e29d.
//
// Solidity: event SignerChanged(address signer)
func (_BadgeMinter *BadgeMinterFilterer) FilterSignerChanged(opts *bind.FilterOpts) (*BadgeMinterSignerChangedIterator, error) {

	logs, sub, err := _BadgeMinter.contract.FilterLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return &BadgeMinterSignerChangedIterator{contract: _BadgeMinter.contract, event: "SignerChanged", logs: logs, sub: sub}, nil
}

// WatchSignerChanged is a free log subscription operation binding the contract event 0x5719a5656c5cfdaafa148ecf366fd3b0a7fae06449ce2a46225977fb7417e29d.
//
// Solidity: event SignerChanged(address signer)
func (_BadgeMinter *BadgeMinterFilterer) WatchSignerChanged(opts *bind.WatchOpts, sink chan<- *BadgeMinterSignerChanged) (event.Subscription, error) {

	logs, sub, err := _BadgeMinter.contract.WatchLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BadgeMinterSignerChanged)
				if err := _BadgeMinter.contract.UnpackLog(event, "SignerChanged", log); err != nil {
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
func (_BadgeMinter *BadgeMinterFilterer) ParseSignerChanged(log types.Log) (*BadgeMinterSignerChanged, error) {
	event := new(BadgeMinterSignerChanged)
	if err := _BadgeMinter.contract.UnpackLog(event, "SignerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
