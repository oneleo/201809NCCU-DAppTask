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
const keystoreFile = './keystore1';
const keystorePasswd = 'nccu';
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
// https://github.com/lendingblock/loan-contracts#using-private-key
const prvKeyProvider = require("truffle-privatekey-provider");

// 定義 Bank.sol 的 abi。
const abi = JSON.parse(fs.readFileSync(cntrtAbiFile).toString());
// 定義 Bank.sol 的 byte code。
const bytecode = '0x' + fs.readFileSync(cntrtBinFile).toString();

// 將 keystore Json 讀出後，再使用 nccu 作為密碼將 Private 解密。
let keyStoreAccount = wallet.fromV3(fs.readFileSync(keystoreFile).toString(), keystorePasswd, true);
// 取得 keystore 的 private key byte 型式。
let prvKey = keyStoreAccount.getPrivateKey();
// 取得 keystore 的 private key hex string 型式。
let prvKeyHex = toHexString(prvKey);
// 印出從 keystore 取得的 private key hex string 型式。
//console.log("Recovered private key:\t", prvKeyHex);

// 根據輸入的 private key hex string 建立一個 accounts provider 提供者。
let provider = new prvKeyProvider(prvKeyHex, url);
// 根據 provider 建立一 web3 連線實體。
let client = new web3(provider);

// 取得 client 內的所有帳號。
client.eth.getAccounts().then(async function (accounts) {
    // deploy contract
    // your code
    console.log("Deploy from account:\t", accounts[0]);
    // 透過 abi 建立一個即將部署在 client 上的 smart contract 實體 bank。
    let bank = await new client.eth.Contract(abi);

    // 給定 bytecode 合約原始碼、arguments 初始值、from 部署者、gas limit、以太幣 value 後，將 smart contrace 部署至以太坊網路。
    let result = await bank
        .deploy({ data: bytecode, arguments: [] })
        .send({ from: accounts[0], gas: gasLmt, gasPrice: gasPrc, value: 0 })
        .catch(function (error) { console.log("Error occurred:\n", error); });

    //This will display the address to which your contract was deployed
    console.log("Contract deployed to:\t", result.options.address);
    // 部署成功後將合約 address 寫至 address.txt
    fs.writeFileSync(cntrtAdrFile, result.options.address);
});

// At termination, 'provider.engine.stop()' should be called to finish the process elegantly.
provider.engine.stop();
