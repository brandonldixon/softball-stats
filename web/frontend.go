package web

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brandonldixon/softball-stats/cmd"
)

// Define slice of players
type Players []cmd.Player

func GenerateWebPageHandler() {
	// Load Permissions
	config, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create http client for dynamodb api
	dynamoClient := dynamodb.NewFromConfig(config)
	tableName := "Nature"

	generateHtml(dynamoClient, tableName)

}

func generateHtml(client *dynamodb.Client, name string) {

	// Read Table
	p, err := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(name),
	})
	if err != nil {
		log.Fatal(err)
	}

	var players Players
	// Read scan output as slice of players
	for _, item := range p.Items {
		var player cmd.Player
		err = attributevalue.UnmarshalMap(item, &player)
		if err != nil {
			log.Fatal(err)
		}
		players = append(players, player)
	}
	sort.Sort(sort.Reverse(players))

	fmt.Println(players)
	for _, player := range players {
		player.Print()
	}

	// Parse HTML template
	tmpl, err := template.New("table").Parse(`
    <!DOCTYPE html>
    <html lang="en">
    
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Softball Stats</title>
        <!--Links to our stylesheets-->
    </head>
    <style>
        .table-container {
            display: flex;
            justify-content: center;
            overflow-x: auto;
        }
    
        .content-table {
            border-collapse: collapse;
            margin: 10px 0;
            font: 0.9em;
        }
    
        .content-table thead tr {
            background-color: #1eaa41;
            color: #ffffff;
            text-align: center;
            font-weight: bold;
        }
    
        .content-table th,
        .content-table td {
            padding: 5px 5px;
        }
    
        .content-table tbody tr {
            border-bottom: 1px solid #dddddd;
            text-align: center;
        }
    
        .content-table tbody tr:nth-of-type(even) {
            background-color: #f3f3f3;
        }
    
        @media screen and (max-width: 600px) {
    
            /* Apply responsive styles for screens up to 600px wide */
            .table-container {
                overflow-x: auto;
                /* Add horizontal scrollbar if the table overflows */
                width: auto;
                /* Allow the container to expand as needed */
            }
    
            table {
                width: 100%;
                /* Ensure the table fills the container */
            }
        }
    </style>
    
    <body>
        <div class="table-container">
            <table class="content-table" border="2">
                <thead>
                    <tr>
                        <th>No.</th>
                        <th>First</th>
                        <th>Last</th>
                        <th>PAs</th>
                        <th>ABs</th>
                        <th>Runs</th>
                        <th>Hits</th>
                        <th>2Bs</th>
                        <th>3Bs</th>
                        <th>HRs</th>
                        <th>RBIs</th>
                        <th>BBs</th>
                        <th>AVG</th>
                        <th>OBP</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .}}
                    <tr>
                        <td>{{.JerseyNumber}}</td>
                        <td>{{.FirstName}}</td>
                        <td>{{.LastName}}</td>
                        <td>{{.PlateAppearances}}</td>
                        <td>{{.AtBats}}</td>
                        <td>{{.Runs}}</td>
                        <td>{{.Hits}}</td>
                        <td>{{.Doubles}}</td>
                        <td>{{.Triples}}</td>
                        <td>{{.HomeRuns}}</td>
                        <td>{{.RBIs}}</td>
                        <td>{{.Walks}}</td>
                        <td>{{.BattingAverage}}</td>
                        <td>{{.OnBasePercentage}}</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
        </div>
    </body>
    
    </html>
		    `)
	if err != nil {
		// Handle error
		panic(err)
	}

	f, err := os.Create("./webpage.html")
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = tmpl.Execute(f, players)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}

// Length function
func (p Players) Len() int {
	return len(p)
}

// Less function
func (p Players) Less(i, j int) bool {
	return p[i].BattingAverage < p[j].BattingAverage
}

// Swap function
func (p Players) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
