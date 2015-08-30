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
	router.HandleFunc("/look/{direction}", lookHandler).Methods("POST")
	router.HandleFunc("/move/{direction}", moveHandler).Methods("POST")
	router.HandleFunc("/leds/{color:(green)}", ledHandler).Methods("POST", "DELETE")
	router.HandleFunc("/distance", distanceHandler).Methods("GET")
	router.HandleFunc("/status", statusHandler).Methods("GET")
	http.Handle("/", router)

	host := hostname + ":" + port
	http.ListenAndServe(host, router)
}
