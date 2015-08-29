package main

import (
	"os"
	"os/signal"
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
		w.WriteHeader(500)
		panic(err)
	}
	w.WriteHeader(200)
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter moveHandler")
	direction := mux.Vars(r)["direction"]
	err := bot.Move(direction)
	if err != nil {
		w.WriteHeader(500)
		panic(err)
	}
	w.WriteHeader(200)
}

func ledHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter ledHandler")
	color := mux.Vars(r)["color"]

	var err error
	if r.Method == "POST" {
		err = bot.LedOn(color)
	} else if r.Method == "DELETE" {
		err = bot.LedOff(color)
	}

	if err != nil {
		w.WriteHeader(500)
		panic(err)
	}
	w.WriteHeader(200)
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

	bot.Test()

	defer bot.Close()

	// catch interrupts so we close GPIO on Ctrl-C
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println("Received an interrupt, stopping...")
			bot.Close()
			os.Exit(0)
		}
	}()

	router := mux.NewRouter()
	router.HandleFunc("/look/{direction}", lookHandler).Methods("GET")
	router.HandleFunc("/move/{direction}", moveHandler).Methods("GET")
	router.HandleFunc("/leds/{color}", ledHandler).Methods("POST", "DELETE")
	http.Handle("/", router)

	host := hostname + ":" + port
	http.ListenAndServe(host, router)
}
