package classification

import (
	"feed-processor/lib"
	"fmt"
	"strings"
)

const (
	minLen = 1
	maxLen = 10
)

// ClassifyFeedback classifies the feedback into categories using a machine learning model
func ClassifyFeedback(description string) ([]string, error) {
	// load the machine learning model from a file or a remote location
	model, err := loadModel()
	if err != nil {
		return nil, fmt.Errorf("failed to load model: %w", err)
	}
	// preprocess the description to prepare it for the model
	preprocessedDescription := preprocessDescription(description)
	// make a prediction using the model
	prediction, err := model.Predict(preprocessedDescription)
	if err != nil {
		return nil, fmt.Errorf("failed to make prediction: %w", err)
	}
	// convert the prediction to a slice of categories
	categories := convertPredictionToCategories(prediction)
	return categories, nil
}

// loadModel loads the machine learning model from a file or a remote location
func loadModel() (Model, error) {
	// TODO: implement this function
	return nil, nil
}

// preprocessDescription preprocesses the feedback description for the machine learning model
func preprocessDescription(description string) interface{} {
	normalizedDescription := normalizeText(description)
	filteredDescription := filterText(normalizedDescription, minLen, maxLen)
	preprocessedDescription := convertToNumeric(filteredDescription)
	return preprocessedDescription
}

// convertPredictionToCategories converts the prediction from the machine learning model to a slice of categories
func convertPredictionToCategories(prediction interface{}) []string {
	// assert that the prediction is a slice of strings
	categories, ok := prediction.([]string)
	if !ok {
		return []string{}
	}
	return categories
}

// convertToNumeric converts a string to a numerical representation using a bag of words approach
func convertToNumeric(text string) []float64 {
	// split the text into words
	words := strings.Split(text, " ")
	// create a map of word counts
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}
	// create a slice of word counts for each unique word in the text
	var wordCountSlice []float64
	for _, word := range lib.UniqueWords(words) {
		wordCountSlice = append(wordCountSlice, float64(wordCounts[word]))
	}
	return wordCountSlice
}

// normalizeText removes punctuation, stop words, and makes the text lowercase
func normalizeText(text string) string {
	// remove punctuation
	text = lib.RemovePunctuation(text)
	// remove stop words
	text = lib.RemoveStopWords(text)
	// make the text lowercase
	text = strings.ToLower(text)
	return text
}

// filterText removes words that are too short or too long
func filterText(text string, minLength int, maxLength int) string {
	// split the text into words
	words := strings.Split(text, " ")
	// create a slice of valid words
	var validWords []string
	for _, word := range words {
		if len(word) >= minLength && len(word) <= maxLength {
			validWords = append(validWords, word)
		}
	}
	// join the valid words into a string
	return strings.Join(validWords, " ")
}
