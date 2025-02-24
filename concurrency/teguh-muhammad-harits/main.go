package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Wallet struct {
	balance int
	mu      sync.Mutex
}

func (w *Wallet) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	w.mu.Lock()
	defer w.mu.Unlock()
	w.balance += amount
	fmt.Println("Deposit:", amount, "Saldo:", w.balance)
}

func (w *Wallet) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.balance >= amount {
		w.balance -= amount
		fmt.Println("Withdraw:", amount, "Saldo:", w.balance)
	} else {
		fmt.Println("Saldo tidak mencukupi untuk withdraw", amount)
	}
}

func demonstrateGoroutines(wallet *Wallet, wg *sync.WaitGroup) {
	fmt.Println("\n1. Demonstrasi Goroutine")
	wg.Add(2)
	go wallet.Deposit(500, wg)
	go wallet.Withdraw(300, wg)
	wg.Wait()
}

func demonstrateChannels() {
	fmt.Println("\n2. Demonstrasi Channel (Buffered dan Unbuffered)")
	transactions := make(chan string, 3)
	transactions <- "Deposit 200"
	transactions <- "Withdraw 150"
	transactions <- "Deposit 100"
	close(transactions)

	for transaction := range transactions {
		fmt.Println("Transaksi diproses:", transaction)
	}
}

func demonstrateWaitGroup(wallet *Wallet, wg *sync.WaitGroup) {
	fmt.Println("\n3. Demonstrasi sync.WaitGroup")
	wg.Add(2)
	go wallet.Deposit(700, wg)
	go wallet.Withdraw(400, wg)
	wg.Wait()
	fmt.Println("Saldo akhir:", wallet.balance)
}

func demonstrateMutex() {
	fmt.Println("\n4. Demonstrasi sync.Mutex")
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			fmt.Println("Counter:", counter)
			mu.Unlock()
		}()
	}
	wg.Wait()
}

func demonstrateSelectTimeout() {
	fmt.Println("\n5. Demonstrasi Select dan Timeout")
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Transaksi diproses"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Menerima pesan:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Tidak ada respons dalam 1 detik")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("=== Implementasi E-Wallet dengan Concurrency ===")

	wallet := &Wallet{balance: 1000}
	var wg sync.WaitGroup

	demonstrateGoroutines(wallet, &wg)
	demonstrateChannels()
	demonstrateWaitGroup(wallet, &wg)
	demonstrateMutex()
	demonstrateSelectTimeout()
}
