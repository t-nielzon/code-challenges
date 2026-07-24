package main

import "sync"

func Merge(c ...chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	// spawn a goroutine for each input channel to forward messages concurrently
	for _, ch := range c {
		wg.Add(1)
		go func(channel chan string) {
			defer wg.Done()
			for msg := range channel {
				out <- msg
			}
		}(ch)
	}

	// wait for all input channels to be consumed, then close the output channel
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}