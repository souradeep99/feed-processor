package lib

import (
	"strings"
	"unicode"
)

// UniqueWords returns the unique words in a slice of words
func UniqueWords(words []string) []string {
	uniqueWordsMap := make(map[string]bool)
	for _, word := range words {
		uniqueWordsMap[word] = true
	}
	var uniqueWords []string
	for word := range uniqueWordsMap {
		uniqueWords = append(uniqueWords, word)
	}
	return uniqueWords
}

// RemovePunctuation removes punctuation from a string
func RemovePunctuation(text string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, text)
}

// RemoveStopWords removes stop words from a string
func RemoveStopWords(text string) string {
	// create a slice of stop words
	stopWords := []string{"a", "an", "the", "and", "or", "but", "so"}
	// split the text into words
	words := strings.Split(text, " ")
	// create a slice of non-stop words
	var nonStopWords []string
	for _, word := range words {
		if !Contains(stopWords, word) {
			nonStopWords = append(nonStopWords, word)
		}
	}
	// join the non-stop words into a string
	return strings.Join(nonStopWords, " ")
}

// Contains checks if a slice contains a value
func Contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
