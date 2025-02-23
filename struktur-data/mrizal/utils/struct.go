package utils

import "fmt"

type Player struct {
	name   string
	role   string
	height int
}

type Team struct {
	name        string
	description string
	player      []Player
}

func StructImplement() {

	// single struct
	newPlayer := Player{
		name:   "Ronaldo",
		role:   "Striker",
		height: 185,
	}
	
	fmt.Println(newPlayer)
	// array of struct
	dataOfTeam := []Player{
		{
			name: "Messi",
			role: "Penyerang",
			height: 170,
		},
		{
			name: "Pepe",
			role: "Defender",
			height: 190,
		},
		 {
			name: "Van",
			role: "Keeper",
			height: 190,
		},
	}
   
	// struct of struct 
	newTeam := Team {
		name: "Persija",
		description: "new team with high budget",
		player: dataOfTeam,
	}
	fmt.Printf("\n\nTim Baru\n")
	fmt.Println(newTeam)
}