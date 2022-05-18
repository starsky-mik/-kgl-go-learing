package maps

import (
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	"github.com/starsky-mik/kgl-go-learing/lesson/maps"
	"testing"
)

func TestSet(t *testing.T) {
	dict := maps.Dictionary{"foo": "bar"}

	t.Run("check maps.dictionary.Set function", func(t *testing.T) {
		err := dict.Set("foo", "baz")
		testingTools.AssertNoError(t, err)

		res, err := dict.Search("foo")
		testingTools.AssertNoError(t, err)
		testingTools.AssertEquals(t, res, "baz")
	})
}
