# feastogether

### 介绍 ✨

歡迎來到 feastogether，此作品提供饗賓餐旅線上快速訂位，本搶定位程式僅供學術和研究目的使用，開發者不對其准確性、可靠性、完整性、合法性以及使用者使用本程式產生的任何後果承擔責任。

請勿將本程式用於非法用途或從事任何違法活動( EX : 蝦皮販賣定位轉讓 )，任何因使用本程式而從事違法活動所導致的法律責任由使用者自行承擔，開發者概不負責。

### 開始 ✨

```bash
git clone https://github.com/zhong1016/feastogether.git
cd ./feastogether
```

修改 config.ini 相關資料

執行程式 :

```bash
go run main.go
```

### 目錄結構 ✨

```text
- client        # booking 程式
- config        # 讀取 ini
- fetch         # 封裝 Http request
- main.go       # 專案入口
- config.ini    # 用戶與餐廳資料
- README.md     # 專案說明文檔
- go.mod        # module 與 依賴
- go.sum        # 檢查依賴版本
```

### 最後 ✨

如果此作品有幫助到你/妳，請給我一個 Star！它會使我更有動力 ：）
