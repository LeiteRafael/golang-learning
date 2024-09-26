package repositories

import (
	"context"
	"gamelib/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameRepository struct {
	Collection *mongo.Collection
}

func NewGameRepository(db *mongo.Database) *GameRepository {
	return &GameRepository{
		Collection: db.Collection("games"),
	}
}

func (repo *GameRepository) Create(game *models.Game) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return repo.Collection.InsertOne(ctx, game)
}

func (repo *GameRepository) FindByID(id primitive.ObjectID) (*models.Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var game models.Game
	err := repo.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (repo *GameRepository) FindAll() ([]models.Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var games []models.Game
	cursor, err := repo.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &games); err != nil {
		log.Fatal(err)
	}
	return games, nil
}
