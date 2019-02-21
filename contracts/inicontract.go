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
const FairswapABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"EncInputRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accept\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitmentofkey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"senderaddrs\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"recieveraddrs\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CircuitMerkleRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initializerecieveraddrs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"nextStage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_commitmentofkey\",\"type\":\"bytes32\"},{\"name\":\"_CircuitMerkleRoot\",\"type\":\"bytes32\"},{\"name\":\"_EncInputRoot\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// FairswapBin is the compiled bytecode used for deploying new contracts.
const FairswapBin = `0x60806040526007805460ff1916905534801561001a57600080fd5b5060405160808061044d8339810180604052608081101561003a57600080fd5b5080516020820151604083015160609093015160048054600160a060020a03191633179055600283905560038290556000849055600181905591929091610088640100000000610091810204565b505050506100d4565b60075460ff1660048111156100a257fe5b60010160048111156100b057fe5b6007805460ff191660018360048111156100c657fe5b021790555060b44201600655565b61036a806100e36000396000f3fe6080604052600436106100b9576000357c010000000000000000000000000000000000000000000000000000000090048063a035b1fe11610081578063a035b1fe14610157578063a784df2a1461016c578063b1c9fe6e14610181578063c811d63c146101ba578063e0e5f68f146101cf578063ee3743ab146101e4576100b9565b8063062a1398146100be5780632852b71c146100e55780634c68d890146100ef57806370dea79a14610104578063945ef14114610119575b600080fd5b3480156100ca57600080fd5b506100d36101f9565b60408051918252519081900360200190f35b6100ed6101ff565b005b3480156100fb57600080fd5b506100d3610281565b34801561011057600080fd5b506100d3610287565b34801561012557600080fd5b5061012e61028d565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561016357600080fd5b506100d36102a9565b34801561017857600080fd5b5061012e6102af565b34801561018d57600080fd5b506101966102cb565b604051808260048111156101a657fe5b60ff16815260200191505060405180910390f35b3480156101c657600080fd5b506100d36102d4565b3480156101db57600080fd5b506100ed6102da565b3480156101f057600080fd5b506100ed6102fb565b60015481565b60055473ffffffffffffffffffffffffffffffffffffffff1660018060075460ff16600481111561022c57fe5b1461023657600080fd5b600654421061024457600080fd5b3373ffffffffffffffffffffffffffffffffffffffff83161461026657600080fd5b60025434101561027557600080fd5b61027d6102fb565b5050565b60035481565b60065481565b60045473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b60055473ffffffffffffffffffffffffffffffffffffffff1681565b60075460ff1681565b60005481565b6005805473ffffffffffffffffffffffffffffffffffffffff191633179055565b60075460ff16600481111561030c57fe5b600101600481111561031a57fe5b6007805460ff1916600183600481111561033057fe5b021790555060b4420160065556fea165627a7a7230582079648d7d24a433d2bcd839f4f725cf3d92546cd9adea33704e7d7ab8671f290c0029`

// DeployFairswap deploys a new Ethereum contract, binding an instance of Fairswap to it.
func DeployFairswap(auth *bind.TransactOpts, backend bind.ContractBackend, _price *big.Int, _commitmentofkey [32]byte, _CircuitMerkleRoot [32]byte, _EncInputRoot [32]byte) (common.Address, *types.Transaction, *Fairswap, error) {
	parsed, err := abi.JSON(strings.NewReader(FairswapABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FairswapBin), backend, _price, _commitmentofkey, _CircuitMerkleRoot, _EncInputRoot)
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
