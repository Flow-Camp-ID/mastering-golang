package utils

import (
	"fmt"
)

// Unbuffered Channel
func RunUnbufferedChannel(pesananChannel chan int, totalPesanan int) {
	for i := 1; i <= totalPesanan; i++ {
		fmt.Printf("Mengirim pesanan %d ke dapur...\n", i)
		pesananChannel <- i
	}
	close(pesananChannel)
}

// Buffered Channel
func RunBufferedChannel(pesananChannel chan int, totalPesanan int) {
	for i := 1; i <= totalPesanan; i++ {
		fmt.Printf("Mengirim pesanan %d ke dapur (Buffered)...\n", i)
		pesananChannel <- i
	}
	close(pesananChannel)
}
