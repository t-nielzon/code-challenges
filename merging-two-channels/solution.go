package main

import "sync"

func Merge(a <-chan string, b <-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		for msg := range a {
			out <- msg
		}
	}()
	
	go func() {
		defer wg.Done()
		for msg := range b {
			out <- msg
		}
	}()
	
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}