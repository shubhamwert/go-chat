package main

import (
	chatApp "chatApp/ChatRoom"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Rooms map[int]*chatApp.Room

func main() {
	Rooms = make(map[int]*chatApp.Room)
	Rooms[0] = &chatApp.Room{RoomId: 0, Messages: []chatApp.Message{}}
	Rooms[0].AddMessage("HERE")
	Rooms[0].AddMessage("MSG")
	Rooms[0].AddMessage("THERE")
	Rooms[0].AddMessage("MSG")

	fmt.Println("Starting")

	router := gin.Default()

	router.GET("/health-check", HealthCheck)
	router.GET("/room/:roomId", getRoom)
	router.GET("/room/:roomId/msg", testMsg)
	router.Run()
}
func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Success")
}
func getRoom(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		fmt.Println("error converting id")
	}
	msg := Rooms[roomId].GetMessages()

	c.IndentedJSON(http.StatusOK, msg)
	// conn = LiveFlow(c.Read)
}
func testMsg(c *gin.Context) {
	// msg := chatApp.Message{Msg: "This is a test message", Id: 0, From: "System"}
	roomId, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		fmt.Println("Error Happened")
		return
	}

	Rooms[roomId].AddConnection(c.Writer, c.Request)
	// Rooms[roomId].BroadCastMessage(msg)
}
