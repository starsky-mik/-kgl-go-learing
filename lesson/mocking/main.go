package mocking

func Countdown(operations CountdownOutputProxy) {
	for i := 3; i >= 1; i-- {
		operations.Sleep()
		_, _ = operations.Write(i)
	}

	operations.Sleep()
	_, _ = operations.Write("Go!")
}
