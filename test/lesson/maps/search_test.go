package maps

import (
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	"github.com/starsky-mik/kgl-go-learing/lesson/maps"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := maps.Dictionary{"foo": "bar"}

	t.Run("check maps.Dictionary.Search function", func(t *testing.T) {
		res, err := dict.Search("foo")

		testingTools.AssertNoError(t, err)
		testingTools.AssertEquals(t, res, "bar")
	})
}

func TestSearchErrors(t *testing.T) {
	dict := maps.Dictionary{"foo": "bar"}

	t.Run("check maps.Dictionary.Search function errors", func(t *testing.T) {
		res, err := dict.Search("baz")

		testingTools.AssertEquals(t, res, "")
		testingTools.AssertError(t, err, maps.KeyNotExistsError)
	})
}
