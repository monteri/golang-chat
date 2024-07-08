package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/join", func(c *gin.Context) {
		room := c.PostForm("room")
		username := c.PostForm("username")
		addUserToRoom(username, room)
		c.HTML(http.StatusOK, "chat.html", gin.H{"room": room, "username": username})
	})

	r.GET("/ws/:room", func(c *gin.Context) {
		handleWebSocket(c)
	})

	return r
}
