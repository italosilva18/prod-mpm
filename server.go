package main

import (
	"net/http"

	"github.com/italosilva18/prod-mpm/routing"
	"github.com/labstack/echo/v4"
)

func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", index)
	e.GET("/pont", routing.Get.Novo)
	e.Logger.Fatal(e.Start(":1323"))
}
