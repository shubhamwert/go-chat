package chatApp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Room struct {
	Messages []Message `json:"messages"`
	RoomId   int       `json:"Id"`
	conn     []*websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (R *Room) GetMessages() []Message {
	return R.Messages

}
func (R *Room) AddMessage(msg string) []Message {
	R.Messages = append(R.Messages, Message{Msg: msg, Id: len(R.Messages), From: "system"})

	return R.Messages

}

func (R *Room) AddConnection(w http.ResponseWriter, r *http.Request) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection")
		fmt.Println(err)

		return err
	}
	R.conn = append(R.conn, ws)
	go R.ReadMessage(ws)
	return nil
}

func (R *Room) ReceiveMessage(msg Message) {
	closedConnection := []int{}
	for i, conn := range R.conn {
		err := conn.WriteJSON(msg)
		if err != nil {
			closedConnection = append(closedConnection, i)
			fmt.Println("error messaging", err)

		}
	}
	for _, v := range closedConnection {
		R.conn = append(R.conn[:v], R.conn[v+1:]...)
	}

}
func (R *Room) BroadCastMessage(msg Message) {
	fmt.Println("SENDING MSG")
	R.ReceiveMessage(msg)

}
func (R *Room) ReadMessage(ws *websocket.Conn) {

	for {
		var message Message
		err := ws.ReadJSON(&message)
		if err != nil {
			fmt.Println(err)

			return
		}
		R.ReceiveMessage(message)
	}
}
