all: 
	@echo "make build|start|status|stop"
build:
	go build -o shorturl
	@echo "构建成功"
start:
	nohup ./shorturl 2>&1 >> shorturl.log 2>&1 /dev/null &
	@echo "服务已启动"
status:
	@echo "查看进程"
	ps -ef | grep -w shorturl | grep -v 'grep'
stop:
	killall shorturl
	@echo "服务已停止"
help:
	@echo "make build - 构建"
	@echo "make start - 启动服务"
	@echo "make status - 查看进程"
	@echo "make stop - 停止服务"