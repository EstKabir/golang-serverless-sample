package product

type Repository interface {
	FindById(id string) (*Model, error)
	Create(endpointData *Model) (*Model, error)
	Update(id string, endpointData *Model) (*Model, error)
	FindAll() ([]*Model, error)
	Delete(id string) error
}
