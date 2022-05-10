package main

import (
	"os"
)

func main() {
	newTodos := os.Args[1:]
	Write(newTodos)
}
