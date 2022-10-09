package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name         string         `json:"name"`
	Remark       string         `json:"remark"`
	ResourceId 	 int
	Resource Resource `json:"resource"`
}

func CreateBook(book *Book) (err error) {
	err = DB.Create(&book).Error
	return err
}

func Get() {

}
