package todos

import (
	"testing"
)

func Test_Repository(t *testing.T) {
	r := TodoRepository{}

	if len(r.todos) != 0 {
		t.Error("Expected 0, got", len(r.todos))
	}

	todo := r.Save(&Todo{Name: "homework"})
	if todo.Id <= 0 {
		t.Error("Expected > 0, got", todo.Id)
	}
	if todo.Name != "homework" {
		t.Error("Expected homework, got", todo.Name)
	}
	if len(r.todos) != 1 {
		t.Error("Expected 1, got", len(r.todos))
	}

	id1 := todo.Id

	todo = r.Find(id1)
	if todo == nil {
		t.Error("Expected todo, got nil")
	}

	todo = r.Find(id1 - 1)
	if todo != nil {
		t.Error("Expected nil, got", todo)
	}

	todos := r.FindAll()
	if len(todos) != 1 {
		t.Error("Expected 1, got", len(todos))
	}

	id2 := r.Save(&Todo{Name: "taxes"}).Id
	if len(r.todos) != 2 {
		t.Error("Expected 2, got", len(r.todos))
	}

	r.Delete(id1)
	todos = r.FindAll()
	if len(todos) != 1 {
		t.Error("Expected 1, got", len(todos))
	}
	if todos[0].Id != id2 {
		t.Errorf("Expected %d, got %d", id1, id1)
	}
	if todos[0].Name != "taxes" {
		t.Error("Expected taxes, got", todos[0].Name)
	}

	r.Delete(id2)
	todos = r.FindAll()
	if len(todos) != 0 {
		t.Error("Expected 0, got", len(todos))
	}
}
