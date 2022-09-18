package controller

import (
	"enochjs.github.com/go-spider/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddBook(ctx *gin.Context) {
	var book models.Book

	ctx.BindJSON(&book)

	err := models.CreateBook(&book)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    book,
		})
	}
}
