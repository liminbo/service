package http

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"google.golang.org/grpc"
	"micro-service/app/service/attachment/api"
	gw "micro-service/app/service/attachment/api/v1/gateway"
	"micro-service/app/service/attachment/internal/gomicro"
	"micro-service/app/service/attachment/internal/service"
	"micro-service/library/http"
	"sync"
)

var svr *service.Service

// New new a bm server.
func New(s *service.Service, waiter *sync.WaitGroup) (err error) {
	service := web.NewService(
		web.Name(api.WebAppID),
	)

	// 接入pprof
	http.StartPerf(service)
	// grpc-gateway start
	ctx := context.Background()
	// Register gRPC server endpoint
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err = gw.RegisterAttachmentServiceHandlerFromEndpoint(ctx, mux,  gomicro.MicroService.Server().Options().Address, opts)
	if err != nil {
		logger.Errorf("grpc-gateway init err:%v", err)
	}
	service.Handle("/", mux)
	// grpc-gateway end

	if err := service.Init(); err != nil {
		logger.Errorf("service init err:%v", err)
	}

	if err := service.Init(web.Address(":8904")); err != nil {
		logger.Errorf("service init err:%v", err)
	}

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		if err := service.Run(); err != nil {
			logger.Errorf("service run err:%v", err)
		}
	}()

	svr = s
	return
}
