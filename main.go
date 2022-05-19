package main

import (
	. "github.com/starsky-mik/kgl-go-learing/lesson/di"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(GreetController)))
}
