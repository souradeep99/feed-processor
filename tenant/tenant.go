package tenant

import (
	"errors"
	"feed-processor/integrators"
)

// Tenant represents a tenant in the system
type Tenant struct {
	ID          string
	Name        string
	Integrators []integrators.Integrator
}

// TenantStore is an interface for storing and retrieving tenant information
type TenantStore interface {
	GetTenant(id string) (*Tenant, error)
	SaveTenant(tenants []*Tenant)
}

// MemoryTenantStore is a simple in-memory implementation of TenantSto re
type MemoryTenantStore struct {
	tenants map[string]*Tenant
}

// NewMemoryTenantStore creates a new MemoryTenantStore
func NewMemoryTenantStore() TenantStore {
	return &MemoryTenantStore{
		tenants: make(map[string]*Tenant),
	}
}

// GetTenant retrieves a tenant by ID
func (store *MemoryTenantStore) GetTenant(id string) (*Tenant, error) {
	tenant, ok := store.tenants[id]
	if !ok {
		return nil, errors.New("tenant not found")
	}
	return tenant, nil
}

// SaveTenant stores a tenant in the store
func (store *MemoryTenantStore) SaveTenant(tenants []*Tenant) {
	for _, tenant := range tenants {
		store.tenants[tenant.ID] = tenant
	}
}
