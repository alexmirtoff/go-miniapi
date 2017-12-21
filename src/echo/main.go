package main

import (
	//"net/http"
	"github.com/labstack/echo"
)


type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Bar() {	
	//bar := "brr"
}

func main() {
	e := echo.New()

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return h(cc)
		}
	})

	e.GET("/", func(c echo.Context) error {
		cc := c.(*CustomContext)
		
		return cc.String(200, "l")
	})


	e.Logger.Fatal(e.Start(":8080"))
}

// func getRoot(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }
