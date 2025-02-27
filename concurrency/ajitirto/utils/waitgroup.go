package utils

import (
	"time"
	"sync"
	"fmt"
)

func LombaLariWaitgroup(nama string, kecepatan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(nama, "sedang berlari...") // Pesan sedang berlari

	time.Sleep(time.Duration(1000/kecepatan) * time.Millisecond) // Simulasikan waktu lari

	fmt.Println(nama, "selesai!") // Pesan selesai
}