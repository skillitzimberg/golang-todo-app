package main

import (
	"fmt"
	"os"
)

type Storage interface {
	Create(interface{})
}

type Todo struct {
	desc string
}

func NewTodo(desc string) (todo Todo) {
	todo.desc = desc
	return
}

type FileDB struct {
	source string
}

func NewFileDB(source string) (fdb FileDB) {
	fdb.source = source
	return
}

func (db FileDB) Create(todo Todo)

func main() {
	newTodos := os.Args[1:]
	fmt.Print(newTodos)
}
