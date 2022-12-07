// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SecondLiveMedal

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SpaceStationABI is the input ABI used to generate the binding from.
const SpaceStationABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc20Fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"EventActivateCampaign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dummyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_mintTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_canTranster\",\"type\":\"bool\"}],\"name\":\"EventClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"campaignSetter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"treasuryManager\",\"type\":\"address\"}],\"name\":\"UpdateManager\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc20Fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_platformFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"activateCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaignFeeConfigs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"erc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc20Fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"platformFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"campaignSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"caps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dummyId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_level\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_mintTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyReceiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"_royaltyFeeNumerator\",\"type\":\"uint96\"},{\"internalType\":\"bool\",\"name\":\"_canTranster\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_campaignSetter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_treasuryManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"numClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_campaignSetter\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_treasuryManager\",\"type\":\"address\"}],\"name\":\"setManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryManager\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// SpaceStationBin is the compiled bytecode used for deploying new contracts.
var SpaceStationBin = "0x608060405234801561001057600080fd5b5061001a33610024565b6001600755610076565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6114f6806100856000396000f3fe6080604052600436106100e15760003560e01c80639e34070f1161007f578063e32ead2d11610059578063e32ead2d1461028c578063f2fde38b146102f7578063f6a6209614610317578063f8c8765e1461033757600080fd5b80639e34070f146101ff578063c37c656d1461023f578063e164adef1461026c57600080fd5b80633cea70d9116100bb5780633cea70d9146101715780636b5da37b14610191578063715018a6146101cc5780638da5cb5b146101e157600080fd5b806315171d06146100ed5780631eccf1021461012a578063238ac9331461014c57600080fd5b366100e857005b600080fd5b3480156100f957600080fd5b5060095461010d906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561013657600080fd5b5061014a6101453660046111af565b610357565b005b34801561015857600080fd5b5060085461010d9061010090046001600160a01b031681565b34801561017d57600080fd5b50600a5461010d906001600160a01b031681565b34801561019d57600080fd5b506101be6101ac3660046111f7565b600c6020526000908152604090205481565b604051908152602001610121565b3480156101d857600080fd5b5061014a61046b565b3480156101ed57600080fd5b506006546001600160a01b031661010d565b34801561020b57600080fd5b5061022f61021a3660046111f7565b600b6020526000908152604090205460ff1681565b6040519015158152602001610121565b34801561024b57600080fd5b506101be61025a3660046111f7565b600d6020526000908152604090205481565b34801561027857600080fd5b506101be610287366004611277565b6104a1565b34801561029857600080fd5b506102d26102a73660046111f7565b600e602052600090815260409020805460018201546002909201546001600160a01b03909116919083565b604080516001600160a01b039094168452602084019290925290820152606001610121565b34801561030357600080fd5b5061014a610312366004611343565b61080a565b34801561032357600080fd5b5061014a610332366004611360565b6108a5565b34801561034357600080fd5b5061014a6103523660046113ab565b610962565b6009546001600160a01b031633146103b65760405162461bcd60e51b815260206004820152601c60248201527f4f6e6c792063616d706169676e5365747465722063616e2063616c6c0000000060448201526064015b60405180910390fd5b60408051606080820183526001600160a01b03878116808452602080850189815285870189815260008d8152600e8452888120975188546001600160a01b03191696169590951787559051600187015551600290950194909455600d845290849020859055835189815292830152918101859052908101839052608081018290527feea389ac416aa738ce024ddbda928dd6734088c64666196949c5465fc811dc5d9060a00160405180910390a15050505050565b6006546001600160a01b031633146104955760405162461bcd60e51b81526004016103ad90611407565b61049f6000610a5f565b565b60006002600754036104f55760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016103ad565b60026007556000898152600b602052604090205460ff161561054c5760405162461bcd60e51b815260206004820152601060248201526f416c726561647920436c61696d65642160801b60448201526064016103ad565b60008b8152600d6020526040902054156105b55760008b8152600d6020908152604080832054600c90925290912054106105b55760405162461bcd60e51b815260206004820152600a602482015269436c61696d20456e642160b01b60448201526064016103ad565b6105cf6105c88c8c8c8c8c8c8c8c610ab1565b8484610b5a565b61060f5760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b60448201526064016103ad565b6000898152600b60209081526040808320805460ff1916600117905580516080810182528e81529182018b81526001600160a01b038b81168484018181528a151560608701908152945163325306f560e21b815260048101929092528551602483015292516044820152915181166064830152915115156084820152919291908d169063c94c1bd49060a4016020604051808303816000875af11580156106ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106de919061143c565b90506001600160601b0387161561076157604051635944c75360e01b8152600481018290526001600160a01b0389811660248301526001600160601b03891660448301528d1690635944c75390606401600060405180830381600087803b15801561074857600080fd5b505af115801561075c573d6000803e3d6000fd5b505050505b60008d8152600c6020526040812080549161077b8361146b565b919050555061078b8d6001610bc0565b60408051828152602081018f90529081018c9052606081018b90526001600160a01b03808e1660808301528a1660a082015286151560c08201528c907f4acbe73b979f0a6452384be471194eef86ad06759695d9d54c94191fb5f29e8c9060e00160405180910390a15060016007559c9b505050505050505050505050565b6006546001600160a01b031633146108345760405162461bcd60e51b81526004016103ad90611407565b6001600160a01b0381166108995760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103ad565b6108a281610a5f565b50565b6006546001600160a01b031633146108cf5760405162461bcd60e51b81526004016103ad90611407565b60088054610100600160a81b0319166101006001600160a01b0386811691820292909217909255600980546001600160a01b0319908116868416908117909255600a805490911692851692831790556040805193845260208401919091528201527fea28c9444370d269c92153ac32abd8ab32f1f963064504e4c41c2b6b726840949060600160405180910390a1505050565b60085460ff16156109b55760405162461bcd60e51b815260206004820181905260248201527f696e697469616c697a653a20416c726561647920696e697469616c697a65642160448201526064016103ad565b6109be84610a5f565b610a076040518060400160405280600a8152602001695365636f6e644c69766560b01b815250604051806040016040528060058152602001640312e302e360dc1b815250610dd9565b60088054600980546001600160a01b03199081166001600160a01b0396871617909155600a8054909116938516939093179092556001600160a81b0319909116610100939092169290920260ff191617600117905550565b600680546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b604080517fcf0be1d1407500823cfdecc7a49d7db290f935da391acdd844e2517557689dbe60208201529081018990526001600160a01b0380891660608301526080820188905260a0820187905280861660c0830152841660e08201526001600160601b038316610100820152811515610120820152600090610b4d906101400160405160208183030381529060405280519060200120610ea4565b9998505050505050505050565b600854604080516020601f850181900481028201810190925283815260009261010090046001600160a01b031691610bae9187918790879081908401838280828437600092019190915250610eeb92505050565b6001600160a01b031614949350505050565b60008111610c085760405162461bcd60e51b815260206004820152601560248201527404d757374206d696e74206d6f7265207468616e203605c1b60448201526064016103ad565b6000828152600e6020908152604091829020825160608101845281546001600160a01b03168152600182015492810192909252600201549181018290529015610cdb576040810151610c5a9083610f66565b341015610ca05760405162461bcd60e51b8152602060048201526014602482015273125b9cdd59999a58da595b9d0814185e5b595b9d60621b60448201526064016103ad565b600a546040516001600160a01b03909116903480156108fc02916000818181858888f19350505050158015610cd9573d6000803e3d6000fd5b505b602081015115610dd4578051600a5460208301516001600160a01b03928316926323b872dd923392911690610d109087610f66565b6040516001600160e01b031960e086901b1681526001600160a01b03938416600482015292909116602483015260448201526064016020604051808303816000875af1158015610d64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d889190611484565b610dd45760405162461bcd60e51b815260206004820152601860248201527f5472616e73666572206572633230466565206661696c6564000000000000000060448201526064016103ad565b505050565b60005460ff1615610e385760405162461bcd60e51b8152602060048201526024808201527f656970496e697469616c697a65643a20416c726561647920696e697469616c696044820152637a65642160e01b60648201526084016103ad565b8151602080840191909120825191830191909120600382905560048190557f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f46600255610e86818484610f79565b60019081556005919091556000805460ff1916909117905550505050565b6000610eae610fc3565b60405161190160f01b6020820152602281019190915260428101839052606201604051602081830303815290604052805190602001209050919050565b60008151604114610f3e5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016103ad565b60208201516040830151606084015160001a610f5c86828585610ff1565b9695505050505050565b6000610f7282846114a1565b9392505050565b600083838346604080516020810195909552840192909252606083015260808201523060a082015260c0016040516020818303038152906040528051906020012090509392505050565b6000600254610fcf4690565b03610fdb575060015490565b610fec600554600354600454610f79565b905090565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a082111561106e5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016103ad565b8360ff16601b148061108357508360ff16601c145b6110da5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c604482015261756560f01b60648201526084016103ad565b6040805160008082526020820180845288905260ff871692820192909252606081018590526080810184905260019060a0016020604051602081039080840390855afa15801561112e573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166111915760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016103ad565b95945050505050565b6001600160a01b03811681146108a257600080fd5b600080600080600060a086880312156111c757600080fd5b8535945060208601356111d98161119a565b94979496505050506040830135926060810135926080909101359150565b60006020828403121561120957600080fd5b5035919050565b80151581146108a257600080fd5b803561122981611210565b919050565b60008083601f84011261124057600080fd5b50813567ffffffffffffffff81111561125857600080fd5b60208301915083602082850101111561127057600080fd5b9250929050565b6000806000806000806000806000806101208b8d03121561129757600080fd5b8a35995060208b01356112a98161119a565b985060408b0135975060608b0135965060808b01356112c78161119a565b955060a08b01356112d78161119a565b945060c08b01356001600160601b03811681146112f357600080fd5b935061130160e08c0161121e565b92506101008b013567ffffffffffffffff81111561131e57600080fd5b61132a8d828e0161122e565b915080935050809150509295989b9194979a5092959850565b60006020828403121561135557600080fd5b8135610f728161119a565b60008060006060848603121561137557600080fd5b83356113808161119a565b925060208401356113908161119a565b915060408401356113a08161119a565b809150509250925092565b600080600080608085870312156113c157600080fd5b84356113cc8161119a565b935060208501356113dc8161119a565b925060408501356113ec8161119a565b915060608501356113fc8161119a565b939692955090935050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60006020828403121561144e57600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b60006001820161147d5761147d611455565b5060010190565b60006020828403121561149657600080fd5b8151610f7281611210565b60008160001904831182151516156114bb576114bb611455565b50029056fea264697066735822122083bc551171e6eca1732a39ba54e2d9d60589f527b970cb27d6371f73f7c3eda364736f6c634300080e0033"

// DeploySpaceStation deploys a new Ethereum contract, binding an instance of SpaceStation to it.
func DeploySpaceStation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SpaceStation, error) {
	parsed, err := abi.JSON(strings.NewReader(SpaceStationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SpaceStationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SpaceStation{SpaceStationCaller: SpaceStationCaller{contract: contract}, SpaceStationTransactor: SpaceStationTransactor{contract: contract}, SpaceStationFilterer: SpaceStationFilterer{contract: contract}}, nil
}

// SpaceStation is an auto generated Go binding around an Ethereum contract.
type SpaceStation struct {
	SpaceStationCaller     // Read-only binding to the contract
	SpaceStationTransactor // Write-only binding to the contract
	SpaceStationFilterer   // Log filterer for contract events
}

// SpaceStationCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpaceStationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpaceStationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpaceStationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpaceStationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpaceStationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpaceStationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpaceStationSession struct {
	Contract     *SpaceStation     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpaceStationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpaceStationCallerSession struct {
	Contract *SpaceStationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SpaceStationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpaceStationTransactorSession struct {
	Contract     *SpaceStationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SpaceStationRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpaceStationRaw struct {
	Contract *SpaceStation // Generic contract binding to access the raw methods on
}

// SpaceStationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpaceStationCallerRaw struct {
	Contract *SpaceStationCaller // Generic read-only contract binding to access the raw methods on
}

// SpaceStationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpaceStationTransactorRaw struct {
	Contract *SpaceStationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpaceStation creates a new instance of SpaceStation, bound to a specific deployed contract.
func NewSpaceStation(address common.Address, backend bind.ContractBackend) (*SpaceStation, error) {
	contract, err := bindSpaceStation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SpaceStation{SpaceStationCaller: SpaceStationCaller{contract: contract}, SpaceStationTransactor: SpaceStationTransactor{contract: contract}, SpaceStationFilterer: SpaceStationFilterer{contract: contract}}, nil
}

// NewSpaceStationCaller creates a new read-only instance of SpaceStation, bound to a specific deployed contract.
func NewSpaceStationCaller(address common.Address, caller bind.ContractCaller) (*SpaceStationCaller, error) {
	contract, err := bindSpaceStation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpaceStationCaller{contract: contract}, nil
}

// NewSpaceStationTransactor creates a new write-only instance of SpaceStation, bound to a specific deployed contract.
func NewSpaceStationTransactor(address common.Address, transactor bind.ContractTransactor) (*SpaceStationTransactor, error) {
	contract, err := bindSpaceStation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpaceStationTransactor{contract: contract}, nil
}

// NewSpaceStationFilterer creates a new log filterer instance of SpaceStation, bound to a specific deployed contract.
func NewSpaceStationFilterer(address common.Address, filterer bind.ContractFilterer) (*SpaceStationFilterer, error) {
	contract, err := bindSpaceStation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpaceStationFilterer{contract: contract}, nil
}

// bindSpaceStation binds a generic wrapper to an already deployed contract.
func bindSpaceStation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SpaceStationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpaceStation *SpaceStationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpaceStation.Contract.SpaceStationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpaceStation *SpaceStationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpaceStation.Contract.SpaceStationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpaceStation *SpaceStationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpaceStation.Contract.SpaceStationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpaceStation *SpaceStationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpaceStation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpaceStation *SpaceStationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpaceStation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpaceStation *SpaceStationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpaceStation.Contract.contract.Transact(opts, method, params...)
}

// CampaignFeeConfigs is a free data retrieval call binding the contract method 0xe32ead2d.
//
// Solidity: function campaignFeeConfigs(uint256 ) view returns(address erc20, uint256 erc20Fee, uint256 platformFee)
func (_SpaceStation *SpaceStationCaller) CampaignFeeConfigs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Erc20       common.Address
	Erc20Fee    *big.Int
	PlatformFee *big.Int
}, error) {
	var out []interface{}
	err := _SpaceStation.contract.Call(opts, &out, "campaignFeeConfigs", arg0)

	outstruct := new(struct {
		Erc20       common.Address
		Erc20Fee    *big.Int
		PlatformFee *big.Int
	})

	outstruct.Erc20 = out[0].(common.Address)
	outstruct.Erc20Fee = out[1].(*big.Int)
	outstruct.PlatformFee = out[2].(*big.Int)

	return *outstruct, err

}

// CampaignFeeConfigs is a free data retrieval call binding the contract method 0xe32ead2d.
//
// Solidity: function campaignFeeConfigs(uint256 ) view returns(address erc20, uint256 erc20Fee, uint256 platformFee)
func (_SpaceStation *SpaceStationSession) CampaignFeeConfigs(arg0 *big.Int) (struct {
	Erc20       common.Address
	Erc20Fee    *big.Int
	PlatformFee *big.Int
}, error) {
	return _SpaceStation.Contract.CampaignFeeConfigs(&_SpaceStation.CallOpts, arg0)
}

// CampaignFeeConfigs is a free data retrieval call binding the contract method 0xe32ead2d.
//
// Solidity: function campaignFeeConfigs(uint256 ) view returns(address erc20, uint256 erc20Fee, uint256 platformFee)
func (_SpaceStation *SpaceStationCallerSession) CampaignFeeConfigs(arg0 *big.Int) (struct {
	Erc20       common.Address
	Erc20Fee    *big.Int
	PlatformFee *big.Int
}, error) {
	return _SpaceStation.Contract.CampaignFeeConfigs(&_SpaceStation.CallOpts, arg0)
}

// CampaignSetter is a free data retrieval call binding the contract method 0x15171d06.
//
// Solidity: function campaignSetter() view returns(address)
func (_SpaceStation *SpaceStationCaller) CampaignSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpaceStation.contract.Call(opts, &out, "campaignSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CampaignSetter is a free data retrieval call binding the contract method 0x15171d06.
//
// Solidity: function campaignSetter() view returns(address)
func (_SpaceStation *SpaceStationSession) CampaignSetter() (common.Address, error) {
	return _SpaceStation.Contract.CampaignSetter(&_SpaceStation.CallOpts)
}

// CampaignSetter is a free data retrieval call binding the contract method 0x15171d06.
//
// Solidity: function campaignSetter() view returns(address)
func (_SpaceStation *SpaceStationCallerSession) CampaignSetter() (common.Address, error) {
	return _SpaceStation.Contract.CampaignSetter(&_SpaceStation.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0xc37c656d.
//
// Solidity: function caps(uint256 ) view returns(uint256)
func (_SpaceStation *SpaceStationCaller) Caps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpaceStation.contract.Call(opts, &out, "caps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Caps is a free data retrieval call binding the contract method 0xc37c656d.
//
// Solidity: function caps(uint256 ) view returns(uint256)
func (_SpaceStation *SpaceStationSession) Caps(arg0 *big.Int) (*big.Int, error) {
	return _SpaceStation.Contract.Caps(&_SpaceStation.CallOpts, arg0)
}

// Caps is a free data retrieval call binding the contract method 0xc37c656d.
//
// Solidity: function caps(uint256 ) view returns(uint256)
func (_SpaceStation *SpaceStationCallerSession) Caps(arg0 *big.Int) (*big.Int, error) {
	return _SpaceStation.Contract.Caps(&_SpaceStation.CallOpts, arg0)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 ) view returns(bool)
func (_SpaceStation *SpaceStationCaller) IsClaimed(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _SpaceStation.contract.Call(opts, &out, "isClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 ) view returns(bool)
func (_SpaceStation *SpaceStationSession) IsClaimed(arg0 *big.Int) (bool, error) {
	return _SpaceStation.Contract.IsClaimed(&_SpaceStation.CallOpts, arg0)
}

// IsClaimed is a free data retrieval call binding the contract method 0x9e34070f.
//
// Solidity: function isClaimed(uint256 ) view returns(bool)
func (_SpaceStation *SpaceStationCallerSession) IsClaimed(arg0 *big.Int) (bool, error) {
	return _SpaceStation.Contract.IsClaimed(&_SpaceStation.CallOpts, arg0)
}

// NumClaimed is a free data retrieval call binding the contract method 0x6b5da37b.
//
// Solidity: function numClaimed(uint256 ) view returns(uint256)
func (_SpaceStation *SpaceStationCaller) NumClaimed(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpaceStation.contract.Call(opts, &out, "numClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumClaimed is a free data retrieval call binding the contract method 0x6b5da37b.
//
// Solidity: function numClaimed(uint256 ) view returns(uint256)
func (_SpaceStation *SpaceStationSession) NumClaimed(arg0 *big.Int) (*big.Int, error) {
	return _SpaceStation.Contract.NumClaimed(&_SpaceStation.CallOpts, arg0)
}

// NumClaimed is a free data retrieval call binding the contract method 0x6b5da37b.
//
// Solidity: function numClaimed(uint256 ) view returns(uint256)
func (_SpaceStation *SpaceStationCallerSession) NumClaimed(arg0 *big.Int) (*big.Int, error) {
	return _SpaceStation.Contract.NumClaimed(&_SpaceStation.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpaceStation *SpaceStationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpaceStation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpaceStation *SpaceStationSession) Owner() (common.Address, error) {
	return _SpaceStation.Contract.Owner(&_SpaceStation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SpaceStation *SpaceStationCallerSession) Owner() (common.Address, error) {
	return _SpaceStation.Contract.Owner(&_SpaceStation.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_SpaceStation *SpaceStationCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpaceStation.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_SpaceStation *SpaceStationSession) Signer() (common.Address, error) {
	return _SpaceStation.Contract.Signer(&_SpaceStation.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_SpaceStation *SpaceStationCallerSession) Signer() (common.Address, error) {
	return _SpaceStation.Contract.Signer(&_SpaceStation.CallOpts)
}

// TreasuryManager is a free data retrieval call binding the contract method 0x3cea70d9.
//
// Solidity: function treasuryManager() view returns(address)
func (_SpaceStation *SpaceStationCaller) TreasuryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SpaceStation.contract.Call(opts, &out, "treasuryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryManager is a free data retrieval call binding the contract method 0x3cea70d9.
//
// Solidity: function treasuryManager() view returns(address)
func (_SpaceStation *SpaceStationSession) TreasuryManager() (common.Address, error) {
	return _SpaceStation.Contract.TreasuryManager(&_SpaceStation.CallOpts)
}

// TreasuryManager is a free data retrieval call binding the contract method 0x3cea70d9.
//
// Solidity: function treasuryManager() view returns(address)
func (_SpaceStation *SpaceStationCallerSession) TreasuryManager() (common.Address, error) {
	return _SpaceStation.Contract.TreasuryManager(&_SpaceStation.CallOpts)
}

// ActivateCampaign is a paid mutator transaction binding the contract method 0x1eccf102.
//
// Solidity: function activateCampaign(uint256 _pid, address _erc20, uint256 _erc20Fee, uint256 _platformFee, uint256 cap) returns()
func (_SpaceStation *SpaceStationTransactor) ActivateCampaign(opts *bind.TransactOpts, _pid *big.Int, _erc20 common.Address, _erc20Fee *big.Int, _platformFee *big.Int, cap *big.Int) (*types.Transaction, error) {
	return _SpaceStation.contract.Transact(opts, "activateCampaign", _pid, _erc20, _erc20Fee, _platformFee, cap)
}

// ActivateCampaign is a paid mutator transaction binding the contract method 0x1eccf102.
//
// Solidity: function activateCampaign(uint256 _pid, address _erc20, uint256 _erc20Fee, uint256 _platformFee, uint256 cap) returns()
func (_SpaceStation *SpaceStationSession) ActivateCampaign(_pid *big.Int, _erc20 common.Address, _erc20Fee *big.Int, _platformFee *big.Int, cap *big.Int) (*types.Transaction, error) {
	return _SpaceStation.Contract.ActivateCampaign(&_SpaceStation.TransactOpts, _pid, _erc20, _erc20Fee, _platformFee, cap)
}

// ActivateCampaign is a paid mutator transaction binding the contract method 0x1eccf102.
//
// Solidity: function activateCampaign(uint256 _pid, address _erc20, uint256 _erc20Fee, uint256 _platformFee, uint256 cap) returns()
func (_SpaceStation *SpaceStationTransactorSession) ActivateCampaign(_pid *big.Int, _erc20 common.Address, _erc20Fee *big.Int, _platformFee *big.Int, cap *big.Int) (*types.Transaction, error) {
	return _SpaceStation.Contract.ActivateCampaign(&_SpaceStation.TransactOpts, _pid, _erc20, _erc20Fee, _platformFee, cap)
}

// Claim is a paid mutator transaction binding the contract method 0xe164adef.
//
// Solidity: function claim(uint256 _pid, address _nft, uint256 _dummyId, uint256 _level, address _mintTo, address _royaltyReceiver, uint96 _royaltyFeeNumerator, bool _canTranster, bytes _signature) returns(uint256)
func (_SpaceStation *SpaceStationTransactor) Claim(opts *bind.TransactOpts, _pid *big.Int, _nft common.Address, _dummyId *big.Int, _level *big.Int, _mintTo common.Address, _royaltyReceiver common.Address, _royaltyFeeNumerator *big.Int, _canTranster bool, _signature []byte) (*types.Transaction, error) {
	return _SpaceStation.contract.Transact(opts, "claim", _pid, _nft, _dummyId, _level, _mintTo, _royaltyReceiver, _royaltyFeeNumerator, _canTranster, _signature)
}

// Claim is a paid mutator transaction binding the contract method 0xe164adef.
//
// Solidity: function claim(uint256 _pid, address _nft, uint256 _dummyId, uint256 _level, address _mintTo, address _royaltyReceiver, uint96 _royaltyFeeNumerator, bool _canTranster, bytes _signature) returns(uint256)
func (_SpaceStation *SpaceStationSession) Claim(_pid *big.Int, _nft common.Address, _dummyId *big.Int, _level *big.Int, _mintTo common.Address, _royaltyReceiver common.Address, _royaltyFeeNumerator *big.Int, _canTranster bool, _signature []byte) (*types.Transaction, error) {
	return _SpaceStation.Contract.Claim(&_SpaceStation.TransactOpts, _pid, _nft, _dummyId, _level, _mintTo, _royaltyReceiver, _royaltyFeeNumerator, _canTranster, _signature)
}

// Claim is a paid mutator transaction binding the contract method 0xe164adef.
//
// Solidity: function claim(uint256 _pid, address _nft, uint256 _dummyId, uint256 _level, address _mintTo, address _royaltyReceiver, uint96 _royaltyFeeNumerator, bool _canTranster, bytes _signature) returns(uint256)
func (_SpaceStation *SpaceStationTransactorSession) Claim(_pid *big.Int, _nft common.Address, _dummyId *big.Int, _level *big.Int, _mintTo common.Address, _royaltyReceiver common.Address, _royaltyFeeNumerator *big.Int, _canTranster bool, _signature []byte) (*types.Transaction, error) {
	return _SpaceStation.Contract.Claim(&_SpaceStation.TransactOpts, _pid, _nft, _dummyId, _level, _mintTo, _royaltyReceiver, _royaltyFeeNumerator, _canTranster, _signature)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _signer, address _campaignSetter, address _treasuryManager) returns()
func (_SpaceStation *SpaceStationTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _signer common.Address, _campaignSetter common.Address, _treasuryManager common.Address) (*types.Transaction, error) {
	return _SpaceStation.contract.Transact(opts, "initialize", _owner, _signer, _campaignSetter, _treasuryManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _signer, address _campaignSetter, address _treasuryManager) returns()
func (_SpaceStation *SpaceStationSession) Initialize(_owner common.Address, _signer common.Address, _campaignSetter common.Address, _treasuryManager common.Address) (*types.Transaction, error) {
	return _SpaceStation.Contract.Initialize(&_SpaceStation.TransactOpts, _owner, _signer, _campaignSetter, _treasuryManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _owner, address _signer, address _campaignSetter, address _treasuryManager) returns()
func (_SpaceStation *SpaceStationTransactorSession) Initialize(_owner common.Address, _signer common.Address, _campaignSetter common.Address, _treasuryManager common.Address) (*types.Transaction, error) {
	return _SpaceStation.Contract.Initialize(&_SpaceStation.TransactOpts, _owner, _signer, _campaignSetter, _treasuryManager)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SpaceStation *SpaceStationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpaceStation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SpaceStation *SpaceStationSession) RenounceOwnership() (*types.Transaction, error) {
	return _SpaceStation.Contract.RenounceOwnership(&_SpaceStation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SpaceStation *SpaceStationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SpaceStation.Contract.RenounceOwnership(&_SpaceStation.TransactOpts)
}

// SetManager is a paid mutator transaction binding the contract method 0xf6a62096.
//
// Solidity: function setManager(address _signer, address _campaignSetter, address _treasuryManager) returns()
func (_SpaceStation *SpaceStationTransactor) SetManager(opts *bind.TransactOpts, _signer common.Address, _campaignSetter common.Address, _treasuryManager common.Address) (*types.Transaction, error) {
	return _SpaceStation.contract.Transact(opts, "setManager", _signer, _campaignSetter, _treasuryManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xf6a62096.
//
// Solidity: function setManager(address _signer, address _campaignSetter, address _treasuryManager) returns()
func (_SpaceStation *SpaceStationSession) SetManager(_signer common.Address, _campaignSetter common.Address, _treasuryManager common.Address) (*types.Transaction, error) {
	return _SpaceStation.Contract.SetManager(&_SpaceStation.TransactOpts, _signer, _campaignSetter, _treasuryManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xf6a62096.
//
// Solidity: function setManager(address _signer, address _campaignSetter, address _treasuryManager) returns()
func (_SpaceStation *SpaceStationTransactorSession) SetManager(_signer common.Address, _campaignSetter common.Address, _treasuryManager common.Address) (*types.Transaction, error) {
	return _SpaceStation.Contract.SetManager(&_SpaceStation.TransactOpts, _signer, _campaignSetter, _treasuryManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SpaceStation *SpaceStationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SpaceStation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SpaceStation *SpaceStationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SpaceStation.Contract.TransferOwnership(&_SpaceStation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SpaceStation *SpaceStationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SpaceStation.Contract.TransferOwnership(&_SpaceStation.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SpaceStation *SpaceStationTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpaceStation.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SpaceStation *SpaceStationSession) Receive() (*types.Transaction, error) {
	return _SpaceStation.Contract.Receive(&_SpaceStation.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SpaceStation *SpaceStationTransactorSession) Receive() (*types.Transaction, error) {
	return _SpaceStation.Contract.Receive(&_SpaceStation.TransactOpts)
}

// SpaceStationEventActivateCampaignIterator is returned from FilterEventActivateCampaign and is used to iterate over the raw logs and unpacked data for EventActivateCampaign events raised by the SpaceStation contract.
type SpaceStationEventActivateCampaignIterator struct {
	Event *SpaceStationEventActivateCampaign // Event containing the contract specifics and raw log

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
func (it *SpaceStationEventActivateCampaignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpaceStationEventActivateCampaign)
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
		it.Event = new(SpaceStationEventActivateCampaign)
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
func (it *SpaceStationEventActivateCampaignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpaceStationEventActivateCampaignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpaceStationEventActivateCampaign represents a EventActivateCampaign event raised by the SpaceStation contract.
type SpaceStationEventActivateCampaign struct {
	Pid         *big.Int
	Erc20       common.Address
	Erc20Fee    *big.Int
	PlatformFee *big.Int
	Cap         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEventActivateCampaign is a free log retrieval operation binding the contract event 0xeea389ac416aa738ce024ddbda928dd6734088c64666196949c5465fc811dc5d.
//
// Solidity: event EventActivateCampaign(uint256 pid, address erc20, uint256 erc20Fee, uint256 platformFee, uint256 cap)
func (_SpaceStation *SpaceStationFilterer) FilterEventActivateCampaign(opts *bind.FilterOpts) (*SpaceStationEventActivateCampaignIterator, error) {

	logs, sub, err := _SpaceStation.contract.FilterLogs(opts, "EventActivateCampaign")
	if err != nil {
		return nil, err
	}
	return &SpaceStationEventActivateCampaignIterator{contract: _SpaceStation.contract, event: "EventActivateCampaign", logs: logs, sub: sub}, nil
}

// WatchEventActivateCampaign is a free log subscription operation binding the contract event 0xeea389ac416aa738ce024ddbda928dd6734088c64666196949c5465fc811dc5d.
//
// Solidity: event EventActivateCampaign(uint256 pid, address erc20, uint256 erc20Fee, uint256 platformFee, uint256 cap)
func (_SpaceStation *SpaceStationFilterer) WatchEventActivateCampaign(opts *bind.WatchOpts, sink chan<- *SpaceStationEventActivateCampaign) (event.Subscription, error) {

	logs, sub, err := _SpaceStation.contract.WatchLogs(opts, "EventActivateCampaign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpaceStationEventActivateCampaign)
				if err := _SpaceStation.contract.UnpackLog(event, "EventActivateCampaign", log); err != nil {
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

// ParseEventActivateCampaign is a log parse operation binding the contract event 0xeea389ac416aa738ce024ddbda928dd6734088c64666196949c5465fc811dc5d.
//
// Solidity: event EventActivateCampaign(uint256 pid, address erc20, uint256 erc20Fee, uint256 platformFee, uint256 cap)
func (_SpaceStation *SpaceStationFilterer) ParseEventActivateCampaign(log types.Log) (*SpaceStationEventActivateCampaign, error) {
	event := new(SpaceStationEventActivateCampaign)
	if err := _SpaceStation.contract.UnpackLog(event, "EventActivateCampaign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpaceStationEventClaimIterator is returned from FilterEventClaim and is used to iterate over the raw logs and unpacked data for EventClaim events raised by the SpaceStation contract.
type SpaceStationEventClaimIterator struct {
	Event *SpaceStationEventClaim // Event containing the contract specifics and raw log

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
func (it *SpaceStationEventClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpaceStationEventClaim)
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
		it.Event = new(SpaceStationEventClaim)
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
func (it *SpaceStationEventClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpaceStationEventClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpaceStationEventClaim represents a EventClaim event raised by the SpaceStation contract.
type SpaceStationEventClaim struct {
	NftID       *big.Int
	Pid         *big.Int
	DummyId     *big.Int
	Level       *big.Int
	Nft         common.Address
	MintTo      common.Address
	CanTranster bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEventClaim is a free log retrieval operation binding the contract event 0x4acbe73b979f0a6452384be471194eef86ad06759695d9d54c94191fb5f29e8c.
//
// Solidity: event EventClaim(uint256 nftID, uint256 _pid, uint256 _dummyId, uint256 _level, address _nft, address _mintTo, bool _canTranster)
func (_SpaceStation *SpaceStationFilterer) FilterEventClaim(opts *bind.FilterOpts) (*SpaceStationEventClaimIterator, error) {

	logs, sub, err := _SpaceStation.contract.FilterLogs(opts, "EventClaim")
	if err != nil {
		return nil, err
	}
	return &SpaceStationEventClaimIterator{contract: _SpaceStation.contract, event: "EventClaim", logs: logs, sub: sub}, nil
}

// WatchEventClaim is a free log subscription operation binding the contract event 0x4acbe73b979f0a6452384be471194eef86ad06759695d9d54c94191fb5f29e8c.
//
// Solidity: event EventClaim(uint256 nftID, uint256 _pid, uint256 _dummyId, uint256 _level, address _nft, address _mintTo, bool _canTranster)
func (_SpaceStation *SpaceStationFilterer) WatchEventClaim(opts *bind.WatchOpts, sink chan<- *SpaceStationEventClaim) (event.Subscription, error) {

	logs, sub, err := _SpaceStation.contract.WatchLogs(opts, "EventClaim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpaceStationEventClaim)
				if err := _SpaceStation.contract.UnpackLog(event, "EventClaim", log); err != nil {
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

// ParseEventClaim is a log parse operation binding the contract event 0x4acbe73b979f0a6452384be471194eef86ad06759695d9d54c94191fb5f29e8c.
//
// Solidity: event EventClaim(uint256 nftID, uint256 _pid, uint256 _dummyId, uint256 _level, address _nft, address _mintTo, bool _canTranster)
func (_SpaceStation *SpaceStationFilterer) ParseEventClaim(log types.Log) (*SpaceStationEventClaim, error) {
	event := new(SpaceStationEventClaim)
	if err := _SpaceStation.contract.UnpackLog(event, "EventClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpaceStationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SpaceStation contract.
type SpaceStationOwnershipTransferredIterator struct {
	Event *SpaceStationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SpaceStationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpaceStationOwnershipTransferred)
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
		it.Event = new(SpaceStationOwnershipTransferred)
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
func (it *SpaceStationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpaceStationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpaceStationOwnershipTransferred represents a OwnershipTransferred event raised by the SpaceStation contract.
type SpaceStationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SpaceStation *SpaceStationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SpaceStationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SpaceStation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SpaceStationOwnershipTransferredIterator{contract: _SpaceStation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SpaceStation *SpaceStationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SpaceStationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SpaceStation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpaceStationOwnershipTransferred)
				if err := _SpaceStation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SpaceStation *SpaceStationFilterer) ParseOwnershipTransferred(log types.Log) (*SpaceStationOwnershipTransferred, error) {
	event := new(SpaceStationOwnershipTransferred)
	if err := _SpaceStation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpaceStationUpdateManagerIterator is returned from FilterUpdateManager and is used to iterate over the raw logs and unpacked data for UpdateManager events raised by the SpaceStation contract.
type SpaceStationUpdateManagerIterator struct {
	Event *SpaceStationUpdateManager // Event containing the contract specifics and raw log

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
func (it *SpaceStationUpdateManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpaceStationUpdateManager)
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
		it.Event = new(SpaceStationUpdateManager)
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
func (it *SpaceStationUpdateManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpaceStationUpdateManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpaceStationUpdateManager represents a UpdateManager event raised by the SpaceStation contract.
type SpaceStationUpdateManager struct {
	Signer          common.Address
	CampaignSetter  common.Address
	TreasuryManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateManager is a free log retrieval operation binding the contract event 0xea28c9444370d269c92153ac32abd8ab32f1f963064504e4c41c2b6b72684094.
//
// Solidity: event UpdateManager(address signer, address campaignSetter, address treasuryManager)
func (_SpaceStation *SpaceStationFilterer) FilterUpdateManager(opts *bind.FilterOpts) (*SpaceStationUpdateManagerIterator, error) {

	logs, sub, err := _SpaceStation.contract.FilterLogs(opts, "UpdateManager")
	if err != nil {
		return nil, err
	}
	return &SpaceStationUpdateManagerIterator{contract: _SpaceStation.contract, event: "UpdateManager", logs: logs, sub: sub}, nil
}

// WatchUpdateManager is a free log subscription operation binding the contract event 0xea28c9444370d269c92153ac32abd8ab32f1f963064504e4c41c2b6b72684094.
//
// Solidity: event UpdateManager(address signer, address campaignSetter, address treasuryManager)
func (_SpaceStation *SpaceStationFilterer) WatchUpdateManager(opts *bind.WatchOpts, sink chan<- *SpaceStationUpdateManager) (event.Subscription, error) {

	logs, sub, err := _SpaceStation.contract.WatchLogs(opts, "UpdateManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpaceStationUpdateManager)
				if err := _SpaceStation.contract.UnpackLog(event, "UpdateManager", log); err != nil {
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

// ParseUpdateManager is a log parse operation binding the contract event 0xea28c9444370d269c92153ac32abd8ab32f1f963064504e4c41c2b6b72684094.
//
// Solidity: event UpdateManager(address signer, address campaignSetter, address treasuryManager)
func (_SpaceStation *SpaceStationFilterer) ParseUpdateManager(log types.Log) (*SpaceStationUpdateManager, error) {
	event := new(SpaceStationUpdateManager)
	if err := _SpaceStation.contract.UnpackLog(event, "UpdateManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
