package server

import "github.com/gin-gonic/gin"

// Подготовить HTTP запросы
func prepare() {
	fillProducts()

	router.GET("/", getIndex)
	router.GET("/api/products", getProducts)
}

func getIndex(ctx *gin.Context) {
	ctx.HTML(200, "index", gin.H{})
}

func getProducts(ctx *gin.Context) {
	ctx.JSON(200, products)
}
