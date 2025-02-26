package utils
import "fmt"

func Channel() {

	// case 1 channel basic
	ch := make(chan int)

	go func() {
		ch <- 42 // user a beli barang CH
		ch <- 30 // user b beli barang CH
	}()
	
	value := <-ch
	fmt.Println(value)

	// case 2 buffered channel
	bufferedChannel := make(chan string, 3)
	bufferedChannel <- "Halo"
	bufferedChannel <- "Taufik"
	bufferedChannel <- "Oke"

	fmt.Println("First message: ", <-bufferedChannel)
	fmt.Println("Second message: ", <-bufferedChannel)
	fmt.Println("Third message: ", <-bufferedChannel)

}