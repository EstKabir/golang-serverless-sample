package product

type Service interface {
	FindById(id string) (*Model, error)
	Create(endpointConfig *Model) (*Model, error)
	Update(id string, endpointConfig *Model) (*Model, error)
	FindAll() ([]*Model, error)
	Delete(id string) error
}
