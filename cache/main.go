package main

import (
	"fmt"
	"sync"
	"time"
)

// Function to calculate fibonacci
func Fibonacci(n int) int {
	if n <= 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Memory holds a function and a map of results
type Memory struct {
	f     Function               // Function to be used
	cache map[int]FunctionResult // Map of results for a given key
	lock  sync.Mutex             // Mutex to protect the cache edit
}

// A function has to recive a value and return a value and an error
type Function func(key int) (interface{}, error)

// The result of a function
type FunctionResult struct {
	value interface{}
	err   error
}

// NewCache creates a new cache
func NewCache(f Function) *Memory {
	return &Memory{f, make(map[int]FunctionResult), sync.Mutex{}}
}

// Get returns the value for a given key
func (m *Memory) Get(key int) (interface{}, error) {
	// Block until the cache is verified
	m.lock.Lock()
	// Check if the value is in the cache
	res, exist := m.cache[key]
	m.lock.Unlock()
	// If the value is not in the cache, calculate it
	if !exist {
		// Block until the cache is edited
		m.lock.Lock()
		res.value, res.err = m.f(key) // Calculate the value
		m.cache[key] = res            // Store the value in the cache
		m.lock.Unlock()
	}
	return res.value, res.err
}

// Function to be used in the cache
func GetFibonacci(key int) (interface{}, error) {
	return Fibonacci(key), nil
}

func main() {
	// Create a cache and some values
	cache := NewCache(GetFibonacci)
	values := []int{42, 40, 41, 42, 38, 41}

	// Create a wait group
	var wg sync.WaitGroup

	// For each value to calculate, get the value and print the time it took to calculate
	for _, v := range values {
		// Funtion anonimous to be used in the cache
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			start := time.Now()

			value, err := cache.Get(v)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Key %d, Value: %d, time: %s \n", v, value, time.Since(start))
		}(v)
	}
	wg.Wait()
}
