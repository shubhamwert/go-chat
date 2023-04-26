package chatApp

import (
	dbhelper "chatApp/dbHelper"
	"fmt"
	"log"
)

var dbConn *dbhelper.Connection

func CreateConnection(path string) {
	dbConn = new(dbhelper.Connection)
	dbConn.LoadConfig(path)
	dbConn.Connect()

}

func InsertMessage(msgs []Message, RoomId int, tableName string) {
	query := " (Msg,From_Msg,RoomId) VALUES"

	for _, msg := range msgs {
		query = fmt.Sprintf(" %s ( '%s', '%s',%d);", query, msg.Msg, msg.From, RoomId)

	}
	result, err := dbConn.InsertQuery(tableName, query, "")

	if err != nil {
		log.Default().Println(err)
	}
	defer result.Close()

}

func GetMessage(RoomId int, tableName string) ([]Message, error) {
	rows, err := dbConn.GetCondition(tableName, fmt.Sprintf("roomId=%d", RoomId), "")
	if err != nil {
		log.Default().Println(err)
		return nil, err

	}
	defer rows.Close()
	msgs := []Message{}
	for rows.Next() {
		msg := Message{}
		var roomId int
		rows.Scan(&msg.Id, &msg.From, &msg.Msg, roomId)
		fmt.Println(msg)
		msgs = append(msgs, msg)
	}

	return msgs, nil

}
func CloseConnection() {
	dbConn.CloseConnection()
}
