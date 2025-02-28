package utils

import (
	"fmt"
	"sync"
)

// Struktur yang menyimpan stok makanan dengan Mutex
type StokMakanan struct {
	mu   sync.Mutex
	Stok int
}

// Fungsi untuk mengurangi stok makanan dengan Mutex
func (s *StokMakanan) KurangiStok(jumlah int, nomorPesanan int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.Stok >= jumlah {
		fmt.Printf("Pesanan %d: Stok cukup, mengurangi %d...\n", nomorPesanan, jumlah)
		s.Stok -= jumlah
	} else {
		fmt.Printf("Pesanan %d: Stok tidak cukup!\n", nomorPesanan)
	}
}
