package main

import (
	"fmt"
	"net/http"
	"websocket/internal/app"
	"websocket/internal/transport"
)

func main() {
	http.HandleFunc("/echo", app.WebSocketHandler)
	http.HandleFunc("/", transport.HttpHandler)

	fmt.Println("Server 8080-portda ishga tushdi...")
	http.ListenAndServe(":8080", nil)
}
