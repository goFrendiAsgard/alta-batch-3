package appRoute

import (
	"gofrendi/structureExample/appController"
	"gofrendi/structureExample/appMiddleware"
	"gofrendi/structureExample/appModel"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func HandleRoutes(e *echo.Echo, personModel appModel.PersonModel) {
	e.GET("/", appController.Hello)
	e.GET("/add/:angka1/:angka2/", appController.Add)
	e.GET("/add/:angka1/:angka2", appController.Add)

	personController := appController.NewPersonController(personModel)
	e.POST("/persons", personController.Add)
	e.POST("/persons/", personController.Add)

	eAuth := e.Group("")
	// curl --location --request GET 'localhost:8080/persons' \
	// --header 'Authorization: Basic YWRtaW46YWRtaW4='

	// eAuth.Use(middleware.BasicAuth(appMiddleware.DummyBasicAuth))
	eAuth.Use(middleware.BasicAuth(appMiddleware.MakePersonBasicAuth(personModel)))

	eAuth.GET("/persons", personController.GetAll)
	eAuth.GET("/persons/", personController.GetAll)
	eAuth.PUT("/persons/:id", personController.Edit)
	eAuth.PUT("/persons/:id/", personController.Edit)
}
