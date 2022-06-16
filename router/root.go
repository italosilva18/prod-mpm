package router

import (
	"github.com/italosilva18/prod-mpm/controllers"
	"github.com/labstack/echo/v4"
)

func Root(e *echo.Echo) {

	e.POST("/add", controllers.Users.Add)
	e.GET("/show:id", controllers.Users.Show)
	e.PUT("/update:id", controllers.Users.Update)
	e.DELETE("/delete:id", controllers.Users.Delete)
	e.GET("/pont", controllers.Get.Novo)

}
