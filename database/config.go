package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoURI() string {
	err := godotenv.Load()
	mongoURI := os.Getenv("MONGO_URI")

	if mongoURI == "" || err != nil {
		log.Fatal("Erro ao pegar MONGO_URI")
	}
	return mongoURI
}

func ConnectMongoDB() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := getMongoURI()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Erro ao conectar ao MongoDB:", err)
	}

	return client.Database("gamedb")
}
