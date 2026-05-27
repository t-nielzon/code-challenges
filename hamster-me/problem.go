// # Introduction
//
// Hamsters are rodents belonging to the subfamily Cricetinae. And you could
// have skipped the introduction as it is entirely unrelated to your task. xD
//
// # Task
//
// Write a function that accepts two inputs: `code` and `message` and returns
// an encrypted string from `message` using the `code`.
//
// The `code` is a string that generates a key. All letters from `code` get
// number 1. All letters which directly follow letters from `code` get number 2
// (unless they already have a smaller number assigned), etc.
//
// How does the encoding work using the `hamster` code?
//
//	a => a1
//	b => a2
//	c => a3
//	d => a4
//	e => e1
//	f => e2
//	...
//
// And applying it to strings:
//
//	hamsterMe('hamster', 'hamster') => h1a1m1s1t1e1r1
//	hamsterMe('hamster', 'helpme')  => h1e1h5m4m1e1
//
// If there is no `a` in the `code`, just add the missing letters after the last
// available letter (in alphabetic order) in the `code` (the alphabet wraps).
//
// # Additional notes
//
// The `code` will have at least 1 letter.
// Duplication of letters in `code` is possible and should be handled.
// The `code` and `message` consist of only lowercase letters.

package kata

func HamsterMe(code, message string) string {
	return ""
}