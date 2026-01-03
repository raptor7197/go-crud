package main

import (
	"log"
	"net/http"

	"go-crud/handlers"
	"go-crud/storage"
)

func main() {
	store := storage.NewTodoStore()
	todoHandler := handlers.NewTodoHandler(store)

	mux := http.NewServeMux()

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.GetAllTodos(w, r)
		case http.MethodPost:
			todoHandler.CreateTodo(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.GetTodoByID(w, r)
		case http.MethodPut:
			todoHandler.UpdateTodo(w, r)
		case http.MethodDelete:
			todoHandler.DeleteTodo(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("server running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
