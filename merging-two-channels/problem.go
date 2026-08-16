package main

/*
# Merging Two Channels

Write a function `func Merge(a <-chan string, b <-chan string) <-chan string`,
which takes two read-only channels and returns a new channel. All messages from
channel `a` and `b` must be forwarded to the new channel. Once `a` and `b` are
both closed, also the returned channel must be closed.

The order of the forwarded messages doesn't matter, but you should consume from
both incoming channels concurrently.
*/

func Merge(a <-chan string, b <-chan string) <-chan string {
	// TODO: implement
	return nil
}