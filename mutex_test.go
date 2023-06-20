package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x:=0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func(){
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ",x)
}

type BankAccount struct{
	RWMutext sync.RWMutex
	balance int
}

func (account *BankAccount) addBalance(amount int){
	account.RWMutext.Lock()
	account.balance = account.balance + amount
	account.RWMutext.Unlock()
}

func (account *BankAccount) getBalance() int {
	account.RWMutext.RLock()
	balance:= account.balance
	account.RWMutext.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account:= BankAccount{}

	for i := 0; i < 100; i++ {
		go func(){
			for j := 0; j < 100; j++ {
				account.addBalance(1)
				fmt.Println(account.getBalance())
			}
		}()
	}
	time.Sleep(5*time.Second)
	fmt.Println("Total Account = ",account.getBalance())

}