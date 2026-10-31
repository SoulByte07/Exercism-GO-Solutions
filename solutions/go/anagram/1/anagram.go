package anagram

import (
	"slices"
	"strings"
)

// sortRunes converts a string to lowercase runes, sorts them, and returns them as a sorted string.
func sortRunes(s string) string {
	runes := []rune(strings.ToLower(s))
	slices.Sort(runes)
	return string(runes)
}

// Detect filters candidates to find all valid anagrams of the subject word.
func Detect(subject string, candidates []string) []string {
	// Initialize res as an empty slice so it serializes to [] rather than nil if no match is found
	res := []string{}

	lowerSubject := strings.ToLower(subject)
	sortedSubject := sortRunes(subject)

	for _, cand := range candidates {
		lowerCand := strings.ToLower(cand)

		// Rule 1: A word is not an anagram of itself (case-insensitive)
		if lowerCand == lowerSubject {
			continue
		}

		// Rule 2: Length shortcut (optimization: different rune counts cannot be anagrams)
		if len([]rune(lowerCand)) != len([]rune(lowerSubject)) {
			continue
		}

		// Rule 3: Must share the exact same character distribution
		if sortRunes(cand) == sortedSubject {
			// Append the candidate keeping its original letter casing
			res = append(res, cand)
		}
	}

	return res
}