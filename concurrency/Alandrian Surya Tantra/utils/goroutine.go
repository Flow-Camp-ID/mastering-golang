package utils

import (
	"fmt"
	"time"
)

// Fungsi untuk menjalankan pesanan secara konkuren
func prosesPesanan(nomorPesanan int) {
	fmt.Printf("Pesanan %d sedang diproses...\n", nomorPesanan)
	time.Sleep(time.Second * 2)
	fmt.Printf("Pesanan %d selesai!\n", nomorPesanan)
}

// Menjalankan beberapa Goroutine
func RunGoroutineExample() {
	go prosesPesanan(1)
	go prosesPesanan(2)
	go prosesPesanan(3)
	time.Sleep(time.Second * 3) // Menunggu agar Goroutine selesai sebelum lanjut
}
