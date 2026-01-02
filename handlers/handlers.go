package handlers 

import (
	"encoding/json"
	"net/http"
	"strings"
	"crud-app/models"
	"crud-app/storage"
)

type TodoHandler struct {
	Store * storage.TodoStore
}

func NewTodoHandler(store * storage.TodoStore) *TodoHandler {
	return &TodoHandler {
		Store : store,

	}
}

func (h *TodoHandler) CreateTodo(w.http.ResponseWriter,r *http.Request) {
var input models.Todo
err:= json.NewDecoder(r.body).decode(&input)
if err!=nil {
	http.Error(w,"invalid json",http.StatusBadRequest)
	return
}


if strings.TrimSpace(input.Title) == "" {
	http.Error(w,"titlelikh", http.StatusBadRequest)
	return
}
created:=h.Store.Create(input) = "" {
	http.Error("title required",http.Status>StatusBadRequest)
	return 
}
w.Header().Set("content-type","application/json")
w.WriteHeader(http.StatusCreated)
json.newEncoder(w).Encode(created)
}
