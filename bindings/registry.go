// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ProviderRegistryProvider is an auto generated low-level Go binding around an user-defined struct.
type ProviderRegistryProvider struct {
	ProviderAddress common.Address
	Cpu             *big.Int
	MemoryMB        *big.Int
	PricePerHour    *big.Int
	Reputation      *big.Int
	Active          bool
}

// ProviderRegistryMetaData contains all meta data concerning the ProviderRegistry contract.
var ProviderRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"ProviderDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"memoryMB\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerHour\",\"type\":\"uint256\"}],\"name\":\"ProviderRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newScore\",\"type\":\"uint256\"}],\"name\":\"ReputationUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deregisterProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllProviders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"}],\"name\":\"getProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"providerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryMB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerHour\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structProviderRegistry.Provider\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"providerList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"providers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"providerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryMB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerHour\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cpu\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryMB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerHour\",\"type\":\"uint256\"}],\"name\":\"registerProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"}],\"name\":\"setMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"updatePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"providerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newScore\",\"type\":\"uint256\"}],\"name\":\"updateReputation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProviderRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProviderRegistryMetaData.ABI instead.
var ProviderRegistryABI = ProviderRegistryMetaData.ABI

// ProviderRegistry is an auto generated Go binding around an Ethereum contract.
type ProviderRegistry struct {
	ProviderRegistryCaller     // Read-only binding to the contract
	ProviderRegistryTransactor // Write-only binding to the contract
	ProviderRegistryFilterer   // Log filterer for contract events
}

// ProviderRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProviderRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProviderRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProviderRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProviderRegistrySession struct {
	Contract     *ProviderRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProviderRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProviderRegistryCallerSession struct {
	Contract *ProviderRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProviderRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProviderRegistryTransactorSession struct {
	Contract     *ProviderRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProviderRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProviderRegistryRaw struct {
	Contract *ProviderRegistry // Generic contract binding to access the raw methods on
}

// ProviderRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProviderRegistryCallerRaw struct {
	Contract *ProviderRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProviderRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProviderRegistryTransactorRaw struct {
	Contract *ProviderRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProviderRegistry creates a new instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistry(address common.Address, backend bind.ContractBackend) (*ProviderRegistry, error) {
	contract, err := bindProviderRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistry{ProviderRegistryCaller: ProviderRegistryCaller{contract: contract}, ProviderRegistryTransactor: ProviderRegistryTransactor{contract: contract}, ProviderRegistryFilterer: ProviderRegistryFilterer{contract: contract}}, nil
}

// NewProviderRegistryCaller creates a new read-only instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProviderRegistryCaller, error) {
	contract, err := bindProviderRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryCaller{contract: contract}, nil
}

// NewProviderRegistryTransactor creates a new write-only instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProviderRegistryTransactor, error) {
	contract, err := bindProviderRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryTransactor{contract: contract}, nil
}

// NewProviderRegistryFilterer creates a new log filterer instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProviderRegistryFilterer, error) {
	contract, err := bindProviderRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryFilterer{contract: contract}, nil
}

// bindProviderRegistry binds a generic wrapper to an already deployed contract.
func bindProviderRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProviderRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderRegistry *ProviderRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderRegistry.Contract.ProviderRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderRegistry *ProviderRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderRegistry *ProviderRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderRegistry *ProviderRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderRegistry *ProviderRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderRegistry *ProviderRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetAllProviders is a free data retrieval call binding the contract method 0x3bb4497c.
//
// Solidity: function getAllProviders() view returns(address[])
func (_ProviderRegistry *ProviderRegistryCaller) GetAllProviders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getAllProviders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllProviders is a free data retrieval call binding the contract method 0x3bb4497c.
//
// Solidity: function getAllProviders() view returns(address[])
func (_ProviderRegistry *ProviderRegistrySession) GetAllProviders() ([]common.Address, error) {
	return _ProviderRegistry.Contract.GetAllProviders(&_ProviderRegistry.CallOpts)
}

// GetAllProviders is a free data retrieval call binding the contract method 0x3bb4497c.
//
// Solidity: function getAllProviders() view returns(address[])
func (_ProviderRegistry *ProviderRegistryCallerSession) GetAllProviders() ([]common.Address, error) {
	return _ProviderRegistry.Contract.GetAllProviders(&_ProviderRegistry.CallOpts)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address providerAddr) view returns((address,uint256,uint256,uint256,uint256,bool))
func (_ProviderRegistry *ProviderRegistryCaller) GetProvider(opts *bind.CallOpts, providerAddr common.Address) (ProviderRegistryProvider, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getProvider", providerAddr)

	if err != nil {
		return *new(ProviderRegistryProvider), err
	}

	out0 := *abi.ConvertType(out[0], new(ProviderRegistryProvider)).(*ProviderRegistryProvider)

	return out0, err

}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address providerAddr) view returns((address,uint256,uint256,uint256,uint256,bool))
func (_ProviderRegistry *ProviderRegistrySession) GetProvider(providerAddr common.Address) (ProviderRegistryProvider, error) {
	return _ProviderRegistry.Contract.GetProvider(&_ProviderRegistry.CallOpts, providerAddr)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address providerAddr) view returns((address,uint256,uint256,uint256,uint256,bool))
func (_ProviderRegistry *ProviderRegistryCallerSession) GetProvider(providerAddr common.Address) (ProviderRegistryProvider, error) {
	return _ProviderRegistry.Contract.GetProvider(&_ProviderRegistry.CallOpts, providerAddr)
}

// GetProviderCount is a free data retrieval call binding the contract method 0x46ce4175.
//
// Solidity: function getProviderCount() view returns(uint256)
func (_ProviderRegistry *ProviderRegistryCaller) GetProviderCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getProviderCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProviderCount is a free data retrieval call binding the contract method 0x46ce4175.
//
// Solidity: function getProviderCount() view returns(uint256)
func (_ProviderRegistry *ProviderRegistrySession) GetProviderCount() (*big.Int, error) {
	return _ProviderRegistry.Contract.GetProviderCount(&_ProviderRegistry.CallOpts)
}

// GetProviderCount is a free data retrieval call binding the contract method 0x46ce4175.
//
// Solidity: function getProviderCount() view returns(uint256)
func (_ProviderRegistry *ProviderRegistryCallerSession) GetProviderCount() (*big.Int, error) {
	return _ProviderRegistry.Contract.GetProviderCount(&_ProviderRegistry.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_ProviderRegistry *ProviderRegistryCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_ProviderRegistry *ProviderRegistrySession) Marketplace() (common.Address, error) {
	return _ProviderRegistry.Contract.Marketplace(&_ProviderRegistry.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_ProviderRegistry *ProviderRegistryCallerSession) Marketplace() (common.Address, error) {
	return _ProviderRegistry.Contract.Marketplace(&_ProviderRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProviderRegistry *ProviderRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProviderRegistry *ProviderRegistrySession) Owner() (common.Address, error) {
	return _ProviderRegistry.Contract.Owner(&_ProviderRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProviderRegistry *ProviderRegistryCallerSession) Owner() (common.Address, error) {
	return _ProviderRegistry.Contract.Owner(&_ProviderRegistry.CallOpts)
}

// ProviderList is a free data retrieval call binding the contract method 0x390bfd3c.
//
// Solidity: function providerList(uint256 ) view returns(address)
func (_ProviderRegistry *ProviderRegistryCaller) ProviderList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "providerList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProviderList is a free data retrieval call binding the contract method 0x390bfd3c.
//
// Solidity: function providerList(uint256 ) view returns(address)
func (_ProviderRegistry *ProviderRegistrySession) ProviderList(arg0 *big.Int) (common.Address, error) {
	return _ProviderRegistry.Contract.ProviderList(&_ProviderRegistry.CallOpts, arg0)
}

// ProviderList is a free data retrieval call binding the contract method 0x390bfd3c.
//
// Solidity: function providerList(uint256 ) view returns(address)
func (_ProviderRegistry *ProviderRegistryCallerSession) ProviderList(arg0 *big.Int) (common.Address, error) {
	return _ProviderRegistry.Contract.ProviderList(&_ProviderRegistry.CallOpts, arg0)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(address providerAddress, uint256 cpu, uint256 memoryMB, uint256 pricePerHour, uint256 reputation, bool active)
func (_ProviderRegistry *ProviderRegistryCaller) Providers(opts *bind.CallOpts, arg0 common.Address) (struct {
	ProviderAddress common.Address
	Cpu             *big.Int
	MemoryMB        *big.Int
	PricePerHour    *big.Int
	Reputation      *big.Int
	Active          bool
}, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "providers", arg0)

	outstruct := new(struct {
		ProviderAddress common.Address
		Cpu             *big.Int
		MemoryMB        *big.Int
		PricePerHour    *big.Int
		Reputation      *big.Int
		Active          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProviderAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Cpu = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MemoryMB = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PricePerHour = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Reputation = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(address providerAddress, uint256 cpu, uint256 memoryMB, uint256 pricePerHour, uint256 reputation, bool active)
func (_ProviderRegistry *ProviderRegistrySession) Providers(arg0 common.Address) (struct {
	ProviderAddress common.Address
	Cpu             *big.Int
	MemoryMB        *big.Int
	PricePerHour    *big.Int
	Reputation      *big.Int
	Active          bool
}, error) {
	return _ProviderRegistry.Contract.Providers(&_ProviderRegistry.CallOpts, arg0)
}

// Providers is a free data retrieval call binding the contract method 0x0787bc27.
//
// Solidity: function providers(address ) view returns(address providerAddress, uint256 cpu, uint256 memoryMB, uint256 pricePerHour, uint256 reputation, bool active)
func (_ProviderRegistry *ProviderRegistryCallerSession) Providers(arg0 common.Address) (struct {
	ProviderAddress common.Address
	Cpu             *big.Int
	MemoryMB        *big.Int
	PricePerHour    *big.Int
	Reputation      *big.Int
	Active          bool
}, error) {
	return _ProviderRegistry.Contract.Providers(&_ProviderRegistry.CallOpts, arg0)
}

// DeregisterProvider is a paid mutator transaction binding the contract method 0x0ce8a33b.
//
// Solidity: function deregisterProvider() returns()
func (_ProviderRegistry *ProviderRegistryTransactor) DeregisterProvider(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "deregisterProvider")
}

// DeregisterProvider is a paid mutator transaction binding the contract method 0x0ce8a33b.
//
// Solidity: function deregisterProvider() returns()
func (_ProviderRegistry *ProviderRegistrySession) DeregisterProvider() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.DeregisterProvider(&_ProviderRegistry.TransactOpts)
}

// DeregisterProvider is a paid mutator transaction binding the contract method 0x0ce8a33b.
//
// Solidity: function deregisterProvider() returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) DeregisterProvider() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.DeregisterProvider(&_ProviderRegistry.TransactOpts)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0x1bace2ff.
//
// Solidity: function registerProvider(uint256 cpu, uint256 memoryMB, uint256 pricePerHour) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) RegisterProvider(opts *bind.TransactOpts, cpu *big.Int, memoryMB *big.Int, pricePerHour *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "registerProvider", cpu, memoryMB, pricePerHour)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0x1bace2ff.
//
// Solidity: function registerProvider(uint256 cpu, uint256 memoryMB, uint256 pricePerHour) returns()
func (_ProviderRegistry *ProviderRegistrySession) RegisterProvider(cpu *big.Int, memoryMB *big.Int, pricePerHour *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.RegisterProvider(&_ProviderRegistry.TransactOpts, cpu, memoryMB, pricePerHour)
}

// RegisterProvider is a paid mutator transaction binding the contract method 0x1bace2ff.
//
// Solidity: function registerProvider(uint256 cpu, uint256 memoryMB, uint256 pricePerHour) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) RegisterProvider(cpu *big.Int, memoryMB *big.Int, pricePerHour *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.RegisterProvider(&_ProviderRegistry.TransactOpts, cpu, memoryMB, pricePerHour)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address _marketplace) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) SetMarketplace(opts *bind.TransactOpts, _marketplace common.Address) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "setMarketplace", _marketplace)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address _marketplace) returns()
func (_ProviderRegistry *ProviderRegistrySession) SetMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.SetMarketplace(&_ProviderRegistry.TransactOpts, _marketplace)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address _marketplace) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) SetMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.SetMarketplace(&_ProviderRegistry.TransactOpts, _marketplace)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 newPrice) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) UpdatePrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "updatePrice", newPrice)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 newPrice) returns()
func (_ProviderRegistry *ProviderRegistrySession) UpdatePrice(newPrice *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.UpdatePrice(&_ProviderRegistry.TransactOpts, newPrice)
}

// UpdatePrice is a paid mutator transaction binding the contract method 0x8d6cc56d.
//
// Solidity: function updatePrice(uint256 newPrice) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) UpdatePrice(newPrice *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.UpdatePrice(&_ProviderRegistry.TransactOpts, newPrice)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0xf5c91a08.
//
// Solidity: function updateReputation(address providerAddr, uint256 newScore) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) UpdateReputation(opts *bind.TransactOpts, providerAddr common.Address, newScore *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "updateReputation", providerAddr, newScore)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0xf5c91a08.
//
// Solidity: function updateReputation(address providerAddr, uint256 newScore) returns()
func (_ProviderRegistry *ProviderRegistrySession) UpdateReputation(providerAddr common.Address, newScore *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.UpdateReputation(&_ProviderRegistry.TransactOpts, providerAddr, newScore)
}

// UpdateReputation is a paid mutator transaction binding the contract method 0xf5c91a08.
//
// Solidity: function updateReputation(address providerAddr, uint256 newScore) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) UpdateReputation(providerAddr common.Address, newScore *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.UpdateReputation(&_ProviderRegistry.TransactOpts, providerAddr, newScore)
}

// ProviderRegistryProviderDeregisteredIterator is returned from FilterProviderDeregistered and is used to iterate over the raw logs and unpacked data for ProviderDeregistered events raised by the ProviderRegistry contract.
type ProviderRegistryProviderDeregisteredIterator struct {
	Event *ProviderRegistryProviderDeregistered // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryProviderDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryProviderDeregistered)
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
		it.Event = new(ProviderRegistryProviderDeregistered)
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
func (it *ProviderRegistryProviderDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryProviderDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryProviderDeregistered represents a ProviderDeregistered event raised by the ProviderRegistry contract.
type ProviderRegistryProviderDeregistered struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderDeregistered is a free log retrieval operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterProviderDeregistered(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryProviderDeregisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ProviderDeregistered", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryProviderDeregisteredIterator{contract: _ProviderRegistry.contract, event: "ProviderDeregistered", logs: logs, sub: sub}, nil
}

// WatchProviderDeregistered is a free log subscription operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchProviderDeregistered(opts *bind.WatchOpts, sink chan<- *ProviderRegistryProviderDeregistered, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ProviderDeregistered", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryProviderDeregistered)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderDeregistered", log); err != nil {
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

// ParseProviderDeregistered is a log parse operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) ParseProviderDeregistered(log types.Log) (*ProviderRegistryProviderDeregistered, error) {
	event := new(ProviderRegistryProviderDeregistered)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryProviderRegisteredIterator is returned from FilterProviderRegistered and is used to iterate over the raw logs and unpacked data for ProviderRegistered events raised by the ProviderRegistry contract.
type ProviderRegistryProviderRegisteredIterator struct {
	Event *ProviderRegistryProviderRegistered // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryProviderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryProviderRegistered)
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
		it.Event = new(ProviderRegistryProviderRegistered)
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
func (it *ProviderRegistryProviderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryProviderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryProviderRegistered represents a ProviderRegistered event raised by the ProviderRegistry contract.
type ProviderRegistryProviderRegistered struct {
	Provider     common.Address
	Cpu          *big.Int
	MemoryMB     *big.Int
	PricePerHour *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProviderRegistered is a free log retrieval operation binding the contract event 0x84324aff969f0e0ccdeead118a9c77219798dcc7d430a64725f3c3eb8996ee73.
//
// Solidity: event ProviderRegistered(address indexed provider, uint256 cpu, uint256 memoryMB, uint256 pricePerHour)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterProviderRegistered(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryProviderRegisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ProviderRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryProviderRegisteredIterator{contract: _ProviderRegistry.contract, event: "ProviderRegistered", logs: logs, sub: sub}, nil
}

// WatchProviderRegistered is a free log subscription operation binding the contract event 0x84324aff969f0e0ccdeead118a9c77219798dcc7d430a64725f3c3eb8996ee73.
//
// Solidity: event ProviderRegistered(address indexed provider, uint256 cpu, uint256 memoryMB, uint256 pricePerHour)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchProviderRegistered(opts *bind.WatchOpts, sink chan<- *ProviderRegistryProviderRegistered, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ProviderRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryProviderRegistered)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderRegistered", log); err != nil {
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

// ParseProviderRegistered is a log parse operation binding the contract event 0x84324aff969f0e0ccdeead118a9c77219798dcc7d430a64725f3c3eb8996ee73.
//
// Solidity: event ProviderRegistered(address indexed provider, uint256 cpu, uint256 memoryMB, uint256 pricePerHour)
func (_ProviderRegistry *ProviderRegistryFilterer) ParseProviderRegistered(log types.Log) (*ProviderRegistryProviderRegistered, error) {
	event := new(ProviderRegistryProviderRegistered)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryReputationUpdatedIterator is returned from FilterReputationUpdated and is used to iterate over the raw logs and unpacked data for ReputationUpdated events raised by the ProviderRegistry contract.
type ProviderRegistryReputationUpdatedIterator struct {
	Event *ProviderRegistryReputationUpdated // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryReputationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryReputationUpdated)
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
		it.Event = new(ProviderRegistryReputationUpdated)
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
func (it *ProviderRegistryReputationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryReputationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryReputationUpdated represents a ReputationUpdated event raised by the ProviderRegistry contract.
type ProviderRegistryReputationUpdated struct {
	Provider common.Address
	NewScore *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReputationUpdated is a free log retrieval operation binding the contract event 0xfc577563f1b9a0461e24abef1e1fcc0d33d3d881f20b5df6dda59de4aae2c821.
//
// Solidity: event ReputationUpdated(address indexed provider, uint256 newScore)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterReputationUpdated(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryReputationUpdatedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ReputationUpdated", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryReputationUpdatedIterator{contract: _ProviderRegistry.contract, event: "ReputationUpdated", logs: logs, sub: sub}, nil
}

// WatchReputationUpdated is a free log subscription operation binding the contract event 0xfc577563f1b9a0461e24abef1e1fcc0d33d3d881f20b5df6dda59de4aae2c821.
//
// Solidity: event ReputationUpdated(address indexed provider, uint256 newScore)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchReputationUpdated(opts *bind.WatchOpts, sink chan<- *ProviderRegistryReputationUpdated, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ReputationUpdated", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryReputationUpdated)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
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

// ParseReputationUpdated is a log parse operation binding the contract event 0xfc577563f1b9a0461e24abef1e1fcc0d33d3d881f20b5df6dda59de4aae2c821.
//
// Solidity: event ReputationUpdated(address indexed provider, uint256 newScore)
func (_ProviderRegistry *ProviderRegistryFilterer) ParseReputationUpdated(log types.Log) (*ProviderRegistryReputationUpdated, error) {
	event := new(ProviderRegistryReputationUpdated)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
