all: 
	@echo "make build|start|stop"
build:
	go build -o shorturl
	@echo "构建成功"
start:
	nohup ./shorturl 2>&1 >> shorturl.log 2>&1 /dev/null &
	@echo "服务已启动"
stop:
	killall shorturl
	@echo "服务已停止"
help:
	@echo "make build - 构建"
	@echo "make start - 启动服务"
	@echo "make stop - 停止服务"