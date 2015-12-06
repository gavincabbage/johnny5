package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting static server")
	dir := "/home/gavin/Development/johnny5/dist/web"
	panic(http.ListenAndServe("127.0.0.1:8080", http.FileServer(http.Dir(dir))))
}
