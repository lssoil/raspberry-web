package main

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to lssoil's home"))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8000", nil)
}
