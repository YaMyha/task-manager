package task

type fakeRepo struct {
	tasks []Task
}

func (f *fakeRepo) GetAll() ([]Task, error) {
	return f.tasks, nil
}

func (f *fakeRepo) Save(tasks []Task) error {
	f.tasks = tasks
	return nil
}
