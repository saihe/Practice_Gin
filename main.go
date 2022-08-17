package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		log.Println(c.Request)
		c.JSON(200, gin.H{"message": "Hello gin !"})
	})

	r.Run(":8080")
}
