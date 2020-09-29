package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const contextTracerKey = "Tracer-context"

// TracerWrapper tracer 中间件
func TracerWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		sp := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path)
		tracer := opentracing.GlobalTracer()
		md := make(map[string]string)
		md["ydt.userid"] = "2323"

		spanCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err == nil {
			sp = opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
			tracer = sp.Tracer()
		}
		defer sp.Finish()

		if err := tracer.Inject(sp.Context(),
			opentracing.TextMap,
			opentracing.TextMapCarrier(md)); err != nil {
			logger.Errorf("opentracing err:%v", err)
		}

		ctx := context.TODO()
		ctx = opentracing.ContextWithSpan(ctx, sp)
		ctx = metadata.NewContext(ctx, md)
		c.Set(contextTracerKey, ctx)

		c.Next()

		statusCode := c.Writer.Status()
		ext.HTTPStatusCode.Set(sp, uint16(statusCode))
		ext.HTTPMethod.Set(sp, c.Request.Method)
		ext.HTTPUrl.Set(sp, c.Request.URL.EscapedPath())
		if statusCode >= http.StatusInternalServerError {
			ext.Error.Set(sp, true)
		}

	}
}

// ContextWithSpan 返回context
func ContextWithSpan(c *gin.Context) (ctx context.Context, ok bool) {
	v, exist := c.Get(contextTracerKey)
	if exist == false {
		ok = false
		ctx = context.TODO()
		return
	}

	ctx, ok = v.(context.Context)
	return
}
