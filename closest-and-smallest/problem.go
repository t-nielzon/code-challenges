/*
#### Input

- a string `strng` of n positive numbers (n = 0 or n >= 2)

Let us call weight of a number the sum of its digits.
For example `99` will have "weight" `18`, `100` will have "weight" `1`.

Two numbers are "close" if the difference of their weights is small.

#### Task:
For each number in `strng` calculate its "weight" and then find *two* numbers
of `strng` that have:

- the smallest difference of weights ie that are the closest
- with the smallest weights
- and with the smallest indices (or ranks, numbered from 0) in `strng`

#### Output:

- a string in Go mimicking an array of two subarrays:
  "[(weight, index, number), (weight, index, number)]"
  or "[(), ()]" if n == 0.

The two subarrays are sorted in ascending order by their number weights if
these weights are different, by their indexes in the string if they have the
same weights.
*/

package kata

func Closest(strng string) string {
	return ""
}