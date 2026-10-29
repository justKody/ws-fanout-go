package main

import (
	"encoding/json"
	"log"
	"sync"
)

type Message struct {
	ClientID string
	Text     string
	SendTo   string
}

type WSMessage struct {
	Text    string      `json:"text"`
	SendTo  string      `json:"send_to"`
	Headers interface{} `json: "HEADERS"`
}

type Hub struct {
	sync.RWMutex
	clients    map[*Client]bool
	broadcast  chan *Message
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.Lock()
			h.clients[client] = true
			h.Unlock()

			log.Printf("client registered %s", client.id)
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				close(client.send)
				delete(h.clients, client)
			}
		case msg := <-h.broadcast:
			payload, err := json.Marshal(msg)
			if err != nil {
				log.Printf("broadcast: %v", err)
				continue
			}
			if msg.SendTo != "" {
				if client := h.findClient(msg.SendTo); client != nil {
					client.send <- payload
				}
			}

			h.Lock()
			for client := range h.clients {
				select {
				case client.send <- payload:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
			h.Unlock()
		}
	}
}

func (h *Hub) findClient(id string) *Client {
	for client := range h.clients {
		if client.id == id {
			return client
		}
	}
	return nil
}
