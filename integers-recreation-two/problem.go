/*
We are still with squared integers.

Given 4 integers a, b, c, d we form the sum of the squares of a and b
and then the sum of the squares of c and d. We multiply the two sums hence a number n and we try to
decompose n in a sum of two squares e and f (e and f integers >= 0) so that n = e² + f².

More: e and f must result only from sums (or differences) of products between on the one hand (a, b)
and on the other (c, d) each of a, b, c, d taken only once.

prod2sum(1, 2, 1, 3) should return [[1, 7], [5, 5]]
prod2sum(2, 3, 4, 5) should return [[2, 23], [7, 22]]
prod2sum(1, 2, 2, 3) should return [[1, 8], [4, 7]]
prod2sum(1, 1, 3, 5) should return [[2, 8]]
*/
package kata

func Prod2Sum(a, b, c, d int) [][2]int {

}