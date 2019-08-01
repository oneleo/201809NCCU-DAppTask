// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hub

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

// AddressUtilsABI is the input ABI used to generate the binding from.
const AddressUtilsABI = "[]"

// AddressUtilsBin is the compiled bytecode used for deploying new contracts.
const AddressUtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058206787267780ef8f389fd0bd4ef2e5f378b4a68443b8cef6d93cd20cbd097b5aa80029`

// DeployAddressUtils deploys a new Ethereum contract, binding an instance of AddressUtils to it.
func DeployAddressUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUtils{AddressUtilsCaller: AddressUtilsCaller{contract: contract}, AddressUtilsTransactor: AddressUtilsTransactor{contract: contract}, AddressUtilsFilterer: AddressUtilsFilterer{contract: contract}}, nil
}

// AddressUtils is an auto generated Go binding around an Ethereum contract.
type AddressUtils struct {
	AddressUtilsCaller     // Read-only binding to the contract
	AddressUtilsTransactor // Write-only binding to the contract
	AddressUtilsFilterer   // Log filterer for contract events
}

// AddressUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUtilsSession struct {
	Contract     *AddressUtils     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUtilsCallerSession struct {
	Contract *AddressUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AddressUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUtilsTransactorSession struct {
	Contract     *AddressUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AddressUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUtilsRaw struct {
	Contract *AddressUtils // Generic contract binding to access the raw methods on
}

// AddressUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUtilsCallerRaw struct {
	Contract *AddressUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUtilsTransactorRaw struct {
	Contract *AddressUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUtils creates a new instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtils(address common.Address, backend bind.ContractBackend) (*AddressUtils, error) {
	contract, err := bindAddressUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUtils{AddressUtilsCaller: AddressUtilsCaller{contract: contract}, AddressUtilsTransactor: AddressUtilsTransactor{contract: contract}, AddressUtilsFilterer: AddressUtilsFilterer{contract: contract}}, nil
}

// NewAddressUtilsCaller creates a new read-only instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsCaller(address common.Address, caller bind.ContractCaller) (*AddressUtilsCaller, error) {
	contract, err := bindAddressUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsCaller{contract: contract}, nil
}

// NewAddressUtilsTransactor creates a new write-only instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUtilsTransactor, error) {
	contract, err := bindAddressUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsTransactor{contract: contract}, nil
}

// NewAddressUtilsFilterer creates a new log filterer instance of AddressUtils, bound to a specific deployed contract.
func NewAddressUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUtilsFilterer, error) {
	contract, err := bindAddressUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUtilsFilterer{contract: contract}, nil
}

// bindAddressUtils binds a generic wrapper to an already deployed contract.
func bindAddressUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUtils *AddressUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUtils.Contract.AddressUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUtils *AddressUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUtils.Contract.AddressUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUtils *AddressUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUtils.Contract.AddressUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUtils *AddressUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUtils *AddressUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUtils *AddressUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUtils.Contract.contract.Transact(opts, method, params...)
}

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC165Bin is the compiled bytecode used for deploying new contracts.
const ERC165Bin = `0x`

// DeployERC165 deploys a new Ethereum contract, binding an instance of ERC165 to it.
func DeployERC165(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC165, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC165Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC165.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, _interfaceId)
}

// ERC721ABI is the input ABI used to generate the binding from.
const ERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721Bin is the compiled bytecode used for deploying new contracts.
const ERC721Bin = `0x`

// DeployERC721 deploys a new Ethereum contract, binding an instance of ERC721 to it.
func DeployERC721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721 *ERC721Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721 *ERC721CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721 *ERC721Caller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721 *ERC721Session) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721 *ERC721CallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.GetApproved(&_ERC721.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721 *ERC721Caller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721 *ERC721Session) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721 *ERC721CallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721.Contract.IsApprovedForAll(&_ERC721.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721 *ERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721 *ERC721Session) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721 *ERC721CallerSession) Name() (string, error) {
	return _ERC721.Contract.Name(&_ERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721 *ERC721Session) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721 *ERC721CallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721 *ERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721 *ERC721Session) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721 *ERC721CallerSession) Symbol() (string, error) {
	return _ERC721.Contract.Symbol(&_ERC721.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721 *ERC721Caller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721 *ERC721Session) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721 *ERC721CallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenByIndex(&_ERC721.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721 *ERC721Caller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721 *ERC721Session) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721 *ERC721CallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721.Contract.TokenOfOwnerByIndex(&_ERC721.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721 *ERC721Caller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721 *ERC721Session) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721 *ERC721CallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721.Contract.TokenURI(&_ERC721.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721 *ERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721 *ERC721Session) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721 *ERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Session) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721 *ERC721TransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721 *ERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721 *ERC721Session) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721 *ERC721TransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721.Contract.SafeTransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721 *ERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721 *ERC721Session) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721 *ERC721TransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721.Contract.SetApprovalForAll(&_ERC721.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721 *ERC721Session) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721ApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721 contract.
type ERC721ApprovalForAllIterator struct {
	Event *ERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ApprovalForAll)
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
		it.Event = new(ERC721ApprovalForAll)
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
func (it *ERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ApprovalForAll represents a ApprovalForAll event raised by the ERC721 contract.
type ERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721 *ERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721ApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalForAllIterator{contract: _ERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721 *ERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721ApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ApprovalForAll)
				if err := _ERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721TransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721BasicABI is the input ABI used to generate the binding from.
const ERC721BasicABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721BasicBin is the compiled bytecode used for deploying new contracts.
const ERC721BasicBin = `0x`

// DeployERC721Basic deploys a new Ethereum contract, binding an instance of ERC721Basic to it.
func DeployERC721Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Basic{ERC721BasicCaller: ERC721BasicCaller{contract: contract}, ERC721BasicTransactor: ERC721BasicTransactor{contract: contract}, ERC721BasicFilterer: ERC721BasicFilterer{contract: contract}}, nil
}

// ERC721Basic is an auto generated Go binding around an Ethereum contract.
type ERC721Basic struct {
	ERC721BasicCaller     // Read-only binding to the contract
	ERC721BasicTransactor // Write-only binding to the contract
	ERC721BasicFilterer   // Log filterer for contract events
}

// ERC721BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721BasicSession struct {
	Contract     *ERC721Basic      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721BasicCallerSession struct {
	Contract *ERC721BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC721BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721BasicTransactorSession struct {
	Contract     *ERC721BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC721BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721BasicRaw struct {
	Contract *ERC721Basic // Generic contract binding to access the raw methods on
}

// ERC721BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721BasicCallerRaw struct {
	Contract *ERC721BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721BasicTransactorRaw struct {
	Contract *ERC721BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Basic creates a new instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721Basic(address common.Address, backend bind.ContractBackend) (*ERC721Basic, error) {
	contract, err := bindERC721Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Basic{ERC721BasicCaller: ERC721BasicCaller{contract: contract}, ERC721BasicTransactor: ERC721BasicTransactor{contract: contract}, ERC721BasicFilterer: ERC721BasicFilterer{contract: contract}}, nil
}

// NewERC721BasicCaller creates a new read-only instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC721BasicCaller, error) {
	contract, err := bindERC721Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicCaller{contract: contract}, nil
}

// NewERC721BasicTransactor creates a new write-only instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721BasicTransactor, error) {
	contract, err := bindERC721Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTransactor{contract: contract}, nil
}

// NewERC721BasicFilterer creates a new log filterer instance of ERC721Basic, bound to a specific deployed contract.
func NewERC721BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721BasicFilterer, error) {
	contract, err := bindERC721Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicFilterer{contract: contract}, nil
}

// bindERC721Basic binds a generic wrapper to an already deployed contract.
func bindERC721Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Basic *ERC721BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Basic.Contract.ERC721BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Basic *ERC721BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Basic.Contract.ERC721BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Basic *ERC721BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Basic.Contract.ERC721BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Basic *ERC721BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Basic *ERC721BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Basic *ERC721BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Basic *ERC721BasicCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Basic *ERC721BasicSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Basic.Contract.BalanceOf(&_ERC721Basic.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Basic *ERC721BasicCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Basic.Contract.BalanceOf(&_ERC721Basic.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Basic *ERC721BasicCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Basic *ERC721BasicSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.GetApproved(&_ERC721Basic.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Basic *ERC721BasicCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.GetApproved(&_ERC721Basic.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Basic *ERC721BasicCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Basic *ERC721BasicSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Basic.Contract.IsApprovedForAll(&_ERC721Basic.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Basic *ERC721BasicCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Basic.Contract.IsApprovedForAll(&_ERC721Basic.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Basic *ERC721BasicCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Basic *ERC721BasicSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.OwnerOf(&_ERC721Basic.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Basic *ERC721BasicCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Basic.Contract.OwnerOf(&_ERC721Basic.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Basic *ERC721BasicCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Basic.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Basic *ERC721BasicSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Basic.Contract.SupportsInterface(&_ERC721Basic.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Basic *ERC721BasicCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Basic.Contract.SupportsInterface(&_ERC721Basic.CallOpts, _interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.Approve(&_ERC721Basic.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.Approve(&_ERC721Basic.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Basic *ERC721BasicTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Basic *ERC721BasicSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SafeTransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SafeTransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Basic *ERC721BasicTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Basic *ERC721BasicSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SetApprovalForAll(&_ERC721Basic.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Basic.Contract.SetApprovalForAll(&_ERC721Basic.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.TransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Basic *ERC721BasicTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Basic.Contract.TransferFrom(&_ERC721Basic.TransactOpts, _from, _to, _tokenId)
}

// ERC721BasicApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Basic contract.
type ERC721BasicApprovalIterator struct {
	Event *ERC721BasicApproval // Event containing the contract specifics and raw log

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
func (it *ERC721BasicApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicApproval)
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
		it.Event = new(ERC721BasicApproval)
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
func (it *ERC721BasicApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicApproval represents a Approval event raised by the ERC721Basic contract.
type ERC721BasicApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721Basic *ERC721BasicFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721BasicApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Basic.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicApprovalIterator{contract: _ERC721Basic.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721Basic *ERC721BasicFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721BasicApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Basic.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicApproval)
				if err := _ERC721Basic.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721BasicApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Basic contract.
type ERC721BasicApprovalForAllIterator struct {
	Event *ERC721BasicApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721BasicApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicApprovalForAll)
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
		it.Event = new(ERC721BasicApprovalForAll)
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
func (it *ERC721BasicApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicApprovalForAll represents a ApprovalForAll event raised by the ERC721Basic contract.
type ERC721BasicApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Basic *ERC721BasicFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721BasicApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Basic.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicApprovalForAllIterator{contract: _ERC721Basic.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Basic *ERC721BasicFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721BasicApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Basic.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicApprovalForAll)
				if err := _ERC721Basic.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Basic contract.
type ERC721BasicTransferIterator struct {
	Event *ERC721BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTransfer)
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
		it.Event = new(ERC721BasicTransfer)
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
func (it *ERC721BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTransfer represents a Transfer event raised by the ERC721Basic contract.
type ERC721BasicTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721Basic *ERC721BasicFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721BasicTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Basic.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTransferIterator{contract: _ERC721Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721Basic *ERC721BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721BasicTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Basic.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTransfer)
				if err := _ERC721Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721BasicTokenABI is the input ABI used to generate the binding from.
const ERC721BasicTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721BasicTokenBin is the compiled bytecode used for deploying new contracts.
const ERC721BasicTokenBin = `0x608060405234801561001057600080fd5b506100437f01ffc9a70000000000000000000000000000000000000000000000000000000064010000000061007a810204565b6100757f80ac58cd0000000000000000000000000000000000000000000000000000000064010000000061007a810204565b6100e6565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156100a957600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b610991806100f56000396000f3006080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146100b3578063081812fc146100e9578063095ea7b31461011d57806319fa8f501461014357806323b872dd1461017557806342842e0e1461019f5780636352211e146101c957806370a08231146101e1578063a22cb46514610214578063b88d4fde1461023a578063e985e9c5146102a9575b600080fd5b3480156100bf57600080fd5b506100d5600160e060020a0319600435166102d0565b604080519115158252519081900360200190f35b3480156100f557600080fd5b506101016004356102ef565b60408051600160a060020a039092168252519081900360200190f35b34801561012957600080fd5b50610141600160a060020a036004351660243561030a565b005b34801561014f57600080fd5b506101586103c0565b60408051600160e060020a03199092168252519081900360200190f35b34801561018157600080fd5b50610141600160a060020a03600435811690602435166044356103e4565b3480156101ab57600080fd5b50610141600160a060020a0360043581169060243516604435610472565b3480156101d557600080fd5b50610101600435610493565b3480156101ed57600080fd5b50610202600160a060020a03600435166104bd565b60408051918252519081900360200190f35b34801561022057600080fd5b50610141600160a060020a036004351660243515156104f0565b34801561024657600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261014194600160a060020a0381358116956024803590921695604435953695608494019181908401838280828437509497506105749650505050505050565b3480156102b557600080fd5b506100d5600160a060020a036004358116906024351661059c565b600160e060020a03191660009081526020819052604090205460ff1690565b600090815260026020526040902054600160a060020a031690565b600061031582610493565b9050600160a060020a03838116908216141561033057600080fd5b33600160a060020a038216148061034c575061034c813361059c565b151561035757600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081565b6103ee33826105ca565b15156103f957600080fd5b600160a060020a038216151561040e57600080fd5b6104188382610629565b610422838261069a565b61042c8282610730565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b61048e8383836020604051908101604052806000815250610574565b505050565b600081815260016020526040812054600160a060020a03168015156104b757600080fd5b92915050565b6000600160a060020a03821615156104d457600080fd5b50600160a060020a031660009081526003602052604090205490565b600160a060020a03821633141561050657600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b61057f8484846103e4565b61058b848484846107c0565b151561059657600080fd5b50505050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b6000806105d683610493565b905080600160a060020a031684600160a060020a03161480610611575083600160a060020a0316610606846102ef565b600160a060020a0316145b806106215750610621818561059c565b949350505050565b81600160a060020a031661063c82610493565b600160a060020a03161461064f57600080fd5b600081815260026020526040902054600160a060020a031615610696576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b81600160a060020a03166106ad82610493565b600160a060020a0316146106c057600080fd5b600160a060020a0382166000908152600360205260409020546106ea90600163ffffffff61092d16565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600081815260016020526040902054600160a060020a03161561075257600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03881690811790915584526003909152909120546107a091610944565b600160a060020a0390921660009081526003602052604090209190915550565b6000806107d585600160a060020a031661095d565b15156107e45760019150610924565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b8381101561087757818101518382015260200161085f565b50505050905090810190601f1680156108a45780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b1580156108c657600080fd5b505af11580156108da573d6000803e3d6000fd5b505050506040513d60208110156108f057600080fd5b5051600160e060020a031981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b6000808383111561093d57600080fd5b5050900390565b60008282018381101561095657600080fd5b9392505050565b6000903b11905600a165627a7a72305820a63697fffb7202f495af60c881500053e568c479c80681022391c4b3133530c00029`

// DeployERC721BasicToken deploys a new Ethereum contract, binding an instance of ERC721BasicToken to it.
func DeployERC721BasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721BasicToken{ERC721BasicTokenCaller: ERC721BasicTokenCaller{contract: contract}, ERC721BasicTokenTransactor: ERC721BasicTokenTransactor{contract: contract}, ERC721BasicTokenFilterer: ERC721BasicTokenFilterer{contract: contract}}, nil
}

// ERC721BasicToken is an auto generated Go binding around an Ethereum contract.
type ERC721BasicToken struct {
	ERC721BasicTokenCaller     // Read-only binding to the contract
	ERC721BasicTokenTransactor // Write-only binding to the contract
	ERC721BasicTokenFilterer   // Log filterer for contract events
}

// ERC721BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721BasicTokenSession struct {
	Contract     *ERC721BasicToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721BasicTokenCallerSession struct {
	Contract *ERC721BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721BasicTokenTransactorSession struct {
	Contract     *ERC721BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721BasicTokenRaw struct {
	Contract *ERC721BasicToken // Generic contract binding to access the raw methods on
}

// ERC721BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721BasicTokenCallerRaw struct {
	Contract *ERC721BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721BasicTokenTransactorRaw struct {
	Contract *ERC721BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721BasicToken creates a new instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicToken(address common.Address, backend bind.ContractBackend) (*ERC721BasicToken, error) {
	contract, err := bindERC721BasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicToken{ERC721BasicTokenCaller: ERC721BasicTokenCaller{contract: contract}, ERC721BasicTokenTransactor: ERC721BasicTokenTransactor{contract: contract}, ERC721BasicTokenFilterer: ERC721BasicTokenFilterer{contract: contract}}, nil
}

// NewERC721BasicTokenCaller creates a new read-only instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicTokenCaller(address common.Address, caller bind.ContractCaller) (*ERC721BasicTokenCaller, error) {
	contract, err := bindERC721BasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenCaller{contract: contract}, nil
}

// NewERC721BasicTokenTransactor creates a new write-only instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721BasicTokenTransactor, error) {
	contract, err := bindERC721BasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenTransactor{contract: contract}, nil
}

// NewERC721BasicTokenFilterer creates a new log filterer instance of ERC721BasicToken, bound to a specific deployed contract.
func NewERC721BasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721BasicTokenFilterer, error) {
	contract, err := bindERC721BasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenFilterer{contract: contract}, nil
}

// bindERC721BasicToken binds a generic wrapper to an already deployed contract.
func bindERC721BasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721BasicToken *ERC721BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721BasicToken.Contract.ERC721BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721BasicToken *ERC721BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.ERC721BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721BasicToken *ERC721BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.ERC721BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721BasicToken *ERC721BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721BasicToken *ERC721BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721BasicToken *ERC721BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721BasicToken *ERC721BasicTokenCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721BasicToken *ERC721BasicTokenSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721BasicToken.Contract.InterfaceIdERC165(&_ERC721BasicToken.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721BasicToken.Contract.InterfaceIdERC165(&_ERC721BasicToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721BasicToken *ERC721BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721BasicToken *ERC721BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721BasicToken.Contract.BalanceOf(&_ERC721BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721BasicToken.Contract.BalanceOf(&_ERC721BasicToken.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.GetApproved(&_ERC721BasicToken.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.GetApproved(&_ERC721BasicToken.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721BasicToken.Contract.IsApprovedForAll(&_ERC721BasicToken.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721BasicToken.Contract.IsApprovedForAll(&_ERC721BasicToken.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.OwnerOf(&_ERC721BasicToken.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721BasicToken.Contract.OwnerOf(&_ERC721BasicToken.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721BasicToken.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721BasicToken.Contract.SupportsInterface(&_ERC721BasicToken.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721BasicToken *ERC721BasicTokenCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721BasicToken.Contract.SupportsInterface(&_ERC721BasicToken.CallOpts, _interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.Approve(&_ERC721BasicToken.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.Approve(&_ERC721BasicToken.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SafeTransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SafeTransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SetApprovalForAll(&_ERC721BasicToken.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.SetApprovalForAll(&_ERC721BasicToken.TransactOpts, _to, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.TransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721BasicToken *ERC721BasicTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721BasicToken.Contract.TransferFrom(&_ERC721BasicToken.TransactOpts, _from, _to, _tokenId)
}

// ERC721BasicTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721BasicToken contract.
type ERC721BasicTokenApprovalIterator struct {
	Event *ERC721BasicTokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTokenApproval)
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
		it.Event = new(ERC721BasicTokenApproval)
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
func (it *ERC721BasicTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTokenApproval represents a Approval event raised by the ERC721BasicToken contract.
type ERC721BasicTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721BasicTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenApprovalIterator{contract: _ERC721BasicToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721BasicTokenApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTokenApproval)
				if err := _ERC721BasicToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721BasicTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721BasicToken contract.
type ERC721BasicTokenApprovalForAllIterator struct {
	Event *ERC721BasicTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTokenApprovalForAll)
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
		it.Event = new(ERC721BasicTokenApprovalForAll)
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
func (it *ERC721BasicTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTokenApprovalForAll represents a ApprovalForAll event raised by the ERC721BasicToken contract.
type ERC721BasicTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721BasicTokenApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenApprovalForAllIterator{contract: _ERC721BasicToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721BasicTokenApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTokenApprovalForAll)
				if err := _ERC721BasicToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721BasicToken contract.
type ERC721BasicTokenTransferIterator struct {
	Event *ERC721BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721BasicTokenTransfer)
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
		it.Event = new(ERC721BasicTokenTransfer)
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
func (it *ERC721BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721BasicTokenTransfer represents a Transfer event raised by the ERC721BasicToken contract.
type ERC721BasicTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721BasicTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721BasicTokenTransferIterator{contract: _ERC721BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721BasicToken *ERC721BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721BasicTokenTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721BasicToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721BasicTokenTransfer)
				if err := _ERC721BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721EnumerableABI is the input ABI used to generate the binding from.
const ERC721EnumerableABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721EnumerableBin is the compiled bytecode used for deploying new contracts.
const ERC721EnumerableBin = `0x`

// DeployERC721Enumerable deploys a new Ethereum contract, binding an instance of ERC721Enumerable to it.
func DeployERC721Enumerable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Enumerable, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721EnumerableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

// ERC721Enumerable is an auto generated Go binding around an Ethereum contract.
type ERC721Enumerable struct {
	ERC721EnumerableCaller     // Read-only binding to the contract
	ERC721EnumerableTransactor // Write-only binding to the contract
	ERC721EnumerableFilterer   // Log filterer for contract events
}

// ERC721EnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721EnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721EnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721EnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721EnumerableSession struct {
	Contract     *ERC721Enumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721EnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721EnumerableCallerSession struct {
	Contract *ERC721EnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC721EnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721EnumerableTransactorSession struct {
	Contract     *ERC721EnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC721EnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721EnumerableRaw struct {
	Contract *ERC721Enumerable // Generic contract binding to access the raw methods on
}

// ERC721EnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721EnumerableCallerRaw struct {
	Contract *ERC721EnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721EnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721EnumerableTransactorRaw struct {
	Contract *ERC721EnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Enumerable creates a new instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721Enumerable(address common.Address, backend bind.ContractBackend) (*ERC721Enumerable, error) {
	contract, err := bindERC721Enumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Enumerable{ERC721EnumerableCaller: ERC721EnumerableCaller{contract: contract}, ERC721EnumerableTransactor: ERC721EnumerableTransactor{contract: contract}, ERC721EnumerableFilterer: ERC721EnumerableFilterer{contract: contract}}, nil
}

// NewERC721EnumerableCaller creates a new read-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableCaller(address common.Address, caller bind.ContractCaller) (*ERC721EnumerableCaller, error) {
	contract, err := bindERC721Enumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableCaller{contract: contract}, nil
}

// NewERC721EnumerableTransactor creates a new write-only instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721EnumerableTransactor, error) {
	contract, err := bindERC721Enumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransactor{contract: contract}, nil
}

// NewERC721EnumerableFilterer creates a new log filterer instance of ERC721Enumerable, bound to a specific deployed contract.
func NewERC721EnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721EnumerableFilterer, error) {
	contract, err := bindERC721Enumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableFilterer{contract: contract}, nil
}

// bindERC721Enumerable binds a generic wrapper to an already deployed contract.
func bindERC721Enumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721EnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.ERC721EnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.ERC721EnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Enumerable *ERC721EnumerableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Enumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Enumerable *ERC721EnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Enumerable.Contract.BalanceOf(&_ERC721Enumerable.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Enumerable *ERC721EnumerableCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Enumerable *ERC721EnumerableSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.GetApproved(&_ERC721Enumerable.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Enumerable.Contract.IsApprovedForAll(&_ERC721Enumerable.CallOpts, _owner, _operator)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Enumerable *ERC721EnumerableCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Enumerable *ERC721EnumerableSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Enumerable.Contract.OwnerOf(&_ERC721Enumerable.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Enumerable.Contract.SupportsInterface(&_ERC721Enumerable.CallOpts, _interfaceId)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenByIndex(&_ERC721Enumerable.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Enumerable.Contract.TokenOfOwnerByIndex(&_ERC721Enumerable.CallOpts, _owner, _index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Enumerable.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Enumerable *ERC721EnumerableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Enumerable.Contract.TotalSupply(&_ERC721Enumerable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.Approve(&_ERC721Enumerable.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SafeTransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.SetApprovalForAll(&_ERC721Enumerable.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Enumerable *ERC721EnumerableTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Enumerable.Contract.TransferFrom(&_ERC721Enumerable.TransactOpts, _from, _to, _tokenId)
}

// ERC721EnumerableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalIterator struct {
	Event *ERC721EnumerableApproval // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApproval)
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
		it.Event = new(ERC721EnumerableApproval)
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
func (it *ERC721EnumerableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApproval represents a Approval event raised by the ERC721Enumerable contract.
type ERC721EnumerableApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721EnumerableApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalIterator{contract: _ERC721Enumerable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApproval)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721EnumerableApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAllIterator struct {
	Event *ERC721EnumerableApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableApprovalForAll)
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
		it.Event = new(ERC721EnumerableApprovalForAll)
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
func (it *ERC721EnumerableApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableApprovalForAll represents a ApprovalForAll event raised by the ERC721Enumerable contract.
type ERC721EnumerableApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721EnumerableApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableApprovalForAllIterator{contract: _ERC721Enumerable.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableApprovalForAll)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721EnumerableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Enumerable contract.
type ERC721EnumerableTransferIterator struct {
	Event *ERC721EnumerableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721EnumerableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721EnumerableTransfer)
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
		it.Event = new(ERC721EnumerableTransfer)
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
func (it *ERC721EnumerableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721EnumerableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721EnumerableTransfer represents a Transfer event raised by the ERC721Enumerable contract.
type ERC721EnumerableTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721Enumerable *ERC721EnumerableFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721EnumerableTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721EnumerableTransferIterator{contract: _ERC721Enumerable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721Enumerable *ERC721EnumerableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721EnumerableTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Enumerable.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721EnumerableTransfer)
				if err := _ERC721Enumerable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721MetadataABI is the input ABI used to generate the binding from.
const ERC721MetadataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_operator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721MetadataBin is the compiled bytecode used for deploying new contracts.
const ERC721MetadataBin = `0x`

// DeployERC721Metadata deploys a new Ethereum contract, binding an instance of ERC721Metadata to it.
func DeployERC721Metadata(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Metadata, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721MetadataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

// ERC721Metadata is an auto generated Go binding around an Ethereum contract.
type ERC721Metadata struct {
	ERC721MetadataCaller     // Read-only binding to the contract
	ERC721MetadataTransactor // Write-only binding to the contract
	ERC721MetadataFilterer   // Log filterer for contract events
}

// ERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721MetadataSession struct {
	Contract     *ERC721Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721MetadataCallerSession struct {
	Contract *ERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721MetadataTransactorSession struct {
	Contract     *ERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721MetadataRaw struct {
	Contract *ERC721Metadata // Generic contract binding to access the raw methods on
}

// ERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721MetadataCallerRaw struct {
	Contract *ERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721MetadataTransactorRaw struct {
	Contract *ERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Metadata creates a new instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721Metadata(address common.Address, backend bind.ContractBackend) (*ERC721Metadata, error) {
	contract, err := bindERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

// NewERC721MetadataCaller creates a new read-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*ERC721MetadataCaller, error) {
	contract, err := bindERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataCaller{contract: contract}, nil
}

// NewERC721MetadataTransactor creates a new write-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MetadataTransactor, error) {
	contract, err := bindERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataTransactor{contract: contract}, nil
}

// NewERC721MetadataFilterer creates a new log filterer instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MetadataFilterer, error) {
	contract, err := bindERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataFilterer{contract: contract}, nil
}

// bindERC721Metadata binds a generic wrapper to an already deployed contract.
func bindERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.ERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Metadata *ERC721MetadataCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Metadata *ERC721MetadataSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(_balance uint256)
func (_ERC721Metadata *ERC721MetadataCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Metadata.Contract.BalanceOf(&_ERC721Metadata.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Metadata *ERC721MetadataCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Metadata *ERC721MetadataSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(_operator address)
func (_ERC721Metadata *ERC721MetadataCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.GetApproved(&_ERC721Metadata.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Metadata.Contract.IsApprovedForAll(&_ERC721Metadata.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721Metadata *ERC721MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721Metadata *ERC721MetadataSession) Name() (string, error) {
	return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(_name string)
func (_ERC721Metadata *ERC721MetadataCallerSession) Name() (string, error) {
	return _ERC721Metadata.Contract.Name(&_ERC721Metadata.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Metadata *ERC721MetadataCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Metadata *ERC721MetadataSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(_owner address)
func (_ERC721Metadata *ERC721MetadataCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Metadata.Contract.OwnerOf(&_ERC721Metadata.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Metadata.Contract.SupportsInterface(&_ERC721Metadata.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Metadata *ERC721MetadataCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Metadata.Contract.SupportsInterface(&_ERC721Metadata.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721Metadata *ERC721MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721Metadata *ERC721MetadataSession) Symbol() (string, error) {
	return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(_symbol string)
func (_ERC721Metadata *ERC721MetadataCallerSession) Symbol() (string, error) {
	return _ERC721Metadata.Contract.Symbol(&_ERC721Metadata.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721Metadata *ERC721MetadataCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Metadata.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721Metadata *ERC721MetadataSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721Metadata *ERC721MetadataCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Metadata.Contract.TokenURI(&_ERC721Metadata.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.Approve(&_ERC721Metadata.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Metadata *ERC721MetadataSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SafeTransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Metadata *ERC721MetadataSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.SetApprovalForAll(&_ERC721Metadata.TransactOpts, _operator, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Metadata *ERC721MetadataTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.TransferFrom(&_ERC721Metadata.TransactOpts, _from, _to, _tokenId)
}

// ERC721MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Metadata contract.
type ERC721MetadataApprovalIterator struct {
	Event *ERC721MetadataApproval // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataApproval)
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
		it.Event = new(ERC721MetadataApproval)
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
func (it *ERC721MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataApproval represents a Approval event raised by the ERC721Metadata contract.
type ERC721MetadataApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721MetadataApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataApprovalIterator{contract: _ERC721Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataApproval)
				if err := _ERC721Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721MetadataApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Metadata contract.
type ERC721MetadataApprovalForAllIterator struct {
	Event *ERC721MetadataApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataApprovalForAll)
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
		it.Event = new(ERC721MetadataApprovalForAll)
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
func (it *ERC721MetadataApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataApprovalForAll represents a ApprovalForAll event raised by the ERC721Metadata contract.
type ERC721MetadataApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721MetadataApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataApprovalForAllIterator{contract: _ERC721Metadata.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721MetadataApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataApprovalForAll)
				if err := _ERC721Metadata.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Metadata contract.
type ERC721MetadataTransferIterator struct {
	Event *ERC721MetadataTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MetadataTransfer)
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
		it.Event = new(ERC721MetadataTransfer)
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
func (it *ERC721MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MetadataTransfer represents a Transfer event raised by the ERC721Metadata contract.
type ERC721MetadataTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721Metadata *ERC721MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721MetadataTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataTransferIterator{contract: _ERC721Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721Metadata *ERC721MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721MetadataTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Metadata.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MetadataTransfer)
				if err := _ERC721Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721ReceiverABI is the input ABI used to generate the binding from.
const ERC721ReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721ReceiverBin is the compiled bytecode used for deploying new contracts.
const ERC721ReceiverBin = `0x`

// DeployERC721Receiver deploys a new Ethereum contract, binding an instance of ERC721Receiver to it.
func DeployERC721Receiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Receiver, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ReceiverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721ReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Receiver{ERC721ReceiverCaller: ERC721ReceiverCaller{contract: contract}, ERC721ReceiverTransactor: ERC721ReceiverTransactor{contract: contract}, ERC721ReceiverFilterer: ERC721ReceiverFilterer{contract: contract}}, nil
}

// ERC721Receiver is an auto generated Go binding around an Ethereum contract.
type ERC721Receiver struct {
	ERC721ReceiverCaller     // Read-only binding to the contract
	ERC721ReceiverTransactor // Write-only binding to the contract
	ERC721ReceiverFilterer   // Log filterer for contract events
}

// ERC721ReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721ReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721ReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721ReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721ReceiverSession struct {
	Contract     *ERC721Receiver   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721ReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721ReceiverCallerSession struct {
	Contract *ERC721ReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721ReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721ReceiverTransactorSession struct {
	Contract     *ERC721ReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721ReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721ReceiverRaw struct {
	Contract *ERC721Receiver // Generic contract binding to access the raw methods on
}

// ERC721ReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721ReceiverCallerRaw struct {
	Contract *ERC721ReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721ReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721ReceiverTransactorRaw struct {
	Contract *ERC721ReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Receiver creates a new instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721Receiver(address common.Address, backend bind.ContractBackend) (*ERC721Receiver, error) {
	contract, err := bindERC721Receiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Receiver{ERC721ReceiverCaller: ERC721ReceiverCaller{contract: contract}, ERC721ReceiverTransactor: ERC721ReceiverTransactor{contract: contract}, ERC721ReceiverFilterer: ERC721ReceiverFilterer{contract: contract}}, nil
}

// NewERC721ReceiverCaller creates a new read-only instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721ReceiverCaller(address common.Address, caller bind.ContractCaller) (*ERC721ReceiverCaller, error) {
	contract, err := bindERC721Receiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ReceiverCaller{contract: contract}, nil
}

// NewERC721ReceiverTransactor creates a new write-only instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721ReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721ReceiverTransactor, error) {
	contract, err := bindERC721Receiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ReceiverTransactor{contract: contract}, nil
}

// NewERC721ReceiverFilterer creates a new log filterer instance of ERC721Receiver, bound to a specific deployed contract.
func NewERC721ReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721ReceiverFilterer, error) {
	contract, err := bindERC721Receiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721ReceiverFilterer{contract: contract}, nil
}

// bindERC721Receiver binds a generic wrapper to an already deployed contract.
func bindERC721Receiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Receiver *ERC721ReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Receiver.Contract.ERC721ReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Receiver *ERC721ReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.ERC721ReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Receiver *ERC721ReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.ERC721ReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Receiver *ERC721ReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Receiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Receiver *ERC721ReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Receiver *ERC721ReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.contract.Transact(opts, method, params...)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(_operator address, _from address, _tokenId uint256, _data bytes) returns(bytes4)
func (_ERC721Receiver *ERC721ReceiverTransactor) OnERC721Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Receiver.contract.Transact(opts, "onERC721Received", _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(_operator address, _from address, _tokenId uint256, _data bytes) returns(bytes4)
func (_ERC721Receiver *ERC721ReceiverSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.OnERC721Received(&_ERC721Receiver.TransactOpts, _operator, _from, _tokenId, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(_operator address, _from address, _tokenId uint256, _data bytes) returns(bytes4)
func (_ERC721Receiver *ERC721ReceiverTransactorSession) OnERC721Received(_operator common.Address, _from common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Receiver.Contract.OnERC721Received(&_ERC721Receiver.TransactOpts, _operator, _from, _tokenId, _data)
}

// ERC721TokenABI is the input ABI used to generate the binding from.
const ERC721TokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// ERC721TokenBin is the compiled bytecode used for deploying new contracts.
const ERC721TokenBin = `0x60806040523480156200001157600080fd5b50604051620011133803806200111383398101604052805160208201519082019101620000677f01ffc9a70000000000000000000000000000000000000000000000000000000064010000000062000137810204565b6200009b7f80ac58cd0000000000000000000000000000000000000000000000000000000064010000000062000137810204565b8151620000b0906005906020850190620001a4565b508051620000c6906006906020840190620001a4565b50620000fb7f780e9d630000000000000000000000000000000000000000000000000000000064010000000062000137810204565b6200012f7f5b5e139f0000000000000000000000000000000000000000000000000000000064010000000062000137810204565b505062000249565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200016757600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001e757805160ff191683800117855562000217565b8280016001018555821562000217579182015b8281111562000217578251825591602001919060010190620001fa565b506200022592915062000229565b5090565b6200024691905b8082111562000225576000815560010162000230565b90565b610eba80620002596000396000f3006080604052600436106100f05763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146100f557806306fdde031461012b578063081812fc146101b5578063095ea7b3146101e957806318160ddd1461020f57806319fa8f501461023657806323b872dd146102685780632f745c591461029257806342842e0e146102b65780634f6ccce7146102e05780636352211e146102f857806370a082311461031057806395d89b4114610331578063a22cb46514610346578063b88d4fde1461036c578063c87b56dd146103db578063e985e9c5146103f3575b600080fd5b34801561010157600080fd5b50610117600160e060020a03196004351661041a565b604080519115158252519081900360200190f35b34801561013757600080fd5b50610140610439565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017a578181015183820152602001610162565b50505050905090810190601f1680156101a75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101c157600080fd5b506101cd6004356104d0565b60408051600160a060020a039092168252519081900360200190f35b3480156101f557600080fd5b5061020d600160a060020a03600435166024356104eb565b005b34801561021b57600080fd5b506102246105a1565b60408051918252519081900360200190f35b34801561024257600080fd5b5061024b6105a7565b60408051600160e060020a03199092168252519081900360200190f35b34801561027457600080fd5b5061020d600160a060020a03600435811690602435166044356105cb565b34801561029e57600080fd5b50610224600160a060020a0360043516602435610659565b3480156102c257600080fd5b5061020d600160a060020a03600435811690602435166044356106a6565b3480156102ec57600080fd5b506102246004356106c7565b34801561030457600080fd5b506101cd6004356106fc565b34801561031c57600080fd5b50610224600160a060020a0360043516610726565b34801561033d57600080fd5b50610140610759565b34801561035257600080fd5b5061020d600160a060020a036004351660243515156107ba565b34801561037857600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261020d94600160a060020a03813581169560248035909216956044359536956084940191819084018382808284375094975061083e9650505050505050565b3480156103e757600080fd5b50610140600435610866565b3480156103ff57600080fd5b50610117600160a060020a036004358116906024351661091b565b600160e060020a03191660009081526020819052604090205460ff1690565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104c55780601f1061049a576101008083540402835291602001916104c5565b820191906000526020600020905b8154815290600101906020018083116104a857829003601f168201915b505050505090505b90565b600090815260026020526040902054600160a060020a031690565b60006104f6826106fc565b9050600160a060020a03838116908216141561051157600080fd5b33600160a060020a038216148061052d575061052d813361091b565b151561053857600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b60095490565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081565b6105d53382610949565b15156105e057600080fd5b600160a060020a03821615156105f557600080fd5b6105ff83826109a8565b6106098382610a19565b6106138282610b20565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b600061066483610726565b821061066f57600080fd5b600160a060020a038316600090815260076020526040902080548390811061069357fe5b9060005260206000200154905092915050565b6106c2838383602060405190810160405280600081525061083e565b505050565b60006106d16105a1565b82106106dc57600080fd5b60098054839081106106ea57fe5b90600052602060002001549050919050565b600081815260016020526040812054600160a060020a031680151561072057600080fd5b92915050565b6000600160a060020a038216151561073d57600080fd5b50600160a060020a031660009081526003602052604090205490565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104c55780601f1061049a576101008083540402835291602001916104c5565b600160a060020a0382163314156107d057600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b6108498484846105cb565b61085584848484610b69565b151561086057600080fd5b50505050565b606061087182610cd6565b151561087c57600080fd5b6000828152600b602090815260409182902080548351601f60026000196101006001861615020190931692909204918201849004840281018401909452808452909183018282801561090f5780601f106108e45761010080835404028352916020019161090f565b820191906000526020600020905b8154815290600101906020018083116108f257829003601f168201915b50505050509050919050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b600080610955836106fc565b905080600160a060020a031684600160a060020a03161480610990575083600160a060020a0316610985846104d0565b600160a060020a0316145b806109a057506109a0818561091b565b949350505050565b81600160a060020a03166109bb826106fc565b600160a060020a0316146109ce57600080fd5b600081815260026020526040902054600160a060020a031615610a15576000818152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff191690555b5050565b6000806000610a288585610cf3565b600084815260086020908152604080832054600160a060020a0389168452600790925290912054909350610a6390600163ffffffff610d8916565b600160a060020a038616600090815260076020526040902080549193509083908110610a8b57fe5b90600052602060002001549050806007600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515610acb57fe5b6000918252602080832090910192909255600160a060020a0387168152600790915260409020805490610b02906000198301610e51565b50600093845260086020526040808520859055908452909220555050565b6000610b2c8383610da0565b50600160a060020a039091166000908152600760209081526040808320805460018101825590845282842081018590559383526008909152902055565b600080610b7e85600160a060020a0316610e30565b1515610b8d5760019150610ccd565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015610c20578181015183820152602001610c08565b50505050905090810190601f168015610c4d5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015610c6f57600080fd5b505af1158015610c83573d6000803e3d6000fd5b505050506040513d6020811015610c9957600080fd5b5051600160e060020a031981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b600090815260016020526040902054600160a060020a0316151590565b81600160a060020a0316610d06826106fc565b600160a060020a031614610d1957600080fd5b600160a060020a038216600090815260036020526040902054610d4390600163ffffffff610d8916565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b60008083831115610d9957600080fd5b5050900390565b600081815260016020526040902054600160a060020a031615610dc257600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388169081179091558452600390915290912054610e1091610e38565b600160a060020a0390921660009081526003602052604090209190915550565b6000903b1190565b600082820183811015610e4a57600080fd5b9392505050565b8154818355818111156106c2576000838152602090206106c29181019083016104cd91905b80821115610e8a5760008155600101610e76565b50905600a165627a7a723058206e01b4dff858a5f1224de40c3cd5ec9f7fdab45c5a3200f9a11ce0ebcc6a76990029`

// DeployERC721Token deploys a new Ethereum contract, binding an instance of ERC721Token to it.
func DeployERC721Token(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *ERC721Token, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721TokenBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Token{ERC721TokenCaller: ERC721TokenCaller{contract: contract}, ERC721TokenTransactor: ERC721TokenTransactor{contract: contract}, ERC721TokenFilterer: ERC721TokenFilterer{contract: contract}}, nil
}

// ERC721Token is an auto generated Go binding around an Ethereum contract.
type ERC721Token struct {
	ERC721TokenCaller     // Read-only binding to the contract
	ERC721TokenTransactor // Write-only binding to the contract
	ERC721TokenFilterer   // Log filterer for contract events
}

// ERC721TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721TokenSession struct {
	Contract     *ERC721Token      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721TokenCallerSession struct {
	Contract *ERC721TokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC721TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TokenTransactorSession struct {
	Contract     *ERC721TokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC721TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721TokenRaw struct {
	Contract *ERC721Token // Generic contract binding to access the raw methods on
}

// ERC721TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721TokenCallerRaw struct {
	Contract *ERC721TokenCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TokenTransactorRaw struct {
	Contract *ERC721TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Token creates a new instance of ERC721Token, bound to a specific deployed contract.
func NewERC721Token(address common.Address, backend bind.ContractBackend) (*ERC721Token, error) {
	contract, err := bindERC721Token(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Token{ERC721TokenCaller: ERC721TokenCaller{contract: contract}, ERC721TokenTransactor: ERC721TokenTransactor{contract: contract}, ERC721TokenFilterer: ERC721TokenFilterer{contract: contract}}, nil
}

// NewERC721TokenCaller creates a new read-only instance of ERC721Token, bound to a specific deployed contract.
func NewERC721TokenCaller(address common.Address, caller bind.ContractCaller) (*ERC721TokenCaller, error) {
	contract, err := bindERC721Token(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenCaller{contract: contract}, nil
}

// NewERC721TokenTransactor creates a new write-only instance of ERC721Token, bound to a specific deployed contract.
func NewERC721TokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721TokenTransactor, error) {
	contract, err := bindERC721Token(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenTransactor{contract: contract}, nil
}

// NewERC721TokenFilterer creates a new log filterer instance of ERC721Token, bound to a specific deployed contract.
func NewERC721TokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721TokenFilterer, error) {
	contract, err := bindERC721Token(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenFilterer{contract: contract}, nil
}

// bindERC721Token binds a generic wrapper to an already deployed contract.
func bindERC721Token(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Token *ERC721TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Token.Contract.ERC721TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Token *ERC721TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Token.Contract.ERC721TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Token *ERC721TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Token.Contract.ERC721TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Token *ERC721TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Token *ERC721TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Token *ERC721TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Token.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721Token *ERC721TokenCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721Token *ERC721TokenSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721Token.Contract.InterfaceIdERC165(&_ERC721Token.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721Token *ERC721TokenCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721Token.Contract.InterfaceIdERC165(&_ERC721Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721Token *ERC721TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721Token *ERC721TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Token.Contract.BalanceOf(&_ERC721Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Token.Contract.BalanceOf(&_ERC721Token.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721Token *ERC721TokenCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721Token *ERC721TokenSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.GetApproved(&_ERC721Token.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721Token *ERC721TokenCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.GetApproved(&_ERC721Token.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Token *ERC721TokenCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Token *ERC721TokenSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Token.Contract.IsApprovedForAll(&_ERC721Token.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_ERC721Token *ERC721TokenCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721Token.Contract.IsApprovedForAll(&_ERC721Token.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Token *ERC721TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Token *ERC721TokenSession) Name() (string, error) {
	return _ERC721Token.Contract.Name(&_ERC721Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721Token *ERC721TokenCallerSession) Name() (string, error) {
	return _ERC721Token.Contract.Name(&_ERC721Token.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721Token *ERC721TokenCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721Token *ERC721TokenSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.OwnerOf(&_ERC721Token.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721Token *ERC721TokenCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Token.Contract.OwnerOf(&_ERC721Token.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Token *ERC721TokenCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Token *ERC721TokenSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Token.Contract.SupportsInterface(&_ERC721Token.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721Token *ERC721TokenCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721Token.Contract.SupportsInterface(&_ERC721Token.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Token *ERC721TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Token *ERC721TokenSession) Symbol() (string, error) {
	return _ERC721Token.Contract.Symbol(&_ERC721Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721Token *ERC721TokenCallerSession) Symbol() (string, error) {
	return _ERC721Token.Contract.Symbol(&_ERC721Token.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721Token *ERC721TokenCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721Token *ERC721TokenSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenByIndex(&_ERC721Token.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenByIndex(&_ERC721Token.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_ERC721Token *ERC721TokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_ERC721Token *ERC721TokenSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenOfOwnerByIndex(&_ERC721Token.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Token.Contract.TokenOfOwnerByIndex(&_ERC721Token.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721Token *ERC721TokenCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721Token *ERC721TokenSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Token.Contract.TokenURI(&_ERC721Token.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_ERC721Token *ERC721TokenCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721Token.Contract.TokenURI(&_ERC721Token.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Token *ERC721TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Token *ERC721TokenSession) TotalSupply() (*big.Int, error) {
	return _ERC721Token.Contract.TotalSupply(&_ERC721Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721Token *ERC721TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Token.Contract.TotalSupply(&_ERC721Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Token *ERC721TokenTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Token *ERC721TokenSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.Approve(&_ERC721Token.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721Token *ERC721TokenTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.Approve(&_ERC721Token.TransactOpts, _to, _tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Token *ERC721TokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Token *ERC721TokenSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Token.Contract.SafeTransferFrom(&_ERC721Token.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_ERC721Token *ERC721TokenTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721Token.Contract.SafeTransferFrom(&_ERC721Token.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721Token *ERC721TokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721Token *ERC721TokenSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Token.Contract.SetApprovalForAll(&_ERC721Token.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_ERC721Token *ERC721TokenTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721Token.Contract.SetApprovalForAll(&_ERC721Token.TransactOpts, _to, _approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Token *ERC721TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Token *ERC721TokenSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.TransferFrom(&_ERC721Token.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_ERC721Token *ERC721TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Token.Contract.TransferFrom(&_ERC721Token.TransactOpts, _from, _to, _tokenId)
}

// ERC721TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Token contract.
type ERC721TokenApprovalIterator struct {
	Event *ERC721TokenApproval // Event containing the contract specifics and raw log

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
func (it *ERC721TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenApproval)
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
		it.Event = new(ERC721TokenApproval)
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
func (it *ERC721TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenApproval represents a Approval event raised by the ERC721Token contract.
type ERC721TokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721Token *ERC721TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenApprovalIterator{contract: _ERC721Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721Token *ERC721TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721TokenApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenApproval)
				if err := _ERC721Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721TokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721Token contract.
type ERC721TokenApprovalForAllIterator struct {
	Event *ERC721TokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721TokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenApprovalForAll)
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
		it.Event = new(ERC721TokenApprovalForAll)
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
func (it *ERC721TokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenApprovalForAll represents a ApprovalForAll event raised by the ERC721Token contract.
type ERC721TokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Token *ERC721TokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721TokenApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Token.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenApprovalForAllIterator{contract: _ERC721Token.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721Token *ERC721TokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721TokenApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721Token.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenApprovalForAll)
				if err := _ERC721Token.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Token contract.
type ERC721TokenTransferIterator struct {
	Event *ERC721TokenTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721TokenTransfer)
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
		it.Event = new(ERC721TokenTransfer)
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
func (it *ERC721TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721TokenTransfer represents a Transfer event raised by the ERC721Token contract.
type ERC721TokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721Token *ERC721TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721TokenTransferIterator{contract: _ERC721Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721Token *ERC721TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721TokenTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721TokenTransfer)
				if err := _ERC721Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// HubABI is the input ABI used to generate the binding from.
const HubABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"dueTime\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createEvent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rate\",\"type\":\"uint16\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEvents\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"serviceFeeRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getEventList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"events\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// HubBin is the compiled bytecode used for deploying new contracts.
const HubBin = `0x60806040526003805461ffff19166101f417905534801561001f57600080fd5b50336000908152602081905260409020805460ff19166001179055613a80806100496000396000f3006080604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630374ab0b81146100b057806312065fe0146100da57806329ff0773146101015780633ccfd60b1461011d578063404aa8f914610132578063493280d21461014757806361d1bc94146101885780636ccf16b0146101b4578063b5f8558c146101cc578063b6db75a0146101e1578063d0e30db01461020a575b005b3480156100bc57600080fd5b506100ae602460048035828101929101359035604435606435610212565b3480156100e657600080fd5b506100ef61038c565b60408051918252519081900360200190f35b34801561010d57600080fd5b506100ae61ffff60043516610391565b34801561012957600080fd5b506100ae61042c565b34801561013e57600080fd5b506100ef6104df565b34801561015357600080fd5b5061015f6004356104e5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561019457600080fd5b5061019d61050d565b6040805161ffff9092168252519081900360200190f35b3480156101c057600080fd5b5061015f600435610517565b3480156101d857600080fd5b506100ef6105da565b3480156101ed57600080fd5b506101f66105e0565b604080519115158252519081900360200190f35b6100ae6105f6565b60028054600101905560035460009086908690309061ffff168787876102366105f8565b73ffffffffffffffffffffffffffffffffffffffff8616602082015261ffff85166040820152606081018490526080810183905260a0810182905260c080825281018790528060e08101898980828437820191505098505050505050505050604051809103906000f0801580156102b1573d6000803e3d6000fd5b50604080517fc85e0be2000000000000000000000000000000000000000000000000000000008152336004820152905191925073ffffffffffffffffffffffffffffffffffffffff83169163c85e0be29160248082019260009290919082900301818387803b15801561032357600080fd5b505af1158015610337573d6000803e3d6000fd5b50506002546000908152600160205260409020805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff949094169390931790925550505050505050565b303190565b3360009081526020819052604090205460ff16151560011461041457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792061646d696e2063616e206163636573732069740000000000000000604482015290519081900360640190fd5b6003805461ffff191661ffff92909216919091179055565b3360009081526020819052604090205460ff1615156001146104af57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f6f6e6c792061646d696e2063616e206163636573732069740000000000000000604482015290519081900360640190fd5b6040513390303180156108fc02916000818181858888f193505050501580156104dc573d6000803e3d6000fd5b50565b60025490565b60016020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60035461ffff1681565b6002546000908211156105b157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f696e707574206d757374206265206c657373207468616e206f7220657175616c60448201527f206576656e747300000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060009081526001602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b60025481565b3360009081526020819052604090205460ff1690565b565b60405161344c8061060983390190560060806040526000600e819055601281905566470de4df82000060135566016bcc41e900006014556015805460ff19166001179055600a60165560178054620a000062ffffff19909116179055601a805461ffff19908116606417909155601b829055601c8054821660f0179055601d829055601e80549091169055601f553480156200008a57600080fd5b506040516200344c3803806200344c83398101604090815281516020830151918301516060840151608085015160a0860151939095019491929091908580620000fc7f01ffc9a70000000000000000000000000000000000000000000000000000000064010000000062000228810204565b620001307f80ac58cd0000000000000000000000000000000000000000000000000000000064010000000062000228810204565b81516200014590600590602085019062000295565b5080516200015b90600690602084019062000295565b50620001907f780e9d630000000000000000000000000000000000000000000000000000000064010000000062000228810204565b620001c47f5b5e139f0000000000000000000000000000000000000000000000000000000064010000000062000228810204565b5050600c8054600160a060020a031916600160a060020a039690961695909517909455336000908152600d60205260409020805460ff19166001179055601a805461ffff191661ffff9490941693909317909255601855601955601355506200033a565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200025857600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002d857805160ff191683800117855562000308565b8280016001018555821562000308579182015b8281111562000308578251825591602001919060010190620002eb565b50620003169291506200031a565b5090565b6200033791905b8082111562000316576000815560010162000321565b90565b613102806200034a6000396000f3006080604052600436106102bb5763ffffffff60e060020a60003504166301ffc9a781146102c057806302dce02c146102f657806303a865ee1461031357806306fdde0314610333578063081812fc146103bd57806308c0204b146103f1578063095ea7b31461040957806309ec6cc71461042d57806318160ddd1461044557806319fa8f501461046c5780631aa3a0081461049e57806321858521146104a657806323037a85146104bb57806323a132a8146104d357806323b872dd146104e857806324bcdfbd146105125780632f745c591461053e5780633ccfd60b146105625780633d3a0a181461057757806342842e0e1461058c57806346609d7f146105b65780634f6ccce7146105cb57806358f06098146105e35780635f3d7fa1146105fb5780636352211e146106105780636b52b73e1461062857806370a082311461063d57806371efdc211461065e57806378e9792514610676578063840f0e521461068b57806387638f9e146106a0578063890fb12b146106c45780638abdf5aa146106ef5780638b424267146107045780638dc5193e1461071957806391b7f5ed1461073157806395d89b41146107495780639fdccfb81461075e578063a035b1fe1461077a578063a22cb4651461078f578063a31ff5a4146107b5578063a4d66daf146107d6578063a78b1426146107eb578063a807411e14610800578063ad7a672f14610815578063b88d4fde1461082a578063b9ced5fe14610899578063bd853041146108ba578063c85e0be2146108cf578063c87b56dd146108f0578063c94e26bc14610908578063dd6388fb14610929578063df1dd8261461094a578063e4f01a4614610955578063e692eca81461096d578063e941fa78146109ad578063e985e9c5146109c2578063ef9089d6146109e9578063f1521f91146109fe578063f659cc0514610a1f578063ff3f19b514610a37575b600080fd5b3480156102cc57600080fd5b506102e2600160e060020a031960043516610a52565b604080519115158252519081900360200190f35b34801561030257600080fd5b50610311600435602435610a71565b005b34801561031f57600080fd5b506103116004803560248101910135610b6f565b34801561033f57600080fd5b50610348610cab565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561038257818101518382015260200161036a565b50505050905090810190601f1680156103af5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103c957600080fd5b506103d5600435610d42565b60408051600160a060020a039092168252519081900360200190f35b3480156103fd57600080fd5b50610311600435610d5d565b34801561041557600080fd5b50610311600160a060020a0360043516602435610e3b565b34801561043957600080fd5b50610311600435610ef1565b34801561045157600080fd5b5061045a610f9f565b60408051918252519081900360200190f35b34801561047857600080fd5b50610481610fa5565b60408051600160e060020a03199092168252519081900360200190f35b610311610fc9565b3480156104b257600080fd5b5061045a6112ae565b3480156104c757600080fd5b506103116004356112b4565b3480156104df57600080fd5b506102e2611313565b3480156104f457600080fd5b50610311600160a060020a03600435811690602435166044356114e0565b34801561051e57600080fd5b5061052761156e565b6040805161ffff9092168252519081900360200190f35b34801561054a57600080fd5b5061045a600160a060020a0360043516602435611578565b34801561056e57600080fd5b506103116115c5565b34801561058357600080fd5b5061045a611662565b34801561059857600080fd5b50610311600160a060020a0360043581169060243516604435611668565b3480156105c257600080fd5b506102e2611684565b3480156105d757600080fd5b5061045a60043561168d565b3480156105ef57600080fd5b506103116004356116c2565b34801561060757600080fd5b5061045a611735565b34801561061c57600080fd5b506103d560043561173b565b34801561063457600080fd5b50610527611765565b34801561064957600080fd5b5061045a600160a060020a036004351661176f565b34801561066a57600080fd5b506102e26004356117a2565b34801561068257600080fd5b5061045a6117b7565b34801561069757600080fd5b506102e26117bd565b3480156106ac57600080fd5b50610311600435600160a060020a03602435166117cb565b3480156106d057600080fd5b506106d9611aa0565b6040805160ff9092168252519081900360200190f35b3480156106fb57600080fd5b5061045a611aaf565b34801561071057600080fd5b5061045a611ab5565b34801561072557600080fd5b50610311600435611abb565b34801561073d57600080fd5b50610311600435611b2e565b34801561075557600080fd5b50610348611b8d565b34801561076a57600080fd5b5061031161ffff60043516611bee565b34801561078657600080fd5b5061045a611c60565b34801561079b57600080fd5b50610311600160a060020a03600435166024351515611c66565b3480156107c157600080fd5b5061045a600160a060020a0360043516611cea565b3480156107e257600080fd5b506106d9611cfc565b3480156107f757600080fd5b506102e2611d05565b34801561080c57600080fd5b50610527611d13565b34801561082157600080fd5b5061045a611d1d565b34801561083657600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261031194600160a060020a038135811695602480359092169560443595369560849401918190840183828082843750949750611d239650505050505050565b3480156108a557600080fd5b506102e2600160a060020a0360043516611d45565b3480156108c657600080fd5b5061045a611d5a565b3480156108db57600080fd5b50610311600160a060020a0360043516611d60565b3480156108fc57600080fd5b50610348600435611dde565b34801561091457600080fd5b50610311600160a060020a0360043516611e93565b34801561093557600080fd5b50610311600160a060020a0360043516612259565b610311600435612349565b34801561096157600080fd5b506102e260043561250f565b34801561097957600080fd5b50610985600435612524565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b3480156109b957600080fd5b5061031161254f565b3480156109ce57600080fd5b506102e2600160a060020a0360043581169060243516612735565b3480156109f557600080fd5b5061045a612763565b348015610a0a57600080fd5b50610311600160a060020a0360043516612769565b348015610a2b57600080fd5b50610311600435612880565b348015610a4357600080fd5b5061031160ff600435166128f3565b600160e060020a03191660009081526020819052604090205460ff1690565b6000610a7c8361173b565b905033600160a060020a0382161480610a9a5750610a9a8133612735565b1515610af0576040805160e560020a62461bcd02815260206004820152601360248201527f6f6e6c79206f776e65722063616e2073656c6c00000000000000000000000000604482015290519081900360640190fd5b50600e8054600190810191829055604080516060810182523380825260208083019788528284019687526000958652600f81528386209251835473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178355965182850155945160029091015592825260109093522080549091019055565b336000908152600d602052604081205460ff161515600114610bc9576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b60175462010000900460ff16821115610c52576040805160e560020a62461bcd02815260206004820152602560248201527f73657420746f6f206d616e79207469636b65747320696e207468652073616d6560448201527f2074696d65000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060005b60ff8116821115610ca657600160116000858560ff8616818110610c7657fe5b60209081029290920135835250810191909152604001600020805460ff1916911515919091179055600101610c56565b505050565b60058054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d375780601f10610d0c57610100808354040283529160200191610d37565b820191906000526020600020905b815481529060010190602001808311610d1a57829003601f168201915b505050505090505b90565b600090815260026020526040902054600160a060020a031690565b336000908152600d602052604090205460ff161515600114610db7576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b6016548111610e36576040805160e560020a62461bcd02815260206004820152602a60248201527f6d6178417474656e64656573206d75737420626520677265617465722074686160448201527f6e206f726967696e616c00000000000000000000000000000000000000000000606482015290519081900360840190fd5b601655565b6000610e468261173b565b9050600160a060020a038381169082161415610e6157600080fd5b33600160a060020a0382161480610e7d5750610e7d8133612735565b1515610e8857600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b610ef961304e565b506000818152600f602090815260409182902082516060810184528154600160a060020a0316808252600183015493820193909352600290910154928101929092523314610f91576040805160e560020a62461bcd02815260206004820152601d60248201527f6f6e6c79206f776e65722063616e2063616e63656c2074726164696e67000000604482015290519081900360640190fd5b610f9b82336129e9565b5050565b60095490565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081565b6000610ff4610fe5601454601254612a4990919063ffffffff16565b6013549063ffffffff612a8216565b6013819055341015611050576040805160e560020a62461bcd02815260206004820152601060248201527f6e6f7420656e6f7567682076616c756500000000000000000000000000000000604482015290519081900360640190fd5b60155460ff1661105f3361176f565b11156110b5576040805160e560020a62461bcd02815260206004820152601660248201527f746f6f206d616e79207469636b657420616d6f756e7400000000000000000000604482015290519081900360640190fd5b6016546012541115611111576040805160e560020a62461bcd02815260206004820152601860248201527f746f6f206d616e79207469636b65742072656769737465720000000000000000604482015290519081900360640190fd5b6018541561117357601854421015611173576040805160e560020a62461bcd02815260206004820152601e60248201527f7469636b65742073616c65206973206e6f742073746172746564207965740000604482015290519081900360640190fd5b601954156111d5576019544211156111d5576040805160e560020a62461bcd02815260206004820152601460248201527f7469636b65742073616c6520697320656e646564000000000000000000000000604482015290519081900360640190fd5b601a546013546112169161120791612710916111fb919061ffff1663ffffffff612a4916565b9063ffffffff612a9416565b601b549063ffffffff612a8216565b601b55601c5460135461124e9161123f91612710916111fb919061ffff1663ffffffff612a4916565b601d549063ffffffff612a8216565b601d5560125461125f903390612ab7565b60135461127390349063ffffffff612b0616565b905060008111156112ab57604051339082156108fc029083906000818181858888f19350505050158015610f9b573d6000803e3d6000fd5b50565b60125481565b336000908152600d602052604090205460ff16151560011461130e576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601455565b60006001815b600c60009054906101000a9004600160a060020a0316600160a060020a031663404aa8f96040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561136c57600080fd5b505af1158015611380573d6000803e3d6000fd5b505050506040513d602081101561139657600080fd5b505182116114d657600c54604080517f493280d2000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a039092169163493280d2916024808201926020929091908290030181600087803b15801561140557600080fd5b505af1158015611419573d6000803e3d6000fd5b505050506040513d602081101561142f57600080fd5b5051604080517f840f0e520000000000000000000000000000000000000000000000000000000081529051919250600160a060020a0383169163840f0e52916004808201926020929091908290030181600087803b15801561149057600080fd5b505af11580156114a4573d6000803e3d6000fd5b505050506040513d60208110156114ba57600080fd5b505115156114cb57600092506114db565b600190910190611319565b600192505b505090565b6114ea3382612b1d565b15156114f557600080fd5b600160a060020a038216151561150a57600080fd5b6115148382612b7c565b61151e8382612beb565b6115288282612cf2565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b60175461ffff1681565b60006115838361176f565b821061158e57600080fd5b600160a060020a03831660009081526007602052604090208054839081106115b257fe5b9060005260206000200154905092915050565b336000908152600d602052604090205460ff16151560011461161f576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601b5433906108fc9061163a9030319063ffffffff612b0616565b6040518115909202916000818181858888f193505050501580156112ab573d6000803e3d6000fd5b60195481565b610ca68383836020604051908101604052806000815250611d23565b601e5460ff1681565b6000611697610f9f565b82106116a257600080fd5b60098054839081106116b057fe5b90600052602060002001549050919050565b336000908152600d602052604090205460ff16151560011461171c576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601b5461172f908263ffffffff612a8216565b601b5550565b60165481565b600081815260016020526040812054600160a060020a031680151561175f57600080fd5b92915050565b601c5461ffff1681565b6000600160a060020a038216151561178657600080fd5b50600160a060020a031660009081526003602052604090205490565b60116020526000908152604090205460ff1681565b60185481565b601e54610100900460ff1681565b60006117d68361173b565b905033600160a060020a03821614806117f457506117f48133612735565b80611807575033600160a060020a038316145b151561185d576040805160e560020a62461bcd02815260206004820152601360248201527f6f6e6c79206f776e65722063616e2073656c6c00000000000000000000000000604482015290519081900360640190fd5b611865611313565b15156001146118be576040805160e560020a62461bcd02815260206004820152601260248201527f6d75737420626520616c6c20736574746c650000000000000000000000000000604482015290519081900360640190fd5b600c54600160a060020a03161515611920576040805160e560020a62461bcd02815260206004820152601b60248201527f68756220636f6e747261637420646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b6019541580159061193357506019544210155b1515611989576040805160e560020a62461bcd02815260206004820152601160248201527f7469636b6574206973206f6e2073616c65000000000000000000000000000000604482015290519081900360640190fd5b601e5460ff1615156001146119e8576040805160e560020a62461bcd02815260206004820152601860248201527f74686973207469636b6574206973206c6f7365207465616d0000000000000000604482015290519081900360640190fd5b33600160a060020a0383161415611a3657601d54604051600160a060020a0384169180156108fc02916000818181858888f19350505050158015611a30573d6000803e3d6000fd5b50610ca6565b80600160a060020a03166108fc611a726012546111fb601d54611a66601b54601f54612b0690919063ffffffff16565b9063ffffffff612b0616565b6040518115909202916000818181858888f19350505050158015611a9a573d6000803e3d6000fd5b50505050565b60175462010000900460ff1681565b601b5481565b601d5481565b336000908152600d602052604090205460ff161515600114611b15576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601d54611b28908263ffffffff612a8216565b601d5550565b336000908152600d602052604090205460ff161515600114611b88576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601355565b60068054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d375780601f10610d0c57610100808354040283529160200191610d37565b336000908152600d602052604090205460ff161515600114611c48576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b6017805461ffff191661ffff92909216919091179055565b60135481565b600160a060020a038216331415611c7c57600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b60106020526000908152604090205481565b60155460ff1681565b601e54610100900460ff1690565b601a5461ffff1681565b601f5481565b611d2e8484846114e0565b611d3a84848484612d3b565b1515611a9a57600080fd5b600d6020526000908152604090205460ff1681565b600e5481565b336000908152600d602052604090205460ff161515600114611dba576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600160a060020a03166000908152600d60205260409020805460ff19166001179055565b6060611de982612ea8565b1515611df457600080fd5b6000828152600b602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015611e875780601f10611e5c57610100808354040283529160200191611e87565b820191906000526020600020905b815481529060010190602001808311611e6a57829003601f168201915b50505050509050919050565b336000908152600d602052604081205460ff161515600114611eed576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600c54600160a060020a03161515611f4f576040805160e560020a62461bcd02815260206004820152601b60248201527f68756220636f6e747261637420646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b60195415801590611f6257506019544210155b1515611fb8576040805160e560020a62461bcd02815260206004820152601160248201527f7469636b6574206973206f6e2073616c65000000000000000000000000000000604482015290519081900360640190fd5b601e54610100900460ff1615612018576040805160e560020a62461bcd02815260206004820152600e60248201527f616c726561647920736574746c65000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03821630146122335781905080600160a060020a031663f1521f9161204560125461173b565b6040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050600060405180830381600087803b15801561209057600080fd5b505af11580156120a4573d6000803e3d6000fd5b5050505080600160a060020a03166358f06098601b546040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156120f357600080fd5b505af1158015612107573d6000803e3d6000fd5b5050505080600160a060020a0316638dc5193e601d546040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561215657600080fd5b505af115801561216a573d6000803e3d6000fd5b5050604080517ff659cc05000000000000000000000000000000000000000000000000000000008152303160048201529051600160a060020a038516935063f659cc059250602480830192600092919082900301818387803b1580156121cf57600080fd5b505af11580156121e3573d6000803e3d6000fd5b5050604051600160a060020a0385169250303180156108fc029250906000818181858888f1935050505015801561221e573d6000803e3d6000fd5b50601e805461ffff1916610100179055610f9b565b61223d3031612880565b601e805461ff001960ff19909116600117166101001790555050565b336000908152600d602052604090205460ff1615156001146122b3576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600160a060020a0381166000908152600d602052604090205460ff161515600114612328576040805160e560020a62461bcd02815260206004820152601760248201527f6f6e6c7920686f73742063616e2062652072656d6f7665000000000000000000604482015290519081900360640190fd5b600160a060020a03166000908152600d60205260409020805460ff19169055565b61235161304e565b6000828152600f60205260408120600101541515612393576040805160e560020a62461bcd028152602060048201526000602482015290519081900360640190fd5b6000838152600f602090815260409182902082516060810184528154600160a060020a03168152600182015492810192909252600201549181018290529250341015612429576040805160e560020a62461bcd02815260206004820152601a60248201527f7469636b6574207072696365206973206e6f7420656e6f756768000000000000604482015290519081900360640190fd5b81516017546040808501519051600160a060020a039093169261ffff909216900380156108fc02916000818181858888f19350505050158015612470573d6000803e3d6000fd5b50601754601b5461248a9161ffff1663ffffffff612a8216565b601b5560408201516124a390349063ffffffff612b0616565b905060008111156124dd57604051339082156108fc029083906000818181858888f193505050501580156124db573d6000803e3d6000fd5b505b6124ef826020015183600001516129e9565b61250182600001518360200151612beb565b610ca6338360200151612cf2565b60009081526011602052604090205460ff1690565b600f60205260009081526040902080546001820154600290920154600160a060020a03909116919083565b336000908152600d602052604090205460ff1615156001146125a9576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600c54600160a060020a0316151561260b576040805160e560020a62461bcd02815260206004820152601b60248201527f68756220636f6e747261637420646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b601b54600010612665576040805160e560020a62461bcd02815260206004820152601b60248201527f7365727669636520666565206973206e6f742072657175697265640000000000604482015290519081900360640190fd5b601b54303110156126c0576040805160e560020a62461bcd02815260206004820152601560248201527f62616c616e6365206973206e6f7420656e6f7567680000000000000000000000604482015290519081900360640190fd5b600c60009054906101000a9004600160a060020a0316600160a060020a031663d0e30db0601b546040518263ffffffff1660e060020a0281526004016000604051808303818588803b15801561271557600080fd5b505af1158015612729573d6000803e3d6000fd5b50506000601b55505050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b60145481565b336000908152600d602052604090205460ff1615156001146127c3576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600160a060020a0381161515612849576040805160e560020a62461bcd02815260206004820152602160248201527f6c617374206c6f736520636f6e747261637420646f6573206e6f74206578697360448201527f7400000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b601e8054600160a060020a03909216620100000275ffffffffffffffffffffffffffffffffffffffff000019909216919091179055565b336000908152600d602052604090205460ff1615156001146128da576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601f546128ed908263ffffffff612a8216565b601f5550565b336000908152600d602052604090205460ff16151560011461294d576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b60155460ff908116908216116129d3576040805160e560020a62461bcd02815260206004820152602360248201527f6c696d6974206d7573742062652067726561746572207468616e206f7269676960448201527f6e616c0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6015805460ff191660ff92909216919091179055565b600160a060020a031660009081526010602090815260408083208054600019908101909155938352600f9091528120805473ffffffffffffffffffffffffffffffffffffffff191681556001810182905560020155600e80549091019055565b600080831515612a5c5760009150612a7b565b50828202828482811515612a6c57fe5b0414612a7757600080fd5b8091505b5092915050565b600082820183811015612a7757600080fd5b600080808311612aa357600080fd5b8284811515612aae57fe5b04949350505050565b612ac18282612ec5565b600980546000838152600a60205260408120829055600182018355919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af015550565b60008083831115612b1657600080fd5b5050900390565b600080612b298361173b565b905080600160a060020a031684600160a060020a03161480612b64575083600160a060020a0316612b5984610d42565b600160a060020a0316145b80612b745750612b748185612735565b949350505050565b81600160a060020a0316612b8f8261173b565b600160a060020a031614612ba257600080fd5b600081815260026020526040902054600160a060020a031615610f9b576000908152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff1916905550565b6000806000612bfa8585612f20565b600084815260086020908152604080832054600160a060020a0389168452600790925290912054909350612c3590600163ffffffff612b0616565b600160a060020a038616600090815260076020526040902080549193509083908110612c5d57fe5b90600052602060002001549050806007600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515612c9d57fe5b6000918252602080832090910192909255600160a060020a0387168152600790915260409020805490612cd4906000198301613079565b50600093845260086020526040808520859055908452909220555050565b6000612cfe8383612fb6565b50600160a060020a039091166000908152600760209081526040808320805460018101825590845282842081018590559383526008909152902055565b600080612d5085600160a060020a0316613046565b1515612d5f5760019150612e9f565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015612df2578181015183820152602001612dda565b50505050905090810190601f168015612e1f5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015612e4157600080fd5b505af1158015612e55573d6000803e3d6000fd5b505050506040513d6020811015612e6b57600080fd5b5051600160e060020a031981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b600090815260016020526040902054600160a060020a0316151590565b600160a060020a0382161515612eda57600080fd5b612ee48282612cf2565b6040518190600160a060020a038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b81600160a060020a0316612f338261173b565b600160a060020a031614612f4657600080fd5b600160a060020a038216600090815260036020526040902054612f7090600163ffffffff612b0616565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600081815260016020526040902054600160a060020a031615612fd857600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038816908117909155845260039091529091205461302691612a82565b600160a060020a0390921660009081526003602052604090209190915550565b6000903b1190565b6060604051908101604052806000600160a060020a0316815260200160008152602001600081525090565b815481835581811115610ca657600083815260209020610ca6918101908301610d3f91905b808211156130b2576000815560010161309e565b509056006f6e6c7920686f73742063616e20616363657373206974000000000000000000a165627a7a72305820a50eccb2e8220b420dfb51a940355e1e621485c4a4c1fe232972472fab0ca1b90029a165627a7a72305820a664be4a06cfc814a1affa6e97cb10a3b8076fbd31b6f46e6e81105e7ea822730029`

// DeployHub deploys a new Ethereum contract, binding an instance of Hub to it.
func DeployHub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hub, error) {
	parsed, err := abi.JSON(strings.NewReader(HubABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hub{HubCaller: HubCaller{contract: contract}, HubTransactor: HubTransactor{contract: contract}, HubFilterer: HubFilterer{contract: contract}}, nil
}

// Hub is an auto generated Go binding around an Ethereum contract.
type Hub struct {
	HubCaller     // Read-only binding to the contract
	HubTransactor // Write-only binding to the contract
	HubFilterer   // Log filterer for contract events
}

// HubCaller is an auto generated read-only Go binding around an Ethereum contract.
type HubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HubSession struct {
	Contract     *Hub              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HubCallerSession struct {
	Contract *HubCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HubTransactorSession struct {
	Contract     *HubTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HubRaw is an auto generated low-level Go binding around an Ethereum contract.
type HubRaw struct {
	Contract *Hub // Generic contract binding to access the raw methods on
}

// HubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HubCallerRaw struct {
	Contract *HubCaller // Generic read-only contract binding to access the raw methods on
}

// HubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HubTransactorRaw struct {
	Contract *HubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHub creates a new instance of Hub, bound to a specific deployed contract.
func NewHub(address common.Address, backend bind.ContractBackend) (*Hub, error) {
	contract, err := bindHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hub{HubCaller: HubCaller{contract: contract}, HubTransactor: HubTransactor{contract: contract}, HubFilterer: HubFilterer{contract: contract}}, nil
}

// NewHubCaller creates a new read-only instance of Hub, bound to a specific deployed contract.
func NewHubCaller(address common.Address, caller bind.ContractCaller) (*HubCaller, error) {
	contract, err := bindHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HubCaller{contract: contract}, nil
}

// NewHubTransactor creates a new write-only instance of Hub, bound to a specific deployed contract.
func NewHubTransactor(address common.Address, transactor bind.ContractTransactor) (*HubTransactor, error) {
	contract, err := bindHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HubTransactor{contract: contract}, nil
}

// NewHubFilterer creates a new log filterer instance of Hub, bound to a specific deployed contract.
func NewHubFilterer(address common.Address, filterer bind.ContractFilterer) (*HubFilterer, error) {
	contract, err := bindHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HubFilterer{contract: contract}, nil
}

// bindHub binds a generic wrapper to an already deployed contract.
func bindHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hub *HubRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hub.Contract.HubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hub *HubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.Contract.HubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hub *HubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hub.Contract.HubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hub *HubCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hub *HubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hub *HubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hub.Contract.contract.Transact(opts, method, params...)
}

// EventList is a free data retrieval call binding the contract method 0x493280d2.
//
// Solidity: function eventList( uint256) constant returns(address)
func (_Hub *HubCaller) EventList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Hub.contract.Call(opts, out, "eventList", arg0)
	return *ret0, err
}

// EventList is a free data retrieval call binding the contract method 0x493280d2.
//
// Solidity: function eventList( uint256) constant returns(address)
func (_Hub *HubSession) EventList(arg0 *big.Int) (common.Address, error) {
	return _Hub.Contract.EventList(&_Hub.CallOpts, arg0)
}

// EventList is a free data retrieval call binding the contract method 0x493280d2.
//
// Solidity: function eventList( uint256) constant returns(address)
func (_Hub *HubCallerSession) EventList(arg0 *big.Int) (common.Address, error) {
	return _Hub.Contract.EventList(&_Hub.CallOpts, arg0)
}

// Events is a free data retrieval call binding the contract method 0xb5f8558c.
//
// Solidity: function events() constant returns(uint256)
func (_Hub *HubCaller) Events(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hub.contract.Call(opts, out, "events")
	return *ret0, err
}

// Events is a free data retrieval call binding the contract method 0xb5f8558c.
//
// Solidity: function events() constant returns(uint256)
func (_Hub *HubSession) Events() (*big.Int, error) {
	return _Hub.Contract.Events(&_Hub.CallOpts)
}

// Events is a free data retrieval call binding the contract method 0xb5f8558c.
//
// Solidity: function events() constant returns(uint256)
func (_Hub *HubCallerSession) Events() (*big.Int, error) {
	return _Hub.Contract.Events(&_Hub.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Hub *HubCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hub.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Hub *HubSession) GetBalance() (*big.Int, error) {
	return _Hub.Contract.GetBalance(&_Hub.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Hub *HubCallerSession) GetBalance() (*big.Int, error) {
	return _Hub.Contract.GetBalance(&_Hub.CallOpts)
}

// GetEventList is a free data retrieval call binding the contract method 0x6ccf16b0.
//
// Solidity: function getEventList(id uint256) constant returns(address)
func (_Hub *HubCaller) GetEventList(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Hub.contract.Call(opts, out, "getEventList", id)
	return *ret0, err
}

// GetEventList is a free data retrieval call binding the contract method 0x6ccf16b0.
//
// Solidity: function getEventList(id uint256) constant returns(address)
func (_Hub *HubSession) GetEventList(id *big.Int) (common.Address, error) {
	return _Hub.Contract.GetEventList(&_Hub.CallOpts, id)
}

// GetEventList is a free data retrieval call binding the contract method 0x6ccf16b0.
//
// Solidity: function getEventList(id uint256) constant returns(address)
func (_Hub *HubCallerSession) GetEventList(id *big.Int) (common.Address, error) {
	return _Hub.Contract.GetEventList(&_Hub.CallOpts, id)
}

// GetEvents is a free data retrieval call binding the contract method 0x404aa8f9.
//
// Solidity: function getEvents() constant returns(uint256)
func (_Hub *HubCaller) GetEvents(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hub.contract.Call(opts, out, "getEvents")
	return *ret0, err
}

// GetEvents is a free data retrieval call binding the contract method 0x404aa8f9.
//
// Solidity: function getEvents() constant returns(uint256)
func (_Hub *HubSession) GetEvents() (*big.Int, error) {
	return _Hub.Contract.GetEvents(&_Hub.CallOpts)
}

// GetEvents is a free data retrieval call binding the contract method 0x404aa8f9.
//
// Solidity: function getEvents() constant returns(uint256)
func (_Hub *HubCallerSession) GetEvents() (*big.Int, error) {
	return _Hub.Contract.GetEvents(&_Hub.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() constant returns(bool)
func (_Hub *HubCaller) IsAdmin(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Hub.contract.Call(opts, out, "isAdmin")
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() constant returns(bool)
func (_Hub *HubSession) IsAdmin() (bool, error) {
	return _Hub.Contract.IsAdmin(&_Hub.CallOpts)
}

// IsAdmin is a free data retrieval call binding the contract method 0xb6db75a0.
//
// Solidity: function isAdmin() constant returns(bool)
func (_Hub *HubCallerSession) IsAdmin() (bool, error) {
	return _Hub.Contract.IsAdmin(&_Hub.CallOpts)
}

// ServiceFeeRate is a free data retrieval call binding the contract method 0x61d1bc94.
//
// Solidity: function serviceFeeRate() constant returns(uint16)
func (_Hub *HubCaller) ServiceFeeRate(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Hub.contract.Call(opts, out, "serviceFeeRate")
	return *ret0, err
}

// ServiceFeeRate is a free data retrieval call binding the contract method 0x61d1bc94.
//
// Solidity: function serviceFeeRate() constant returns(uint16)
func (_Hub *HubSession) ServiceFeeRate() (uint16, error) {
	return _Hub.Contract.ServiceFeeRate(&_Hub.CallOpts)
}

// ServiceFeeRate is a free data retrieval call binding the contract method 0x61d1bc94.
//
// Solidity: function serviceFeeRate() constant returns(uint16)
func (_Hub *HubCallerSession) ServiceFeeRate() (uint16, error) {
	return _Hub.Contract.ServiceFeeRate(&_Hub.CallOpts)
}

// CreateEvent is a paid mutator transaction binding the contract method 0x0374ab0b.
//
// Solidity: function createEvent(name string, startTime uint256, dueTime uint256, price uint256) returns()
func (_Hub *HubTransactor) CreateEvent(opts *bind.TransactOpts, name string, startTime *big.Int, dueTime *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "createEvent", name, startTime, dueTime, price)
}

// CreateEvent is a paid mutator transaction binding the contract method 0x0374ab0b.
//
// Solidity: function createEvent(name string, startTime uint256, dueTime uint256, price uint256) returns()
func (_Hub *HubSession) CreateEvent(name string, startTime *big.Int, dueTime *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.CreateEvent(&_Hub.TransactOpts, name, startTime, dueTime, price)
}

// CreateEvent is a paid mutator transaction binding the contract method 0x0374ab0b.
//
// Solidity: function createEvent(name string, startTime uint256, dueTime uint256, price uint256) returns()
func (_Hub *HubTransactorSession) CreateEvent(name string, startTime *big.Int, dueTime *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Hub.Contract.CreateEvent(&_Hub.TransactOpts, name, startTime, dueTime, price)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Hub *HubTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Hub *HubSession) Deposit() (*types.Transaction, error) {
	return _Hub.Contract.Deposit(&_Hub.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Hub *HubTransactorSession) Deposit() (*types.Transaction, error) {
	return _Hub.Contract.Deposit(&_Hub.TransactOpts)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x29ff0773.
//
// Solidity: function setFeeRate(rate uint16) returns()
func (_Hub *HubTransactor) SetFeeRate(opts *bind.TransactOpts, rate uint16) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "setFeeRate", rate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x29ff0773.
//
// Solidity: function setFeeRate(rate uint16) returns()
func (_Hub *HubSession) SetFeeRate(rate uint16) (*types.Transaction, error) {
	return _Hub.Contract.SetFeeRate(&_Hub.TransactOpts, rate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x29ff0773.
//
// Solidity: function setFeeRate(rate uint16) returns()
func (_Hub *HubTransactorSession) SetFeeRate(rate uint16) (*types.Transaction, error) {
	return _Hub.Contract.SetFeeRate(&_Hub.TransactOpts, rate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Hub *HubTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hub.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Hub *HubSession) Withdraw() (*types.Transaction, error) {
	return _Hub.Contract.Withdraw(&_Hub.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Hub *HubTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Hub.Contract.Withdraw(&_Hub.TransactOpts)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820921e517de20ce110365210dd39a28f49cee5d5cfd7da4ce73511e9bd71dbd6e80029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterfaceWithLookupABI is the input ABI used to generate the binding from.
const SupportsInterfaceWithLookupABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// SupportsInterfaceWithLookupBin is the compiled bytecode used for deploying new contracts.
const SupportsInterfaceWithLookupBin = `0x608060405234801561001057600080fd5b506100437f01ffc9a700000000000000000000000000000000000000000000000000000000640100000000610048810204565b6100b4565b7fffffffff00000000000000000000000000000000000000000000000000000000808216141561007757600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b610166806100c36000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a7811461005057806319fa8f501461009b575b600080fd5b34801561005c57600080fd5b506100877bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19600435166100e2565b604080519115158252519081900360200190f35b3480156100a757600080fd5b506100b0610116565b604080517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff199092168252519081900360200190f35b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191660009081526020819052604090205460ff1690565b7f01ffc9a700000000000000000000000000000000000000000000000000000000815600a165627a7a723058201e98a33712224c327f500b1d66ccda209920ddb363a7c5d3d15d9473b03280540029`

// DeploySupportsInterfaceWithLookup deploys a new Ethereum contract, binding an instance of SupportsInterfaceWithLookup to it.
func DeploySupportsInterfaceWithLookup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SupportsInterfaceWithLookup, error) {
	parsed, err := abi.JSON(strings.NewReader(SupportsInterfaceWithLookupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SupportsInterfaceWithLookupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SupportsInterfaceWithLookup{SupportsInterfaceWithLookupCaller: SupportsInterfaceWithLookupCaller{contract: contract}, SupportsInterfaceWithLookupTransactor: SupportsInterfaceWithLookupTransactor{contract: contract}, SupportsInterfaceWithLookupFilterer: SupportsInterfaceWithLookupFilterer{contract: contract}}, nil
}

// SupportsInterfaceWithLookup is an auto generated Go binding around an Ethereum contract.
type SupportsInterfaceWithLookup struct {
	SupportsInterfaceWithLookupCaller     // Read-only binding to the contract
	SupportsInterfaceWithLookupTransactor // Write-only binding to the contract
	SupportsInterfaceWithLookupFilterer   // Log filterer for contract events
}

// SupportsInterfaceWithLookupCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportsInterfaceWithLookupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportsInterfaceWithLookupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SupportsInterfaceWithLookupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupportsInterfaceWithLookupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupportsInterfaceWithLookupSession struct {
	Contract     *SupportsInterfaceWithLookup // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// SupportsInterfaceWithLookupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupportsInterfaceWithLookupCallerSession struct {
	Contract *SupportsInterfaceWithLookupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// SupportsInterfaceWithLookupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupportsInterfaceWithLookupTransactorSession struct {
	Contract     *SupportsInterfaceWithLookupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// SupportsInterfaceWithLookupRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupRaw struct {
	Contract *SupportsInterfaceWithLookup // Generic contract binding to access the raw methods on
}

// SupportsInterfaceWithLookupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupCallerRaw struct {
	Contract *SupportsInterfaceWithLookupCaller // Generic read-only contract binding to access the raw methods on
}

// SupportsInterfaceWithLookupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupportsInterfaceWithLookupTransactorRaw struct {
	Contract *SupportsInterfaceWithLookupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupportsInterfaceWithLookup creates a new instance of SupportsInterfaceWithLookup, bound to a specific deployed contract.
func NewSupportsInterfaceWithLookup(address common.Address, backend bind.ContractBackend) (*SupportsInterfaceWithLookup, error) {
	contract, err := bindSupportsInterfaceWithLookup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SupportsInterfaceWithLookup{SupportsInterfaceWithLookupCaller: SupportsInterfaceWithLookupCaller{contract: contract}, SupportsInterfaceWithLookupTransactor: SupportsInterfaceWithLookupTransactor{contract: contract}, SupportsInterfaceWithLookupFilterer: SupportsInterfaceWithLookupFilterer{contract: contract}}, nil
}

// NewSupportsInterfaceWithLookupCaller creates a new read-only instance of SupportsInterfaceWithLookup, bound to a specific deployed contract.
func NewSupportsInterfaceWithLookupCaller(address common.Address, caller bind.ContractCaller) (*SupportsInterfaceWithLookupCaller, error) {
	contract, err := bindSupportsInterfaceWithLookup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupportsInterfaceWithLookupCaller{contract: contract}, nil
}

// NewSupportsInterfaceWithLookupTransactor creates a new write-only instance of SupportsInterfaceWithLookup, bound to a specific deployed contract.
func NewSupportsInterfaceWithLookupTransactor(address common.Address, transactor bind.ContractTransactor) (*SupportsInterfaceWithLookupTransactor, error) {
	contract, err := bindSupportsInterfaceWithLookup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupportsInterfaceWithLookupTransactor{contract: contract}, nil
}

// NewSupportsInterfaceWithLookupFilterer creates a new log filterer instance of SupportsInterfaceWithLookup, bound to a specific deployed contract.
func NewSupportsInterfaceWithLookupFilterer(address common.Address, filterer bind.ContractFilterer) (*SupportsInterfaceWithLookupFilterer, error) {
	contract, err := bindSupportsInterfaceWithLookup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupportsInterfaceWithLookupFilterer{contract: contract}, nil
}

// bindSupportsInterfaceWithLookup binds a generic wrapper to an already deployed contract.
func bindSupportsInterfaceWithLookup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SupportsInterfaceWithLookupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterfaceWithLookupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterfaceWithLookupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterfaceWithLookupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SupportsInterfaceWithLookup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupportsInterfaceWithLookup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupportsInterfaceWithLookup.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _SupportsInterfaceWithLookup.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupSession) InterfaceIdERC165() ([4]byte, error) {
	return _SupportsInterfaceWithLookup.Contract.InterfaceIdERC165(&_SupportsInterfaceWithLookup.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _SupportsInterfaceWithLookup.Contract.InterfaceIdERC165(&_SupportsInterfaceWithLookup.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SupportsInterfaceWithLookup.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterface(&_SupportsInterfaceWithLookup.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_SupportsInterfaceWithLookup *SupportsInterfaceWithLookupCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _SupportsInterfaceWithLookup.Contract.SupportsInterface(&_SupportsInterfaceWithLookup.CallOpts, _interfaceId)
}

// TicketSaleABI is the input ABI used to generate the binding from.
const TicketSaleABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ticketId\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"requestTrading\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ticketIds\",\"type\":\"uint256[]\"}],\"name\":\"setUsedTickets\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxAttendees\",\"type\":\"uint256\"}],\"name\":\"setMaxAttendees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tradeId\",\"type\":\"uint256\"}],\"name\":\"cancelTrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tickets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nextPrice\",\"type\":\"uint256\"}],\"name\":\"setNextPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allSettle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tradeFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dueTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWin\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_serviceFee\",\"type\":\"uint256\"}],\"name\":\"addServiceFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAttendees\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardFeeRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedTickets\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSettle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_ticketId\",\"type\":\"uint256\"},{\"name\":\"_lastLose\",\"type\":\"address\"}],\"name\":\"Reward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxMarkedTickets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"serviceFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_rewardFee\",\"type\":\"uint256\"}],\"name\":\"addRewardFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fee\",\"type\":\"uint16\"}],\"name\":\"setTradeFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tradingTicketsOfOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"limit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIsSettle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"serviceFeeRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hosts\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tradings\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"host\",\"type\":\"address\"}],\"name\":\"setHost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_winTicket\",\"type\":\"address\"}],\"name\":\"Win\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeHost\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tradeId\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ticketId\",\"type\":\"uint256\"}],\"name\":\"isUsedTicket\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tradingList\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"ticketId\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lastLose\",\"type\":\"address\"}],\"name\":\"setLastLose\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_totalBalance\",\"type\":\"uint256\"}],\"name\":\"addTotalBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_limit\",\"type\":\"uint8\"}],\"name\":\"setLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"hubAddr\",\"type\":\"address\"},{\"name\":\"ratio\",\"type\":\"uint16\"},{\"name\":\"_startTime\",\"type\":\"uint256\"},{\"name\":\"_dueTime\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// TicketSaleBin is the compiled bytecode used for deploying new contracts.
const TicketSaleBin = `0x60806040526000600e819055601281905566470de4df82000060135566016bcc41e900006014556015805460ff19166001179055600a60165560178054620a000062ffffff19909116179055601a805461ffff19908116606417909155601b829055601c8054821660f0179055601d829055601e80549091169055601f553480156200008a57600080fd5b506040516200344c3803806200344c83398101604090815281516020830151918301516060840151608085015160a0860151939095019491929091908580620000fc7f01ffc9a70000000000000000000000000000000000000000000000000000000064010000000062000228810204565b620001307f80ac58cd0000000000000000000000000000000000000000000000000000000064010000000062000228810204565b81516200014590600590602085019062000295565b5080516200015b90600690602084019062000295565b50620001907f780e9d630000000000000000000000000000000000000000000000000000000064010000000062000228810204565b620001c47f5b5e139f0000000000000000000000000000000000000000000000000000000064010000000062000228810204565b5050600c8054600160a060020a031916600160a060020a039690961695909517909455336000908152600d60205260409020805460ff19166001179055601a805461ffff191661ffff9490941693909317909255601855601955601355506200033a565b7fffffffff0000000000000000000000000000000000000000000000000000000080821614156200025857600080fd5b7fffffffff00000000000000000000000000000000000000000000000000000000166000908152602081905260409020805460ff19166001179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002d857805160ff191683800117855562000308565b8280016001018555821562000308579182015b8281111562000308578251825591602001919060010190620002eb565b50620003169291506200031a565b5090565b6200033791905b8082111562000316576000815560010162000321565b90565b613102806200034a6000396000f3006080604052600436106102bb5763ffffffff60e060020a60003504166301ffc9a781146102c057806302dce02c146102f657806303a865ee1461031357806306fdde0314610333578063081812fc146103bd57806308c0204b146103f1578063095ea7b31461040957806309ec6cc71461042d57806318160ddd1461044557806319fa8f501461046c5780631aa3a0081461049e57806321858521146104a657806323037a85146104bb57806323a132a8146104d357806323b872dd146104e857806324bcdfbd146105125780632f745c591461053e5780633ccfd60b146105625780633d3a0a181461057757806342842e0e1461058c57806346609d7f146105b65780634f6ccce7146105cb57806358f06098146105e35780635f3d7fa1146105fb5780636352211e146106105780636b52b73e1461062857806370a082311461063d57806371efdc211461065e57806378e9792514610676578063840f0e521461068b57806387638f9e146106a0578063890fb12b146106c45780638abdf5aa146106ef5780638b424267146107045780638dc5193e1461071957806391b7f5ed1461073157806395d89b41146107495780639fdccfb81461075e578063a035b1fe1461077a578063a22cb4651461078f578063a31ff5a4146107b5578063a4d66daf146107d6578063a78b1426146107eb578063a807411e14610800578063ad7a672f14610815578063b88d4fde1461082a578063b9ced5fe14610899578063bd853041146108ba578063c85e0be2146108cf578063c87b56dd146108f0578063c94e26bc14610908578063dd6388fb14610929578063df1dd8261461094a578063e4f01a4614610955578063e692eca81461096d578063e941fa78146109ad578063e985e9c5146109c2578063ef9089d6146109e9578063f1521f91146109fe578063f659cc0514610a1f578063ff3f19b514610a37575b600080fd5b3480156102cc57600080fd5b506102e2600160e060020a031960043516610a52565b604080519115158252519081900360200190f35b34801561030257600080fd5b50610311600435602435610a71565b005b34801561031f57600080fd5b506103116004803560248101910135610b6f565b34801561033f57600080fd5b50610348610cab565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561038257818101518382015260200161036a565b50505050905090810190601f1680156103af5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103c957600080fd5b506103d5600435610d42565b60408051600160a060020a039092168252519081900360200190f35b3480156103fd57600080fd5b50610311600435610d5d565b34801561041557600080fd5b50610311600160a060020a0360043516602435610e3b565b34801561043957600080fd5b50610311600435610ef1565b34801561045157600080fd5b5061045a610f9f565b60408051918252519081900360200190f35b34801561047857600080fd5b50610481610fa5565b60408051600160e060020a03199092168252519081900360200190f35b610311610fc9565b3480156104b257600080fd5b5061045a6112ae565b3480156104c757600080fd5b506103116004356112b4565b3480156104df57600080fd5b506102e2611313565b3480156104f457600080fd5b50610311600160a060020a03600435811690602435166044356114e0565b34801561051e57600080fd5b5061052761156e565b6040805161ffff9092168252519081900360200190f35b34801561054a57600080fd5b5061045a600160a060020a0360043516602435611578565b34801561056e57600080fd5b506103116115c5565b34801561058357600080fd5b5061045a611662565b34801561059857600080fd5b50610311600160a060020a0360043581169060243516604435611668565b3480156105c257600080fd5b506102e2611684565b3480156105d757600080fd5b5061045a60043561168d565b3480156105ef57600080fd5b506103116004356116c2565b34801561060757600080fd5b5061045a611735565b34801561061c57600080fd5b506103d560043561173b565b34801561063457600080fd5b50610527611765565b34801561064957600080fd5b5061045a600160a060020a036004351661176f565b34801561066a57600080fd5b506102e26004356117a2565b34801561068257600080fd5b5061045a6117b7565b34801561069757600080fd5b506102e26117bd565b3480156106ac57600080fd5b50610311600435600160a060020a03602435166117cb565b3480156106d057600080fd5b506106d9611aa0565b6040805160ff9092168252519081900360200190f35b3480156106fb57600080fd5b5061045a611aaf565b34801561071057600080fd5b5061045a611ab5565b34801561072557600080fd5b50610311600435611abb565b34801561073d57600080fd5b50610311600435611b2e565b34801561075557600080fd5b50610348611b8d565b34801561076a57600080fd5b5061031161ffff60043516611bee565b34801561078657600080fd5b5061045a611c60565b34801561079b57600080fd5b50610311600160a060020a03600435166024351515611c66565b3480156107c157600080fd5b5061045a600160a060020a0360043516611cea565b3480156107e257600080fd5b506106d9611cfc565b3480156107f757600080fd5b506102e2611d05565b34801561080c57600080fd5b50610527611d13565b34801561082157600080fd5b5061045a611d1d565b34801561083657600080fd5b50604080516020601f60643560048181013592830184900484028501840190955281845261031194600160a060020a038135811695602480359092169560443595369560849401918190840183828082843750949750611d239650505050505050565b3480156108a557600080fd5b506102e2600160a060020a0360043516611d45565b3480156108c657600080fd5b5061045a611d5a565b3480156108db57600080fd5b50610311600160a060020a0360043516611d60565b3480156108fc57600080fd5b50610348600435611dde565b34801561091457600080fd5b50610311600160a060020a0360043516611e93565b34801561093557600080fd5b50610311600160a060020a0360043516612259565b610311600435612349565b34801561096157600080fd5b506102e260043561250f565b34801561097957600080fd5b50610985600435612524565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b3480156109b957600080fd5b5061031161254f565b3480156109ce57600080fd5b506102e2600160a060020a0360043581169060243516612735565b3480156109f557600080fd5b5061045a612763565b348015610a0a57600080fd5b50610311600160a060020a0360043516612769565b348015610a2b57600080fd5b50610311600435612880565b348015610a4357600080fd5b5061031160ff600435166128f3565b600160e060020a03191660009081526020819052604090205460ff1690565b6000610a7c8361173b565b905033600160a060020a0382161480610a9a5750610a9a8133612735565b1515610af0576040805160e560020a62461bcd02815260206004820152601360248201527f6f6e6c79206f776e65722063616e2073656c6c00000000000000000000000000604482015290519081900360640190fd5b50600e8054600190810191829055604080516060810182523380825260208083019788528284019687526000958652600f81528386209251835473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178355965182850155945160029091015592825260109093522080549091019055565b336000908152600d602052604081205460ff161515600114610bc9576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b60175462010000900460ff16821115610c52576040805160e560020a62461bcd02815260206004820152602560248201527f73657420746f6f206d616e79207469636b65747320696e207468652073616d6560448201527f2074696d65000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b5060005b60ff8116821115610ca657600160116000858560ff8616818110610c7657fe5b60209081029290920135835250810191909152604001600020805460ff1916911515919091179055600101610c56565b505050565b60058054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d375780601f10610d0c57610100808354040283529160200191610d37565b820191906000526020600020905b815481529060010190602001808311610d1a57829003601f168201915b505050505090505b90565b600090815260026020526040902054600160a060020a031690565b336000908152600d602052604090205460ff161515600114610db7576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b6016548111610e36576040805160e560020a62461bcd02815260206004820152602a60248201527f6d6178417474656e64656573206d75737420626520677265617465722074686160448201527f6e206f726967696e616c00000000000000000000000000000000000000000000606482015290519081900360840190fd5b601655565b6000610e468261173b565b9050600160a060020a038381169082161415610e6157600080fd5b33600160a060020a0382161480610e7d5750610e7d8133612735565b1515610e8857600080fd5b600082815260026020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387811691821790925591518593918516917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b610ef961304e565b506000818152600f602090815260409182902082516060810184528154600160a060020a0316808252600183015493820193909352600290910154928101929092523314610f91576040805160e560020a62461bcd02815260206004820152601d60248201527f6f6e6c79206f776e65722063616e2063616e63656c2074726164696e67000000604482015290519081900360640190fd5b610f9b82336129e9565b5050565b60095490565b7f01ffc9a70000000000000000000000000000000000000000000000000000000081565b6000610ff4610fe5601454601254612a4990919063ffffffff16565b6013549063ffffffff612a8216565b6013819055341015611050576040805160e560020a62461bcd02815260206004820152601060248201527f6e6f7420656e6f7567682076616c756500000000000000000000000000000000604482015290519081900360640190fd5b60155460ff1661105f3361176f565b11156110b5576040805160e560020a62461bcd02815260206004820152601660248201527f746f6f206d616e79207469636b657420616d6f756e7400000000000000000000604482015290519081900360640190fd5b6016546012541115611111576040805160e560020a62461bcd02815260206004820152601860248201527f746f6f206d616e79207469636b65742072656769737465720000000000000000604482015290519081900360640190fd5b6018541561117357601854421015611173576040805160e560020a62461bcd02815260206004820152601e60248201527f7469636b65742073616c65206973206e6f742073746172746564207965740000604482015290519081900360640190fd5b601954156111d5576019544211156111d5576040805160e560020a62461bcd02815260206004820152601460248201527f7469636b65742073616c6520697320656e646564000000000000000000000000604482015290519081900360640190fd5b601a546013546112169161120791612710916111fb919061ffff1663ffffffff612a4916565b9063ffffffff612a9416565b601b549063ffffffff612a8216565b601b55601c5460135461124e9161123f91612710916111fb919061ffff1663ffffffff612a4916565b601d549063ffffffff612a8216565b601d5560125461125f903390612ab7565b60135461127390349063ffffffff612b0616565b905060008111156112ab57604051339082156108fc029083906000818181858888f19350505050158015610f9b573d6000803e3d6000fd5b50565b60125481565b336000908152600d602052604090205460ff16151560011461130e576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601455565b60006001815b600c60009054906101000a9004600160a060020a0316600160a060020a031663404aa8f96040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561136c57600080fd5b505af1158015611380573d6000803e3d6000fd5b505050506040513d602081101561139657600080fd5b505182116114d657600c54604080517f493280d2000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a039092169163493280d2916024808201926020929091908290030181600087803b15801561140557600080fd5b505af1158015611419573d6000803e3d6000fd5b505050506040513d602081101561142f57600080fd5b5051604080517f840f0e520000000000000000000000000000000000000000000000000000000081529051919250600160a060020a0383169163840f0e52916004808201926020929091908290030181600087803b15801561149057600080fd5b505af11580156114a4573d6000803e3d6000fd5b505050506040513d60208110156114ba57600080fd5b505115156114cb57600092506114db565b600190910190611319565b600192505b505090565b6114ea3382612b1d565b15156114f557600080fd5b600160a060020a038216151561150a57600080fd5b6115148382612b7c565b61151e8382612beb565b6115288282612cf2565b8082600160a060020a031684600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b60175461ffff1681565b60006115838361176f565b821061158e57600080fd5b600160a060020a03831660009081526007602052604090208054839081106115b257fe5b9060005260206000200154905092915050565b336000908152600d602052604090205460ff16151560011461161f576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601b5433906108fc9061163a9030319063ffffffff612b0616565b6040518115909202916000818181858888f193505050501580156112ab573d6000803e3d6000fd5b60195481565b610ca68383836020604051908101604052806000815250611d23565b601e5460ff1681565b6000611697610f9f565b82106116a257600080fd5b60098054839081106116b057fe5b90600052602060002001549050919050565b336000908152600d602052604090205460ff16151560011461171c576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601b5461172f908263ffffffff612a8216565b601b5550565b60165481565b600081815260016020526040812054600160a060020a031680151561175f57600080fd5b92915050565b601c5461ffff1681565b6000600160a060020a038216151561178657600080fd5b50600160a060020a031660009081526003602052604090205490565b60116020526000908152604090205460ff1681565b60185481565b601e54610100900460ff1681565b60006117d68361173b565b905033600160a060020a03821614806117f457506117f48133612735565b80611807575033600160a060020a038316145b151561185d576040805160e560020a62461bcd02815260206004820152601360248201527f6f6e6c79206f776e65722063616e2073656c6c00000000000000000000000000604482015290519081900360640190fd5b611865611313565b15156001146118be576040805160e560020a62461bcd02815260206004820152601260248201527f6d75737420626520616c6c20736574746c650000000000000000000000000000604482015290519081900360640190fd5b600c54600160a060020a03161515611920576040805160e560020a62461bcd02815260206004820152601b60248201527f68756220636f6e747261637420646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b6019541580159061193357506019544210155b1515611989576040805160e560020a62461bcd02815260206004820152601160248201527f7469636b6574206973206f6e2073616c65000000000000000000000000000000604482015290519081900360640190fd5b601e5460ff1615156001146119e8576040805160e560020a62461bcd02815260206004820152601860248201527f74686973207469636b6574206973206c6f7365207465616d0000000000000000604482015290519081900360640190fd5b33600160a060020a0383161415611a3657601d54604051600160a060020a0384169180156108fc02916000818181858888f19350505050158015611a30573d6000803e3d6000fd5b50610ca6565b80600160a060020a03166108fc611a726012546111fb601d54611a66601b54601f54612b0690919063ffffffff16565b9063ffffffff612b0616565b6040518115909202916000818181858888f19350505050158015611a9a573d6000803e3d6000fd5b50505050565b60175462010000900460ff1681565b601b5481565b601d5481565b336000908152600d602052604090205460ff161515600114611b15576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601d54611b28908263ffffffff612a8216565b601d5550565b336000908152600d602052604090205460ff161515600114611b88576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601355565b60068054604080516020601f6002600019610100600188161502019095169490940493840181900481028201810190925282815260609390929091830182828015610d375780601f10610d0c57610100808354040283529160200191610d37565b336000908152600d602052604090205460ff161515600114611c48576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b6017805461ffff191661ffff92909216919091179055565b60135481565b600160a060020a038216331415611c7c57600080fd5b336000818152600460209081526040808320600160a060020a03871680855290835292819020805460ff1916861515908117909155815190815290519293927f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31929181900390910190a35050565b60106020526000908152604090205481565b60155460ff1681565b601e54610100900460ff1690565b601a5461ffff1681565b601f5481565b611d2e8484846114e0565b611d3a84848484612d3b565b1515611a9a57600080fd5b600d6020526000908152604090205460ff1681565b600e5481565b336000908152600d602052604090205460ff161515600114611dba576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600160a060020a03166000908152600d60205260409020805460ff19166001179055565b6060611de982612ea8565b1515611df457600080fd5b6000828152600b602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015611e875780601f10611e5c57610100808354040283529160200191611e87565b820191906000526020600020905b815481529060010190602001808311611e6a57829003601f168201915b50505050509050919050565b336000908152600d602052604081205460ff161515600114611eed576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600c54600160a060020a03161515611f4f576040805160e560020a62461bcd02815260206004820152601b60248201527f68756220636f6e747261637420646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b60195415801590611f6257506019544210155b1515611fb8576040805160e560020a62461bcd02815260206004820152601160248201527f7469636b6574206973206f6e2073616c65000000000000000000000000000000604482015290519081900360640190fd5b601e54610100900460ff1615612018576040805160e560020a62461bcd02815260206004820152600e60248201527f616c726561647920736574746c65000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03821630146122335781905080600160a060020a031663f1521f9161204560125461173b565b6040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a03168152602001915050600060405180830381600087803b15801561209057600080fd5b505af11580156120a4573d6000803e3d6000fd5b5050505080600160a060020a03166358f06098601b546040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156120f357600080fd5b505af1158015612107573d6000803e3d6000fd5b5050505080600160a060020a0316638dc5193e601d546040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561215657600080fd5b505af115801561216a573d6000803e3d6000fd5b5050604080517ff659cc05000000000000000000000000000000000000000000000000000000008152303160048201529051600160a060020a038516935063f659cc059250602480830192600092919082900301818387803b1580156121cf57600080fd5b505af11580156121e3573d6000803e3d6000fd5b5050604051600160a060020a0385169250303180156108fc029250906000818181858888f1935050505015801561221e573d6000803e3d6000fd5b50601e805461ffff1916610100179055610f9b565b61223d3031612880565b601e805461ff001960ff19909116600117166101001790555050565b336000908152600d602052604090205460ff1615156001146122b3576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600160a060020a0381166000908152600d602052604090205460ff161515600114612328576040805160e560020a62461bcd02815260206004820152601760248201527f6f6e6c7920686f73742063616e2062652072656d6f7665000000000000000000604482015290519081900360640190fd5b600160a060020a03166000908152600d60205260409020805460ff19169055565b61235161304e565b6000828152600f60205260408120600101541515612393576040805160e560020a62461bcd028152602060048201526000602482015290519081900360640190fd5b6000838152600f602090815260409182902082516060810184528154600160a060020a03168152600182015492810192909252600201549181018290529250341015612429576040805160e560020a62461bcd02815260206004820152601a60248201527f7469636b6574207072696365206973206e6f7420656e6f756768000000000000604482015290519081900360640190fd5b81516017546040808501519051600160a060020a039093169261ffff909216900380156108fc02916000818181858888f19350505050158015612470573d6000803e3d6000fd5b50601754601b5461248a9161ffff1663ffffffff612a8216565b601b5560408201516124a390349063ffffffff612b0616565b905060008111156124dd57604051339082156108fc029083906000818181858888f193505050501580156124db573d6000803e3d6000fd5b505b6124ef826020015183600001516129e9565b61250182600001518360200151612beb565b610ca6338360200151612cf2565b60009081526011602052604090205460ff1690565b600f60205260009081526040902080546001820154600290920154600160a060020a03909116919083565b336000908152600d602052604090205460ff1615156001146125a9576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600c54600160a060020a0316151561260b576040805160e560020a62461bcd02815260206004820152601b60248201527f68756220636f6e747261637420646f6573206e6f742065786973740000000000604482015290519081900360640190fd5b601b54600010612665576040805160e560020a62461bcd02815260206004820152601b60248201527f7365727669636520666565206973206e6f742072657175697265640000000000604482015290519081900360640190fd5b601b54303110156126c0576040805160e560020a62461bcd02815260206004820152601560248201527f62616c616e6365206973206e6f7420656e6f7567680000000000000000000000604482015290519081900360640190fd5b600c60009054906101000a9004600160a060020a0316600160a060020a031663d0e30db0601b546040518263ffffffff1660e060020a0281526004016000604051808303818588803b15801561271557600080fd5b505af1158015612729573d6000803e3d6000fd5b50506000601b55505050565b600160a060020a03918216600090815260046020908152604080832093909416825291909152205460ff1690565b60145481565b336000908152600d602052604090205460ff1615156001146127c3576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b600160a060020a0381161515612849576040805160e560020a62461bcd02815260206004820152602160248201527f6c617374206c6f736520636f6e747261637420646f6573206e6f74206578697360448201527f7400000000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b601e8054600160a060020a03909216620100000275ffffffffffffffffffffffffffffffffffffffff000019909216919091179055565b336000908152600d602052604090205460ff1615156001146128da576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b601f546128ed908263ffffffff612a8216565b601f5550565b336000908152600d602052604090205460ff16151560011461294d576040805160e560020a62461bcd02815260206004820152601760248201526000805160206130b7833981519152604482015290519081900360640190fd5b60155460ff908116908216116129d3576040805160e560020a62461bcd02815260206004820152602360248201527f6c696d6974206d7573742062652067726561746572207468616e206f7269676960448201527f6e616c0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6015805460ff191660ff92909216919091179055565b600160a060020a031660009081526010602090815260408083208054600019908101909155938352600f9091528120805473ffffffffffffffffffffffffffffffffffffffff191681556001810182905560020155600e80549091019055565b600080831515612a5c5760009150612a7b565b50828202828482811515612a6c57fe5b0414612a7757600080fd5b8091505b5092915050565b600082820183811015612a7757600080fd5b600080808311612aa357600080fd5b8284811515612aae57fe5b04949350505050565b612ac18282612ec5565b600980546000838152600a60205260408120829055600182018355919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af015550565b60008083831115612b1657600080fd5b5050900390565b600080612b298361173b565b905080600160a060020a031684600160a060020a03161480612b64575083600160a060020a0316612b5984610d42565b600160a060020a0316145b80612b745750612b748185612735565b949350505050565b81600160a060020a0316612b8f8261173b565b600160a060020a031614612ba257600080fd5b600081815260026020526040902054600160a060020a031615610f9b576000908152600260205260409020805473ffffffffffffffffffffffffffffffffffffffff1916905550565b6000806000612bfa8585612f20565b600084815260086020908152604080832054600160a060020a0389168452600790925290912054909350612c3590600163ffffffff612b0616565b600160a060020a038616600090815260076020526040902080549193509083908110612c5d57fe5b90600052602060002001549050806007600087600160a060020a0316600160a060020a0316815260200190815260200160002084815481101515612c9d57fe5b6000918252602080832090910192909255600160a060020a0387168152600790915260409020805490612cd4906000198301613079565b50600093845260086020526040808520859055908452909220555050565b6000612cfe8383612fb6565b50600160a060020a039091166000908152600760209081526040808320805460018101825590845282842081018590559383526008909152902055565b600080612d5085600160a060020a0316613046565b1515612d5f5760019150612e9f565b6040517f150b7a020000000000000000000000000000000000000000000000000000000081523360048201818152600160a060020a03898116602485015260448401889052608060648501908152875160848601528751918a169463150b7a0294938c938b938b93909160a490910190602085019080838360005b83811015612df2578181015183820152602001612dda565b50505050905090810190601f168015612e1f5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b158015612e4157600080fd5b505af1158015612e55573d6000803e3d6000fd5b505050506040513d6020811015612e6b57600080fd5b5051600160e060020a031981167f150b7a020000000000000000000000000000000000000000000000000000000014925090505b50949350505050565b600090815260016020526040902054600160a060020a0316151590565b600160a060020a0382161515612eda57600080fd5b612ee48282612cf2565b6040518190600160a060020a038416906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b81600160a060020a0316612f338261173b565b600160a060020a031614612f4657600080fd5b600160a060020a038216600090815260036020526040902054612f7090600163ffffffff612b0616565b600160a060020a03909216600090815260036020908152604080832094909455918152600190915220805473ffffffffffffffffffffffffffffffffffffffff19169055565b600081815260016020526040902054600160a060020a031615612fd857600080fd5b6000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038816908117909155845260039091529091205461302691612a82565b600160a060020a0390921660009081526003602052604090209190915550565b6000903b1190565b6060604051908101604052806000600160a060020a0316815260200160008152602001600081525090565b815481835581811115610ca657600083815260209020610ca6918101908301610d3f91905b808211156130b2576000815560010161309e565b509056006f6e6c7920686f73742063616e20616363657373206974000000000000000000a165627a7a72305820a50eccb2e8220b420dfb51a940355e1e621485c4a4c1fe232972472fab0ca1b90029`

// DeployTicketSale deploys a new Ethereum contract, binding an instance of TicketSale to it.
func DeployTicketSale(auth *bind.TransactOpts, backend bind.ContractBackend, name string, hubAddr common.Address, ratio uint16, _startTime *big.Int, _dueTime *big.Int, _price *big.Int) (common.Address, *types.Transaction, *TicketSale, error) {
	parsed, err := abi.JSON(strings.NewReader(TicketSaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TicketSaleBin), backend, name, hubAddr, ratio, _startTime, _dueTime, _price)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TicketSale{TicketSaleCaller: TicketSaleCaller{contract: contract}, TicketSaleTransactor: TicketSaleTransactor{contract: contract}, TicketSaleFilterer: TicketSaleFilterer{contract: contract}}, nil
}

// TicketSale is an auto generated Go binding around an Ethereum contract.
type TicketSale struct {
	TicketSaleCaller     // Read-only binding to the contract
	TicketSaleTransactor // Write-only binding to the contract
	TicketSaleFilterer   // Log filterer for contract events
}

// TicketSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type TicketSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TicketSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TicketSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TicketSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TicketSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TicketSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TicketSaleSession struct {
	Contract     *TicketSale       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TicketSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TicketSaleCallerSession struct {
	Contract *TicketSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TicketSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TicketSaleTransactorSession struct {
	Contract     *TicketSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TicketSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type TicketSaleRaw struct {
	Contract *TicketSale // Generic contract binding to access the raw methods on
}

// TicketSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TicketSaleCallerRaw struct {
	Contract *TicketSaleCaller // Generic read-only contract binding to access the raw methods on
}

// TicketSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TicketSaleTransactorRaw struct {
	Contract *TicketSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTicketSale creates a new instance of TicketSale, bound to a specific deployed contract.
func NewTicketSale(address common.Address, backend bind.ContractBackend) (*TicketSale, error) {
	contract, err := bindTicketSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TicketSale{TicketSaleCaller: TicketSaleCaller{contract: contract}, TicketSaleTransactor: TicketSaleTransactor{contract: contract}, TicketSaleFilterer: TicketSaleFilterer{contract: contract}}, nil
}

// NewTicketSaleCaller creates a new read-only instance of TicketSale, bound to a specific deployed contract.
func NewTicketSaleCaller(address common.Address, caller bind.ContractCaller) (*TicketSaleCaller, error) {
	contract, err := bindTicketSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TicketSaleCaller{contract: contract}, nil
}

// NewTicketSaleTransactor creates a new write-only instance of TicketSale, bound to a specific deployed contract.
func NewTicketSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*TicketSaleTransactor, error) {
	contract, err := bindTicketSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TicketSaleTransactor{contract: contract}, nil
}

// NewTicketSaleFilterer creates a new log filterer instance of TicketSale, bound to a specific deployed contract.
func NewTicketSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*TicketSaleFilterer, error) {
	contract, err := bindTicketSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TicketSaleFilterer{contract: contract}, nil
}

// bindTicketSale binds a generic wrapper to an already deployed contract.
func bindTicketSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TicketSaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TicketSale *TicketSaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TicketSale.Contract.TicketSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TicketSale *TicketSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicketSale.Contract.TicketSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TicketSale *TicketSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TicketSale.Contract.TicketSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TicketSale *TicketSaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TicketSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TicketSale *TicketSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicketSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TicketSale *TicketSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TicketSale.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_TicketSale *TicketSaleCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_TicketSale *TicketSaleSession) InterfaceIdERC165() ([4]byte, error) {
	return _TicketSale.Contract.InterfaceIdERC165(&_TicketSale.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_TicketSale *TicketSaleCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _TicketSale.Contract.InterfaceIdERC165(&_TicketSale.CallOpts)
}

// AllSettle is a free data retrieval call binding the contract method 0x23a132a8.
//
// Solidity: function allSettle() constant returns(bool)
func (_TicketSale *TicketSaleCaller) AllSettle(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "allSettle")
	return *ret0, err
}

// AllSettle is a free data retrieval call binding the contract method 0x23a132a8.
//
// Solidity: function allSettle() constant returns(bool)
func (_TicketSale *TicketSaleSession) AllSettle() (bool, error) {
	return _TicketSale.Contract.AllSettle(&_TicketSale.CallOpts)
}

// AllSettle is a free data retrieval call binding the contract method 0x23a132a8.
//
// Solidity: function allSettle() constant returns(bool)
func (_TicketSale *TicketSaleCallerSession) AllSettle() (bool, error) {
	return _TicketSale.Contract.AllSettle(&_TicketSale.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_TicketSale *TicketSaleCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_TicketSale *TicketSaleSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TicketSale.Contract.BalanceOf(&_TicketSale.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TicketSale.Contract.BalanceOf(&_TicketSale.CallOpts, _owner)
}

// DueTime is a free data retrieval call binding the contract method 0x3d3a0a18.
//
// Solidity: function dueTime() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) DueTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "dueTime")
	return *ret0, err
}

// DueTime is a free data retrieval call binding the contract method 0x3d3a0a18.
//
// Solidity: function dueTime() constant returns(uint256)
func (_TicketSale *TicketSaleSession) DueTime() (*big.Int, error) {
	return _TicketSale.Contract.DueTime(&_TicketSale.CallOpts)
}

// DueTime is a free data retrieval call binding the contract method 0x3d3a0a18.
//
// Solidity: function dueTime() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) DueTime() (*big.Int, error) {
	return _TicketSale.Contract.DueTime(&_TicketSale.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_TicketSale *TicketSaleCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_TicketSale *TicketSaleSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _TicketSale.Contract.GetApproved(&_TicketSale.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_TicketSale *TicketSaleCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _TicketSale.Contract.GetApproved(&_TicketSale.CallOpts, _tokenId)
}

// GetIsSettle is a free data retrieval call binding the contract method 0xa78b1426.
//
// Solidity: function getIsSettle() constant returns(bool)
func (_TicketSale *TicketSaleCaller) GetIsSettle(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "getIsSettle")
	return *ret0, err
}

// GetIsSettle is a free data retrieval call binding the contract method 0xa78b1426.
//
// Solidity: function getIsSettle() constant returns(bool)
func (_TicketSale *TicketSaleSession) GetIsSettle() (bool, error) {
	return _TicketSale.Contract.GetIsSettle(&_TicketSale.CallOpts)
}

// GetIsSettle is a free data retrieval call binding the contract method 0xa78b1426.
//
// Solidity: function getIsSettle() constant returns(bool)
func (_TicketSale *TicketSaleCallerSession) GetIsSettle() (bool, error) {
	return _TicketSale.Contract.GetIsSettle(&_TicketSale.CallOpts)
}

// Hosts is a free data retrieval call binding the contract method 0xb9ced5fe.
//
// Solidity: function hosts( address) constant returns(bool)
func (_TicketSale *TicketSaleCaller) Hosts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "hosts", arg0)
	return *ret0, err
}

// Hosts is a free data retrieval call binding the contract method 0xb9ced5fe.
//
// Solidity: function hosts( address) constant returns(bool)
func (_TicketSale *TicketSaleSession) Hosts(arg0 common.Address) (bool, error) {
	return _TicketSale.Contract.Hosts(&_TicketSale.CallOpts, arg0)
}

// Hosts is a free data retrieval call binding the contract method 0xb9ced5fe.
//
// Solidity: function hosts( address) constant returns(bool)
func (_TicketSale *TicketSaleCallerSession) Hosts(arg0 common.Address) (bool, error) {
	return _TicketSale.Contract.Hosts(&_TicketSale.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_TicketSale *TicketSaleCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_TicketSale *TicketSaleSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _TicketSale.Contract.IsApprovedForAll(&_TicketSale.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_TicketSale *TicketSaleCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _TicketSale.Contract.IsApprovedForAll(&_TicketSale.CallOpts, _owner, _operator)
}

// IsSettle is a free data retrieval call binding the contract method 0x840f0e52.
//
// Solidity: function isSettle() constant returns(bool)
func (_TicketSale *TicketSaleCaller) IsSettle(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "isSettle")
	return *ret0, err
}

// IsSettle is a free data retrieval call binding the contract method 0x840f0e52.
//
// Solidity: function isSettle() constant returns(bool)
func (_TicketSale *TicketSaleSession) IsSettle() (bool, error) {
	return _TicketSale.Contract.IsSettle(&_TicketSale.CallOpts)
}

// IsSettle is a free data retrieval call binding the contract method 0x840f0e52.
//
// Solidity: function isSettle() constant returns(bool)
func (_TicketSale *TicketSaleCallerSession) IsSettle() (bool, error) {
	return _TicketSale.Contract.IsSettle(&_TicketSale.CallOpts)
}

// IsUsedTicket is a free data retrieval call binding the contract method 0xe4f01a46.
//
// Solidity: function isUsedTicket(_ticketId uint256) constant returns(bool)
func (_TicketSale *TicketSaleCaller) IsUsedTicket(opts *bind.CallOpts, _ticketId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "isUsedTicket", _ticketId)
	return *ret0, err
}

// IsUsedTicket is a free data retrieval call binding the contract method 0xe4f01a46.
//
// Solidity: function isUsedTicket(_ticketId uint256) constant returns(bool)
func (_TicketSale *TicketSaleSession) IsUsedTicket(_ticketId *big.Int) (bool, error) {
	return _TicketSale.Contract.IsUsedTicket(&_TicketSale.CallOpts, _ticketId)
}

// IsUsedTicket is a free data retrieval call binding the contract method 0xe4f01a46.
//
// Solidity: function isUsedTicket(_ticketId uint256) constant returns(bool)
func (_TicketSale *TicketSaleCallerSession) IsUsedTicket(_ticketId *big.Int) (bool, error) {
	return _TicketSale.Contract.IsUsedTicket(&_TicketSale.CallOpts, _ticketId)
}

// IsWin is a free data retrieval call binding the contract method 0x46609d7f.
//
// Solidity: function isWin() constant returns(bool)
func (_TicketSale *TicketSaleCaller) IsWin(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "isWin")
	return *ret0, err
}

// IsWin is a free data retrieval call binding the contract method 0x46609d7f.
//
// Solidity: function isWin() constant returns(bool)
func (_TicketSale *TicketSaleSession) IsWin() (bool, error) {
	return _TicketSale.Contract.IsWin(&_TicketSale.CallOpts)
}

// IsWin is a free data retrieval call binding the contract method 0x46609d7f.
//
// Solidity: function isWin() constant returns(bool)
func (_TicketSale *TicketSaleCallerSession) IsWin() (bool, error) {
	return _TicketSale.Contract.IsWin(&_TicketSale.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() constant returns(uint8)
func (_TicketSale *TicketSaleCaller) Limit(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "limit")
	return *ret0, err
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() constant returns(uint8)
func (_TicketSale *TicketSaleSession) Limit() (uint8, error) {
	return _TicketSale.Contract.Limit(&_TicketSale.CallOpts)
}

// Limit is a free data retrieval call binding the contract method 0xa4d66daf.
//
// Solidity: function limit() constant returns(uint8)
func (_TicketSale *TicketSaleCallerSession) Limit() (uint8, error) {
	return _TicketSale.Contract.Limit(&_TicketSale.CallOpts)
}

// MaxAttendees is a free data retrieval call binding the contract method 0x5f3d7fa1.
//
// Solidity: function maxAttendees() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) MaxAttendees(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "maxAttendees")
	return *ret0, err
}

// MaxAttendees is a free data retrieval call binding the contract method 0x5f3d7fa1.
//
// Solidity: function maxAttendees() constant returns(uint256)
func (_TicketSale *TicketSaleSession) MaxAttendees() (*big.Int, error) {
	return _TicketSale.Contract.MaxAttendees(&_TicketSale.CallOpts)
}

// MaxAttendees is a free data retrieval call binding the contract method 0x5f3d7fa1.
//
// Solidity: function maxAttendees() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) MaxAttendees() (*big.Int, error) {
	return _TicketSale.Contract.MaxAttendees(&_TicketSale.CallOpts)
}

// MaxMarkedTickets is a free data retrieval call binding the contract method 0x890fb12b.
//
// Solidity: function maxMarkedTickets() constant returns(uint8)
func (_TicketSale *TicketSaleCaller) MaxMarkedTickets(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "maxMarkedTickets")
	return *ret0, err
}

// MaxMarkedTickets is a free data retrieval call binding the contract method 0x890fb12b.
//
// Solidity: function maxMarkedTickets() constant returns(uint8)
func (_TicketSale *TicketSaleSession) MaxMarkedTickets() (uint8, error) {
	return _TicketSale.Contract.MaxMarkedTickets(&_TicketSale.CallOpts)
}

// MaxMarkedTickets is a free data retrieval call binding the contract method 0x890fb12b.
//
// Solidity: function maxMarkedTickets() constant returns(uint8)
func (_TicketSale *TicketSaleCallerSession) MaxMarkedTickets() (uint8, error) {
	return _TicketSale.Contract.MaxMarkedTickets(&_TicketSale.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TicketSale *TicketSaleCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TicketSale *TicketSaleSession) Name() (string, error) {
	return _TicketSale.Contract.Name(&_TicketSale.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TicketSale *TicketSaleCallerSession) Name() (string, error) {
	return _TicketSale.Contract.Name(&_TicketSale.CallOpts)
}

// NextPrice is a free data retrieval call binding the contract method 0xef9089d6.
//
// Solidity: function nextPrice() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) NextPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "nextPrice")
	return *ret0, err
}

// NextPrice is a free data retrieval call binding the contract method 0xef9089d6.
//
// Solidity: function nextPrice() constant returns(uint256)
func (_TicketSale *TicketSaleSession) NextPrice() (*big.Int, error) {
	return _TicketSale.Contract.NextPrice(&_TicketSale.CallOpts)
}

// NextPrice is a free data retrieval call binding the contract method 0xef9089d6.
//
// Solidity: function nextPrice() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) NextPrice() (*big.Int, error) {
	return _TicketSale.Contract.NextPrice(&_TicketSale.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_TicketSale *TicketSaleCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_TicketSale *TicketSaleSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _TicketSale.Contract.OwnerOf(&_TicketSale.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_TicketSale *TicketSaleCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _TicketSale.Contract.OwnerOf(&_TicketSale.CallOpts, _tokenId)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "price")
	return *ret0, err
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_TicketSale *TicketSaleSession) Price() (*big.Int, error) {
	return _TicketSale.Contract.Price(&_TicketSale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) Price() (*big.Int, error) {
	return _TicketSale.Contract.Price(&_TicketSale.CallOpts)
}

// RewardFee is a free data retrieval call binding the contract method 0x8b424267.
//
// Solidity: function rewardFee() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) RewardFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "rewardFee")
	return *ret0, err
}

// RewardFee is a free data retrieval call binding the contract method 0x8b424267.
//
// Solidity: function rewardFee() constant returns(uint256)
func (_TicketSale *TicketSaleSession) RewardFee() (*big.Int, error) {
	return _TicketSale.Contract.RewardFee(&_TicketSale.CallOpts)
}

// RewardFee is a free data retrieval call binding the contract method 0x8b424267.
//
// Solidity: function rewardFee() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) RewardFee() (*big.Int, error) {
	return _TicketSale.Contract.RewardFee(&_TicketSale.CallOpts)
}

// RewardFeeRatio is a free data retrieval call binding the contract method 0x6b52b73e.
//
// Solidity: function rewardFeeRatio() constant returns(uint16)
func (_TicketSale *TicketSaleCaller) RewardFeeRatio(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "rewardFeeRatio")
	return *ret0, err
}

// RewardFeeRatio is a free data retrieval call binding the contract method 0x6b52b73e.
//
// Solidity: function rewardFeeRatio() constant returns(uint16)
func (_TicketSale *TicketSaleSession) RewardFeeRatio() (uint16, error) {
	return _TicketSale.Contract.RewardFeeRatio(&_TicketSale.CallOpts)
}

// RewardFeeRatio is a free data retrieval call binding the contract method 0x6b52b73e.
//
// Solidity: function rewardFeeRatio() constant returns(uint16)
func (_TicketSale *TicketSaleCallerSession) RewardFeeRatio() (uint16, error) {
	return _TicketSale.Contract.RewardFeeRatio(&_TicketSale.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) ServiceFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "serviceFee")
	return *ret0, err
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() constant returns(uint256)
func (_TicketSale *TicketSaleSession) ServiceFee() (*big.Int, error) {
	return _TicketSale.Contract.ServiceFee(&_TicketSale.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) ServiceFee() (*big.Int, error) {
	return _TicketSale.Contract.ServiceFee(&_TicketSale.CallOpts)
}

// ServiceFeeRatio is a free data retrieval call binding the contract method 0xa807411e.
//
// Solidity: function serviceFeeRatio() constant returns(uint16)
func (_TicketSale *TicketSaleCaller) ServiceFeeRatio(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "serviceFeeRatio")
	return *ret0, err
}

// ServiceFeeRatio is a free data retrieval call binding the contract method 0xa807411e.
//
// Solidity: function serviceFeeRatio() constant returns(uint16)
func (_TicketSale *TicketSaleSession) ServiceFeeRatio() (uint16, error) {
	return _TicketSale.Contract.ServiceFeeRatio(&_TicketSale.CallOpts)
}

// ServiceFeeRatio is a free data retrieval call binding the contract method 0xa807411e.
//
// Solidity: function serviceFeeRatio() constant returns(uint16)
func (_TicketSale *TicketSaleCallerSession) ServiceFeeRatio() (uint16, error) {
	return _TicketSale.Contract.ServiceFeeRatio(&_TicketSale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "startTime")
	return *ret0, err
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_TicketSale *TicketSaleSession) StartTime() (*big.Int, error) {
	return _TicketSale.Contract.StartTime(&_TicketSale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) StartTime() (*big.Int, error) {
	return _TicketSale.Contract.StartTime(&_TicketSale.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_TicketSale *TicketSaleCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_TicketSale *TicketSaleSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _TicketSale.Contract.SupportsInterface(&_TicketSale.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_TicketSale *TicketSaleCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _TicketSale.Contract.SupportsInterface(&_TicketSale.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TicketSale *TicketSaleCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TicketSale *TicketSaleSession) Symbol() (string, error) {
	return _TicketSale.Contract.Symbol(&_TicketSale.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TicketSale *TicketSaleCallerSession) Symbol() (string, error) {
	return _TicketSale.Contract.Symbol(&_TicketSale.CallOpts)
}

// Tickets is a free data retrieval call binding the contract method 0x21858521.
//
// Solidity: function tickets() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) Tickets(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "tickets")
	return *ret0, err
}

// Tickets is a free data retrieval call binding the contract method 0x21858521.
//
// Solidity: function tickets() constant returns(uint256)
func (_TicketSale *TicketSaleSession) Tickets() (*big.Int, error) {
	return _TicketSale.Contract.Tickets(&_TicketSale.CallOpts)
}

// Tickets is a free data retrieval call binding the contract method 0x21858521.
//
// Solidity: function tickets() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) Tickets() (*big.Int, error) {
	return _TicketSale.Contract.Tickets(&_TicketSale.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_TicketSale *TicketSaleCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_TicketSale *TicketSaleSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _TicketSale.Contract.TokenByIndex(&_TicketSale.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _TicketSale.Contract.TokenByIndex(&_TicketSale.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_TicketSale *TicketSaleCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_TicketSale *TicketSaleSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _TicketSale.Contract.TokenOfOwnerByIndex(&_TicketSale.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _TicketSale.Contract.TokenOfOwnerByIndex(&_TicketSale.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_TicketSale *TicketSaleCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_TicketSale *TicketSaleSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _TicketSale.Contract.TokenURI(&_TicketSale.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_TicketSale *TicketSaleCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _TicketSale.Contract.TokenURI(&_TicketSale.CallOpts, _tokenId)
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) TotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "totalBalance")
	return *ret0, err
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() constant returns(uint256)
func (_TicketSale *TicketSaleSession) TotalBalance() (*big.Int, error) {
	return _TicketSale.Contract.TotalBalance(&_TicketSale.CallOpts)
}

// TotalBalance is a free data retrieval call binding the contract method 0xad7a672f.
//
// Solidity: function totalBalance() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) TotalBalance() (*big.Int, error) {
	return _TicketSale.Contract.TotalBalance(&_TicketSale.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TicketSale *TicketSaleSession) TotalSupply() (*big.Int, error) {
	return _TicketSale.Contract.TotalSupply(&_TicketSale.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) TotalSupply() (*big.Int, error) {
	return _TicketSale.Contract.TotalSupply(&_TicketSale.CallOpts)
}

// TradeFee is a free data retrieval call binding the contract method 0x24bcdfbd.
//
// Solidity: function tradeFee() constant returns(uint16)
func (_TicketSale *TicketSaleCaller) TradeFee(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "tradeFee")
	return *ret0, err
}

// TradeFee is a free data retrieval call binding the contract method 0x24bcdfbd.
//
// Solidity: function tradeFee() constant returns(uint16)
func (_TicketSale *TicketSaleSession) TradeFee() (uint16, error) {
	return _TicketSale.Contract.TradeFee(&_TicketSale.CallOpts)
}

// TradeFee is a free data retrieval call binding the contract method 0x24bcdfbd.
//
// Solidity: function tradeFee() constant returns(uint16)
func (_TicketSale *TicketSaleCallerSession) TradeFee() (uint16, error) {
	return _TicketSale.Contract.TradeFee(&_TicketSale.CallOpts)
}

// TradingList is a free data retrieval call binding the contract method 0xe692eca8.
//
// Solidity: function tradingList( uint256) constant returns(owner address, ticketId uint256, value uint256)
func (_TicketSale *TicketSaleCaller) TradingList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner    common.Address
	TicketId *big.Int
	Value    *big.Int
}, error) {
	ret := new(struct {
		Owner    common.Address
		TicketId *big.Int
		Value    *big.Int
	})
	out := ret
	err := _TicketSale.contract.Call(opts, out, "tradingList", arg0)
	return *ret, err
}

// TradingList is a free data retrieval call binding the contract method 0xe692eca8.
//
// Solidity: function tradingList( uint256) constant returns(owner address, ticketId uint256, value uint256)
func (_TicketSale *TicketSaleSession) TradingList(arg0 *big.Int) (struct {
	Owner    common.Address
	TicketId *big.Int
	Value    *big.Int
}, error) {
	return _TicketSale.Contract.TradingList(&_TicketSale.CallOpts, arg0)
}

// TradingList is a free data retrieval call binding the contract method 0xe692eca8.
//
// Solidity: function tradingList( uint256) constant returns(owner address, ticketId uint256, value uint256)
func (_TicketSale *TicketSaleCallerSession) TradingList(arg0 *big.Int) (struct {
	Owner    common.Address
	TicketId *big.Int
	Value    *big.Int
}, error) {
	return _TicketSale.Contract.TradingList(&_TicketSale.CallOpts, arg0)
}

// TradingTicketsOfOwner is a free data retrieval call binding the contract method 0xa31ff5a4.
//
// Solidity: function tradingTicketsOfOwner( address) constant returns(uint256)
func (_TicketSale *TicketSaleCaller) TradingTicketsOfOwner(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "tradingTicketsOfOwner", arg0)
	return *ret0, err
}

// TradingTicketsOfOwner is a free data retrieval call binding the contract method 0xa31ff5a4.
//
// Solidity: function tradingTicketsOfOwner( address) constant returns(uint256)
func (_TicketSale *TicketSaleSession) TradingTicketsOfOwner(arg0 common.Address) (*big.Int, error) {
	return _TicketSale.Contract.TradingTicketsOfOwner(&_TicketSale.CallOpts, arg0)
}

// TradingTicketsOfOwner is a free data retrieval call binding the contract method 0xa31ff5a4.
//
// Solidity: function tradingTicketsOfOwner( address) constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) TradingTicketsOfOwner(arg0 common.Address) (*big.Int, error) {
	return _TicketSale.Contract.TradingTicketsOfOwner(&_TicketSale.CallOpts, arg0)
}

// Tradings is a free data retrieval call binding the contract method 0xbd853041.
//
// Solidity: function tradings() constant returns(uint256)
func (_TicketSale *TicketSaleCaller) Tradings(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "tradings")
	return *ret0, err
}

// Tradings is a free data retrieval call binding the contract method 0xbd853041.
//
// Solidity: function tradings() constant returns(uint256)
func (_TicketSale *TicketSaleSession) Tradings() (*big.Int, error) {
	return _TicketSale.Contract.Tradings(&_TicketSale.CallOpts)
}

// Tradings is a free data retrieval call binding the contract method 0xbd853041.
//
// Solidity: function tradings() constant returns(uint256)
func (_TicketSale *TicketSaleCallerSession) Tradings() (*big.Int, error) {
	return _TicketSale.Contract.Tradings(&_TicketSale.CallOpts)
}

// UsedTickets is a free data retrieval call binding the contract method 0x71efdc21.
//
// Solidity: function usedTickets( uint256) constant returns(bool)
func (_TicketSale *TicketSaleCaller) UsedTickets(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TicketSale.contract.Call(opts, out, "usedTickets", arg0)
	return *ret0, err
}

// UsedTickets is a free data retrieval call binding the contract method 0x71efdc21.
//
// Solidity: function usedTickets( uint256) constant returns(bool)
func (_TicketSale *TicketSaleSession) UsedTickets(arg0 *big.Int) (bool, error) {
	return _TicketSale.Contract.UsedTickets(&_TicketSale.CallOpts, arg0)
}

// UsedTickets is a free data retrieval call binding the contract method 0x71efdc21.
//
// Solidity: function usedTickets( uint256) constant returns(bool)
func (_TicketSale *TicketSaleCallerSession) UsedTickets(arg0 *big.Int) (bool, error) {
	return _TicketSale.Contract.UsedTickets(&_TicketSale.CallOpts, arg0)
}

// Reward is a paid mutator transaction binding the contract method 0x87638f9e.
//
// Solidity: function Reward(_ticketId uint256, _lastLose address) returns()
func (_TicketSale *TicketSaleTransactor) Reward(opts *bind.TransactOpts, _ticketId *big.Int, _lastLose common.Address) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "Reward", _ticketId, _lastLose)
}

// Reward is a paid mutator transaction binding the contract method 0x87638f9e.
//
// Solidity: function Reward(_ticketId uint256, _lastLose address) returns()
func (_TicketSale *TicketSaleSession) Reward(_ticketId *big.Int, _lastLose common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.Reward(&_TicketSale.TransactOpts, _ticketId, _lastLose)
}

// Reward is a paid mutator transaction binding the contract method 0x87638f9e.
//
// Solidity: function Reward(_ticketId uint256, _lastLose address) returns()
func (_TicketSale *TicketSaleTransactorSession) Reward(_ticketId *big.Int, _lastLose common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.Reward(&_TicketSale.TransactOpts, _ticketId, _lastLose)
}

// Win is a paid mutator transaction binding the contract method 0xc94e26bc.
//
// Solidity: function Win(_winTicket address) returns()
func (_TicketSale *TicketSaleTransactor) Win(opts *bind.TransactOpts, _winTicket common.Address) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "Win", _winTicket)
}

// Win is a paid mutator transaction binding the contract method 0xc94e26bc.
//
// Solidity: function Win(_winTicket address) returns()
func (_TicketSale *TicketSaleSession) Win(_winTicket common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.Win(&_TicketSale.TransactOpts, _winTicket)
}

// Win is a paid mutator transaction binding the contract method 0xc94e26bc.
//
// Solidity: function Win(_winTicket address) returns()
func (_TicketSale *TicketSaleTransactorSession) Win(_winTicket common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.Win(&_TicketSale.TransactOpts, _winTicket)
}

// AddRewardFee is a paid mutator transaction binding the contract method 0x8dc5193e.
//
// Solidity: function addRewardFee(_rewardFee uint256) returns()
func (_TicketSale *TicketSaleTransactor) AddRewardFee(opts *bind.TransactOpts, _rewardFee *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "addRewardFee", _rewardFee)
}

// AddRewardFee is a paid mutator transaction binding the contract method 0x8dc5193e.
//
// Solidity: function addRewardFee(_rewardFee uint256) returns()
func (_TicketSale *TicketSaleSession) AddRewardFee(_rewardFee *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.AddRewardFee(&_TicketSale.TransactOpts, _rewardFee)
}

// AddRewardFee is a paid mutator transaction binding the contract method 0x8dc5193e.
//
// Solidity: function addRewardFee(_rewardFee uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) AddRewardFee(_rewardFee *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.AddRewardFee(&_TicketSale.TransactOpts, _rewardFee)
}

// AddServiceFee is a paid mutator transaction binding the contract method 0x58f06098.
//
// Solidity: function addServiceFee(_serviceFee uint256) returns()
func (_TicketSale *TicketSaleTransactor) AddServiceFee(opts *bind.TransactOpts, _serviceFee *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "addServiceFee", _serviceFee)
}

// AddServiceFee is a paid mutator transaction binding the contract method 0x58f06098.
//
// Solidity: function addServiceFee(_serviceFee uint256) returns()
func (_TicketSale *TicketSaleSession) AddServiceFee(_serviceFee *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.AddServiceFee(&_TicketSale.TransactOpts, _serviceFee)
}

// AddServiceFee is a paid mutator transaction binding the contract method 0x58f06098.
//
// Solidity: function addServiceFee(_serviceFee uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) AddServiceFee(_serviceFee *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.AddServiceFee(&_TicketSale.TransactOpts, _serviceFee)
}

// AddTotalBalance is a paid mutator transaction binding the contract method 0xf659cc05.
//
// Solidity: function addTotalBalance(_totalBalance uint256) returns()
func (_TicketSale *TicketSaleTransactor) AddTotalBalance(opts *bind.TransactOpts, _totalBalance *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "addTotalBalance", _totalBalance)
}

// AddTotalBalance is a paid mutator transaction binding the contract method 0xf659cc05.
//
// Solidity: function addTotalBalance(_totalBalance uint256) returns()
func (_TicketSale *TicketSaleSession) AddTotalBalance(_totalBalance *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.AddTotalBalance(&_TicketSale.TransactOpts, _totalBalance)
}

// AddTotalBalance is a paid mutator transaction binding the contract method 0xf659cc05.
//
// Solidity: function addTotalBalance(_totalBalance uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) AddTotalBalance(_totalBalance *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.AddTotalBalance(&_TicketSale.TransactOpts, _totalBalance)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_TicketSale *TicketSaleTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_TicketSale *TicketSaleSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.Approve(&_TicketSale.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.Approve(&_TicketSale.TransactOpts, _to, _tokenId)
}

// CancelTrade is a paid mutator transaction binding the contract method 0x09ec6cc7.
//
// Solidity: function cancelTrade(_tradeId uint256) returns()
func (_TicketSale *TicketSaleTransactor) CancelTrade(opts *bind.TransactOpts, _tradeId *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "cancelTrade", _tradeId)
}

// CancelTrade is a paid mutator transaction binding the contract method 0x09ec6cc7.
//
// Solidity: function cancelTrade(_tradeId uint256) returns()
func (_TicketSale *TicketSaleSession) CancelTrade(_tradeId *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.CancelTrade(&_TicketSale.TransactOpts, _tradeId)
}

// CancelTrade is a paid mutator transaction binding the contract method 0x09ec6cc7.
//
// Solidity: function cancelTrade(_tradeId uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) CancelTrade(_tradeId *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.CancelTrade(&_TicketSale.TransactOpts, _tradeId)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_TicketSale *TicketSaleTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_TicketSale *TicketSaleSession) Register() (*types.Transaction, error) {
	return _TicketSale.Contract.Register(&_TicketSale.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_TicketSale *TicketSaleTransactorSession) Register() (*types.Transaction, error) {
	return _TicketSale.Contract.Register(&_TicketSale.TransactOpts)
}

// RemoveHost is a paid mutator transaction binding the contract method 0xdd6388fb.
//
// Solidity: function removeHost(addr address) returns()
func (_TicketSale *TicketSaleTransactor) RemoveHost(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "removeHost", addr)
}

// RemoveHost is a paid mutator transaction binding the contract method 0xdd6388fb.
//
// Solidity: function removeHost(addr address) returns()
func (_TicketSale *TicketSaleSession) RemoveHost(addr common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.RemoveHost(&_TicketSale.TransactOpts, addr)
}

// RemoveHost is a paid mutator transaction binding the contract method 0xdd6388fb.
//
// Solidity: function removeHost(addr address) returns()
func (_TicketSale *TicketSaleTransactorSession) RemoveHost(addr common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.RemoveHost(&_TicketSale.TransactOpts, addr)
}

// RequestTrading is a paid mutator transaction binding the contract method 0x02dce02c.
//
// Solidity: function requestTrading(_ticketId uint256, _value uint256) returns()
func (_TicketSale *TicketSaleTransactor) RequestTrading(opts *bind.TransactOpts, _ticketId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "requestTrading", _ticketId, _value)
}

// RequestTrading is a paid mutator transaction binding the contract method 0x02dce02c.
//
// Solidity: function requestTrading(_ticketId uint256, _value uint256) returns()
func (_TicketSale *TicketSaleSession) RequestTrading(_ticketId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.RequestTrading(&_TicketSale.TransactOpts, _ticketId, _value)
}

// RequestTrading is a paid mutator transaction binding the contract method 0x02dce02c.
//
// Solidity: function requestTrading(_ticketId uint256, _value uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) RequestTrading(_ticketId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.RequestTrading(&_TicketSale.TransactOpts, _ticketId, _value)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_TicketSale *TicketSaleTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_TicketSale *TicketSaleSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TicketSale.Contract.SafeTransferFrom(&_TicketSale.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_TicketSale *TicketSaleTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _TicketSale.Contract.SafeTransferFrom(&_TicketSale.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_TicketSale *TicketSaleTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_TicketSale *TicketSaleSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _TicketSale.Contract.SetApprovalForAll(&_TicketSale.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_TicketSale *TicketSaleTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _TicketSale.Contract.SetApprovalForAll(&_TicketSale.TransactOpts, _to, _approved)
}

// SetHost is a paid mutator transaction binding the contract method 0xc85e0be2.
//
// Solidity: function setHost(host address) returns()
func (_TicketSale *TicketSaleTransactor) SetHost(opts *bind.TransactOpts, host common.Address) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "setHost", host)
}

// SetHost is a paid mutator transaction binding the contract method 0xc85e0be2.
//
// Solidity: function setHost(host address) returns()
func (_TicketSale *TicketSaleSession) SetHost(host common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.SetHost(&_TicketSale.TransactOpts, host)
}

// SetHost is a paid mutator transaction binding the contract method 0xc85e0be2.
//
// Solidity: function setHost(host address) returns()
func (_TicketSale *TicketSaleTransactorSession) SetHost(host common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.SetHost(&_TicketSale.TransactOpts, host)
}

// SetLastLose is a paid mutator transaction binding the contract method 0xf1521f91.
//
// Solidity: function setLastLose(_lastLose address) returns()
func (_TicketSale *TicketSaleTransactor) SetLastLose(opts *bind.TransactOpts, _lastLose common.Address) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "setLastLose", _lastLose)
}

// SetLastLose is a paid mutator transaction binding the contract method 0xf1521f91.
//
// Solidity: function setLastLose(_lastLose address) returns()
func (_TicketSale *TicketSaleSession) SetLastLose(_lastLose common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.SetLastLose(&_TicketSale.TransactOpts, _lastLose)
}

// SetLastLose is a paid mutator transaction binding the contract method 0xf1521f91.
//
// Solidity: function setLastLose(_lastLose address) returns()
func (_TicketSale *TicketSaleTransactorSession) SetLastLose(_lastLose common.Address) (*types.Transaction, error) {
	return _TicketSale.Contract.SetLastLose(&_TicketSale.TransactOpts, _lastLose)
}

// SetLimit is a paid mutator transaction binding the contract method 0xff3f19b5.
//
// Solidity: function setLimit(_limit uint8) returns()
func (_TicketSale *TicketSaleTransactor) SetLimit(opts *bind.TransactOpts, _limit uint8) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "setLimit", _limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0xff3f19b5.
//
// Solidity: function setLimit(_limit uint8) returns()
func (_TicketSale *TicketSaleSession) SetLimit(_limit uint8) (*types.Transaction, error) {
	return _TicketSale.Contract.SetLimit(&_TicketSale.TransactOpts, _limit)
}

// SetLimit is a paid mutator transaction binding the contract method 0xff3f19b5.
//
// Solidity: function setLimit(_limit uint8) returns()
func (_TicketSale *TicketSaleTransactorSession) SetLimit(_limit uint8) (*types.Transaction, error) {
	return _TicketSale.Contract.SetLimit(&_TicketSale.TransactOpts, _limit)
}

// SetMaxAttendees is a paid mutator transaction binding the contract method 0x08c0204b.
//
// Solidity: function setMaxAttendees(_maxAttendees uint256) returns()
func (_TicketSale *TicketSaleTransactor) SetMaxAttendees(opts *bind.TransactOpts, _maxAttendees *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "setMaxAttendees", _maxAttendees)
}

// SetMaxAttendees is a paid mutator transaction binding the contract method 0x08c0204b.
//
// Solidity: function setMaxAttendees(_maxAttendees uint256) returns()
func (_TicketSale *TicketSaleSession) SetMaxAttendees(_maxAttendees *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.SetMaxAttendees(&_TicketSale.TransactOpts, _maxAttendees)
}

// SetMaxAttendees is a paid mutator transaction binding the contract method 0x08c0204b.
//
// Solidity: function setMaxAttendees(_maxAttendees uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) SetMaxAttendees(_maxAttendees *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.SetMaxAttendees(&_TicketSale.TransactOpts, _maxAttendees)
}

// SetNextPrice is a paid mutator transaction binding the contract method 0x23037a85.
//
// Solidity: function setNextPrice(_nextPrice uint256) returns()
func (_TicketSale *TicketSaleTransactor) SetNextPrice(opts *bind.TransactOpts, _nextPrice *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "setNextPrice", _nextPrice)
}

// SetNextPrice is a paid mutator transaction binding the contract method 0x23037a85.
//
// Solidity: function setNextPrice(_nextPrice uint256) returns()
func (_TicketSale *TicketSaleSession) SetNextPrice(_nextPrice *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.SetNextPrice(&_TicketSale.TransactOpts, _nextPrice)
}

// SetNextPrice is a paid mutator transaction binding the contract method 0x23037a85.
//
// Solidity: function setNextPrice(_nextPrice uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) SetNextPrice(_nextPrice *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.SetNextPrice(&_TicketSale.TransactOpts, _nextPrice)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(_price uint256) returns()
func (_TicketSale *TicketSaleTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(_price uint256) returns()
func (_TicketSale *TicketSaleSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.SetPrice(&_TicketSale.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(_price uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.SetPrice(&_TicketSale.TransactOpts, _price)
}

// SetTradeFee is a paid mutator transaction binding the contract method 0x9fdccfb8.
//
// Solidity: function setTradeFee(_fee uint16) returns()
func (_TicketSale *TicketSaleTransactor) SetTradeFee(opts *bind.TransactOpts, _fee uint16) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "setTradeFee", _fee)
}

// SetTradeFee is a paid mutator transaction binding the contract method 0x9fdccfb8.
//
// Solidity: function setTradeFee(_fee uint16) returns()
func (_TicketSale *TicketSaleSession) SetTradeFee(_fee uint16) (*types.Transaction, error) {
	return _TicketSale.Contract.SetTradeFee(&_TicketSale.TransactOpts, _fee)
}

// SetTradeFee is a paid mutator transaction binding the contract method 0x9fdccfb8.
//
// Solidity: function setTradeFee(_fee uint16) returns()
func (_TicketSale *TicketSaleTransactorSession) SetTradeFee(_fee uint16) (*types.Transaction, error) {
	return _TicketSale.Contract.SetTradeFee(&_TicketSale.TransactOpts, _fee)
}

// SetUsedTickets is a paid mutator transaction binding the contract method 0x03a865ee.
//
// Solidity: function setUsedTickets(_ticketIds uint256[]) returns()
func (_TicketSale *TicketSaleTransactor) SetUsedTickets(opts *bind.TransactOpts, _ticketIds []*big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "setUsedTickets", _ticketIds)
}

// SetUsedTickets is a paid mutator transaction binding the contract method 0x03a865ee.
//
// Solidity: function setUsedTickets(_ticketIds uint256[]) returns()
func (_TicketSale *TicketSaleSession) SetUsedTickets(_ticketIds []*big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.SetUsedTickets(&_TicketSale.TransactOpts, _ticketIds)
}

// SetUsedTickets is a paid mutator transaction binding the contract method 0x03a865ee.
//
// Solidity: function setUsedTickets(_ticketIds uint256[]) returns()
func (_TicketSale *TicketSaleTransactorSession) SetUsedTickets(_ticketIds []*big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.SetUsedTickets(&_TicketSale.TransactOpts, _ticketIds)
}

// Trade is a paid mutator transaction binding the contract method 0xdf1dd826.
//
// Solidity: function trade(_tradeId uint256) returns()
func (_TicketSale *TicketSaleTransactor) Trade(opts *bind.TransactOpts, _tradeId *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "trade", _tradeId)
}

// Trade is a paid mutator transaction binding the contract method 0xdf1dd826.
//
// Solidity: function trade(_tradeId uint256) returns()
func (_TicketSale *TicketSaleSession) Trade(_tradeId *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.Trade(&_TicketSale.TransactOpts, _tradeId)
}

// Trade is a paid mutator transaction binding the contract method 0xdf1dd826.
//
// Solidity: function trade(_tradeId uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) Trade(_tradeId *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.Trade(&_TicketSale.TransactOpts, _tradeId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_TicketSale *TicketSaleTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_TicketSale *TicketSaleSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.TransferFrom(&_TicketSale.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256) returns()
func (_TicketSale *TicketSaleTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _TicketSale.Contract.TransferFrom(&_TicketSale.TransactOpts, _from, _to, _tokenId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TicketSale *TicketSaleTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TicketSale *TicketSaleSession) Withdraw() (*types.Transaction, error) {
	return _TicketSale.Contract.Withdraw(&_TicketSale.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TicketSale *TicketSaleTransactorSession) Withdraw() (*types.Transaction, error) {
	return _TicketSale.Contract.Withdraw(&_TicketSale.TransactOpts)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_TicketSale *TicketSaleTransactor) WithdrawFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TicketSale.contract.Transact(opts, "withdrawFee")
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_TicketSale *TicketSaleSession) WithdrawFee() (*types.Transaction, error) {
	return _TicketSale.Contract.WithdrawFee(&_TicketSale.TransactOpts)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_TicketSale *TicketSaleTransactorSession) WithdrawFee() (*types.Transaction, error) {
	return _TicketSale.Contract.WithdrawFee(&_TicketSale.TransactOpts)
}

// TicketSaleApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TicketSale contract.
type TicketSaleApprovalIterator struct {
	Event *TicketSaleApproval // Event containing the contract specifics and raw log

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
func (it *TicketSaleApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TicketSaleApproval)
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
		it.Event = new(TicketSaleApproval)
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
func (it *TicketSaleApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TicketSaleApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TicketSaleApproval represents a Approval event raised by the TicketSale contract.
type TicketSaleApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_TicketSale *TicketSaleFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*TicketSaleApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _TicketSale.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TicketSaleApprovalIterator{contract: _TicketSale.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_TicketSale *TicketSaleFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TicketSaleApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _TicketSale.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TicketSaleApproval)
				if err := _TicketSale.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TicketSaleApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TicketSale contract.
type TicketSaleApprovalForAllIterator struct {
	Event *TicketSaleApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TicketSaleApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TicketSaleApprovalForAll)
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
		it.Event = new(TicketSaleApprovalForAll)
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
func (it *TicketSaleApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TicketSaleApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TicketSaleApprovalForAll represents a ApprovalForAll event raised by the TicketSale contract.
type TicketSaleApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_TicketSale *TicketSaleFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*TicketSaleApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _TicketSale.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &TicketSaleApprovalForAllIterator{contract: _TicketSale.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_TicketSale *TicketSaleFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TicketSaleApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _TicketSale.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TicketSaleApprovalForAll)
				if err := _TicketSale.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// TicketSaleTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TicketSale contract.
type TicketSaleTransferIterator struct {
	Event *TicketSaleTransfer // Event containing the contract specifics and raw log

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
func (it *TicketSaleTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TicketSaleTransfer)
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
		it.Event = new(TicketSaleTransfer)
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
func (it *TicketSaleTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TicketSaleTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TicketSaleTransfer represents a Transfer event raised by the TicketSale contract.
type TicketSaleTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_TicketSale *TicketSaleFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*TicketSaleTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _TicketSale.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TicketSaleTransferIterator{contract: _TicketSale.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_TicketSale *TicketSaleFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TicketSaleTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _TicketSale.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TicketSaleTransfer)
				if err := _TicketSale.contract.UnpackLog(event, "Transfer", log); err != nil {
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
