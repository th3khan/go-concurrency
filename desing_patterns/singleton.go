package main

import (
	"fmt"
	"sync"
	"time"
)

type DataBase struct{}

func (DataBase) CreateSingleConnection() {
	fmt.Println("Creating conection to DB...")
	time.Sleep(3 * time.Second)
	fmt.Println("Connection created successfull")
}

var db *DataBase
var lock sync.Mutex

func getDataBaseConection() *DataBase {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Create new conection to DB")
		db = &DataBase{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB Already Created..!")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDataBaseConection()
		}()
	}
	wg.Wait()
}
