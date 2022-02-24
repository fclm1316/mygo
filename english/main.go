package main

import (
	"mygo/english/models"
	"mygo/english/services"
)

func main() {
	models.InitDB()
	services.StartApp()
}
