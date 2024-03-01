// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SignIn

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

// SignInMetaData contains all meta data concerning the SignIn contract.
var SignInMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"}],\"name\":\"eveSignIn\",\"type\":\"event\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212203dbede3dc2cb3d5d7859c0702ec910c91b1e01a6eb111b2cfed1765fe571b26664736f6c63430008000033",
}

// SignInABI is the input ABI used to generate the binding from.
// Deprecated: Use SignInMetaData.ABI instead.
var SignInABI = SignInMetaData.ABI

// SignInBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignInMetaData.Bin instead.
var SignInBin = SignInMetaData.Bin

// DeploySignIn deploys a new Ethereum contract, binding an instance of SignIn to it.
func DeploySignIn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignIn, error) {
	parsed, err := SignInMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignInBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignIn{SignInCaller: SignInCaller{contract: contract}, SignInTransactor: SignInTransactor{contract: contract}, SignInFilterer: SignInFilterer{contract: contract}}, nil
}

// SignIn is an auto generated Go binding around an Ethereum contract.
type SignIn struct {
	SignInCaller     // Read-only binding to the contract
	SignInTransactor // Write-only binding to the contract
	SignInFilterer   // Log filterer for contract events
}

// SignInCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignInCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignInTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignInTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignInFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignInFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignInSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignInSession struct {
	Contract     *SignIn           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignInCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignInCallerSession struct {
	Contract *SignInCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SignInTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignInTransactorSession struct {
	Contract     *SignInTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignInRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignInRaw struct {
	Contract *SignIn // Generic contract binding to access the raw methods on
}

// SignInCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignInCallerRaw struct {
	Contract *SignInCaller // Generic read-only contract binding to access the raw methods on
}

// SignInTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignInTransactorRaw struct {
	Contract *SignInTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignIn creates a new instance of SignIn, bound to a specific deployed contract.
func NewSignIn(address common.Address, backend bind.ContractBackend) (*SignIn, error) {
	contract, err := bindSignIn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignIn{SignInCaller: SignInCaller{contract: contract}, SignInTransactor: SignInTransactor{contract: contract}, SignInFilterer: SignInFilterer{contract: contract}}, nil
}

// NewSignInCaller creates a new read-only instance of SignIn, bound to a specific deployed contract.
func NewSignInCaller(address common.Address, caller bind.ContractCaller) (*SignInCaller, error) {
	contract, err := bindSignIn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignInCaller{contract: contract}, nil
}

// NewSignInTransactor creates a new write-only instance of SignIn, bound to a specific deployed contract.
func NewSignInTransactor(address common.Address, transactor bind.ContractTransactor) (*SignInTransactor, error) {
	contract, err := bindSignIn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignInTransactor{contract: contract}, nil
}

// NewSignInFilterer creates a new log filterer instance of SignIn, bound to a specific deployed contract.
func NewSignInFilterer(address common.Address, filterer bind.ContractFilterer) (*SignInFilterer, error) {
	contract, err := bindSignIn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignInFilterer{contract: contract}, nil
}

// bindSignIn binds a generic wrapper to an already deployed contract.
func bindSignIn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignInABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignIn *SignInRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignIn.Contract.SignInCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignIn *SignInRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignIn.Contract.SignInTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignIn *SignInRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignIn.Contract.SignInTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignIn *SignInCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignIn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignIn *SignInTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignIn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignIn *SignInTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignIn.Contract.contract.Transact(opts, method, params...)
}

// SignInEveSignInIterator is returned from FilterEveSignIn and is used to iterate over the raw logs and unpacked data for EveSignIn events raised by the SignIn contract.
type SignInEveSignInIterator struct {
	Event *SignInEveSignIn // Event containing the contract specifics and raw log

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
func (it *SignInEveSignInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignInEveSignIn)
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
		it.Event = new(SignInEveSignIn)
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
func (it *SignInEveSignInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignInEveSignInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignInEveSignIn represents a EveSignIn event raised by the SignIn contract.
type SignInEveSignIn struct {
	Operator common.Address
	Date     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEveSignIn is a free log retrieval operation binding the contract event 0x2a0cd797412ef7ea19c365def73dc12d3ec05fbe2c714249219e6b2739cce14a.
//
// Solidity: event eveSignIn(address indexed operator, uint256 indexed date)
func (_SignIn *SignInFilterer) FilterEveSignIn(opts *bind.FilterOpts, operator []common.Address, date []*big.Int) (*SignInEveSignInIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _SignIn.contract.FilterLogs(opts, "eveSignIn", operatorRule, dateRule)
	if err != nil {
		return nil, err
	}
	return &SignInEveSignInIterator{contract: _SignIn.contract, event: "eveSignIn", logs: logs, sub: sub}, nil
}

// WatchEveSignIn is a free log subscription operation binding the contract event 0x2a0cd797412ef7ea19c365def73dc12d3ec05fbe2c714249219e6b2739cce14a.
//
// Solidity: event eveSignIn(address indexed operator, uint256 indexed date)
func (_SignIn *SignInFilterer) WatchEveSignIn(opts *bind.WatchOpts, sink chan<- *SignInEveSignIn, operator []common.Address, date []*big.Int) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var dateRule []interface{}
	for _, dateItem := range date {
		dateRule = append(dateRule, dateItem)
	}

	logs, sub, err := _SignIn.contract.WatchLogs(opts, "eveSignIn", operatorRule, dateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignInEveSignIn)
				if err := _SignIn.contract.UnpackLog(event, "eveSignIn", log); err != nil {
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

// ParseEveSignIn is a log parse operation binding the contract event 0x2a0cd797412ef7ea19c365def73dc12d3ec05fbe2c714249219e6b2739cce14a.
//
// Solidity: event eveSignIn(address indexed operator, uint256 indexed date)
func (_SignIn *SignInFilterer) ParseEveSignIn(log types.Log) (*SignInEveSignIn, error) {
	event := new(SignInEveSignIn)
	if err := _SignIn.contract.UnpackLog(event, "eveSignIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
