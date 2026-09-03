/*
Consider the sequence a(1) = 7, a(n) = a(n-1) + gcd(n, a(n-1)) for n >= 2:
7, 8, 9, 10, 15, 18, 19, 20, 21, 22, 33, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 69, 72, 73...

The differences between successive elements form sequence g (with a 1 added at the head):
g: 1, 1, 1, 1, 5, 3, 1, 1, 1, 1, 11, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 23, 3, 1...

Removing the 1s gives sequence p:
p: 5, 3, 11, 3, 23, 3...

Write functions to:
1. an(n): returns first n terms of sequence a
2. gn(n): returns first n terms of sequence g
3. countOnes(n): returns number of 1s in gn(n)
4. pn(n): returns first n distinct primes from gn(n)
5. maxPn(n): returns biggest prime of first n terms of pn(n)
6. anOver(n): returns array of n terms a(i)/i where g(i) != 1
7. anOverAverage(n): returns average of anOver(n) as integer

Only functions 3, 5, and 7 are tested.
*/

function countOnes(n) {
  
}

function maxPn(n) {
  
}

function anOverAverage(n) {
  
}