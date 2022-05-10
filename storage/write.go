package storage

import (
	"fmt"
	"log"
	"os"
)

func Write(newTodos []string) {
	f, err := os.OpenFile("todos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, newTodo := range newTodos {
		f.Write([]byte(fmt.Sprintf("%s\n", newTodo)))
	}
}
