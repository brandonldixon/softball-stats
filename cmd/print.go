package cmd

import "fmt"

// A function to call that prints a player struct
func (p Player) Print() {
	fmt.Printf("\n %+v \n", p)
}
