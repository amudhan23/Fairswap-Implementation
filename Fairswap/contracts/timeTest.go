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

// TimeTestABI is the input ABI used to generate the binding from.
const TimeTestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"Now\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"Nowy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TimeTestBin is the compiled bytecode used for deploying new contracts.
const TimeTestBin = `0x608060405234801561001057600080fd5b5060dd8061001f6000396000f3fe6080604052348015600f57600080fd5b5060043610604e577c0100000000000000000000000000000000000000000000000000000000600035046344d4fd1981146053578063fd6fd6fa14606d575b600080fd5b60596075565b60405160649190608e565b60405180910390f35b6073607b565b005b60005481565b42600055565b60888160a0565b82525050565b60208101609a82846081565b92915050565b9056fea265627a7a72305820c496371a885be6d66b9d0fc33688441bb1b2925ce6783fd0c3b5a95769f0604e6c6578706572696d656e74616cf50037`

// DeployTimeTest deploys a new Ethereum contract, binding an instance of TimeTest to it.
func DeployTimeTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TimeTest, error) {
	parsed, err := abi.JSON(strings.NewReader(TimeTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TimeTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TimeTest{TimeTestCaller: TimeTestCaller{contract: contract}, TimeTestTransactor: TimeTestTransactor{contract: contract}, TimeTestFilterer: TimeTestFilterer{contract: contract}}, nil
}

// TimeTest is an auto generated Go binding around an Ethereum contract.
type TimeTest struct {
	TimeTestCaller     // Read-only binding to the contract
	TimeTestTransactor // Write-only binding to the contract
	TimeTestFilterer   // Log filterer for contract events
}

// TimeTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimeTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimeTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimeTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimeTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimeTestSession struct {
	Contract     *TimeTest         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimeTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimeTestCallerSession struct {
	Contract *TimeTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TimeTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimeTestTransactorSession struct {
	Contract     *TimeTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TimeTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimeTestRaw struct {
	Contract *TimeTest // Generic contract binding to access the raw methods on
}

// TimeTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimeTestCallerRaw struct {
	Contract *TimeTestCaller // Generic read-only contract binding to access the raw methods on
}

// TimeTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimeTestTransactorRaw struct {
	Contract *TimeTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimeTest creates a new instance of TimeTest, bound to a specific deployed contract.
func NewTimeTest(address common.Address, backend bind.ContractBackend) (*TimeTest, error) {
	contract, err := bindTimeTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimeTest{TimeTestCaller: TimeTestCaller{contract: contract}, TimeTestTransactor: TimeTestTransactor{contract: contract}, TimeTestFilterer: TimeTestFilterer{contract: contract}}, nil
}

// NewTimeTestCaller creates a new read-only instance of TimeTest, bound to a specific deployed contract.
func NewTimeTestCaller(address common.Address, caller bind.ContractCaller) (*TimeTestCaller, error) {
	contract, err := bindTimeTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimeTestCaller{contract: contract}, nil
}

// NewTimeTestTransactor creates a new write-only instance of TimeTest, bound to a specific deployed contract.
func NewTimeTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TimeTestTransactor, error) {
	contract, err := bindTimeTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimeTestTransactor{contract: contract}, nil
}

// NewTimeTestFilterer creates a new log filterer instance of TimeTest, bound to a specific deployed contract.
func NewTimeTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TimeTestFilterer, error) {
	contract, err := bindTimeTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimeTestFilterer{contract: contract}, nil
}

// bindTimeTest binds a generic wrapper to an already deployed contract.
func bindTimeTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimeTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeTest *TimeTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TimeTest.Contract.TimeTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeTest *TimeTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeTest.Contract.TimeTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeTest *TimeTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeTest.Contract.TimeTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimeTest *TimeTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TimeTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimeTest *TimeTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimeTest *TimeTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimeTest.Contract.contract.Transact(opts, method, params...)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_TimeTest *TimeTestCaller) Now(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TimeTest.contract.Call(opts, out, "Now")
	return *ret0, err
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_TimeTest *TimeTestSession) Now() (*big.Int, error) {
	return _TimeTest.Contract.Now(&_TimeTest.CallOpts)
}

// Now is a free data retrieval call binding the contract method 0x44d4fd19.
//
// Solidity: function Now() constant returns(uint256)
func (_TimeTest *TimeTestCallerSession) Now() (*big.Int, error) {
	return _TimeTest.Contract.Now(&_TimeTest.CallOpts)
}

// Nowy is a paid mutator transaction binding the contract method 0xfd6fd6fa.
//
// Solidity: function Nowy() returns()
func (_TimeTest *TimeTestTransactor) Nowy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimeTest.contract.Transact(opts, "Nowy")
}

// Nowy is a paid mutator transaction binding the contract method 0xfd6fd6fa.
//
// Solidity: function Nowy() returns()
func (_TimeTest *TimeTestSession) Nowy() (*types.Transaction, error) {
	return _TimeTest.Contract.Nowy(&_TimeTest.TransactOpts)
}

// Nowy is a paid mutator transaction binding the contract method 0xfd6fd6fa.
//
// Solidity: function Nowy() returns()
func (_TimeTest *TimeTestTransactorSession) Nowy() (*types.Transaction, error) {
	return _TimeTest.Contract.Nowy(&_TimeTest.TransactOpts)
}
