package pointers_and_errors

import (
	. "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgl-go-learing/lesson/pointers-and-errors"
	"testing"
)

func TestWallet_Balance(t *testing.T) {
	wallet := Wallet{}
	AssertEquals(t, wallet.Balance(), Bitcoin(0))
}
