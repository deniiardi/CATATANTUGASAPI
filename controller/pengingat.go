package controller

import (
	"CatatanTugas/lib/database"
	"CatatanTugas/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetAllPengingatsController(c echo.Context) error {
	pengingats := database.GetPengingats()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success get all pengingats",
		"data":    pengingats,
	})
}

func GetPengingatByIDController(c echo.Context) error {
	id := c.Param("id")
	pengingat := database.GetPengingatByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success get pengingat by id",
		"data":    pengingat.ID,
	})
}

func DeletePengingatByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeletePengingatByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Pengingat Successfully Deleted",
	})
}

func UpdatePengingatByIDController(c echo.Context) error {
	id := c.Param("id")

	var pengingat model.Pengingat
	if err := c.Bind(&pengingat); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error update pengingat",
			"error":   err.Error(),
		})
	}
	database.UpdatePengingatByID(id, pengingat)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Pengingat Successfully Updated",
	})
}

func CreatePengingatController(c echo.Context) error {
	var newPengingat model.Pengingat
	if err := c.Bind(&newPengingat); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error create pengingat",
			"error":   err.Error(),
		})
	}

	newPengingat = database.CreatePengingat(newPengingat)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Pengingat Successfully Created",
		"data":    newPengingat,
	})
}
