package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitRoutes(){
	router := gin.Default()
	// api := router.Group("/api")


	// hotelGroup := api.Group("/hotel")
	// routes.InitPublicUserRoutes(hotelGroup, hotelController)


	if err := router.Run(":8080"); err != nil {
		log.Fatal()
	}
}
