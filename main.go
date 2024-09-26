package main

import (
	"gamelib/controllers"
	"gamelib/database"
	"gamelib/repositories"
	"gamelib/routes"
	"gamelib/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(db *mongo.Database) *controllers.GameController {
	gameRepository := repositories.NewGameRepository(db)
	gameService := services.NewGameService(gameRepository)
	gameController := controllers.NewGameController(gameService)

	return gameController
}

func setupRouter(gameController *controllers.GameController) *gin.Engine {
	return routes.SetupRouter(gameController)
}

func main() {
	db := database.ConnectMongoDB()

	gameController := initDependencies(db)

	router := setupRouter(gameController)
	router.Run(":8080")
}
