package gweb_test

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestHelloGo3(t *testing.T) {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(wd)

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err = server.ListenAndServeTLS("cert.pem", "private.pem")
	if err != nil {
		log.Fatalln(err)
	}
}
