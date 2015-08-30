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
