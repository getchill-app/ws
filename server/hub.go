package server

import (
	"log"

	"github.com/getchill-app/ws/api"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	url  string
	auth string

	rds *Redis

	// Registered clients.
	clients map[string]*client

	// Clients by token.
	clientsByToken map[string]map[string]*client

	// Inbound messages.
	broadcastCh chan *api.Event

	// Inbound client.
	clientCh chan *registerClient

	// Register client from the clients.
	registerCh chan *client

	// Unregister requests from clients.
	unregisterCh chan *client
}

type registerClient struct {
	client *client
	token  string
}

// NewHub ...
func NewHub(url string, auth string) *Hub {
	h := &Hub{
		url:            url,
		auth:           auth,
		broadcastCh:    make(chan *api.Event),
		clientCh:       make(chan *registerClient, 10),
		registerCh:     make(chan *client),
		unregisterCh:   make(chan *client),
		clients:        make(map[string]*client),
		clientsByToken: make(map[string]map[string]*client),
	}
	return h
}

func (h *Hub) Auth() string {
	return h.auth
}

// Run hub.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.registerCh:
			log.Printf("register %s\n", client.id)
			h.clients[client.id] = client
		case client := <-h.unregisterCh:
			log.Printf("unregister %s\n", client.id)
			if _, ok := h.clients[client.id]; ok {
				h.unregisterClient(client)
				close(client.send)
			}
		case registerClient := <-h.clientCh:
			// log.Printf("register %s\n", auth.client.id)
			h.registerClient(registerClient)
			// auth.client.send <- &api.Event{}
		case event := <-h.broadcastCh:
			clients := h.findClients(event.Token)
			h.send(clients, event)
		}
	}
}

func (h *Hub) send(clients []*client, event *api.Event) {
	for _, client := range clients {
		select {
		case client.send <- event:
			// log.Printf("send %s => %s\n", client.id, event)
		default:
			close(client.send)
			h.unregisterClient(client)
		}
	}
}

func (h *Hub) registerClient(registerClient *registerClient) {
	cl, token := registerClient.client, registerClient.token
	if cl.tokens == nil {
		cl.tokens = []string{}
	}
	cl.tokens = append(cl.tokens, token)
	clients, ok := h.clientsByToken[token]
	if !ok {
		clients = map[string]*client{}
		h.clientsByToken[token] = clients
	}
	clients[cl.id] = cl
}

func (h *Hub) unregisterClient(cl *client) {
	for _, token := range cl.tokens {
		clientsByToken, ok := h.clientsByToken[token]
		if !ok {
			continue
		}
		delete(clientsByToken, cl.id)
	}
	delete(h.clients, cl.id)
	cl.tokens = nil
}

func (h *Hub) findClients(token string) []*client {
	clients, ok := h.clientsByToken[token]
	if !ok {
		return nil
	}
	out := make([]*client, 0, len(clients))
	for _, cl := range clients {
		if _, ok := h.clients[cl.id]; !ok {
			continue
		}
		out = append(out, cl)
	}
	return out
}
