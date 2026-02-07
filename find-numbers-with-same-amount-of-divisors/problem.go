/*
Find Numbers with Same Amount of Divisors

The integers 14 and 15 are contiguous (i.e. the difference between them is 1)
and have the same number of divisors:

14 ----> 1, 2, 7, 14  (4 divisors)
15 ----> 1, 3, 5, 15  (4 divisors)

The next pair of contiguous integers with this property is 21 and 22:

21 -----> 1, 3,  7, 21  (4 divisors)
22 -----> 1, 2, 11, 22  (4 divisors)

We have 8 pairs of integers below 50 having this property.

We need a function that receives two integer parameters, diff and nMax,
and outputs the count of pairs of integers that fulfill this property,
all of them being strictly smaller than nMax.

(diff = 1, nMax =  50) -----> 8
(diff = 3, nMax = 100) -----> 7
*/
package kata

func CountPairsInt(diff, nMax int) int {
}