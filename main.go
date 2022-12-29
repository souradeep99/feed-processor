package main

import (
	"feed-processor/config"
	"feed-processor/orchestrator"
	"feed-processor/repository"
	"feed-processor/tenant"
)

func main() {
	// create a database connection
	db := repository.New()

	store := tenant.NewMemoryTenantStore()

	tenants, err := config.GetTenants()
	if err != nil {
		// handle error
	}
	store.SaveTenant(tenants)

	// create orchestrator and use in the integrators
	orch := orchestrator.New(db, store)

	for _, t := range tenants {
		if err := orch.Integrate(t.ID); err != nil {
			// handle error
		}
	}
}
