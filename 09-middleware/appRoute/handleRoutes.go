package appRoute

import (
	"gofrendi/structureExample/appController"
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
)

func HandleRoutes(e *echo.Echo, personModel appModel.PersonModel) {
	e.GET("/", appController.Hello)
	e.GET("/add/:angka1/:angka2/", appController.Add)
	e.GET("/add/:angka1/:angka2", appController.Add)

	personController := appController.NewPersonController(personModel)
	e.GET("/persons", personController.GetAll)
	e.GET("/persons/", personController.GetAll)
	e.POST("/persons", personController.Add)
	e.POST("/persons/", personController.Add)
	e.PUT("/persons/:id", personController.Edit)
	e.PUT("/persons/:id/", personController.Edit)
}
