package main

import (
	"log"
	"net/http"
	"tldr"
)

func main() {

	router := tldr.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))

}
