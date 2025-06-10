package main

import (
	"employee-service/config"
	"employee-service/database"
	"employee-service/routes"
	"employee-service/seed"

	"github.com/gin-gonic/gin"

	_ "employee-service/cmd/docs"

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
	database.Init()

	r := gin.Default()
	config.LoadEnv()
	config.ConnectDB()
	config.ConnectRedis()

	seed.SeedEmployees()

	routes.RegisterRoutes(r)

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + config.Env("PORT"))
}
