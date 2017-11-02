package main

import (
	"log"
	"net/http"
	"flag"
	"fmt"
)

func main() {

	port := flag.String("port", "/dev/ttys006", "serial port ( example /dev/ttys006, etc)")
	addr := flag.String("addr", "localhost:4567", "web socket host and port ( example localhost:4567)")

	flag.Parse()

	fmt.Println("Websocket starting on: ", *addr)

	initDb()
	hub := newHub()
	go hub.run()
	go r(hub, port)

	http.HandleFunc("/", serveHome)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	http.ServeFile(w, r, "index.html")
}