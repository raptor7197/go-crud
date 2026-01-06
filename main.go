package main

import (
	"log"
	"net/http"
	"go-crud/handlers"
	"go-crud/storage"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
)

func writeError(w http.ResponseWriter,status int , message string ) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{"error":message})
	}


func main() {
	store, err := storage.NewSqliteTodoStore("todo.db")
	if err != nil {
		log.Fatal(err)
	}
	todoHandler := handlers.NewTodoHandler(store)


	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter,r *http.Request) {
		if r.URL.Path!="/" {
			http.NotFound(w,r)
			return
		}
		http.ServeFile(w,r,"html/index.html")
	})

	mux.Handle("/todos",handlers.WithMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				todoHandler.GetAllTodos(w,r)
			case http.MethodPost:
				todoHandler.CreateTodo(w,r)
			default:
				writeError(w,http.StatusMethodNotAllowed,"method not allowed")
			}
		}),
		handlers.RecoveryMiddleware,
		handlers.LoggingMiddleware,
	))
	

	mux.Handle("/todos/",handlers.WithMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				todoHandler.GetTodoByID(w,r)
			case http.MethodPut:
				todoHandler.UpdateTodo(w,r)
			case http.MethodDelete:
				todoHandler.DeleteTodo(w,r)
			default:
				writeError(w,http.StatusMethodNotAllowed,"method not allowed")
			}
		}),
		handlers.RecoveryMiddleware,
		handlers.LoggingMiddleware,
	))

	

	log.Println("server running on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
