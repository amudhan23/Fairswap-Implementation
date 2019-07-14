// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Protocol2ABI is the input ABI used to generate the binding from.
const Protocol2ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"encrypted\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"complaintInside\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"receiverAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decrypted\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiverGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashedEncryptedFileChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"encryptedtext\",\"type\":\"bytes\"}],\"name\":\"Dec\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashedFileChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"noOfFileChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\"}],\"name\":\"keyReveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"EncryptedFileChunk\",\"type\":\"bytes\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"complain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"senderGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashDecrypted\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_hashedFileChunks\",\"type\":\"bytes32[]\"},{\"name\":\"_hashedEncryptedFileChunks\",\"type\":\"bytes32[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// Protocol2Bin is the compiled bytecode used for deploying new contracts.
const Protocol2Bin = `0x60806040819052600d805460ff1916905562001240388190039081908339810180604052620000329190810190620002f3565b600a8054600160a060020a0319163317905580518251146200005357600080fd5b8151600381905560408051828152602080840282010190915290801562000084578160200160208202803883390190505b5080516200009b91600491602090910190620001e5565b5060005b600354811015620000e8578281815181101515620000b957fe5b90602001906020020151600482815481101515620000d357fe5b6000918252602090912001556001016200009f565b5060035460405190808252806020026020018201604052801562000116578160200160208202803883390190505b5080516200012d91600591602090910190620001e5565b5060005b6003548110156200017a5781818151811015156200014b57fe5b906020019060200201516005828154811015156200016557fe5b60009182526020909120015560010162000131565b506000839055620001936401000000006200019c810204565b505050620003bf565b600d5460ff166006811115620001ae57fe5b6001016006811115620001bd57fe5b600d805460ff19166001836006811115620001d457fe5b0217905550600a4204601401600c55565b82805482825590600052602060002090810192821562000223579160200282015b828111156200022357825182559160200191906001019062000206565b506200023192915062000235565b5090565b6200025291905b808211156200023157600081556001016200023c565b90565b6000601f820183136200026757600080fd5b81516200027e62000278826200039e565b62000377565b91508181835260208401935060208101905083856020840282011115620002a457600080fd5b60005b83811015620002d45781620002bd8882620002de565b8452506020928301929190910190600101620002a7565b5050505092915050565b6000620002ec825162000252565b9392505050565b6000806000606084860312156200030957600080fd5b6000620003178686620002de565b93505060208401516001604060020a038111156200033457600080fd5b620003428682870162000255565b92505060408401516001604060020a038111156200035f57600080fd5b6200036d8682870162000255565b9150509250925092565b6040518181016001604060020a03811182821017156200039657600080fd5b604052919050565b60006001604060020a03821115620003b557600080fd5b5060209081020190565b610e7180620003cf6000396000f3fe608060405260043610610147576000357c0100000000000000000000000000000000000000000000000000000000900480635a012304116100c8578063ab01b2a21161008c578063ab01b2a2146102da578063b1c9fe6e146102fa578063d541ebd01461031c578063ee3743ab1461033c578063f6fa345f14610351578063fb374a291461036657610147565b80635a0123041461025b57806370dea79a1461027b5780638b0ca1fe1461029057806395c6a800146102b0578063a035b1fe146102c557610147565b806339ac0e491161010f57806339ac0e49146101da5780633e37aeb6146101fc578063448a9fac1461021157806344d4fd19146102315780634fc847911461024657610147565b80630adf939b1461014c5780630ed17fd61461017757806316fed3e2146101995780632852b71c146101bb5780633943380c146101c5575b600080fd5b34801561015857600080fd5b5061016161037b565b60405161016e9190610d4c565b60405180910390f35b34801561018357600080fd5b5061018c610409565b60405161016e9190610d30565b3480156101a557600080fd5b506101ae610412565b60405161016e9190610d1c565b6101c3610421565b005b3480156101d157600080fd5b506101616104da565b3480156101e657600080fd5b506101ef610534565b60405161016e9190610d3e565b34801561020857600080fd5b506101c361053a565b34801561021d57600080fd5b506101ef61022c366004610c9f565b6105e7565b34801561023d57600080fd5b506101ef610606565b34801561025257600080fd5b506101ae61060c565b34801561026757600080fd5b50610161610276366004610c11565b61061b565b34801561028757600080fd5b506101ef610711565b34801561029c57600080fd5b506101ef6102ab366004610c9f565b610717565b3480156102bc57600080fd5b506101ef610725565b3480156102d157600080fd5b506101ef61072b565b3480156102e657600080fd5b506101c36102f5366004610c11565b610731565b34801561030657600080fd5b5061030f6107a2565b60405161016e9190610d5d565b34801561032857600080fd5b5061018c610337366004610c4e565b6107ab565b34801561034857600080fd5b506101c3610968565b34801561035d57600080fd5b506101c36109ae565b34801561037257600080fd5b506101ef610a52565b6008805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104015780601f106103d657610100808354040283529160200191610401565b820191906000526020600020905b8154815290600101906020018083116103e457829003601f168201915b505050505081565b60065460ff1681565b600b54600160a060020a031681565b33600180600d5460ff16600681111561043657fe5b1461044057600080fd5b600c54600a42041061045157600080fd5b33600160a060020a0383161461046657600080fd5b600b805473ffffffffffffffffffffffffffffffffffffffff1916331790556000543410156104ce5760405133903480156108fc02916000818181858888f193505050501580156104bb573d6000803e3d6000fd5b50600d805460ff191660061790556104d6565b6104d6610968565b5050565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104015780601f106103d657610100808354040283529160200191610401565b60075481565b600b54600160a060020a0316600280600d5460ff16600681111561055a57fe5b1461056457600080fd5b600c54600a42041061057557600080fd5b33600160a060020a0383161461058a57600080fd5b600c54600a42041161059b57600080fd5b600b5460008054604051600160a060020a039093169281156108fc0292818181858888f193505050501580156105d5573d6000803e3d6000fd5b5050600d805460ff1916600517905550565b60058054829081106105f557fe5b600091825260209091200154905081565b60025481565b600a54600160a060020a031681565b60608082516040519080825280601f01601f19166020018201604052801561064a576020820181803883390190505b50905060005b835181101561070a57838181518110151561066757fe5b90602001015160f860020a900460f860020a02600182815460018160011615610100020316600290048110151561069a57fe5b8154600116156106b95790600052602060002090602091828204019190065b9054901a60f860020a021882828151811015156106d257fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600101610650565b5092915050565b600c5481565b60048054829081106105f557fe5b60035481565b60005481565b600a54600160a060020a0316600280600d5460ff16600681111561075157fe5b1461075b57600080fd5b600c54600a42041061076c57600080fd5b33600160a060020a0383161461078157600080fd5b8251610794906001906020860190610b0f565b5061079d610968565b505050565b600d5460ff1681565b600b54600090600160a060020a0316600380600d5460ff1660068111156107ce57fe5b146107d857600080fd5b600c54600a4204106107e957600080fd5b33600160a060020a038316146107fe57600080fd5b6006805460ff19166001179055845161081e906008906020880190610b0f565b5061082885610a58565b600755600580548590811061083957fe5b60009182526020909120015460095561085185610a58565b600580548690811061085f57fe5b90600052602060002001541415156108ca57600d80546006919060ff19166001835b0217905550600a5460008054604051600160a060020a039093169281156108fc0292818181858888f193505050501580156108c0573d6000803e3d6000fd5b5060009250610960565b6108db6108d68661061b565b610a58565b60048054869081106108e957fe5b906000526020600020015414151561094c57600d805460ff19166005179055600b5460008054604051600160a060020a03909316926108fc82150292818181858888f19350505050158015610942573d6000803e3d6000fd5b5060019250610960565b600d80546006919060ff1916600183610881565b505092915050565b600d5460ff16600681111561097957fe5b600101600681111561098757fe5b600d805460ff1916600183600681111561099d57fe5b0217905550600a4204601401600c55565b600a54600160a060020a0316600380600d5460ff1660068111156109ce57fe5b146109d857600080fd5b600c54600a4204106109e957600080fd5b33600160a060020a038316146109fe57600080fd5b600c54600a420411610a0f57600080fd5b600a5460008054604051600160a060020a039093169281156108fc0292818181858888f19350505050158015610a49573d6000803e3d6000fd5b506104d6610968565b60095481565b600080606083516040519080825280601f01601f191660200182016040528015610a89576020820181803883390190505b50905060005b8451811015610aff578481815181101515610aa657fe5b90602001015160f860020a900460f860020a028282815181101515610ac757fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600101610a8f565b5080516020909101209392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b5057805160ff1916838001178555610b7d565b82800160010185558215610b7d579182015b82811115610b7d578251825591602001919060010190610b62565b50610b89929150610b8d565b5090565b610ba791905b80821115610b895760008155600101610b93565b90565b6000601f82018313610bbb57600080fd5b8135610bce610bc982610d92565b610d6b565b91508082526020830160208301858383011115610bea57600080fd5b610bf5838284610df1565b50505092915050565b6000610c0a8235610ba7565b9392505050565b600060208284031215610c2357600080fd5b813567ffffffffffffffff811115610c3a57600080fd5b610c4684828501610baa565b949350505050565b60008060408385031215610c6157600080fd5b823567ffffffffffffffff811115610c7857600080fd5b610c8485828601610baa565b9250506020610c9585828601610bfe565b9150509250929050565b600060208284031215610cb157600080fd5b6000610c468484610bfe565b610cc681610dbe565b82525050565b610cc681610dc9565b610cc681610ba7565b6000610ce982610dba565b808452610cfd816020860160208601610dfd565b610d0681610e2d565b9093016020019392505050565b610cc681610de6565b60208101610d2a8284610cbd565b92915050565b60208101610d2a8284610ccc565b60208101610d2a8284610cd5565b60208082528101610c0a8184610cde565b60208101610d2a8284610d13565b60405181810167ffffffffffffffff81118282101715610d8a57600080fd5b604052919050565b600067ffffffffffffffff821115610da957600080fd5b506020601f91909101601f19160190565b5190565b6000610d2a82610dda565b151590565b600060078210610b8957fe5b600160a060020a031690565b6000610d2a82610dce565b82818337506000910152565b60005b83811015610e18578181015183820152602001610e00565b83811115610e27576000848401525b50505050565b601f01601f19169056fea265627a7a723058209a80b75ce506e43e4e856dfb8004ddc5edd77596c569de4d37e4909804fc3bcf6c6578706572696d656e74616cf50037`

// DeployProtocol2 deploys a new Ethereum contract, binding an instance of Protocol2 to it.
func DeployProtocol2(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int, _hashedFileChunks [][32]byte, _hashedEncryptedFileChunks [][32]byte) (common.Address, *types.Transaction, *Protocol2, error) {
	parsed, err := abi.JSON(strings.NewReader(Protocol2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Protocol2Bin), backend, _price, _hashedFileChunks, _hashedEncryptedFileChunks)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Protocol2{Protocol2Caller: Protocol2Caller{contract: contract}, Protocol2Transactor: Protocol2Transactor{contract: contract}, Protocol2Filterer: Protocol2Filterer{contract: contract}}, nil
}

// Protocol2 is an auto generated Go binding around an Ethereum contract.
type Protocol2 struct {
	Protocol2Caller     // Read-only binding to the contract
	Protocol2Transactor // Write-only binding to the contract
	Protocol2Filterer   // Log filterer for contract events
}

// Protocol2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Protocol2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Protocol2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Protocol2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Protocol2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Protocol2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Protocol2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Protocol2Session struct {
	Contract     *Protocol2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Protocol2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Protocol2CallerSession struct {
	Contract *Protocol2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Protocol2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Protocol2TransactorSession struct {
	Contract     *Protocol2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Protocol2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Protocol2Raw struct {
	Contract *Protocol2 // Generic contract binding to access the raw methods on
}

// Protocol2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Protocol2CallerRaw struct {
	Contract *Protocol2Caller // Generic read-only contract binding to access the raw methods on
}

// Protocol2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Protocol2TransactorRaw struct {
	Contract *Protocol2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocol2 creates a new instance of Protocol2, bound to a specific deployed contract.
func NewProtocol2(address common.Address, backend bind.ContractBackend) (*Protocol2, error) {
	contract, err := bindProtocol2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Protocol2{Protocol2Caller: Protocol2Caller{contract: contract}, Protocol2Transactor: Protocol2Transactor{contract: contract}, Protocol2Filterer: Protocol2Filterer{contract: contract}}, nil
}

// NewProtocol2Caller creates a new read-only instance of Protocol2, bound to a specific deployed contract.
func NewProtocol2Caller(address common.Address, caller bind.ContractCaller) (*Protocol2Caller, error) {
	contract, err := bindProtocol2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Protocol2Caller{contract: contract}, nil
}

// NewProtocol2Transactor creates a new write-only instance of Protocol2, bound to a specific deployed contract.
func NewProtocol2Transactor(address common.Address, transactor bind.ContractTransactor) (*Protocol2Transactor, error) {
	contract, err := bindProtocol2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Protocol2Transactor{contract: contract}, nil
}

// NewProtocol2Filterer creates a new log filterer instance of Protocol2, bound to a specific deployed contract.
func NewProtocol2Filterer(address common.Address, filterer bind.ContractFilterer) (*Protocol2Filterer, error) {
	contract, err := bindProtocol2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Protocol2Filterer{contract: contract}, nil
}

// bindProtocol2 binds a generic wrapper to an already deployed contract.
func bindProtocol2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Protocol2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol2 *Protocol2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol2.Contract.Protocol2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol2 *Protocol2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol2.Contract.Protocol2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol2 *Protocol2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol2.Contract.Protocol2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Protocol2 *Protocol2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Protocol2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Protocol2 *Protocol2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Protocol2 *Protocol2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Protocol2.Contract.contract.Transact(opts, method, params...)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_Protocol2 *Protocol2Caller) Now(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "Now")
	return *ret0, err
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_Protocol2 *Protocol2Session) Now() (*big.Int, error) {
	return _Protocol2.Contract.Now(&_Protocol2.CallOpts)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_Protocol2 *Protocol2CallerSession) Now() (*big.Int, error) {
	return _Protocol2.Contract.Now(&_Protocol2.CallOpts)
}

// ComplaintInside is a free data retrieval call binding the contract method 0x0ed17fd6.
//
// Solidity: function complaintInside() constant returns(bool)
func (_Protocol2 *Protocol2Caller) ComplaintInside(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "complaintInside")
	return *ret0, err
}

// ComplaintInside is a free data retrieval call binding the contract method 0x0ed17fd6.
//
// Solidity: function complaintInside() constant returns(bool)
func (_Protocol2 *Protocol2Session) ComplaintInside() (bool, error) {
	return _Protocol2.Contract.ComplaintInside(&_Protocol2.CallOpts)
}

// ComplaintInside is a free data retrieval call binding the contract method 0x0ed17fd6.
//
// Solidity: function complaintInside() constant returns(bool)
func (_Protocol2 *Protocol2CallerSession) ComplaintInside() (bool, error) {
	return _Protocol2.Contract.ComplaintInside(&_Protocol2.CallOpts)
}

// Decrypted is a free data retrieval call binding the contract method 0x39ac0e49.
//
// Solidity: function decrypted() constant returns(bytes32)
func (_Protocol2 *Protocol2Caller) Decrypted(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "decrypted")
	return *ret0, err
}

// Decrypted is a free data retrieval call binding the contract method 0x39ac0e49.
//
// Solidity: function decrypted() constant returns(bytes32)
func (_Protocol2 *Protocol2Session) Decrypted() ([32]byte, error) {
	return _Protocol2.Contract.Decrypted(&_Protocol2.CallOpts)
}

// Decrypted is a free data retrieval call binding the contract method 0x39ac0e49.
//
// Solidity: function decrypted() constant returns(bytes32)
func (_Protocol2 *Protocol2CallerSession) Decrypted() ([32]byte, error) {
	return _Protocol2.Contract.Decrypted(&_Protocol2.CallOpts)
}

// Encrypted is a free data retrieval call binding the contract method 0x0adf939b.
//
// Solidity: function encrypted() constant returns(bytes)
func (_Protocol2 *Protocol2Caller) Encrypted(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "encrypted")
	return *ret0, err
}

// Encrypted is a free data retrieval call binding the contract method 0x0adf939b.
//
// Solidity: function encrypted() constant returns(bytes)
func (_Protocol2 *Protocol2Session) Encrypted() ([]byte, error) {
	return _Protocol2.Contract.Encrypted(&_Protocol2.CallOpts)
}

// Encrypted is a free data retrieval call binding the contract method 0x0adf939b.
//
// Solidity: function encrypted() constant returns(bytes)
func (_Protocol2 *Protocol2CallerSession) Encrypted() ([]byte, error) {
	return _Protocol2.Contract.Encrypted(&_Protocol2.CallOpts)
}

// HashDecrypted is a free data retrieval call binding the contract method 0xfb374a29.
//
// Solidity: function hashDecrypted() constant returns(bytes32)
func (_Protocol2 *Protocol2Caller) HashDecrypted(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "hashDecrypted")
	return *ret0, err
}

// HashDecrypted is a free data retrieval call binding the contract method 0xfb374a29.
//
// Solidity: function hashDecrypted() constant returns(bytes32)
func (_Protocol2 *Protocol2Session) HashDecrypted() ([32]byte, error) {
	return _Protocol2.Contract.HashDecrypted(&_Protocol2.CallOpts)
}

// HashDecrypted is a free data retrieval call binding the contract method 0xfb374a29.
//
// Solidity: function hashDecrypted() constant returns(bytes32)
func (_Protocol2 *Protocol2CallerSession) HashDecrypted() ([32]byte, error) {
	return _Protocol2.Contract.HashDecrypted(&_Protocol2.CallOpts)
}

// HashedEncryptedFileChunks is a free data retrieval call binding the contract method 0x448a9fac.
//
// Solidity: function hashedEncryptedFileChunks(uint256 ) constant returns(bytes32)
func (_Protocol2 *Protocol2Caller) HashedEncryptedFileChunks(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "hashedEncryptedFileChunks", arg0)
	return *ret0, err
}

// HashedEncryptedFileChunks is a free data retrieval call binding the contract method 0x448a9fac.
//
// Solidity: function hashedEncryptedFileChunks(uint256 ) constant returns(bytes32)
func (_Protocol2 *Protocol2Session) HashedEncryptedFileChunks(arg0 *big.Int) ([32]byte, error) {
	return _Protocol2.Contract.HashedEncryptedFileChunks(&_Protocol2.CallOpts, arg0)
}

// HashedEncryptedFileChunks is a free data retrieval call binding the contract method 0x448a9fac.
//
// Solidity: function hashedEncryptedFileChunks(uint256 ) constant returns(bytes32)
func (_Protocol2 *Protocol2CallerSession) HashedEncryptedFileChunks(arg0 *big.Int) ([32]byte, error) {
	return _Protocol2.Contract.HashedEncryptedFileChunks(&_Protocol2.CallOpts, arg0)
}

// HashedFileChunks is a free data retrieval call binding the contract method 0x8b0ca1fe.
//
// Solidity: function hashedFileChunks(uint256 ) constant returns(bytes32)
func (_Protocol2 *Protocol2Caller) HashedFileChunks(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "hashedFileChunks", arg0)
	return *ret0, err
}

// HashedFileChunks is a free data retrieval call binding the contract method 0x8b0ca1fe.
//
// Solidity: function hashedFileChunks(uint256 ) constant returns(bytes32)
func (_Protocol2 *Protocol2Session) HashedFileChunks(arg0 *big.Int) ([32]byte, error) {
	return _Protocol2.Contract.HashedFileChunks(&_Protocol2.CallOpts, arg0)
}

// HashedFileChunks is a free data retrieval call binding the contract method 0x8b0ca1fe.
//
// Solidity: function hashedFileChunks(uint256 ) constant returns(bytes32)
func (_Protocol2 *Protocol2CallerSession) HashedFileChunks(arg0 *big.Int) ([32]byte, error) {
	return _Protocol2.Contract.HashedFileChunks(&_Protocol2.CallOpts, arg0)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes)
func (_Protocol2 *Protocol2Caller) Key(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "key")
	return *ret0, err
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes)
func (_Protocol2 *Protocol2Session) Key() ([]byte, error) {
	return _Protocol2.Contract.Key(&_Protocol2.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes)
func (_Protocol2 *Protocol2CallerSession) Key() ([]byte, error) {
	return _Protocol2.Contract.Key(&_Protocol2.CallOpts)
}

// NoOfFileChunks is a free data retrieval call binding the contract method 0x95c6a800.
//
// Solidity: function noOfFileChunks() constant returns(uint256)
func (_Protocol2 *Protocol2Caller) NoOfFileChunks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "noOfFileChunks")
	return *ret0, err
}

// NoOfFileChunks is a free data retrieval call binding the contract method 0x95c6a800.
//
// Solidity: function noOfFileChunks() constant returns(uint256)
func (_Protocol2 *Protocol2Session) NoOfFileChunks() (*big.Int, error) {
	return _Protocol2.Contract.NoOfFileChunks(&_Protocol2.CallOpts)
}

// NoOfFileChunks is a free data retrieval call binding the contract method 0x95c6a800.
//
// Solidity: function noOfFileChunks() constant returns(uint256)
func (_Protocol2 *Protocol2CallerSession) NoOfFileChunks() (*big.Int, error) {
	return _Protocol2.Contract.NoOfFileChunks(&_Protocol2.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Protocol2 *Protocol2Caller) Phase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "phase")
	return *ret0, err
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Protocol2 *Protocol2Session) Phase() (uint8, error) {
	return _Protocol2.Contract.Phase(&_Protocol2.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Protocol2 *Protocol2CallerSession) Phase() (uint8, error) {
	return _Protocol2.Contract.Phase(&_Protocol2.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Protocol2 *Protocol2Caller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Protocol2 *Protocol2Session) Price() (*big.Int, error) {
	return _Protocol2.Contract.Price(&_Protocol2.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Protocol2 *Protocol2CallerSession) Price() (*big.Int, error) {
	return _Protocol2.Contract.Price(&_Protocol2.CallOpts)
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() constant returns(address)
func (_Protocol2 *Protocol2Caller) ReceiverAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "receiverAddress")
	return *ret0, err
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() constant returns(address)
func (_Protocol2 *Protocol2Session) ReceiverAddress() (common.Address, error) {
	return _Protocol2.Contract.ReceiverAddress(&_Protocol2.CallOpts)
}

// ReceiverAddress is a free data retrieval call binding the contract method 0x16fed3e2.
//
// Solidity: function receiverAddress() constant returns(address)
func (_Protocol2 *Protocol2CallerSession) ReceiverAddress() (common.Address, error) {
	return _Protocol2.Contract.ReceiverAddress(&_Protocol2.CallOpts)
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() constant returns(address)
func (_Protocol2 *Protocol2Caller) SenderAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "senderAddress")
	return *ret0, err
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() constant returns(address)
func (_Protocol2 *Protocol2Session) SenderAddress() (common.Address, error) {
	return _Protocol2.Contract.SenderAddress(&_Protocol2.CallOpts)
}

// SenderAddress is a free data retrieval call binding the contract method 0x4fc84791.
//
// Solidity: function senderAddress() constant returns(address)
func (_Protocol2 *Protocol2CallerSession) SenderAddress() (common.Address, error) {
	return _Protocol2.Contract.SenderAddress(&_Protocol2.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Protocol2 *Protocol2Caller) Timeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "timeout")
	return *ret0, err
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Protocol2 *Protocol2Session) Timeout() (*big.Int, error) {
	return _Protocol2.Contract.Timeout(&_Protocol2.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Protocol2 *Protocol2CallerSession) Timeout() (*big.Int, error) {
	return _Protocol2.Contract.Timeout(&_Protocol2.CallOpts)
}

// Dec is a paid mutator transaction binding the contract method 0x5a012304.
//
// Solidity: function Dec(bytes encryptedtext) returns(bytes)
func (_Protocol2 *Protocol2Transactor) Dec(opts *bind.TransactOpts, encryptedtext []byte) (*types.Transaction, error) {
	return _Protocol2.contract.Transact(opts, "Dec", encryptedtext)
}

// Dec is a paid mutator transaction binding the contract method 0x5a012304.
//
// Solidity: function Dec(bytes encryptedtext) returns(bytes)
func (_Protocol2 *Protocol2Session) Dec(encryptedtext []byte) (*types.Transaction, error) {
	return _Protocol2.Contract.Dec(&_Protocol2.TransactOpts, encryptedtext)
}

// Dec is a paid mutator transaction binding the contract method 0x5a012304.
//
// Solidity: function Dec(bytes encryptedtext) returns(bytes)
func (_Protocol2 *Protocol2TransactorSession) Dec(encryptedtext []byte) (*types.Transaction, error) {
	return _Protocol2.Contract.Dec(&_Protocol2.TransactOpts, encryptedtext)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Protocol2 *Protocol2Transactor) Accept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol2.contract.Transact(opts, "accept")
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Protocol2 *Protocol2Session) Accept() (*types.Transaction, error) {
	return _Protocol2.Contract.Accept(&_Protocol2.TransactOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Protocol2 *Protocol2TransactorSession) Accept() (*types.Transaction, error) {
	return _Protocol2.Contract.Accept(&_Protocol2.TransactOpts)
}

// Complain is a paid mutator transaction binding the contract method 0xd541ebd0.
//
// Solidity: function complain(bytes EncryptedFileChunk, uint256 index) returns(bool)
func (_Protocol2 *Protocol2Transactor) Complain(opts *bind.TransactOpts, EncryptedFileChunk []byte, index *big.Int) (*types.Transaction, error) {
	return _Protocol2.contract.Transact(opts, "complain", EncryptedFileChunk, index)
}

// Complain is a paid mutator transaction binding the contract method 0xd541ebd0.
//
// Solidity: function complain(bytes EncryptedFileChunk, uint256 index) returns(bool)
func (_Protocol2 *Protocol2Session) Complain(EncryptedFileChunk []byte, index *big.Int) (*types.Transaction, error) {
	return _Protocol2.Contract.Complain(&_Protocol2.TransactOpts, EncryptedFileChunk, index)
}

// Complain is a paid mutator transaction binding the contract method 0xd541ebd0.
//
// Solidity: function complain(bytes EncryptedFileChunk, uint256 index) returns(bool)
func (_Protocol2 *Protocol2TransactorSession) Complain(EncryptedFileChunk []byte, index *big.Int) (*types.Transaction, error) {
	return _Protocol2.Contract.Complain(&_Protocol2.TransactOpts, EncryptedFileChunk, index)
}

// KeyReveal is a paid mutator transaction binding the contract method 0xab01b2a2.
//
// Solidity: function keyReveal(bytes _key) returns()
func (_Protocol2 *Protocol2Transactor) KeyReveal(opts *bind.TransactOpts, _key []byte) (*types.Transaction, error) {
	return _Protocol2.contract.Transact(opts, "keyReveal", _key)
}

// KeyReveal is a paid mutator transaction binding the contract method 0xab01b2a2.
//
// Solidity: function keyReveal(bytes _key) returns()
func (_Protocol2 *Protocol2Session) KeyReveal(_key []byte) (*types.Transaction, error) {
	return _Protocol2.Contract.KeyReveal(&_Protocol2.TransactOpts, _key)
}

// KeyReveal is a paid mutator transaction binding the contract method 0xab01b2a2.
//
// Solidity: function keyReveal(bytes _key) returns()
func (_Protocol2 *Protocol2TransactorSession) KeyReveal(_key []byte) (*types.Transaction, error) {
	return _Protocol2.Contract.KeyReveal(&_Protocol2.TransactOpts, _key)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Protocol2 *Protocol2Transactor) NextStage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol2.contract.Transact(opts, "nextStage")
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Protocol2 *Protocol2Session) NextStage() (*types.Transaction, error) {
	return _Protocol2.Contract.NextStage(&_Protocol2.TransactOpts)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Protocol2 *Protocol2TransactorSession) NextStage() (*types.Transaction, error) {
	return _Protocol2.Contract.NextStage(&_Protocol2.TransactOpts)
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Protocol2 *Protocol2Transactor) ReceiverGetEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol2.contract.Transact(opts, "receiverGetEther")
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Protocol2 *Protocol2Session) ReceiverGetEther() (*types.Transaction, error) {
	return _Protocol2.Contract.ReceiverGetEther(&_Protocol2.TransactOpts)
}

// ReceiverGetEther is a paid mutator transaction binding the contract method 0x3e37aeb6.
//
// Solidity: function receiverGetEther() returns()
func (_Protocol2 *Protocol2TransactorSession) ReceiverGetEther() (*types.Transaction, error) {
	return _Protocol2.Contract.ReceiverGetEther(&_Protocol2.TransactOpts)
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Protocol2 *Protocol2Transactor) SenderGetEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol2.contract.Transact(opts, "senderGetEther")
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Protocol2 *Protocol2Session) SenderGetEther() (*types.Transaction, error) {
	return _Protocol2.Contract.SenderGetEther(&_Protocol2.TransactOpts)
}

// SenderGetEther is a paid mutator transaction binding the contract method 0xf6fa345f.
//
// Solidity: function senderGetEther() returns()
func (_Protocol2 *Protocol2TransactorSession) SenderGetEther() (*types.Transaction, error) {
	return _Protocol2.Contract.SenderGetEther(&_Protocol2.TransactOpts)
}
