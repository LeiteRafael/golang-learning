package services

import (
	"gamelib/models"
	"gamelib/repositories"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameService struct {
	Repository *repositories.GameRepository
}

func NewGameService(repo *repositories.GameRepository) *GameService {
	return &GameService{
		Repository: repo,
	}
}

func (service *GameService) CreateGame(game *models.Game) (*models.Game, error) {
	game.ID = primitive.NewObjectID()
	_, err := service.Repository.Create(game)
	if err != nil {
		return nil, err
	}
	return game, nil
}

func (service *GameService) GetGameByID(id string) (*models.Game, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return service.Repository.FindByID(objID)
}

func (service *GameService) GetAllGames() ([]models.Game, error) {
	return service.Repository.FindAll()
}
