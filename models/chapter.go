package models

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Url    string `json:"url"`
}
