package main

import (
	"os"

	"github.com/skillitzimberg/golang-todo-app/storage"
)

func main() {
	newTodos := os.Args[1:]
	storage.Write(newTodos)
}
