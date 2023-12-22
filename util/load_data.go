package util

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gauravshinde1816/card_tracker/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx context.Context

const DB_NAME = "card_tracker"

func LoadData(client *mongo.Client) {
	// Drop the existing collection
	client.Database(DB_NAME).Collection("card_status").Drop(ctx)

	// Open and read CSV files from the data folder
	files := []string{"card_delivery_data.csv", "delivery_exception_data.csv", "card_pickup_data.csv", "card_returned_data.csv"}

	for _, file := range files {
		filePath := fmt.Sprintf("data/%s", file)
		csvFile, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer csvFile.Close()

		reader := csv.NewReader(csvFile)

		// Read and skip the header
		_, err = reader.Read()
		if err != nil {
			log.Fatal(err)
		}

		// Read and insert data into MongoDB
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			switch file {
			case "card_delivery_data.csv":
				saveCardStatus(record[1], record[2], "DELIVERED", record[3], "", client)
			case "delivery_exception_data.csv":
				saveCardStatus(record[1], record[2], "EXCEPTION", record[3], record[4], client)
			case "card_pickup_data.csv":
				saveCardStatus(record[1], record[2], "PICKED_UP", record[3], "", client)
			case "card_returned_data.csv":
				saveCardStatus(record[1], record[2], "RETURNED", record[3], "", client)
			}
		}
	}
}

func saveCardStatus(cardID, userContact, status, timestamp string, comment string, client *mongo.Client) {
	var parsedTime time.Time
	if len(timestamp) == 25 { // "02-01-2006 3:04 PM" format
		layout := "02-01-2006 3:04 PM"
		parsedTime, _ = time.Parse(layout, timestamp)
	} else if len(timestamp) == 20 {
		parsedTime, _ = time.Parse(time.RFC3339, timestamp)
	} else {
		parsedTime = time.Now()
	}

	// Extract user ID from phone number
	userID := strings.ReplaceAll(userContact, "\"", "")

	// Insert data into MongoDB
	cardStatus := models.CardStatus{
		CardID:    cardID,
		UserID:    userID,
		Status:    status,
		Comment:   comment,
		Timestamp: parsedTime,
	}

	collection := client.Database(DB_NAME).Collection("card_status")
	_, err := collection.InsertOne(ctx, cardStatus)
	if err != nil {
		log.Fatal(err)
	}
}
