package utils

import (
	"fmt"
	"sync"
)

// Menjalankan beberapa tugas dengan WaitGroup
func RunWaitGroupExample(totalPesanan int) {
	var wg sync.WaitGroup

	for i := 1; i <= totalPesanan; i++ {
		wg.Add(1)
		go func(nomor int) {
			defer wg.Done()
			fmt.Printf("Pesanan %d sedang diproses...\n", nomor)
		}(i)
	}

	wg.Wait()
	fmt.Println("Semua pesanan telah selesai diproses!")
}
