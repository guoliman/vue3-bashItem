# 生成二进制包

# 一行实现
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o vue3-bashItem -x main.go

# 多行实现
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o vue3-bashItem -x main.go

