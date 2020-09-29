package grpc

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"micro-service/library/conf/env"
)

type ServerConfig struct {
	Addr string
}

type ClientConfig struct {
	Addr string
}

func NewClient() (c client.Client){
	c = grpc.NewClient()
	err := c.Init(client.Registry(etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = env.RegistryAddress
	})))
	if err != nil{
		panic(err)
	}
	return
}