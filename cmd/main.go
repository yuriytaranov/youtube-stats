package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/yuriytaranov/youtube-stats/internal/pkg/websocket"
)

// homePage will be a simple "hello World" style page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// our new stats function which will expose any YouTube
// stats via a websocket connection
func stats(w http.ResponseWriter, r *http.Request) {
	// we call our new websocket package Upgrade
	// function in order to upgrade the connection
	// from a standard HTTP connection to a websocket one
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Println(w, "%+v\n", err)
	}
	// we the call our Writer function
	// which continually polls and writes the results
	// to this websocket connection
	go websocket.Writer(ws)
}

// setupRoutes handles setting up our servers
// routes and matching them to their respective
// functions
func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)
	// here we kick off our server on localhost:8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// our main function
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("YouTube Subscriber Monitor")
	// call setup routes
	setupRoutes()
}
