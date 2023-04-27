package chatApp

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	Messages []Message `json:"messages"`
	RoomId   int       `json:"Id"`
	conn     []*webSocketConn
	RoomName string `json:"RoomName"`
}
type webSocketConn struct {
	mu   sync.RWMutex
	conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (R *Room) GetMessages() []Message {
	return R.Messages

}
func (R *Room) AddMessage(msg string, usr string, saveMsg bool) []Message {

	msgF := Message{Msg: msg, Id: len(R.Messages), From: usr}
	fmt.Println("Inserting")

	if saveMsg {
		R.AddMessageDB([]Message{msgF})
	}
	fmt.Println("Insert complete")
	go R.ReceiveMessage([]Message{msgF})

	return R.Messages

}

func (R *Room) AddConnection(w http.ResponseWriter, r *http.Request) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	var u User
	err2 := ws.ReadJSON(&u)

	R.conn = append(R.conn, &webSocketConn{conn: ws})
	if err2 != nil {
		fmt.Println("ERROR   ", err2)
	}
	fmt.Println(u)
	if err != nil {
		fmt.Println("Error upgrading connection")
		fmt.Println(err)

		return err
	}
	R.LoadMessages(ws, u.Username)

	R.AddMessage(fmt.Sprintf("New user connected %s", u.Username), u.Username, false)

	go R.ReadMessage(ws, u)
	return nil
}

func (R *Room) ReceiveMessage(msg []Message) {
	closedConnection := []int{}

	for i, conn := range R.conn {
		go func(conn *webSocketConn) {
			conn.mu.Lock()
			defer conn.mu.Unlock()
			err := conn.conn.WriteJSON(msg)

			if err != nil {
				closedConnection = append(closedConnection, i)
				fmt.Println("error messaging", err)
				conn.conn.Close()

			}
		}(conn)
	}
	for _, v := range closedConnection {
		R.conn = append(R.conn[:v], R.conn[v+1:]...)
	}

}

func (R *Room) ReadMessage(ws *websocket.Conn, u User) {
	for {
		var message Message
		err := ws.ReadJSON(&message)
		if err != nil {
			log.Default().Println("Connection REad Error", err)
			ws.Close()
			return
		}
		if len(message.Msg) <= 1 {
			continue
		}
		message.From = u.Username

		if err != nil {
			fmt.Println(err)

			return
		}
		go R.AddMessage(message.Msg, u.Username, true)
	}
}
func (R *Room) LoadMessages(ws *websocket.Conn, u string) error {
	var err error
	fmt.Println("ROom id is ", R.RoomId)
	R.Messages, err = GetMessage(R.RoomId, "message")
	fmt.Println("Loaded message")
	if err != nil {
		return err
	}
	err = ws.WriteJSON(R.Messages)

	return err
}
func (R *Room) AddMessageDB(msg []Message) {
	log.Default().Println("Adding message to db")
	InsertMessage(msg, R.RoomId, "message")
}
