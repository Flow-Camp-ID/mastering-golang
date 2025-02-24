package main

import (
	"fmt"
	"sync"
	"time"
)

type TicketOrder struct {
	id          int
	event       string
	processTime time.Duration
}

type Locker struct {
	orderQueue chan TicketOrder
	wg         sync.WaitGroup
}

type SequentialLocker struct {
	orders []TicketOrder
}

func NewLocker() *Locker {
	return &Locker{
		orderQueue: make(chan TicketOrder, 10),
	}
}

func NewSequentialLocker() *SequentialLocker {
	return &SequentialLocker{
		orders: make([]TicketOrder, 0),
	}
}

// case concurrency
func (l *Locker) Process(orderID int) {
	defer l.wg.Done()
	for order := range l.orderQueue {
		fmt.Printf("[Concurrency] Loker %d mulai memproses tiket %s untuk pesanan #%d\n",
			orderID, order.event, order.id)
		time.Sleep(order.processTime * time.Second) // sebuah jeda
		fmt.Printf("[Concurrency] Loker %d selesai memproses tiket %s untuk pesanan #%d\n",
			orderID, order.event, order.id)
	}
}

// case sequential
func (l *SequentialLocker) Process(orderID int, order TicketOrder) {
	fmt.Printf("[Sequential] Loker %d mulai memproses tiket %s untuk pesanan #%d\n",
		orderID, order.event, order.id)
	time.Sleep(order.processTime * time.Second) // sebuah jeda
	fmt.Printf("[Sequential] Loker %d selesai memproses tiket %s untuk pesanan #%d\n",
		orderID, order.event, order.id)
}

// case concurrency
func RunConcurrent(orders []TicketOrder) time.Duration {
	start := time.Now()

	locker := NewLocker()

	numLockers := 3
	locker.wg.Add(numLockers)

	for i := 1; i <= numLockers; i++ {
		go locker.Process(i)
	}

	for _, order := range orders {
		locker.orderQueue <- order
	}

	close(locker.orderQueue)
	locker.wg.Wait()

	return time.Since(start)
}

// case sequential
func RunSequential(orders []TicketOrder) time.Duration {
	start := time.Now()

	locker := NewSequentialLocker()

	for _, order := range orders {
		locker.orders = append(locker.orders, order)
		locker.Process(1, order)
	}

	return time.Since(start)
}

func ConcurrencyCase() {
	orders := []TicketOrder{
		{1, "Concert A", 2},
		{2, "Concert B", 3},
		{3, "Concert C", 1},
		{4, "Concert D", 1},
		{5, "Concert E", 3},
		{6, "Concert F", 3},
	}

	// runnning runtime
	fmt.Println("Sequential Process")
	seqTime := RunSequential(orders)

	fmt.Println("Concurrency Process")
	concTime := RunConcurrent(orders)

	fmt.Println("Perbandingan")
	fmt.Printf("Waktu sequential: %v\n", seqTime)
	fmt.Printf("Waktu concurrency: %v\n", concTime)
	fmt.Printf("Selisih waktu: %v\n", seqTime-concTime)

}

func main() {
	ConcurrencyCase()
}
