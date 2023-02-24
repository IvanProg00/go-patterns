package singleton

import (
	"sync"

	"github.com/IvanProg00/go-patterns/singleton/db"
)

func Run() {
	database := db.New()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			database.Connect()
		}()
	}

	wg.Wait()
}
