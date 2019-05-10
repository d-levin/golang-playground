package todos

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

var repository = New()

var allowedMethods = []string{http.MethodGet, http.MethodPost, http.MethodDelete}

func Handler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		handleMethodGet(&w, req)
	case http.MethodPost:
		handleMethodPost(&w, req)
	case http.MethodDelete:
		handleMethodDelete(&w, req)
	default:
		w.Header().Add("Allow", strings.Join(allowedMethods, ", "))
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleMethodGet(w *http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) > 2 && parts[2] != "" {
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			(*w).WriteHeader(http.StatusBadRequest)
			return
		}
		getTodoById(id, w)
	} else {
		getAllTodos(w)
	}
}

func handleMethodPost(w *http.ResponseWriter, req *http.Request) {
	createTodo(w, req)
}

func handleMethodDelete(w *http.ResponseWriter, req *http.Request) {
	parts := strings.Split(req.URL.Path, "/")
	if len(parts) > 2 && parts[2] != "" {
		id, err := strconv.Atoi(parts[2])
		if err != nil {
			(*w).WriteHeader(http.StatusBadRequest)
			return
		}
		deleteTodoById(id)
	} else {
		(*w).WriteHeader(http.StatusBadRequest)
	}
}

func getAllTodos(w *http.ResponseWriter) {
	bytes, err := json.Marshal(repository.FindAll())
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}
	(*w).Header().Add("Content-Type", "application/json")
	_, _ = (*w).Write(bytes)
}

func getTodoById(id int, w *http.ResponseWriter) {
	bytes, err := json.Marshal(repository.Find(id))
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}
	(*w).Header().Add("Content-Type", "application/json")
	_, _ = (*w).Write(bytes)
}

func createTodo(w *http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	todo := Todo{}
	err := decoder.Decode(&todo)
	if err != nil {
		(*w).WriteHeader(http.StatusBadRequest)
		return
	}
	saved := repository.Save(&todo)
	bytes, err := json.Marshal(saved)
	(*w).Header().Add("Content-Type", "application/json")
	_, _ = (*w).Write(bytes)
}

func deleteTodoById(id int) {
	repository.Delete(id)
}
