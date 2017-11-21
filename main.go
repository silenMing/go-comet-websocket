package main

import (
	"net/http"
  "log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type connection struct{
  ws *websocket.Coon
}

func handler(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Println(err)
        return
    }

}

func (c *connection) writer() {
    for message := range c.send {
        err := c.ws.WriteMessage(websocket.TextMessage, message)
        if err != nil {
            break
        }
    }
    c.ws.Close()
}

func (c *connection) reader() {
    for {
      _,message,err := c.ws.ReadMessage()
      if err != nil {
        log.Println(err)
        break
      }
      h.broadcast <- message
    }
    c.ws.Close()
}