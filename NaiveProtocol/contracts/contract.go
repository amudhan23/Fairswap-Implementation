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
const Protocol2ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"receiverAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiverGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashedEncryptedFileChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"encryptedtext\",\"type\":\"bytes\"}],\"name\":\"Dec\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashedFileChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"noOfFileChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\"}],\"name\":\"keyReveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"EncryptedFileChunk\",\"type\":\"bytes\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"complain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"senderGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Nowy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_hashedFileChunks\",\"type\":\"bytes32[]\"},{\"name\":\"_hashedEncryptedFileChunks\",\"type\":\"bytes32[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// Protocol2Bin is the compiled bytecode used for deploying new contracts.
const Protocol2Bin = `0x608060408190526009805460ff1916905562001110388190039081908339810180604052620000329190810190620002f0565b60068054600160a060020a0319163317905580518251146200005357600080fd5b8151600381905560408051828152602080840282010190915290801562000084578160200160208202803883390190505b5080516200009b91600491602090910190620001e2565b5060005b600354811015620000e8578281815181101515620000b957fe5b90602001906020020151600482815481101515620000d357fe5b6000918252602090912001556001016200009f565b5060035460405190808252806020026020018201604052801562000116578160200160208202803883390190505b5080516200012d91600591602090910190620001e2565b5060005b6003548110156200017a5781818151811015156200014b57fe5b906020019060200201516005828154811015156200016557fe5b60009182526020909120015560010162000131565b506000839055620001936401000000006200019c810204565b505050620003bc565b60095460ff166004811115620001ae57fe5b6001016004811115620001bd57fe5b6009805460ff19166001836004811115620001d457fe5b0217905550603c4201600855565b82805482825590600052602060002090810192821562000220579160200282015b828111156200022057825182559160200191906001019062000203565b506200022e92915062000232565b5090565b6200024f91905b808211156200022e576000815560010162000239565b90565b6000601f820183136200026457600080fd5b81516200027b62000275826200039b565b62000374565b91508181835260208401935060208101905083856020840282011115620002a157600080fd5b60005b83811015620002d15781620002ba8882620002db565b8452506020928301929190910190600101620002a4565b5050505092915050565b6000620002e982516200024f565b9392505050565b6000806000606084860312156200030657600080fd5b6000620003148686620002db565b93505060208401516001604060020a038111156200033157600080fd5b6200033f8682870162000252565b92505060408401516001604060020a038111156200035c57600080fd5b6200036a8682870162000252565b9150509250925092565b6040518181016001604060020a03811182821017156200039357600080fd5b604052919050565b60006001604060020a03821115620003b257600080fd5b5060209081020190565b610d4480620003cc6000396000f3fe608060405260043610610126576000357c0100000000000000000000000000000000000000000000000000000000900480638b0ca1fe116100b2578063b1c9fe6e11610081578063b1c9fe6e1461028d578063d541ebd0146102af578063ee3743ab146102dc578063f6fa345f146102f1578063fd6fd6fa1461030657610126565b80638b0ca1fe1461022357806395c6a80014610243578063a035b1fe14610258578063ab01b2a21461026d57610126565b8063448a9fac116100f9578063448a9fac1461019757806344d4fd19146101c45780634fc84791146101d95780635a012304146101ee57806370dea79a1461020e57610126565b806316fed3e21461012b5780632852b71c146101565780633943380c146101605780633e37aeb614610182575b600080fd5b34801561013757600080fd5b5061014061031b565b60405161014d9190610bef565b60405180910390f35b61015e61032a565b005b34801561016c57600080fd5b506101756103d3565b60405161014d9190610c1f565b34801561018e57600080fd5b5061015e610460565b3480156101a357600080fd5b506101b76101b2366004610b72565b6104f9565b60405161014d9190610c11565b3480156101d057600080fd5b506101b7610518565b3480156101e557600080fd5b5061014061051e565b3480156101fa57600080fd5b50610175610209366004610ae4565b61052d565b34801561021a57600080fd5b506101b7610623565b34801561022f57600080fd5b506101b761023e366004610b72565b610629565b34801561024f57600080fd5b506101b7610637565b34801561026457600080fd5b506101b761063d565b34801561027957600080fd5b5061015e610288366004610ae4565b610643565b34801561029957600080fd5b506102a26106b1565b60405161014d9190610c30565b3480156102bb57600080fd5b506102cf6102ca366004610b21565b6106ba565b60405161014d9190610c03565b3480156102e857600080fd5b5061015e61085b565b3480156102fd57600080fd5b5061015e61089e565b34801561031257600080fd5b5061015e610925565b600754600160a060020a031681565b3360018060095460ff16600481111561033f57fe5b1461034957600080fd5b600854421061035757600080fd5b33600160a060020a0383161461036c57600080fd5b6000543410156103a85760405133903480156108fc02916000818181858888f193505050501580156103a2573d6000803e3d6000fd5b506103cf565b6007805473ffffffffffffffffffffffffffffffffffffffff1916331790556103cf61085b565b5050565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104585780601f1061042d57610100808354040283529160200191610458565b820191906000526020600020905b81548152906001019060200180831161043b57829003601f168201915b505050505081565b600754600160a060020a031660028060095460ff16600481111561048057fe5b1461048a57600080fd5b600854421161049857600080fd5b33600160a060020a038316146104ad57600080fd5b60075460008054604051600160a060020a039093169281156108fc0292818181858888f193505050501580156104e7573d6000803e3d6000fd5b50506009805460ff1916600417905550565b600580548290811061050757fe5b600091825260209091200154905081565b60025481565b600654600160a060020a031681565b60608082516040519080825280601f01601f19166020018201604052801561055c576020820181803883390190505b50905060005b835181101561061c57838181518110151561057957fe5b90602001015160f860020a900460f860020a0260018281546001816001161561010002031660029004811015156105ac57fe5b8154600116156105cb5790600052602060002090602091828204019190065b9054901a60f860020a021882828151811015156105e457fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600101610562565b5092915050565b60085481565b600480548290811061050757fe5b60035481565b60005481565b600654600160a060020a031660028060095460ff16600481111561066357fe5b1461066d57600080fd5b600854421061067b57600080fd5b33600160a060020a0383161461069057600080fd5b82516106a39060019060208601906109e2565b506106ac61085b565b505050565b60095460ff1681565b600754600090600160a060020a031660038060095460ff1660048111156106dd57fe5b146106e757600080fd5b60085442106106f557600080fd5b33600160a060020a0383161461070a57600080fd5b6107138561092b565b600580548690811061072157fe5b9060005260206000200154141515610784576009805460ff1916600417905560065460008054604051600160a060020a03909316926108fc82150292818181858888f1935050505015801561077a573d6000803e3d6000fd5b5060009250610853565b6107956107908661052d565b61092b565b60048054869081106107a357fe5b9060005260206000200154141515610806576009805460ff1916600417905560075460008054604051600160a060020a03909316926108fc82150292818181858888f193505050501580156107fc573d6000803e3d6000fd5b5060019250610853565b6009805460ff1916600417905560065460008054604051600160a060020a03909316926108fc82150292818181858888f1935050505015801561084d573d6000803e3d6000fd5b50600092505b505092915050565b60095460ff16600481111561086c57fe5b600101600481111561087a57fe5b6009805460ff1916600183600481111561089057fe5b0217905550603c4201600855565b600654600160a060020a031660038060095460ff1660048111156108be57fe5b146108c857600080fd5b60085442116108d657600080fd5b33600160a060020a038316146108eb57600080fd5b60065460008054604051600160a060020a039093169281156108fc0292818181858888f193505050501580156104e7573d6000803e3d6000fd5b42600255565b600080606083516040519080825280601f01601f19166020018201604052801561095c576020820181803883390190505b50905060005b84518110156109d257848181518110151561097957fe5b90602001015160f860020a900460f860020a02828281518110151561099a57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600101610962565b5080516020909101209392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a2357805160ff1916838001178555610a50565b82800160010185558215610a50579182015b82811115610a50578251825591602001919060010190610a35565b50610a5c929150610a60565b5090565b610a7a91905b80821115610a5c5760008155600101610a66565b90565b6000601f82018313610a8e57600080fd5b8135610aa1610a9c82610c65565b610c3e565b91508082526020830160208301858383011115610abd57600080fd5b610ac8838284610cc4565b50505092915050565b6000610add8235610a7a565b9392505050565b600060208284031215610af657600080fd5b813567ffffffffffffffff811115610b0d57600080fd5b610b1984828501610a7d565b949350505050565b60008060408385031215610b3457600080fd5b823567ffffffffffffffff811115610b4b57600080fd5b610b5785828601610a7d565b9250506020610b6885828601610ad1565b9150509250929050565b600060208284031215610b8457600080fd5b6000610b198484610ad1565b610b9981610c91565b82525050565b610b9981610c9c565b610b9981610a7a565b6000610bbc82610c8d565b808452610bd0816020860160208601610cd0565b610bd981610d00565b9093016020019392505050565b610b9981610cb9565b60208101610bfd8284610b90565b92915050565b60208101610bfd8284610b9f565b60208101610bfd8284610ba8565b60208082528101610add8184610bb1565b60208101610bfd8284610be6565b60405181810167ffffffffffffffff81118282101715610c5d57600080fd5b604052919050565b600067ffffffffffffffff821115610c7c57600080fd5b506020601f91909101601f19160190565b5190565b6000610bfd82610cad565b151590565b600060058210610a5c57fe5b600160a060020a031690565b6000610bfd82610ca1565b82818337506000910152565b60005b83811015610ceb578181015183820152602001610cd3565b83811115610cfa576000848401525b50505050565b601f01601f19169056fea265627a7a723058205e70e0c9e01f709d90fbec3f91933201dd076ef31e5578a7f3420539d509573c6c6578706572696d656e74616cf50037`

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

// Nowy is a paid mutator transaction binding the contract method 0xfd6fd6fa.
//
// Solidity: function Nowy() returns()
func (_Protocol2 *Protocol2Transactor) Nowy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Protocol2.contract.Transact(opts, "Nowy")
}

// Nowy is a paid mutator transaction binding the contract method 0xfd6fd6fa.
//
// Solidity: function Nowy() returns()
func (_Protocol2 *Protocol2Session) Nowy() (*types.Transaction, error) {
	return _Protocol2.Contract.Nowy(&_Protocol2.TransactOpts)
}

// Nowy is a paid mutator transaction binding the contract method 0xfd6fd6fa.
//
// Solidity: function Nowy() returns()
func (_Protocol2 *Protocol2TransactorSession) Nowy() (*types.Transaction, error) {
	return _Protocol2.Contract.Nowy(&_Protocol2.TransactOpts)
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
