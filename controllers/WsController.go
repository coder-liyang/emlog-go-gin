package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		HandshakeTimeout: 0,
		ReadBufferSize:   0,
		WriteBufferSize:  0,
		WriteBufferPool:  nil,
		Subprotocols:     nil,
		Error:            nil,
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: false,
	}
)

func WsHandler(c *gin.Context) {
	var (
		err error
		conn *websocket.Conn
		data []byte
	)
	if conn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err !=nil {
		panic(err)
	}

	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		data = append([]byte("From service: "), data...)
		//conn.WriteJSON([]string{"a","b","c"});
		if err = conn.WriteMessage(websocket.TextMessage, data); err !=nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
	panic(err)

}
