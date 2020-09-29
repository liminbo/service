package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/logger"
)

// dao dao.
type Dao struct {
	db *gorm.DB
}

// New new a dao and return.
func New(db *gorm.DB) (d *Dao, cf func(), err error) {
	d = &Dao{
		db: db,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *Dao) Close() {
	logger.Info("close dao")
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}
