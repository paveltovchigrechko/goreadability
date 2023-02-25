// Package ari provides functions to calculate the Coleman–Liau index (CLI) for texts readability.
//
// See https://en.wikipedia.org/wiki/Coleman%E2%80%93Liau_index for the details.
package cli

import (
	"errors"
	"goreadability/stats"
	"math"
)

// CalculateCLI accepts a string and returns the Coleman–Liau index (CLI) for it.
// The calculated CLI is rounded to the first decimal point.
func CalculateCLI(s string) (float64, error) {
	if len(s) == 0 {
		return 0, errors.New("Empty string.")
	}

	characters := float64(stats.CountCharacters(s))
	words := float64(stats.CountWords(s))
	sentences := float64(stats.CountSentences(s))

	if words == 0 {
		return 0, errors.New("No words were parsed. Cannot calculate CLI.")
	}

	cli := 5.88*(characters/words) - 29.6*(sentences/words) - 15.8
	cli = math.Round(cli*10) / 10
	return cli, nil
}
