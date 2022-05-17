package pointers_and_errors

import (
	testingTools "github.com/starsky-mik/kgl-go-learing/internal/pkg/testing_tools"
	pointersAndErrors "github.com/starsky-mik/kgl-go-learing/lesson/pointers-and-errors"
	"testing"
)

func TestWallet_Withdraw_Errors(t *testing.T) {
	wallet := pointersAndErrors.Wallet{}
	wallet.Deposit(pointersAndErrors.Bitcoin(10))
	err := wallet.Withdraw(pointersAndErrors.Bitcoin(100))

	testingTools.AssertEquals(t, wallet.Balance(), pointersAndErrors.Bitcoin(10))
	testingTools.AssertError(t, err, pointersAndErrors.InsufficientBalanceError)
}
