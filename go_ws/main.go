package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type updateWS struct {
	upgraders *websocket.Upgrader
}

// CREATE VARIABLE WEBSOCKET

// var clients []websocket.Conn

func main() {
	fmt.Printf("type of a is %T\n", upgrader)

	// CREATE ENDPOIND FOR CONNECT WEBSOKCET
	http.HandleFunc("/websoc", func(w http.ResponseWriter, r *http.Request) {
		// INITIALIZE CONFIG
		conn, _ := upgrader.Upgrade(w, r, nil)

		defer conn.Close()

		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			conn.WriteMessage(websocket.TextMessage, []byte("Hello World"))
		}

	})
	// SEND YOU HTML FILE FOR OPEN TO BROWSER
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
		// w,r IS WRITE AND DELETE YOU INDEX.HTML

	})
	println("You server run 8080")
	http.ListenAndServe(":8080", nil)

}
