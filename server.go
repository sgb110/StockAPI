package main

import (
	"stockApi/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/prices", controllers.GetPrices)
	e.GET("/price/:id", controllers.GetPrice)
	e.GET("/price", controllers.GetPriceByQuery)
	e.POST("/price", controllers.Save)
	e.PUT("/price", controllers.SaveJson)

	e.Logger.Fatal(e.Start(":8000"))
}
