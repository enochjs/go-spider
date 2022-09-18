package routes

import (
	"enochjs.github.com/go-spider/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	//r.LoadHTMLGlob("templates/*")
	//
	//r.Static("/static", "static")
	//
	//r.GET("/", controller.IndexHtml)
	//
	//book := r.Group("/api/book")
	//book.POST("/add", controller.AddBook)

	site := r.Group("/api/site")
	site.POST("/add", controller.AddSite)

	//r.NoRoute(controller.IndexHtml)

	return r

}
