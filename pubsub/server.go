package pubsub

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

var upgrader = websocket.Upgrader{}
var ps = &PubSub{}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {

	id := uuid.Must(uuid.NewV4()).String()

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()

	// add new client
	client := Client{
		Id:   id,
		Conn: c,
	}
	ps.AddClient(client)

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)

			// handle remove client and subscriptions
			ps.RemoveClient(client)

			break
		}
		log.Printf("recv: %s, clients: %d", message, len(ps.Clients))

		ps.HandleReceivedMessage(&client, mt, message)

		if err != nil {
			log.Println("write:", err)
			break
		}
	}

}
