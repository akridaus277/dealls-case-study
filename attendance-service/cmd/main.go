package main

import (
	"attendance-service/config"
	"attendance-service/database"
	"attendance-service/routes"
	"attendance-service/seed"

	"github.com/gin-gonic/gin"

	_ "attendance-service/cmd/docs"

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

	config.ConnectRedis()

	routes.RegisterRoutes(r)

	seed.SeedAttendancePeriod()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + config.Env("PORT"))
}
