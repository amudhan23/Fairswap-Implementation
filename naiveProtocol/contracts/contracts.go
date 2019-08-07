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
const Protocol2ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"receiverAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiverGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashedEncryptedFileChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"encryptedtext\",\"type\":\"bytes\"}],\"name\":\"Dec\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validComplain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hashedFileChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"noOfFileChunks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes\"}],\"name\":\"keyReveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"EncryptedFileChunk\",\"type\":\"bytes\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"complain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"senderGetEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Nowy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_hashedFileChunks\",\"type\":\"bytes32[]\"},{\"name\":\"_hashedEncryptedFileChunks\",\"type\":\"bytes32[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"}]"

// Protocol2Bin is the compiled bytecode used for deploying new contracts.
const Protocol2Bin = `0x608060408190526009805460ff191690556200116b388190039081908339810180604052620000329190810190620002f5565b6006805461010060a860020a031916336101000217905580518251146200005857600080fd5b8151600381905560408051828152602080840282010190915290801562000089578160200160208202803883390190505b508051620000a091600491602090910190620001e7565b5060005b600354811015620000ed578281815181101515620000be57fe5b90602001906020020151600482815481101515620000d857fe5b600091825260209091200155600101620000a4565b506003546040519080825280602002602001820160405280156200011b578160200160208202803883390190505b5080516200013291600591602090910190620001e7565b5060005b6003548110156200017f5781818151811015156200015057fe5b906020019060200201516005828154811015156200016a57fe5b60009182526020909120015560010162000136565b50600083905562000198640100000000620001a1810204565b505050620003c1565b60095460ff166004811115620001b357fe5b6001016004811115620001c257fe5b6009805460ff19166001836004811115620001d957fe5b021790555060144201600855565b82805482825590600052602060002090810192821562000225579160200282015b828111156200022557825182559160200191906001019062000208565b506200023392915062000237565b5090565b6200025491905b808211156200023357600081556001016200023e565b90565b6000601f820183136200026957600080fd5b8151620002806200027a82620003a0565b62000379565b91508181835260208401935060208101905083856020840282011115620002a657600080fd5b60005b83811015620002d65781620002bf8882620002e0565b8452506020928301929190910190600101620002a9565b5050505092915050565b6000620002ee825162000254565b9392505050565b6000806000606084860312156200030b57600080fd5b6000620003198686620002e0565b93505060208401516001604060020a038111156200033657600080fd5b620003448682870162000257565b92505060408401516001604060020a038111156200036157600080fd5b6200036f8682870162000257565b9150509250925092565b6040518181016001604060020a03811182821017156200039857600080fd5b604052919050565b60006001604060020a03821115620003b757600080fd5b5060209081020190565b610d9a80620003d16000396000f3fe608060405260043610610131576000357c01000000000000000000000000000000000000000000000000000000009004806376aa657f116100bd578063b1c9fe6e11610081578063b1c9fe6e146102ba578063d541ebd0146102dc578063ee3743ab146102fc578063f6fa345f14610311578063fd6fd6fa1461032657610131565b806376aa657f1461022e5780638b0ca1fe1461025057806395c6a80014610270578063a035b1fe14610285578063ab01b2a21461029a57610131565b8063448a9fac11610104578063448a9fac146101a257806344d4fd19146101cf5780634fc84791146101e45780635a012304146101f957806370dea79a1461021957610131565b806316fed3e2146101365780632852b71c146101615780633943380c1461016b5780633e37aeb61461018d575b600080fd5b34801561014257600080fd5b5061014b61033b565b6040516101589190610c45565b60405180910390f35b61016961034a565b005b34801561017757600080fd5b506101806103f3565b6040516101589190610c75565b34801561019957600080fd5b50610169610480565b3480156101ae57600080fd5b506101c26101bd366004610bc8565b610519565b6040516101589190610c67565b3480156101db57600080fd5b506101c2610538565b3480156101f057600080fd5b5061014b61053e565b34801561020557600080fd5b50610180610214366004610b3a565b610552565b34801561022557600080fd5b506101c2610648565b34801561023a57600080fd5b5061024361064e565b6040516101589190610c59565b34801561025c57600080fd5b506101c261026b366004610bc8565b610657565b34801561027c57600080fd5b506101c2610665565b34801561029157600080fd5b506101c261066b565b3480156102a657600080fd5b506101696102b5366004610b3a565b610671565b3480156102c657600080fd5b506102cf6106e4565b6040516101589190610c86565b3480156102e857600080fd5b506102436102f7366004610b77565b6106ed565b34801561030857600080fd5b506101696108a8565b34801561031d57600080fd5b506101696108eb565b34801561033257600080fd5b5061016961097b565b600754600160a060020a031681565b3360018060095460ff16600481111561035f57fe5b1461036957600080fd5b600854421061037757600080fd5b33600160a060020a0383161461038c57600080fd5b6000543410156103c85760405133903480156108fc02916000818181858888f193505050501580156103c2573d6000803e3d6000fd5b506103ef565b6007805473ffffffffffffffffffffffffffffffffffffffff1916331790556103ef6108a8565b5050565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104785780601f1061044d57610100808354040283529160200191610478565b820191906000526020600020905b81548152906001019060200180831161045b57829003601f168201915b505050505081565b600754600160a060020a031660028060095460ff1660048111156104a057fe5b146104aa57600080fd5b60085442116104b857600080fd5b33600160a060020a038316146104cd57600080fd5b60075460008054604051600160a060020a039093169281156108fc0292818181858888f19350505050158015610507573d6000803e3d6000fd5b50506009805460ff1916600417905550565b600580548290811061052757fe5b600091825260209091200154905081565b60025481565b6006546101009004600160a060020a031681565b60608082516040519080825280601f01601f191660200182016040528015610581576020820181803883390190505b50905060005b835181101561064157838181518110151561059e57fe5b90602001015160f860020a900460f860020a0260018281546001816001161561010002031660029004811015156105d157fe5b8154600116156105f05790600052602060002090602091828204019190065b9054901a60f860020a0218828281518110151561060957fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600101610587565b5092915050565b60085481565b60065460ff1681565b600480548290811061052757fe5b60035481565b60005481565b6006546101009004600160a060020a031660028060095460ff16600481111561069657fe5b146106a057600080fd5b60085442106106ae57600080fd5b33600160a060020a038316146106c357600080fd5b82516106d6906001906020860190610a38565b506106df6108a8565b505050565b60095460ff1681565b600754600090600160a060020a031660038060095460ff16600481111561071057fe5b1461071a57600080fd5b600854421061072857600080fd5b33600160a060020a0383161461073d57600080fd5b61074685610981565b600580548690811061075457fe5b90600052602060002001541415156107bd5760098054600460ff1990911617905560065460008054604051610100909304600160a060020a0316926108fc82150292818181858888f193505050501580156107b3573d6000803e3d6000fd5b50600092506108a0565b6107ce6107c986610552565b610981565b60048054869081106107dc57fe5b906000526020600020015414151561084d576009805460ff1916600417905560075460008054604051600160a060020a03909316926108fc82150292818181858888f19350505050158015610835573d6000803e3d6000fd5b506006805460ff1916600190811790915592506108a0565b60098054600460ff1990911617905560065460008054604051610100909304600160a060020a0316926108fc82150292818181858888f1935050505015801561089a573d6000803e3d6000fd5b50600092505b505092915050565b60095460ff1660048111156108b957fe5b60010160048111156108c757fe5b6009805460ff191660018360048111156108dd57fe5b021790555060144201600855565b6006546101009004600160a060020a031660038060095460ff16600481111561091057fe5b1461091a57600080fd5b600854421161092857600080fd5b33600160a060020a0383161461093d57600080fd5b60065460008054604051610100909304600160a060020a0316926108fc82150292818181858888f19350505050158015610507573d6000803e3d6000fd5b42600255565b600080606083516040519080825280601f01601f1916602001820160405280156109b2576020820181803883390190505b50905060005b8451811015610a285784818151811015156109cf57fe5b90602001015160f860020a900460f860020a0282828151811015156109f057fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506001016109b8565b5080516020909101209392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a7957805160ff1916838001178555610aa6565b82800160010185558215610aa6579182015b82811115610aa6578251825591602001919060010190610a8b565b50610ab2929150610ab6565b5090565b610ad091905b80821115610ab25760008155600101610abc565b90565b6000601f82018313610ae457600080fd5b8135610af7610af282610cbb565b610c94565b91508082526020830160208301858383011115610b1357600080fd5b610b1e838284610d1a565b50505092915050565b6000610b338235610ad0565b9392505050565b600060208284031215610b4c57600080fd5b813567ffffffffffffffff811115610b6357600080fd5b610b6f84828501610ad3565b949350505050565b60008060408385031215610b8a57600080fd5b823567ffffffffffffffff811115610ba157600080fd5b610bad85828601610ad3565b9250506020610bbe85828601610b27565b9150509250929050565b600060208284031215610bda57600080fd5b6000610b6f8484610b27565b610bef81610ce7565b82525050565b610bef81610cf2565b610bef81610ad0565b6000610c1282610ce3565b808452610c26816020860160208601610d26565b610c2f81610d56565b9093016020019392505050565b610bef81610d0f565b60208101610c538284610be6565b92915050565b60208101610c538284610bf5565b60208101610c538284610bfe565b60208082528101610b338184610c07565b60208101610c538284610c3c565b60405181810167ffffffffffffffff81118282101715610cb357600080fd5b604052919050565b600067ffffffffffffffff821115610cd257600080fd5b506020601f91909101601f19160190565b5190565b6000610c5382610d03565b151590565b600060058210610ab257fe5b600160a060020a031690565b6000610c5382610cf7565b82818337506000910152565b60005b83811015610d41578181015183820152602001610d29565b83811115610d50576000848401525b50505050565b601f01601f19169056fea265627a7a723058208c412637c80353d5da7b30a9ccb97e01d074dfc1f2d6a3e4d6b2a5d32e997cd36c6578706572696d656e74616cf50037`

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

// ValidComplain is a free data retrieval call binding the contract method 0x76aa657f.
//
// Solidity: function validComplain() constant returns(bool)
func (_Protocol2 *Protocol2Caller) ValidComplain(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Protocol2.contract.Call(opts, out, "validComplain")
	return *ret0, err
}

// ValidComplain is a free data retrieval call binding the contract method 0x76aa657f.
//
// Solidity: function validComplain() constant returns(bool)
func (_Protocol2 *Protocol2Session) ValidComplain() (bool, error) {
	return _Protocol2.Contract.ValidComplain(&_Protocol2.CallOpts)
}

// ValidComplain is a free data retrieval call binding the contract method 0x76aa657f.
//
// Solidity: function validComplain() constant returns(bool)
func (_Protocol2 *Protocol2CallerSession) ValidComplain() (bool, error) {
	return _Protocol2.Contract.ValidComplain(&_Protocol2.CallOpts)
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
