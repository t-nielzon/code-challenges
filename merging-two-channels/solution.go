package main

import "sync"

func Merge(a <-chan string, b <-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for msg := range a {
			out <- msg
		}
		wg.Done()
	}()

	go func() {
		for msg := range b {
			out <- msg
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}