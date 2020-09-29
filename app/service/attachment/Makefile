
.PHONY: proto
proto:
	cd ./api/v1 && make proto

.PHONY: build
build:
	cd ./cmd && go build -o ./service-attachment ./main.go

.PHONY: pm2_run
pm2_run:
	cd ./cmd && pm2 reload pm2.json