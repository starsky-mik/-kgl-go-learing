package hello_world

import (
	"fmt"
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	helloWorld "github.com/starsky-mik/kgl-go-learing/lesson/hello-world"
	"testing"
)

func TestSayGreetingsTo(t *testing.T) {
	t.Run("saying hello to Mike", func(t *testing.T) {
		testingTools.AssertEquals(t, helloWorld.SayGreetingsTo("Mike", helloWorld.English), "Hello, Mike!")
	})

	t.Run("saying hello to Mike", func(t *testing.T) {
		testingTools.AssertEquals(t, helloWorld.SayGreetingsTo("Mike", helloWorld.Spanish), "Hola, Mike!")
	})

	t.Run("saying hello to Mike", func(t *testing.T) {
		testingTools.AssertEquals(t, helloWorld.SayGreetingsTo("Mike", helloWorld.French), "Bonjour, Mike!")
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		testingTools.AssertEquals(t, helloWorld.SayGreetingsTo("", helloWorld.English), "Hello, World!")
	})
}

func ExampleSayGreetingsTo() {
	greetings := helloWorld.SayGreetingsTo("Bob", helloWorld.English)
	fmt.Println(greetings)
	// Output: Hello, Bob!
}

func BenchmarkSayGreetingsTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloWorld.SayGreetingsTo("Bob", helloWorld.English)
	}
}
