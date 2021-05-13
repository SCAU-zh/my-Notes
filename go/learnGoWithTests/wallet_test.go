package learnGoWithTests

import "testing"

func TestWallet(t *testing.T) {

	//测试存钱
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	//测试转账成功，余额足够
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	//测试转账失败，余额不足
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(30)
		want := Bitcoin(20)
		assertBalance(t, wallet, want)
		assertError(t, err, "cannot withdraw, insufficient funds")
	})

}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want string) {
	if got == nil {
		t.Fatal("wanted an error but didnt get one")
	}
	if got.Error() != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("not error but  get one")
	}
}
