// solution.go
package kata

import "sync"

func Merge(a <-chan string, b <-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		for val := range a {
			out <- val
		}
	}()
	
	go func() {
		defer wg.Done()
		for val := range b {
			out <- val
		}
	}()
	
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}