package task

import "testing"

func TestService_Add(t *testing.T) {
	repo := &fakeRepo{}
	service := NewService(repo)

	err := service.Add("New Task")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(repo.tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(repo.tasks))
	}

	task := repo.tasks[0]
	if task.Title != "New Task" {
		t.Errorf("Expected task title 'New Task', got '%s'", task.Title)
	}

	if task.Done {
		t.Errorf("Expected task to be not done")
	}
}

func TestService_Add_EmptyTitle(t *testing.T) {
	repo := &fakeRepo{}
	service := NewService(repo)

	err := service.Add("")
	if err == nil {
		t.Fatalf("Expected error for empty title, got nil")
	}

	if len(repo.tasks) != 0 {
		t.Fatalf("Expected 0 tasks, got %d", len(repo.tasks))
	}
}
