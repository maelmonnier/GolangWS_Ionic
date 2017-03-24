package main

import (
"net/http"

"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"username" xml:"username" form:"username" query:"username"`
	Email string `json:"email" xml:"email" form:"email" query:"name"`
}

// ServerHeader middleware adds a `Server` header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		return next(c)
	}
}

func main() {
	e := echo.New()

	// Server header
	e.Use(ServerHeader)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.OPTIONS("/users", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderAllow, "POST")
		c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://localhost:8000")
		c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Content-Type")
		return c.String(http.StatusOK, "Just post of a new user for the moment")
	})

	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

