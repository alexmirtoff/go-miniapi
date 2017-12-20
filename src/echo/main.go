package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	// nameVar := "user"
	// passVar := "passs"
	//dataVar := "test"

	e := echo.New()

	//e.GET("/", getRoot)
	e.GET("/get", func(c echo.Context, dataVar string) error{
		return dataVar
	})



	e.Logger.Fatal(e.Start(":8080"))
}

func getRoot(c echo.Context, s string) error {
	return c.String(http.StatusOK, "Hello, World!")
}
