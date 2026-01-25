package storage

import (
	"encoding/json"
	"os"
	"task-manager/internal/task"
)

type JSONStorage struct {
	FilePath string
}

func NewJSONStorage(path string) *JSONStorage {
	return &JSONStorage{FilePath: path}
}

func (s *JSONStorage) GetAll() ([]task.Task, error) {
	data, err := os.ReadFile(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []task.Task{}, nil
		}
		return nil, err
	}

	var tasks []task.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *JSONStorage) Save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, data, 0644)
}
