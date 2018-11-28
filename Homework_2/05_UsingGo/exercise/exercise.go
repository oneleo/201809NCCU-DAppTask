package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"

	"github.com/oneleo/201809-Ethereum-DApp-Development-Homework/Homework_2/05_UsingGo/contract"
	"github.com/oneleo/201809-Ethereum-DApp-Development-Homework/Homework_2/05_UsingGo/file"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	debug    bool   = false
	gasLimit uint64 = 300000
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

func main() {

	// 打開賬戶私鑰文件。
	keyJson1, err := ioutil.ReadFile(keystorFile1)
	if err != nil {
		log.Fatal("Key json read error: ", err, "\n")
	}
	keyJson2, err := ioutil.ReadFile(keystorFile2)
	if err != nil {
		log.Fatal("Key json read error: ", err, "\n")
	}

	// 解析私鑰文件。
	keyWrapper1, err := keystore.DecryptKey(keyJson1, keystorAuth)
	if err != nil {
		log.Fatal("Key decrypt error: ", err, "\n")
	}
	keyWrapper2, err := keystore.DecryptKey(keyJson2, keystorAuth)
	if err != nil {
		log.Fatal("Key decrypt error: ", err, "\n")
	}

	// 取得 Private Key ECDSA Struct 結構。
	privateKeyECDSA1 := keyWrapper1.PrivateKey
	privateKeyECDSA2 := keyWrapper2.PrivateKey

	if debug == true {
		prvKeyStr1 := hexutil.Encode(crypto.FromECDSA(privateKeyECDSA1))
		fmt.Print("Address1:\t", keyWrapper1.Address.Hex(), "\nPrivate key:\t", prvKeyStr1, "\n")
		prvKeyStr2 := hexutil.Encode(crypto.FromECDSA(privateKeyECDSA2))
		fmt.Print("Address2:\t", keyWrapper2.Address.Hex(), "\nPrivate key:\t", prvKeyStr2, "\n")
	}

	// 連線至指定的以太坊網路。
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	// 從指定的檔案，取得已部署的智能合約地址。
	stringsTemp, err := tools.FileToStrings(addressFile)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(stringsTemp[0])

	fmt.Print("Contract address =\t", contractAddress.Hex(), "\n")

	// 建立與智能合約溝通的實體 Instance。
	instance, err := contract.NewBank(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("--------------------------------------------------\n\n")
	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 := call(keyWrapper1.Address)

	fmt.Print("1、查看目前 Account 1 儲存在 Bank 合約內的以太幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的以太幣。
	bankBalance, err := instance.GetBankBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	bankBalance.Div(bankBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nBank Balance =\t", bankBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
	auth1, err := sign(client, privateKeyECDSA1, 3, 1000000)
	if err != nil {
		log.Fatal(err, "\n")
	}

	fmt.Print("2、Account 1 將 3 顆以太幣存至 Bank 合約內。\n")

	// 將以太幣儲存至 Bank 合約內。
	depositInfo, err := instance.Deposit(auth1)
	if err != nil {
		log.Fatal(err, "\n")
	}

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nDeposit info:\t", depositInfo, "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("3、查看目前 Account 1 儲存在 Bank 合約內的以太幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的以太幣。
	bankBalance, err = instance.GetBankBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	bankBalance.Div(bankBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nBank Balance =\t", bankBalance.String(), "\n\n")

	fmt.Print("--------------------------------------------------\n\n")
	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
	auth1, err = sign(client, privateKeyECDSA1, 0, 1000000)
	if err != nil {
		log.Fatal(err, "\n")
	}

	fmt.Print("4、Account 1 將 Bank 帳戶內的 1 顆以太幣轉至 Account 2 的 Bank 帳戶。\n")

	// Account 1 將 Bank 帳戶內的 1 顆以太幣轉至 Account 2 的 Bank 帳戶。
	transferInfo, err := instance.Transfer(auth1, keyWrapper2.Address, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nTransfer info:\t", transferInfo, "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("5、查看目前 Account 1 儲存在 Bank 合約內的以太幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的以太幣。
	bankBalance, err = instance.GetBankBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	bankBalance.Div(bankBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nBank Balance =\t", bankBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts2 := call(keyWrapper2.Address)

	fmt.Print("6、查看目前 Account 2 儲存在 Bank 合約內的以太幣。\n")

	// 查看目前 Account 2 儲存在 Bank 合約內的以太幣。
	bankBalance, err = instance.GetBankBalance(opts2)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	bankBalance.Div(bankBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 2 =\t", keyWrapper2.Address.Hex(), "\nBank Balance =\t", bankBalance.String(), "\n\n")

	fmt.Print("--------------------------------------------------\n\n")
	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
	auth1, err = sign(client, privateKeyECDSA1, 0, 1000000)
	if err != nil {
		log.Fatal(err, "\n")
	}

	fmt.Print("7、Account 1 將 Bank 帳戶內的 1 顆以太幣領出。", "\n")

	// Account 1 將 Bank 帳戶內的 1 顆以太幣領出。
	withdrawInfo, err := instance.Withdraw(auth1, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nDeposit info:\t", withdrawInfo, "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("8、查看目前 Account 1 儲存在 Bank 合約內的以太幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的以太幣。
	bankBalance, err = instance.GetBankBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	bankBalance.Div(bankBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nBank Balance =\t", bankBalance.String(), "\n\n")

	fmt.Print("--------------------------------------------------\n\n")
	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("9、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的 Bank Coin 幣。
	coinBalance, err := instance.GetCoinBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	coinBalance.Div(coinBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t\t", keyWrapper1.Address.Hex(), "\nBank_Coin Balance =\t", coinBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
	auth1, err = sign(client, privateKeyECDSA1, 0, 1000000)
	if err != nil {
		log.Fatal(err, "\n")
	}

	fmt.Print("10、Account 1 鑄出 3 顆 Bank_Coin 幣至 Bank 帳戶。", "\n")

	// Account 1 鑄出 3 顆 Bank_Coin 幣至 Bank 帳戶。
	mintInfo, err := instance.Mint(auth1, big.NewInt(3))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nMint info:\t", mintInfo, "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("11、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的 Bank Coin 幣。
	coinBalance, err = instance.GetCoinBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	coinBalance.Div(coinBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t\t", keyWrapper1.Address.Hex(), "\nBank_Coin Balance =\t", coinBalance.String(), "\n\n")

	fmt.Print("--------------------------------------------------\n\n")
	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
	auth1, err = sign(client, privateKeyECDSA1, 0, 1000000)
	if err != nil {
		log.Fatal(err, "\n")
	}

	fmt.Print("12、Account 1 空投 1 顆 Bank_Coin 幣至 Account 2 的 Bank 帳戶內。", "\n")

	// Account 1 空投 1 顆 Bank_Coin 幣至 Account 2 的 Bank 帳戶內。
	transferCoinInfo, err := instance.TransferCoin(auth1, keyWrapper2.Address, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Account 1 =\t\t", keyWrapper1.Address.Hex(), "\nTransfer Bank_Coin info:\t", transferCoinInfo, "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("13、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的 Bank Coin 幣。
	coinBalance, err = instance.GetCoinBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	coinBalance.Div(coinBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t\t", keyWrapper1.Address.Hex(), "\nBank_Coin Balance =\t", coinBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts2 = call(keyWrapper2.Address)

	fmt.Print("14、查看目前 Account 2 儲存在 Bank 合約內的 Bank_Coin 幣。\n")

	// 查看目前 Account 2 儲存在 Bank 合約內的 Bank Coin 幣。
	coinBalance, err = instance.GetCoinBalance(opts2)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	coinBalance.Div(coinBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 2 =\t\t", keyWrapper2.Address.Hex(), "\nBank_Coin Balance =\t", coinBalance.String(), "\n\n")

	fmt.Print("--------------------------------------------------\n\n")
	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("15、查看目前 Account 1 儲存在 Bank 合約內的以太幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的以太幣。
	bankBalance, err = instance.GetBankBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	bankBalance.Div(bankBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nBank Balance =\t", bankBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts2 = call(keyWrapper2.Address)

	fmt.Print("16、查看目前 Account 2 儲存在 Bank 合約內的以太幣。\n")

	// 查看目前 Account 2 儲存在 Bank 合約內的以太幣。
	bankBalance, err = instance.GetBankBalance(opts2)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	bankBalance.Div(bankBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 2 =\t", keyWrapper2.Address.Hex(), "\nBank Balance =\t", bankBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("17、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的 Bank Coin 幣。
	coinBalance, err = instance.GetCoinBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	coinBalance.Div(coinBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t\t", keyWrapper1.Address.Hex(), "\nBank_Coin Balance =\t", coinBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts2 = call(keyWrapper2.Address)

	fmt.Print("18、查看目前 Account 2 儲存在 Bank 合約內的 Bank_Coin 幣。\n")

	// 查看目前 Account 2 儲存在 Bank 合約內的 Bank Coin 幣。
	coinBalance, err = instance.GetCoinBalance(opts2)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	coinBalance.Div(coinBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 2 =\t\t", keyWrapper2.Address.Hex(), "\nBank_Coin Balance =\t", coinBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
	auth2, err := sign(client, privateKeyECDSA2, 0, 1000000)
	if err != nil {
		log.Fatal(err, "\n")
	}

	fmt.Print("19、Account 2 覺得 Bank_Coin 幣會大漲，所以用 Bank 帳戶裡的 1 顆以太幣，向 Account 1 購買 1 顆 Bank_Coin 幣，並存至 Bank 帳戶內。\n")

	// Account 2 覺得 Bank_Coin 幣會大漲，所以用 Bank 帳戶裡的 1 顆以太幣，向 Account 1 購買 1 顆 Bank_Coin 幣，並存至 Bank 帳戶內。
	buyCoinInfo, err := instance.Buy(auth2, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Account 2 =\t\t", keyWrapper2.Address.Hex(), "\nBuy Bank_Coin info:\t", buyCoinInfo, "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("20、查看目前 Account 1 儲存在 Bank 合約內的以太幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的以太幣。
	bankBalance, err = instance.GetBankBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	bankBalance.Div(bankBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t", keyWrapper1.Address.Hex(), "\nBank Balance =\t", bankBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts2 = call(keyWrapper2.Address)

	fmt.Print("21、查看目前 Account 2 儲存在 Bank 合約內的以太幣。\n")

	// 查看目前 Account 2 儲存在 Bank 合約內的以太幣。
	bankBalance, err = instance.GetBankBalance(opts2)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	bankBalance.Div(bankBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 2 =\t", keyWrapper2.Address.Hex(), "\nBank Balance =\t", bankBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts1 = call(keyWrapper1.Address)

	fmt.Print("22、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。\n")

	// 查看目前 Account 1 儲存在 Bank 合約內的 Bank Coin 幣。
	coinBalance, err = instance.GetCoinBalance(opts1)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	coinBalance.Div(coinBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 1 =\t\t", keyWrapper1.Address.Hex(), "\nBank_Coin Balance =\t", coinBalance.String(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
	opts2 = call(keyWrapper2.Address)

	fmt.Print("23、查看目前 Account 2 儲存在 Bank 合約內的 Bank_Coin 幣。\n")

	// 查看目前 Account 2 儲存在 Bank 合約內的 Bank Coin 幣。
	coinBalance, err = instance.GetCoinBalance(opts2)
	if err != nil {
		log.Fatal(err)
	}

	// 1 ether = 1,000,000,000,000,000,000 wei
	coinBalance.Div(coinBalance, big.NewInt(int64(math.Pow10(18))))

	fmt.Print("Account 2 =\t\t", keyWrapper2.Address.Hex(), "\nBank_Coin Balance =\t", coinBalance.String(), "\n\n")

	fmt.Print("--------------------------------------------------\n\n")
	// --------------------------------------------------

	fmt.Print("24、查看目前 Bank 合約的擁有者。\n")

	// 查看目前 Bank 合約的擁有者。
	contractOwner, err := instance.GetOwner(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Bank contract owner =\t", contractOwner.Hex(), "\n\n")

	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
	auth1, err = sign(client, privateKeyECDSA1, 0, 1000000)
	if err != nil {
		log.Fatal(err, "\n")
	}

	fmt.Print("25、Account 1 覺得 Bank 沒有價值，於是將 Bank 的擁有權給充滿信仰的 Account 2。\n")

	// Account 1 覺得 Bank 沒有價值，於是將 Bank 的擁有權給充滿信仰的 Account 2。
	transferOwnerInfo, err := instance.TransferOwner(auth1, keyWrapper2.Address)
	if err != nil {
		log.Fatal(err, "\n")
	}

	fmt.Print("Account 1 =\t\t\t", keyWrapper1.Address.Hex(), "\nTransfer contract owner info:\t", transferOwnerInfo, "\n\n")

	// --------------------------------------------------

	fmt.Print("26、查看目前 Bank 合約的擁有者。\n")

	// 查看目前 Bank 合約的擁有者。
	contractOwner, err = instance.GetOwner(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Bank contract owner =\t", contractOwner.Hex(), "\n\n")

	fmt.Print("--------------------------------------------------\n\n")
	// --------------------------------------------------

	// 建立用於與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
	auth1, err = sign(client, privateKeyECDSA1, 0, 1000000)
	if err != nil {
		log.Fatal(err, "\n")
	}

	// Account 1 將 Bank 帳戶內的 1 顆以太幣領出。
	withdrawInfo, err = instance.Withdraw(auth1, big.NewInt(2))
	if err != nil {
		log.Fatal(err)
	}

}

// 	sign 用於建立與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
func sign(client *ethclient.Client, prvKeyECDSA *ecdsa.PrivateKey, ether int64, gasLimit uint64) (auth *bind.TransactOpts, err error) {
	// 從 Private Key ECDSA Struct 取得公鑰結構 Public Key ECDSA Struct，裡頭包含了一組橢圓曲線參數，及一組公鑰曲線點座標 Q = d_A * G。下述計算方法擇一。
	// 方法一：
	// pubKeyECDSA := &prvKeyECDSA.PublicKey
	// 方法二：
	// /*
	pubKey := prvKeyECDSA.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Error casting public key to ECDSA")
	}

	// 取得 Address。
	address := crypto.PubkeyToAddress(*pubKeyECDSA)

	// 取得此地址下一組的交易流水號（Nonce）。
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		return nil, errors.New("Error casting nonce from client with given specific address")
	}

	// 取得目前以太坊網路建議的 Gas price。
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.New("Can't get the suggestion gas price from client")
	}

	// 建立用於與智能合約溝通用的憑證（用於會改變以太坊網路狀態的智能合約方法）。
	auth = bind.NewKeyedTransactor(prvKeyECDSA)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(ether * int64(math.Pow10(18))) // 1 ether = 1,000,000,000,000,000,000 wei
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice // unit: Gwei

	return auth, nil
}

// call 用來建立用於與智能合約溝通用的憑證（用於不會改變以太坊網路狀態的智能合約方法）。
func call(address common.Address) (opts *bind.CallOpts) {

	opts = new(bind.CallOpts)
	opts.From = address
	opts.Context = context.Background()
	opts.Pending = false
	return opts
}
