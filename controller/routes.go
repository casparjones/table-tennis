package controller

import (
	"github.com/gin-gonic/gin"
)

type RouteType struct {
	Route  string
	Method string
}

type RouterManager struct {
	routes map[RouteType]gin.HandlerFunc
}

func (r RouterManager) Get() map[RouteType]gin.HandlerFunc {
	return r.routes
}

func (r RouterManager) AddRoute(rType RouteType, controller gin.HandlerFunc) {
	r.routes[rType] = controller
}

func NewRoutes() RouterManager {
	return RouterManager{
		map[RouteType]gin.HandlerFunc{
			{"/", "GET"}:                  Spiel.Index,
			{"/settings", "GET"}:          Spiel.Settings,
			{"/game/info", "GET"}:         Spiel.Info,
			{"/game/switch-mode", "GET"}:  Spiel.SwitchMode,
			{"/game/ball/1", "GET"}:       Spiel.Ball2Player1,
			{"/game/ball/2", "GET"}:       Spiel.Ball2Player2,
			{"/point/inc/player1", "GET"}: Points.IncPlayer1,
			{"/point/inc/player2", "GET"}: Points.IncPlayer2,
			{"/point/dec/player1", "GET"}: Points.DecPlayer1,
			{"/point/dec/player2", "GET"}: Points.DecPlayer2,
			{"/point/reset", "GET"}:       Points.Reset,
			{"/player1", "POST"}:          Player.Player1,
			{"/player2", "POST"}:          Player.Player2,
		},
	}
}
