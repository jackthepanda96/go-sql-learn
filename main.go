package main

import (
	"fmt"
	"todo/configs"
	"todo/internal/controllers"
	"todo/internal/models"
)

func main() {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}

	um := models.NewUserModel(connection)
	uc := controllers.NewUserController(um)

	uc.Login()

}
