package gweb_test

import (
	"fmt"
	"log"
	"net/http"
	"nuts/gweb"
	"nuts/gweb/data"
	"os"
	"testing"
)

func TestHelloGo2(t *testing.T) {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(wd)

	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)

	mux.HandleFunc("/authenticate", gweb.Authenticate)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	if threads, err := data.Threads(); err == nil {
		_, err := gweb.Session(w, r)
		if err == nil {
			err = gweb.GenerateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			err = gweb.GenerateHTML(w, threads, "layout", "private.navbar", "index")
		}
	} else {
		log.Fatalln(err)
	}
}
