package utils

import (
	"fmt"
	"sync"
	"time"
)

// RoastingKopi menyimpan informasi tentang proses pemanggangan kopi.
type RoastingKopi struct {
    id               int           // ID pesanan kopi
    jenis_kopi       string        // Jenis kopi yang digunakan
    suhu_air         string        // Suhu air yang digunakan untuk menyeduh kopi
    jenis_roasting    string        // Jenis pemanggangan kopi
    waktu_penyeduhan time.Duration  // Waktu yang dibutuhkan untuk menyeduh kopi
}

// CoffeeShop merepresentasikan sebuah kedai kopi dan antrean pemanggangan kopi menggunakan goroutines.
type CoffeeShop struct {
    roastingQueue chan RoastingKopi // Channel untuk antrean proses roasting
    wg           sync.WaitGroup      // WaitGroup untuk menunggu semua goroutine selesai
}

// SequentialCoffeeShop merepresentasikan kedai kopi yang melakukan pemanggangan secara berurutan.
type SequentialCoffeeShop struct {
    roastings []RoastingKopi // Daftar pemanggangan kopi dalam urutan
}

// NewCoffeeShop menginisialisasi dan mengembalikan instansi CoffeeShop dengan ukuran buffer tertentu untuk channel.
func NewCoffeeShop(bufferSize int) *CoffeeShop {
    return &CoffeeShop{
        roastingQueue: make(chan RoastingKopi, bufferSize),
    }
}

// NewSequentialCoffeeShop menginisialisasi dan mengembalikan instansi SequentialCoffeeShop.
func NewSequentialCoffeeShop() *SequentialCoffeeShop {
    return &SequentialCoffeeShop{
        roastings: make([]RoastingKopi, 0),
    }
}

// Roasting memproses pemanggangan kopi secara bersamaan. Menggunakan goroutine dan WaitGroup untuk sinkronisasi.
func (r *CoffeeShop) Roasting(roastingId int) {
    defer r.wg.Done() // Menandakan goroutine telah selesai saat fungsi keluar
    for roasting := range r.roastingQueue {
        fmt.Printf("[Concurrency] Barista %d mulai membuat kopi %s dengan jenis roasting %s untuk pesanan #%d\n",
            roastingId, roasting.jenis_kopi, roasting.jenis_roasting, roasting.id)
        time.Sleep(roasting.waktu_penyeduhan * time.Second) // Simulasi waktu penyeduhan
        fmt.Printf("[Concurrency] Barista %d Selesai membuat kopi %s untuk pesanan #%d\n",
            roastingId, roasting.jenis_kopi, roasting.id)
    }
}

// Roasting memproses pemanggangan kopi secara berurutan tanpa menggunakan goroutine.
func (r *SequentialCoffeeShop) Roasting(roastingId int, roasting RoastingKopi) {
    fmt.Printf("[Sequential] Barista %d mulai membuat kopi %s dengan jenis roasting %s untuk pesanan #%d\n",
        roastingId, roasting.jenis_kopi, roasting.jenis_roasting, roasting.id)
    time.Sleep(roasting.waktu_penyeduhan * time.Second) // Simulasi waktu penyeduhan
    fmt.Printf("[Sequential] Barista %d selesai membuat kopi %s untuk pesanan #%d\n",
        roastingId, roasting.jenis_kopi, roasting.id)
}

// RunConcurrent menjalankan pemrosesan pemanggangan kopi secara bersamaan dengan menggunakan goroutines dan buffered channel.
func RunConcurrent(roastings []RoastingKopi, bufferSize int) time.Duration {
    start := time.Now() // Mencatat waktu mulai proses

    coffeShop := NewCoffeeShop(bufferSize) // Membuat kedai kopi dengan ukuran buffer

    numCooks := 3 // Jumlah barista yang akan bekerja
    coffeShop.wg.Add(numCooks) // Menambahkan jumlah goroutine yang harus ditunggu

    // Menjalankan goroutine untuk setiap barista
    for i := 1; i <= numCooks; i++ {
        go coffeShop.Roasting(i)
    }

    // Mengirim setiap pemanggangan ke dalam antrean
    for _, roasting := range roastings {
        coffeShop.roastingQueue <- roasting
    }

    close(coffeShop.roastingQueue) // Menutup channel setelah semua pemanggangan dikirim
    coffeShop.wg.Wait() // Menunggu semua goroutine selesai

    return time.Since(start) // Mengembalikan durasi waktu pemrosesan
}

// RunSequential menjalankan pemrograman pemanggangan kopi secara berurutan.
func RunSequential(roastings []RoastingKopi) time.Duration {
    start := time.Now() // Mencatat waktu mulai proses

    coffeShop := NewSequentialCoffeeShop() // Membuat kedai kopi sequential

    // Memproses setiap pemanggangan kopi secara berurutan
    for _, roasting := range roastings {
        coffeShop.roastings = append(coffeShop.roastings, roasting) // Menambahkan pemanggangan ke antrean
        coffeShop.Roasting(1, roasting) // Memanggil fungsi Roasting untuk memproses pemanggangan
    }

    return time.Since(start) // Mengembalikan durasi waktu pemrosesan
}

// PesanKopi menghasilkan beberapa jenis kopi, menjalankan proses secara sequential dan concurrent, serta membandingkan waktu.
func PesanKopi() {
    // Daftar pesanan kopi yang akan diproses
    roastings := []RoastingKopi{
        {1, "Arabica", "90-96°C", "(Light, Medium, Dark)", 4},
        {2, "Robusta", "92-98°C", "(Medium, Dark)", 5},
        {3, "Liberica", "92-96°C", "(Medium, Dark)", 6},
        {4, "Excelsa", "90-95°C", "(Light, Medium", 5},
        {5, "Espresso Blend", "88-94°C", "(Medium-Dark, Dark)", 1},
        {6, "Turkish Coffee", "85-95°C", "(Dark)", 3},
    }

    // Menjalankan proses pemanggangan secara sequential
    fmt.Println("Sequential Process")
    seqTime := RunSequential(roastings)

    // Menjalankan proses pemanggangan secara concurrent dengan buffered channel
    fmt.Println("[Buffered] Concurrency Process")
    concTimebuffered := RunConcurrent(roastings, 10)

    // Menjalankan proses pemanggangan secara concurrent dengan unbuffered channel
    fmt.Println("[Unbuffered] Concurrency Process")
    concTimeUnbufferd := RunConcurrent(roastings, 0)

    // Menampilkan perbandingan waktu antara pemrosesan sequential dan concurrent
    fmt.Println("Perbandingan")
    fmt.Printf("Waktu sequential: %v\n", seqTime)
    fmt.Printf("Waktu concurrency[Buffered]: %v\n", concTimebuffered)
    fmt.Printf("Waktu concurrency[Unbuffered]: %v\n", concTimeUnbufferd)
    fmt.Printf("Selisih waktu sequential & concurrency[Buffered]: %v\n", seqTime-concTimebuffered)
    fmt.Printf("Selisih waktu sequential & concurrency[Unbuffered]: %v\n", seqTime-concTimeUnbufferd)
}