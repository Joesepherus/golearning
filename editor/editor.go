package editor

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections
	},
}

var clients = make(map[*websocket.Conn]bool)

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./editor/index.html")
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true
	log.Println("clients", clients)

	for {
		// Read message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			break
		}

		// Broadcast message to all clients
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func Run() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ws", handleConnections)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
