
.PHONY: proto
proto:
	protoc -I .  -I ../ -I ../../../../../../third_party/protobuf  -I ../../../../../../third_party/protobuf/googleapis --go_out=Mattachment.proto=service.yidoutang.com/app/service/attachment/api/v1,plugins=grpc:. --grpc-gateway_out=Mattachment.proto=service.yidoutang.com/app/service/attachment/api/v1,logtostderr=true:. --swagger_out=logtostderr=true:. attachment_gateway.proto