package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

var (
	bot CoreBot
	hostname, port string = "0.0.0.0", "8080"
)

func lookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter lookHandler")
	direction := mux.Vars(r)["direction"]
	err := bot.Look(direction)
	if err != nil {
		panic(err)
	}
	w.Write(nil)
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter moveHandler")
	direction := mux.Vars(r)["direction"]
	err := bot.Move(direction)
	if err != nil {
		panic(err)
	}
	w.Write(nil)
}

func errorHandler() {
	fmt.Println("enter errorHandler")
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	os.Exit(1)
}

func main() {

	defer errorHandler()

	bot = NewCoreBot()

	router := mux.NewRouter()
	router.HandleFunc("/look/{direction}", lookHandler).Methods("GET")
	router.HandleFunc("/move/{direction}", moveHandler).Methods("GET")
	http.Handle("/", router)

	host := hostname + ":" + port
	http.ListenAndServe(host, router)
}
