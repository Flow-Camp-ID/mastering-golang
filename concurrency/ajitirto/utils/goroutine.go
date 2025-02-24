package utils

import (
	"fmt"
)

//goroutin
func LombaLari(kecepatan int)  {
	// validasi supaya user tidak input angka negatif atau 0
	if kecepatan <= 0 {
		fmt.Println("Kecepatan tidak boleh 0 atau angka negatif, nilai input = ", kecepatan)
		return
	}

	// lakukan lari
	fmt.Println("Lari dengan kecepetan", kecepatan, "m/s")
}