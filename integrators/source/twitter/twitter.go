package twitter

import (
	"feed-processor/feedback"
	"feed-processor/integrators"
	"feed-processor/repository"
	"time"
)

// TwitterIntegrator is an implementation of the Integrator interface for Twitter.
type TwitterIntegrator struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// NewTwitterIntegrator returns a new instance of Integrator.
func NewTwitterIntegrator(
	consumerKey string,
	consumerSecret string,
	accessToken string,
	accessSecret string,
) integrators.Integrator {
	return &TwitterIntegrator{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		AccessToken:    accessToken,
		AccessSecret:   accessSecret,
	}
}

// FetchData fetches feedback records from the Twitter source.
func (t *TwitterIntegrator) FetchData(startTime time.Time, endTime time.Time) (interface{}, error) {
	// TODO: Implement code to fetch feedback data from Twitter.
	// we can use the go-twitter library (https://github.com/dghubble/go-twitter)
	// to make requests to the Twitter API.
	return nil, nil
}

// ProcessData processes the raw data from the FetchData method and transforms it into a uniform internal structure
func (t *TwitterIntegrator) ProcessData(rawData interface{}) ([]*feedback.Feedback, error) {
	// TODO: Implement code to process the raw feedback data into a uniform internal structure.
	// we can enrich the feedback with additional information from external sources,
	// classify the feedback, or perform other transformations as needed.
	return nil, nil
}

// StoreData stores the processed feedback records in the database.
func (t *TwitterIntegrator) StoreData(records []*feedback.Feedback, db repository.RepositoryStore) error {
	// TODO: implement this method
	// insert the processed feedback records into the database
	err := db.StoreData(records)
	return err
}
