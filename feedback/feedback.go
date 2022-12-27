package feedback

import "time"

// Feedback represents a feedback record.
type Feedback struct {
	ID          int64
	Username    string
	Type        string
	Language    string
	Description string
	Tenant      *Tenant
	Source      *Source
	Categories  []string
	Data        []byte
	Metadata    map[string]interface{}
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Tenant represents a tenant.
type Tenant struct {
	ID   int64
	Name string
	// other fields for storing tenant-specific data and metadata
}

// Source represents a feedback source.
type Source struct {
	ID   int64
	Name string
}
