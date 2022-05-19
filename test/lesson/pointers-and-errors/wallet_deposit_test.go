package pointers_and_errors

import (
	. "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgl-go-learing/lesson/pointers-and-errors"
	"testing"
)

func TestWallet_Deposit(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	AssertEquals(t, wallet.Balance(), Bitcoin(10))
}
