package main

import (
	"fmt"
	"sync"
	"time"
)

// function to simulate a work hard process
func ExpensiveFibonacci(n int) int {
	fmt.Printf("Calculating fibonacci for (%d)\n", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool       // map to keep track of in progress requests
	IsPending  map[int][]chan int // map to keep track of pending requests
	Lock       sync.RWMutex
}

func (s *Service) Work(job int) {
	s.Lock.RLock()
	exist := s.InProgress[job]

	if exist {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)
		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()
		fmt.Printf("Waiting for job (%d) to complete\n", job)
		result := <-response
		fmt.Printf("Job (%d) completed with result (%d)\n", job, result)
		return
	}
	s.Lock.RUnlock()
	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculating fibonacci for (%d)\n", job)
	result := ExpensiveFibonacci(job)
	s.Lock.RLock()
	pendingWorkers, exist := s.IsPending[job]
	s.Lock.RUnlock()
	if exist {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result
		}
		fmt.Printf("Result sent to pending workers for job (%d) ready.!\n", job)
	}
	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
		Lock:       sync.RWMutex{},
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 6, 6, 7, 7, 3}
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, num := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(num)
	}
	wg.Wait()
}
