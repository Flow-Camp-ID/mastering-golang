package utils

import "fmt"

func MapImplement() {
	// map dengan tipe data integer sebagai index dan value 
	nilaiSiswa := map[int]int{
		1: 9,
		2: 8,
		3: 7,
		4: 8,
		5: 10,
	}

	fmt.Println("Nilai siswa berdasarkan index")
	for index, value := range nilaiSiswa {
		fmt.Printf("%d->%d\n", index, value)
	}
}