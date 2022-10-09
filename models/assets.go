package models

import "gorm.io/gorm"

type Assets struct {
	gorm.Model
	Url         string         `json:"url"`
	ChapterId 	int						`json:"chapterId"`
	Index  			int 					`json:"index"`
}