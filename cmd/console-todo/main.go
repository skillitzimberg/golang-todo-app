package main

import (
	"os"

	"golang-todo-app/storage"
)

func WriteTodo(storage storage.Storage, todo string) {
	storage.Create(todo)
}

func main() {
	newTodos := os.Args[1:]

	fdb := storage.NewFileDB("todos.txt")
	for _, newTodo := range newTodos {
		WriteTodo(fdb, newTodo)
	}
}
