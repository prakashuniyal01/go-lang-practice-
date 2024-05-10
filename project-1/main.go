package main

import (
	"log"
	"trademarkia/common"
	"trademarkia/repository"
	"trademarkia/routes"

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
