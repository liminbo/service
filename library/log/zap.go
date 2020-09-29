package log

import (
	"github.com/micro/go-micro/v2/logger"
	mzap "github.com/micro/go-plugins/logger/zap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


func NewZap(path []string) logger.Logger{
	// 生产环境配置
	productConfig := zap.NewProductionConfig()
	// 格式化时间
	productConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志输出路径
	productConfig.OutputPaths = path

	// 生成go-micro logger
	log, err := mzap.NewLogger(mzap.WithConfig(productConfig), mzap.WithCallerSkip(2))
	if err != nil{
		panic(err)
	}
	return log
}