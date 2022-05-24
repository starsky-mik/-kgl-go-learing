package hello_world

import (
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/hello-world"
	"testing"
)

func TestGetGreetingsPrefix(t *testing.T) {
	t.Run("check helloWorld.GetGreetingsPrefix function", func(t *testing.T) {
		AssertEquals(t, GetGreetingsPrefix(English), "Hello")
	})
	t.Run("check helloWorld.GetGreetingsPrefix function", func(t *testing.T) {
		AssertEquals(t, GetGreetingsPrefix(Spanish), "Hola")
	})
	t.Run("check helloWorld.GetGreetingsPrefix function", func(t *testing.T) {
		AssertEquals(t, GetGreetingsPrefix(French), "Bonjour")
	})
}
