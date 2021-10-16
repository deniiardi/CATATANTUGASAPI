package controller

import (
	"CatatanTugas/lib/database"
	"CatatanTugas/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllCatatansController(c echo.Context) error {
	catatans := database.GetCatatans()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success get all catatans",
		"data":    catatans,
	})
}

func GetCatatanByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	catatans, err := database.GetCatatansByID(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get catatan by ID",
		"data":    catatans,
	})
}

func DeleteCatatanByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteCatatanByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Catatan Successfully Deleted",
	})
}

func UpdateCatatanByIDController(c echo.Context) error {
	id := c.Param("id")

	var catatan model.Catatan
	if err := c.Bind(&catatan); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error update catatan",
			"error":   err.Error(),
		})
	}
	database.UpdateCatatanByID(id, catatan)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Catatan Successfully Updated",
	})
}

func CreateCatatanController(c echo.Context) error {
	var newCatatan model.Catatan
	if err := c.Bind(&newCatatan); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error create catatan",
			"error":   err.Error(),
		})
	}

	newCatatan = database.CreateCatatan(newCatatan)
	return c.JSON(http.StatusOK, echo.Map{
		"message":   "Catatan Successfully Created",
		"catatanID": newCatatan.ID,
	})
}
