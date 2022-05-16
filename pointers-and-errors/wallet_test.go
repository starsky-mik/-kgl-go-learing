package pointers_and_errors

import "testing"

func assertEquals(t testing.TB, actual, expected any) {
	t.Helper()

	if actual != expected {
		t.Errorf("expected '%s' but got '%s'", expected, actual)
	}
}

func assertError(t testing.TB, err, expected error) {
	t.Helper()

	if err == nil {
		t.Errorf("expected error \"%v\" but got \"%v\"", expected, err)
	}

	if err != expected {
		t.Errorf("expected error \"%v\" but got \"%v\"", expected, err)
	}
}

func TestWallet_Balance(t *testing.T) {
	wallet := Wallet{}
	assertEquals(t, wallet.Balance(), Bitcoin(0))
}

func TestWallet_Deposit(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))
	assertEquals(t, wallet.Balance(), Bitcoin(10))
}

func TestWallet_Withdraw(t *testing.T) {
	wallet := Wallet{balance: Bitcoin(20)}
	_ = wallet.Withdraw(Bitcoin(10))
	assertEquals(t, wallet.Balance(), Bitcoin(10))
}

func TestWallet_Withdraw_Errors(t *testing.T) {
	wallet := Wallet{balance: Bitcoin(10)}
	err := wallet.Withdraw(Bitcoin(100))
	assertEquals(t, wallet.Balance(), Bitcoin(10))
	assertError(t, err, InsufficientBalanceError)
}
