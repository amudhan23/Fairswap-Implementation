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

// FairswapABI is the input ABI used to generate the binding from.
const FairswapABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EncInputRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"circuitGatesOperation\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"key\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitmentofkey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"complaint\",\"type\":\"bytes32[5][]\"},{\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"complain\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderaddrs\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"complianBool\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"recieveraddrs\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CircuitMerkleRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"revealKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initializerecieveraddrs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"complaint\",\"type\":\"bytes32[5]\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"Mverify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_commitmentofkey\",\"type\":\"bytes32\"},{\"name\":\"_CircuitMerkleRoot\",\"type\":\"bytes32\"},{\"name\":\"_EncInputRoot\",\"type\":\"bytes32\"},{\"name\":\"circuitGatesOperationArray\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FairswapBin is the compiled bytecode used for deploying new contracts.
const FairswapBin = `0x6080604052600a805460ff191690553480156200001b57600080fd5b5060405162001033380380620010338339810180604052620000419190810190620001a8565b60078054600160a060020a0319163317905560028590556003849055600083815560018390555b8151811015620000a55781818151811015156200008157fe5b60209081029091018101516000838152600690925260409091205560010162000068565b50620000b9640100000000620000c4810204565b50505050506200028b565b600a5460ff166004811115620000d657fe5b6001016004811115620000e557fe5b600a805460ff19166001836004811115620000fc57fe5b021790555060b44201600955565b6000601f820183136200011c57600080fd5b8151620001336200012d8262000267565b62000240565b915081818352602084019350602081019050838560208402820111156200015957600080fd5b60005b8381101562000189578162000172888262000193565b84525060209283019291909101906001016200015c565b5050505092915050565b6000620001a1825162000288565b9392505050565b600080600080600060a08688031215620001c157600080fd5b6000620001cf888862000193565b9550506020620001e28882890162000193565b9450506040620001f58882890162000193565b9350506060620002088882890162000193565b92505060808601516001604060020a038111156200022557600080fd5b62000233888289016200010a565b9150509295509295909350565b6040518181016001604060020a03811182821017156200025f57600080fd5b604052919050565b60006001604060020a038211156200027e57600080fd5b5060209081020190565b90565b610d98806200029b6000396000f3fe60806040526004361061011b576000357c010000000000000000000000000000000000000000000000000000000090048063a035b1fe116100b2578063c811d63c11610081578063c811d63c14610264578063d654723614610279578063e0e5f68f14610299578063ee3743ab146102ae578063f94acc36146102c35761011b565b8063a035b1fe14610203578063a0fb101314610218578063a784df2a1461022d578063b1c9fe6e146102425761011b565b80634c68d890116100ee5780634c68d8901461018a57806370dea79a1461019f5780637896c64e146101b4578063945ef141146101e15761011b565b8063062a1398146101205780630b90457d1461014b5780632852b71c1461016b5780633943380c14610175575b600080fd5b34801561012c57600080fd5b506101356102e3565b6040516101429190610ca2565b60405180910390f35b34801561015757600080fd5b50610135610166366004610c36565b6102e9565b6101736102fb565b005b34801561018157600080fd5b50610135610363565b34801561019657600080fd5b50610135610369565b3480156101ab57600080fd5b5061013561036f565b3480156101c057600080fd5b506101d46101cf366004610b9d565b610375565b6040516101429190610c94565b3480156101ed57600080fd5b506101f6610526565b6040516101429190610c86565b34801561020f57600080fd5b50610135610535565b34801561022457600080fd5b506101d461053b565b34801561023957600080fd5b506101f6610544565b34801561024e57600080fd5b50610257610553565b6040516101429190610cb0565b34801561027057600080fd5b5061013561055c565b34801561028557600080fd5b50610173610294366004610c36565b610562565b3480156102a557600080fd5b506101736105c1565b3480156102ba57600080fd5b506101736105e2565b3480156102cf57600080fd5b506101d46102de366004610c06565b610625565b60015481565b60066020526000908152604090205481565b600854600160a060020a0316600180600a5460ff16600481111561031b57fe5b1461032557600080fd5b600954421061033357600080fd5b33600160a060020a0383161461034857600080fd5b60025434101561035757600080fd5b61035f6105e2565b5050565b60045481565b60035481565b60095481565b6000806000606060008086600081518110151561038e57fe5b90602001906020020151905060018851036040519080825280602002602001820160405280156103c8578160200160208202803883390190505b5092506104058860008151811015156103dd57fe5b906020019060200201518860008151811015156103f657fe5b90602001906020020151610625565b151561041057600080fd5b6006600088600081518110151561042357fe5b90602001906020020151815260200190815260200160002054935061046188600081518110151561045057fe5b602090810291909101015151610816565b915060005b60018951038110156104da5761049b898281518110151561048357fe5b9060200190602002015189838151811015156103f657fe5b15156104a657600080fd5b6104ba898260010181518110151561045057fe5b84828151811015156104c857fe5b60209081029091010152600101610466565b506104e584846108cf565b94508482141561050c5750506005805460ff19166001908117909155935061052092505050565b50506005805460ff19169055506000925050505b92915050565b600754600160a060020a031681565b60025481565b60055460ff1681565b600854600160a060020a031681565b600a5460ff1681565b60005481565b600754600160a060020a0316600280600a5460ff16600481111561058257fe5b1461058c57600080fd5b600954421061059a57600080fd5b33600160a060020a038316146105af57600080fd5b60048390556105bc6105e2565b505050565b6008805473ffffffffffffffffffffffffffffffffffffffff191633179055565b600a5460ff1660048111156105f357fe5b600101600481111561060157fe5b600a805460ff1916600183600481111561061757fe5b021790555060b44201600955565b815160408051818152606080820183526000939284928392600592909190602082018180388339019050509050600092505b6001820383101561080457600283900a87811604600114156107385760005b60208110156106d55788600185016005811061068e57fe5b6020020151816020811015156106a057fe5b1a60f860020a0282828151811015156106b557fe5b906020010190600160f860020a031916908160001a905350600101610676565b5060005b6020811015610726578581602081106106ee57fe5b1a60f860020a02828260200181518110151561070657fe5b906020010190600160f860020a031916908160001a9053506001016106d9565b508051602082012094508493506107f9565b60005b60208110156107855785816020811061075057fe5b1a60f860020a02828281518110151561076557fe5b906020010190600160f860020a031916908160001a90535060010161073b565b5060005b60208110156107eb578860018501600581106107a157fe5b6020020151816020811015156107b357fe5b1a60f860020a0282826020018151811015156107cb57fe5b906020010190600160f860020a031916908160001a905350600101610789565b508051602082012094508493505b600190920191610657565b60015485149550505050505092915050565b604080516021808252606082810190935260009291839182919060208201818038833901905050925060005b602081101561088e57600454816020811061085957fe5b1a60f860020a02848281518110151561086e57fe5b906020010190600160f860020a031916908160001a905350600101610842565b508251600090849060209081106108a157fe5b906020010190600160f860020a031916908160001a9053505081516020909201919091208318915050919050565b60008083600214156109d75760408051818152606081810183529160208201818038833901905050905060005b60208110156109605784600081518110151561091457fe5b906020019060200201518160208110151561092b57fe5b1a60f860020a02828281518110151561094057fe5b906020010190600160f860020a031916908160001a9053506001016108fc565b5060005b60208110156109cb5784600181518110151561097c57fe5b906020019060200201518160208110151561099357fe5b1a60f860020a0282826020018151811015156109ab57fe5b906020010190600160f860020a031916908160001a905350600101610964565b50805160209091012090505b9392505050565b6000601f820183136109ef57600080fd5b8135610a026109fd82610ce5565b610cbe565b915081818352602084019350602081019050838560a0840282011115610a2757600080fd5b60005b83811015610a555781610a3d8882610a5f565b84525060209092019160a09190910190600101610a2a565b5050505092915050565b6000601f82018313610a7057600080fd5b6005610a7e6109fd82610d06565b91508183856020840282011115610a9457600080fd5b60005b83811015610a555781610aaa8882610b91565b8452506020928301929190910190600101610a97565b6000601f82018313610ad157600080fd5b6005610adf6109fd82610d06565b91508183856020840282011115610af557600080fd5b60005b83811015610a555781610b0b8882610b91565b8452506020928301929190910190600101610af8565b6000601f82018313610b3257600080fd5b8135610b406109fd82610ce5565b91508181835260208401935060208101905083856020840282011115610b6557600080fd5b60005b83811015610a555781610b7b8882610b91565b8452506020928301929190910190600101610b68565b60006109d78235610d34565b60008060408385031215610bb057600080fd5b823567ffffffffffffffff811115610bc757600080fd5b610bd3858286016109de565b925050602083013567ffffffffffffffff811115610bf057600080fd5b610bfc85828601610b21565b9150509250929050565b60008060c08385031215610c1957600080fd5b6000610c258585610ac0565b92505060a0610bfc85828601610b91565b600060208284031215610c4857600080fd5b6000610c548484610b91565b949350505050565b610c6581610d24565b82525050565b610c6581610d2f565b610c6581610d34565b610c6581610d53565b602081016105208284610c5c565b602081016105208284610c6b565b602081016105208284610c74565b602081016105208284610c7d565b60405181810167ffffffffffffffff81118282101715610cdd57600080fd5b604052919050565b600067ffffffffffffffff821115610cfc57600080fd5b5060209081020190565b600067ffffffffffffffff821115610d1d57600080fd5b5060200290565b600061052082610d47565b151590565b90565b600060058210610d4357fe5b5090565b600160a060020a031690565b600061052082610d3756fea265627a7a723058202fb250ccd9384bd159cfe3f361d6dc66d364f9679e294b9fff92f53753305de06c6578706572696d656e74616cf50037`

// DeployFairswap deploys a new Ethereum contract, binding an instance of Fairswap to it.
func DeployFairswap(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int, _commitmentofkey [32]byte, _CircuitMerkleRoot [32]byte, _EncInputRoot [32]byte, circuitGatesOperationArray []*big.Int) (common.Address, *types.Transaction, *Fairswap, error) {
	parsed, err := abi.JSON(strings.NewReader(FairswapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FairswapBin), backend, _price, _commitmentofkey, _CircuitMerkleRoot, _EncInputRoot, circuitGatesOperationArray)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fairswap{FairswapCaller: FairswapCaller{contract: contract}, FairswapTransactor: FairswapTransactor{contract: contract}, FairswapFilterer: FairswapFilterer{contract: contract}}, nil
}

// Fairswap is an auto generated Go binding around an Ethereum contract.
type Fairswap struct {
	FairswapCaller     // Read-only binding to the contract
	FairswapTransactor // Write-only binding to the contract
	FairswapFilterer   // Log filterer for contract events
}

// FairswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type FairswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FairswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FairswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FairswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FairswapSession struct {
	Contract     *Fairswap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FairswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FairswapCallerSession struct {
	Contract *FairswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FairswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FairswapTransactorSession struct {
	Contract     *FairswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FairswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type FairswapRaw struct {
	Contract *Fairswap // Generic contract binding to access the raw methods on
}

// FairswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FairswapCallerRaw struct {
	Contract *FairswapCaller // Generic read-only contract binding to access the raw methods on
}

// FairswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FairswapTransactorRaw struct {
	Contract *FairswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFairswap creates a new instance of Fairswap, bound to a specific deployed contract.
func NewFairswap(address common.Address, backend bind.ContractBackend) (*Fairswap, error) {
	contract, err := bindFairswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fairswap{FairswapCaller: FairswapCaller{contract: contract}, FairswapTransactor: FairswapTransactor{contract: contract}, FairswapFilterer: FairswapFilterer{contract: contract}}, nil
}

// NewFairswapCaller creates a new read-only instance of Fairswap, bound to a specific deployed contract.
func NewFairswapCaller(address common.Address, caller bind.ContractCaller) (*FairswapCaller, error) {
	contract, err := bindFairswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FairswapCaller{contract: contract}, nil
}

// NewFairswapTransactor creates a new write-only instance of Fairswap, bound to a specific deployed contract.
func NewFairswapTransactor(address common.Address, transactor bind.ContractTransactor) (*FairswapTransactor, error) {
	contract, err := bindFairswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FairswapTransactor{contract: contract}, nil
}

// NewFairswapFilterer creates a new log filterer instance of Fairswap, bound to a specific deployed contract.
func NewFairswapFilterer(address common.Address, filterer bind.ContractFilterer) (*FairswapFilterer, error) {
	contract, err := bindFairswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FairswapFilterer{contract: contract}, nil
}

// bindFairswap binds a generic wrapper to an already deployed contract.
func bindFairswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FairswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fairswap *FairswapRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fairswap.Contract.FairswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fairswap *FairswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.Contract.FairswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fairswap *FairswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fairswap.Contract.FairswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fairswap *FairswapCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Fairswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fairswap *FairswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fairswap *FairswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fairswap.Contract.contract.Transact(opts, method, params...)
}

// CircuitMerkleRoot is a free data retrieval call binding the contract method 0xc811d63c.
//
// Solidity: function CircuitMerkleRoot() constant returns(bytes32)
func (_Fairswap *FairswapCaller) CircuitMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "CircuitMerkleRoot")
	return *ret0, err
}

// CircuitMerkleRoot is a free data retrieval call binding the contract method 0xc811d63c.
//
// Solidity: function CircuitMerkleRoot() constant returns(bytes32)
func (_Fairswap *FairswapSession) CircuitMerkleRoot() ([32]byte, error) {
	return _Fairswap.Contract.CircuitMerkleRoot(&_Fairswap.CallOpts)
}

// CircuitMerkleRoot is a free data retrieval call binding the contract method 0xc811d63c.
//
// Solidity: function CircuitMerkleRoot() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) CircuitMerkleRoot() ([32]byte, error) {
	return _Fairswap.Contract.CircuitMerkleRoot(&_Fairswap.CallOpts)
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() constant returns(bytes32)
func (_Fairswap *FairswapCaller) EncInputRoot(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "EncInputRoot")
	return *ret0, err
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() constant returns(bytes32)
func (_Fairswap *FairswapSession) EncInputRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncInputRoot(&_Fairswap.CallOpts)
}

// EncInputRoot is a free data retrieval call binding the contract method 0x062a1398.
//
// Solidity: function EncInputRoot() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) EncInputRoot() ([32]byte, error) {
	return _Fairswap.Contract.EncInputRoot(&_Fairswap.CallOpts)
}

// Mverify is a free data retrieval call binding the contract method 0xf94acc36.
//
// Solidity: function Mverify(bytes32[5] complaint, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapCaller) Mverify(opts *bind.CallOpts, complaint [5][32]byte, _index *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "Mverify", complaint, _index)
	return *ret0, err
}

// Mverify is a free data retrieval call binding the contract method 0xf94acc36.
//
// Solidity: function Mverify(bytes32[5] complaint, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapSession) Mverify(complaint [5][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.Mverify(&_Fairswap.CallOpts, complaint, _index)
}

// Mverify is a free data retrieval call binding the contract method 0xf94acc36.
//
// Solidity: function Mverify(bytes32[5] complaint, uint256 _index) constant returns(bool)
func (_Fairswap *FairswapCallerSession) Mverify(complaint [5][32]byte, _index *big.Int) (bool, error) {
	return _Fairswap.Contract.Mverify(&_Fairswap.CallOpts, complaint, _index)
}

// CircuitGatesOperation is a free data retrieval call binding the contract method 0x0b90457d.
//
// Solidity: function circuitGatesOperation(uint256 ) constant returns(uint256)
func (_Fairswap *FairswapCaller) CircuitGatesOperation(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "circuitGatesOperation", arg0)
	return *ret0, err
}

// CircuitGatesOperation is a free data retrieval call binding the contract method 0x0b90457d.
//
// Solidity: function circuitGatesOperation(uint256 ) constant returns(uint256)
func (_Fairswap *FairswapSession) CircuitGatesOperation(arg0 *big.Int) (*big.Int, error) {
	return _Fairswap.Contract.CircuitGatesOperation(&_Fairswap.CallOpts, arg0)
}

// CircuitGatesOperation is a free data retrieval call binding the contract method 0x0b90457d.
//
// Solidity: function circuitGatesOperation(uint256 ) constant returns(uint256)
func (_Fairswap *FairswapCallerSession) CircuitGatesOperation(arg0 *big.Int) (*big.Int, error) {
	return _Fairswap.Contract.CircuitGatesOperation(&_Fairswap.CallOpts, arg0)
}

// Commitmentofkey is a free data retrieval call binding the contract method 0x4c68d890.
//
// Solidity: function commitmentofkey() constant returns(bytes32)
func (_Fairswap *FairswapCaller) Commitmentofkey(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "commitmentofkey")
	return *ret0, err
}

// Commitmentofkey is a free data retrieval call binding the contract method 0x4c68d890.
//
// Solidity: function commitmentofkey() constant returns(bytes32)
func (_Fairswap *FairswapSession) Commitmentofkey() ([32]byte, error) {
	return _Fairswap.Contract.Commitmentofkey(&_Fairswap.CallOpts)
}

// Commitmentofkey is a free data retrieval call binding the contract method 0x4c68d890.
//
// Solidity: function commitmentofkey() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) Commitmentofkey() ([32]byte, error) {
	return _Fairswap.Contract.Commitmentofkey(&_Fairswap.CallOpts)
}

// ComplianBool is a free data retrieval call binding the contract method 0xa0fb1013.
//
// Solidity: function complianBool() constant returns(bool)
func (_Fairswap *FairswapCaller) ComplianBool(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "complianBool")
	return *ret0, err
}

// ComplianBool is a free data retrieval call binding the contract method 0xa0fb1013.
//
// Solidity: function complianBool() constant returns(bool)
func (_Fairswap *FairswapSession) ComplianBool() (bool, error) {
	return _Fairswap.Contract.ComplianBool(&_Fairswap.CallOpts)
}

// ComplianBool is a free data retrieval call binding the contract method 0xa0fb1013.
//
// Solidity: function complianBool() constant returns(bool)
func (_Fairswap *FairswapCallerSession) ComplianBool() (bool, error) {
	return _Fairswap.Contract.ComplianBool(&_Fairswap.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes32)
func (_Fairswap *FairswapCaller) Key(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "key")
	return *ret0, err
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes32)
func (_Fairswap *FairswapSession) Key() ([32]byte, error) {
	return _Fairswap.Contract.Key(&_Fairswap.CallOpts)
}

// Key is a free data retrieval call binding the contract method 0x3943380c.
//
// Solidity: function key() constant returns(bytes32)
func (_Fairswap *FairswapCallerSession) Key() ([32]byte, error) {
	return _Fairswap.Contract.Key(&_Fairswap.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Fairswap *FairswapCaller) Phase(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "phase")
	return *ret0, err
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Fairswap *FairswapSession) Phase() (uint8, error) {
	return _Fairswap.Contract.Phase(&_Fairswap.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() constant returns(uint8)
func (_Fairswap *FairswapCallerSession) Phase() (uint8, error) {
	return _Fairswap.Contract.Phase(&_Fairswap.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Fairswap *FairswapCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Fairswap *FairswapSession) Price() (*big.Int, error) {
	return _Fairswap.Contract.Price(&_Fairswap.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) Price() (*big.Int, error) {
	return _Fairswap.Contract.Price(&_Fairswap.CallOpts)
}

// Recieveraddrs is a free data retrieval call binding the contract method 0xa784df2a.
//
// Solidity: function recieveraddrs() constant returns(address)
func (_Fairswap *FairswapCaller) Recieveraddrs(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "recieveraddrs")
	return *ret0, err
}

// Recieveraddrs is a free data retrieval call binding the contract method 0xa784df2a.
//
// Solidity: function recieveraddrs() constant returns(address)
func (_Fairswap *FairswapSession) Recieveraddrs() (common.Address, error) {
	return _Fairswap.Contract.Recieveraddrs(&_Fairswap.CallOpts)
}

// Recieveraddrs is a free data retrieval call binding the contract method 0xa784df2a.
//
// Solidity: function recieveraddrs() constant returns(address)
func (_Fairswap *FairswapCallerSession) Recieveraddrs() (common.Address, error) {
	return _Fairswap.Contract.Recieveraddrs(&_Fairswap.CallOpts)
}

// Senderaddrs is a free data retrieval call binding the contract method 0x945ef141.
//
// Solidity: function senderaddrs() constant returns(address)
func (_Fairswap *FairswapCaller) Senderaddrs(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "senderaddrs")
	return *ret0, err
}

// Senderaddrs is a free data retrieval call binding the contract method 0x945ef141.
//
// Solidity: function senderaddrs() constant returns(address)
func (_Fairswap *FairswapSession) Senderaddrs() (common.Address, error) {
	return _Fairswap.Contract.Senderaddrs(&_Fairswap.CallOpts)
}

// Senderaddrs is a free data retrieval call binding the contract method 0x945ef141.
//
// Solidity: function senderaddrs() constant returns(address)
func (_Fairswap *FairswapCallerSession) Senderaddrs() (common.Address, error) {
	return _Fairswap.Contract.Senderaddrs(&_Fairswap.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Fairswap *FairswapCaller) Timeout(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Fairswap.contract.Call(opts, out, "timeout")
	return *ret0, err
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Fairswap *FairswapSession) Timeout() (*big.Int, error) {
	return _Fairswap.Contract.Timeout(&_Fairswap.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() constant returns(uint256)
func (_Fairswap *FairswapCallerSession) Timeout() (*big.Int, error) {
	return _Fairswap.Contract.Timeout(&_Fairswap.CallOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Fairswap *FairswapTransactor) Accept(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "accept")
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Fairswap *FairswapSession) Accept() (*types.Transaction, error) {
	return _Fairswap.Contract.Accept(&_Fairswap.TransactOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x2852b71c.
//
// Solidity: function accept() returns()
func (_Fairswap *FairswapTransactorSession) Accept() (*types.Transaction, error) {
	return _Fairswap.Contract.Accept(&_Fairswap.TransactOpts)
}

// Complain is a paid mutator transaction binding the contract method 0x7896c64e.
//
// Solidity: function complain(bytes32[5][] complaint, uint256[] indices) returns(bool)
func (_Fairswap *FairswapTransactor) Complain(opts *bind.TransactOpts, complaint [][5][32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "complain", complaint, indices)
}

// Complain is a paid mutator transaction binding the contract method 0x7896c64e.
//
// Solidity: function complain(bytes32[5][] complaint, uint256[] indices) returns(bool)
func (_Fairswap *FairswapSession) Complain(complaint [][5][32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _Fairswap.Contract.Complain(&_Fairswap.TransactOpts, complaint, indices)
}

// Complain is a paid mutator transaction binding the contract method 0x7896c64e.
//
// Solidity: function complain(bytes32[5][] complaint, uint256[] indices) returns(bool)
func (_Fairswap *FairswapTransactorSession) Complain(complaint [][5][32]byte, indices []*big.Int) (*types.Transaction, error) {
	return _Fairswap.Contract.Complain(&_Fairswap.TransactOpts, complaint, indices)
}

// Initializerecieveraddrs is a paid mutator transaction binding the contract method 0xe0e5f68f.
//
// Solidity: function initializerecieveraddrs() returns()
func (_Fairswap *FairswapTransactor) Initializerecieveraddrs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "initializerecieveraddrs")
}

// Initializerecieveraddrs is a paid mutator transaction binding the contract method 0xe0e5f68f.
//
// Solidity: function initializerecieveraddrs() returns()
func (_Fairswap *FairswapSession) Initializerecieveraddrs() (*types.Transaction, error) {
	return _Fairswap.Contract.Initializerecieveraddrs(&_Fairswap.TransactOpts)
}

// Initializerecieveraddrs is a paid mutator transaction binding the contract method 0xe0e5f68f.
//
// Solidity: function initializerecieveraddrs() returns()
func (_Fairswap *FairswapTransactorSession) Initializerecieveraddrs() (*types.Transaction, error) {
	return _Fairswap.Contract.Initializerecieveraddrs(&_Fairswap.TransactOpts)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapTransactor) NextStage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "nextStage")
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapSession) NextStage() (*types.Transaction, error) {
	return _Fairswap.Contract.NextStage(&_Fairswap.TransactOpts)
}

// NextStage is a paid mutator transaction binding the contract method 0xee3743ab.
//
// Solidity: function nextStage() returns()
func (_Fairswap *FairswapTransactorSession) NextStage() (*types.Transaction, error) {
	return _Fairswap.Contract.NextStage(&_Fairswap.TransactOpts)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapTransactor) RevealKey(opts *bind.TransactOpts, _key [32]byte) (*types.Transaction, error) {
	return _Fairswap.contract.Transact(opts, "revealKey", _key)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapSession) RevealKey(_key [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.RevealKey(&_Fairswap.TransactOpts, _key)
}

// RevealKey is a paid mutator transaction binding the contract method 0xd6547236.
//
// Solidity: function revealKey(bytes32 _key) returns()
func (_Fairswap *FairswapTransactorSession) RevealKey(_key [32]byte) (*types.Transaction, error) {
	return _Fairswap.Contract.RevealKey(&_Fairswap.TransactOpts, _key)
}
