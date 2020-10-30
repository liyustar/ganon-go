package gweb_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHelloGo(t *testing.T) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}