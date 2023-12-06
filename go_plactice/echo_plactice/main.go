package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/users/:id", getUser)

	// use prefix
	ad := e.Group("/admin")
	ad.GET("/", hello)
	ad.GET("/users/:id", getUser)

	type User struct {
		Name  string `json:"name" xml:"name" form:"name" query:"name"`
		Email string `json:"email" xml:"email" form:"email" query:"email"`
	}
	
	e.POST("/User", func(c echo.Context) error {
		u := new(User)
		// Bind method binds data from the request body
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	return c.String(http.StatusOK, "team:" + c.Param("id") + ", member:" + c.QueryParam("member"))
}
