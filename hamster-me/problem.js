/*
# Hamster me

Write a function that accepts two inputs: `code` and `message` and returns
an encrypted string from `message` using the `code`.

The `code` is a string that generates the key in the way shown below:

 1  | h a m s t e r
 2  | i b n   u f
 3  | j c o   v g
 4  | k d p   w
 5  | l   q   x
 6  |         y
 7  |         z

All letters from `code` get number 1. All letters which directly follow
letters from `code` get number 2 (unless they already have a smaller number),
etc.

Encoding: each letter maps to its column's code letter + its row number.

The `code` will have at least 1 letter.
Duplication of letters in `code` is possible and should be handled.
The `code` and `message` consist of only lowercase letters.
*/

function hamsterMe(code, message) {
}