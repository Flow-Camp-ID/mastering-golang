package utils

import (
	"fmt"
	"sync"
	"time"
)

// mendifinisikan data order kendaraan terdiri dari ID, Pat Nomor, Nama Kendaraa, dan Waktu Pengerjaan
type Order struct {
	Id       int
	CarName  string
	PlatNo   string
	WashTime time.Duration
}

type CarWash struct {

	// inisialisasi channel
	OrderQueue chan Order

	// sync wait group digunakan utk menunggu waktu sampai func beres proses
	Wg sync.WaitGroup
}

func NewCarWash() *CarWash {

	//set pointer untuk antrian kendaraan
	return &CarWash{
		OrderQueue: make(chan Order, 10),
	}

}

// concurrency process
func (r *CarWash) Washing(WashID int) {

	//menampilkan proses pengerjaan operator
	defer r.Wg.Done()
	for Order := range r.OrderQueue {
		fmt.Printf("Opeartor %d mulai mencuci %s dengan nopol %s untuk order #%d\n", WashID, Order.CarName, Order.PlatNo, Order.Id)
		time.Sleep(Order.WashTime * time.Second)
		fmt.Printf("Operator %d selesai mencuci %s dengan nopol %s untuk order #%d\n", WashID, Order.CarName, Order.PlatNo, Order.Id)
	}
}

// run concurrent
func RunConcurrent(Orders []Order) time.Duration {

	// Jalanin go routine pengerjaan car wash

	Start := time.Now()
	CarWash := NewCarWash()

	NumWashing := 5
	CarWash.Wg.Add(NumWashing)

	// Masukin data operator
	for i := 1; i <= NumWashing; i++ {
		go CarWash.Washing(i)
	}

	// Masukin data orderan kendaraan
	for _, Order := range Orders {
		CarWash.OrderQueue <- Order
	}

	// ketika sudah dijalankan semua diclose supaya mencegah rice condition
	close(CarWash.OrderQueue)
	CarWash.Wg.Wait()

	return time.Since(Start)
}

func ConcurrencyCase(Orders []Order) {

	// mulai menjalankan order dari antrian kendaraan yang datang
	fmt.Println("Concurency Process untuk car wash")
	ConcTime := RunConcurrent(Orders)
	fmt.Printf("Wantu concurrency yang dibutuhkan : %v\n", ConcTime)
}
