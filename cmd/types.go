package cmd

// The struct that represents a team.
// This struct will have basic identifying information for a Team like Name, Coach, State, etc
// A team will be made up of players
type Team struct {
	TeamID    string
	TeamName  string
	CoachName string
	State     string
}

// The struct that represents a player.
// This struct will have basic identifying information like Name and Jersey Number
// But it will also contain the batting stats of the player as a sub struct
type Player struct {
	//PlayerID     string
	FirstName    string `dynamodbav:"FirstName"`
	LastName     string `dynamodbav:"LastName"`
	PlayerName   string `dynamodbav:"PlayerName"`
	JerseyNumber string `dynamodbav:"JerseyNumber"`
	//TeamMembership string // This links the player to a team (Player Item to the Team Table)
	Stats `dynamodbav:"Stats"`
}

// This is the struct that represents the batting stats of a player
// This struct will be contained in the player struct
// This is broken out from the player struct because Players may have multiple stat lines from multiple teams
// This may need to be changed to a map to make DynamoDB operations easier since it will be a map in the table
type Stats struct {
	PlateAppearances int
	AtBats           int
	Hits             int
	Walks            int
	Singles          int
	Doubles          int
	Triples          int
	HomeRuns         int
	RBIs             int
	Runs             int
	BattingAverage   float64
	OnBasePercentage float64
}
