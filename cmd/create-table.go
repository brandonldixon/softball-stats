package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func CreateTableHandler() (string, error) {
	// Prompt for Table Name
	var n string
	fmt.Println("Enter Name of Team to Create Table: ")
	fmt.Scanln(&n)

	// Load Permissions
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}

	// Create http client for dynamodb api
	dynamoClient := dynamodb.NewFromConfig(config)

	// Check to see if the table already exists
	exists, err := tableExists(n, dynamoClient)
	if err != nil {
		return "", err
	}
	if !exists {
		log.Printf("Creating table %v...\n", n)
		_, err := createTable(n, dynamoClient)
		if err != nil {
			return "", err
		} else {
			log.Printf("Created Table %v.\n", n)
		}
	} else {
		return "", errors.New("table already exists")
	}
	return "", nil
}

// TableExists determines whether a DynamoDB table exists.
func tableExists(n string, client *dynamodb.Client) (bool, error) {
	exists := true
	_, err := client.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(n)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", n)
			err = nil
		} else {
			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", n, err)
		}
		exists = false
	}
	return exists, err
}

// Create Table function
func createTable(n string, client *dynamodb.Client) (*types.TableDescription, error) {
	// Create Table
	var tableDesc *types.TableDescription
	table, err := client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Player Name"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("Jersey Number"),
				AttributeType: types.ScalarAttributeTypeN,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Player Name"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("Jersey Number"),
				KeyType:       types.KeyTypeRange,
			},
		},
		TableName:   aws.String(n),
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		log.Printf("Couldn't create table %v. Here's why: %v\n", n, err)
	} else {
		waiter := dynamodb.NewTableExistsWaiter(client)
		err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
			TableName: aws.String(n)}, 5*time.Minute)
		if err != nil {
			log.Printf("Wait for table exists failed. Here's why: %v\n", err)
		}
		tableDesc = table.TableDescription
	}
	return tableDesc, err
}
