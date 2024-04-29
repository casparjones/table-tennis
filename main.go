package main

import (
	"github.com/joho/godotenv"
	"tt-points/app"
	"tt-points/controller"
)

func main() {
	godotenv.Load()

	web := app.NewWeb()
	engine := web.Start()

	socket := app.NewSocket()
	socket.Start(engine)
	controller.SetBroadcaster(&socket)

	engine.Run()
}
