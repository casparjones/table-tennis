package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"tt-points/controller"
	"tt-points/game"
)

type Web struct {
	Port string
}

func (w *Web) Start() *gin.Engine {
	appEngine := gin.Default()
	appEngine.LoadHTMLGlob("templates/*.twig")
	appEngine.Static("/assets", "./assets")

	// appEngine.Use(cors.Default())
	routes := controller.Routes.Get()
	for routeConfig, handle := range routes {
		if routeConfig.Method == "GET" {
			appEngine.GET(routeConfig.Route, handle)
		} else if routeConfig.Method == "POST" {
			appEngine.POST(routeConfig.Route, handle)
		} else {
			appEngine.GET(routeConfig.Route, handle)
		}

	}

	return appEngine
}

type Socket struct {
	Port string
	ws   []*websocket.Conn
}

func remove(s []*websocket.Conn, i int) []*websocket.Conn {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (s *Socket) Broadcast(game game.Game) {
	var removeItems []int
	if s.ws != nil {
		for i, ws := range s.ws {
			err := ws.WriteJSON(game)
			if err != nil {
				removeItems = append(removeItems, i)
				log.Println(err.Error())
			}
		}

		for _, i := range removeItems {
			s.ws = remove(s.ws, i)
		}
	}
}

func (s *Socket) Start(r *gin.Engine) {
	upgrade := websocket.Upgrader{}
	r.GET("/ws", func(c *gin.Context) {
		//upgrade get request to websocket protocol
		ws, err := upgrade.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		s.ws = append(s.ws, ws)

		// defer s.ws.Close()
		for {
			//Read Message from client
			mt, message, err := ws.ReadMessage()
			if err != nil {
				fmt.Println(err)
				break
			}
			//If client message is ping will return pong
			if string(message) == "ping" {
				message = []byte("pong")
				//Response message to client
				err = ws.WriteMessage(mt, message)
				if err != nil {
					fmt.Println(err)
					break
				}
			}
		}
	})
}

func NewSocket() Socket {
	port := os.Getenv("SOCKET_PORT")
	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "8080"
	}

	return Socket{port, nil}
}

func NewWeb() Web {
	port := os.Getenv("PORT")
	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "8080"
	}

	return Web{port}
}
