package main

import (
	"enochjs.github.com/go-spider/models"
	"enochjs.github.com/go-spider/routes"
)

func main() {

	models.InitDb()

	r := routes.InitRouter()

	r.Run(":9001")
}
