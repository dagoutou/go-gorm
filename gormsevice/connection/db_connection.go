package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() (db *gorm.DB) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:12345678@tcp(127.0.0.1:3306)/go-web?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	return db
}
