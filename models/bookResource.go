package models

import "gorm.io/gorm"

// 资源对应站点url

type BookResource struct {
	gorm.Model
	Url string `json:"url"`
	//SiteId uint   `json:"siteId"`
	//Site   Site   `json:"site"`
	Remark string `json:"remark"`
	BookId uint   `json:"bookId"`
}
