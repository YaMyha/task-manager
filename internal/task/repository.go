package task

type Repository interface {
	GetAll() ([]Task, error)
	Save([]Task) error
}
