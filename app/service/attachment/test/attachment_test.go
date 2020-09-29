package test

import (
	"context"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"micro-service/app/service/attachment/api"
	v1 "micro-service/app/service/attachment/api/v1"
	"testing"
)
func Benchmark_Detail(b *testing.B) {
	var n int
	ctx := context.Background()
	microSerice := micro.NewService()
	attachmentService := v1.NewAttachmentService(api.AppID, microSerice.Client())
	for i := 0; i < b.N; i++ {
		rep, _ := attachmentService.AttachmentDetailByIds(ctx, &v1.AttachmentDetailByIdsReq{
			Ids:                  []int32{9382},
		})
		logger.Infof("data:%v", rep)
		n++
	}
}
