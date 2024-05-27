package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

type BankAccount struct {
	RWMutext sync.RWMutex
	Balance  int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutext.Lock()
	account.Balance = account.Balance + amount
	account.RWMutext.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutext.RLock()
	balance := account.Balance
	account.RWMutext.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total balance", account.GetBalance())
}
