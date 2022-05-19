package di

import (
	"bytes"
	"github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	"github.com/starsky-mik/kgl-go-learing/lesson/di"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("check di.Greet function", func(t *testing.T) {
		buffer := bytes.Buffer{}
		err := di.Greet(&buffer, "Mike")

		testing_tools.AssertNoError(t, err)
		testing_tools.AssertEquals(t, buffer.String(), "Hello, Mike!")
	})
}
