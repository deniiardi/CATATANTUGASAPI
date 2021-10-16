package controller

import (
	"CatatanTugas/lib/database"
	"CatatanTugas/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllPenggunasController(c echo.Context) error {
	penggunas := database.GetPenggunas()
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Success get all penggunas",
		"data":    penggunas,
	})
}

func DeletePenggunaByIDController(c echo.Context) error {
	id := c.Param("id")
	database.DeletePenggunaByID(id)

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Pengguna Successfully Deleted",
	})
}

func UpdatePenggunaByIDController(c echo.Context) error {
	id := c.Param("id")

	var pengguna model.Pengguna
	if err := c.Bind(&pengguna); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Fail to update pengguna",
			"error":   err.Error(),
		})
	}
	database.UpdatePenggunaByID(id, pengguna)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Pengguna Successfully Updated",
	})
}

func CreatePenggunaController(c echo.Context) error {
	var newPengguna model.Pengguna
	if err := c.Bind(&newPengguna); err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"message": "Fail to create pengguna",
			"error":   err.Error(),
		})
	}

	newPengguna = database.CreatePengguna(newPengguna)
	newPengguna.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"message": "Pengguna successfully created",
	})

}

// func LoginPenggunaController(c echo.Context) error {
// 	var pengguna model.Pengguna

// 	isValid := database.IsValid(pengguna.Email, pengguna.Name)
// 	if !isValid {
// 		return c.String(http.StatusBadRequest, "Email atau Password Salah")
// 	}

// 	claims := jwt.MapClaims{}

// 	claims["penggunaID"] = pengguna.Email
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte("legal"))

// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.String(http.StatusOK, tokenString)
// }

func LoginPenggunaController(c echo.Context) error {
	pengguna := model.Pengguna{}
	c.Bind(&pengguna)

	token, err := database.LoginPengguna(&pengguna)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Login Success",
		"token":  token,
	})

}

func GetPenggunaByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	penggunas, err := database.GetDetailPengguna(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":   "Login Success",
		"message":  "Success get pengguna by id",
		"pengguna": penggunas,
	})

}
