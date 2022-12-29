package orchestrator

import (
	"errors"
	"feed-processor/feedback"
	"feed-processor/repository"
	"feed-processor/tenant"
	"fmt"
	"time"
)

// OrchestratorStore represents the interface for the orchestrator store.
type OrchestratorStore interface {
	Integrate(tenantID string) error
}

// Orchestrator is an object that is responsible for coordinating
// the integration of various feedback sources for a given tenant.
type Orchestrator struct {
	DB          repository.RepositoryStore
	TenantStore tenant.TenantStore
}

// New creates a new instance of the orchestrator store.
func New(
	db repository.RepositoryStore,
	tenantStore tenant.TenantStore,
) OrchestratorStore {
	return &Orchestrator{
		DB:          db,
		TenantStore: tenantStore,
	}
}

// Integrate integrates feedback records from various sources.
func (o *Orchestrator) Integrate(tenantID string) error {
	// retrieve the tenant information
	tenant, err := o.TenantStore.GetTenant(tenantID)
	if err != nil {
		return err
	}
	if tenant == nil {
		return errors.New(fmt.Sprintf("tenant with ID %s not found", tenantID))
	}
	if len(tenant.Integrators) == 0 {
		return errors.New(fmt.Sprintf("tenant with ID %q has no integrators", tenantID))
	}
	for _, i := range tenant.Integrators {
		// TODO: implement logic of time range
		// Fetch the data from the integrator.
		records, err := i.FetchData(time.Now(), time.Now())
		if err != nil {
			return err
		}
		// Process the data.
		feedbacks, err := i.ProcessData(records)
		for i, _ := range feedbacks {
			feedbacks[i].Tenant = &feedback.Tenant{
				ID:   tenant.ID,
				Name: tenant.Name,
			}
		}
		if err != nil {
			return err
		}
		// Store the data in the repository.
		if err := i.StoreData(feedbacks, o.DB); err != nil {
			return err
		}
	}
	return nil
}
