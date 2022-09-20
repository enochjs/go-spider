package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Site struct {
	gorm.Model
	Name   string `json:"name"`
	Url    string `json:"url"`
	Remark string `json:"remark"`
	Script string `json:"script"` // 使用枚举
}

func CreateSite(site *Site) (err error) {
	fmt.Printf("create site = %v", site)
	err = DB.Debug().Create(&site).Error
	return err
}
