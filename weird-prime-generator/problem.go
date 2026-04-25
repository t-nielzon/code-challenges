/*
Consider the sequence a(1) = 7, a(n) = a(n-1) + gcd(n, a(n-1)) for n >= 2:

7, 8, 9, 10, 15, 18, 19, 20, 21, 22, 33, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 69, 72, 73...

Let us take the differences between successive elements of the sequence and
get a second sequence g: 1, 1, 1, 5, 3, 1, 1, 1, 1, 11, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 23, 3, 1...

For the sake of uniformity of the lengths of sequences we add a 1 at the head of g:
g: 1, 1, 1, 1, 5, 3, 1, 1, 1, 1, 11, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 23, 3, 1...

Removing the 1s gives a third sequence:
p: 5, 3, 11, 3, 23, 3...
where you can see prime numbers.

Task:
- countOnes(n): number of 1s in g(n) (with the 1 added at the head)
- maxPn(n): biggest prime among the first n distinct primes from p
- anOverAverage(n): integer average of a(i)/i over the first n indices i with g(i) != 1
*/

package kata

func countOnes(n int) int {
	return 0
}

func maxPn(n int) int {
	return 0
}

func anOverAverage(n int) int {
	return 0
}