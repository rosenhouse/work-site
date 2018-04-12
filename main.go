package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", okay)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func okay(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, os.Getenv("DEFAULT_REDIRECT"), http.StatusTemporaryRedirect)
}
