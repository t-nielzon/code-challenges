/*
Be u(n) a sequence beginning with:

u[1]  = 1,  u[2]  = 1,  u[3]  = 2,  u[4]  = 3,  u[5]  = 3,  u[6] = 4,
u[7]  = 5,  u[8]  = 5,  u[9]  = 6,  u[10] = 6,  u[11] = 6,  u[12] = 8,
u[13] = 8,  u[14] = 8,  u[15] = 10, u[16] = 9,  u[17] = 10, u[18] = 11,
u[19] = 11, u[20] = 12, u[21] = 12, u[22] = 12, u[23] = 12 etc...

How is u[8] calculated?
We have u[7] = 5 and u[6] = 4. Go backwards from index 8:
  8 - 5 = 3 and 8 - 4 = 4
  u[3] = 2 and u[4] = 3 hence u[8] = 2 + 3 = 5.

Task:
  0) u(n) = u[n - u[n-1]] + u[n - u[n-2]]

  1) LengthSupUK(n, k) returns the number of terms u[i] >= k with 1 <= i <= n.

  2) Comp(n) returns the number of times a term is less than its predecessor
     up to and including u[n].

Examples:
  LengthSupUK(23, 12) => 4
  LengthSupUK(50, 10) => 35
  LengthSupUK(500, 100) => 304

  Comp(23) => 1
  Comp(100) => 22
  Comp(200) => 63
*/
package kata

func LengthSupUK(n, k int) int {
}

func Comp(n int) int {
}