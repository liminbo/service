package test

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"micro-service/app/service/attachment/api"
	"github.com/micro/go-micro/v2/registry"
	consul "github.com/micro/go-plugins/registry/consul/v2"
	v1 "micro-service/app/service/attachment/api/v1"
	"testing"
)
func Benchmark_Detail(b *testing.B) {
	var n int
	ctx := context.Background()
	microSerice := micro.NewService(
		micro.Registry(consul.NewRegistry(func(options *registry.Options) {
			options.Addrs = []string{"127.0.0.1:8500"}
		})),
		)
	attachmentService := v1.NewAttachmentService(api.AppID, microSerice.Client())
	for i := 0; i < b.N; i++ {
		rep, _ := attachmentService.AttachmentDetailByIds(ctx, &v1.AttachmentDetailByIdsReq{
			Ids:                  []int32{9382},
		})
		logger.Infof("data:%v", rep)
		n++
	}
}
