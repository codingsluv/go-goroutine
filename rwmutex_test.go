package belajargoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRwMutex(t *testing.T) {
	acc := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				acc.AddBalance(1)
				fmt.Println("Current balance:", acc.GetBalance())
			}
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Final balance:", acc.GetBalance())
}
