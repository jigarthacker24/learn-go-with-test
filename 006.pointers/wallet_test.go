package pointers

import (
	"testing"
)

func checkBalanceAssert(tb testing.TB, w Wallet, want Bitcoin) {
	tb.Helper()
	got := w.Balance()
	if got != want {
		tb.Errorf("Balance is not as expected. got: %s, want: %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Error("Error expected. No error received")
	}

	if got != want {
		t.Errorf("Incorrect error. got:%q, want:%q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("No error expected. Error received: %q", got)
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		checkBalanceAssert(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 25}
		err := wallet.Withdraw(Bitcoin(10))
		checkBalanceAssert(t, wallet, Bitcoin(15))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient fund", func(t *testing.T) {
		initialFund := Bitcoin(25)
		wallet := Wallet{balance: initialFund}
		err := wallet.Withdraw(Bitcoin(100))

		checkBalanceAssert(t, wallet, initialFund)
		assertError(t, err, ErrNotEnoughFund)
	})
}
