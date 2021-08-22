package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// type Book struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Author string  `json:"author"`
// 	Price  float64 `json:"price"`
// }

// var books = []Book{
// 	{ID: "1", Title: "It", Author: "Stephen King", Price: 49.98},
// 	{ID: "2", Title: "Insomnia", Author: "Stephen King", Price: 29.49},
// 	{ID: "3", Title: "The Green Mile", Author: "Stephen King", Price: 78.50},
// 	{ID: "4", Title: "The Masque of the Red Death", Author: "Edgar Allan Poe", Price: 24.99},
// 	{ID: "5", Title: "A Dream Within a Dream", Author: "Edgar Allan Poe", Price: 14.99},
// }

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var MI MongoInstance

func ConnectDB() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")

	MI = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DB")),
	}
}

// func getHome(c *fiber.Ctx) error {
// 	return c.SendString("Hello, Go!")
// }
