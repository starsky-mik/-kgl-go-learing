package di

import (
	"fmt"
	. "github.com/starsky-mik/kgl-go-learing/lesson/hello-world"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string) error {
	_, err := fmt.Fprintf(w, SayGreetingsTo(name, English))
	return err
}

func GreetController(w http.ResponseWriter, _ *http.Request) {
	_ = Greet(w, "Mike")
}
