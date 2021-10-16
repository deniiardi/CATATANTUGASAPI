package route

import (
	"CatatanTugas/controller"
	"CatatanTugas/middleware"

	"github.com/labstack/echo"
)

func NewGambars(e *echo.Echo) {
	e.GET("/catatans/gambars", controller.GetAllGambarsController, middleware.AuthJWT)
	e.POST("/catatans/gambars", controller.CreateGambarController, middleware.AuthJWT)
	e.GET("/catatans/gambars/:id", controller.GetGambarByIDController, middleware.AuthJWT)
	e.DELETE("/catatans/gambars/:id", controller.DeleteGambarByIDController, middleware.AuthJWT)
	e.PUT("/catatans/gambars/:id", controller.UpdateGambarByIDController, middleware.AuthJWT)
}
