package main

import (
	"db"
	"log"
	"net/http"
	"tldr"
)

func main() {

	db.Init()
	defer db.Close()
	router := tldr.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))

}
