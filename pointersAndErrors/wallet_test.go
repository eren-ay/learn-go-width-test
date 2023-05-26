package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	myWallet := Wallet{}
	checkBalance := func(t testing.TB, wallet Wallet, want float64) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()

		if err != nil {
			t.Fatal("error is  ", err.Error())
		}
	}

	t.Run("Wallet deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet = myWallet
		deposit := 10.0
		myWallet.Deposit(deposit)
		want := wallet.Balance() + deposit
		checkBalance(t, myWallet, want)
	})

	t.Run("Wallet withdraw", func(t *testing.T) {
		wallet := myWallet
		withdrawAmount := 50.0
		err := myWallet.Withdraw(withdrawAmount)
		want := wallet.Balance() - withdrawAmount
		assertError(t, err)
		checkBalance(t, myWallet, want)
	})

}
