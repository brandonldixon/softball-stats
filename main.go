package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/brandonldixon/softball-stats/cmd"
)

func main() {

	/*
		brandon := cmd.Player{
			FirstName:    "Brandon",
			LastName:     "Dixon",
			JerseyNumber: 1,
		}
		brandon.Stats.AtBats = 4
		brandon.Stats.Hits = 3
		brandon.Stats.Walks = 2

		brandon.CalculateBattingAverage()
		brandon.CalculateOnBasePercentage()
		brandon.Print()

		brandon.UpdateStats(4, 3, 1, 1, 0, 1, 2)
		brandon.CalculateBattingAverage()
		brandon.CalculateOnBasePercentage()
		brandon.Print()
	*/
	/*

		player2 := cmd.CreatePlayer()
		player2.UpdateStats(4, 4, 2, 1, 0, 1, 5)
		fmt.Println(player2)

	*/

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
		cmd.CreatePlayer()
		fmt.Println("Creating Player")
	case "update-stats":
		//cmd.UpdateStats()
		fmt.Println("Updating Player")
	default:
		fmt.Println("Unrecognized Command", *cmdPointer)
	}

	// Data Calls
	/*

		type Player struct {
			Name           string                 `dynamodbav:"name"`
			BattingAverage int                    `dynamodbav:"battingaverage"`
			Info           map[string]interface{} `dynamodbav:"info"`
		}
	*/

	/*

		// GetKey returns the composite primary key of the movie in a format that can be
		// sent to DynamoDB.
		func (player Player) GetKey() map[string]types.AttributeValue {
			title, err := attributevalue.Marshal(movie.Title)
			if err != nil {
				panic(err)
			}
			year, err := attributevalue.Marshal(movie.Year)
			if err != nil {
				panic(err)
			}
			return map[string]types.AttributeValue{"title": title, "year": year}
		}
	*/
}
