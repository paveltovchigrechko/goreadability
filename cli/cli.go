// Package ari provides functions to calculate the Coleman–Liau index (CLI) for texts readability.
//
// See https://en.wikipedia.org/wiki/Coleman%E2%80%93Liau_index for the details.
package cli

import (
	"goreadability/stats"
	"math"
)

// CalculateCLI accepts a string and returns the Coleman–Liau index (CLI) for it.
// The calculated CLI is rounded to the first decimal point.
func CalculateCLI(s string) float64 {
	characters := float64(stats.CountCharacters(s))
	words := float64(stats.CountWords(s))
	sentences := float64(stats.CountSentences(s))
	cli := 0.0588*(characters/words)*100 - 0.296*(sentences/words)*100 - 15.8
	cli = math.Round(cli*10) / 10
	return cli
}
