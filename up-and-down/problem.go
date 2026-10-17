package main

import (
	"strings"
	"unicode/utf8"
)

/*
TASK: Arrange a string of words into an "up and down" pattern based on word lengths.

Given a string s of space-separated words, return a string t where the word lengths follow a pattern:
length(t[0]) <= length(t[1]) >= length(t[2]) <= length(t[3]) >= ...

RULES:
- You can only swap adjacent or previously-consecutive words
- Minimize the number of swaps
- Format result: positions 0,2,4,... are lowercase; positions 1,3,5,... are uppercase

EXAMPLE:
"after be arrived two My so" → "be ARRIVED two AFTER my SO"
where lengths are: 2 <= 7 >= 3 <= 5 >= 2 <= 2
*/

func arrange(s string) string {
	// Implement the arrange function
}