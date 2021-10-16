package route

import (
	"CatatanTugas/controller"
	"CatatanTugas/middleware"

	"github.com/labstack/echo"
)

func NewPengingats(e *echo.Echo) {
	e.GET("/catatans/pengingats", controller.GetAllPengingatsController, middleware.AuthJWT)
	e.POST("/catatans/pengingats", controller.CreatePengingatController, middleware.AuthJWT)
	e.GET("/catatans/pengingats/:id", controller.GetPengingatByIDController, middleware.AuthJWT)
	e.DELETE("/catatans/pengingats/:id", controller.DeletePengingatByIDController, middleware.AuthJWT)
	e.PUT("/catatans/pengingats/:id", controller.UpdatePengingatByIDController, middleware.AuthJWT)
}
