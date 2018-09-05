# golang短网址

## 启动
```
go get -u github.com/convee/shorturl

cd $GOPATH/src/github.com/convee/shorturl

go build -o shorturl //构建

./shorturl  -addr 127.0.0.1:8002 //监听本地8002端口

```
## 生成短网址

```
curl http://127.0.0.1:8002/gen?longurl=http://convee.cn

{
    "shorturl": "2xZj"
    "longurl": "http://convee.cn"
}
```


## 短网址访问