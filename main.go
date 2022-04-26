package main

import (
	"github.com/nao11aihara/echo-api-sample/app/controllers"
	"github.com/nao11aihara/echo-api-sample/app/models"
)

func main() {
	models.ConnectDb()
	controllers.ConfigServer()
}
