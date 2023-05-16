package main

import (
	"myapp/src/db"
	"myapp/src/model"

	"github.com/labstack/echo/v4"
)

func main() {

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	e := echo.New()

	e.Logger.Fatal(e.Start(":8000"))
}