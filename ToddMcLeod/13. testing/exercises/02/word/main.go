// Package word provides custom functions for working with strings and words
package word

import "strings"

// UseCOunt returns the number of time each word is used in a string
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string
func Count(s string) int {
	return len(strings.Fields(s))
}
