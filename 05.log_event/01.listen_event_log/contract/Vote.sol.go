// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Vote

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

// VoteMetaData contains all meta data concerning the Vote contract.
var VoteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"proposal\",\"type\":\"uint8\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_proposal\",\"type\":\"uint8\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voted\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516108573803806108578339818101604052810190610032919061007a565b80600181905550506100a7565b600080fd5b6000819050919050565b61005781610044565b811461006257600080fd5b50565b6000815190506100748161004e565b92915050565b6000602082840312156100905761008f61003f565b5b600061009e84828501610065565b91505092915050565b6107a1806100b66000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063aec2ccae1161005b578063aec2ccae146100dc578063b3f98adc1461010c578063e168c3ec14610128578063fc79c9d1146101465761007d565b80633197cbb6146100825780636bd4c024146100a05780639a59dd93146100be575b600080fd5b61008a610164565b6040516100979190610473565b60405180910390f35b6100a861016a565b6040516100b59190610473565b60405180910390f35b6100c6610170565b6040516100d39190610473565b60405180910390f35b6100f660048036038101906100f191906104f1565b610176565b604051610103919061053a565b60405180910390f35b61012660048036038101906101219190610581565b610196565b005b610130610430565b60405161013d9190610473565b60405180910390f35b61014e610454565b60405161015b9190610473565b60405180910390f35b60015481565b60045481565b60025481565b60006020528060005260406000206000915054906101000a900460ff1681565b60015442106101da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d19061060b565b60405180910390fd5b60018160ff16101580156101f2575060038160ff1611155b610231576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022890610677565b60405180910390fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905060008160ff1614610311578160ff168160ff16146103105760018160ff16036102c157600260008154809291906102b7906106c6565b919050555061030f565b60028160ff16036102e957600360008154809291906102df906106c6565b919050555061030e565b60038160ff160361030d5760046000815480929190610307906106c6565b91905055505b5b5b5b5b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908360ff16021790555060018260ff16036103905760026000815480929190610386906106ef565b91905055506103de565b60028260ff16036103b857600360008154809291906103ae906106ef565b91905055506103dd565b60038260ff16036103dc57600460008154809291906103d6906106ef565b91905055505b5b5b3373ffffffffffffffffffffffffffffffffffffffff167f14075e33dffdc00c1fcaf75d0f86d667170be57fb566b12b41b499e5fa53b12583604051610424919061053a565b60405180910390a25050565b60006004546003546002546104459190610737565b61044f9190610737565b905090565b60035481565b6000819050919050565b61046d8161045a565b82525050565b60006020820190506104886000830184610464565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104be82610493565b9050919050565b6104ce816104b3565b81146104d957600080fd5b50565b6000813590506104eb816104c5565b92915050565b6000602082840312156105075761050661048e565b5b6000610515848285016104dc565b91505092915050565b600060ff82169050919050565b6105348161051e565b82525050565b600060208201905061054f600083018461052b565b92915050565b61055e8161051e565b811461056957600080fd5b50565b60008135905061057b81610555565b92915050565b6000602082840312156105975761059661048e565b5b60006105a58482850161056c565b91505092915050565b600082825260208201905092915050565b7f566f746520657870697265642e00000000000000000000000000000000000000600082015250565b60006105f5600d836105ae565b9150610600826105bf565b602082019050919050565b60006020820190508181036000830152610624816105e8565b9050919050565b7f496e76616c69642070726f706f73616c2e000000000000000000000000000000600082015250565b60006106616011836105ae565b915061066c8261062b565b602082019050919050565b6000602082019050818103600083015261069081610654565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006106d18261045a565b9150600082036106e4576106e3610697565b5b600182039050919050565b60006106fa8261045a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361072c5761072b610697565b5b600182019050919050565b60006107428261045a565b915061074d8361045a565b925082820190508082111561076557610764610697565b5b9291505056fea264697066735822122018dff0b38468015d2395defc6c4f80a49fbb204ef143bd6fc211eb9b2c96b08164736f6c63430008110033",
}

// VoteABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteMetaData.ABI instead.
var VoteABI = VoteMetaData.ABI

// VoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoteMetaData.Bin instead.
var VoteBin = VoteMetaData.Bin

// DeployVote deploys a new Ethereum contract, binding an instance of Vote to it.
func DeployVote(auth *bind.TransactOpts, backend bind.ContractBackend, _endTime *big.Int) (common.Address, *types.Transaction, *Vote, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoteBin), backend, _endTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Vote *VoteCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Vote *VoteSession) EndTime() (*big.Int, error) {
	return _Vote.Contract.EndTime(&_Vote.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Vote *VoteCallerSession) EndTime() (*big.Int, error) {
	return _Vote.Contract.EndTime(&_Vote.CallOpts)
}

// ProposalA is a free data retrieval call binding the contract method 0x9a59dd93.
//
// Solidity: function proposalA() view returns(uint256)
func (_Vote *VoteCaller) ProposalA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "proposalA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalA is a free data retrieval call binding the contract method 0x9a59dd93.
//
// Solidity: function proposalA() view returns(uint256)
func (_Vote *VoteSession) ProposalA() (*big.Int, error) {
	return _Vote.Contract.ProposalA(&_Vote.CallOpts)
}

// ProposalA is a free data retrieval call binding the contract method 0x9a59dd93.
//
// Solidity: function proposalA() view returns(uint256)
func (_Vote *VoteCallerSession) ProposalA() (*big.Int, error) {
	return _Vote.Contract.ProposalA(&_Vote.CallOpts)
}

// ProposalB is a free data retrieval call binding the contract method 0xfc79c9d1.
//
// Solidity: function proposalB() view returns(uint256)
func (_Vote *VoteCaller) ProposalB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "proposalB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalB is a free data retrieval call binding the contract method 0xfc79c9d1.
//
// Solidity: function proposalB() view returns(uint256)
func (_Vote *VoteSession) ProposalB() (*big.Int, error) {
	return _Vote.Contract.ProposalB(&_Vote.CallOpts)
}

// ProposalB is a free data retrieval call binding the contract method 0xfc79c9d1.
//
// Solidity: function proposalB() view returns(uint256)
func (_Vote *VoteCallerSession) ProposalB() (*big.Int, error) {
	return _Vote.Contract.ProposalB(&_Vote.CallOpts)
}

// ProposalC is a free data retrieval call binding the contract method 0x6bd4c024.
//
// Solidity: function proposalC() view returns(uint256)
func (_Vote *VoteCaller) ProposalC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "proposalC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalC is a free data retrieval call binding the contract method 0x6bd4c024.
//
// Solidity: function proposalC() view returns(uint256)
func (_Vote *VoteSession) ProposalC() (*big.Int, error) {
	return _Vote.Contract.ProposalC(&_Vote.CallOpts)
}

// ProposalC is a free data retrieval call binding the contract method 0x6bd4c024.
//
// Solidity: function proposalC() view returns(uint256)
func (_Vote *VoteCallerSession) ProposalC() (*big.Int, error) {
	return _Vote.Contract.ProposalC(&_Vote.CallOpts)
}

// Voted is a free data retrieval call binding the contract method 0xaec2ccae.
//
// Solidity: function voted(address ) view returns(uint8)
func (_Vote *VoteCaller) Voted(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "voted", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Voted is a free data retrieval call binding the contract method 0xaec2ccae.
//
// Solidity: function voted(address ) view returns(uint8)
func (_Vote *VoteSession) Voted(arg0 common.Address) (uint8, error) {
	return _Vote.Contract.Voted(&_Vote.CallOpts, arg0)
}

// Voted is a free data retrieval call binding the contract method 0xaec2ccae.
//
// Solidity: function voted(address ) view returns(uint8)
func (_Vote *VoteCallerSession) Voted(arg0 common.Address) (uint8, error) {
	return _Vote.Contract.Voted(&_Vote.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0xe168c3ec.
//
// Solidity: function votes() view returns(uint256)
func (_Vote *VoteCaller) Votes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "votes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0xe168c3ec.
//
// Solidity: function votes() view returns(uint256)
func (_Vote *VoteSession) Votes() (*big.Int, error) {
	return _Vote.Contract.Votes(&_Vote.CallOpts)
}

// Votes is a free data retrieval call binding the contract method 0xe168c3ec.
//
// Solidity: function votes() view returns(uint256)
func (_Vote *VoteCallerSession) Votes() (*big.Int, error) {
	return _Vote.Contract.Votes(&_Vote.CallOpts)
}

// Vote is a paid mutator transaction binding the contract method 0xb3f98adc.
//
// Solidity: function vote(uint8 _proposal) returns()
func (_Vote *VoteTransactor) Vote(opts *bind.TransactOpts, _proposal uint8) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "vote", _proposal)
}

// Vote is a paid mutator transaction binding the contract method 0xb3f98adc.
//
// Solidity: function vote(uint8 _proposal) returns()
func (_Vote *VoteSession) Vote(_proposal uint8) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, _proposal)
}

// Vote is a paid mutator transaction binding the contract method 0xb3f98adc.
//
// Solidity: function vote(uint8 _proposal) returns()
func (_Vote *VoteTransactorSession) Vote(_proposal uint8) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, _proposal)
}

// VoteVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the Vote contract.
type VoteVotedIterator struct {
	Event *VoteVoted // Event containing the contract specifics and raw log

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
func (it *VoteVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteVoted)
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
		it.Event = new(VoteVoted)
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
func (it *VoteVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteVoted represents a Voted event raised by the Vote contract.
type VoteVoted struct {
	Voter    common.Address
	Proposal uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x14075e33dffdc00c1fcaf75d0f86d667170be57fb566b12b41b499e5fa53b125.
//
// Solidity: event Voted(address indexed voter, uint8 proposal)
func (_Vote *VoteFilterer) FilterVoted(opts *bind.FilterOpts, voter []common.Address) (*VoteVotedIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "Voted", voterRule)
	if err != nil {
		return nil, err
	}
	return &VoteVotedIterator{contract: _Vote.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x14075e33dffdc00c1fcaf75d0f86d667170be57fb566b12b41b499e5fa53b125.
//
// Solidity: event Voted(address indexed voter, uint8 proposal)
func (_Vote *VoteFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *VoteVoted, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "Voted", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteVoted)
				if err := _Vote.contract.UnpackLog(event, "Voted", log); err != nil {
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

// ParseVoted is a log parse operation binding the contract event 0x14075e33dffdc00c1fcaf75d0f86d667170be57fb566b12b41b499e5fa53b125.
//
// Solidity: event Voted(address indexed voter, uint8 proposal)
func (_Vote *VoteFilterer) ParseVoted(log types.Log) (*VoteVoted, error) {
	event := new(VoteVoted)
	if err := _Vote.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
