package di

import (
	"micro-service/app/service/attachment/internal/dao"
	"micro-service/app/service/attachment/internal/gomicro"
	"micro-service/app/service/attachment/internal/server/grpc"
	"micro-service/app/service/attachment/internal/server/http"
	"micro-service/app/service/attachment/internal/service"
	"sync"
)

var Waiter sync.WaitGroup

// 初始化相关资源，并且注入
func InitApp() (func(), error) {

	db, closeDB, err := dao.NewDB()
	if err != nil {
		return nil, err
	}

	daoDao, closeDao, err := dao.New(db)
	if err != nil {
		closeDB()
		return nil, err
	}

	serviceService, closeService, err := service.New(daoDao)
	if err != nil {
		closeDao()
		closeDB()
		return nil, err
	}

	err = grpc.New(gomicro.MicroService, serviceService, &Waiter)
	if err != nil {
		closeService()
		closeDao()
		closeDB()
		return nil, err
	}

	err = http.New(serviceService, &Waiter)
	if err != nil {
		closeService()
		closeDao()
		closeDB()
		return nil, err
	}

	return func() {
		closeService()
		closeDao()
		closeDB()
	}, nil
}
