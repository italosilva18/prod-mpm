package routing

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var Get *getRouting

type getRouting struct {
}

func (t *getRouting) Novo(c echo.Context) error {
	return c.JSON(http.StatusOK, "ponteiro funcionando")
}
