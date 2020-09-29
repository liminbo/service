package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"micro-service/library/database/mysql"
)


// NewMySQL new db and retry connection when has error.
func NewMySQL(c *mysql.Config) (db *gorm.DB) {

	dsn := fmt.Sprintf(
		"%s:%v@tcp(%s:%v)/%s?charset=%s&parseTime=True&loc=Local",
		c.User,c.Password, c.Host,c.Port, c.Database, c.Charset)
	db, err := gorm.Open("mysql", dsn)

	// 数据库链接不上，直接报错
	if err !=nil{
		panic(err)
	}

	db.DB().SetMaxIdleConns(c.IdleConn)
	db.DB().SetMaxOpenConns(c.MaxConn)
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return c.TablePrefix + defaultTableName
	}
	return db
}
