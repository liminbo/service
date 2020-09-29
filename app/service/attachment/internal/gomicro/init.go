package gomicro

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-micro/v2/registry"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"micro-service/app/service/attachment/api"
	"micro-service/app/service/attachment/conf"
)

var MicroService micro.Service

func InitMicroService() micro.Service {
	MicroService = micro.NewService(
		micro.Name(api.AppID),
		micro.Version("latest"),

		micro.Flags(
			&cli.StringFlag{
				Name:  "conf_path",
				Value: "./",
				Usage: "配置文件目录",
			},
		),
		micro.Action(func(ctx *cli.Context) (err error) {
			confPath := ctx.String("conf_path")
			logger.Infof("config:%v", confPath)
			// 加载配置文件
			conf.SetPath(confPath)
			return err
		}),
	)

	MicroService.Init()
	return MicroService
}
