package cmd

import "fmt"

// A function to add a new player to a team roster
// The function prompts for player First Name, Last Name, and Jersey Number
// This creates a new player struct for stats to be added for the player
func CreatePlayer() *Player {
	p := Player{}
	var i, j string
	var k int
	fmt.Println("Add a New Player to Roster.")
	fmt.Println("Enter Player First Name: ")
	fmt.Scanln(&i)
	fmt.Println("Enter Player Last Name: ")
	fmt.Scanln(&j)
	fmt.Println("Enter Player Jersey Number: ")
	fmt.Scanln(&k)
	p.FirstName = i
	p.LastName = j
	p.JerseyNumber = k
	//p.PlayerID = GenerateId()
	return &p
}
