package maps

import (
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/maps"
	"testing"
)

func TestSet(t *testing.T) {
	dict := Dictionary{"foo": "bar"}

	t.Run("check maps.dictionary.Set function", func(t *testing.T) {
		err := dict.Set("foo", "baz")
		AssertNoError(t, err)

		res, err := dict.Search("foo")
		AssertNoError(t, err)
		AssertEquals(t, res, "baz")
	})
}
