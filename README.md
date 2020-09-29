# go微服务



## 环境要求
- Golang >= 1.14
- consul >= 1.8.1
- jaeger >= 1.18

## 安装consul（服务注册中心）

1. 下载链接：https://www.consul.io/downloads.html

2. 启动conul命令：

   ```shell
   ./consul agent -server -data-dir=/data/consul -bootstrap -ui -advertise=127.0.0.1 -client=0.0.0.0
   ```

   访问：http://127.0.0.1:8500/ui/



## 使用docker安装jaeger（链路追踪）

1. 安装命令：

   ```shell
   docker pull jaegertracing/all-in-one
   docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 9411:9411 jaegertracing/all-in-one:latest
   ```

2. 查看链路追踪后台 http://127.0.0.1:16686/search



## 安装proto（用于生成protobuf）

1. 下载链接：https://github.com/protocolbuffers/protobuf/releases

2、下载grpc-gateway: https://github.com/grpc-ecosystem/grpc-gateway/releases/tag/v1.14.5
    把下载好的文件改名为protoc-gen-grpc-gateway

3. 安装go依赖

   ```shell
   go get -u github.com/golang/protobuf/proto
   go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2   // 兼容grpc-gateway版本
   go get -u github.com/micro/micro/v2/cmd/protoc-gen-micro
   // grpc-gateway相关
   go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
   // 根据protobuf生成文档
   go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
   // 参数验证
   go get -u github.com/envoyproxy/protoc-gen-validate@0.4.1
   ```

​   生成的bin 应该存放在$PATH:$GOPATH/bin这里

4. 生成protobuf文件的命令：

   ```shell
   protoc --proto_path=. --micro_out=. --go_out=. api.proto
   ```

### 

## 首次运行

配置环境变量 :

```
export GOROOT=/usr/local/go // go软件包的安装路径
export GOPATH=/data/golang  // go项目存放位置 /data/golang下面需要有 pkg src bin 这三个目录 service项目代码放在src目录下
export PATH=$PATH:$GOPATH/bin // 安装go工具存放的二进制文件
export PATH=$PATH:$GOROOT/bin // go相关二进制文件
export GOPROXY=https://goproxy.cn // 配置go mod 代理，方便下载包
   
export MICRO_REGISTRY=etcd // 使用etcd为注册中心
export MICRO_REGISTRY_ADDRESS=127.0.0.1:2379 // 注册中心地址 多个用逗号隔开
export MICRO_TRACER_ADDRESS=127.0.0.1:6831 // 链路追踪上报地址
export MICRO_SERVICE_ENV=dev // 服务环境(根据环境获取对应的配置文件) dev:本地环境 test:测试环境 prod:生产环境
```



## 其他
bloomrpc（RPC调试工具）：https://github.com/uw-labs/bloomrpc

