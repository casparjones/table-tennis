package controller

import "github.com/gin-gonic/gin"

type PointsController struct{}

func (pc PointsController) IncPlayer1(c *gin.Context) {
	Game.IncPoints(1)
	SendToBroadcast(Game)
	c.JSON(200, Ok)
}

func (pc PointsController) IncPlayer2(c *gin.Context) {
	Game.IncPoints(2)
	SendToBroadcast(Game)
	c.JSON(200, Ok)
}

func (pc PointsController) DecPlayer1(c *gin.Context) {
	Game.DecPoints(1)
	SendToBroadcast(Game)
	c.JSON(200, Ok)
}

func (pc PointsController) DecPlayer2(c *gin.Context) {
	Game.DecPoints(2)
	SendToBroadcast(Game)
	c.JSON(200, Ok)
}

func (pc PointsController) Reset(c *gin.Context) {
	Game.Reset()
	SendToBroadcast(Game)
	c.JSON(200, Ok)
}

var Points = PointsController{}
