package todos

import (
	"math"
	"math/rand"
	"sort"
)

type Todo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type TodoRepository struct {
	todos []*Todo
}

func New() *TodoRepository {
	return &TodoRepository{todos: []*Todo{}}
}

func (t *TodoRepository) Save(todo *Todo) *Todo {
	todo.Id = rand.Intn(math.MaxInt64)
	t.todos = append(t.todos, todo)
	return todo
}

func (t *TodoRepository) Find(id int) *Todo {
	for _, todo := range t.todos {
		if todo.Id == id {
			return todo
		}
	}
	return nil
}

func (t *TodoRepository) FindAll() []*Todo {
	return t.todos
}

func (t *TodoRepository) Delete(id int) {
	i := sort.Search(len(t.todos), func(i int) bool { return t.todos[i].Id >= id })

	if i < len(t.todos) && t.todos[i].Id == id {
		t.todos = append(t.todos[:i], t.todos[i+1:]...)
	}
}
