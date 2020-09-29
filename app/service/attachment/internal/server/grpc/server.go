package grpc

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	v1 "micro-service/app/service/attachment/api/v1"
	"micro-service/app/service/attachment/internal/service"
	"sync"
	"time"
)

// New new a grpc server.
func New(microService micro.Service, svc *service.Service, waiter *sync.WaitGroup) (err error) {
	if err = v1.RegisterAttachmentServiceHandler(microService.Server(), svc); err != nil {
		return
	}

	waiter.Add(1)
	go func() {
		defer waiter.Done()
		if err := microService.Run(); err != nil {
			logger.Errorf("service run err:%v", err)
		}
	}()
	time.Sleep(time.Second*2)
	return
}
