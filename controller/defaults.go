package controller

import (
	"tt-points/game"
)

type Broadcaster interface {
	Broadcast(game game.Game)
}

var (
	Ok     = map[string]string{"status": "ok"}
	Game   = game.NewGame()
	Routes = NewRoutes()
	bc     Broadcaster
)

func SetBroadcaster(newBc Broadcaster) {
	bc = newBc
}

func SendToBroadcast(obj *game.Game) {
	if bc != nil {
		bc.Broadcast(obj.GetGame())
	}
}
