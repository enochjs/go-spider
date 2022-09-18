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
	fmt.Printf("create site 1= %v", site)

	site1 := Site{
		Name:   "enochjs",
		Url:    "1222",
		Remark: "string",
		Script: "test",
	}
	err = DB.Debug().Create(&site1).Error
	return err
}
