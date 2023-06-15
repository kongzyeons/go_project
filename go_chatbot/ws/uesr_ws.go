package ws

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
)

// type UserWs struct {
// 	Conn     *websocket.Conn
// 	Message  chan *Message
// 	ID       string `json:"id"`
// 	RoomID   string `json:"roomId"`
// 	Username string `json:"username"`
// }

// type Message struct {
// 	Content  string `json:"content"`
// 	RoomID   string `json:"roomId"`
// 	Username string `json:"username"`
// }

var clients []websocket.Conn

func GetWs(ctx *websocket.Conn) {

	// ctx.WriteMessage(websocket.TextMessage, []byte("6666"))ß
	// for {
	// 	_, msg, err := ctx.ReadMessage()
	// 	if err != nil {
	// 		return
	// 	}
	// 	fmt.Printf("%s send: %s\n", ctx.RemoteAddr(), string(msg))
	// 	if err = ctx.WriteMessage(websocket.TextMessage, msg); err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }

	clients = append(clients, *ctx)

	// LOOP IF CLIENT SEND TO SERVERß
	for {
		// READ MESSAGE FROM BROWSER
		msgType, msg, err := ctx.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		// PRINT MESSAGE IN YOU CONSOLE TERMINAL
		fmt.Printf("%s send: %s\n", ctx.RemoteAddr(), string(msg))

		if err = clients[0].WriteMessage(msgType, msg); err != nil {
			return

		}

		// LOOP IF MESSAGE FOUND AND SEND AGAIN TO CLIENT FOR
		// WRITE IN YOU BROWSER
		// for _, client := range clients {
		// 	if err = client.WriteMessage(msgType, msg); err != nil {
		// 		return

		// 	}
		// }

	}

}
