package mutex

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu    sync.Mutex
	count int
)

func increment() {
	mu.Lock() // Lock the mutex to ensure exclusive access to the critical section
	count++
	mu.Unlock() // Unlock the mutex when done
}

func Start() {
	start := time.Now() // Capture the start ti
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(20 * time.Millisecond)
			increment()
		}()
	}

	wg.Wait()
	duration := time.Since(start) // Measure the elapsed time
	fmt.Println("Goroutine Count:", count)
	fmt.Println("Goroutine  Time taken:", duration)

	start = time.Now() // Capture the start ti

	for i := 0; i < 1000; i++ {
		increment()
		time.Sleep(20 * time.Millisecond)
	}
	duration = time.Since(start) // Measure the elapsed time
	fmt.Println("Count:", count)
	fmt.Println("Time taken:", duration)

}
