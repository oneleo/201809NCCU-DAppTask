// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BankABI is the input ABI used to generate the binding from.
const BankABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"etherValue\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCoinBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBankBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"coinValue\",\"type\":\"uint256\"}],\"name\":\"transferCoin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinValue\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"etherValue\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinValue\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TransferEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MintEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BuyCoinEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TransferCoinEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TransferOwnerEvent\",\"type\":\"event\"}]"

// BankBin is the compiled bytecode used for deploying new contracts.
//const BankBin = `0x608060405260008054600160a060020a03191633179055610a22806100256000396000f3fe6080604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632e1a7d4d81146100b357806341c0e1b5146100df5780634fb2e45d146100f457806356fbd78f146101275780637b83b50b1461014e578063893d20e8146101635780638dde60fa14610194578063a0712d68146101cd578063a9059cbb146101f7578063d0e30db014610230578063d96a094a14610238575b600080fd5b3480156100bf57600080fd5b506100dd600480360360208110156100d657600080fd5b5035610262565b005b3480156100eb57600080fd5b506100dd610381565b34801561010057600080fd5b506100dd6004803603602081101561011757600080fd5b5035600160a060020a03166103e6565b34801561013357600080fd5b5061013c6104d1565b60408051918252519081900360200190f35b34801561015a57600080fd5b5061013c6104e4565b34801561016f57600080fd5b506101786104f7565b60408051600160a060020a039092168252519081900360200190f35b3480156101a057600080fd5b506100dd600480360360408110156101b757600080fd5b50600160a060020a038135169060200135610506565b3480156101d957600080fd5b506100dd600480360360208110156101f057600080fd5b5035610610565b34801561020357600080fd5b506100dd6004803603604081101561021a57600080fd5b50600160a060020a0381351690602001356106d6565b6100dd6107e0565b34801561024457600080fd5b506100dd6004803603602081101561025b57600080fd5b5035610838565b33600090815260016020526040902054670de0b6b3a76400008202908111156102fb576040805160e560020a62461bcd02815260206004820152602360248201527f596f75722065746865722062616c616e63657320617265206e6f7420656e6f7560448201527f67682e0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b604051339082156108fc029083906000818181858888f19350505050158015610328573d6000803e3d6000fd5b5033600081815260016020908152604091829020805485900390558151858152429181019190915281517f5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab929181900390910190a25050565b600054600160a060020a031633146103e3576040805160e560020a62461bcd02815260206004820152601260248201527f596f7520617265206e6f74206f776e65722e0000000000000000000000000000604482015290519081900360640190fd5b33ff5b600054600160a060020a03163314610448576040805160e560020a62461bcd02815260206004820152601260248201527f596f7520617265206e6f74206f776e65722e0000000000000000000000000000604482015290519081900360640190fd5b600160a060020a038116151561045d57600080fd5b600054604080514281529051600160a060020a038085169316917f587a4fcff87b7be11c779eb502f8b2584f996387d8b8cda0e5113fef424f7316919081900360200190a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b3360009081526002602052604090205490565b3360009081526001602052604090205490565b600054600160a060020a031690565b33600090815260026020526040902054670de0b6b3a764000082029081111561059f576040805160e560020a62461bcd02815260206004820152602760248201527f596f75722062616e6b20636f696e2062616c616e63657320617265206e6f742060448201527f656e6f7567682e00000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b33600081815260026020908152604080832080548690039055600160a060020a038716808452928190208054860190558051868152429281019290925280519293927f941d755df54ad0234b406209d0c923107cabf6d4f1ce335b8ae5d89d6a28c2d29281900390910190a3505050565b600054600160a060020a03163314610672576040805160e560020a62461bcd02815260206004820152601260248201527f596f7520617265206e6f74206f776e65722e0000000000000000000000000000604482015290519081900360640190fd5b336000818152600260209081526040918290208054670de0b6b3a764000086029081019091558251858152429281019290925282519093927f8069ef4945469d029cc32e222031bccdc99b2eaaf4ee374cd268012f7ddee907928290030190a25050565b33600090815260016020526040902054670de0b6b3a764000082029081111561076f576040805160e560020a62461bcd02815260206004820152602360248201527f596f75722065746865722062616c616e63657320617265206e6f7420656e6f7560448201527f67682e0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b33600081815260016020908152604080832080548690039055600160a060020a038716808452928190208054860190558051868152429281019290925280519293927fbabc8cd3bd6701ee99131f374fd2ab4ea66f48dc4e4182ed78fecb0502e44dd69281900390910190a3505050565b336000818152600160209081526040918290208054349081019091558251908152429181019190915281517fad40ae5dc69974ba932d08b0a608e89109412d41d04850f5196f144875ae2660929181900390910190a2565b60008054600160a060020a0316815260026020526040902054670de0b6b3a76400008202908111156108da576040805160e560020a62461bcd02815260206004820152602a60248201527f4f776e657227732062616e6b20636f696e2062616c616e63657320617265206e60448201527f6f7420656e6f7567682e00000000000000000000000000000000000000000000606482015290519081900360840190fd5b33600090815260016020526040902054811115610967576040805160e560020a62461bcd02815260206004820152602360248201527f596f75722065746865722062616c616e63657320617265206e6f7420656e6f7560448201527f67682e0000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000818152600160209081526040808320805486900390558254600160a060020a03908116845281842080548701905584845260028352818420805487019055835416835291829020805485900390558151858152429181019190915281517f4c5ad1aea676c1e1613de5416105424342b84655de046903409dea58418bedff929181900390910190a2505056fea165627a7a72305820063f5a6735fc1ccd7c898926ee5036708bc3e55ec9df1b24a7554dd1ede3ee0a0029`
var BankBin string = ""

func SetBankBin(solBin string) {
	BankBin = solBin
}

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := abi.JSON(strings.NewReader(BankABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if BankBin == "" {
		return common.Address{}, nil, nil, errors.New("BankBin is empty")
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// GetBankBalance is a free data retrieval call binding the contract method 0x7b83b50b.
//
// Solidity: function getBankBalance() constant returns(uint256)
func (_Bank *BankCaller) GetBankBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bank.contract.Call(opts, out, "getBankBalance")
	return *ret0, err
}

// GetBankBalance is a free data retrieval call binding the contract method 0x7b83b50b.
//
// Solidity: function getBankBalance() constant returns(uint256)
func (_Bank *BankSession) GetBankBalance() (*big.Int, error) {
	return _Bank.Contract.GetBankBalance(&_Bank.CallOpts)
}

// GetBankBalance is a free data retrieval call binding the contract method 0x7b83b50b.
//
// Solidity: function getBankBalance() constant returns(uint256)
func (_Bank *BankCallerSession) GetBankBalance() (*big.Int, error) {
	return _Bank.Contract.GetBankBalance(&_Bank.CallOpts)
}

// GetCoinBalance is a free data retrieval call binding the contract method 0x56fbd78f.
//
// Solidity: function getCoinBalance() constant returns(uint256)
func (_Bank *BankCaller) GetCoinBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bank.contract.Call(opts, out, "getCoinBalance")
	return *ret0, err
}

// GetCoinBalance is a free data retrieval call binding the contract method 0x56fbd78f.
//
// Solidity: function getCoinBalance() constant returns(uint256)
func (_Bank *BankSession) GetCoinBalance() (*big.Int, error) {
	return _Bank.Contract.GetCoinBalance(&_Bank.CallOpts)
}

// GetCoinBalance is a free data retrieval call binding the contract method 0x56fbd78f.
//
// Solidity: function getCoinBalance() constant returns(uint256)
func (_Bank *BankCallerSession) GetCoinBalance() (*big.Int, error) {
	return _Bank.Contract.GetCoinBalance(&_Bank.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Bank *BankCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bank.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Bank *BankSession) GetOwner() (common.Address, error) {
	return _Bank.Contract.GetOwner(&_Bank.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_Bank *BankCallerSession) GetOwner() (common.Address, error) {
	return _Bank.Contract.GetOwner(&_Bank.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(coinValue uint256) returns()
func (_Bank *BankTransactor) Buy(opts *bind.TransactOpts, coinValue *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "buy", coinValue)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(coinValue uint256) returns()
func (_Bank *BankSession) Buy(coinValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Buy(&_Bank.TransactOpts, coinValue)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(coinValue uint256) returns()
func (_Bank *BankTransactorSession) Buy(coinValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Buy(&_Bank.TransactOpts, coinValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Bank *BankSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Bank *BankTransactorSession) Deposit() (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Bank *BankTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Bank *BankSession) Kill() (*types.Transaction, error) {
	return _Bank.Contract.Kill(&_Bank.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_Bank *BankTransactorSession) Kill() (*types.Transaction, error) {
	return _Bank.Contract.Kill(&_Bank.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(coinValue uint256) returns()
func (_Bank *BankTransactor) Mint(opts *bind.TransactOpts, coinValue *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "mint", coinValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(coinValue uint256) returns()
func (_Bank *BankSession) Mint(coinValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Mint(&_Bank.TransactOpts, coinValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(coinValue uint256) returns()
func (_Bank *BankTransactorSession) Mint(coinValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Mint(&_Bank.TransactOpts, coinValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, etherValue uint256) returns()
func (_Bank *BankTransactor) Transfer(opts *bind.TransactOpts, to common.Address, etherValue *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "transfer", to, etherValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, etherValue uint256) returns()
func (_Bank *BankSession) Transfer(to common.Address, etherValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Transfer(&_Bank.TransactOpts, to, etherValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, etherValue uint256) returns()
func (_Bank *BankTransactorSession) Transfer(to common.Address, etherValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Transfer(&_Bank.TransactOpts, to, etherValue)
}

// TransferCoin is a paid mutator transaction binding the contract method 0x8dde60fa.
//
// Solidity: function transferCoin(to address, coinValue uint256) returns()
func (_Bank *BankTransactor) TransferCoin(opts *bind.TransactOpts, to common.Address, coinValue *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "transferCoin", to, coinValue)
}

// TransferCoin is a paid mutator transaction binding the contract method 0x8dde60fa.
//
// Solidity: function transferCoin(to address, coinValue uint256) returns()
func (_Bank *BankSession) TransferCoin(to common.Address, coinValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.TransferCoin(&_Bank.TransactOpts, to, coinValue)
}

// TransferCoin is a paid mutator transaction binding the contract method 0x8dde60fa.
//
// Solidity: function transferCoin(to address, coinValue uint256) returns()
func (_Bank *BankTransactorSession) TransferCoin(to common.Address, coinValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.TransferCoin(&_Bank.TransactOpts, to, coinValue)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Bank *BankTransactor) TransferOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "transferOwner", newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Bank *BankSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Bank.Contract.TransferOwner(&_Bank.TransactOpts, newOwner)
}

// TransferOwner is a paid mutator transaction binding the contract method 0x4fb2e45d.
//
// Solidity: function transferOwner(newOwner address) returns()
func (_Bank *BankTransactorSession) TransferOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Bank.Contract.TransferOwner(&_Bank.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(etherValue uint256) returns()
func (_Bank *BankTransactor) Withdraw(opts *bind.TransactOpts, etherValue *big.Int) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "withdraw", etherValue)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(etherValue uint256) returns()
func (_Bank *BankSession) Withdraw(etherValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts, etherValue)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(etherValue uint256) returns()
func (_Bank *BankTransactorSession) Withdraw(etherValue *big.Int) (*types.Transaction, error) {
	return _Bank.Contract.Withdraw(&_Bank.TransactOpts, etherValue)
}

// BankBuyCoinEventIterator is returned from FilterBuyCoinEvent and is used to iterate over the raw logs and unpacked data for BuyCoinEvent events raised by the Bank contract.
type BankBuyCoinEventIterator struct {
	Event *BankBuyCoinEvent // Event containing the contract specifics and raw log

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
func (it *BankBuyCoinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankBuyCoinEvent)
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
		it.Event = new(BankBuyCoinEvent)
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
func (it *BankBuyCoinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankBuyCoinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankBuyCoinEvent represents a BuyCoinEvent event raised by the Bank contract.
type BankBuyCoinEvent struct {
	From      common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBuyCoinEvent is a free log retrieval operation binding the contract event 0x4c5ad1aea676c1e1613de5416105424342b84655de046903409dea58418bedff.
//
// Solidity: e BuyCoinEvent(from indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) FilterBuyCoinEvent(opts *bind.FilterOpts, from []common.Address) (*BankBuyCoinEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bank.contract.FilterLogs(opts, "BuyCoinEvent", fromRule)
	if err != nil {
		return nil, err
	}
	return &BankBuyCoinEventIterator{contract: _Bank.contract, event: "BuyCoinEvent", logs: logs, sub: sub}, nil
}

// WatchBuyCoinEvent is a free log subscription operation binding the contract event 0x4c5ad1aea676c1e1613de5416105424342b84655de046903409dea58418bedff.
//
// Solidity: e BuyCoinEvent(from indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) WatchBuyCoinEvent(opts *bind.WatchOpts, sink chan<- *BankBuyCoinEvent, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bank.contract.WatchLogs(opts, "BuyCoinEvent", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankBuyCoinEvent)
				if err := _Bank.contract.UnpackLog(event, "BuyCoinEvent", log); err != nil {
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

// BankDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the Bank contract.
type BankDepositEventIterator struct {
	Event *BankDepositEvent // Event containing the contract specifics and raw log

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
func (it *BankDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankDepositEvent)
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
		it.Event = new(BankDepositEvent)
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
func (it *BankDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankDepositEvent represents a DepositEvent event raised by the Bank contract.
type BankDepositEvent struct {
	From      common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0xad40ae5dc69974ba932d08b0a608e89109412d41d04850f5196f144875ae2660.
//
// Solidity: e DepositEvent(from indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) FilterDepositEvent(opts *bind.FilterOpts, from []common.Address) (*BankDepositEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bank.contract.FilterLogs(opts, "DepositEvent", fromRule)
	if err != nil {
		return nil, err
	}
	return &BankDepositEventIterator{contract: _Bank.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0xad40ae5dc69974ba932d08b0a608e89109412d41d04850f5196f144875ae2660.
//
// Solidity: e DepositEvent(from indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *BankDepositEvent, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bank.contract.WatchLogs(opts, "DepositEvent", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankDepositEvent)
				if err := _Bank.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// BankMintEventIterator is returned from FilterMintEvent and is used to iterate over the raw logs and unpacked data for MintEvent events raised by the Bank contract.
type BankMintEventIterator struct {
	Event *BankMintEvent // Event containing the contract specifics and raw log

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
func (it *BankMintEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankMintEvent)
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
		it.Event = new(BankMintEvent)
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
func (it *BankMintEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankMintEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankMintEvent represents a MintEvent event raised by the Bank contract.
type BankMintEvent struct {
	From      common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintEvent is a free log retrieval operation binding the contract event 0x8069ef4945469d029cc32e222031bccdc99b2eaaf4ee374cd268012f7ddee907.
//
// Solidity: e MintEvent(from indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) FilterMintEvent(opts *bind.FilterOpts, from []common.Address) (*BankMintEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bank.contract.FilterLogs(opts, "MintEvent", fromRule)
	if err != nil {
		return nil, err
	}
	return &BankMintEventIterator{contract: _Bank.contract, event: "MintEvent", logs: logs, sub: sub}, nil
}

// WatchMintEvent is a free log subscription operation binding the contract event 0x8069ef4945469d029cc32e222031bccdc99b2eaaf4ee374cd268012f7ddee907.
//
// Solidity: e MintEvent(from indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) WatchMintEvent(opts *bind.WatchOpts, sink chan<- *BankMintEvent, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bank.contract.WatchLogs(opts, "MintEvent", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankMintEvent)
				if err := _Bank.contract.UnpackLog(event, "MintEvent", log); err != nil {
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

// BankTransferCoinEventIterator is returned from FilterTransferCoinEvent and is used to iterate over the raw logs and unpacked data for TransferCoinEvent events raised by the Bank contract.
type BankTransferCoinEventIterator struct {
	Event *BankTransferCoinEvent // Event containing the contract specifics and raw log

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
func (it *BankTransferCoinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankTransferCoinEvent)
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
		it.Event = new(BankTransferCoinEvent)
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
func (it *BankTransferCoinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankTransferCoinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankTransferCoinEvent represents a TransferCoinEvent event raised by the Bank contract.
type BankTransferCoinEvent struct {
	From      common.Address
	To        common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferCoinEvent is a free log retrieval operation binding the contract event 0x941d755df54ad0234b406209d0c923107cabf6d4f1ce335b8ae5d89d6a28c2d2.
//
// Solidity: e TransferCoinEvent(from indexed address, to indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) FilterTransferCoinEvent(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BankTransferCoinEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bank.contract.FilterLogs(opts, "TransferCoinEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BankTransferCoinEventIterator{contract: _Bank.contract, event: "TransferCoinEvent", logs: logs, sub: sub}, nil
}

// WatchTransferCoinEvent is a free log subscription operation binding the contract event 0x941d755df54ad0234b406209d0c923107cabf6d4f1ce335b8ae5d89d6a28c2d2.
//
// Solidity: e TransferCoinEvent(from indexed address, to indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) WatchTransferCoinEvent(opts *bind.WatchOpts, sink chan<- *BankTransferCoinEvent, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bank.contract.WatchLogs(opts, "TransferCoinEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankTransferCoinEvent)
				if err := _Bank.contract.UnpackLog(event, "TransferCoinEvent", log); err != nil {
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

// BankTransferEventIterator is returned from FilterTransferEvent and is used to iterate over the raw logs and unpacked data for TransferEvent events raised by the Bank contract.
type BankTransferEventIterator struct {
	Event *BankTransferEvent // Event containing the contract specifics and raw log

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
func (it *BankTransferEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankTransferEvent)
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
		it.Event = new(BankTransferEvent)
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
func (it *BankTransferEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankTransferEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankTransferEvent represents a TransferEvent event raised by the Bank contract.
type BankTransferEvent struct {
	From      common.Address
	To        common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferEvent is a free log retrieval operation binding the contract event 0xbabc8cd3bd6701ee99131f374fd2ab4ea66f48dc4e4182ed78fecb0502e44dd6.
//
// Solidity: e TransferEvent(from indexed address, to indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) FilterTransferEvent(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BankTransferEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bank.contract.FilterLogs(opts, "TransferEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BankTransferEventIterator{contract: _Bank.contract, event: "TransferEvent", logs: logs, sub: sub}, nil
}

// WatchTransferEvent is a free log subscription operation binding the contract event 0xbabc8cd3bd6701ee99131f374fd2ab4ea66f48dc4e4182ed78fecb0502e44dd6.
//
// Solidity: e TransferEvent(from indexed address, to indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) WatchTransferEvent(opts *bind.WatchOpts, sink chan<- *BankTransferEvent, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bank.contract.WatchLogs(opts, "TransferEvent", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankTransferEvent)
				if err := _Bank.contract.UnpackLog(event, "TransferEvent", log); err != nil {
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

// BankTransferOwnerEventIterator is returned from FilterTransferOwnerEvent and is used to iterate over the raw logs and unpacked data for TransferOwnerEvent events raised by the Bank contract.
type BankTransferOwnerEventIterator struct {
	Event *BankTransferOwnerEvent // Event containing the contract specifics and raw log

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
func (it *BankTransferOwnerEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankTransferOwnerEvent)
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
		it.Event = new(BankTransferOwnerEvent)
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
func (it *BankTransferOwnerEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankTransferOwnerEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankTransferOwnerEvent represents a TransferOwnerEvent event raised by the Bank contract.
type BankTransferOwnerEvent struct {
	OldOwner  common.Address
	NewOwner  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferOwnerEvent is a free log retrieval operation binding the contract event 0x587a4fcff87b7be11c779eb502f8b2584f996387d8b8cda0e5113fef424f7316.
//
// Solidity: e TransferOwnerEvent(oldOwner indexed address, newOwner indexed address, timestamp uint256)
func (_Bank *BankFilterer) FilterTransferOwnerEvent(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*BankTransferOwnerEventIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bank.contract.FilterLogs(opts, "TransferOwnerEvent", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BankTransferOwnerEventIterator{contract: _Bank.contract, event: "TransferOwnerEvent", logs: logs, sub: sub}, nil
}

// WatchTransferOwnerEvent is a free log subscription operation binding the contract event 0x587a4fcff87b7be11c779eb502f8b2584f996387d8b8cda0e5113fef424f7316.
//
// Solidity: e TransferOwnerEvent(oldOwner indexed address, newOwner indexed address, timestamp uint256)
func (_Bank *BankFilterer) WatchTransferOwnerEvent(opts *bind.WatchOpts, sink chan<- *BankTransferOwnerEvent, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bank.contract.WatchLogs(opts, "TransferOwnerEvent", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankTransferOwnerEvent)
				if err := _Bank.contract.UnpackLog(event, "TransferOwnerEvent", log); err != nil {
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

// BankWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the Bank contract.
type BankWithdrawEventIterator struct {
	Event *BankWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *BankWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankWithdrawEvent)
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
		it.Event = new(BankWithdrawEvent)
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
func (it *BankWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankWithdrawEvent represents a WithdrawEvent event raised by the Bank contract.
type BankWithdrawEvent struct {
	From      common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: e WithdrawEvent(from indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) FilterWithdrawEvent(opts *bind.FilterOpts, from []common.Address) (*BankWithdrawEventIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bank.contract.FilterLogs(opts, "WithdrawEvent", fromRule)
	if err != nil {
		return nil, err
	}
	return &BankWithdrawEventIterator{contract: _Bank.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: e WithdrawEvent(from indexed address, value uint256, timestamp uint256)
func (_Bank *BankFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *BankWithdrawEvent, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Bank.contract.WatchLogs(opts, "WithdrawEvent", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankWithdrawEvent)
				if err := _Bank.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
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
