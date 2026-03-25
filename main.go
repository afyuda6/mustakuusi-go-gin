package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mustakuusi-go-gin/controllers"
	"net/http"
	"os"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// @title Mustakuusi API
// @version 1.0
// @host localhost:3004
// @BasePath /
func main() {
	dsn := "host=localhost user=postgres password=mustakuusi dbname=mustakuusi port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	}

	gamesController := controllers.GamesController{DB: db}
	charactersController := controllers.CharactersController{DB: db}

	r := gin.Default()

	r.GET("/games", gamesController.Index)
	r.GET("/characters", charactersController.Index)

	r.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/robots.txt" {
			c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(`User-agent: *
Disallow:`))
			c.Abort()
			return
		}
		c.Next()
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, ErrorResponse{
			Status:  404,
			Message: "Not Found",
		})
	})

	port := "3004"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	r.Run(":" + port)
}
