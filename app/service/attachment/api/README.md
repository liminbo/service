# Demo

##生成proto文件
```
protoc --proto_path=. --micro_out=. --go_out=. api.proto
```

##生成文档命令（文档目录：app/apidoc/v1/）
```
protoc --doc_out=./../../../../apidoc/v1 --doc_opt=html,attachment.html *.proto
```
