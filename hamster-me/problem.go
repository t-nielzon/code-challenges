// problem.go
/*
# Introduction

Hamsters are rodents belonging to the subfamily Cricetinae. The subfamily contains about 25 species, classified in six or seven genera. They have become established as popular small house pets, and, partly because they are easy to breed in captivity, hamsters are often used as laboratory animals.

# Task

Write a function that accepts two inputs: code and message and returns an encrypted string from message using the code.

The code is a string that generates the key where:
- All letters from code get number 1
- All letters which directly follow letters from code get number 2 (unless they already have a smaller number assigned), etc.

For example, with code "hamster":
  hamsterMe('hamster', 'hamster')   => h1a1m1s1t1e1r1
  hamsterMe('hamster', 'helpme')    => h1e1h5m4m1e1

If a letter is not in the code, it is added after the last available letter in the code (in alphabetic order).

# Notes
- The code will have at least 1 letter.
- Duplication of letters in code is possible and should be handled.
- The code and message consist of only lowercase letters.
*/

package main

func hamsterMe(code string, message string) string {
	// TODO: Implement the hamster encryption
	return ""
}