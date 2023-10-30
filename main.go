package main

import (
	"log"
	"snjat/common"
	"snjat/repository"
	"snjat/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	_, err := common.LoadConfig()
	if err != nil {
		log.Println("error loading config", err)
		return
	}

	repository.SeedDatabase()

	router := gin.Default()

	routes.UserRoutes(router)

	router.Run()

}
