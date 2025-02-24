package main

import (
	"fmt"
	"restaurant/utils"
	"sync"
)

func main() {
	fmt.Println("=== Memulai Pemrosesan Pesanan Restoran ===")

	// 1. Menjalankan Goroutine
	fmt.Println("\n[1] Menjalankan pesanan dengan Goroutine:")
	utils.RunGoroutineExample()

	// 2. Menggunakan Unbuffered Channel
	fmt.Println("\n[2] Menggunakan Unbuffered Channel:")
	unbufferedPesananChannel := make(chan int)
	go utils.RunUnbufferedChannel(unbufferedPesananChannel, 3)

	for pesanan := range unbufferedPesananChannel {
		fmt.Printf("Koki mulai memasak pesanan %d...\n", pesanan)
	}

	// 3. Menggunakan Buffered Channel
	fmt.Println("\n[3] Menggunakan Buffered Channel:")
	bufferedPesananChannel := make(chan int, 3)
	go utils.RunBufferedChannel(bufferedPesananChannel, 5)

	for pesanan := range bufferedPesananChannel {
		fmt.Printf("Koki sedang memasak pesanan %d (Buffered)...\n", pesanan)
	}

	// 4. Menggunakan Mutex
	fmt.Println("\n[4] Menggunakan Mutex untuk mengontrol stok makanan:")
	stok := utils.StokMakanan{Stok: 5}

	var wg sync.WaitGroup
	for i := 1; i <= 7; i++ {
		wg.Add(1)
		go func(nomor int) {
			defer wg.Done()
			stok.KurangiStok(1, nomor)
		}(i)
	}
	wg.Wait()

	// 5. Menggunakan Select dengan Timeout
	fmt.Println("\n[5] Menggunakan Select untuk menangani Timeout:")
	utils.RunSelectTimeoutExample()

	fmt.Println("\n=== Semua pesanan selesai! ===")
}
