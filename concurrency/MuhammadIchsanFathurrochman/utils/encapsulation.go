package utils

import "fmt"

type Identitas struct {
	Nama string // public 
	umur int // private
	alamat string // private
}

func NewIdentity(nama string, umur int, alamat string) Identitas {
	return Identitas{
		Nama: nama,
		umur: umur,
		alamat: alamat,
	}
}

// getter
func (p Identitas) GetAge() int {
	return p.umur
}

func (p *Identitas) SetAge(umur int) error {
	if umur < 0 || umur > 150 {
		return fmt.Errorf("umur tidak valid")
	}
	p.umur = umur
	return nil
}

func Encapsulation() {
	identitas := NewIdentity("Taufik", 23, "Jakarta")

	fmt.Println("Nama: " + identitas.Nama)
	// fmt.Println(identitas.GetAge())

	err := identitas.SetAge(151)

	if err == nil {
		fmt.Println(identitas.GetAge())
	}else {
		fmt.Println("Umur tidak valid")
	}

	// fmt.Println(identitas.GetAge())

}