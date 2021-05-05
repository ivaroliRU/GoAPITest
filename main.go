package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	// create echo instance
	e := echo.New()

	// simple health check
	e.GET("/api/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Ok")
	})

	/*e.GET("/index", Controllers.Index)
	e.GET("/show/:id", Controllers.Show)
	e.POST("/store", Controllers.Store)
	e.PUT("/update/:id", Controllers.Update)
	e.DELETE("/delete/:id", Controllers.Delete)*/

	//start server
	e.Logger.Fatal(e.Start(":8000"))
}