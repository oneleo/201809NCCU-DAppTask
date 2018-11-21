// 使用 JavaScript 的嚴格模式。
'use strict';

// 定義 Byte 轉 Hex String 函數。
function toHexString(byteArray) {
    return Array.prototype.map.call(byteArray, function (byte) {
        return ('0' + (byte & 0xFF).toString(16)).slice(-2);
    }).join('');
}

// 定義要連線的以太坊網路。
//const url = 'http://localhost:8545';
const url = 'https://rinkeby.infura.io';
//const url = 'wss://rinkeby.infura.io/ws';

// 定義 keystore 存放位置及解開密碼。
const keystoreFile = '../keystore1';
const keystorePasswd = 'nccu';
// 設置智能合約 ABI 及 BIN 的放置地址。
const cntrtAbiFile = '../02_Contract/Bank.abi';
// 存放部署後的智能合約地址。
const cntrtAdrFile = '../address.txt';
// 設置 gas limit 及 gas price。
const gasLmt = 3400000;
const gasPrc = 2000000000;

// 定義所需的套件。
const wallet = require('ethereumjs-wallet');
const fs = require('fs');
const web3 = require('web3');
const prvKeyProvider = require("truffle-privatekey-provider");

// 定義 Bank.sol 的 abi。
const abi = JSON.parse(fs.readFileSync(cntrtAbiFile).toString());
// 取得智能合約地址。
const cntrtAdr = fs.readFileSync(cntrtAdrFile).toString();

// 將 keystore Json 讀出後，再使用 nccu 作為密碼將 Private 解密。
let keyStoreAccount = wallet.fromV3(fs.readFileSync(keystoreFile).toString(), keystorePasswd, true);
// 取得 keystore 的 private key byte 型式。
let prvKey = keyStoreAccount.getPrivateKey();
// 取得 keystore 的 private key hex string 型式。
let prvKeyHex = toHexString(prvKey);

// 根據輸入的 private key hex string 建立一個 accounts provider 提供者。
let provider = new prvKeyProvider(prvKeyHex, url);
// 根據 provider 建立一 web3 連線實體。
let client = new web3(provider);

// 取得 client 內的所有帳號。
client.eth.getAccounts().then(async function (accounts) {

    // 透過 abi 建立一個即將部署在 client 上的 smart contract 實體 bank。
    let bank = await new client.eth.Contract(abi, cntrtAdr);

    // accounts[0] withdraw 1 ether
    await bank.methods.withdraw(1)
        .send({ from: accounts[0], gas: gasLmt, gasPrice: gasPrc, value: 0 })
        .then(function (receipt) { console.log("Contract address:\t", cntrtAdr, "\nNow address:\t\t", accounts[0], "\nWithdraw from bank (info):\n", receipt) })
        .catch(function (error) { console.log("Error occurred:\n", error); });
});

// At termination, 'provider.engine.stop()' should be called to finish the process elegantly.
provider.engine.stop();
