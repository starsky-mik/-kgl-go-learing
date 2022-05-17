package hello_world

import (
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	helloWorld "github.com/starsky-mik/kgl-go-learing/lesson/hello-world"
	"testing"
)

func TestGetGreetingsPrefix(t *testing.T) {
	t.Run("check helloWorld.GetGreetingsPrefix function", func(t *testing.T) {
		testingTools.AssertEquals(t, helloWorld.GetGreetingsPrefix(helloWorld.English), "Hello")
	})
	t.Run("check helloWorld.GetGreetingsPrefix function", func(t *testing.T) {
		testingTools.AssertEquals(t, helloWorld.GetGreetingsPrefix(helloWorld.Spanish), "Hola")
	})
	t.Run("check helloWorld.GetGreetingsPrefix function", func(t *testing.T) {
		testingTools.AssertEquals(t, helloWorld.GetGreetingsPrefix(helloWorld.French), "Bonjour")
	})
}
