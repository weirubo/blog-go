package model

import (
	"fmt"
	"github.com/weirubo/blog-go/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 创建数据库
func NewDBEngine(databaseConfig *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		databaseConfig.UserName,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.DBName,
		databaseConfig.Charset,
		databaseConfig.ParseTime,
		databaseConfig.Loc)), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   databaseConfig.TablePrefix,
		SingularTable: true,
	}})
	if err != nil {
		fmt.Println("NewDBEngine error:", err)
		return nil, err
	}
	return db, nil
}
