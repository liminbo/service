# proto相关命令操作

##生成proto文件
```
protoc -I .  -I ../ -I ../../../../../../third_party/protobuf  -I ../../../../../../third_party/protobuf/googleapis --go_out=Mrequest.proto=service.yidoutang.com/app/service/mtest/api/v1,Mresponse.proto=service.yidoutang.com/app/service/mtest/api/v1,Mfolder.proto=service.yidoutang.com/app/service/mtest/api/v1,plugins=grpc:. --grpc-gateway_out=Mrequest.proto=service.yidoutang.com/app/service/mtest/api/v1,Mresponse.proto=service.yidoutang.com/app/service/mtest/api/v1,Mfolder.proto=service.yidoutang.com/app/service/mtest/api/v1,logtostderr=true:. --swagger_out=logtostderr=true:. gateway.proto
```