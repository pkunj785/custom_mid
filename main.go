package main

import (
	h "curs_custom/handler"
	m "curs_custom/middlewares"
	"log"
	"net/http"

	"github.com/justinas/alice"
)

func main() {

	port := ":8080"

	originalHandler := http.HandlerFunc(h.Handle)

	chain := alice.New(m.FilterContenType, m.SetServerTimeCookie).Then(originalHandler)

	http.Handle("/lesson", chain)
	log.Fatal(http.ListenAndServe(port, nil))
}
