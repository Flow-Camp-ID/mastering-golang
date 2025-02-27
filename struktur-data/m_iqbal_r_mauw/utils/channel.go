package utils

import "fmt"

func Channel() {
	fmt.Println("")
	fmt.Println("Channel")
	fmt.Println("=======")

	// Membut Channel
	messages := make(chan string) // inisialisasi channel messages
	go func() {                   // goroutine
		messages <- "Hello" // kirim data ke channel messages
	}()
	msg := <-messages // terima data dari channel messages
	fmt.Println(msg)  // print msg

	// Channel dengan Buffer
	bufferedMessages := make(chan string, 2) // inisialisasi channel bufferedMessages dengan buffer 2
	go func() {                              // goroutine
		bufferedMessages <- "Hello" // kirim data ke channel bufferedMessages
		bufferedMessages <- "World" // kirim data ke channel bufferedMessages
	}()
	fmt.Println(<-bufferedMessages) // terima data dari channel bufferedMessages
	fmt.Println(<-bufferedMessages) // terima data dari channel bufferedMessages

	// Channel dengan Range
	rangeMessages := make(chan string, 2) // inisialisasi channel rangeMessages dengan buffer 2
	rangeMessages <- "Hello"              // kirim data ke channel rangeMessages
	rangeMessages <- "World"              // kirim data ke channel rangeMessages
	close(rangeMessages)                  // tutup channel rangeMessages
	for message := range rangeMessages {  // loop rangeMessages
		fmt.Println(message) // print message
	}
}
