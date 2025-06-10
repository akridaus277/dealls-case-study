package main

import (
	"audit-service/config"
	"audit-service/database"
	"audit-service/kafka"
	"audit-service/repositories"
	"audit-service/routes"
	service "audit-service/services"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "audit-service/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Authentication Service API
//	@version		1.0
//	@description	API for user authentication and management
//	@host			localhost:8080
//	@BasePath		/api

//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

// @schemes	http
func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found. Using system environment variables.")
	}

	// Init config, db, redis, dll
	config.LoadEnv()
	database.Init()
	db := config.ConnectDB()
	config.ConnectRedis()

	// Setup repository & service
	repo := repositories.NewAuditRepository(db)
	svc := service.NewAuditService(repo)

	// Start Kafka consumer di goroutine
	ctx := context.Background()
	go kafka.StartConsumer(svc, ctx)

	// Setup routes
	r := gin.Default()
	routes.RegisterRoutes(r)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run server
	r.Run(":" + config.Env("PORT"))
}
