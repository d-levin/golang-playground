package todos

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_Handler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/todos", nil)
	if err != nil {
		t.Error(err)
	}
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(responseRecorder, req)
	decoder := json.NewDecoder(responseRecorder.Body)
	var todos []*Todo
	err = decoder.Decode(&todos)
	if err != nil {
		t.Error("Expected nil, got", err)
	}
	if len(todos) != 0 {
		t.Error("Expected 0, got", len(todos))
	}

	body := bytes.NewReader([]byte(`{"name": "task-1"}`))
	req, err = http.NewRequest(http.MethodPost, "/todos", body)
	if err != nil {
		t.Error(err)
	}
	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, req)
	decoder = json.NewDecoder(responseRecorder.Body)
	var todo *Todo
	err = decoder.Decode(&todo)
	if err != nil {
		t.Error("Expected nil, got", err)
	}
	if todo.Id <= 0 {
		t.Error("Expected > 0, got", todo.Id)
	}
	if todo.Name != "task-1" {
		t.Error("Expected task-1, got", todo.Name)
	}

	req, err = http.NewRequest(http.MethodGet, "/todos", nil)
	if err != nil {
		t.Error(err)
	}
	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, req)
	decoder = json.NewDecoder(responseRecorder.Body)
	err = decoder.Decode(&todos)
	if err != nil {
		t.Error("Expected nil, got", err)
	}
	if len(todos) != 1 {
		t.Error("Expected 1, got", len(todos))
	}
	if todos[0].Id != todo.Id {
		t.Errorf("Expected %d, got %d", todo.Id, todos[0].Id)
	}
	if todos[0].Name != todo.Name {
		t.Errorf("Expected %s, got %s", todo.Name, todos[0].Name)
	}

	id, name := todo.Id, todo.Name
	req, err = http.NewRequest(http.MethodGet, "/todos/"+strconv.Itoa(id), nil)
	if err != nil {
		t.Error(err)
	}
	responseRecorder = httptest.NewRecorder()
	handler.ServeHTTP(responseRecorder, req)
	decoder = json.NewDecoder(responseRecorder.Body)
	err = decoder.Decode(&todo)
	if err != nil {
		t.Error("Expected nil, got", err)
	}
	if todo.Id != id {
		t.Errorf("Expected %d, got %d", id, todo.Id)
	}
	if todo.Name != name {
		t.Errorf("Expected %s, got %s", name, todo.Name)
	}
}
