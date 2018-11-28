## 5、使用 Go 語言完成本次作業

### （1）部署智能合約至 Ganache private net

- 部署 Bank 智能合約程式原始碼請參考：

	- [deploy.go](./deploy/deploy.go)
	- [contract.go](./deploy/contract.go)

- 部署 Bank 智能合約程式執行結果：

```
API server listening at: 127.0.0.1:18207
Deploy from:		0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Contract address:	0x906637756CD20f232d9CB85D45A28d1143292bad
Transaction hash:	0xeb48d06fede3326e8a9a33af372676403f2148cf3210bf1dcea81cfe6324034f
Instance info:		&{{0xc00011f680} {0xc00011f680} {0xc00011f680}}
```

### （2）將本次作業使用 Go 語言完成

- exercise.go 程式原始碼請參考：

	- [exercise.go](./exercise/exercise.go)

- exercise.go 程式原始碼執行結果：

```
API server listening at: 127.0.0.1:43538
Contract address =	0x906637756CD20f232d9CB85D45A28d1143292bad

--------------------------------------------------

1、查看目前 Account 1 儲存在 Bank 合約內的以太幣。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank Balance =	0

2、Account 1 將 3 顆以太幣存至 Bank 合約內。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Deposit info:	&{{1 0xc02015c220 1000000 0xc020158280 0xc02015c200 [208 227 13 176] 0xc02015c3a0 0xc02015c360 0xc02015c380 <nil>} {<nil>} {<nil>} {<nil>}}

3、查看目前 Account 1 儲存在 Bank 合約內的以太幣。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank Balance =	3

--------------------------------------------------

4、Account 1 將 Bank 帳戶內的 1 顆以太幣轉至 Account 2 的 Bank 帳戶。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Transfer info:	&{{2 0xc02015c820 1000000 0xc0201587a0 0xc02015c800 [169 5 156 187 0 0 0 0 0 0 0 0 0 0 0 0 24 100 46 228 240 65 151 131 45 172 119 0 201 183 47 212 175 239 21 135 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1] 0xc02015c940 0xc02015c900 0xc02015c920 <nil>} {<nil>} {<nil>} {<nil>}}

5、查看目前 Account 1 儲存在 Bank 合約內的以太幣。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank Balance =	2

6、查看目前 Account 2 儲存在 Bank 合約內的以太幣。
Account 2 =	0x18642eE4f04197832Dac7700c9b72fd4aFEf1587
Bank Balance =	1

--------------------------------------------------

7、Account 1 將 Bank 帳戶內的 1 顆以太幣領出。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Deposit info:	&{{3 0xc0201c6580 1000000 0xc000011180 0xc0201c6560 [46 26 125 77 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1] 0xc0201c66a0 0xc0201c6660 0xc0201c6680 <nil>} {<nil>} {<nil>} {<nil>}}

8、查看目前 Account 1 儲存在 Bank 合約內的以太幣。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank Balance =	1

--------------------------------------------------

9、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。
Account 1 =		0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank_Coin Balance =	0

10、Account 1 鑄出 3 顆 Bank_Coin 幣至 Bank 帳戶。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Mint info:	&{{4 0xc0201c6b80 1000000 0xc000011680 0xc0201c6b60 [160 113 45 104 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 3] 0xc0201c6ca0 0xc0201c6c60 0xc0201c6c80 <nil>} {<nil>} {<nil>} {<nil>}}

11、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。
Account 1 =		0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank_Coin Balance =	3

--------------------------------------------------

12、Account 1 空投 1 顆 Bank_Coin 幣至 Account 2 的 Bank 帳戶內。
Account 1 =		0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Transfer Bank_Coin info:	&{{5 0xc0201c7060 1000000 0xc000011b80 0xc0201c7040 [141 222 96 250 0 0 0 0 0 0 0 0 0 0 0 0 24 100 46 228 240 65 151 131 45 172 119 0 201 183 47 212 175 239 21 135 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1] 0xc0201c7180 0xc0201c7140 0xc0201c7160 <nil>} {<nil>} {<nil>} {<nil>}}

13、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。
Account 1 =		0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank_Coin Balance =	2

14、查看目前 Account 2 儲存在 Bank 合約內的 Bank_Coin 幣。
Account 2 =		0x18642eE4f04197832Dac7700c9b72fd4aFEf1587
Bank_Coin Balance =	1

--------------------------------------------------

15、查看目前 Account 1 儲存在 Bank 合約內的以太幣。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank Balance =	1

16、查看目前 Account 2 儲存在 Bank 合約內的以太幣。
Account 2 =	0x18642eE4f04197832Dac7700c9b72fd4aFEf1587
Bank Balance =	1

17、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。
Account 1 =		0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank_Coin Balance =	2

18、查看目前 Account 2 儲存在 Bank 合約內的 Bank_Coin 幣。
Account 2 =		0x18642eE4f04197832Dac7700c9b72fd4aFEf1587
Bank_Coin Balance =	1

19、Account 2 覺得 Bank_Coin 幣會大漲，所以用 Bank 帳戶裡的 1 顆以太幣，向 Account 1 購買 1 顆 Bank_Coin 幣，並存至 Bank 帳戶內。
Account 2 =		0x18642eE4f04197832Dac7700c9b72fd4aFEf1587
Buy Bank_Coin info:	&{{0 0xc0201c7900 1000000 0xc0201ec2a0 0xc0201c78e0 [217 106 9 74 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1] 0xc0201c7a20 0xc0201c79e0 0xc0201c7a00 <nil>} {<nil>} {<nil>} {<nil>}}

20、查看目前 Account 1 儲存在 Bank 合約內的以太幣。
Account 1 =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank Balance =	2

21、查看目前 Account 2 儲存在 Bank 合約內的以太幣。
Account 2 =	0x18642eE4f04197832Dac7700c9b72fd4aFEf1587
Bank Balance =	0

22、查看目前 Account 1 儲存在 Bank 合約內的 Bank_Coin 幣。
Account 1 =		0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Bank_Coin Balance =	1

23、查看目前 Account 2 儲存在 Bank 合約內的 Bank_Coin 幣。
Account 2 =		0x18642eE4f04197832Dac7700c9b72fd4aFEf1587
Bank_Coin Balance =	2

--------------------------------------------------

24、查看目前 Bank 合約的擁有者。
Bank contract owner =	0x131746C3CaB2C6525DCCb1854458CdeE8199098F

25、Account 1 覺得 Bank 沒有價值，於是將 Bank 的擁有權給充滿信仰的 Account 2。
Account 1 =			0x131746C3CaB2C6525DCCb1854458CdeE8199098F
Transfer contract owner info:	&{{6 0xc020200120 1000000 0xc0201ecaa0 0xc020200100 [79 178 228 93 0 0 0 0 0 0 0 0 0 0 0 0 24 100 46 228 240 65 151 131 45 172 119 0 201 183 47 212 175 239 21 135] 0xc020200240 0xc020200200 0xc020200220 <nil>} {<nil>} {<nil>} {<nil>}}

26、查看目前 Bank 合約的擁有者。
Bank contract owner =	0x18642eE4f04197832Dac7700c9b72fd4aFEf1587

--------------------------------------------------
```