package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, wanted Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != wanted {
			t.Errorf("got %s wanted %s", got, wanted)
		}
	}

	assertError := func(t *testing.T, got error, wanted error) {
		t.Helper()
		if got == nil {
			t.Errorf("No error but expeted %s", wanted)
		}

		if got != wanted {
			t.Errorf("got error %s wanted %s", got, wanted)
		}

	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(23)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(13))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficent", func(t *testing.T) {
		const startingBalance = 20
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(100)
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficentFunds)
	})

}

func assertNoError(t *testing.T, e error) {
	t.Helper()
	if e != nil {
		t.Errorf("Throw error %s while not expected!", e)
	}

}
