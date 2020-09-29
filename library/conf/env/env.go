package env

import (
	"os"
	"strings"
)

// env configuration.
var (
	// 链路追踪地址
	TracerAddress string
	// 服务注册中心
	RegistryAddress []string
)

func init(){
	TracerAddress = os.Getenv("MICRO_TRACER_ADDRESS")

	registryAddress := os.Getenv("MICRO_REGISTRY_ADDRESS")
	if registryAddress != ""{
		RegistryAddress = strings.Split(registryAddress, ",")
	}
}
