// problem.go
/*
Merging Two Channels
Difficulty: 6 kyu

Write a function Merge(a <-chan string, b <-chan string) <-chan string,
which takes two read-only channels and returns a new channel.

All messages from channel a and b must be forwarded to the new channel.
Once a and b are both closed, also the returned channel must be closed.

The order of the forwarded messages doesn't matter, but you should consume
from both incoming channels concurrently.

Example:
- channel a contains 3 messages: "foo", "bar", "baz"
- channel b contains 2 messages: "hello", "world"
- merged channel c must return all 5 messages from a and b
- channel c must be closed after all messages are consumed
*/

package kata

func Merge(a <-chan string, b <-chan string) <-chan string {
}