package wallet

import (
	"errors"
	"fmt"
)

// 给予现有的类型创建一个类型
type Bitcoin int

type Wallet struct {
	balance Bitcoin // 这个是私有的， 不能在外部使用
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address fo balance in deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

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
