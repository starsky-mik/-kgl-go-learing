package pointers_and_errors

type WalletError uint8

const (
	InsufficientBalanceError WalletError = iota
)

func (w WalletError) Error() string {
	switch w {
	case InsufficientBalanceError:
		return "Insufficient balance error"
	}

	return "Undefined error"
}
