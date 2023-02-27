// Package ari provides functions and types to calculate the automated readability index (ARI) for texts.
//
// See https://en.wikipedia.org/wiki/Automated_readability_index for the details.
package ari

import (
	"errors"
	"goreadability/stats"
	"math"
)

// ====== Types ======

// AriResult represents the minimal age and grade to be able to read the text according automated readability index calculation.
type AriResult struct {
	// score      int8
	age        string
	gradeLevel string
}

// ariTable maps the ARI score to AriResult.
var ariTable = map[int]AriResult{
	1: {
		"5-6",
		"Kindengarden",
	},
	2: {
		"6-7",
		"First Grade",
	},
	3: {
		"7-8",
		"Second Grade",
	},
	4: {
		"8-9",
		"Third Grade",
	},
	5: {
		"9-10",
		"Forth Grade",
	},
	6: {
		"10-11",
		"Fifth Grade",
	},
	7: {
		"11-12",
		"Sixth Grade",
	},
	8: {
		"12-13",
		"Seventh Grade",
	},
	9: {
		"13-14",
		"Eighth Grade",
	},
	10: {
		"14-15",
		"Ninth Grade",
	},
	11: {
		"15-16",
		"Tenth Grade",
	},
	12: {
		"16-17",
		"Eleventh Grade",
	},
	13: {
		"17-18",
		"Twelfth Grade",
	},
	14: {
		"18-22",
		"College student",
	},
}

// ====== Functions ======

// CalculateARI accepts a non-empty string and returns the automated readability index (ARI) of it. The string has to have at least one word and at least one sentence (ended with `.`, `?`, `!`, or `...`)
// The result is always rounded up to the nearest whole number.
func CalculateAri(s string) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("Empty string.")
	}
	characters := float64(stats.CountCharacters(s))
	words := float64(stats.CountWords(s))
	sentences := float64(stats.CountSentences(s))

	if words == 0 || sentences == 0 {
		return 0, errors.New("No words of sentences in text. Cannot calculate ARI")
	}

	ariFloat := 4.71*(characters/words) + 0.5*(words/sentences) - 21.43
	// fmt.Println("Rough ARI:", ariFloat)
	score := int(math.Ceil(ariFloat))
	return score, nil
}

// ConvertARItoGrades accepts an ARI score as integer and returns the mapped to the score age and grade as strings.
//
// If no structure found, returns {"Unknown", "Unknown"}.
func ConvertAriToGrades(score int) (age, grade string) {
	if score > 14 {
		return "22+", "Professor level"
	}
	if _, ok := ariTable[score]; !ok {
		return "Unknown", "Unknown"
	}
	return ariTable[score].age, ariTable[score].gradeLevel
}
