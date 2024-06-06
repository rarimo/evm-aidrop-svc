// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ERC20TransferWithPermitMetaData contains all meta data concerning the ERC20TransferWithPermit contract.
var ERC20TransferWithPermitMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v_\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s_\",\"type\":\"bytes32\"}],\"name\":\"transferWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC20TransferWithPermitABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20TransferWithPermitMetaData.ABI instead.
var ERC20TransferWithPermitABI = ERC20TransferWithPermitMetaData.ABI

// ERC20TransferWithPermit is an auto generated Go binding around an Ethereum contract.
type ERC20TransferWithPermit struct {
	ERC20TransferWithPermitCaller     // Read-only binding to the contract
	ERC20TransferWithPermitTransactor // Write-only binding to the contract
	ERC20TransferWithPermitFilterer   // Log filterer for contract events
}

// ERC20TransferWithPermitCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TransferWithPermitCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TransferWithPermitTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TransferWithPermitTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TransferWithPermitFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TransferWithPermitFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TransferWithPermitSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TransferWithPermitSession struct {
	Contract     *ERC20TransferWithPermit // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20TransferWithPermitCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TransferWithPermitCallerSession struct {
	Contract *ERC20TransferWithPermitCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// ERC20TransferWithPermitTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransferWithPermitTransactorSession struct {
	Contract     *ERC20TransferWithPermitTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// ERC20TransferWithPermitRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TransferWithPermitRaw struct {
	Contract *ERC20TransferWithPermit // Generic contract binding to access the raw methods on
}

// ERC20TransferWithPermitCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TransferWithPermitCallerRaw struct {
	Contract *ERC20TransferWithPermitCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransferWithPermitTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransferWithPermitTransactorRaw struct {
	Contract *ERC20TransferWithPermitTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TransferWithPermit creates a new instance of ERC20TransferWithPermit, bound to a specific deployed contract.
func NewERC20TransferWithPermit(address common.Address, backend bind.ContractBackend) (*ERC20TransferWithPermit, error) {
	contract, err := bindERC20TransferWithPermit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferWithPermit{ERC20TransferWithPermitCaller: ERC20TransferWithPermitCaller{contract: contract}, ERC20TransferWithPermitTransactor: ERC20TransferWithPermitTransactor{contract: contract}, ERC20TransferWithPermitFilterer: ERC20TransferWithPermitFilterer{contract: contract}}, nil
}

// NewERC20TransferWithPermitCaller creates a new read-only instance of ERC20TransferWithPermit, bound to a specific deployed contract.
func NewERC20TransferWithPermitCaller(address common.Address, caller bind.ContractCaller) (*ERC20TransferWithPermitCaller, error) {
	contract, err := bindERC20TransferWithPermit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferWithPermitCaller{contract: contract}, nil
}

// NewERC20TransferWithPermitTransactor creates a new write-only instance of ERC20TransferWithPermit, bound to a specific deployed contract.
func NewERC20TransferWithPermitTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TransferWithPermitTransactor, error) {
	contract, err := bindERC20TransferWithPermit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferWithPermitTransactor{contract: contract}, nil
}

// NewERC20TransferWithPermitFilterer creates a new log filterer instance of ERC20TransferWithPermit, bound to a specific deployed contract.
func NewERC20TransferWithPermitFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TransferWithPermitFilterer, error) {
	contract, err := bindERC20TransferWithPermit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferWithPermitFilterer{contract: contract}, nil
}

// bindERC20TransferWithPermit binds a generic wrapper to an already deployed contract.
func bindERC20TransferWithPermit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20TransferWithPermitMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TransferWithPermit *ERC20TransferWithPermitRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TransferWithPermit.Contract.ERC20TransferWithPermitCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TransferWithPermit *ERC20TransferWithPermitRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.ERC20TransferWithPermitTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TransferWithPermit *ERC20TransferWithPermitRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.ERC20TransferWithPermitTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TransferWithPermit *ERC20TransferWithPermitCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TransferWithPermit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TransferWithPermit *ERC20TransferWithPermitTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TransferWithPermit *ERC20TransferWithPermitTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TransferWithPermit *ERC20TransferWithPermitCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TransferWithPermit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TransferWithPermit *ERC20TransferWithPermitSession) Owner() (common.Address, error) {
	return _ERC20TransferWithPermit.Contract.Owner(&_ERC20TransferWithPermit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20TransferWithPermit *ERC20TransferWithPermitCallerSession) Owner() (common.Address, error) {
	return _ERC20TransferWithPermit.Contract.Owner(&_ERC20TransferWithPermit.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TransferWithPermit *ERC20TransferWithPermitTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TransferWithPermit *ERC20TransferWithPermitSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.RenounceOwnership(&_ERC20TransferWithPermit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20TransferWithPermit *ERC20TransferWithPermitTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.RenounceOwnership(&_ERC20TransferWithPermit.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TransferWithPermit *ERC20TransferWithPermitTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TransferWithPermit *ERC20TransferWithPermitSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.TransferOwnership(&_ERC20TransferWithPermit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20TransferWithPermit *ERC20TransferWithPermitTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.TransferOwnership(&_ERC20TransferWithPermit.TransactOpts, newOwner)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0xfdd95894.
//
// Solidity: function transferWithPermit(address token_, address owner_, address recipient_, uint256 amount_, uint256 feeAmount_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns()
func (_ERC20TransferWithPermit *ERC20TransferWithPermitTransactor) TransferWithPermit(opts *bind.TransactOpts, token_ common.Address, owner_ common.Address, recipient_ common.Address, amount_ *big.Int, feeAmount_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.contract.Transact(opts, "transferWithPermit", token_, owner_, recipient_, amount_, feeAmount_, deadline_, v_, r_, s_)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0xfdd95894.
//
// Solidity: function transferWithPermit(address token_, address owner_, address recipient_, uint256 amount_, uint256 feeAmount_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns()
func (_ERC20TransferWithPermit *ERC20TransferWithPermitSession) TransferWithPermit(token_ common.Address, owner_ common.Address, recipient_ common.Address, amount_ *big.Int, feeAmount_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.TransferWithPermit(&_ERC20TransferWithPermit.TransactOpts, token_, owner_, recipient_, amount_, feeAmount_, deadline_, v_, r_, s_)
}

// TransferWithPermit is a paid mutator transaction binding the contract method 0xfdd95894.
//
// Solidity: function transferWithPermit(address token_, address owner_, address recipient_, uint256 amount_, uint256 feeAmount_, uint256 deadline_, uint8 v_, bytes32 r_, bytes32 s_) returns()
func (_ERC20TransferWithPermit *ERC20TransferWithPermitTransactorSession) TransferWithPermit(token_ common.Address, owner_ common.Address, recipient_ common.Address, amount_ *big.Int, feeAmount_ *big.Int, deadline_ *big.Int, v_ uint8, r_ [32]byte, s_ [32]byte) (*types.Transaction, error) {
	return _ERC20TransferWithPermit.Contract.TransferWithPermit(&_ERC20TransferWithPermit.TransactOpts, token_, owner_, recipient_, amount_, feeAmount_, deadline_, v_, r_, s_)
}

// ERC20TransferWithPermitOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20TransferWithPermit contract.
type ERC20TransferWithPermitOwnershipTransferredIterator struct {
	Event *ERC20TransferWithPermitOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20TransferWithPermitOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TransferWithPermitOwnershipTransferred)
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
		it.Event = new(ERC20TransferWithPermitOwnershipTransferred)
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
func (it *ERC20TransferWithPermitOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferWithPermitOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TransferWithPermitOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20TransferWithPermit contract.
type ERC20TransferWithPermitOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TransferWithPermit *ERC20TransferWithPermitFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20TransferWithPermitOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TransferWithPermit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferWithPermitOwnershipTransferredIterator{contract: _ERC20TransferWithPermit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20TransferWithPermit *ERC20TransferWithPermitFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20TransferWithPermitOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20TransferWithPermit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TransferWithPermitOwnershipTransferred)
				if err := _ERC20TransferWithPermit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20TransferWithPermit *ERC20TransferWithPermitFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20TransferWithPermitOwnershipTransferred, error) {
	event := new(ERC20TransferWithPermitOwnershipTransferred)
	if err := _ERC20TransferWithPermit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
