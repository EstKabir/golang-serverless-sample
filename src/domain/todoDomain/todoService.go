package todoDomain

type Service interface {
	FindById(id string) (*Model, error)
	Create(todo *Model) (*Model, error)
	Update(id string, todo *Model) (*Model, error)
	FindAll() (*[]Model, error)
	Delete(id string) error
}
