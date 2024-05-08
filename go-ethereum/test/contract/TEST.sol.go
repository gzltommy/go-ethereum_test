// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// TESTMetaData contains all meta data concerning the TEST contract.
var TESTMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"StakedNFTs\",\"type\":\"event\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212205cff8ccb5568d16632a092690fd65147c4fb34649105a461d97eeedb4730738764736f6c63430007060033",
}

// TESTABI is the input ABI used to generate the binding from.
// Deprecated: Use TESTMetaData.ABI instead.
var TESTABI = TESTMetaData.ABI

// TESTBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TESTMetaData.Bin instead.
var TESTBin = TESTMetaData.Bin

// DeployTEST deploys a new Ethereum contract, binding an instance of TEST to it.
func DeployTEST(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TEST, error) {
	parsed, err := TESTMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TESTBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TEST{TESTCaller: TESTCaller{contract: contract}, TESTTransactor: TESTTransactor{contract: contract}, TESTFilterer: TESTFilterer{contract: contract}}, nil
}

// TEST is an auto generated Go binding around an Ethereum contract.
type TEST struct {
	TESTCaller     // Read-only binding to the contract
	TESTTransactor // Write-only binding to the contract
	TESTFilterer   // Log filterer for contract events
}

// TESTCaller is an auto generated read-only Go binding around an Ethereum contract.
type TESTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TESTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TESTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TESTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TESTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TESTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TESTSession struct {
	Contract     *TEST             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TESTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TESTCallerSession struct {
	Contract *TESTCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TESTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TESTTransactorSession struct {
	Contract     *TESTTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TESTRaw is an auto generated low-level Go binding around an Ethereum contract.
type TESTRaw struct {
	Contract *TEST // Generic contract binding to access the raw methods on
}

// TESTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TESTCallerRaw struct {
	Contract *TESTCaller // Generic read-only contract binding to access the raw methods on
}

// TESTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TESTTransactorRaw struct {
	Contract *TESTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTEST creates a new instance of TEST, bound to a specific deployed contract.
func NewTEST(address common.Address, backend bind.ContractBackend) (*TEST, error) {
	contract, err := bindTEST(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TEST{TESTCaller: TESTCaller{contract: contract}, TESTTransactor: TESTTransactor{contract: contract}, TESTFilterer: TESTFilterer{contract: contract}}, nil
}

// NewTESTCaller creates a new read-only instance of TEST, bound to a specific deployed contract.
func NewTESTCaller(address common.Address, caller bind.ContractCaller) (*TESTCaller, error) {
	contract, err := bindTEST(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TESTCaller{contract: contract}, nil
}

// NewTESTTransactor creates a new write-only instance of TEST, bound to a specific deployed contract.
func NewTESTTransactor(address common.Address, transactor bind.ContractTransactor) (*TESTTransactor, error) {
	contract, err := bindTEST(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TESTTransactor{contract: contract}, nil
}

// NewTESTFilterer creates a new log filterer instance of TEST, bound to a specific deployed contract.
func NewTESTFilterer(address common.Address, filterer bind.ContractFilterer) (*TESTFilterer, error) {
	contract, err := bindTEST(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TESTFilterer{contract: contract}, nil
}

// bindTEST binds a generic wrapper to an already deployed contract.
func bindTEST(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TESTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TEST *TESTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TEST.Contract.TESTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TEST *TESTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TEST.Contract.TESTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TEST *TESTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TEST.Contract.TESTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TEST *TESTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TEST.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TEST *TESTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TEST.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TEST *TESTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TEST.Contract.contract.Transact(opts, method, params...)
}

// TESTStakedNFTsIterator is returned from FilterStakedNFTs and is used to iterate over the raw logs and unpacked data for StakedNFTs events raised by the TEST contract.
type TESTStakedNFTsIterator struct {
	Event *TESTStakedNFTs // Event containing the contract specifics and raw log

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
func (it *TESTStakedNFTsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TESTStakedNFTs)
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
		it.Event = new(TESTStakedNFTs)
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
func (it *TESTStakedNFTsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TESTStakedNFTsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TESTStakedNFTs represents a StakedNFTs event raised by the TEST contract.
type TESTStakedNFTs struct {
	User    common.Address
	Ids     []*big.Int
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakedNFTs is a free log retrieval operation binding the contract event 0xc90a5a2ae8a0891e7dca598cd9c609f2b61a580415fc7c7ce18ea1fdad82faa8.
//
// Solidity: event StakedNFTs(address indexed user, uint256[] ids, uint256[] amounts)
func (_TEST *TESTFilterer) FilterStakedNFTs(opts *bind.FilterOpts, user []common.Address) (*TESTStakedNFTsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TEST.contract.FilterLogs(opts, "StakedNFTs", userRule)
	if err != nil {
		return nil, err
	}
	return &TESTStakedNFTsIterator{contract: _TEST.contract, event: "StakedNFTs", logs: logs, sub: sub}, nil
}

// WatchStakedNFTs is a free log subscription operation binding the contract event 0xc90a5a2ae8a0891e7dca598cd9c609f2b61a580415fc7c7ce18ea1fdad82faa8.
//
// Solidity: event StakedNFTs(address indexed user, uint256[] ids, uint256[] amounts)
func (_TEST *TESTFilterer) WatchStakedNFTs(opts *bind.WatchOpts, sink chan<- *TESTStakedNFTs, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _TEST.contract.WatchLogs(opts, "StakedNFTs", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TESTStakedNFTs)
				if err := _TEST.contract.UnpackLog(event, "StakedNFTs", log); err != nil {
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

// ParseStakedNFTs is a log parse operation binding the contract event 0xc90a5a2ae8a0891e7dca598cd9c609f2b61a580415fc7c7ce18ea1fdad82faa8.
//
// Solidity: event StakedNFTs(address indexed user, uint256[] ids, uint256[] amounts)
func (_TEST *TESTFilterer) ParseStakedNFTs(log types.Log) (*TESTStakedNFTs, error) {
	event := new(TESTStakedNFTs)
	if err := _TEST.contract.UnpackLog(event, "StakedNFTs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
