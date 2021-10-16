package controller

import (
	"CatatanTugas/lib/database"
	"CatatanTugas/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllLabelCatatansController(c echo.Context) error {
	labelcatatans := database.GetLabelCatatans()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success get all labelcatatans",
		"data":    labelcatatans,
	})
}

func GetLabelCatatanByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	labelcatatans, err := database.GetLabelCatatansByID(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get by labelcatatan id",
		"data":    labelcatatans,
	})
}

func DeleteLabelCatatanByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeleteLabelCatatanByID(id)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "LabelCatatan Successfully Deleted",
	})
}

func UpdateLabelCatatanByIDController(c echo.Context) error {
	id := c.Param("id")

	var labelcatatan model.LabelCatatan
	if err := c.Bind(&labelcatatan); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error update labelcatatan",
			"error":   err.Error(),
		})
	}
	database.UpdateLabelCatatanByID(id, labelcatatan)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "LabelCatatan Successfully Updated",
	})
}

func CreateLabelCatatanController(c echo.Context) error {
	var newLabelCatatan model.LabelCatatan
	if err := c.Bind(&newLabelCatatan); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "error create labelcatatan",
			"error":   err.Error(),
		})
	}

	newLabelCatatan = database.CreateLabelCatatan(newLabelCatatan)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "LabelCatatan Successfully Created",
		"data":    newLabelCatatan,
	})
}
