package tldr

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to TL;DR pages!")
}

func TLDRIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TL;DR Index!")
	tldrs := Pages{
		Page{Name: "gcc"},
		Page{Name: "tar"},
	}

	json.NewEncoder(w).Encode(tldrs)
}

func TLDRItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tldrName := vars["tldrName"]
	fmt.Fprintln(w, "TL;DR name:", tldrName)
}
