package utils

import "fmt"

func SliceFunction(){
	// simpan array komoditas
	komoditas := [...]string{"Kopi","Cokelat","Kentang","Belerang","Tape"}
	// membuat slice dari index 1-3 
	slice := komoditas[1:4]
	
	//cetak variable slice
	fmt.Println(slice)
}