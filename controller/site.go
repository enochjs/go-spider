package controller

import (
	"enochjs.github.com/go-spider/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddSite(ctx *gin.Context) {
	println("come in")
	var site models.Site

	ctx.BindJSON(&site)

	//println("site", site)
	fmt.Printf("site=%v", site)

	err := models.CreateSite(&site)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    site,
		})
	}
}
