package controller

import (
	"fmt"
	"net/http"

	"gofrendi/structureExample/model"

	"github.com/labstack/echo/v4"
)

type PersonModel interface {
	GetAll() ([]model.Person, error)
	Add(model.Person) (model.Person, error)
}

type PersonController struct {
	model PersonModel
}

func NewPersonController(m PersonModel) PersonController {
	return PersonController{model: m}
}

func (pc PersonController) GetAll(c echo.Context) error {
	allPersons, err := pc.model.GetAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot get persons")
	}
	return c.JSON(http.StatusOK, allPersons)
}

func (pc PersonController) Add(c echo.Context) error {
	var person model.Person
	if err := c.Bind(&person); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "invalid person structure")
	}
	person, err := pc.model.Add(person)
	if err != nil {
		return c.String(http.StatusInternalServerError, "cannot save person")
	}
	return c.JSON(http.StatusOK, person)
}
