package main

import (
	. "github.com/starsky-mik/kgd-go-learing/lesson/di"
	"github.com/starsky-mik/kgd-go-learing/lesson/mocking"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mocking.Countdown(mocking.DefaultCountdownOutputProxy{
		Out:      os.Stdout,
		Interval: 1 * time.Second,
	})
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(GreetController)))
}
