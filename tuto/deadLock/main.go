package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	balance int
	mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) {
	a.balance -= val
}

func (a *Account) Deposit(val int) {
	a.balance += val
}

func (a *Account) Balance() int {
	balance := a.balance
	return balance
}

var accounts []*Account

func Transfer(sender, reciver int, money int) {
	accounts[sender].mutex.Lock() // sender를 lock 걸었는데, reciver도 lock를 걸어서 서로 사용을 못함
	// deadlock 이 걸림, 그런데 보통 간헐적으로 발생하여 찾기 어려움
	// 그러므로 deadlock을 막기 위해서는 lock을 아예 작거나 크게 잡아야함
	// 작게 잡는다 : 잡았다가 빨리 푼다
	// 크게 잡는다 : 아예 멀리에 락을 걸어 한 사람씩 포크를 집어오게함
	// 컨베이어 벨트를 생각하면 좋음 -> 자기가 맡은 파트만 관리
	// = 생산자-소비자 패턴 = Producer-Consumer pattern
	fmt.Println("Lock", sender)
	time.Sleep(1000 * time.Millisecond)
	accounts[reciver].mutex.Lock()
	fmt.Println("Lock", reciver)

	accounts[sender].Widthdraw(money)
	accounts[reciver].Deposit(money)

	accounts[sender].mutex.Unlock()
	accounts[reciver].mutex.Unlock()

	fmt.Println("Transfer", sender, reciver, money)
}

func GetTotalBalance() int {
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	return total
}

func RandomTransfer() {
	var sender, balance int
	for {
		sender = rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 {
			break
		}
	}

	var reciver int
	for {
		reciver = rand.Intn(len(accounts))
		if sender != reciver {
			break
		}
	}

	money := rand.Intn(balance)
	Transfer(sender, reciver, money)
}

func GoTransfer() {
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Println("Total: ", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ {
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}})
	}

	go func() {
		for {
			Transfer(0, 1, 100)
		}
	}()

	go func() {
		for {
			Transfer(1, 0, 100)
		}
	}()

	for {
		//PrintTotalBalance()
		time.Sleep(100 * time.Millisecond)
	}
}
