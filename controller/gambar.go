package controller

import (
	"CatatanTugas/lib/database"
	"CatatanTugas/model"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllGambarsController(c echo.Context) error {
	gambars := database.GetGambars()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success get all gambars",
		"data":    gambars,
	})
}

func GetGambarByIDController(c echo.Context) error {
	id := c.Param("id")
	gambar := database.GetGambarByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success get gambar By ID",
		"data":    gambar.ID,
	})
}

func DeleteGambarByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteGambarByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Gambar Successfully Deleted",
	})
}

func UpdateGambarByIDController(c echo.Context) error {
	id := c.Param("id")

	var gambar model.Gambar
	if err := c.Bind(&gambar); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error update gambar",
			"error":   err.Error(),
		})
	}
	database.UpdateGambarByID(id, gambar)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Gambar Successfully Updated",
	})
}

func CreateGambarController(c echo.Context) error {
	var newGambar model.Gambar
	file, err := c.FormFile("gambar")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	//destination
	// dst, err := os.Create(file.Filename)
	// if err != nil {
	// 	return err
	// }
	// defer dst.Close()

	// this is path which  we want to store the file
	f, err := os.OpenFile("gambars/"+file.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer f.Close()

	if _, err = io.Copy(f, src); err != nil {
		return err
	}

	newGambar.Gambar_name = file.Filename

	catatanid, err := strconv.Atoi(c.FormValue("catatanid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}
	newGambar.CatatanID = uint(catatanid)

	newGambar = database.CreateGambar(newGambar)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Gambar Successfully Uploaded",
		"data":    newGambar,
	})
}
