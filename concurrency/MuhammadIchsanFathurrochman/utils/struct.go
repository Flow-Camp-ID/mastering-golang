package utils

import "fmt" 

type DataDiri struct {
	firstName string
	lastName string
	email string
}

type SocialMedia struct {
	instagram string
	youtube string
}

type Peoples struct {
	data DataDiri
	Sosmed SocialMedia
	role string
}

func StructFunc() {

	var people []interface {}

	people = append(people, DataDiri{
		firstName: "Taufik",
		lastName: "Mulyawan",
		email: "taufikmulyawan",
	})

	people = append(people, SocialMedia{
		instagram: "taufik.dev",
		youtube: "Flow Camp ID",
	})

	fmt.Println(people)

	nama1 := Peoples{
		data: DataDiri{
			firstName: "Taufik",
			lastName: "Mulyawan",
			email: "taufikmulyawan",
		},
		Sosmed: SocialMedia{
			instagram: "taufik.dev",
			youtube: "Flow Camp ID",
		},
		role: "Backend",
	}

	fmt.Println(nama1)


	// people := []DataDiri {
	// 	{
	// 	firstName: "Taufik",
	// 	lastName: "Mulyawan",
	// 	email: "taufikmulyawan@gmail.com",
	// 	},
	// 	{
	// 		firstName: "Opung",
	// 		lastName: "Opung",
	// 		email: "opung@gmail.com",
	// 		},
	// }

	// fmt.Println(people)
}