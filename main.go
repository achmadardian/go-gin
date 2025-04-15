package main

import (
	"go-gin/config"
	"go-gin/handlers"
	"go-gin/repositories"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	healthcheckHandler := handlers.NewHealthcheck()
	userRepo := repositories.NewUserRepository(config.DB)
	userHandler := handlers.NewUserHandler(*userRepo)

	api := router.Group("api")
	{
		api.GET("/", healthcheckHandler.GetHealth)

		users := api.Group("users")
		{
			users.GET("/", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetById)
			users.POST("/", userHandler.Save)
		}
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load env")
	}
	log.Print("success load env")

	dsn := os.Getenv("DB_DSN")
	config.InitDB(dsn)

	log.Print("server running at http://localhost:8080")
	router.Run()
}
