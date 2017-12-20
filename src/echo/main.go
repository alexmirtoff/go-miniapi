package main

import (
	"net/http"
	
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", getRoot)
	e.GET("/user/:id", getUser)

	e.Logger.Fatal(e.Start(":8080"))
}

func getRoot(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
