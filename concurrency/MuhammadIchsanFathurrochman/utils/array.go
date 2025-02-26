package utils

import "fmt"

func ArrayFunction() {
	// case 1 - Array Fixed
	var numbers [5]int = [5]int{1,2}
	fmt.Println(numbers)

	// case 2 - implisit size
	name := [...]string{"Taufik", "Test"}
	fmt.Println(name)

	// case 3 - interface array
	testing := [...]interface{}{"Taufik", 20, true}
	fmt.Println(testing[1])

	// case 4 - specific init array
	makanan := [5]string{0: "Mie", 3: "Ayam Goreng"}
	fmt.Println(makanan[3])

	// case 5 - multidimensi array
	matrix := [2][3]int{
		{1, 5, 2}, // index ke 0, [0][0] -> 1, [0][1] -> 5, [0][2] -> 2
		{9, 2, 5}, // index ke 1, [1][0] -> 9, [1][1] -> 2, [1][2] -> 5
	}

	fmt.Println(matrix[1][0])

	matrix[1][0] = 100

	fmt.Println(matrix[1][0])

}