package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

const (
	REDIS_HOST string = "127.0.0.1:4290"
)

var MOVE map[string]byte = map[string]byte{
	"forward" : 10,
	"back" : 11,
	"left" : 12,
	"right" : 13,
	"stop" : 14,
}

var LOOK map[string]byte = map[string]byte{
	"center" : 20,
	"left" : 21,
	"right" : 22,
	"up" : 23,
	"down" : 24,
}

var (
	arduino1, arduino2 byte
	hostname, port string
	bus embd.I2CBus
	router *mux.Router
)

func lookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENTER lookHandler")
	direction := mux.Vars(r)["direction"]
	bus.WriteByte(arduino2, LOOK[direction])
	w.Write(nil)
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENTER moveHandler")
	direction := mux.Vars(r)["direction"]
	bus.WriteByte(arduino1, MOVE[direction])
	w.Write(nil)
}

func errorHandler() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	os.Exit(1)
}

func main() {

	defer errorHandler()

	bus = embd.NewI2CBus(1)

	router = mux.NewRouter()
	router.HandleFunc("/look/{direction}", lookHandler).Methods("GET")
	router.HandleFunc("/move/{direction}", moveHandler).Methods("GET")
	http.Handle("/", router)

	host := "127.0.0.1:8080"
	http.ListenAndServe(host, router)
}
