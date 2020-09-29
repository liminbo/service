package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/opentracing/opentracing-go"
	"micro-service/app/service/attachment/api"
	conf "micro-service/app/service/attachment/conf"
	"micro-service/app/service/attachment/internal/di"
	"micro-service/app/service/attachment/internal/gomicro"
	"micro-service/library/conf/env"
	"micro-service/library/log"
	"micro-service/library/net/rpc"
	"micro-service/library/net/trace"
	"time"

	tracer "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

func main() {
	// 初始化 micro service
	microService := gomicro.InitMicroService()

	// 初始化配置
	conf.Init()

	// 切换到zap日志组件 start
	outputPath := []string{"stderr"}
	logger.Infof("application.log_path:%s", conf.Conf.LogPath)
	if conf.Conf.LogPath != "" {
		outputPath = append(outputPath, conf.Conf.LogPath)
	}
	logger.DefaultLogger = log.NewZap(outputPath)
	// 切换到zap日志组件 end

	// 链路追踪 start
	t, closer, err := trace.NewJaegerTracer(api.AppID, env.TracerAddress)
	if err != nil {
		logger.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(t)
	// 链路追踪 end

	microService.Init(
		//挂载链路
		micro.WrapHandler(tracer.NewHandlerWrapper(t)),
		micro.WrapCall(tracer.NewCallWrapper(t)),
		micro.WrapHandler(rpc.LogHandler()),
	)

	// 初始化app
	logger.Infof("%v start", conf.Conf.ServiceName)
	closeFunc, err := di.InitApp()
	if err != nil {
		panic(err)
	}

	// 等待所有任务完成，优雅退出
	di.Waiter.Wait()
	closeFunc()
	time.Sleep(time.Second * 1)
}
