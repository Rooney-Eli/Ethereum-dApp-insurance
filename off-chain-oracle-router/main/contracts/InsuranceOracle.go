// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package InsuranceOracle

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// InsuranceOracleABI is the input ABI used to generate the binding from.
const InsuranceOracleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"PolicyInitialized\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"declarePolicyActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPolicyIsClaimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setPolicyIsClaimable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"testConnection\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// InsuranceOracleBin is the compiled bytecode used for deploying new contracts.
var InsuranceOracleBin = "0x608060405260008060146101000a81548160ff021916908315150217905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506102d48061006d6000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80632dd308d9146100675780633412a15c1461009757806338cc4831146100b95780638da5cb5b146101035780639d8b3e241461014d578063f714c06c1461017d575b600080fd5b6100956004803603602081101561007d57600080fd5b8101908080351515906020019092919050505061019f565b005b61009f610215565b604051808215151515815260200191505060405180910390f35b6100c161021e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61010b610226565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61017b6004803603602081101561016357600080fd5b8101908080351515906020019092919050505061024b565b005b610185610289565b604051808215151515815260200191505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101f857600080fd5b80600060146101000a81548160ff02191690831515021790555050565b60006001905090565b600030905090565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f7f80f6442c3355ae66fd57fc95715c623b267bff31b0f17d8de16c010e0510e981604051808215151515815260200191505060405180910390a150565b60008060149054906101000a900460ff1690509056fea265627a7a72315820056e0a2703b526a51faed4974b3963892a1c565bb487768d4dfa152a0d55c93264736f6c63430005100032"

// DeployInsuranceOracle deploys a new Ethereum contract, binding an instance of InsuranceOracle to it.
func DeployInsuranceOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *InsuranceOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(InsuranceOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(InsuranceOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InsuranceOracle{InsuranceOracleCaller: InsuranceOracleCaller{contract: contract}, InsuranceOracleTransactor: InsuranceOracleTransactor{contract: contract}, InsuranceOracleFilterer: InsuranceOracleFilterer{contract: contract}}, nil
}

// InsuranceOracle is an auto generated Go binding around an Ethereum contract.
type InsuranceOracle struct {
	InsuranceOracleCaller     // Read-only binding to the contract
	InsuranceOracleTransactor // Write-only binding to the contract
	InsuranceOracleFilterer   // Log filterer for contract events
}

// InsuranceOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type InsuranceOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsuranceOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InsuranceOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsuranceOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InsuranceOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InsuranceOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InsuranceOracleSession struct {
	Contract     *InsuranceOracle  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InsuranceOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InsuranceOracleCallerSession struct {
	Contract *InsuranceOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// InsuranceOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InsuranceOracleTransactorSession struct {
	Contract     *InsuranceOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// InsuranceOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type InsuranceOracleRaw struct {
	Contract *InsuranceOracle // Generic contract binding to access the raw methods on
}

// InsuranceOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InsuranceOracleCallerRaw struct {
	Contract *InsuranceOracleCaller // Generic read-only contract binding to access the raw methods on
}

// InsuranceOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InsuranceOracleTransactorRaw struct {
	Contract *InsuranceOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInsuranceOracle creates a new instance of InsuranceOracle, bound to a specific deployed contract.
func NewInsuranceOracle(address common.Address, backend bind.ContractBackend) (*InsuranceOracle, error) {
	contract, err := bindInsuranceOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InsuranceOracle{InsuranceOracleCaller: InsuranceOracleCaller{contract: contract}, InsuranceOracleTransactor: InsuranceOracleTransactor{contract: contract}, InsuranceOracleFilterer: InsuranceOracleFilterer{contract: contract}}, nil
}

// NewInsuranceOracleCaller creates a new read-only instance of InsuranceOracle, bound to a specific deployed contract.
func NewInsuranceOracleCaller(address common.Address, caller bind.ContractCaller) (*InsuranceOracleCaller, error) {
	contract, err := bindInsuranceOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InsuranceOracleCaller{contract: contract}, nil
}

// NewInsuranceOracleTransactor creates a new write-only instance of InsuranceOracle, bound to a specific deployed contract.
func NewInsuranceOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*InsuranceOracleTransactor, error) {
	contract, err := bindInsuranceOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InsuranceOracleTransactor{contract: contract}, nil
}

// NewInsuranceOracleFilterer creates a new log filterer instance of InsuranceOracle, bound to a specific deployed contract.
func NewInsuranceOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*InsuranceOracleFilterer, error) {
	contract, err := bindInsuranceOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InsuranceOracleFilterer{contract: contract}, nil
}

// bindInsuranceOracle binds a generic wrapper to an already deployed contract.
func bindInsuranceOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InsuranceOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsuranceOracle *InsuranceOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InsuranceOracle.Contract.InsuranceOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsuranceOracle *InsuranceOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsuranceOracle.Contract.InsuranceOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsuranceOracle *InsuranceOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsuranceOracle.Contract.InsuranceOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InsuranceOracle *InsuranceOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InsuranceOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InsuranceOracle *InsuranceOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InsuranceOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InsuranceOracle *InsuranceOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InsuranceOracle.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() view returns(address)
func (_InsuranceOracle *InsuranceOracleCaller) GetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InsuranceOracle.contract.Call(opts, &out, "getAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() view returns(address)
func (_InsuranceOracle *InsuranceOracleSession) GetAddress() (common.Address, error) {
	return _InsuranceOracle.Contract.GetAddress(&_InsuranceOracle.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x38cc4831.
//
// Solidity: function getAddress() view returns(address)
func (_InsuranceOracle *InsuranceOracleCallerSession) GetAddress() (common.Address, error) {
	return _InsuranceOracle.Contract.GetAddress(&_InsuranceOracle.CallOpts)
}

// GetPolicyIsClaimable is a free data retrieval call binding the contract method 0xf714c06c.
//
// Solidity: function getPolicyIsClaimable() view returns(bool)
func (_InsuranceOracle *InsuranceOracleCaller) GetPolicyIsClaimable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InsuranceOracle.contract.Call(opts, &out, "getPolicyIsClaimable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPolicyIsClaimable is a free data retrieval call binding the contract method 0xf714c06c.
//
// Solidity: function getPolicyIsClaimable() view returns(bool)
func (_InsuranceOracle *InsuranceOracleSession) GetPolicyIsClaimable() (bool, error) {
	return _InsuranceOracle.Contract.GetPolicyIsClaimable(&_InsuranceOracle.CallOpts)
}

// GetPolicyIsClaimable is a free data retrieval call binding the contract method 0xf714c06c.
//
// Solidity: function getPolicyIsClaimable() view returns(bool)
func (_InsuranceOracle *InsuranceOracleCallerSession) GetPolicyIsClaimable() (bool, error) {
	return _InsuranceOracle.Contract.GetPolicyIsClaimable(&_InsuranceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InsuranceOracle *InsuranceOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InsuranceOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InsuranceOracle *InsuranceOracleSession) Owner() (common.Address, error) {
	return _InsuranceOracle.Contract.Owner(&_InsuranceOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InsuranceOracle *InsuranceOracleCallerSession) Owner() (common.Address, error) {
	return _InsuranceOracle.Contract.Owner(&_InsuranceOracle.CallOpts)
}

// TestConnection is a free data retrieval call binding the contract method 0x3412a15c.
//
// Solidity: function testConnection() pure returns(bool)
func (_InsuranceOracle *InsuranceOracleCaller) TestConnection(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InsuranceOracle.contract.Call(opts, &out, "testConnection")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TestConnection is a free data retrieval call binding the contract method 0x3412a15c.
//
// Solidity: function testConnection() pure returns(bool)
func (_InsuranceOracle *InsuranceOracleSession) TestConnection() (bool, error) {
	return _InsuranceOracle.Contract.TestConnection(&_InsuranceOracle.CallOpts)
}

// TestConnection is a free data retrieval call binding the contract method 0x3412a15c.
//
// Solidity: function testConnection() pure returns(bool)
func (_InsuranceOracle *InsuranceOracleCallerSession) TestConnection() (bool, error) {
	return _InsuranceOracle.Contract.TestConnection(&_InsuranceOracle.CallOpts)
}

// DeclarePolicyActive is a paid mutator transaction binding the contract method 0x9d8b3e24.
//
// Solidity: function declarePolicyActive(bool value) returns()
func (_InsuranceOracle *InsuranceOracleTransactor) DeclarePolicyActive(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _InsuranceOracle.contract.Transact(opts, "declarePolicyActive", value)
}

// DeclarePolicyActive is a paid mutator transaction binding the contract method 0x9d8b3e24.
//
// Solidity: function declarePolicyActive(bool value) returns()
func (_InsuranceOracle *InsuranceOracleSession) DeclarePolicyActive(value bool) (*types.Transaction, error) {
	return _InsuranceOracle.Contract.DeclarePolicyActive(&_InsuranceOracle.TransactOpts, value)
}

// DeclarePolicyActive is a paid mutator transaction binding the contract method 0x9d8b3e24.
//
// Solidity: function declarePolicyActive(bool value) returns()
func (_InsuranceOracle *InsuranceOracleTransactorSession) DeclarePolicyActive(value bool) (*types.Transaction, error) {
	return _InsuranceOracle.Contract.DeclarePolicyActive(&_InsuranceOracle.TransactOpts, value)
}

// SetPolicyIsClaimable is a paid mutator transaction binding the contract method 0x2dd308d9.
//
// Solidity: function setPolicyIsClaimable(bool value) returns()
func (_InsuranceOracle *InsuranceOracleTransactor) SetPolicyIsClaimable(opts *bind.TransactOpts, value bool) (*types.Transaction, error) {
	return _InsuranceOracle.contract.Transact(opts, "setPolicyIsClaimable", value)
}

// SetPolicyIsClaimable is a paid mutator transaction binding the contract method 0x2dd308d9.
//
// Solidity: function setPolicyIsClaimable(bool value) returns()
func (_InsuranceOracle *InsuranceOracleSession) SetPolicyIsClaimable(value bool) (*types.Transaction, error) {
	return _InsuranceOracle.Contract.SetPolicyIsClaimable(&_InsuranceOracle.TransactOpts, value)
}

// SetPolicyIsClaimable is a paid mutator transaction binding the contract method 0x2dd308d9.
//
// Solidity: function setPolicyIsClaimable(bool value) returns()
func (_InsuranceOracle *InsuranceOracleTransactorSession) SetPolicyIsClaimable(value bool) (*types.Transaction, error) {
	return _InsuranceOracle.Contract.SetPolicyIsClaimable(&_InsuranceOracle.TransactOpts, value)
}

// InsuranceOraclePolicyInitializedIterator is returned from FilterPolicyInitialized and is used to iterate over the raw logs and unpacked data for PolicyInitialized events raised by the InsuranceOracle contract.
type InsuranceOraclePolicyInitializedIterator struct {
	Event *InsuranceOraclePolicyInitialized // Event containing the contract specifics and raw log

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
func (it *InsuranceOraclePolicyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InsuranceOraclePolicyInitialized)
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
		it.Event = new(InsuranceOraclePolicyInitialized)
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
func (it *InsuranceOraclePolicyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InsuranceOraclePolicyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InsuranceOraclePolicyInitialized represents a PolicyInitialized event raised by the InsuranceOracle contract.
type InsuranceOraclePolicyInitialized struct {
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPolicyInitialized is a free log retrieval operation binding the contract event 0x7f80f6442c3355ae66fd57fc95715c623b267bff31b0f17d8de16c010e0510e9.
//
// Solidity: event PolicyInitialized(bool isActive)
func (_InsuranceOracle *InsuranceOracleFilterer) FilterPolicyInitialized(opts *bind.FilterOpts) (*InsuranceOraclePolicyInitializedIterator, error) {

	logs, sub, err := _InsuranceOracle.contract.FilterLogs(opts, "PolicyInitialized")
	if err != nil {
		return nil, err
	}
	return &InsuranceOraclePolicyInitializedIterator{contract: _InsuranceOracle.contract, event: "PolicyInitialized", logs: logs, sub: sub}, nil
}

// WatchPolicyInitialized is a free log subscription operation binding the contract event 0x7f80f6442c3355ae66fd57fc95715c623b267bff31b0f17d8de16c010e0510e9.
//
// Solidity: event PolicyInitialized(bool isActive)
func (_InsuranceOracle *InsuranceOracleFilterer) WatchPolicyInitialized(opts *bind.WatchOpts, sink chan<- *InsuranceOraclePolicyInitialized) (event.Subscription, error) {

	logs, sub, err := _InsuranceOracle.contract.WatchLogs(opts, "PolicyInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InsuranceOraclePolicyInitialized)
				if err := _InsuranceOracle.contract.UnpackLog(event, "PolicyInitialized", log); err != nil {
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

// ParsePolicyInitialized is a log parse operation binding the contract event 0x7f80f6442c3355ae66fd57fc95715c623b267bff31b0f17d8de16c010e0510e9.
//
// Solidity: event PolicyInitialized(bool isActive)
func (_InsuranceOracle *InsuranceOracleFilterer) ParsePolicyInitialized(log types.Log) (*InsuranceOraclePolicyInitialized, error) {
	event := new(InsuranceOraclePolicyInitialized)
	if err := _InsuranceOracle.contract.UnpackLog(event, "PolicyInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
