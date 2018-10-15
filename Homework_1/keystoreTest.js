// 使用 JavaScript 的嚴格模式。
'use strict';

// 引入 npm-library。
const Wallet = require('ethereumjs-wallet');
const keccak256 = require('js-sha3').keccak256;
const fs = require('fs');

const keystoreFile = "./Homework_1/keystore";

// 使用 nccu 作為密碼將 Private 解密。
let recoverWallet = Wallet.fromV3(fs.readFileSync(keystoreFile).toString(), 'nccu', true); // Wrong Password。
let recoverPrivKey = recoverWallet.getPrivateKey();
let recoverPrivKeyHex = toHexString(recoverPrivKey);
console.log("recover:\n" + recoverPrivKeyHex);