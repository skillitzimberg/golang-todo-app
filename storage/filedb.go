package storage

import (
	"fmt"
	"log"
	"os"
)

type FileDB struct {
	source string
}

func NewFileDB(source string) (fdb FileDB) {
	fdb.source = source
	return
}

func (db FileDB) Create(todo string) {
	f, err := os.OpenFile(db.source, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}

	f.Write([]byte(fmt.Sprintf("%s\n", todo)))
}
