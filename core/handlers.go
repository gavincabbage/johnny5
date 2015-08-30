package main

import (
    "os"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func lookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter lookHandler")
	direction := mux.Vars(r)["direction"]
	err := bot.Look(direction)
    statusCode := 200
	if err != nil {
		statusCode = 500
	}
	w.WriteHeader(statusCode)
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("enter moveHandler")
	direction := mux.Vars(r)["direction"]
	err := bot.Move(direction)
    statusCode := 200
	if err != nil {
		statusCode = 500
	}
	w.WriteHeader(statusCode)
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
    statusCode := 200
	if err != nil {
		statusCode = 500
	}
	w.WriteHeader(statusCode)
}

func distanceHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(204)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(204)
}

func errorHandler() {
	fmt.Println("enter errorHandler")
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	os.Exit(1)
}
