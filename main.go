package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/ehsundar/deeds/internal/server"
)

var (
	port = flag.Int("port", 8000, "port")
)

func main() {
	os.MkdirAll("images/", os.ModePerm)
	os.MkdirAll("tokens/", os.ModePerm)

	srv := &server.Server{}

	http.HandleFunc("/upload", srv.HandleFrom)
	http.HandleFunc("/confirm", srv.HandleConfirm)
	http.HandleFunc("/view", srv.HandleView)

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), nil)
}
