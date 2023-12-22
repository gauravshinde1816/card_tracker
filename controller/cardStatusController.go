package controller

import (
	"context"
	"log"

	"github.com/gauravshinde1816/card_tracker/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardController struct {
	clinet *mongo.Client
}

var ctx context.Context

const DB_NAME = "card_tracker"

func NewCardController(client *mongo.Client) *CardController {
	return &CardController{clinet: client}
}

func (cardController *CardController) GetCardStatus(c *gin.Context) {
	// Retrieve card status based on user's phone number or card ID
	userID := c.Query("user_id")
	cardID := c.Query("card_id")

	var filter bson.M
	if userID != "" {
		filter = bson.M{"user_id": userID}
	} else if cardID != "" {
		filter = bson.M{"card_id": cardID}
	} else {
		c.JSON(400, gin.H{"error": "Please provide user_id or card_id parameter"})
		return
	}

	// Query MongoDB for card status
	collection := cardController.clinet.Database(DB_NAME).Collection("card_status")
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	defer cursor.Close(ctx)

	var cardStatusList []models.CardStatus
	if err = cursor.All(ctx, &cardStatusList); err != nil {
		log.Fatal(err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	// Return card status as JSON
	c.JSON(200, gin.H{"card status": cardStatusList})
}
