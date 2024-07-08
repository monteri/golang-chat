package main

import (
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"sync"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	rooms     = make(map[string][]*websocket.Conn)
	mu        = new(sync.Mutex)
	db        *gorm.DB
	cassandra *gocql.Session
)

func main() {
	db = setupDatabase()
	cassandra = setupCassandra()

	r := setupRouter()
	r.Run(":8080")
}
