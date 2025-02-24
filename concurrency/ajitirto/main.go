package main

import (
	"ajitirto/utils"
	"fmt"
	"sync"
	"time"
)

func implementGoroutine(){
	// ########## BEGIN: Goroutine ####################
	
	kecepatanUser1 := 5 // membuat variable sekalin validasi
	kecepatanUser2 := 10
	kecepatanUser3 := -1

	fmt.Println("lomba lari di mulai");fmt.Print("\n")

	// Penerapan goroutine
	go utils.LombaLari(kecepatanUser1) 
	go utils.LombaLari(kecepatanUser2)
	go utils.LombaLari(kecepatanUser3)

	time.Sleep(1 * time.Second)

	fmt.Print("\n")
	fmt.Println("lomba lari selesai")
	// ########## END: Goroutine ####################
}

func implementChannel(){

	// ########### BEGIN: Channel ####################

	// hasilChan := make(chan string) // unbuffered 
	hasilChan := make(chan string, 2) // buffered : kapasitas buffer

	go utils.LombaLariChannel("Pelari 1", 5, hasilChan)
	go utils.LombaLariChannel("Pelari 2", 10, hasilChan)

	fmt.Println(<-hasilChan) // Terima dan cetak hasil
	fmt.Println(<-hasilChan)
	
	// ########### END: Channel ####################
}

func implementWaitGroup() {
	// ########### BEGIN: waitgroup  ####################

	var wg sync.WaitGroup // deklarasi waitgroup

	pelari := []struct { // deklarasi struct terdiri dari 4 data pelari
			nama      string
			kecepatan int
	}{
			{"Pelari 1", 5},
			{"Pelari 2", 10},
			{"Pelari 3", 3},
			{"Pelari 4", 7},
	}

	fmt.Println("Siap!")

	wg.Add(len(pelari)) // manambah counter

	for _, p := range pelari { // loping dari 4 data pelari
			go utils.LombaLariWaitgroup(p.nama, p.kecepatan, &wg)
	}

	wg.Wait() // menunggu goroutine sampai selasai

	fmt.Println("Semua pelari selesai!")

	// ########### END: waitgroup  ####################
}

func main() {

	implementGoroutine()
	// implementChannel()
	// implementWaitGroup()

}
