package main

import (
	"go-restapi/apps/product"
	"go-restapi/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database.LoadConfig()
	database.LoadDB()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")
	product.RegisterRoute(api, "/product")

	database.DB.AutoMigrate(&product.Product{})

	err := r.Run(":" + database.ENV.PORT)
	if err != nil {
		log.Fatal("? Could not load apps to port"+database.ENV.PORT, err)
	}
}
