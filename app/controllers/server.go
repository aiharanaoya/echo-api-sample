package controllers

import "github.com/labstack/echo/v4"

// ルーティング設定・起動
func ConfigServer() {
	e := echo.New()

	api := e.Group("/api")

	api.GET("/products", getProducts)
	api.GET("/products/:id", getProduct)
	api.POST("/products", postProduct)
	api.PUT("/products/:id", putProduct)
	api.DELETE("/products/:id", deleteProduct)

	e.Start(":8080")
}
