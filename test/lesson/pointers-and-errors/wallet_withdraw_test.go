package pointers_and_errors

import (
	. "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgl-go-learing/lesson/pointers-and-errors"
	"testing"
)

func TestWallet_Withdraw(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(20))
	AssertNoError(t, wallet.Withdraw(Bitcoin(10)))
	AssertEquals(t, wallet.Balance(), Bitcoin(10))
}

func TestWallet_Withdraw_Errors(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	err := wallet.Withdraw(Bitcoin(100))

	AssertEquals(t, wallet.Balance(), Bitcoin(10))
	AssertError(t, err, InsufficientBalanceError)
}
