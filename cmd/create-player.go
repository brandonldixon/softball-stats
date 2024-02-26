package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// A function to add a new player to a team roster
// The function prompts for player First Name, Last Name, and Jersey Number
// This creates a new player struct for stats to be added for the player
func CreatePlayerHandler() (string, error) {
	p := Player{}
	var i, j, k, teamName string
	fmt.Println("Add a New Player to Roster.")
	fmt.Println("Enter Player First Name: ")
	fmt.Scanln(&i)
	fmt.Println("Enter Player Last Name: ")
	fmt.Scanln(&j)
	fmt.Println("Enter Player Jersey Number: ")
	fmt.Scanln(&k)
	fmt.Println("Enter Team Name for Player: ")
	fmt.Scanln(&teamName)
	p.FirstName = i
	p.LastName = j
	p.PlayerName = p.FirstName + p.LastName
	p.JerseyNumber = k
	fmt.Printf("Creating Player\n Name: %s\n Number: %s\n", p.PlayerName, p.JerseyNumber)
	//p.PlayerID = GenerateId()
	//return &p

	// Load Permissions
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	// Create http client for dynamodb api
	dynamoClient := dynamodb.NewFromConfig(config)

	// Create New Player
	createPlayer(teamName, p, dynamoClient)
	if err != nil {
		fmt.Println(err)
	}
	return "", nil
}

func createPlayer(teamName string, p Player, client *dynamodb.Client) error {
	item, err := attributevalue.MarshalMap(p)
	if err != nil {
		return err
	}
	_, err = client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(teamName), Item: item,
	})
	if err != nil {
		return err
	}
	return nil
}
