package main

import (
	"fmt"
	"gofrendi/structureExample/controller"
	"gofrendi/structureExample/model"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Hello world")
	e := echo.New()
	e.GET("/", controller.Hello)
	e.GET("/add/:angka1/:angka2", controller.Add)

	personMemModel := model.NewPersonMemModel()
	personController := controller.NewPersonController(personMemModel)
	e.GET("/person/", personController.GetAll)
	e.POST("/person/", personController.Add)

	e.Start(":8080")

}
