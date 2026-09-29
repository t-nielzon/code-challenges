// Introduction
// Go has a powerful built-in data type: channel. In some cases you want process messages
// from different data sources (e.g. UDP packages, TCP messages, HTTP requests) and you want
// to consume all in the same way. Then it might make sense to merge multiple channels into
// a single one. This is what we do here.
//
// Learning Goal
// In this kata we learn to merge two channels into a single one by utilizing concurrency patterns.
//
// Task
// Write a function Merge(a <-chan string, b <-chan string) <-chan string, which takes two
// read-only channels and returns a new channel. All messages from channel a and b must be
// forwarded to the new channel. Once a and b are both closed, also the returned channel must
// be closed.
//
// The order of the forwarded messages doesn't matter, but you should consume from both
// incoming channels concurrently.

package kata

func Merge(a <-chan string, b <-chan string) <-chan string {
}