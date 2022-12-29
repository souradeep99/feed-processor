package repository

// RepositoryStore is an interface for a repository store.
type RepositoryStore interface {
	StoreData(data interface{}) error
}

type Repository struct {
	// TODO: Add other fields as needed.
}

// New returns a new instance of Repository.
func New() RepositoryStore {
	return &Repository{}
}

// StoreData stores data in the repository.
func (r *Repository) StoreData(data interface{}) error {
	return nil
}
