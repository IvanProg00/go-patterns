package db

import (
	"fmt"
	"sync"
)

type Database struct{}

var (
	db *Database
	mx = &sync.Mutex{}
)

func Connect() *Database {
	mx.Lock()
	defer mx.Unlock()

	if db != nil {
		fmt.Println("already connected to database")
	} else {
		fmt.Println("connecting to database")
		db = &Database{}
	}

	return db
}
