package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Item : Model Item
type Item struct {
	Name  string `Json:"name"`
	Price int    `Json:"price"`
}

// GetPrices :  Return Item Prices
func GetPrices(c echo.Context) error {
	Items := make([]Item, 0)
	for i := 1; i < 10; i++ {
		var item Item
		item.Name = "Item"
		item.Price = i * 1000

		Items = append(Items, item)
	}
	return c.JSON(http.StatusOK, Items)
}

//GetPrice :  by Item id
func GetPrice(c echo.Context) error {
	id := c.Param("id")
	var item Item
	item.Name = fmt.Sprintf("item-%s", id)
	return c.JSON(http.StatusOK, item)
}

//GetPriceByQuery : Get by query string
func GetPriceByQuery(c echo.Context) error {
	name := c.QueryParam("name")
	var item Item
	item.Name = name
	return c.JSON(http.StatusOK, item)
}

//Save : Save To Array
func Save(c echo.Context) error {
	name := c.FormValue("name")
	price := c.FormValue("price")
	var item Item
	item.Name = name
	item.Price, _ = strconv.Atoi(price)

	return c.JSON(http.StatusAccepted, item)
}

//SaveJson : Save Json To Array
func SaveJson(c echo.Context) error {

	var item Item
	c.Bind(&item)

	return c.JSON(http.StatusAccepted, item)
}
