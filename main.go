package main

import (
	chatApp "chatApp/ChatRoom"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Rooms map[int]*chatApp.Room

// var RoomsList map[*chatApp.User]*chatApp.Room

func main() {
	chatApp.CreateConnection("config.yaml")
	defer chatApp.CloseConnection()
	Rooms = make(map[int]*chatApp.Room)
	Rooms[0] = &chatApp.Room{RoomId: 0, Messages: []chatApp.Message{}, RoomName: "Genysis"}

	Rooms[1] = &chatApp.Room{RoomId: 1, Messages: []chatApp.Message{}, RoomName: "Genysis2"}
	Rooms[2] = &chatApp.Room{RoomId: 2, Messages: []chatApp.Message{}, RoomName: "Genysis3"}
	fmt.Println("Starting")

	router := gin.Default()

	router.GET("/health-check", HealthCheck)
	router.GET("/room/:roomId", GetMessages)
	router.GET("/room/:roomId/connect", Connect)
	router.GET("/room/createNewRoom", CreateRoom)

	router.GET("/", rooms)
	// router.GET("/login", login)
	router.Run()
}
func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Success")
}
func GetMessages(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("roomId"))
	if err != nil {
		fmt.Println("error converting id")
	}
	msg := Rooms[roomId].GetMessages()

	c.IndentedJSON(http.StatusOK, msg)
}
func Connect(c *gin.Context) {
	// msg := chatApp.Message{Msg: "This is a test message", Id: 0, From: "System"}
	roomId, err := strconv.Atoi(c.Param("roomId"))

	if err != nil {
		fmt.Println("Error Happened")
		return
	}

	Rooms[roomId].AddConnection(c.Writer, c.Request)
	// Rooms[roomId].BroadCastMessage(msg)
}
func CreateRoom(c *gin.Context) {
	roomNo := len(Rooms)
	fmt.Println(roomNo)
	fmt.Println(c.Query("roomname"))
	roomName := c.Query("roomname")
	Rooms[roomNo] = &chatApp.Room{RoomId: roomNo, Messages: []chatApp.Message{}, RoomName: roomName}
	c.IndentedJSON(http.StatusOK, "{status:ok,room:created}")
}
func rooms(c *gin.Context) {
	r := make(map[int]string, 0)

	for k, v := range Rooms {
		r[k] = v.RoomName
		fmt.Println(r)

	}
	c.Header("Access-Control-Allow-Origin", "*")

	c.IndentedJSON(http.StatusOK, r)
}

// func login(c *gin.Context) {

// 	var (
// 		// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
// 		key   = []byte(RandStringBytesMaskImprSrc(64))
// 		store = sessions.NewCookieStore(key)
// 	)

// 	session, _ := store.Get(c.Request, "login-connect")

// 	// Authentication goes here
// 	// ...

// 	// Set user as authenticated
// 	session.Values["authenticated"] = true
// 	session.Save(c.Request, c.Writer)
// }

// var src = rand.NewSource(time.Now().UnixNano())

// const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// const (
// 	letterIdxBits = 6                    // 6 bits to represent a letter index
// 	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
// 	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
// )

// func RandStringBytesMaskImprSrc(n int) string {
// 	b := make([]byte, n)
// 	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
// 	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
// 		if remain == 0 {
// 			cache, remain = src.Int63(), letterIdxMax
// 		}
// 		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
// 			b[i] = letterBytes[idx]
// 			i--
// 		}
// 		cache >>= letterIdxBits
// 		remain--
// 	}

// 	return string(b)
// }
