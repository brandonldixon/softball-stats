package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/brandonldixon/softball-stats/cmd"
	"github.com/brandonldixon/softball-stats/web"
)

func main() {

	// Take in flag for command
	cmdPointer := flag.String("cmd", "", "input command")
	flag.Parse()
	if *cmdPointer == "" {
		fmt.Println("Please provide a command.")
		flag.Usage()
		return
	}

	switch *cmdPointer {
	case "create-table":
		_, err := cmd.CreateTableHandler()
		if err != nil {
			log.Fatal(err)
		}
	case "create-player":
		cmd.CreatePlayerHandler()
	case "update-player":
		cmd.UpdatePlayerHandler()
		//fmt.Println("Updating Player")
	case "generate-webpage":
		web.GenerateWebPageHandler()
	default:
		fmt.Println("Unrecognized Command", *cmdPointer)
	}
}
