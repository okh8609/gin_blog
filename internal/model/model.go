package model

import (
	"database/sql"
	"fmt"

	"github.com/okh8609/gin_blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBEngine(c *setting.DatabaseSetting) (*gorm.DB, error) {
	conn_str := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		c.UserName,
		c.Password,
		c.Host,
		c.DBName,
		c.Charset,
		c.ParseTime,
	)

	sqlDB, err := sql.Open(c.DBName, conn_str)
	if err != nil {
		return nil, err
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err1 := gormDB.DB()
	if err1 != nil {
		return nil, err1
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	return gormDB, nil
}
