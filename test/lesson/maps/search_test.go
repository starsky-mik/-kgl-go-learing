package maps

import (
	. "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgl-go-learing/lesson/maps"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"foo": "bar"}

	t.Run("check maps.Dictionary.Search function", func(t *testing.T) {
		res, err := dict.Search("foo")

		AssertNoError(t, err)
		AssertEquals(t, res, "bar")
	})
}

func TestSearchErrors(t *testing.T) {
	dict := Dictionary{"foo": "bar"}

	t.Run("check maps.Dictionary.Search function errors", func(t *testing.T) {
		res, err := dict.Search("baz")

		AssertEquals(t, res, "")
		AssertError(t, err, KeyNotExistsError)
	})
}
