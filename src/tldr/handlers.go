package tldr

import (
	"db"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"model"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to TL;DR pages!")
}

func tldrIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(db.Tldrs); err != nil {
		panic(err)
	}
}

func tldrItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tldrName := vars["tldrName"]
	fmt.Fprintln(w, "TL;DR name:", tldrName)
}

func tldrCreate(w http.ResponseWriter, r *http.Request) {
	var tldrPage model.Page
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &tldrPage); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	db.AddPage(tldrPage)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(tldrPage); err != nil {
		panic(err)
	}
}
