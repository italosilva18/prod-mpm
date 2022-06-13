package router

import (
	"github.com/italosilva18/prod-mpm/controllers"
	"github.com/labstack/echo/v4"
)

func Root(e *echo.Echo) {

	e.POST("/add", controllers.Add)
	e.GET("/show", controllers.Users.Show)
	e.PUT("/update", controllers.Users.update)
	e.DELETE("/delete", controllers.Users.delete)
	e.GET("/pont", controllers.Get.Novo)
	e.GET("/2", controllers.Users.Sim)
}
