// @title D4 Microservice API
// @version 1.0
// @description API for user management, authentication, calculator services, AI calls, and async CSV ingestion.
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"os"

	docs "github.com/dimapog/jwt-microservice/docs"
	"github.com/dimapog/jwt-microservice/internal/ai"
	"github.com/dimapog/jwt-microservice/internal/auth"
	"github.com/dimapog/jwt-microservice/internal/calculator"
	"github.com/dimapog/jwt-microservice/internal/csv"
	"github.com/dimapog/jwt-microservice/internal/user"
	"github.com/dimapog/jwt-microservice/utils"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	utils.LoadEnvVariables()
	utils.ConnectToDB()
	utils.SyncDB()
	if err := user.Migrate(); err != nil {
		panic(err)
	}
	if err := csv.Migrate(); err != nil {
		panic(err)
	}
}

func main() {
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = os.Getenv("HOST") + ":" + os.Getenv("PORT")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
