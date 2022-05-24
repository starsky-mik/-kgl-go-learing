package maps

import (
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/maps"
	"testing"
)

func TestRemove(t *testing.T) {
	t.Run("check maps.dictionary.Add function", func(t *testing.T) {
		dict := Dictionary{"foo": "bar"}
		err := dict.Remove("foo")
		AssertNoError(t, err)

		res, err := dict.Search("foo")
		AssertEquals(t, res, "")
		AssertError(t, err, KeyNotExistsError)
	})
}

func TestRemoveErrors(t *testing.T) {
	t.Run("check maps.dictionary.Remove function", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Remove("foo")
		AssertError(t, err, KeyNotExistsError)
	})
}
