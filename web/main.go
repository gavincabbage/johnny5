package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting static server")
	dir := "/home/johnny"
	panic(http.ListenAndServe("0.0.0.0:8080", http.FileServer(http.Dir(dir))))
}
