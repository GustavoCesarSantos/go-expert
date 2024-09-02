package main

import (
	"rest-gin-postgresql/controllers"
	"rest-gin-postgresql/db"
	"rest-gin-postgresql/repositories"
	"rest-gin-postgresql/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
    dbConnection, err := db.ConnectDB()
    if err != nil {
        panic(err)
    }
    ProductRepository := repositories.NewProductRepository(dbConnection) 
    ProductUseCases := usecases.NewProductUseCase(ProductRepository)
    ProductController := controllers.NewProductController(ProductUseCases)
    server := gin.Default()
    server.GET("/ping", func(ctx *gin.Context) {
        ctx.JSON(200, gin.H{
            "message": "pong",
        })
    })
    server.GET("/products", ProductController.GetProducts)
    server.POST("/products", ProductController.CreateProduct)
    server.GET("/products/:id", ProductController.GetProduct)
    server.Run(":8000")
}
