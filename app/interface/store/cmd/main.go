package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	//"micro-service/app/service/attachment/api"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"micro-service/app/interface/store/middleware"
	"micro-service/library/conf/env"
	pprof "micro-service/library/http"
	"micro-service/library/net/trace"
	"strings"
	"time"
)

var (
	cli client.Client
)

func main() {
	service := web.NewService(
		web.Name("go.micro.web.store"),
	)
	cli = service.Options().Service.Client()

	// 接入pprof
	pprof.StartPerf(service)

	if err := service.Init(); err != nil {
		logger.Errorf("service init err:%v", err)
	}

	if err := service.Init(web.Address("127.0.0.1:8080")); err != nil {
		logger.Errorf("service init err:%v", err)
	}

	// 链路追踪 start
	t, closer, err := trace.NewJaegerTracer("go.micro.api.store", env.TracerAddress)
	if err != nil {
		logger.Fatalf("opentracing tracer create error:%v", err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(t)
	// 链路追踪 end

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery(), middleware.TracerWrapper())

	r := router.Group("/user")
	r.GET("/bo", boHandlefunc)

	service.Handle("/", router)

	if err := service.Run(); err != nil {
		logger.Errorf("service run err:%v", err)
	}
}

func boHandlefunc(ctx *gin.Context) {
	//c, _ := middleware.ContextWithSpan(ctx)
	//svr := v1.NewStoreService(api.AppID, cli)
	//rep, _ := svr.OfflineStoreDetail(c, &v1.StoreDetailReq{StoreId: 358520})
	//str := fmt.Sprintf("data:%v", rep)
	//ctx.String(200, str)
}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
