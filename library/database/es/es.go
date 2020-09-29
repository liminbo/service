package es

import (
	"fmt"
	"github.com/micro/go-micro/v2/logger"
	elastic "gopkg.in/olivere/elastic.v5"
	"time"
)

type Config struct {
	Addr          string          // write data source name.
	Timeout time.Duration
}

type ElasticSearch struct {
	Client *elastic.Client
	Timeout time.Duration
}

type ElasticLog struct {

}

func (e *ElasticLog) Printf(format string, v ...interface{}){
	logger.Infof(format, v...)
}

func New(c *Config) (es *ElasticSearch){

	client, err := elastic.NewClient(elastic.SetURL(c.Addr), elastic.SetSniff(false))
	// 调试开发可以把这个方法替换成上面那个，然后就可以开启把elasticSearch的请求参数和结果打印到日志里
	//esLog := &ElasticLog{}
	//client, err := elastic.NewClient(elastic.SetURL(c.Addr), elastic.SetSniff(false),elastic.SetTraceLog(esLog))

	if err != nil {
		panic(fmt.Sprintf("es:集群连接失败, cluster: %v", err))
	}
	es = &ElasticSearch{
		Client:  client,
		Timeout: time.Duration(c.Timeout),
	}
	return
}