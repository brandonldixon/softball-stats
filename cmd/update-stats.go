package cmd

import "math"

// A method that updates the stats of a player
// The function takes in several arguments, and then updates the struct by adding the new stats to the existing stats
func (p *Player) UpdateStats(newAtBats, newHits, newSingles, newDoubles, newTriples, newHomeRuns, newRbis int) {
	(*p).Stats.AtBats += newAtBats
	(*p).Stats.Hits += newHits
	(*p).Stats.Singles += newSingles
	(*p).Stats.Doubles += newDoubles
	(*p).Stats.Triples += newTriples
	(*p).Stats.HomeRuns += newHomeRuns
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
