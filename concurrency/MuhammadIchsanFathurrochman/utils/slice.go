package utils

import "fmt"

func Slice() {
	// membuat slices
	number := []int{1,2,3,4,5,6,7,8}
	buah := make([]string, 3) // hanya bisa diisi oleh [0] , [1] , [2]

	array := [5]int{10, 2, 4, 12, 2}
	slice1 := array[1:3]
	slice2 := array[:3]

	buah = append(buah, "Semangka")
	buah = append(buah, "Mangga", "Apel")

	fmt.Println(number)
	fmt.Println(buah)
	fmt.Println(slice1)
	fmt.Println(slice2)

}