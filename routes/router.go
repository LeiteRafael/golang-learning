package routes

import (
	"gamelib/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(gameController *controllers.GameController) *gin.Engine {
	router := gin.Default()

	router.POST("/games", gameController.CreateGame)

	router.GET("/games/:id", gameController.GetGameByID)

	router.GET("/games", gameController.GetAllGames)

	return router
}
