package main

import (
	"db"
	"log"
	"net/http"
	"tldr"
)

func main() {

	router := tldr.NewRouter()
	db.Init()
	log.Fatal(http.ListenAndServe(":8081", router))

}
