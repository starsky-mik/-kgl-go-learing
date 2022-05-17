package pointers_and_errors

import (
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	pointersAndErrors "github.com/starsky-mik/kgl-go-learing/lesson/pointers-and-errors"
	"testing"
)

func TestWallet_Balance(t *testing.T) {
	wallet := pointersAndErrors.Wallet{}
	testingTools.AssertEquals(t, wallet.Balance(), pointersAndErrors.Bitcoin(0))
}
