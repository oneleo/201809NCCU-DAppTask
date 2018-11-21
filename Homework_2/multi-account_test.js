// 使用 JavaScript 的嚴格模式。
'use strict';

// 定義 Byte 轉 Hex String 函數。
function toHexString(byteArray) {
    return Array.prototype.map.call(byteArray, function (byte) {
        return ('0' + (byte & 0xFF).toString(16)).slice(-2);
    }).join('');
}

// 定義 Hex String 轉 Byte 函數。
function toByteArray(hexString) {
    var result = [];
    while (hexString.length >= 2) {
        result.push(parseInt(hexString.substring(0, 2), 16));
        hexString = hexString.substring(2, hexString.length);
    }
    return result;
}

// 定義要連線的以太坊網路。
//const url = 'http://localhost:8545';
const url = 'https://rinkeby.infura.io';
//const url = 'wss://rinkeby.infura.io/ws';

// 定義 keystore 存放位置及解開密碼。
const keystoreFile1 = './keystore1';
const keystoreFile2 = './keystore2';
const keystorePasswd1 = 'nccu';
const keystorePasswd2 = 'nccu';
// 設置智能合約 ABI 及 BIN 的放置地址。
const cntrtAbiFile = './02_Contract/Bank.abi';
const cntrtBinFile = './02_Contract/Bank.bin';
// 存放部署後的智能合約地址。
const cntrtAdrFile = './address.txt';
// 設置 gas limit 及 gas price。
const gasLmt = 3400000;
const gasPrc = 2000000000;

// 定義所需的套件。
const wallet = require('ethereumjs-wallet');
const fs = require('fs');
const web3 = require('web3');
// Reference:
// https://github.com/trufflesuite/truffle-hdwallet-provider
// https://github.com/nosuchip/truffle-privatekey-provider
// https://ethereum.stackexchange.com/questions/46203/deploy-contract-on-live-testnet-using-private-key-truffle
const prvKeyProvider = require("truffle-privatekey-provider");

// 定義 Bank.sol 的 abi。
const abi = JSON.parse(fs.readFileSync(cntrtAbiFile).toString());
// 定義 Bank.sol 的 byte code。
const bytecode = '0x' + fs.readFileSync(cntrtBinFile).toString();

// 將 keystore Json 讀出後，再使用 nccu 作為密碼將 Private 解密。
let keyStoreAccount1 = wallet.fromV3(fs.readFileSync(keystoreFile1).toString(), keystorePasswd1, true);
let keyStoreAccount2 = wallet.fromV3(fs.readFileSync(keystoreFile2).toString(), keystorePasswd2, true);
// 取得 keystore 的 private key byte 型式。
let prvKey1 = keyStoreAccount1.getPrivateKey();
let prvKey2 = keyStoreAccount2.getPrivateKey();
// 取得 keystore 的 private key hex string 型式。
let prvKeyHex1 = toHexString(prvKey1);
let prvKeyHex2 = toHexString(prvKey2);
// 印出從 keystore 取得的 private key hex string 型式。
console.log("Recovered private key1:\t", prvKeyHex1);
console.log("Recovered private key2:\t", prvKeyHex2);

// 根據輸入的 private key hex string 建立一個 accounts provider 提供者。
let provider1 = new prvKeyProvider(prvKeyHex1, url);
let provider2 = new prvKeyProvider(prvKeyHex2, url);
// 根據 provider 建立一 web3 連線實體。
let client1 = new web3(provider1);
let client2 = new web3(provider2);

/*
// Doesn't work.
let acnt2 = client.eth.accounts.privateKeyToAccount("0x" + prvKeyHex2);
client.eth.accounts.wallet.add(acnt2);
client.eth.defaultAccount = acnt2.address;
*/

// 取得 client 內的所有帳號。
client1.eth.getAccounts().then(async function (accounts) {
    console.log("Account 1:", accounts[0]);
});

client2.eth.getAccounts().then(async function (accounts) {
    console.log("Account 2:", accounts[0]);
});

// At termination, 'provider.engine.stop()' should be called to finish the process elegantly.
provider1.engine.stop();
provider2.engine.stop();
