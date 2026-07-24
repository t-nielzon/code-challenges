/*
Merging N Channels

Go has a powerful built-in data type: channel. In some cases you want to process messages from different
data sources (e.g. UDP packages, TCP messages, HTTP requests) and you want to consume all in the same way.
Then it might make sense to merge multiple channels into a single one.

Task:
Write a function Merge that takes any amount of channels and returns a new channel.
All messages from the input channels must be forwarded to the new channel.
Once all input channels are closed, also the returned channel must be closed.
The order of the forwarded messages doesn't matter, but you should consume from all incoming channels concurrently.
*/

package main

func Merge(c ...chan string) <-chan string {
	// TODO: implement
	return nil
}