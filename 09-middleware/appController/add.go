package appController

import (
	"fmt"
	"gofrendi/structureExample/arithmetic"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Add(c echo.Context) error {
	angka1, err := strconv.Atoi(c.Param("angka1"))
	if err != nil {
		return c.String(http.StatusBadRequest, "angka1 tidak valid")
	}
	angka2, err := strconv.Atoi(c.Param("angka2"))
	if err != nil {
		return c.String(http.StatusBadRequest, "angka2 tidak valid")
	}
	result := fmt.Sprintf("%d", arithmetic.Add(angka1, angka2))
	return c.String(http.StatusOK, result)
}
