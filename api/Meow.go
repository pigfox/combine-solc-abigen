// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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
	_ = abi.ConvertType
)

// MeowMetaData contains all meta data concerning the Meow contract.
var MeowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"myFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}][{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"myFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}][{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"myFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}][{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"myFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b50606a80601a5f395ff3fe6080604052348015600e575f80fd5b50600436106026575f3560e01c8063c3780a3a14602a575b5f80fd5b603260035f55565b00fea26469706673582212209b498247d95b06f9aa10532b25e9928ac955a571fbad018b3d81df5c16a5913964736f6c634300081900336080604052348015600e575f80fd5b50606a80601a5f395ff3fe6080604052348015600e575f80fd5b50600436106026575f3560e01c8063c3780a3a14602a575b5f80fd5b603260035f55565b00fea26469706673582212204d91517e152fd222d9d185cba0f638c374c2c77d3a5a9c0857a92f455a8d0eb564736f6c634300081900336080604052348015600e575f80fd5b50606a80601a5f395ff3fe6080604052348015600e575f80fd5b50600436106026575f3560e01c8063c3780a3a14602a575b5f80fd5b603260035f55565b00fea26469706673582212209b498247d95b06f9aa10532b25e9928ac955a571fbad018b3d81df5c16a5913964736f6c634300081900336080604052348015600e575f80fd5b50606a80601a5f395ff3fe6080604052348015600e575f80fd5b50600436106026575f3560e01c8063c3780a3a14602a575b5f80fd5b603260035f55565b00fea26469706673582212204d91517e152fd222d9d185cba0f638c374c2c77d3a5a9c0857a92f455a8d0eb564736f6c63430008190033",
}

// MeowABI is the input ABI used to generate the binding from.
// Deprecated: Use MeowMetaData.ABI instead.
var MeowABI = MeowMetaData.ABI

// MeowBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MeowMetaData.Bin instead.
var MeowBin = MeowMetaData.Bin

// DeployMeow deploys a new Ethereum contract, binding an instance of Meow to it.
func DeployMeow(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Meow, error) {
	parsed, err := MeowMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MeowBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Meow{MeowCaller: MeowCaller{contract: contract}, MeowTransactor: MeowTransactor{contract: contract}, MeowFilterer: MeowFilterer{contract: contract}}, nil
}

// Meow is an auto generated Go binding around an Ethereum contract.
type Meow struct {
	MeowCaller     // Read-only binding to the contract
	MeowTransactor // Write-only binding to the contract
	MeowFilterer   // Log filterer for contract events
}

// MeowCaller is an auto generated read-only Go binding around an Ethereum contract.
type MeowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MeowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MeowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MeowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MeowSession struct {
	Contract     *Meow             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MeowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MeowCallerSession struct {
	Contract *MeowCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MeowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MeowTransactorSession struct {
	Contract     *MeowTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MeowRaw is an auto generated low-level Go binding around an Ethereum contract.
type MeowRaw struct {
	Contract *Meow // Generic contract binding to access the raw methods on
}

// MeowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MeowCallerRaw struct {
	Contract *MeowCaller // Generic read-only contract binding to access the raw methods on
}

// MeowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MeowTransactorRaw struct {
	Contract *MeowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMeow creates a new instance of Meow, bound to a specific deployed contract.
func NewMeow(address common.Address, backend bind.ContractBackend) (*Meow, error) {
	contract, err := bindMeow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Meow{MeowCaller: MeowCaller{contract: contract}, MeowTransactor: MeowTransactor{contract: contract}, MeowFilterer: MeowFilterer{contract: contract}}, nil
}

// NewMeowCaller creates a new read-only instance of Meow, bound to a specific deployed contract.
func NewMeowCaller(address common.Address, caller bind.ContractCaller) (*MeowCaller, error) {
	contract, err := bindMeow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MeowCaller{contract: contract}, nil
}

// NewMeowTransactor creates a new write-only instance of Meow, bound to a specific deployed contract.
func NewMeowTransactor(address common.Address, transactor bind.ContractTransactor) (*MeowTransactor, error) {
	contract, err := bindMeow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MeowTransactor{contract: contract}, nil
}

// NewMeowFilterer creates a new log filterer instance of Meow, bound to a specific deployed contract.
func NewMeowFilterer(address common.Address, filterer bind.ContractFilterer) (*MeowFilterer, error) {
	contract, err := bindMeow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MeowFilterer{contract: contract}, nil
}

// bindMeow binds a generic wrapper to an already deployed contract.
func bindMeow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MeowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meow *MeowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meow.Contract.MeowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meow *MeowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meow.Contract.MeowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meow *MeowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meow.Contract.MeowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Meow *MeowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Meow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Meow *MeowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Meow *MeowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Meow.Contract.contract.Transact(opts, method, params...)
}

// MyFunction is a paid mutator transaction binding the contract method 0xc3780a3a.
//
// Solidity: function myFunction() returns()
func (_Meow *MeowTransactor) MyFunction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Meow.contract.Transact(opts, "myFunction")
}

// MyFunction is a paid mutator transaction binding the contract method 0xc3780a3a.
//
// Solidity: function myFunction() returns()
func (_Meow *MeowSession) MyFunction() (*types.Transaction, error) {
	return _Meow.Contract.MyFunction(&_Meow.TransactOpts)
}

// MyFunction is a paid mutator transaction binding the contract method 0xc3780a3a.
//
// Solidity: function myFunction() returns()
func (_Meow *MeowTransactorSession) MyFunction() (*types.Transaction, error) {
	return _Meow.Contract.MyFunction(&_Meow.TransactOpts)
}
