package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
	"golang-chat/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// InitializeDatabase initializes the GORM database connection and runs migrations
func setupDatabase() *gorm.DB {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = db.AutoMigrate(&models.UserRoom{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	return db
}

func setupCassandra() *gocql.Session {
	fmt.Println("CASSANDRA_HOST: %s", os.Getenv("CASSANDRA_HOST"))
	cluster := gocql.NewCluster()
	cluster.Hosts = []string{os.Getenv("CASSANDRA_HOST")}

	// Set keyspace and consistency level
	cluster.Keyspace = "chatapp"
	cluster.Consistency = gocql.Quorum

	// Create a session to interact with Cassandra
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	// Return the session
	return session
}

// addUserToRoom adds a user to a room in the database
func addUserToRoom(username, room string) {
	userRoom := models.UserRoom{Username: username, Room: room}
	result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&userRoom)
	if result.Error != nil {
		panic(result.Error)
	}
}

// storeChatMessage stores a chat message in Cassandra
func storeChatMessage(room string, message []byte) {
	err := cassandra.Query(`INSERT INTO chat_messages (room, message) VALUES (?, ?)`,
		room, string(message)).Exec()
	if err != nil {
		fmt.Println("Failed to store chat message:", err)
	}
}
