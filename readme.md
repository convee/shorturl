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
curl http://127.0.0.1:8002/genUrl?longurl=https://convee.cn

```

- 成功：
```
{
    "error": 0,
    "msg": "ok",
    "data": {
        "longurl": "https://convee.cn",
        "shorturl": "abc"
    }
}
```
- 失败：
```
{
    "error": 1,
    "msg": "no result",
    "data": null
}
```


## 根据短网址获取长网址

```
curl http://127.0.0.1:8002/getUrl?shorturl=abc1

```

- 成功：
```
{
    "error": 0,
    "msg": "ok",
    "data": {
        "longurl": "https://convee.cn",
        "shorturl": "abc"
    }
}
```
- 失败：
```
{
    "error": 1,
    "msg": "no result",
    "data": null
}
```