package main

import (
	"github.com/gauravshinde1816/card_tracker/controller"
	"github.com/gauravshinde1816/card_tracker/db"
	"github.com/gauravshinde1816/card_tracker/util"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to MongoDB
	client := db.InitDB()

	// Initialize Gin router
	router := gin.Default()

	// Load data into MongoDB
	util.LoadData(client)

	//Get new card status controller
	cardStatusController := controller.NewCardController(client)

	// Define API endpoint
	router.GET("/get_card_status", cardStatusController.GetCardStatus)

	// Run the server
	router.Run(":8080")
}
