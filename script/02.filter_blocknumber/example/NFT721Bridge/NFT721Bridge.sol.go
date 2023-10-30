// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NFT721Bridge

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

// NFT721BridgeMetaData contains all meta data concerning the NFT721Bridge contract.
var NFT721BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"sourceChain\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"sendChain\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ReceiveNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"TransferNFT\",\"type\":\"event\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea264697066735822122062470309d7995963741e9abe3a3f9a84927bcaa9beb66e6000e8f4324603e65d64736f6c63430008110033",
}

// NFT721BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use NFT721BridgeMetaData.ABI instead.
var NFT721BridgeABI = NFT721BridgeMetaData.ABI

// NFT721BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NFT721BridgeMetaData.Bin instead.
var NFT721BridgeBin = NFT721BridgeMetaData.Bin

// DeployNFT721Bridge deploys a new Ethereum contract, binding an instance of NFT721Bridge to it.
func DeployNFT721Bridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NFT721Bridge, error) {
	parsed, err := NFT721BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NFT721BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NFT721Bridge{NFT721BridgeCaller: NFT721BridgeCaller{contract: contract}, NFT721BridgeTransactor: NFT721BridgeTransactor{contract: contract}, NFT721BridgeFilterer: NFT721BridgeFilterer{contract: contract}}, nil
}

// NFT721Bridge is an auto generated Go binding around an Ethereum contract.
type NFT721Bridge struct {
	NFT721BridgeCaller     // Read-only binding to the contract
	NFT721BridgeTransactor // Write-only binding to the contract
	NFT721BridgeFilterer   // Log filterer for contract events
}

// NFT721BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFT721BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFT721BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFT721BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFT721BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFT721BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFT721BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFT721BridgeSession struct {
	Contract     *NFT721Bridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFT721BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFT721BridgeCallerSession struct {
	Contract *NFT721BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NFT721BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFT721BridgeTransactorSession struct {
	Contract     *NFT721BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NFT721BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFT721BridgeRaw struct {
	Contract *NFT721Bridge // Generic contract binding to access the raw methods on
}

// NFT721BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFT721BridgeCallerRaw struct {
	Contract *NFT721BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// NFT721BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFT721BridgeTransactorRaw struct {
	Contract *NFT721BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFT721Bridge creates a new instance of NFT721Bridge, bound to a specific deployed contract.
func NewNFT721Bridge(address common.Address, backend bind.ContractBackend) (*NFT721Bridge, error) {
	contract, err := bindNFT721Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFT721Bridge{NFT721BridgeCaller: NFT721BridgeCaller{contract: contract}, NFT721BridgeTransactor: NFT721BridgeTransactor{contract: contract}, NFT721BridgeFilterer: NFT721BridgeFilterer{contract: contract}}, nil
}

// NewNFT721BridgeCaller creates a new read-only instance of NFT721Bridge, bound to a specific deployed contract.
func NewNFT721BridgeCaller(address common.Address, caller bind.ContractCaller) (*NFT721BridgeCaller, error) {
	contract, err := bindNFT721Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFT721BridgeCaller{contract: contract}, nil
}

// NewNFT721BridgeTransactor creates a new write-only instance of NFT721Bridge, bound to a specific deployed contract.
func NewNFT721BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*NFT721BridgeTransactor, error) {
	contract, err := bindNFT721Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFT721BridgeTransactor{contract: contract}, nil
}

// NewNFT721BridgeFilterer creates a new log filterer instance of NFT721Bridge, bound to a specific deployed contract.
func NewNFT721BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*NFT721BridgeFilterer, error) {
	contract, err := bindNFT721Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFT721BridgeFilterer{contract: contract}, nil
}

// bindNFT721Bridge binds a generic wrapper to an already deployed contract.
func bindNFT721Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFT721BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT721Bridge *NFT721BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT721Bridge.Contract.NFT721BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT721Bridge *NFT721BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT721Bridge.Contract.NFT721BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT721Bridge *NFT721BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT721Bridge.Contract.NFT721BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT721Bridge *NFT721BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT721Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT721Bridge *NFT721BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT721Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT721Bridge *NFT721BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT721Bridge.Contract.contract.Transact(opts, method, params...)
}

// NFT721BridgeReceiveNFTIterator is returned from FilterReceiveNFT and is used to iterate over the raw logs and unpacked data for ReceiveNFT events raised by the NFT721Bridge contract.
type NFT721BridgeReceiveNFTIterator struct {
	Event *NFT721BridgeReceiveNFT // Event containing the contract specifics and raw log

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
func (it *NFT721BridgeReceiveNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFT721BridgeReceiveNFT)
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
		it.Event = new(NFT721BridgeReceiveNFT)
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
func (it *NFT721BridgeReceiveNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFT721BridgeReceiveNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFT721BridgeReceiveNFT represents a ReceiveNFT event raised by the NFT721Bridge contract.
type NFT721BridgeReceiveNFT struct {
	Nonce       uint64
	SourceToken common.Address
	Token       common.Address
	TokenID     *big.Int
	SourceChain uint16
	SendChain   uint16
	Recipient   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceiveNFT is a free log retrieval operation binding the contract event 0x32aae95950c2e1f2c1a419165ba01c63c49604db10ee1b95d9960c0f5b9b9fa8.
//
// Solidity: event ReceiveNFT(uint64 indexed nonce, address sourceToken, address token, uint256 tokenID, uint16 sourceChain, uint16 sendChain, address recipient)
func (_NFT721Bridge *NFT721BridgeFilterer) FilterReceiveNFT(opts *bind.FilterOpts, nonce []uint64) (*NFT721BridgeReceiveNFTIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _NFT721Bridge.contract.FilterLogs(opts, "ReceiveNFT", nonceRule)
	if err != nil {
		return nil, err
	}
	return &NFT721BridgeReceiveNFTIterator{contract: _NFT721Bridge.contract, event: "ReceiveNFT", logs: logs, sub: sub}, nil
}

// WatchReceiveNFT is a free log subscription operation binding the contract event 0x32aae95950c2e1f2c1a419165ba01c63c49604db10ee1b95d9960c0f5b9b9fa8.
//
// Solidity: event ReceiveNFT(uint64 indexed nonce, address sourceToken, address token, uint256 tokenID, uint16 sourceChain, uint16 sendChain, address recipient)
func (_NFT721Bridge *NFT721BridgeFilterer) WatchReceiveNFT(opts *bind.WatchOpts, sink chan<- *NFT721BridgeReceiveNFT, nonce []uint64) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _NFT721Bridge.contract.WatchLogs(opts, "ReceiveNFT", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFT721BridgeReceiveNFT)
				if err := _NFT721Bridge.contract.UnpackLog(event, "ReceiveNFT", log); err != nil {
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

// ParseReceiveNFT is a log parse operation binding the contract event 0x32aae95950c2e1f2c1a419165ba01c63c49604db10ee1b95d9960c0f5b9b9fa8.
//
// Solidity: event ReceiveNFT(uint64 indexed nonce, address sourceToken, address token, uint256 tokenID, uint16 sourceChain, uint16 sendChain, address recipient)
func (_NFT721Bridge *NFT721BridgeFilterer) ParseReceiveNFT(log types.Log) (*NFT721BridgeReceiveNFT, error) {
	event := new(NFT721BridgeReceiveNFT)
	if err := _NFT721Bridge.contract.UnpackLog(event, "ReceiveNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFT721BridgeTransferNFTIterator is returned from FilterTransferNFT and is used to iterate over the raw logs and unpacked data for TransferNFT events raised by the NFT721Bridge contract.
type NFT721BridgeTransferNFTIterator struct {
	Event *NFT721BridgeTransferNFT // Event containing the contract specifics and raw log

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
func (it *NFT721BridgeTransferNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFT721BridgeTransferNFT)
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
		it.Event = new(NFT721BridgeTransferNFT)
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
func (it *NFT721BridgeTransferNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFT721BridgeTransferNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFT721BridgeTransferNFT represents a TransferNFT event raised by the NFT721Bridge contract.
type NFT721BridgeTransferNFT struct {
	Nonce      uint64
	Token      common.Address
	TokenID    *big.Int
	DstChainId uint16
	Sender     common.Address
	Recipient  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferNFT is a free log retrieval operation binding the contract event 0xe11d2ca26838f15acb41450029a785bb3d6f909b7f622ebf9c45524ded76f411.
//
// Solidity: event TransferNFT(uint64 indexed nonce, address token, uint256 tokenID, uint16 dstChainId, address sender, address recipient)
func (_NFT721Bridge *NFT721BridgeFilterer) FilterTransferNFT(opts *bind.FilterOpts, nonce []uint64) (*NFT721BridgeTransferNFTIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _NFT721Bridge.contract.FilterLogs(opts, "TransferNFT", nonceRule)
	if err != nil {
		return nil, err
	}
	return &NFT721BridgeTransferNFTIterator{contract: _NFT721Bridge.contract, event: "TransferNFT", logs: logs, sub: sub}, nil
}

// WatchTransferNFT is a free log subscription operation binding the contract event 0xe11d2ca26838f15acb41450029a785bb3d6f909b7f622ebf9c45524ded76f411.
//
// Solidity: event TransferNFT(uint64 indexed nonce, address token, uint256 tokenID, uint16 dstChainId, address sender, address recipient)
func (_NFT721Bridge *NFT721BridgeFilterer) WatchTransferNFT(opts *bind.WatchOpts, sink chan<- *NFT721BridgeTransferNFT, nonce []uint64) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _NFT721Bridge.contract.WatchLogs(opts, "TransferNFT", nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFT721BridgeTransferNFT)
				if err := _NFT721Bridge.contract.UnpackLog(event, "TransferNFT", log); err != nil {
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

// ParseTransferNFT is a log parse operation binding the contract event 0xe11d2ca26838f15acb41450029a785bb3d6f909b7f622ebf9c45524ded76f411.
//
// Solidity: event TransferNFT(uint64 indexed nonce, address token, uint256 tokenID, uint16 dstChainId, address sender, address recipient)
func (_NFT721Bridge *NFT721BridgeFilterer) ParseTransferNFT(log types.Log) (*NFT721BridgeTransferNFT, error) {
	event := new(NFT721BridgeTransferNFT)
	if err := _NFT721Bridge.contract.UnpackLog(event, "TransferNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
