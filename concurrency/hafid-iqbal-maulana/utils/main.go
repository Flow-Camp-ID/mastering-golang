package main

import (
	"fmt"
	"sync"
)

// Fungsi untuk memeriksa apakah suatu bilangan adalah bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Fungsi worker untuk memproses bilangan dari channel primes
func primeWorker(primes <-chan int, wg *sync.WaitGroup, results chan<- int) {
	defer wg.Done()
	for n := range primes {
		if isPrime(n) {
			results <- n
		}
	}
}

// Fungsi utama untuk menghasilkan bilangan prima
func generatePrimes(limit int) []int {
	primes := make(chan int, limit)
	results := make(chan int, limit)
	var wg sync.WaitGroup

	// Memulai goroutine untuk worker
	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go primeWorker(primes, &wg, results)
	}

	// Mengirim bilangan ke channel primes
	go func() {
		for i := 2; i <= limit; i++ {
			primes <- i
		}
		close(primes)
	}()

	// Menutup channel results setelah semua worker selesai
	go func() {
		wg.Wait()
		close(results)
	}()

	// Mengumpulkan hasil
	var primesList []int
	for prime := range results {
		primesList = append(primesList, prime)
	}

	return primesList
}

func main() {
	limit := 100
	primes := generatePrimes(limit)
	fmt.Println("Bilangan prima hingga", limit, ":")
	fmt.Println(primes)
}
