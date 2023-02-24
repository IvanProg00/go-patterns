package db

import (
	"fmt"
	"sync"
)

type Database struct {
	mx          sync.Mutex
	isConnected bool
}

func New() *Database {
	return &Database{}
}

func (db *Database) Connect() {
	db.mx.Lock()
	defer db.mx.Unlock()

	if db.isConnected {
		fmt.Println("already connected to database")
	} else {
		fmt.Println("connecting to database")
		db.isConnected = true
	}
}
