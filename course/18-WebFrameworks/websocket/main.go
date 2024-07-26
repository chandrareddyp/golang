package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
its based on the article: https://tutorialedge.net/golang/go-websocket-tutorial/
youtube video : https://www.youtube.com/watch?v=dniVs0xKYKk&t=24s
*/
func main() {
	fmt.Println("Hello, playground")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupRoutes(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home Page")
}

func reader(conn *websocket.Conn){
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(string(p))
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Websocket Endpoint")

	upgrade.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	reader(ws)
}

var wu websocket.Upgrader

var upgrade = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}