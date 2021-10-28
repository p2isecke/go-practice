package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		// Withdraw should return an error if you try to take out more than you have
		err := wallet.Withdraw(Bitcoin(100))

		// ErrInsufficientFunds is defined in the same package in wallet.go
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// check that if the Withdraw is successful that there is no error
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("should NOT have gotten an error")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	// Errors can be nil because the return type of Withdraw will be error, which is an interface
	if got == nil {
		// Will stop the test if invoked and avoid an unnecessary panic because of a nil pointer
		t.Fatal("should have gotten an error")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
