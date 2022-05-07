package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nao11aihara/echo-api-sample/app/models"
)

type Product struct {
	Id	string	`json:"id"`
	Title	string	`json:"title"`
	Price	int	`json:"price"`
}

func getProducts(c echo.Context) error {
	name := c.QueryParam("name")

	fmt.Println("name:", name)
	fmt.Println("商品一覧取得処理")

	product1 := Product{
		Id: "1",
		Title: "商品1",
		Price: 1000,
	}
	product2 := Product{
		Id: "2",
		Title: "商品2",
		Price: 1000,
	}

	products := []Product{product1, product2}

	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	id := c.Param("id")

	fmt.Println("id:", id)
	fmt.Println("商品取得処理")

	product := Product{
		Id: "1",
		Title: "商品1",
		Price: 1000,
	}

	return c.JSON(http.StatusOK, product)
}

func postProduct(c echo.Context) error {
  product := Product{}

  err := c.Bind(&product);
	if err != nil {
    return err
  }

  models.Db.Create(&product)

  return c.JSON(http.StatusCreated, product)
}

func putProduct(c echo.Context) error {
	id := c.Param("id")

	fmt.Println("id:", id)
	fmt.Println("商品更新処理")

	product := Product{
		Id: "1",
		Title: "商品1",
		Price: 1000,
	}

	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	id := c.Param("id")

	fmt.Println("id:", id)
	fmt.Println("商品削除処理")
	
	return c.String(http.StatusNoContent, "")
}
