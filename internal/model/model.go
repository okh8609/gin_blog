package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/okh8609/gin_blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
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

	sqlDB, err := sql.Open(c.DBType, conn_str)
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

	gormDB.Callback().Create().After("gorm:create").Register("time_stamp", createTimestampCallback)
	gormDB.Callback().Update().After("gorm:update").Register("time_stamp", updateTimestampCallback)
	gormDB.Callback().Delete().Replace("gorm:delete", deleteCallback)

	return gormDB, nil
}

// 處理資料表內共用(重複)(相同)(blog_basic_info)的欄位
func createTimestampCallback(db *gorm.DB) {
	if db.Error == nil && db.Statement.Schema != nil {
		nowTime := time.Now().Unix()

		if field := db.Statement.Schema.LookUpField("CreatedOn"); field != nil {
			field.Set(db.Statement.ReflectValue, nowTime)
		}

		if field := db.Statement.Schema.LookUpField("ModifiedOn"); field != nil {
			field.Set(db.Statement.ReflectValue, nowTime)
		}
	}
}

func updateTimestampCallback(db *gorm.DB) {
	if db.Error == nil && db.Statement.Schema != nil {
		if field := db.Statement.Schema.LookUpField("ModifiedOn"); field != nil {
			field.Set(db.Statement.ReflectValue, time.Now().Unix())
		}
	}
}

func deleteCallback(db *gorm.DB) {
	if db.Error == nil {
		deletedOnField := db.Statement.Schema.LookUpField("DeletedOn")
		isDelField := db.Statement.Schema.LookUpField("IsDel")

		if deletedOnField != nil && isDelField != nil {
			// TODO: 軟刪除
			// db.Exec(fmt.Sprintf(
			// 	"UPDATE %v SET `DeletedOn`=%v,`IsDel`=%v WHERE %v",
			// 	scope.QuotedTableName(),

			// 	time.Now().Unix(),
			// 	1,

			// 	addExtraSpaceIfExist(extraOption),
			// ))
		} else {
			callbacks.Delete(db)
		}
	}
}

// func addExtraSpaceIfExist(str string) string {
// 	if str != "" {
// 		return " " + str
// 	}
// 	return ""
// }
