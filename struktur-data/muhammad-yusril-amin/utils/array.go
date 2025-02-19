package utils

import "fmt"

func ArrayFunction(){
	// case 1 - array fixed
	var nomor [5]int = [5]int{34,100,12,33,12}
	// cetak variabel
	fmt.Println(nomor)
	
	// case 2 - implisit size (gak mendefiniskan limit array)
	komoditas := [...]string{"Kopi","Cokelat","Kentang","Mangga"}
	// cetak variabel
	fmt.Println(komoditas)
}
