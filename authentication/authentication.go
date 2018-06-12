package authentication

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type jwtCustomClaims struct {
	jwt.StandardClaims
	Name string `json:"name"`
}

func JWTConfiguration() middleware.JWTConfig {
	jwtconfig := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("strongSecretKey"),
	}
	return jwtconfig
}

func JWTParser(c echo.Context) *jwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	return claims
}

func CreateToken() string {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		"admin",
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("strongSecretKey"))
	if err != nil {
		fmt.Println(err)
	}

	return t
}

func Authentication(c echo.Context) error {
	o := new(LoginUser)
	if err := c.Bind(o); err != nil {
		return err
	}

	if o.Email == "admin@test.com" && o.Password == "password" {
		token := CreateToken()

		return c.JSON(http.StatusOK, echo.Map{
			"X-Token": token,
		})
	}

	return echo.ErrUnauthorized

}
