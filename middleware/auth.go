package middleware

import (
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const key = "legal"

// func JWTauthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		authorizationFromHeader := c.Request().Header.Get("authorization")

// 		tokenString := strings.ReplaceAll(authorizationFromHeader, "Bearer ", "")

// 		claims := jwt.MapClaims{}
// 		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
// 			return []byte("legal"), nil
// 		})

// 		if err != nil && !token.Valid {
// 			return c.String(http.StatusForbidden, "Token Salah")
// 		}

// 		c.Set("email", claims["penggunaID"])
// 		return next(c)
// 	}
// }

func CreateToken(penggunaId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["penggunaId"] = penggunaId
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

// func ExtractTokenPenggunaId(e echo.Context) int {
// 	pengguna := e.Get("pengguna").(*jwt.Token)
// 	if pengguna.Valid {
// 		claims := pengguna.Claims.(jwt.MapClaims)
// 		penggunaId := claims["penggunaId"].(int)
// 		return penggunaId
// 	}
// 	return 0
// }

func AuthJWT(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorizationFromHeader := c.Request().Header.Get("authorization")
		if authorizationFromHeader == "" {
			return c.String(http.StatusForbidden, "Forbidden")
		}

		tokenString := strings.ReplaceAll(authorizationFromHeader, "Bearer ", "")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})
		if err != nil && !token.Valid {
			return c.String(http.StatusForbidden, "Wrong Token")
		}

		c.Set("email", claims["penggunaId"])
		return hf(c)
	}
}
