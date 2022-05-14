package main

import (
	"fmt"
	"log"
	"os"

	"github.com/skillitzimberg/golang-todo-app/model/todo"
)

type Storage interface {
	Create(Todo)
}

type FileDB struct {
	source string
}

func NewFileDB(source string) (fdb FileDB) {
	fdb.source = source
	return
}

func (db FileDB) Create(todo Todo) {
	f, err := os.OpenFile(db.source, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}

	f.Write([]byte(fmt.Sprintf("%s\n", todo.desc)))
}

func WriteTodo(storage Storage, todo Todo) {
	storage.Create(todo)
}

func main() {
	newTodos := os.Args[1:]

	fdb := NewFileDB("todos.txt")
	for _, newTodo := range newTodos {
		WriteTodo(fdb, NewTodo(newTodo))
	}
}