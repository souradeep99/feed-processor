package integrators

import (
	"feed-processor/feedback"
	"feed-processor/repository"
	"time"
)

type PlaystoreIntegrator struct {
	// TODO: Add other fields as needed.
	AppPackageNames []string
}

// NewPlaystoreIntegrator returns a new instance of PlaystoreIntegrator.
func NewPlaystoreIntegrator(appPackageNames []string) Integrator {
	return &PlaystoreIntegrator{
		AppPackageNames: appPackageNames,
	}
}

// FetchData fetches feedback data from the Playstore.
func (p *PlaystoreIntegrator) FetchData(startTime time.Time, endTime time.Time) (interface{}, error) {
	// TODO: Implement code to fetch feedback data from the Playstore.
	// You may want to use the androidpublisher library (https://godoc.org/google.golang.org/api/androidpublisher/v3)
	// to make requests to the Playstore API.
	var fb []*feedback.Feedback
	for _, _ = range p.AppPackageNames {
		// TODO: Implement code to fetch feedback data for a single app.
	}
	return fb, nil
}

// ProcessData processes the raw feedback data into a uniform internal structure.
func (p *PlaystoreIntegrator) ProcessData(rawData interface{}) ([]*feedback.Feedback, error) {
	// TODO: Implement code to process the raw feedback data into a uniform internal structure.
	// You may want to enrich the feedback with additional information from external sources,
	// classify the feedback, or perform other transformations as needed.
	return nil, nil
}

// StoreData stores the processed feedback data.
func (p *PlaystoreIntegrator) StoreData(records []*feedback.Feedback, db repository.RepositoryStore) error {
	// TODO: implement this method
	// insert the processed feedback records into the database
	err := db.StoreData(records)
	return err
}
