package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func handleWebSocket(c *gin.Context) {
	room := c.Param("room")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	mu.Lock()
	rooms[room] = append(rooms[room], conn)
	mu.Unlock()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		broadcastMessage(room, message)
		storeChatMessage(room, message)
	}

	mu.Lock()
	for i, c := range rooms[room] {
		if c == conn {
			rooms[room] = append(rooms[room][:i], rooms[room][i+1:]...)
			break
		}
	}
	mu.Unlock()
}

func broadcastMessage(room string, message []byte) {
	mu.Lock()
	defer mu.Unlock()
	for _, conn := range rooms[room] {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			conn.Close()
		}
	}
}
