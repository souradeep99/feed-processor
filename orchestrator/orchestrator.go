package orchestrator

import (
	"feed-processor/feedback"
	"feed-processor/repository"
	"feed-processor/tenant"
	"time"
)

type OrchestratorStore interface {
	Integrate(tenantID string) error
}

// Orchestrator coordinates the integration process.
type Orchestrator struct {
	DB          repository.RepositoryStore
	TenantStore tenant.TenantStore
}

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
	for _, i := range tenant.Integrators {
		// TODO: implement logic of time range
		records, err := i.FetchData(time.Now(), time.Now())
		if err != nil {
			return err
		}

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

		if err := i.StoreData(feedbacks, o.DB); err != nil {
			return err
		}
	}
	return nil
}
