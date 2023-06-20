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

type UserBalance struct{
	sync.Mutex
	name string
	balance int
}

func (user *UserBalance) Lock(){
	user.Mutex.Lock()
}

func (user *UserBalance) UnLock(){
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int){
	user.balance = user.balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int){
	user1.Lock()
	fmt.Println("Lock User 1 ",user1.name)
	user1.Change(-amount)

	time.Sleep(1*time.Second)

	user2.Lock()
	fmt.Println("Lock User 2 ",user2.name)
	user2.Change(amount)

	time.Sleep(1*time.Second)

	user1.UnLock()
	user2.UnLock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		name: "Sandy",
		balance: 1000000,
	}

	user2 := UserBalance{
		name: "Azhi",
		balance: 1000000,
	}

	go Transfer(&user1,&user2,100000)
	go Transfer(&user2,&user1,200000)

	time.Sleep(10*time.Second)

	fmt.Println("User ",user1.name," Balance ",user1.balance)
	fmt.Println("User ",user2.name," Balance ",user2.balance)

}