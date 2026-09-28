package main

/*
# Hamster me

Write a function that accepts two inputs: code and message and returns an encrypted string
from message using the code. The code is a string that generates the key as shown below:

All letters from code get number 1. All letters which directly follow letters from code
get number 2 (unless they already have a smaller number assigned), etc.

Example key for code "hamster":
 1  | h a m s t e r
 2  | i b n   u f
 3  | j c o   v g
 4  | k d p   w
 5  | l   q   x
 6  |         y
 7  |         z

Encoding works like:
a => a1
b => a2
c => a3
...

hamsterMe('hamster', 'hamster')   => h1a1m1s1t1e1r1
hamsterMe('hamster', 'helpme')    => h1e1h5m4m1e1

If there is no 'a' in the code, add missing letters after the last available letter
(in alphabetic order) in the code.

The code will have at least 1 letter. Duplication in code should be handled.
Code and message consist of only lowercase letters.
*/

func HamsterMe(code string, message string) string {
	// Implementation goes here
	return ""
}