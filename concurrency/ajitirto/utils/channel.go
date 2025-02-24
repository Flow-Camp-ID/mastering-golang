package utils

import (
	"fmt"
	"time"
)

func LombaLariChannel(nama string, kecepatan int, hasilChan chan string) {
	time.Sleep(time.Duration(1000/kecepatan) * time.Millisecond) // Simulasikan waktu lari
	hasilChan <- fmt.Sprintf("%s selesai!", nama)
}