package belajargoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (u *UserBalance) Lock() {
	u.Mutex.Lock()
}

func (u *UserBalance) Unlock() {
	u.Mutex.Unlock()
}

func (u *UserBalance) Change(amount int) {
	u.Balance = u.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user2.Unlock()
	user1.Unlock()

}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "user1",
		Balance: 100000,
	}
	user2 := UserBalance{
		Name:    "user2",
		Balance: 100000,
	}

	go Transfer(&user1, &user2, 10000)
	go Transfer(&user2, &user1, 20000)

	time.Sleep(5 * time.Second)

	fmt.Println("User", user1.Name, "Balance", user1.Balance)
	fmt.Println("User", user2.Name, "Balance", user2.Balance)

}
