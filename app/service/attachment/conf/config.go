package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/micro/go-micro/v2/logger"
	"github.com/spf13/viper"
	"os"
	"micro-service/library/database/es"
	"micro-service/library/database/mysql"
	"micro-service/library/net/trace"
)

// Conf global variable.
var (
	Conf     = &Config{}
	confPath string
)

// Config struct of conf.
type Config struct {
	ServiceName string
	LogPath     string
	MySQL       *mysql.Config
	YdtMySQL    *mysql.Config
	Trace       *trace.Config
	Es          *es.Config
}

// Init int config
func Init() {
	// 配置文件名
	configName := "config"
	if env := os.Getenv("MICRO_SERVICE_ENV"); env != ""{
		configName = fmt.Sprintf("%s.%s", configName, env)
	}
	logger.Infof("config name:%s", configName)

	viper.SetConfigName(configName)
	// 配置文件类型
	viper.SetConfigType("toml")
	// 配置文件目录
	viper.AddConfigPath(confPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	if err := viper.Unmarshal(Conf); err != nil {
		logger.Errorf("watch config failed: %v", err)
	}
	viper.OnConfigChange(func(fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			logger.Errorf("OnConfigChange err: %v", err)
		}
	})
}

func SetPath(path string) {
	confPath = path
}
