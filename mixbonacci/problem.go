/*
Mixbonacci

Mix -nacci sequences using a given pattern p.
Return the first n elements of the mixed sequence.

Rules:
1. The pattern p is given as a list of strings using the pattern mapping below.
2. When n is 0 or p is empty return an empty list.
3. If n is more than the length of p repeat the pattern.

Pattern mapping:
  'fib' -> fibonacci:   0, 1, 1, 2, 3 ...
  'pad' -> padovan:     1, 0, 0, 1, 0 ...
  'jac' -> jacobsthal:  0, 1, 1, 3, 5 ...
  'pel' -> pell:        0, 1, 2, 5, 12 ...
  'tri' -> tribonacci:  0, 0, 1, 1, 2 ...
  'tet' -> tetranacci:  0, 0, 0, 1, 1 ...
*/
package kata

func Mixbonacci(p []string, n int) []int {
}