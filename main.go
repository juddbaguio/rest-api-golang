package main

import (
	"log"
	"net/http"
)

func main() {

	srv := MuxRouter().NewServer()

	log.Fatal(http.ListenAndServe(":3000", srv))
}
