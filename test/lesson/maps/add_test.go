package maps

import (
	. "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgl-go-learing/lesson/maps"
	"testing"
)

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	t.Run("check maps.dictionary.Add function", func(t *testing.T) {
		err := dict.Add("foo", "bar")
		AssertNoError(t, err)

		res, err := dict.Search("foo")
		AssertNoError(t, err)
		AssertEquals(t, res, "bar")
	})
}

func TestAddErrors(t *testing.T) {
	dict := Dictionary{"foo": "bar"}

	t.Run("check maps.dictionary.Add function errors", func(t *testing.T) {
		err := dict.Add("foo", "baz")
		AssertError(t, err, KeyAlreadyExistsError)
	})
}
