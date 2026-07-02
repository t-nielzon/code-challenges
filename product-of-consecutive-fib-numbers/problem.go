// Product of consecutive Fib numbers
//
// The Fibonacci numbers are the numbers in the following integer sequence (Fn):
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, ...
//
// such that:
//   F(0) = 0
//   F(1) = 1
//   F(n) = F(n-1) + F(n-2)
//
// Given a number, say prod (for product), we search two Fibonacci numbers
// F(n) and F(n+1) verifying:
//   F(n) * F(n+1) = prod
//
// Your function takes an integer (prod) and returns an array:
//   - if F(n) * F(n+1) = prod: [F(n), F(n+1), 1]
//   - if you do not find two consecutive F(n) verifying F(n) * F(n+1) = prod:
//     [F(n), F(n+1), 0]
//     where F(n) is the smallest one such as F(n) * F(n+1) > prod.
//
// Examples:
//   714 ---> [21, 34, 1]  since 714 = 21 * 34
//   800 ---> [34, 55, 0]  since 21 * 34 < 800 < 34 * 55

package kata

func ProductFib(prod uint64) [3]uint64 {
	return [3]uint64{}
}