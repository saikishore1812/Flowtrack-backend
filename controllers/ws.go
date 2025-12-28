package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan interface{})

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// -------- HANDLE CONNECTIONS ----------
func HandleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}

// -------- BROADCAST MESSAGES ----------
func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			client.WriteJSON(msg)
		}
	}
}
 