package maps

import (
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	"github.com/starsky-mik/kgl-go-learing/lesson/maps"
	"testing"
)

func TestRemove(t *testing.T) {
	t.Run("check maps.dictionary.Add function", func(t *testing.T) {
		dict := maps.Dictionary{"foo": "bar"}
		err := dict.Remove("foo")
		testingTools.AssertNoError(t, err)

		res, err := dict.Search("foo")
		testingTools.AssertEquals(t, res, "")
		testingTools.AssertError(t, err, maps.KeyNotExistsError)
	})
}

func TestRemoveErrors(t *testing.T) {
	t.Run("check maps.dictionary.Remove function", func(t *testing.T) {
		dict := maps.Dictionary{}
		err := dict.Remove("foo")
		testingTools.AssertError(t, err, maps.KeyNotExistsError)
	})
}
