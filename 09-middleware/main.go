package main

import (
	"fmt"
	"gofrendi/structureExample/config"
	"gofrendi/structureExample/controller"
	"gofrendi/structureExample/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// personModel can be either personMemModel or personDbModel, depends on the configuration
	var personModel model.PersonModel
	switch cfg.Storage {
	case "db":
		db, err := gorm.Open(mysql.Open(cfg.ConnectionString), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		personModel = model.NewPersonDbModel(db)
	case "mem":
		personModel = model.NewPersonMemModel()
	}

	// create new echo instant
	e := echo.New()

	// Log middleware
	//e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.GET("/", controller.Hello)
	e.GET("/add/:angka1/:angka2", controller.Add)

	personController := controller.NewPersonController(personModel)
	e.GET("/persons/", personController.GetAll)
	e.POST("/persons/", personController.Add)
	e.PUT("/persons/:id", personController.Edit)

	if err = e.Start(fmt.Sprintf(":%d", cfg.HttpPort)); err != nil {
		panic(err)
	}
}
