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

// EscrowEscrowEntry is an auto generated low-level Go binding around an user-defined struct.
type EscrowEscrowEntry struct {
	Client   common.Address
	Amount   *big.Int
	Released bool
}

// EscrowMetaData contains all meta data concerning the Escrow contract.
var EscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Released\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"entries\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"released\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"}],\"name\":\"getEntry\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"released\",\"type\":\"bool\"}],\"internalType\":\"structEscrow.EscrowEntry\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"jobId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"marketplaceAddr\",\"type\":\"address\"}],\"name\":\"setMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// EscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use EscrowMetaData.ABI instead.
var EscrowABI = EscrowMetaData.ABI

// Escrow is an auto generated Go binding around an Ethereum contract.
type Escrow struct {
	EscrowCaller     // Read-only binding to the contract
	EscrowTransactor // Write-only binding to the contract
	EscrowFilterer   // Log filterer for contract events
}

// EscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type EscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EscrowSession struct {
	Contract     *Escrow           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EscrowCallerSession struct {
	Contract *EscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EscrowTransactorSession struct {
	Contract     *EscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type EscrowRaw struct {
	Contract *Escrow // Generic contract binding to access the raw methods on
}

// EscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EscrowCallerRaw struct {
	Contract *EscrowCaller // Generic read-only contract binding to access the raw methods on
}

// EscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EscrowTransactorRaw struct {
	Contract *EscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEscrow creates a new instance of Escrow, bound to a specific deployed contract.
func NewEscrow(address common.Address, backend bind.ContractBackend) (*Escrow, error) {
	contract, err := bindEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Escrow{EscrowCaller: EscrowCaller{contract: contract}, EscrowTransactor: EscrowTransactor{contract: contract}, EscrowFilterer: EscrowFilterer{contract: contract}}, nil
}

// NewEscrowCaller creates a new read-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowCaller(address common.Address, caller bind.ContractCaller) (*EscrowCaller, error) {
	contract, err := bindEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowCaller{contract: contract}, nil
}

// NewEscrowTransactor creates a new write-only instance of Escrow, bound to a specific deployed contract.
func NewEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*EscrowTransactor, error) {
	contract, err := bindEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EscrowTransactor{contract: contract}, nil
}

// NewEscrowFilterer creates a new log filterer instance of Escrow, bound to a specific deployed contract.
func NewEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*EscrowFilterer, error) {
	contract, err := bindEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EscrowFilterer{contract: contract}, nil
}

// bindEscrow binds a generic wrapper to an already deployed contract.
func bindEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.EscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.EscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Escrow *EscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Escrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Escrow *EscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Escrow *EscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Escrow.Contract.contract.Transact(opts, method, params...)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(address client, uint256 amount, bool released)
func (_Escrow *EscrowCaller) Entries(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Client   common.Address
	Amount   *big.Int
	Released bool
}, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "entries", arg0)

	outstruct := new(struct {
		Client   common.Address
		Amount   *big.Int
		Released bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Client = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Released = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(address client, uint256 amount, bool released)
func (_Escrow *EscrowSession) Entries(arg0 *big.Int) (struct {
	Client   common.Address
	Amount   *big.Int
	Released bool
}, error) {
	return _Escrow.Contract.Entries(&_Escrow.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xb30906d4.
//
// Solidity: function entries(uint256 ) view returns(address client, uint256 amount, bool released)
func (_Escrow *EscrowCallerSession) Entries(arg0 *big.Int) (struct {
	Client   common.Address
	Amount   *big.Int
	Released bool
}, error) {
	return _Escrow.Contract.Entries(&_Escrow.CallOpts, arg0)
}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 jobId) view returns((address,uint256,bool))
func (_Escrow *EscrowCaller) GetEntry(opts *bind.CallOpts, jobId *big.Int) (EscrowEscrowEntry, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "getEntry", jobId)

	if err != nil {
		return *new(EscrowEscrowEntry), err
	}

	out0 := *abi.ConvertType(out[0], new(EscrowEscrowEntry)).(*EscrowEscrowEntry)

	return out0, err

}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 jobId) view returns((address,uint256,bool))
func (_Escrow *EscrowSession) GetEntry(jobId *big.Int) (EscrowEscrowEntry, error) {
	return _Escrow.Contract.GetEntry(&_Escrow.CallOpts, jobId)
}

// GetEntry is a free data retrieval call binding the contract method 0xbae78d7b.
//
// Solidity: function getEntry(uint256 jobId) view returns((address,uint256,bool))
func (_Escrow *EscrowCallerSession) GetEntry(jobId *big.Int) (EscrowEscrowEntry, error) {
	return _Escrow.Contract.GetEntry(&_Escrow.CallOpts, jobId)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_Escrow *EscrowCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Escrow.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_Escrow *EscrowSession) Marketplace() (common.Address, error) {
	return _Escrow.Contract.Marketplace(&_Escrow.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_Escrow *EscrowCallerSession) Marketplace() (common.Address, error) {
	return _Escrow.Contract.Marketplace(&_Escrow.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 jobId, address client) payable returns()
func (_Escrow *EscrowTransactor) Deposit(opts *bind.TransactOpts, jobId *big.Int, client common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "deposit", jobId, client)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 jobId, address client) payable returns()
func (_Escrow *EscrowSession) Deposit(jobId *big.Int, client common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Deposit(&_Escrow.TransactOpts, jobId, client)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 jobId, address client) payable returns()
func (_Escrow *EscrowTransactorSession) Deposit(jobId *big.Int, client common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Deposit(&_Escrow.TransactOpts, jobId, client)
}

// Refund is a paid mutator transaction binding the contract method 0x7ad226dc.
//
// Solidity: function refund(uint256 jobId, address client) returns()
func (_Escrow *EscrowTransactor) Refund(opts *bind.TransactOpts, jobId *big.Int, client common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "refund", jobId, client)
}

// Refund is a paid mutator transaction binding the contract method 0x7ad226dc.
//
// Solidity: function refund(uint256 jobId, address client) returns()
func (_Escrow *EscrowSession) Refund(jobId *big.Int, client common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Refund(&_Escrow.TransactOpts, jobId, client)
}

// Refund is a paid mutator transaction binding the contract method 0x7ad226dc.
//
// Solidity: function refund(uint256 jobId, address client) returns()
func (_Escrow *EscrowTransactorSession) Refund(jobId *big.Int, client common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Refund(&_Escrow.TransactOpts, jobId, client)
}

// Release is a paid mutator transaction binding the contract method 0x8124fea6.
//
// Solidity: function release(uint256 jobId, address recipient) returns()
func (_Escrow *EscrowTransactor) Release(opts *bind.TransactOpts, jobId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "release", jobId, recipient)
}

// Release is a paid mutator transaction binding the contract method 0x8124fea6.
//
// Solidity: function release(uint256 jobId, address recipient) returns()
func (_Escrow *EscrowSession) Release(jobId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Release(&_Escrow.TransactOpts, jobId, recipient)
}

// Release is a paid mutator transaction binding the contract method 0x8124fea6.
//
// Solidity: function release(uint256 jobId, address recipient) returns()
func (_Escrow *EscrowTransactorSession) Release(jobId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.Release(&_Escrow.TransactOpts, jobId, recipient)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address marketplaceAddr) returns()
func (_Escrow *EscrowTransactor) SetMarketplace(opts *bind.TransactOpts, marketplaceAddr common.Address) (*types.Transaction, error) {
	return _Escrow.contract.Transact(opts, "setMarketplace", marketplaceAddr)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address marketplaceAddr) returns()
func (_Escrow *EscrowSession) SetMarketplace(marketplaceAddr common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetMarketplace(&_Escrow.TransactOpts, marketplaceAddr)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address marketplaceAddr) returns()
func (_Escrow *EscrowTransactorSession) SetMarketplace(marketplaceAddr common.Address) (*types.Transaction, error) {
	return _Escrow.Contract.SetMarketplace(&_Escrow.TransactOpts, marketplaceAddr)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Escrow *EscrowTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Escrow.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Escrow *EscrowSession) Receive() (*types.Transaction, error) {
	return _Escrow.Contract.Receive(&_Escrow.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Escrow *EscrowTransactorSession) Receive() (*types.Transaction, error) {
	return _Escrow.Contract.Receive(&_Escrow.TransactOpts)
}

// EscrowDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Escrow contract.
type EscrowDepositedIterator struct {
	Event *EscrowDeposited // Event containing the contract specifics and raw log

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
func (it *EscrowDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowDeposited)
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
		it.Event = new(EscrowDeposited)
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
func (it *EscrowDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowDeposited represents a Deposited event raised by the Escrow contract.
type EscrowDeposited struct {
	JobId  *big.Int
	Client common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x1599c0fcf897af5babc2bfcf707f5dc050f841b044d97c3251ecec35b9abf80b.
//
// Solidity: event Deposited(uint256 indexed jobId, address indexed client, uint256 amount)
func (_Escrow *EscrowFilterer) FilterDeposited(opts *bind.FilterOpts, jobId []*big.Int, client []common.Address) (*EscrowDepositedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Deposited", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return &EscrowDepositedIterator{contract: _Escrow.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x1599c0fcf897af5babc2bfcf707f5dc050f841b044d97c3251ecec35b9abf80b.
//
// Solidity: event Deposited(uint256 indexed jobId, address indexed client, uint256 amount)
func (_Escrow *EscrowFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EscrowDeposited, jobId []*big.Int, client []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Deposited", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowDeposited)
				if err := _Escrow.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x1599c0fcf897af5babc2bfcf707f5dc050f841b044d97c3251ecec35b9abf80b.
//
// Solidity: event Deposited(uint256 indexed jobId, address indexed client, uint256 amount)
func (_Escrow *EscrowFilterer) ParseDeposited(log types.Log) (*EscrowDeposited, error) {
	event := new(EscrowDeposited)
	if err := _Escrow.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the Escrow contract.
type EscrowRefundedIterator struct {
	Event *EscrowRefunded // Event containing the contract specifics and raw log

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
func (it *EscrowRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowRefunded)
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
		it.Event = new(EscrowRefunded)
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
func (it *EscrowRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowRefunded represents a Refunded event raised by the Escrow contract.
type EscrowRefunded struct {
	JobId  *big.Int
	Client common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x7ca5472b7ea78c2c0141c5a12ee6d170cf4ce8ed06be3d22c8252ddfc7a6a2c4.
//
// Solidity: event Refunded(uint256 indexed jobId, address indexed client, uint256 amount)
func (_Escrow *EscrowFilterer) FilterRefunded(opts *bind.FilterOpts, jobId []*big.Int, client []common.Address) (*EscrowRefundedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Refunded", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return &EscrowRefundedIterator{contract: _Escrow.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x7ca5472b7ea78c2c0141c5a12ee6d170cf4ce8ed06be3d22c8252ddfc7a6a2c4.
//
// Solidity: event Refunded(uint256 indexed jobId, address indexed client, uint256 amount)
func (_Escrow *EscrowFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *EscrowRefunded, jobId []*big.Int, client []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Refunded", jobIdRule, clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowRefunded)
				if err := _Escrow.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x7ca5472b7ea78c2c0141c5a12ee6d170cf4ce8ed06be3d22c8252ddfc7a6a2c4.
//
// Solidity: event Refunded(uint256 indexed jobId, address indexed client, uint256 amount)
func (_Escrow *EscrowFilterer) ParseRefunded(log types.Log) (*EscrowRefunded, error) {
	event := new(EscrowRefunded)
	if err := _Escrow.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EscrowReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the Escrow contract.
type EscrowReleasedIterator struct {
	Event *EscrowReleased // Event containing the contract specifics and raw log

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
func (it *EscrowReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EscrowReleased)
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
		it.Event = new(EscrowReleased)
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
func (it *EscrowReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EscrowReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EscrowReleased represents a Released event raised by the Escrow contract.
type EscrowReleased struct {
	JobId     *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0x3bfce8de0db7450cc169b94323c210e69a36c6a4a58c9f5d96bec4973adce392.
//
// Solidity: event Released(uint256 indexed jobId, address indexed recipient, uint256 amount)
func (_Escrow *EscrowFilterer) FilterReleased(opts *bind.FilterOpts, jobId []*big.Int, recipient []common.Address) (*EscrowReleasedIterator, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Escrow.contract.FilterLogs(opts, "Released", jobIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &EscrowReleasedIterator{contract: _Escrow.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0x3bfce8de0db7450cc169b94323c210e69a36c6a4a58c9f5d96bec4973adce392.
//
// Solidity: event Released(uint256 indexed jobId, address indexed recipient, uint256 amount)
func (_Escrow *EscrowFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *EscrowReleased, jobId []*big.Int, recipient []common.Address) (event.Subscription, error) {

	var jobIdRule []interface{}
	for _, jobIdItem := range jobId {
		jobIdRule = append(jobIdRule, jobIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Escrow.contract.WatchLogs(opts, "Released", jobIdRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EscrowReleased)
				if err := _Escrow.contract.UnpackLog(event, "Released", log); err != nil {
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

// ParseReleased is a log parse operation binding the contract event 0x3bfce8de0db7450cc169b94323c210e69a36c6a4a58c9f5d96bec4973adce392.
//
// Solidity: event Released(uint256 indexed jobId, address indexed recipient, uint256 amount)
func (_Escrow *EscrowFilterer) ParseReleased(log types.Log) (*EscrowReleased, error) {
	event := new(EscrowReleased)
	if err := _Escrow.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
