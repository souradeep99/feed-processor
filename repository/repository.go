package repository

type RepositoryStore interface {
	StoreData(data interface{}) error
}

type Repository struct {
}

func New() RepositoryStore {
	return &Repository{}
}

func (r *Repository) StoreData(data interface{}) error {
	return nil
}
