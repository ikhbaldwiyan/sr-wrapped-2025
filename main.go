package main

import (
	"log"

	"github.com/ikhbaldwiyan/sr-wrapped-2025/config"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/handler"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/repository"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/router"
	"github.com/ikhbaldwiyan/sr-wrapped-2025/service"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.ConnectMongo()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	mostWatchIdnRepo := repository.NewWatchIDNRepository()
	mostWatchIdnService := service.NewWatchIDNService(mostWatchIdnRepo)
	userHandler := handler.NewUserHandler(userService)
	mostWatchIdnHandler := handler.NewWatchIDNHandler(mostWatchIdnService, userService)

	r := router.SetupRouter(userHandler, mostWatchIdnHandler)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
