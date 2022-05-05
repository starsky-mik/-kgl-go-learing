package main

import (
	"fmt"
	"testing"
)

func assertEquals(t testing.TB, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("function result(%q) does`t match expected value(%q)", actual, expected)
	}
}

func TestGetGreetingsPrefix(t *testing.T) {
	t.Run("get english greetings prefix", func(t *testing.T) {
		assertEquals(t, GetGreetingsPrefix(English), "Hello")
	})
	t.Run("get spanish greetings prefix", func(t *testing.T) {
		assertEquals(t, GetGreetingsPrefix(Spanish), "Hola")
	})
	t.Run("get french greetings prefix", func(t *testing.T) {
		assertEquals(t, GetGreetingsPrefix(French), "Bonjour")
	})
}

func TestSayGreetingsTo(t *testing.T) {
	t.Run("saying hello to Mike", func(t *testing.T) {
		assertEquals(t, SayGreetingsTo("Mike", English), "Hello, Mike!")
	})

	t.Run("saying hello to Mike", func(t *testing.T) {
		assertEquals(t, SayGreetingsTo("Mike", Spanish), "Hola, Mike!")
	})

	t.Run("saying hello to Mike", func(t *testing.T) {
		assertEquals(t, SayGreetingsTo("Mike", French), "Bonjour, Mike!")
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		actual := SayGreetingsTo("", English)
		expected := "Hello, World!"
		assertEquals(t, actual, expected)
	})
}

func ExampleSayGreetingsTo() {
	greetings := SayGreetingsTo("Bob", English)
	fmt.Println(greetings)
	// Output: Hello, Bob!
}
