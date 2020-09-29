package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/logger"
	"micro-service/app/service/attachment/conf"
	"micro-service/library/database/orm"
)

func NewDB() (db *gorm.DB, cf func(), err error) {
	db = orm.NewMySQL(conf.Conf.MySQL)
	db.LogMode(true)
	cf = func() {
		logger.Info("close DB")
		db.Close()
	}
	return
}
