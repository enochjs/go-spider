package models

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Name   string `json:"name"`
	Index  int `json:"index"`
	Remark string `json:"remark"`
	Url    string `json:"url"`
	BookId int  	`json:"bookId"`
}
