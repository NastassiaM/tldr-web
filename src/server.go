package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tldr"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to TL;DR pages!")
}

func TLDRIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TL;DR Index!")
	tldrs := tldr.Pages{
		tldr.Page{Name: "gcc"},
		tldr.Page{Name: "tar"},
	}

	json.NewEncoder(w).Encode(tldrs)
}

func TLDRItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tldrName := vars["tldrName"]
	fmt.Fprintln(w, "TL;DR name:", tldrName)
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/tldr", TLDRIndex)
	router.HandleFunc("/tldr/{tldrName}", TLDRItem)

	log.Fatal(http.ListenAndServe(":8081", router))

}
