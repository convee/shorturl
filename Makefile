TIME=$(shell date +"%Y%m%d")
GIT_REVISION=$(shell git show -s --pretty=format:%h)
GO=/usr/local/go/bin/go
PATH=~/bin/shorturl/
all: 
	@echo "make help|build-linux|build-mac|build-windows"
build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(PATH)linux.shorturl.$(TIME).$(GIT_REVISION) -a -v main.go
	@echo "构建成功"
build-mac:
	GOOS=darwin GOARCH=amd64 $(GO) build -o $(PATH)mac.shorturl.$(TIME).$(GIT_REVISION) -a -v main.go
	@echo "构建成功"
build-windows:
	GOOS=windows GOARCH=amd64 $(GO) build -o $(PATH)windows.shorturl.$(TIME).$(GIT_REVISION).exe -a -v main.go
	@echo "构建成功"
help:
	@echo "make build-linux - 构建linux"
	@echo "make build-mac - 构建mac"
	@echo "make build-windows - 构建windows"
	@echo "make all - 查看所有命令"