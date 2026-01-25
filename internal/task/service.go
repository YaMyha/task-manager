package task

import "errors"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Add(title string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}

	tasks, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	tasks = append(tasks, Task{
		ID:    id,
		Title: title,
		Done:  false,
	})

	return s.repo.Save(tasks)
}
