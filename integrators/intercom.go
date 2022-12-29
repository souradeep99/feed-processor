package integrators

import (
	"feed-processor/feedback"
	"feed-processor/repository"
	"time"
)

// IntercomIntegrator is an implementation of the Integrator interface for Intercom.
type IntercomIntegrator struct {
	AppID   string
	APIKey  string
	BaseURL string
}

func NewIntercomIntegrator(
	appID string,
	apiKey string,
	baseURL string,
) Integrator {
	return &IntercomIntegrator{
		AppID:   appID,
		APIKey:  apiKey,
		BaseURL: baseURL,
	}
}

// FetchData fetches feedback records from the Intercom source.
func (t *IntercomIntegrator) FetchData(startTime time.Time, endTime time.Time) (interface{}, error) {
	// TODO: Implement code to fetch feedback data from Intercom.
	// You may want to use the intercom-go library (https://github.com/intercom/intercom-go)
	// to make requests to the Intercom API.
	return nil, nil
}

// ProcessData processes the raw data from the FetchData method and transforms it into a uniform internal structure
func (t *IntercomIntegrator) ProcessData(rawData interface{}) ([]*feedback.Feedback, error) {
	// TODO: Implement code to process the raw feedback data into a uniform internal structure.
	// You may want to enrich the feedback with additional information from external sources,
	// classify the feedback, or perform other transformations as needed.
	return nil, nil
}

// StoreData stores the processed feedback records in the database.
func (t *IntercomIntegrator) StoreData(records []*feedback.Feedback, db repository.RepositoryStore) error {
	// TODO: implement this method
	// insert the processed feedback records into the database
	err := db.StoreData(records)
	return err
}
