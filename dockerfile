# 指定 Docker 映像檔的基本映像檔
FROM golang:1.19.1

# 在容器中建立一個目錄，用於放置應用程式的檔案
WORKDIR /app

# 將本機目錄中的所有檔案複製到容器中的 /app 目錄中
COPY . .

# 下載應用程式的相依套件
RUN go mod download

# 指定容器啟動時要執行的命令
CMD ["go", "run", "main.go"]
