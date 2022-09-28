package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

var db *Database
var lock sync.Mutex

func (Database) CreateSingleConnection() {
	fmt.Println("Creating Singleton for Database")
	time.Sleep(2 * time.Second)
	fmt.Println("Connected successfully")
}

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()

	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("Database already connected")
	}

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}

	wg.Wait()
}
