package chatApp

type Message struct {
	Msg  string `json:"message"`
	Id   int    `json:"Id"`
	From string `json:"From"`
}
