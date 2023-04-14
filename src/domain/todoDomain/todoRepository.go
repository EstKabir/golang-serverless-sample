package todoDomain

type Repository interface {
	GetById(id string) (*Model, error)
	FindAll() (*[]Model, error)
	Create(endpointData *Model) (*Model, error)
	Update(id string, endpointData *Model) (*Model, error)
	Delete(id string) error
}
