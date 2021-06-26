package main

import "net/http"

type TodoServer struct {
}

func (t TodoServer) GetTodos(w http.ResponseWriter, r *http.Request, params GetTodosParams) {
	// our logic to retrieve all todos from a persistent layer
}

func (t TodoServer) CreateTodo(w http.ResponseWriter, r *http.Request) {
	// our logic to store the todo into a persistent layer
}

func (t TodoServer) DeleteTodo(w http.ResponseWriter, r *http.Request, todoId int32) {
	// our logic to delete a todo from the persistent layer
}

func (t TodoServer) UpdateTodo(w http.ResponseWriter, r *http.Request, todoId int32) {
	// our logic to update the todo.
}

func main() {
	s := TodoServer{}
	h := Handler(s)

	http.ListenAndServe(":3000", h)
}
