package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name         string         `json:"name"`
	Remark       string         `json:"remark"`
	BookResource []BookResource `json:"bookResource"`
}

func CreateBook(book *Book) (err error) {
	err = DB.Create(&book).Error
	return err
}

func Get() {

}
