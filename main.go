 package main

 import ( 
	 "log" 
	 "net/http"
	 "go-crud/handlers"
	 "go-crud/storage"
 )

 func main() {
	 store:=storage.newTodoStore()
	 handler := handlers.NewTodoStore(store)

	 http.handleFunc("/todos",func (w http.ResponseWriter,r *http.Request)
	if r.Method == http.MethodPost {
	handler.CreateTodo(w,r)
	return
	}
	http.Error(w,"method not allowed",http.StatusMethodNotAllowed)
 )
 }

 log.Println("server runnin on port 8080")
 http.ListenandServe(":8080",nil)
