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

// 引入 npm-library。
const Wallet = require('ethereumjs-wallet');
const keccak256 = require('js-sha3').keccak256;

// 建立一組 keypair。
const wallet = Wallet.generate();

// 取得 privKey 並印出。
const privKey = wallet.getPrivateKey();
let privKeyHex = toHexString(privKey); // To Hex String。
//console.log("privKey:\n", privKey);
console.log("privKey:\n" + privKeyHex);

// 取得 pubKey 並印出。
const pubKey = wallet.getPublicKey();
let pubKeyHex = toHexString(pubKey); // To Hex String。
//console.log("pubKey:", pubKey);
console.log("pubKey:\n" + pubKeyHex);

// 取得 address 並印出。
//let address = wallet.getAddressString();
//console.log("address:\n" + address);
let public_key_hash = keccak256(pubKey);
let address2 = '0x' + public_key_hash.substring(public_key_hash.length - 20 * 2); // 1 byte = 2 hex digits。
console.log("address:\n" + address2);

// 使用 nccu 作為密碼將 Private 加密。
let keystore = wallet.toV3String('nccu');
console.log("keystore:\n" + keystore);

// 使用 nccu 作為密碼將 Private 解密。
let recoverWallet = Wallet.fromV3(keystore, 'nccu', true);
let recoverPrivKey = recoverWallet.getPrivateKey();
let recoverPrivKeyHex = toHexString(recoverPrivKey);
console.log("recover:\n" + recoverPrivKeyHex);