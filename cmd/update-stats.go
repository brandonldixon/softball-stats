package cmd

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Grab a player from the database
// Handler function to handle the update-stats command
// Function that grabs a player from the database and calls the update stats functions then writes it back to the database
// This is done as a putitem without the conditional

// UpdatePlayer Handler function
func UpdatePlayerHandler() (string, error) {
	var i, j, teamName string
	fmt.Println("Update a Player already on the Roster.")
	fmt.Println("Enter Player First Name: ")
	fmt.Scanln(&i)
	fmt.Println("Enter Player Last Name: ")
	fmt.Scanln(&j)
	fmt.Println("What Team Name does the player play for: ")
	fmt.Scanln(&teamName)
	playerName := i + j
	//fmt.Printf("Creating Player\n Name: %s\n Number: %s\n", p.PlayerName, p.JerseyNumber)
	//p.PlayerID = GenerateId()
	//return &p

	// Load Permissions
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	// Create http client for dynamodb api
	dynamoClient := dynamodb.NewFromConfig(config)

	// Update Player Stats
	err = updatePlayer(teamName, playerName, dynamoClient)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return "", nil
}

// The function that updates the player
func updatePlayer(t, k string, client *dynamodb.Client) error {
	key, err := attributevalue.Marshal(k)
	if err != nil {
		log.Fatal(err)
	}
	item := map[string]types.AttributeValue{"PlayerName": key}
	p, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(t),
		Key:       item,
	})
	if err != nil {
		log.Fatal(err)
	}
	var player Player
	err = attributevalue.UnmarshalMap(p.Item, &player)
	if err != nil {
		fmt.Println(err)
	}
	player.Print()
	player.UpdateStats()
	player.CalculateBattingAverage()
	player.CalculateOnBasePercentage()
	player.Print()

	item, err = attributevalue.MarshalMap(player)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(t),
		Item:      item,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// A method that updates the stats of a player
// The function takes in several arguments, and then updates the struct by adding the new stats to the existing stats
func (p *Player) UpdateStats() {
	var newAtBats, newHits, newWalks, newSingles, newDoubles, newTriples, newHomeRuns, newRbis int

	// Prompt for values
	fmt.Println("Enter the number of At Bats the player had:")
	fmt.Scanln(&newAtBats)
	fmt.Println("Enter the number of Hits the player had:")
	fmt.Scanln(&newHits)
	// Input validation for hits
	if newHits > newAtBats {
		log.Fatal("Cannot have more hits than at bats.")
	}
	// Walks
	fmt.Println("How many walks?")
	fmt.Scanln(&newWalks)
	// Singles
	fmt.Println("How many singles?")
	fmt.Scanln(&newSingles)
	// Input validation for singles
	if newSingles > newHits {
		log.Fatal("Cannot have more singles than hits.")
	}
	// Doubles
	fmt.Println("How many doubles?")
	fmt.Scanln(&newDoubles)
	// Input validation for doubles
	if newDoubles+newSingles > newHits {
		log.Fatal("Cannot have more singles and doubles than hits.")
	}
	// Triples
	fmt.Println("How many triples?")
	fmt.Scanln(&newTriples)
	// Input validation
	if newTriples+newDoubles+newSingles > newHits {
		log.Fatal("Cannot have more triples doubles and singles than hits")
	}
	// Home Runs
	fmt.Println("How many Home Runs?")
	fmt.Scanln(&newHomeRuns)
	// Input validation
	if newHomeRuns+newTriples+newDoubles+newSingles != newHits {
		log.Fatal("Cannot have more home runs, triples, doubles and singles than hits")
	}
	// RBIs
	fmt.Println("How many RBIs?")
	fmt.Scanln(&newRbis)

	(*p).Stats.AtBats += newAtBats
	(*p).Stats.Hits += newHits
	(*p).Stats.Walks += newWalks
	(*p).Stats.Singles += newSingles
	(*p).Stats.Doubles += newDoubles
	(*p).Stats.Triples += newTriples
	(*p).Stats.HomeRuns += newHomeRuns
	(*p).Stats.RBIs += newRbis

}

// A method that calculates the batting average.
// The method cals a function that rounds the float to 3 decimal places.
// The batting average will be calculated without walks, as walks do not count as at bats.
func (p *Player) CalculateBattingAverage() {
	(*p).Stats.BattingAverage = roundFloat(float64((*p).Stats.Hits)/float64((*p).Stats.AtBats), 3)
}

// A method that calculates the on base percentage.
// This method calls the roundFloat function also to round the on base percentage to 3 decimal places.
// The on base percentge includes walks, by adding the walks stat to the numerator and demoninator of the calculation equation
func (p *Player) CalculateOnBasePercentage() {
	(*p).Stats.OnBasePercentage = roundFloat(float64((*p).Stats.Hits+(*p).Stats.Walks)/float64((*p).Stats.AtBats+(*p).Stats.Walks), 3)
}

// A function that exists to round a float64 to 3 decimal places
// This may also be able to be done with a fmt.Print(%.3f) to print the float with 3 decimal places
func roundFloat(value float64, precision uint8) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(value*ratio) / ratio
}
