package utils

import "fmt"

func Pointer() {
	// case 1 - assign pointer ke variabel PTR
	var x int = 50
	var ptr *int = &x
	// var ptr2 *int = &x

	
	// fmt.Println(ptr, " Ini memory address nya")
	// fmt.Println(*ptr, " Ini value dari variabel ptr")

	// edit value dari si PTR
	*ptr = 200

	// case 2 - menggunakan pointer di struct
	dataDiri := &DataDiri{
		firstName: "Taufik",
		lastName: "Mulyawan",
		email: "Taufikmulyawan@gmail.com",
	}

	dataDiri.email = "Testing@gmail.com"

	// fmt.Println(*dataDiri, " Ini value dari variabel dataDiri")

	// case 3 - menggunakan pointer di array
	arr := [3]int{2,3,1}
	arrPtr := &arr
	arrPtr[2] = 100
	fmt.Println((*arrPtr)[1], " Ini value dari variabel dataDiri")

	// case 4 - pointer to pointer
	// var ptrToPtr **int = &ptr
	// fmt.Println(ptr, " Ini memorynya")
	// fmt.Println(*ptrToPtr, " Ini memorynya")

	// fmt.Println(*ptr, " Ini valuenya")
	// fmt.Println(**ptrToPtr, " Ini valuenya")

	// case 5 - pointer to pointer to pointer
	var testing = 50
	var ptrTest *int = &testing  // pointer
	var ptrTest2 **int = &ptrTest // pointer to pointer
	var ptrToPtrToPtr ***int = &ptrTest2 // pointer to pointer to pointer

	fmt.Println(ptrTest, " Ini memorynya")
	fmt.Println(*ptrTest2, " Ini memorynya")
	fmt.Println(**ptrToPtrToPtr, " Ini memorynya")

	fmt.Println(*ptrTest, " Ini valuenya")
	fmt.Println(**ptrTest2, " Ini valuenya")
	fmt.Println(***ptrToPtrToPtr, " Ini valuenya")
}
