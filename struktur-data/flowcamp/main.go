package main

import (
	"fmt"
)

// Struct: Digunakan untuk menyimpan data kompleks
type Person struct {
	Name string
	Age  int
}

// Interface: Mendefinisikan kontrak method tanpa implementasi spesifik
type Describer interface {
	Describe() string
}

// Implementasi Interface pada struct
func (p Person) Describe() string {
	return fmt.Sprintf("Nama: %s, Umur: %d", p.Name, p.Age)
}

// Function untuk demonstrasi Array
func arrayExample() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)
}

// Function untuk demonstrasi Slice
func sliceExample() {
	slice := []int{10, 20, 30}
	slice = append(slice, 40)
	fmt.Println("Slice:", slice)
}

// Function untuk demonstrasi Map
func mapExample() {
	personAge := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Println("Map:", personAge)
}

// Function untuk demonstrasi Pointer
func pointerExample() {
	num := 100
	ptr := &num
	fmt.Println("Pointer Value:", *ptr)
}

// Node untuk Linked List
type Node struct {
	Value int
	Next  *Node
}

// Function untuk demonstrasi Linked List
func linkedListExample() {
	head := &Node{Value: 1}
	head.Next = &Node{Value: 2}
	head.Next.Next = &Node{Value: 3}

	fmt.Print("Linked List: ")
	for current := head; current != nil; current = current.Next {
		fmt.Print(current.Value, " -> ")
	}
	fmt.Println("nil")
}

// Queue menggunakan Slice
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	value := q.items[0]
	q.items = q.items[1:]
	return value
}

func queueExample() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println("Queue Dequeue:", q.Dequeue())
}

// Stack menggunakan Slice
type Stack struct {
	items []int
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

func stackExample() {
	s := Stack{}
	s.Push(5)
	s.Push(10)
	fmt.Println("Stack Pop:", s.Pop())
}

// Channel: Digunakan untuk komunikasi antar goroutine
func channelExample() {
	ch := make(chan string)
	go func() {
		ch <- "Hello from Goroutine"
	}()
	fmt.Println("Channel Received:", <-ch)
}

func main() {
	arrayExample()
	sliceExample()
	mapExample()
	pointerExample()
	linkedListExample()
	queueExample()
	stackExample()
	channelExample()

	p := Person{Name: "John", Age: 28}
	fmt.Println(p.Describe())
}
