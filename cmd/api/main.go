package main

import (
	"fmt"
	"log"

	"github.com/LaviqueDias/api-hotel-go/cmd/api/dependecies"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/database/postgresql"
	"github.com/LaviqueDias/api-hotel-go/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Olá, mundo")
	db, err := postgresql.Connection()
	if err != nil {
		log.Fatal("Err to connect to bd", err)
	}
	defer postgresql.CloseConnection()

	p := db.Ping()
	fmt.Println(p)

	hotelController := dependecies.InitHotelDependencies(db)
	roomController := dependecies.InitRoomDependencies(db)
	guestController := dependecies.InitGuestDependencies(db)

	router := gin.Default()
	routerApi := router.Group("/api")

	hotelGroup := routerApi.Group("/hotel")
	routes.InitHotelRoutes(hotelGroup, hotelController)

	roomGroup := routerApi.Group("/room")
	routes.InitRoomRoutes(roomGroup, roomController)

	guestGroup := routerApi.Group("/guest")
	routes.InitGuestRoutes(guestGroup, guestController)


	if err := router.Run(":8080"); err != nil {
		log.Fatal()
	}
}