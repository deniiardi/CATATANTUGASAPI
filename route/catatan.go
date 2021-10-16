package route

import (
	"CatatanTugas/controller"
	"CatatanTugas/middleware"

	"github.com/labstack/echo"
)

func NewCatatans(e *echo.Echo) {
	e.GET("/catatans", controller.GetAllCatatansController, middleware.AuthJWT)
	e.POST("/catatans", controller.CreateCatatanController, middleware.AuthJWT)
	e.GET("/catatans/:id", controller.GetCatatanByIDController, middleware.AuthJWT)
	e.DELETE("/catatans/:id", controller.DeleteCatatanByIDController, middleware.AuthJWT)
	e.PUT("/catatans/:id", controller.UpdateCatatanByIDController, middleware.AuthJWT)
}
