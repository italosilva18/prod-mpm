package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var Users *userRouting

type userRouting struct {
}

func (t *userRouting) Sim(c echo.Context) error {
	return c.JSON(http.StatusOK, "ponteiro funcionando")
}

var Get *getRouting

type getRouting struct {
}

func (t *getRouting) Novo(c echo.Context) error {
	return c.String(http.StatusOK, "ola")
}
