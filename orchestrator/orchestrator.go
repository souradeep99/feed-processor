package orchestrator

import "feed-processor/integrators"

// Orchestrator coordinates the integration process.
type Orchestrator struct {
	DB *database.DB
}

// Integrate integrates feedback records from various sources.
func (o *Orchestrator) Integrate(integrators []integrators.Integrator) error {
	for _, i := range integrators {
		records, err := i.FetchData()
		if err != nil {
			return err
		}

		records, err = i.ProcessData(records)
		if err != nil {
			return err
		}

		if err := i.StoreData(records, o.DB); err != nil {
			return err
		}
	}
	return nil
}
