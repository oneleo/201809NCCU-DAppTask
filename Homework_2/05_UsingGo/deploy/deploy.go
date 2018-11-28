package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/oneleo/technical-analysis/file"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/oneleo/201809-Ethereum-DApp-Development-Homework/Homework_2/05_UsingGo/contract"

	"log"
)

var (
	debug    bool   = false
	gasLimit uint64 = 1000000
	// Keystore 1 mnemonic:
	// people stand vicious engage cruel solve error power camera cousin maple use
	// Keystore 2 private key:
	// 0xad99cb7d0f056d79ae84779e9649555a2abb48e43f9cbf1287ba03d93c2d33c6
	keystorFile1 string = "../keystore1"
	keystorFile2 string = "../keystore2"
	keystorAuth  string = "nccu"

	// 定義要連線的以太坊網路。
	//url string = "https://rinkeby.infura.io"
	//url string = "wss://rinkeby.infura.io/ws"
	url string = "http://localhost:8545"

	// 定義部署後的合約地址放置位置。
	addressFile string = "../address.txt"
)

//
// Reference:
// https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts
//
// > cd /d "%GOPATH%\src\github.com\oneleo\201809-Ethereum-DApp-Development-Homework\Homework_2\05_UsingGo"
//
// > md ".\contract"
// > abigen --sol ./contract/bank.sol --pkg contract --out ./contract/bank.go
//
// > code "./contract/bank.go"
//
// 修改第 34 列，把原來的內容：
// -------------------- ----- 內容 ----- --------------------
// var BankBin string = `***`
// -------------------- ----- 內容 ----- --------------------
// 改成
// -------------------- ----- 內容 ----- --------------------
// var BankBin string = ""
// func SetBankBin(solBin string) {
// 	BankBin = solBin
// }
// -------------------- ----- 內容 ----- --------------------
//
// 在第 42 列增加內容：
// -------------------- ----- 內容 ----- --------------------
// 	if BankBin == "" {
// 		return common.Address{}, nil, nil, errors.New("BankBin is empty")
// 	}
// -------------------- ----- 內容 ----- --------------------
//

func main() {
	//compile, err := compiler.CompileSolidity("solc.exe", "./contract/bank.sol")
	compile, err := compiler.CompileSolidityString("solc.exe", BankContract)
	if err != nil {
		log.Fatal("Compile solidity error occur: ", err, "\n")
	}
	if debug == true {
		fmt.Print(compile[CompilerName], "\n")
	}

	// 打開賬戶私鑰文件
	keyJson1, err := ioutil.ReadFile(keystorFile1)
	if err != nil {
		log.Fatal("kKey json read error: ", err, "\n")
	}

	// 解析私鑰文件
	keyWrapper1, err := keystore.DecryptKey(keyJson1, keystorAuth)
	if err != nil {
		log.Fatal("Key decrypt error: ", err, "\n")
	}
	prvKey1 := keyWrapper1.PrivateKey

	if debug == true {
		prvKeyStr1 := hexutil.Encode(crypto.FromECDSA(prvKey1))
		fmt.Print("Address1:\t", keyWrapper1.Address.Hex(), "\nPrivate key:\t", prvKeyStr1, "\n")
	}

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := (*prvKey1).Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(prvKey1)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	auth.GasLimit = uint64(gasLimit) // in units
	auth.GasPrice = gasPrice

	contract.SetBankBin((*compile[CompilerName]).Code)
	address, tx, instance, err := contract.DeployBank(auth, client)
	//store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Deploy from:\t\t", keyWrapper1.Address.Hex(), "\n")
	fmt.Print("Contract address:\t", address.Hex(), "\n")
	fmt.Print("Transaction hash:\t", tx.Hash().Hex(), "\n")
	fmt.Print("Instance info:\t\t", instance, "\n")

	stringsTemp := []string{address.Hex()}
	file.StringsToFile(addressFile, stringsTemp)
}
