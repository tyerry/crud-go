package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tyerry/crud-go/src/configuration/logger"
	"github.com/tyerry/crud-go/src/controller"
	"github.com/tyerry/crud-go/src/controller/routes"
	"github.com/tyerry/crud-go/src/model/service"
)

func main() {
	logger.Info("About to start user application")
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}