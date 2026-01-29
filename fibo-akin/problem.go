/*
Be u(n) a sequence beginning with:

u[1]  = 1,  u[2]  = 1,  u[3]  = 2,  u[4]  = 3,  u[5]  = 3,  u[6] = 4,
u[7]  = 5,  u[8]  = 5,  u[9]  = 6,  u[10] = 6,  u[11] = 6,  u[12] = 8,
u[13] = 8,  u[14] = 8,  u[15] = 10, u[16] = 9,  u[17] = 10, u[18] = 11,
u[19] = 11, u[20] = 12, u[21] = 12, u[22] = 12, u[23] = 12 etc...

- How is u[8] calculated?

We have u[7] = 5 and u[6] = 4. These numbers tell us that we have to go backwards
from index 8 to index 8 - 5 = 3 and to index 8 - 4 = 4 so to index 3 and 4.

u[3] = 2 and u[4] = 3 hence u[8] = u[3] + u[4] = 2 + 3 = 5.

- Another example: let us calculate u[13]. At indexes 12 and 11 we have 8 and 6.
Going backwards of 8 and 6 from 13 we get indexes 13 - 8 = 5 and 13 - 6 = 7.

u[5] = 3 and u[7] = 5 so u[13] = u[5] + u[7] = 3 + 5 = 8.

Task:
- 1) Given two numbers n, k (integers > 2) write the function LengthSupUK(n, k)
     returning the number of terms u[i] >= k with 1 <= i <= n.

- 2) Given n (integer > 2) write the function Comp(n) returning the number of times
     where a term of u is less than its predecessor up to and including u[n].
*/
package kata

func LengthSupUK(n, k int) int {
}

func Comp(n int) int {
}