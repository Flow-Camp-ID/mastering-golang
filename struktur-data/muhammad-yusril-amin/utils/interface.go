package utils

import "fmt"

func InterfaceFunction(){
	// Menyimpan data kendaraan dengan tipe data beragam
	kendaraan := [...]interface{}{"Honda brio",25,true}

	// cetak variabel
	fmt.Println(kendaraan)
}