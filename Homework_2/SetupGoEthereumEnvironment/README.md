## 在 Windows 10 作業系統內建置 go-ethereum 開發環境
> 參考：
> [https://github.com/ethereum/go-ethereum/wiki/Building-Ethereum](https://github.com/ethereum/go-ethereum/wiki/Building-Ethereum)

### 緣由

### （1）安裝 Git for Windows

- 進到 Git 官網 [https://git-scm.com/downloads](https://git-scm.com/downloads)，並點選 `Download *.*.* for Windows`（預設是下載到 `%USERPROFILE%\Downloads` 資料夾內）。

![](../images/061.png)

- 檔案下載後，滑鼠連點兩下執行安裝檔。

![](../images/062.png)

- 點選【Next >】。

![](../images/063.png)

- 點選【Next >】。

![](../images/064.png)

- 取消勾選【Additional icons】、【Windows Explorer integration】→ 勾選【Use a TrueType font in all console windows】、【Check daily for Git for Windows updates】 → 點選【Next >】。

![](../images/065.png)

- 點選【Next >】。

![](../images/066.png)

- 點選【Use Vim (the ubiquitous text editor) as Git's default editor】→【Next >】。

![](../images/067.png)

- 點選【Use Git from the Windows Command Prompt】→【Next >】。

![](../images/068.png)

- 點選【Use the OpenSSL library】→【Next >】。

![](../images/069.png)

- 點選【Checkout Windows-style, commit Unix-style line endings】→【Next >】。 

![](../images/070.png)

- 點選【Use MinTTY (the default terminal of MSYS2)】→【Next >】。

![](../images/071.png)

- 勾選【Enable file system caching】、【Enable Git Credential Manager】、【Enable symbolic links】→【Next >】。

![](../images/072.png)

- 勾選【Enable experimental, built-in rebase】、【Enable experimental, built-in stash】→【Install】。

![](../images/073.png)

- 取消勾選【View Release Notes】→ 點選【Finish】。

![](../images/074.png)

### （2）安裝 Golang 程式語言：

- 進到 Go Language 官網 [https://golang.org/dl/](https://golang.org/dl/)，並點選 `go*.*.*.windows-amd64.msi`（預設是下載到 `%USERPROFILE%\Downloads` 資料夾內）。

![](../images/081.png)

- 檔案下載後，滑鼠連點兩下執行安裝檔。

![](../images/082.png)

- 點選【Next】。

![](../images/083.png)

- 點選【Next】。

![](../images/084.png)

- 點選【Next】。

![](../images/085.png)

- 點選【Install】。

![](../images/086.png)

- 點選【Finish】。

![](../images/087.png)

### （3）安裝 MinGW-64（這邊以安裝 msys2-x86_64_20180531.exe 至 C:\msys64 為例）
> 參考：
> [https://shaochien.gitbooks.io/how-to-use-gcc-to-develop-c-cpp-on-windows/content/shell-select-and-install-gcc.html](https://shaochien.gitbooks.io/how-to-use-gcc-to-develop-c-cpp-on-windows/content/shell-select-and-install-gcc.html)
> [https://cwchen.tw/windows-survival/mingw-msys/](https://cwchen.tw/windows-survival/mingw-msys/)
> [https://labs.mediatek.com/zh-tw/faq/FAQ18854](https://labs.mediatek.com/zh-tw/faq/FAQ18854)
> 注意：經過實測，若有安裝 Haskell 等內建專用 gcc 套件的軟體，要先將之刪除，不然後續易會產生像是「realgcc.exe: fatal error: no input files」等問題

- 進到 MinGW-64 官網 [http://www.msys2.org/](http://www.msys2.org/)，並點選 `MSYS2-X86_64-20*.exe`（預設是下載到 `%USERPROFILE%\Downloads` 資料夾內）。

![](../images/091.png)

- 檔案下載後，滑鼠連點兩下執行安裝檔。

![](../images/092.png)

- 點選【Next】。

![](../images/093.png)

- 點選【Next】。

![](../images/094.png)

- 點選【Next】。

![](../images/095.png)

- 勾選【Run MSYS2 64bit now.】→ 點選【Finish】。

![](../images/096.png)

- 此時會自動開啟一個 MSYS 視窗，輸入下方指令更新核心軟體：

```
$ pacman -Syu --noconfirm
```

![](../images/097.png)

- 出現「close your terminal window instead of calling exit」訊息，關閉目前的 MSYS 視窗。

![](../images/098.png)

- 在【開始】→【MSYS2 64bit】→【MSYS2 MSYS】來啟動 MSYS 視窗。

![](../images/099.png)

- 再次輸入下方指令更新核心及其他軟體。
```
$ pacman -Syu --noconfirm
$ pacman -Su --noconfirm
$ pacman -Sy --noconfirm
```

![](../images/100.png)

![](../images/101.png)

- 可以使用下方指令搜尋可安裝的 gcc 套件：
```
$ pacman -Ss gcc
$ pacman -Sl | grep gcc
```
註：將會看到 msys2 的 gcc 套件有三種類可選，這邊以安裝 Mingw64 版的 gcc：
* mingw-w64-i686：32 位元套件 (mingw32)
* mingw-w64-x86_64：64 位元套件 (mingw64)
* 無字首：僅用於 msys 環境中 (msys)

![](../images/102.png)

- 安裝 gcc、make。
```
$ pacman -S mingw-w64-x86_64-gcc make --noconfirm
```

![](../images/103.png)

- MSYS2 會將 gcc 可執行檔放在「C:\msys64\mingw64\bin」內，make 可執行檔會放在「C:\msys64\usr\bin」內。

![](../images/104.png)

![](../images/105.png)

- 設定 Windows 環境變數。點選「Windows 符號」→ 點選【Windows 系統】→ 點選【控制台】。

```
> set MINGW64="C:\msys64\mingw64"
> set MSYS64USR="C:\msys64\usr"
> set PATH="%PATH%;%MINGW64%\bin;%MSYS64USR%\bin"
```

![](../images/106.png)

- 點選【系統及安全性】。

![](../images/107.png)

- 點選【系統】。

![](../images/108.png)

- 點選【進階系統設定】。

![](../images/109.png)

- 在出現的「系統內容」視窗後，點選【進階】標籤 → 點選【環境變數(N)...】。

![](../images/110.png)

- 在上方「XXX 的使用者變數(U)」中，點選【新增(N)...】。

![](../images/111.png)

- 在出現的「新增使用者變數」視窗中，輸入下方數值：

	- 「變數名稱(N):」=【MINGW64】
	- 「變數值(V):」=【C:\msys64\mingw64】→ 點選【確定】

![](../images/112.png)

- 再點選【新增(N)...】。

![](../images/113.png)

- 在出現的「新增使用者變數」視窗中，輸入下方數值：

	- 「變數名稱(N):」=【MSYS64USR】
	- 「變數值(V):」=【C:\msys64\usr】→ 點選【確定】

![](../images/114.png)

- 點選上方【Path】→ 點選【編輯(E)...】。

![](../images/115.png)

- 點選【新增(N)】→ 輸入【%MINGW64%\bin】→ 再點選【新增(N)】→ 輸入【%MSYS64USR%\bin】。

![](../images/116.png)

- 點選【確定】完成環境變數的設置。

![](../images/117.png)

- 開啟命令提示字元（%WINDIR%\system32\cmd.exe）程式

![](../images/118.png)

- 輸入下列指令測試 gcc、make 是否有安裝成功：

```
> gcc -v
> make -v
```
![](../images/119.png)

### （4）安裝 Visual Studio Code 整合編譯環境

- 進到 Visual Studio Code 官網 [https://code.visualstudio.com/download](https://code.visualstudio.com/download)，並點選 `Windows` 版本（預設是下載到 `%USERPROFILE%\Downloads` 資料夾內）。

![](../images/121.png)

- 檔案下載後，滑鼠連點兩下執行安裝檔。

![](../images/122.png)

- 點選【下一步(N) >】。

![](../images/123.png)

- 點選【我接受合約(A)】→ 點選【下一步(N) >】。

![](../images/124.png)

- 點選【下一步(N) >】。

![](../images/125.png)

- 點選【下一步(N) >】。

![](../images/126.png)

- 點選【將「以 Code 開啟」動作加入 Windows 檔案總管檔案的操作功能表中】、【將「以 Code 開啟」動作加入 Windows 檔案總管目錄的操作功能表中】、【針對支援的檔案類型將 Code 註冊為編輯器】→【下一步(N) >】

![](../images/127.png)

- 點選【安裝(I)】。

![](../images/128.png)

- 點選【啟動 Visual Studio Code】→【完成(F)】。

![](../images/129.png)

### （5）安裝 Go Extension for Visual Studio Code
> 套件介紹頁面：[https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)

- 在【礦充功能】頁面下搜尋【go】關鍵字 → 點選【Install】。

![](../images/131.png)

- 安裝完成後，點選【Reload】將 Visual Studio Code 重新啟動。

![](../images/132.png)

- 點選【Ctrl + Shift + P】→ 輸入【Go: Install/Update tools】→【Enter】

![](../images/133.png)

- 勾選左側【勾選方塊】→【OK】來安裝 Go Extension for Visual Studio Code 需要的 go 套件。

![](../images/134.png)

- 要看到「All tools successfully installed. You're ready to Go :).」訊息才算安裝完成。

![](../images/135.png)

### （6）安裝 dotNET SDK for Visual Studio Code（可選，非必要）

- 進到 dotNET SDK 官網 [https://www.microsoft.com/net/learn/get-started/windows#windowscmd](https://www.microsoft.com/net/learn/get-started/windows#windowscmd)，並點選 `Download .NET SDK`（預設是下載到 `%USERPROFILE%\Downloads` 資料夾內）。

![](../images/141.png)

- 檔案下載後，滑鼠連點兩下執行安裝檔。

![](../images/142.png)

- 點選【Install】。

![](../images/143.png)

- 安裝完成後點選【Close】。

![](../images/144.png)

### （7）在命令提示字元（%WINDIR%\system32\cmd.exe）內輸入下列指令來安裝 go-ethereum 編譯環境

```
> go get -u -v golang.org/x/net/context
> go get -u -v github.com/ethereum/go-ethereum
> go install -v github.com/ethereum/go-ethereum/cmd/...
```

![](../images/151.png)

![](../images/152.png)

- 可以在 `%GOPATH%\bin` 中看到 go install 已安裝完成的可執行檔。

![](../images/153.png)


### （8-1）測試：（8-1 及 8-2 擇一）使用命令提示字元（%WINDIR%\system32\cmd.exe）編譯 go 語言連接至以太坊測試網路

- 開啟命令提示字元（%WINDIR%\system32\cmd.exe）

![](../images/161.png)

- 輸入下方指令，在建立一資料夾後，編輯 %GOPATH%\src\github.com\username\myfirstproject\client.go 檔。

```
> md %GOPATH%\src\github.com\username\myfirstproject
> code %GOPATH%\src\github.com\username\myfirstproject\client.go
```

![](../images/162.png)

- 使用 Visual Studio Code 編輯【client.go】檔，輸入以下內容後按下【Ctrl】+【S】鍵儲存。

```
package main

import (
        "fmt"
        "github.com/ethereum/go-ethereum/ethclient"
        "log"
)

func main() {
        //client, err := ethclient.Dial("https://mainnet.infura.io")
        client, err := ethclient.Dial("https://rinkeby.infura.io")
        if err != nil {
                log.Fatal(err, "\n")
        }
        fmt.Println("we have a connection")
        _ = client
}
```

![](../images/163.png)

- 編譯 %GOPATH%\src\github.com\username\myfirstproject\client.go 後執行，確認可順利連至以太坊測試網路

```
> go run github.com\username\myfirstproject

//-------------------------Content-------------------------
we have a connection
//-------------------------Content-------------------------
```

![](../images/164.png)

### （8-2）使用 Visual Studio Code 編譯 go 語言並連接至以太坊測試網路

- 建立資料夾「%GOPATH%\src\github.com\username\myfirstproject」

![](../images/171.png)

![](../images/172.png)

- 使用 Visual Studio Code 開啟「%GOPATH%\src\github.com\username\myfirstproject」。

![](../images/173.png)

- 新增【client.go】檔，

![](../images/174.png)

- 輸入以下內容後按下【Ctrl】+【S】鍵儲存。

```
package main

import (
        "fmt"
        "github.com/ethereum/go-ethereum/ethclient"
        "log"
)

func main() {
        //client, err := ethclient.Dial("https://mainnet.infura.io")
        client, err := ethclient.Dial("https://rinkeby.infura.io")
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("we have a connection")
        _ = client
}
```

![](../images/175.png)

- 點選【偵錯 Debug】→【啟動偵錯 Start Debugging（F5）】即可開始編譯並在下方看到執行結果。
> 註：未來可以直接按【F5】鍵來進行「啟動偵錯 Start Debugging」

![](../images/176.png)

![](../images/177.png)

### （9）建立 Solidity 編譯環境。

- 先至微軟官網 [https://www.microsoft.com/zh-TW/download/details.aspx?id=53587](https://www.microsoft.com/zh-TW/download/details.aspx?id=53587)，並點選【下載】→ 勾選【vc_redist.x64.exe】→【Next】下載 Solidity Compiler 所需的 `Microsoft Visual C++ 2015 可轉散髮套件`（預設是下載到 `%USERPROFILE%\Downloads` 資料夾內）。

![](../images/181.png)

![](../images/182.png)

- 檔案下載後，滑鼠連點兩下執行安裝檔。

![](../images/183.png)

- 點選【我同意授權條款及條件(A)】→【安裝(I)】。

![](../images/184.png)

- 點選【關閉(C)】完成安裝。

![](../images/185.png)
