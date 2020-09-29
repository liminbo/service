module micro-service

go 1.14

require (
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/envoyproxy/protoc-gen-validate v0.4.1 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.8+incompatible // indirect
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/iancoleman/strcase v0.1.1 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.10
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/logger/zap/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/onsi/ginkgo v1.12.3 // indirect
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cobra v1.0.0 // indirect
	github.com/spf13/viper v1.7.0
	github.com/streadway/amqp v1.0.0 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/uber/jaeger-client-go v2.24.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/sys v0.0.0-20200602225109-6fdc65e7d980 // indirect
	golang.org/x/tools v0.0.0-20200903185744-af4cc2cd812e // indirect
	google.golang.org/genproto v0.0.0-20200605102947-12044bf5ea91
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/olivere/elastic.v5 v5.0.85
)

replace google.golang.org/grpc v1.27.0 => google.golang.org/grpc v1.26.0
