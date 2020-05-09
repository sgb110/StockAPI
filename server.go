package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/prices", GetPrices)
	e.GET("/price/:id", GetPrice)
	e.GET("/price", GetPriceByQuery)
	e.POST("/price", Save)
	e.PUT("/price", SaveJson)

	e.Logger.Fatal(e.Start(":8000"))
}
