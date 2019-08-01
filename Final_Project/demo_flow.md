
程式碼
https://bitbucket.org/Starfine/nccu10

假設有二帳號：
帳號一：0x131746C3CaB2C6525DCCb1854458CdeE8199098F
帳號二：0x24Dd213eA94c146940C230269e0B48dBc468b1aF

---

1、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）部署 Hub 合約


2、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 Hub 合約執行 createEvent 函數建立票券

aaa 隊
1545696000（2018/12/25）
1546246200（12/31/2018 @ 8:50am (UTC)）
10000000000

→
0x939b706b5cf79cb397e8075e2bc964e71bbe0993

3、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 Hub 合約執行 createEvent 函數建立票券

bbb 隊
1545696000（2018/12/25）
1546246200（12/31/2018 @ 8:50am (UTC)）
10000000000

→
0xc33fc7dc5ced5b4efd20fe5bcca2e5153e45bdb6

---

4、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 aaa 內執行 setHost 函數
aaa 券的地址

5、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 aaa 內執行 setHost 函數
bbb 券的地址

6、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 bbb 內執行 setHost 函數
aaa 券的地址

7、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 bbb 內執行 setHost 函數
bbb 券的地址

8、在 aaa 內執行 hosts 函數
aaa 券的地址 → true
bbb 券的地址 → true

9、在 bbb 內執行 hosts 函數
aaa 券的地址 → true
bbb 券的地址 → true

---

10、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 aaa 內執行 registry 函數，並送 0.01 以太

11、用 0x24Dd213eA94c146940C230269e0B48dBc468b1aF 在 aaa 內執行 registry 函數，並送 0.01 以太

12、用 0x24Dd213eA94c146940C230269e0B48dBc468b1aF 在 aaa 內執行 registry 函數，並送 0.01 以太

13、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 bbb 內執行 registry 函數，並送 0.01 以太

14、用 0x24Dd213eA94c146940C230269e0B48dBc468b1aF 在 bbb 內執行 registry 函數，並送 0.01 以太

15、用 0x24Dd213eA94c146940C230269e0B48dBc468b1aF 在 bbb 內執行 registry 函數，並送 0.01 以太

16、在 aaa 內執行 ownerOf 函數
1 → 0x131746c3cab2c6525dccb1854458cdee8199098f（HOST）
2 → 0x24dd213ea94c146940c230269e0b48dbc468b1af
3 → 0x24dd213ea94c146940c230269e0b48dbc468b1af

17、在 bbb 內執行 ownerOf 函數
1 → 0x131746c3cab2c6525dccb1854458cdee8199098f（HOST）
2 → 0x24dd213ea94c146940c230269e0b48dbc468b1af
3 → 0x24dd213ea94c146940c230269e0b48dbc468b1af

---

18、售票時間到後，用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 aaa 內執行 Win 函數，
aaa 券的地址（假設 aaa 隊贏了）

19、售票時間到後，用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（HOST）在 bbb 內執行 Win 函數，
aaa 券的地址（假設 aaa 隊贏了）

20、上面會失敗，所以開外掛
（1）用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F 在 bbb 內執行 setIsSettle 函數，
true

---

11、用 0x131746C3CaB2C6525DCCb1854458CdeE8199098F（因為有買 aaa 票券）在 aaa 內執行 Reward 函數
1
0x24Dd213eA94c146940C230269e0B48dBc468b1aF（買輸的那一隊（bbb）的最後那一位）

12、用 0x24dd213ea94c146940c230269e0b48dbc468b1af（因為有買 aaa 票券）在 aaa 內執行 Reward 函數
2
0x24Dd213eA94c146940C230269e0B48dBc468b1aF（買輸的那一隊（bbb）的最後那一位）
3
0x24Dd213eA94c146940C230269e0B48dBc468b1aF（買輸的那一隊（bbb）的最後那一位）


