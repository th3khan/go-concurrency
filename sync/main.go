package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	lock.Lock()
	fmt.Println("Depositing:", amount)
	b := balance
	balance = b + amount
	lock.Unlock()
}

func Balance(lock *sync.RWMutex) int {
	lock.RLock()
	b := balance
	lock.RUnlock()
	return b
}

func ShowBalance(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Balance:", balance)
}

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	for i := 1; i <= 5; i++ {
		wg.Add(2)
		go Deposit(i*100, &wg, &lock)
		go ShowBalance(&wg)
	}
	wg.Wait()
	fmt.Println("------------------------")
	fmt.Println("Final balance:", Balance(&lock))
}
