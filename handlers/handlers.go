package handlers

import (
	"chatApp/helpers"
	"chatApp/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client represents a connected WebSocket client.
type Client struct {
	conn   *websocket.Conn
	userID string
}

var clients = make(map[*Client]bool)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var emp models.User
	w.Header().Set("content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Can't read the request", err)
		return
	}

	if err = helpers.DB.Create(&emp).Error; err != nil {
		http.Error(w, "Can't insert value into the Table", http.StatusInternalServerError)
		fmt.Println("Can't insert value into the Table", err)
		return
	}

	response, err := json.Marshal(emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error During encoding the response", err)
		return
	}

	fmt.Println("Response", string(response))

	w.Write(response)
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		fmt.Println("User ID is required.")
		return
	}

	var user models.User

	result := helpers.DB.Where("id=?", userID).Find(&user)
	if result.Error != nil {
		fmt.Printf("%s does not exist in Database", userID)
		return
	}

	client := &Client{conn, userID}

	clients[client] = true

	log.Printf("User %s connected\n", userID)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("User %s disconnected\n", userID)
			delete(clients, client)
			return
		}

		// Split the message into 2 parts: [receiver : message]
		parts := strings.SplitN(string(p), ":", 2)
		if len(parts) != 2 {
			log.Printf("Invalid message format: %s\n", string(p))
			continue
		}

		receiver := parts[0]
		message := parts[1]

		for c := range clients {
			if c.userID == receiver {
				err := c.conn.WriteMessage(messageType, []byte(message))
				if err != nil {
					log.Println(err)
					return
				}
			}
		}
	}
}
