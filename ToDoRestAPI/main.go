package main

import (
	"fmt"
	"pavlyysh/ToDoList/http"
	"pavlyysh/ToDoList/todo"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed starting server")
	}
}
