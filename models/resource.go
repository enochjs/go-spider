package models

import "gorm.io/gorm"

// 资源对应站点url，如 海贼王漫画
type Resource struct {
	gorm.Model
	Url string `json:"url"`
	Name string `json:"name"`
	// 执行脚本
	Script string `json:"script"`
	Remark string `json:"remark"`
	// 所属站点
	SiteId int
	Site Site
}
