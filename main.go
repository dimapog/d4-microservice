package main

import (
	"github.com/dimapog/jwt-microservice/internal/ai"
	"github.com/dimapog/jwt-microservice/internal/auth"
	"github.com/dimapog/jwt-microservice/internal/calculator"
	"github.com/dimapog/jwt-microservice/internal/csv"
	"github.com/dimapog/jwt-microservice/internal/user"
	"github.com/dimapog/jwt-microservice/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnvVariables()
	utils.ConnectToDB()
	utils.SyncDB()
}

func main() {
	router := gin.Default()

	// Initialize user module
	userRepo := user.NewRepository(utils.DB)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	userHandler.RegisterRoutes(router)

	// Register auth module
	authService := auth.NewService(userService)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(router)

	// Register AI module
	aiService := ai.NewService()
	aiHandler := ai.NewHandler(aiService)
	aiHandler.RegisterRoutes(router)

	// Register Calculator module
	calcService := calculator.NewService(userService)
	calcHandler := calculator.NewHandler(calcService)
	calcHandler.RegisterRoutes(router)

	// Register CSV module
	csvRepo := csv.NewRepository(utils.DB)
	csvService := csv.NewService(csvRepo)
	csvHandler := csv.NewHandler(csvService)
	csvHandler.RegisterRoutes(router)

	router.Run()
}
