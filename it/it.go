// Package `it` provides functions and types to calculate the readability for texts in Italian language.
// 1. Gulpease index (https://it.wikipedia.org/wiki/Indice_Gulpease)
package it

import (
	"errors"
	"goreadability/stats"
	"math"
)

// CalcGulpease accepts a non-empty string and returns the Dale–Chall readability (DCR) formula for it. The string must contain at least one word (a number is considered a word, for example `18.` is valid string) and at least one sentence.
// The calculated result is rounded to the nearest whole number.
func CalcGulpease(s string) (uint, error) {
	if len(s) == 0 {
		return 0, errors.New("Empty string.")
	}

	words := float64(stats.CountWords(s))
	if words == 0 {
		return 0, errors.New("No words were parsed. Cannot calculate Gulpease readability index.")
	}

	characters := float64(stats.CountCharacters(s))
	sentences := float64(stats.CountSentences(s))

	raw_index_gulpease := 89 + ((300*sentences - 10*characters) / words)
	gulpease_index := uint(math.Round(raw_index_gulpease))
	return gulpease_index, nil
}
