package main

import (
	"CatatanTugas/config"
	"CatatanTugas/middleware"
	"CatatanTugas/route"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	config.InitLog()
	config.InitMigration()

	e := echo.New()
	e.Use(middleware.Log)
	route.NewPengguna(e)
	route.NewCatatans(e)
	route.NewLabelCatatan(e)
	route.NewPengingats(e)
	route.NewGambars(e)
	panic(e.Start(":8080"))
}
