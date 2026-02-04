package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("listening on :8888")
	log.Fatal(http.ListenAndServe(":8888", router()))
}
