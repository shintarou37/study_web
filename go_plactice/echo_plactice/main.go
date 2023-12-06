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

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/middleware")
	// Executed for requests in the middleware group
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	// The third argument is executed first
	e.GET("/users", func(c echo.Context) error {
		// Executed after the next method of the track function
		return c.String(http.StatusOK, "/users")
	}, track)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	return c.String(http.StatusOK, "team:"+c.Param("id")+", member:"+c.QueryParam("member"))
}
