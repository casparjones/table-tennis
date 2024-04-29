package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type SpielController struct{}

func (gc SpielController) Info(c *gin.Context) {
	c.JSON(200, Game)
}

func (gc SpielController) SwitchMode(c *gin.Context) {
	Game.SwitchMode()
	SendToBroadcast(Game)
	c.JSON(200, Ok)
}

func (gc SpielController) Ball2Player1(c *gin.Context) {
	Game.Ball = 1
	SendToBroadcast(Game)
	c.JSON(200, Ok)
}

func (gc SpielController) Ball2Player2(c *gin.Context) {
	Game.Ball = 2
	SendToBroadcast(Game)
	c.JSON(200, Ok)
}

func (gc SpielController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.twig", gin.H{
		"title":       "TSC Heimersdorf",
		"socket_host": os.Getenv("SOCKET_HOST"),
	})
}

func (gc SpielController) Settings(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.twig", gin.H{
		"title":       "Settings Punkte",
		"socket_host": os.Getenv("SOCKET_HOST"),
	})
}

var Spiel = SpielController{}
