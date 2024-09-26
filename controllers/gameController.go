package controllers

import (
	"gamelib/models"
	"gamelib/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	Service *services.GameService
}

func NewGameController(service *services.GameService) *GameController {
	return &GameController{
		Service: service,
	}
}

func (ctrl *GameController) CreateGame(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGame, err := ctrl.Service.CreateGame(&game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newGame)
}

func (ctrl *GameController) GetGameByID(c *gin.Context) {
	id := c.Param("id")
	game, err := ctrl.Service.GetGameByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.JSON(http.StatusOK, game)
}

func (ctrl *GameController) GetAllGames(c *gin.Context) {
	games, err := ctrl.Service.GetAllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, games)
}
