package route

import (
	"CatatanTugas/controller"
	"CatatanTugas/middleware"

	"github.com/labstack/echo"
)

func NewPengguna(e *echo.Echo) {
	e.GET("/penggunas", controller.GetAllPenggunasController, middleware.AuthJWT)
	e.POST("/penggunas/login", controller.LoginPenggunaController)
	e.POST("/penggunas", controller.CreatePenggunaController)
	e.GET("/penggunas/:id", controller.GetPenggunaByIDController, middleware.AuthJWT)
	e.DELETE("/penggunas/:id", controller.DeletePenggunaByIDController, middleware.AuthJWT)
	e.PUT("/penggunas/:id", controller.UpdatePenggunaByIDController, middleware.AuthJWT)

}
