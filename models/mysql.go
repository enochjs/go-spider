package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

func InitDb() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/spider?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "spider_",
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	//迁移建表
	DB.AutoMigrate(&Site{})
	DB.AutoMigrate(&Book{})
	DB.AutoMigrate(&BookResource{})
	DB.AutoMigrate(&Chapter{})

}
