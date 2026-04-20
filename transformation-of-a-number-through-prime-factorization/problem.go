/*
Every natural number, n, may have a prime factorization.

We define the geometric derivative of n, as a number with the following value:
For each prime factor p_i with exponent k_i, compute k_i * p_i^(k_i - 1),
then multiply all these values together.

For example: calculate the value of n* for n = 24500.
    24500 = 2² * 5³ * 7²
    n* = (2*2) * (3*5²) * (2*7) = 4200

Make a function, f, that can perform this calculation:
    f(n) ----> n*

Every prime number will have n* = 1.
Every number that does not have an exponent k_i higher than 1 will give n* = 1, too.

Examples:
    f(24500) == 4200
    f(997)   == 1
*/

package kata

func F(n int) int {
}