package learnGoWithTests

import (
	"errors"
	"fmt"
)

// insufficient funds error
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

//自定义比特币
type Bitcoin int

//存钱
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//查看余额
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//转账
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}
