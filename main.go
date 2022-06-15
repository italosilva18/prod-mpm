package main

import (
	"github.com/italosilva18/prod-mpm/repository"
	"github.com/italosilva18/prod-mpm/router"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	router.Root(e)

	repository.ConnectDB()

	e.Logger.Fatal(e.Start(":1323"))
}
