
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. *.proto
	protoc --doc_out=./../../../../apidoc/v1 --doc_opt=html,attachment.html *.proto
	cd ./gateway && make proto