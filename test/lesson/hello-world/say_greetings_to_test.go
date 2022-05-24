package hello_world

import (
	"fmt"
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/hello-world"
	"testing"
)

func TestSayGreetingsTo(t *testing.T) {
	t.Run("saying hello to Mike", func(t *testing.T) {
		AssertEquals(t, SayGreetingsTo("Mike", English), "Hello, Mike!")
	})

	t.Run("saying hello to Mike", func(t *testing.T) {
		AssertEquals(t, SayGreetingsTo("Mike", Spanish), "Hola, Mike!")
	})

	t.Run("saying hello to Mike", func(t *testing.T) {
		AssertEquals(t, SayGreetingsTo("Mike", French), "Bonjour, Mike!")
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		AssertEquals(t, SayGreetingsTo("", English), "Hello, World!")
	})
}

func ExampleSayGreetingsTo() {
	greetings := SayGreetingsTo("Bob", English)
	fmt.Println(greetings)
	// Output: Hello, Bob!
}

func BenchmarkSayGreetingsTo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayGreetingsTo("Bob", English)
	}
}
