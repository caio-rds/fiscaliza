package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type Message struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Anonymous   int       `json:"anonymous"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Street      string    `json:"street"`
	District    string    `json:"district"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Lat         string    `json:"lat"`
	Lon         string    `json:"lon"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

var upg = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsConn *websocket.Conn

func Connections(c *gin.Context) {
	var err error
	wsConn, err = upg.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(wsConn)

	for {
		messageType, message, err := wsConn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		if err = wsConn.WriteMessage(messageType, message); err != nil {
			fmt.Println(err)
			break
		}
	}
}

func SendMessage(message *Message) error {
	if wsConn != nil {
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			return err
		}
		err = wsConn.WriteMessage(websocket.TextMessage, jsonMessage)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("WebSocket connection is not established")
}
