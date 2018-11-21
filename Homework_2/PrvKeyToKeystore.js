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
const wallet = require('ethereumjs-wallet');
const keccak256 = require('js-sha3').keccak256;
const fs = require('fs');
const ethUtil = require('ethereumjs-util');

// 定義 keystore 存放位置及加密密碼。
const keystoreFile1 = './keystore1';
const keystoreFile2 = './keystore2';
const passwd1 = 'nccu';
const passwd2 = 'nccu';

// 在私鑰前加上 "0x"。
let prvKeyStr1 = "ad99cb7d0f056d79ae84779e9649555a2abb48e43f9cbf1287ba03d93c2d33c6";
prvKeyStr1 = "0x" + prvKeyStr1;
let prvKeyStr2 = "1354c43353a9d482d119f8bac9e1bb55ce79436c94c5c8a7144ddc5d20a7276b";
prvKeyStr2 = "0x" + prvKeyStr2;

// 儲存至緩衝。
let prvKeyBuf1 = ethUtil.toBuffer(prvKeyStr1);
let prvKeyBuf2 = ethUtil.toBuffer(prvKeyStr2);

// 將私鑰包成帳號。
let account1 = wallet.fromPrivateKey(prvKeyBuf1);
let account2 = wallet.fromPrivateKey(prvKeyBuf2);

// 取得 privKey。
//let privKey1 = account1.getPrivateKey();
//let privKey2 = account2.getPrivateKey();

// 取得 pubKey。
let pubKey1 = account1.getPublicKey();
let pubKey2 = account2.getPublicKey();

// 將公鑰 hash。
let public_key_hash1 = keccak256(pubKey1);
let public_key_hash2 = keccak256(pubKey2);

// 取得 address。
let address1 = '0x' + public_key_hash1.substring(public_key_hash1.length - 20 * 2); // 1 byte = 2 hex digits。
let address2 = '0x' + public_key_hash2.substring(public_key_hash2.length - 20 * 2); // 1 byte = 2 hex digits。

// 印出地址。
console.log("address1:\t\t", address1);
console.log("address2:\t\t", address2);

// 使用 nccu 作為密碼將 Private 加密。
let keystore1 = account1.toV3String(passwd1);
let keystore2 = account2.toV3String(passwd2);

// 將 keystore1 Json 寫入 keystore1 檔內。
fs.writeFileSync(keystoreFile1, keystore1, 'utf8', function (err) {
    if (err) {
        console.log("An error occured while writing JSON Object to File.");
        return console.error(err);
    }
});

// 將 keystore Json 寫入 03_keystore 檔內。
fs.writeFileSync(keystoreFile2, keystore2, 'utf8', function (err) {
    if (err) {
        console.log("An error occured while writing JSON Object to File.");
        return console.error(err);
    }
});

// --------------------------------------------------

// 將 keystore Json 讀出後，再使用 nccu 作為密碼將 Private 解密。
let recoverWallet1 = wallet.fromV3(fs.readFileSync(keystoreFile1).toString(), passwd1, true);
let recoverWallet2 = wallet.fromV3(fs.readFileSync(keystoreFile2).toString(), passwd2, true);

// 取得私鑰。
let recoverPrivKey1 = recoverWallet1.getPrivateKey();
let recoverPrivKey2 = recoverWallet2.getPrivateKey();

// 取得私鑰 HEX。
let recoverPrivKeyHex1 = toHexString(recoverPrivKey1);
let recoverPrivKeyHex2 = toHexString(recoverPrivKey2);

// 印出私鑰 HEX。
console.log("Recovered private key1:\t", recoverPrivKeyHex1);
console.log("Recovered private key2:\t", recoverPrivKeyHex2);
