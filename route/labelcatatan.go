package route

import (
	"CatatanTugas/controller"
	"CatatanTugas/middleware"

	"github.com/labstack/echo"
)

func NewLabelCatatan(e *echo.Echo) {
	e.GET("/labelcatatans", controller.GetAllLabelCatatansController, middleware.AuthJWT)
	e.POST("/labelcatatans", controller.CreateLabelCatatanController, middleware.AuthJWT)
	e.GET("/labelcatatans/:id", controller.GetLabelCatatanByIDController, middleware.AuthJWT)
	e.DELETE("/labelcatatans/:id", controller.DeleteLabelCatatanByIDController, middleware.AuthJWT)
	e.PUT("/labelcatatans/:id", controller.UpdateLabelCatatanByIDController, middleware.AuthJWT)
}
