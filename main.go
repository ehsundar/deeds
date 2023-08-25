package main

import (
	"github.com/ehsundar/deeds/internal/server"
	"net/http"
	"os"
)

func main() {
	os.MkdirAll("images/", os.ModePerm)
	os.MkdirAll("tokens/", os.ModePerm)

	srv := &server.Server{}

	http.HandleFunc("/upload", srv.HandleFrom)
	http.HandleFunc("/confirm", srv.HandleConfirm)
	http.HandleFunc("/view", srv.HandleView)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	http.ListenAndServe(":8000", nil)
}
