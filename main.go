package main

import (
	"feed-processor/integrators"
	"feed-processor/orchestrator"
)

func main() {
	// create a database connection
	db, err := database.New("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		// handle error
	}
	defer db.Close()

	// create integrators for each feedback source
	discourseIntegrator := &integrators.DiscourseIntegrator{
		BaseURL: "https://meta.discourse.org",
	}

	// create orchestrator and pass in the integrators
	orch := &orchestrator.Orchestrator{
		DB: db,
	}
	if err := orch.Integrate([]integrators.Integrator{
		discourseIntegrator,
	}); err != nil {
		// handle error
	}
}
