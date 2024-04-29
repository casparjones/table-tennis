package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type PlayerController struct{}

func (pc PlayerController) Player1(c *gin.Context) {
	Game.Player1.SetName(c.Request.PostFormValue("name"))
	point, err := strconv.Atoi(c.Request.PostFormValue("points"))
	if err == nil {
		Game.Player1.SetPoints(point)
	}

	SendToBroadcast(Game)
	c.JSON(200, Game.Player1)
}

func (pc PlayerController) Player2(c *gin.Context) {
	Game.Player2.SetName(c.Request.PostFormValue("name"))
	point, err := strconv.Atoi(c.Request.PostFormValue("points"))
	if err == nil {
		Game.Player2.SetPoints(point)
	}

	SendToBroadcast(Game)
	c.JSON(200, Game.Player2)
}

var Player = PlayerController{}
