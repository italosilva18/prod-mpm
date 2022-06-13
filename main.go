package main

import (
	"github.com/italosilva18/prod-mpm/config"
	"github.com/italosilva18/prod-mpm/router"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	config.ConnectDB()

	router.Root(e)

	e.Logger.Fatal(e.Start(":1323"))
}
