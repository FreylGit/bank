package hello_handler

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
}