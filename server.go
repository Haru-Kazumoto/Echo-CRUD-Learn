package main

import (
	"myapp/src/db"
	"myapp/src/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	e := echo.New()

	e.GET("/users", func(c echo.Context) error {
		var users []model.User
		db.Find(&users)
		return c.JSON(http.StatusOK, users)
	})
	
	e.Logger.Fatal(e.Start(":8000"))
}