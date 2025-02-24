package utils

import (
	"fmt"
	"time"
)

// Fungsi menggunakan select untuk menangani timeout
func RunSelectTimeoutExample() {
	pesananChannel := make(chan int)

	go func() {
		time.Sleep(3 * time.Second) // Simulasi dapur lambat
		pesananChannel <- 1
	}()

	select {
	case pesanan := <-pesananChannel:
		fmt.Printf("Pesanan %d diterima oleh dapur!\n", pesanan)
	case <-time.After(2 * time.Second): // Timeout jika lebih dari 2 detik
		fmt.Println("Timeout! Dapur terlalu lama memproses pesanan.")
	}
}
