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

// MarketplaceJob is an auto generated low-level Go binding around an user-defined struct.
type MarketplaceJob struct {
	JobId            *big.Int
	Client           common.Address
	AssignedProvider common.Address
	CpuRequired      *big.Int
	MemoryRequired   *big.Int
	MaxDuration      *big.Int
	MaxPricePerHour  *big.Int
	Status           uint8
	CreatedAt        *big.Int
	StartedAt        *big.Int
	CompletedAt      *big.Int
	ResultHash       [32]byte
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddr\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"escrowAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"schedulerAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"JobAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"JobCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resultHash\",\"type\":\"bytes32\"}],\"name\":\"JobCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"JobFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cpuRequired\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"memoryRequired\",\"type\":\"uint256\"}],\"name\":\"JobPosted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"JobStarted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"assignJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"cancelJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"resultHash\",\"type\":\"bytes32\"}],\"name\":\"completeJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"escrow\",\"outputs\":[{\"internalType\":\"contractEscrow\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"failJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"getJob\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assignedProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cpuRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPricePerHour\",\"type\":\"uint256\"},{\"internalType\":\"enumMarketplace.JobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"completedAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"resultHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMarketplace.Job\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpenJobs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assignedProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cpuRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPricePerHour\",\"type\":\"uint256\"},{\"internalType\":\"enumMarketplace.JobStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"completedAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"resultHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextJobId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"openJobs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cpuRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"memoryRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPricePerHour\",\"type\":\"uint256\"}],\"name\":\"postJob\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractProviderRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"startJob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_Marketplace *MarketplaceCaller) Escrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_Marketplace *MarketplaceSession) Escrow() (common.Address, error) {
	return _Marketplace.Contract.Escrow(&_Marketplace.CallOpts)
}

// Escrow is a free data retrieval call binding the contract method 0xe2fdcc17.
//
// Solidity: function escrow() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Escrow() (common.Address, error) {
	return _Marketplace.Contract.Escrow(&_Marketplace.CallOpts)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint8,uint256,uint256,uint256,bytes32))
func (_Marketplace *MarketplaceCaller) GetJob(opts *bind.CallOpts, jobId *big.Int) (MarketplaceJob, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getJob", jobId)

	if err != nil {
		return *new(MarketplaceJob), err
	}

	out0 := *abi.ConvertType(out[0], new(MarketplaceJob)).(*MarketplaceJob)

	return out0, err

}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint8,uint256,uint256,uint256,bytes32))
func (_Marketplace *MarketplaceSession) GetJob(jobId *big.Int) (MarketplaceJob, error) {
	return _Marketplace.Contract.GetJob(&_Marketplace.CallOpts, jobId)
}

// GetJob is a free data retrieval call binding the contract method 0xbf22c457.
//
// Solidity: function getJob(uint256 jobId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint8,uint256,uint256,uint256,bytes32))
func (_Marketplace *MarketplaceCallerSession) GetJob(jobId *big.Int) (MarketplaceJob, error) {
	return _Marketplace.Contract.GetJob(&_Marketplace.CallOpts, jobId)
}

// GetOpenJobs is a free data retrieval call binding the contract method 0x1321c95e.
//
// Solidity: function getOpenJobs() view returns(uint256[])
func (_Marketplace *MarketplaceCaller) GetOpenJobs(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getOpenJobs")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOpenJobs is a free data retrieval call binding the contract method 0x1321c95e.
//
// Solidity: function getOpenJobs() view returns(uint256[])
func (_Marketplace *MarketplaceSession) GetOpenJobs() ([]*big.Int, error) {
	return _Marketplace.Contract.GetOpenJobs(&_Marketplace.CallOpts)
}

// GetOpenJobs is a free data retrieval call binding the contract method 0x1321c95e.
//
// Solidity: function getOpenJobs() view returns(uint256[])
func (_Marketplace *MarketplaceCallerSession) GetOpenJobs() ([]*big.Int, error) {
	return _Marketplace.Contract.GetOpenJobs(&_Marketplace.CallOpts)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 jobId, address client, address assignedProvider, uint256 cpuRequired, uint256 memoryRequired, uint256 maxDuration, uint256 maxPricePerHour, uint8 status, uint256 createdAt, uint256 startedAt, uint256 completedAt, bytes32 resultHash)
func (_Marketplace *MarketplaceCaller) Jobs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	JobId            *big.Int
	Client           common.Address
	AssignedProvider common.Address
	CpuRequired      *big.Int
	MemoryRequired   *big.Int
	MaxDuration      *big.Int
	MaxPricePerHour  *big.Int
	Status           uint8
	CreatedAt        *big.Int
	StartedAt        *big.Int
	CompletedAt      *big.Int
	ResultHash       [32]byte
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "jobs", arg0)

	outstruct := new(struct {
		JobId            *big.Int
		Client           common.Address
		AssignedProvider common.Address
		CpuRequired      *big.Int
		MemoryRequired   *big.Int
		MaxDuration      *big.Int
		MaxPricePerHour  *big.Int
		Status           uint8
		CreatedAt        *big.Int
		StartedAt        *big.Int
		CompletedAt      *big.Int
		ResultHash       [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.JobId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Client = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AssignedProvider = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CpuRequired = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MemoryRequired = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.MaxDuration = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MaxPricePerHour = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.CreatedAt = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.CompletedAt = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.ResultHash = *abi.ConvertType(out[11], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 jobId, address client, address assignedProvider, uint256 cpuRequired, uint256 memoryRequired, uint256 maxDuration, uint256 maxPricePerHour, uint8 status, uint256 createdAt, uint256 startedAt, uint256 completedAt, bytes32 resultHash)
func (_Marketplace *MarketplaceSession) Jobs(arg0 *big.Int) (struct {
	JobId            *big.Int
	Client           common.Address
	AssignedProvider common.Address
	CpuRequired      *big.Int
	MemoryRequired   *big.Int
	MaxDuration      *big.Int
	MaxPricePerHour  *big.Int
	Status           uint8
	CreatedAt        *big.Int
	StartedAt        *big.Int
	CompletedAt      *big.Int
	ResultHash       [32]byte
}, error) {
	return _Marketplace.Contract.Jobs(&_Marketplace.CallOpts, arg0)
}

// Jobs is a free data retrieval call binding the contract method 0x180aedf3.
//
// Solidity: function jobs(uint256 ) view returns(uint256 jobId, address client, address assignedProvider, uint256 cpuRequired, uint256 memoryRequired, uint256 maxDuration, uint256 maxPricePerHour, uint8 status, uint256 createdAt, uint256 startedAt, uint256 completedAt, bytes32 resultHash)
func (_Marketplace *MarketplaceCallerSession) Jobs(arg0 *big.Int) (struct {
	JobId            *big.Int
	Client           common.Address
	AssignedProvider common.Address
	CpuRequired      *big.Int
	MemoryRequired   *big.Int
	MaxDuration      *big.Int
	MaxPricePerHour  *big.Int
	Status           uint8
	CreatedAt        *big.Int
	StartedAt        *big.Int
	CompletedAt      *big.Int
	ResultHash       [32]byte
}, error) {
	return _Marketplace.Contract.Jobs(&_Marketplace.CallOpts, arg0)
}

// NextJobId is a free data retrieval call binding the contract method 0xb0c2aa5e.
//
// Solidity: function nextJobId() view returns(uint256)
func (_Marketplace *MarketplaceCaller) NextJobId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "nextJobId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextJobId is a free data retrieval call binding the contract method 0xb0c2aa5e.
//
// Solidity: function nextJobId() view returns(uint256)
func (_Marketplace *MarketplaceSession) NextJobId() (*big.Int, error) {
	return _Marketplace.Contract.NextJobId(&_Marketplace.CallOpts)
}

// NextJobId is a free data retrieval call binding the contract method 0xb0c2aa5e.
//
// Solidity: function nextJobId() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) NextJobId() (*big.Int, error) {
	return _Marketplace.Contract.NextJobId(&_Marketplace.CallOpts)
}

// OpenJobs is a free data retrieval call binding the contract method 0xef25bd8f.
//
// Solidity: function openJobs(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) OpenJobs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "openJobs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpenJobs is a free data retrieval call binding the contract method 0xef25bd8f.
//
// Solidity: function openJobs(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceSession) OpenJobs(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.OpenJobs(&_Marketplace.CallOpts, arg0)
}

// OpenJobs is a free data retrieval call binding the contract method 0xef25bd8f.
//
// Solidity: function openJobs(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) OpenJobs(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.OpenJobs(&_Marketplace.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Marketplace *MarketplaceCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Marketplace *MarketplaceSession) Registry() (common.Address, error) {
	return _Marketplace.Contract.Registry(&_Marketplace.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Registry() (common.Address, error) {
	return _Marketplace.Contract.Registry(&_Marketplace.CallOpts)
}

// Scheduler is a free data retrieval call binding the contract method 0xd1ad17bf.
//
// Solidity: function scheduler() view returns(address)
func (_Marketplace *MarketplaceCaller) Scheduler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "scheduler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Scheduler is a free data retrieval call binding the contract method 0xd1ad17bf.
//
// Solidity: function scheduler() view returns(address)
func (_Marketplace *MarketplaceSession) Scheduler() (common.Address, error) {
	return _Marketplace.Contract.Scheduler(&_Marketplace.CallOpts)
}

// Scheduler is a free data retrieval call binding the contract method 0xd1ad17bf.
//
// Solidity: function scheduler() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Scheduler() (common.Address, error) {
	return _Marketplace.Contract.Scheduler(&_Marketplace.CallOpts)
}

// AssignJob is a paid mutator transaction binding the contract method 0xf3a4735a.
//
// Solidity: function assignJob(uint256 jobId, address provider) returns()
func (_Marketplace *MarketplaceTransactor) AssignJob(opts *bind.TransactOpts, jobId *big.Int, provider common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "assignJob", jobId, provider)
}

// AssignJob is a paid mutator transaction binding the contract method 0xf3a4735a.
//
// Solidity: function assignJob(uint256 jobId, address provider) returns()
func (_Marketplace *MarketplaceSession) AssignJob(jobId *big.Int, provider common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.AssignJob(&_Marketplace.TransactOpts, jobId, provider)
}

// AssignJob is a paid mutator transaction binding the contract method 0xf3a4735a.
//
// Solidity: function assignJob(uint256 jobId, address provider) returns()
func (_Marketplace *MarketplaceTransactorSession) AssignJob(jobId *big.Int, provider common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.AssignJob(&_Marketplace.TransactOpts, jobId, provider)
}

// CancelJob is a paid mutator transaction binding the contract method 0x1dffa3dc.
//
// Solidity: function cancelJob(uint256 jobId) returns()
func (_Marketplace *MarketplaceTransactor) CancelJob(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelJob", jobId)
}

// CancelJob is a paid mutator transaction binding the contract method 0x1dffa3dc.
//
// Solidity: function cancelJob(uint256 jobId) returns()
func (_Marketplace *MarketplaceSession) CancelJob(jobId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelJob(&_Marketplace.TransactOpts, jobId)
}

// CancelJob is a paid mutator transaction binding the contract method 0x1dffa3dc.
//
// Solidity: function cancelJob(uint256 jobId) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelJob(jobId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelJob(&_Marketplace.TransactOpts, jobId)
}

// CompleteJob is a paid mutator transaction binding the contract method 0xebb59fad.
//
// Solidity: function completeJob(uint256 jobId, bytes32 resultHash) returns()
func (_Marketplace *MarketplaceTransactor) CompleteJob(opts *bind.TransactOpts, jobId *big.Int, resultHash [32]byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "completeJob", jobId, resultHash)
}

// CompleteJob is a paid mutator transaction binding the contract method 0xebb59fad.
//
// Solidity: function completeJob(uint256 jobId, bytes32 resultHash) returns()
func (_Marketplace *MarketplaceSession) CompleteJob(jobId *big.Int, resultHash [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.CompleteJob(&_Marketplace.TransactOpts, jobId, resultHash)
}

// CompleteJob is a paid mutator transaction binding the contract method 0xebb59fad.
//
// Solidity: function completeJob(uint256 jobId, bytes32 resultHash) returns()
func (_Marketplace *MarketplaceTransactorSession) CompleteJob(jobId *big.Int, resultHash [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.CompleteJob(&_Marketplace.TransactOpts, jobId, resultHash)
}

// FailJob is a paid mutator transaction binding the contract method 0x420bb38f.
//
// Solidity: function failJob(uint256 jobId) returns()
func (_Marketplace *MarketplaceTransactor) FailJob(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "failJob", jobId)
}

// FailJob is a paid mutator transaction binding the contract method 0x420bb38f.
//
// Solidity: function failJob(uint256 jobId) returns()
func (_Marketplace *MarketplaceSession) FailJob(jobId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.FailJob(&_Marketplace.TransactOpts, jobId)
}

// FailJob is a paid mutator transaction binding the contract method 0x420bb38f.
//
// Solidity: function failJob(uint256 jobId) returns()
func (_Marketplace *MarketplaceTransactorSession) FailJob(jobId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.FailJob(&_Marketplace.TransactOpts, jobId)
}

// PostJob is a paid mutator transaction binding the contract method 0xd9ad702a.
//
// Solidity: function postJob(uint256 cpuRequired, uint256 memoryRequired, uint256 maxDuration, uint256 maxPricePerHour) payable returns(uint256)
func (_Marketplace *MarketplaceTransactor) PostJob(opts *bind.TransactOpts, cpuRequired *big.Int, memoryRequired *big.Int, maxDuration *big.Int, maxPricePerHour *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "postJob", cpuRequired, memoryRequired, maxDuration, maxPricePerHour)
}

// PostJob is a paid mutator transaction binding the contract method 0xd9ad702a.
//
// Solidity: function postJob(uint256 cpuRequired, uint256 memoryRequired, uint256 maxDuration, uint256 maxPricePerHour) payable returns(uint256)
func (_Marketplace *MarketplaceSession) PostJob(cpuRequired *big.Int, memoryRequired *big.Int, maxDuration *big.Int, maxPricePerHour *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PostJob(&_Marketplace.TransactOpts, cpuRequired, memoryRequired, maxDuration, maxPricePerHour)
}

// PostJob is a paid mutator transaction binding the contract method 0xd9ad702a.
//
// Solidity: function postJob(uint256 cpuRequired, uint256 memoryRequired, uint256 maxDuration, uint256 maxPricePerHour) payable returns(uint256)
func (_Marketplace *MarketplaceTransactorSession) PostJob(cpuRequired *big.Int, memoryRequired *big.Int, maxDuration *big.Int, maxPricePerHour *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PostJob(&_Marketplace.TransactOpts, cpuRequired, memoryRequired, maxDuration, maxPricePerHour)
}

// StartJob is a paid mutator transaction binding the contract method 0xe1255294.
//
// Solidity: function startJob(uint256 jobId) returns()
func (_Marketplace *MarketplaceTransactor) StartJob(opts *bind.TransactOpts, jobId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "startJob", jobId)
}

// StartJob is a paid mutator transaction binding the contract method 0xe1255294.
//
// Solidity: function startJob(uint256 jobId) returns()
func (_Marketplace *MarketplaceSession) StartJob(jobId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.StartJob(&_Marketplace.TransactOpts, jobId)
}

// StartJob is a paid mutator transaction binding the contract method 0xe1255294.
//
// Solidity: function startJob(uint256 jobId) returns()
func (_Marketplace *MarketplaceTransactorSession) StartJob(jobId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.StartJob(&_Marketplace.TransactOpts, jobId)
}

// MarketplaceJobAssignedIterator is returned from FilterJobAssigned and is used to iterate over the raw logs and unpacked data for JobAssigned events raised by the Marketplace contract.
type MarketplaceJobAssignedIterator struct {
	Event *MarketplaceJobAssigned // Event containing the contract specifics and raw log

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
func (it *MarketplaceJobAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceJobAssigned)
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
		it.Event = new(MarketplaceJobAssigned)
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
func (it *MarketplaceJobAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceJobAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceJobAssigned represents a JobAssigned event raised by the Marketplace contract.
type MarketplaceJobAssigned struct {
	JobId    *big.Int
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterJobAssigned is a free log retrieval operation binding the contract event 0x3738a681a43a0b6a554742a08bae57e0215e8a774217dc2e7de6c9411b6d0ddc.
//
// Solidity: event JobAssigned(uint256 indexed jobId, address indexed provider)
func (_Marketplace *MarketplaceFilterer) FilterJobAssigned(opts *bind.FilterOpts, jobId []*big.Int, provider []common.Address) (*MarketplaceJobAssignedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "JobAssigned", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceJobAssignedIterator{contract: _Marketplace.contract, event: "JobAssigned", logs: logs, sub: sub}, nil
}

// WatchJobAssigned is a free log subscription operation binding the contract event 0x3738a681a43a0b6a554742a08bae57e0215e8a774217dc2e7de6c9411b6d0ddc.
//
// Solidity: event JobAssigned(uint256 indexed jobId, address indexed provider)
func (_Marketplace *MarketplaceFilterer) WatchJobAssigned(opts *bind.WatchOpts, sink chan<- *MarketplaceJobAssigned, jobId []*big.Int, provider []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "JobAssigned", jobIdRule, providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceJobAssigned)
				if err := _Marketplace.contract.UnpackLog(event, "JobAssigned", log); err != nil {
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

// ParseJobAssigned is a log parse operation binding the contract event 0x3738a681a43a0b6a554742a08bae57e0215e8a774217dc2e7de6c9411b6d0ddc.
//
// Solidity: event JobAssigned(uint256 indexed jobId, address indexed provider)
func (_Marketplace *MarketplaceFilterer) ParseJobAssigned(log types.Log) (*MarketplaceJobAssigned, error) {
	event := new(MarketplaceJobAssigned)
	if err := _Marketplace.contract.UnpackLog(event, "JobAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceJobCancelledIterator is returned from FilterJobCancelled and is used to iterate over the raw logs and unpacked data for JobCancelled events raised by the Marketplace contract.
type MarketplaceJobCancelledIterator struct {
	Event *MarketplaceJobCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplaceJobCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceJobCancelled)
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
		it.Event = new(MarketplaceJobCancelled)
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
func (it *MarketplaceJobCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceJobCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceJobCancelled represents a JobCancelled event raised by the Marketplace contract.
type MarketplaceJobCancelled struct {
	JobId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterJobCancelled is a free log retrieval operation binding the contract event 0x58243f872c550ad02802199ade99eb4250f6bb19e9abbf2d3de7b969c0eb5e4f.
//
// Solidity: event JobCancelled(uint256 indexed jobId)
func (_Marketplace *MarketplaceFilterer) FilterJobCancelled(opts *bind.FilterOpts, jobId []*big.Int) (*MarketplaceJobCancelledIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "JobCancelled", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceJobCancelledIterator{contract: _Marketplace.contract, event: "JobCancelled", logs: logs, sub: sub}, nil
}

// WatchJobCancelled is a free log subscription operation binding the contract event 0x58243f872c550ad02802199ade99eb4250f6bb19e9abbf2d3de7b969c0eb5e4f.
//
// Solidity: event JobCancelled(uint256 indexed jobId)
func (_Marketplace *MarketplaceFilterer) WatchJobCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceJobCancelled, jobId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "JobCancelled", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceJobCancelled)
				if err := _Marketplace.contract.UnpackLog(event, "JobCancelled", log); err != nil {
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

// ParseJobCancelled is a log parse operation binding the contract event 0x58243f872c550ad02802199ade99eb4250f6bb19e9abbf2d3de7b969c0eb5e4f.
//
// Solidity: event JobCancelled(uint256 indexed jobId)
func (_Marketplace *MarketplaceFilterer) ParseJobCancelled(log types.Log) (*MarketplaceJobCancelled, error) {
	event := new(MarketplaceJobCancelled)
	if err := _Marketplace.contract.UnpackLog(event, "JobCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceJobCompletedIterator is returned from FilterJobCompleted and is used to iterate over the raw logs and unpacked data for JobCompleted events raised by the Marketplace contract.
type MarketplaceJobCompletedIterator struct {
	Event *MarketplaceJobCompleted // Event containing the contract specifics and raw log

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
func (it *MarketplaceJobCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceJobCompleted)
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
		it.Event = new(MarketplaceJobCompleted)
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
func (it *MarketplaceJobCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceJobCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceJobCompleted represents a JobCompleted event raised by the Marketplace contract.
type MarketplaceJobCompleted struct {
	JobId      *big.Int
	ResultHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterJobCompleted is a free log retrieval operation binding the contract event 0x45c386dc6524a2d9fe630455323c6a39f557c52ab01e886deee20a0b538147ac.
//
// Solidity: event JobCompleted(uint256 indexed jobId, bytes32 resultHash)
func (_Marketplace *MarketplaceFilterer) FilterJobCompleted(opts *bind.FilterOpts, jobId []*big.Int) (*MarketplaceJobCompletedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "JobCompleted", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceJobCompletedIterator{contract: _Marketplace.contract, event: "JobCompleted", logs: logs, sub: sub}, nil
}

// WatchJobCompleted is a free log subscription operation binding the contract event 0x45c386dc6524a2d9fe630455323c6a39f557c52ab01e886deee20a0b538147ac.
//
// Solidity: event JobCompleted(uint256 indexed jobId, bytes32 resultHash)
func (_Marketplace *MarketplaceFilterer) WatchJobCompleted(opts *bind.WatchOpts, sink chan<- *MarketplaceJobCompleted, jobId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "JobCompleted", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceJobCompleted)
				if err := _Marketplace.contract.UnpackLog(event, "JobCompleted", log); err != nil {
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

// ParseJobCompleted is a log parse operation binding the contract event 0x45c386dc6524a2d9fe630455323c6a39f557c52ab01e886deee20a0b538147ac.
//
// Solidity: event JobCompleted(uint256 indexed jobId, bytes32 resultHash)
func (_Marketplace *MarketplaceFilterer) ParseJobCompleted(log types.Log) (*MarketplaceJobCompleted, error) {
	event := new(MarketplaceJobCompleted)
	if err := _Marketplace.contract.UnpackLog(event, "JobCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceJobFailedIterator is returned from FilterJobFailed and is used to iterate over the raw logs and unpacked data for JobFailed events raised by the Marketplace contract.
type MarketplaceJobFailedIterator struct {
	Event *MarketplaceJobFailed // Event containing the contract specifics and raw log

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
func (it *MarketplaceJobFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceJobFailed)
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
		it.Event = new(MarketplaceJobFailed)
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
func (it *MarketplaceJobFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceJobFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceJobFailed represents a JobFailed event raised by the Marketplace contract.
type MarketplaceJobFailed struct {
	JobId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterJobFailed is a free log retrieval operation binding the contract event 0xb75cf56c3b87307029f0059a671fedc39db6212ba4ef07f39d4f053185dc9006.
//
// Solidity: event JobFailed(uint256 indexed jobId)
func (_Marketplace *MarketplaceFilterer) FilterJobFailed(opts *bind.FilterOpts, jobId []*big.Int) (*MarketplaceJobFailedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "JobFailed", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceJobFailedIterator{contract: _Marketplace.contract, event: "JobFailed", logs: logs, sub: sub}, nil
}

// WatchJobFailed is a free log subscription operation binding the contract event 0xb75cf56c3b87307029f0059a671fedc39db6212ba4ef07f39d4f053185dc9006.
//
// Solidity: event JobFailed(uint256 indexed jobId)
func (_Marketplace *MarketplaceFilterer) WatchJobFailed(opts *bind.WatchOpts, sink chan<- *MarketplaceJobFailed, jobId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "JobFailed", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceJobFailed)
				if err := _Marketplace.contract.UnpackLog(event, "JobFailed", log); err != nil {
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

// ParseJobFailed is a log parse operation binding the contract event 0xb75cf56c3b87307029f0059a671fedc39db6212ba4ef07f39d4f053185dc9006.
//
// Solidity: event JobFailed(uint256 indexed jobId)
func (_Marketplace *MarketplaceFilterer) ParseJobFailed(log types.Log) (*MarketplaceJobFailed, error) {
	event := new(MarketplaceJobFailed)
	if err := _Marketplace.contract.UnpackLog(event, "JobFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceJobPostedIterator is returned from FilterJobPosted and is used to iterate over the raw logs and unpacked data for JobPosted events raised by the Marketplace contract.
type MarketplaceJobPostedIterator struct {
	Event *MarketplaceJobPosted // Event containing the contract specifics and raw log

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
func (it *MarketplaceJobPostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceJobPosted)
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
		it.Event = new(MarketplaceJobPosted)
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
func (it *MarketplaceJobPostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceJobPostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceJobPosted represents a JobPosted event raised by the Marketplace contract.
type MarketplaceJobPosted struct {
	JobId          *big.Int
	Client         common.Address
	CpuRequired    *big.Int
	MemoryRequired *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterJobPosted is a free log retrieval operation binding the contract event 0xc395422cf0471e0b91fe523372800edb102f71150f46497121d9b80c3812534d.
//
// Solidity: event JobPosted(uint256 indexed jobId, address indexed client, uint256 cpuRequired, uint256 memoryRequired)
func (_Marketplace *MarketplaceFilterer) FilterJobPosted(opts *bind.FilterOpts, jobId []*big.Int, client []common.Address) (*MarketplaceJobPostedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "JobPosted", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceJobPostedIterator{contract: _Marketplace.contract, event: "JobPosted", logs: logs, sub: sub}, nil
}

// WatchJobPosted is a free log subscription operation binding the contract event 0xc395422cf0471e0b91fe523372800edb102f71150f46497121d9b80c3812534d.
//
// Solidity: event JobPosted(uint256 indexed jobId, address indexed client, uint256 cpuRequired, uint256 memoryRequired)
func (_Marketplace *MarketplaceFilterer) WatchJobPosted(opts *bind.WatchOpts, sink chan<- *MarketplaceJobPosted, jobId []*big.Int, client []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "JobPosted", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceJobPosted)
				if err := _Marketplace.contract.UnpackLog(event, "JobPosted", log); err != nil {
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

// ParseJobPosted is a log parse operation binding the contract event 0xc395422cf0471e0b91fe523372800edb102f71150f46497121d9b80c3812534d.
//
// Solidity: event JobPosted(uint256 indexed jobId, address indexed client, uint256 cpuRequired, uint256 memoryRequired)
func (_Marketplace *MarketplaceFilterer) ParseJobPosted(log types.Log) (*MarketplaceJobPosted, error) {
	event := new(MarketplaceJobPosted)
	if err := _Marketplace.contract.UnpackLog(event, "JobPosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceJobStartedIterator is returned from FilterJobStarted and is used to iterate over the raw logs and unpacked data for JobStarted events raised by the Marketplace contract.
type MarketplaceJobStartedIterator struct {
	Event *MarketplaceJobStarted // Event containing the contract specifics and raw log

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
func (it *MarketplaceJobStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceJobStarted)
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
		it.Event = new(MarketplaceJobStarted)
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
func (it *MarketplaceJobStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceJobStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceJobStarted represents a JobStarted event raised by the Marketplace contract.
type MarketplaceJobStarted struct {
	JobId *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterJobStarted is a free log retrieval operation binding the contract event 0x7e2409eb03fe1e2ea93add34ad736f8cddc41ffdc3c5875587b3e40ab490996a.
//
// Solidity: event JobStarted(uint256 indexed jobId)
func (_Marketplace *MarketplaceFilterer) FilterJobStarted(opts *bind.FilterOpts, jobId []*big.Int) (*MarketplaceJobStartedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "JobStarted", jobIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceJobStartedIterator{contract: _Marketplace.contract, event: "JobStarted", logs: logs, sub: sub}, nil
}

// WatchJobStarted is a free log subscription operation binding the contract event 0x7e2409eb03fe1e2ea93add34ad736f8cddc41ffdc3c5875587b3e40ab490996a.
//
// Solidity: event JobStarted(uint256 indexed jobId)
func (_Marketplace *MarketplaceFilterer) WatchJobStarted(opts *bind.WatchOpts, sink chan<- *MarketplaceJobStarted, jobId []*big.Int) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "JobStarted", jobIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceJobStarted)
				if err := _Marketplace.contract.UnpackLog(event, "JobStarted", log); err != nil {
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

// ParseJobStarted is a log parse operation binding the contract event 0x7e2409eb03fe1e2ea93add34ad736f8cddc41ffdc3c5875587b3e40ab490996a.
//
// Solidity: event JobStarted(uint256 indexed jobId)
func (_Marketplace *MarketplaceFilterer) ParseJobStarted(log types.Log) (*MarketplaceJobStarted, error) {
	event := new(MarketplaceJobStarted)
	if err := _Marketplace.contract.UnpackLog(event, "JobStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
