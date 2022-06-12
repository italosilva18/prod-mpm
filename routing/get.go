package routing

import (
	"github.com/italosilva18/prod-mpm/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	e.GET("/pont", controllers.Get.Novo)
	e.GET("/2", controllers.Users.Sim)
}