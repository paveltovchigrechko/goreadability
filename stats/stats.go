// Package stats provides basic functions to count and calculate strings statistics, such as number of characters, words, sentences, and so on.
package stats

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ====== Types & Consts ======

type TotalStats struct {
	Symbols    int
	Characters int
	Words      int
	Sentences  int
}

var abbreviations = []string{
	"U.S.",

	"Mr.",
	"Messrs.",
	"Mrs.",
	"Mmes.",
	"Ms.",
	"Dr.",
	"Prof.",
	"Capt.",
	"St.",
	"Revd.",
	"Rev.",

	"Jan.",
	"Feb.",
	"Mar.",
	"Apr.",
	"Aug.",
	"Sept.",
	"Oct.",
	"Nov.",
	"Dec.",

	"a.m.",
	"p.m.",
	"i.e.",
	"e.g.",
	"a.d.",
	"b.c.",
	"b.c.e.",
	"c.e.",
	"n.b.",
}

// ====== Methods ======

func (stats TotalStats) Print() {
	fmt.Println("Symbols:\t", stats.Symbols)
	fmt.Println("Characters:\t", stats.Characters)
	fmt.Println("Words:\t\t", stats.Words)
	fmt.Println("Sentences:\t", stats.Sentences)
}

// ====== Functions ======

// For debugging and testing purposes
func CountAllStats(text string) TotalStats {
	var result TotalStats
	result.Symbols = CountSymbols(text)
	result.Characters = CountCharacters(text)
	result.Words = CountWords(text)
	result.Sentences = CountSentences(text)
	return result
}

// CountSymbols accepts a string and returns the number of symbols in it.
// The string should not have trailing spaces before new lines.
// Only new lines do not count as symbols.
// An ellipsis ... counts as one symbol, an ellipsis in brackets [...] counts as three symbols. (?)
func CountSymbols(s string) int {
	if len(s) == 0 {
		return 0
	}
	ellipsis := strings.Count(s, "...")
	newLines := strings.Count(s, "\n")
	total := utf8.RuneCountInString(s) - newLines - 2*ellipsis
	return total
}

// CountCharacters accepts a string and returns the number of characters.
// A character is a letter or a digit.
func CountCharacters(s string) int {
	if len(s) == 0 {
		return 0
	}
	chars := 0
	for _, char := range s {
		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			chars++
		}
	}
	return chars
}

// CountWords accepts a string and returns the number of words in it.
// The string should not have trailing spaces before new lines (e.g. "Word. \nAnother word." isn't counted correctly), nor double newlines (e.g. "Word.\n\nAnother word.")
// Numbers count as a word (for example, "44." returns `1`, and "12 and 43." returns `3`).
// TODO: case with multiple sequential new lines. ("One.\n\nTwo." => must return `2`).
// TODO: En Dash in dates ("1845-1851" should be 2 words(?))
func CountWords(s string) int {
	if len(s) == 0 {
		return 0
	}
	if strings.Count(s, "\n") > 0 {
		s = strings.ReplaceAll(s, "\n", " ")
	}
	words := len(strings.Fields(s))
	return words
}

// CountSentences accepts a string and returns the number of sentences in it.
// TODO: cases "?!", "???", "!!!", "...", "!?" must count as one sentence.
// TODO: case when point is used in abbreviation ("U.S.", "Mr.", "Jr.", "Dec. 9, 1991", see abbreviations above).
// TODO: ellipsis as an omission ("The witnesses reported that the suspect fled the scene ... and headed west toward the highway.")
// TODO: cases with dots in fractions ("10.5 pbs." should return `1`.)
// TODO: general case when there is no space after the finishing point. Should not count as a sentence.
func CountSentences(s string) int {
	if len(s) == 0 {
		return 0
	}

	points := strings.Count(s, ".")
	exclamations := strings.Count(s, "!")
	questions := strings.Count(s, "?")
	//ellipsis := strings.Count(s, "...")
	return points + exclamations + questions //- 2*ellipsis
}
