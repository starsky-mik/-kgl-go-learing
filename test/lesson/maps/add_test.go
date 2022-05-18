package maps

import (
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	"github.com/starsky-mik/kgl-go-learing/lesson/maps"
	"testing"
)

func TestAdd(t *testing.T) {
	dict := maps.Dictionary{}

	t.Run("check maps.dictionary.Add function", func(t *testing.T) {
		err := dict.Add("foo", "bar")
		testingTools.AssertNoError(t, err)

		res, err := dict.Search("foo")
		testingTools.AssertNoError(t, err)
		testingTools.AssertEquals(t, res, "bar")
	})
}

func TestAddErrors(t *testing.T) {
	dict := maps.Dictionary{"foo": "bar"}

	t.Run("check maps.dictionary.Add function errors", func(t *testing.T) {
		err := dict.Add("foo", "baz")
		testingTools.AssertError(t, err, maps.KeyAlreadyExistsError)
	})
}
