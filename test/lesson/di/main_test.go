package di

import (
	"bytes"
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/di"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("check di.Greet function", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := Greet(&buffer, "Mike")

		AssertNoError(t, err)
		AssertEquals(t, buffer.String(), "Hello, Mike!")
	})
}
