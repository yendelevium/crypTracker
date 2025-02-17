package websockets

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/socketio"
	"github.com/gofiber/fiber/v2"
	"github.com/yendelevium/crypTracker/models"
)

// This function will be called when I fetch API Data
// It will send the data to ALL clients via Broadcast
func SendCryptoData(coinData []models.Coin) {
	brodcastData, err := json.Marshal(coinData)
	if err != nil {
		log.Fatalf("Error marshalling WS broadcast to JSON: %s", err)
	}
	log.Println("Hitem BITCHES")
	socketio.Broadcast(brodcastData)
}

// Handler for /ws
// IDK the point of the callback
var WSServer = socketio.New(func(kws *socketio.Websocket) {
	log.Println("Socket Upgradation Endpoint")
	kws.Emit([]byte("Estabilished connection to crypTracker"))
})

// Basically just to setup the socketio.On() events
func WSRouter() func(*fiber.Ctx) error {
	// Multiple event handling supported
	socketio.On(socketio.EventConnect, func(ep *socketio.EventPayload) {
		log.Println("Connected to WS")
	})

	// On disconnect event
	socketio.On(socketio.EventDisconnect, func(ep *socketio.EventPayload) {
		log.Println("Disconnected from WS")
	})

	// On close event
	// This event is called when the server disconnects the user actively with .Close() method
	socketio.On(socketio.EventClose, func(ep *socketio.EventPayload) {
		log.Println("Closed WS")
	})

	// On error event
	socketio.On(socketio.EventError, func(ep *socketio.EventPayload) {
		log.Println("Error with WS")
	})

	// Testing message event
	socketio.On(socketio.EventMessage, func(ep *socketio.EventPayload) {
		log.Println(ep.Data)
		ep.Kws.Emit([]byte("Works?"))
	})

	return WSServer
}
