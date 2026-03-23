/*
Split and then add both sides of an array together.

Inspired by the Fold an Array kata. This one is sort of similar but a little different.

## Task
You will receive an array as parameter that contains 1 or more integers and a number n.

Step 1: Split the array in two:
  [1, 2, 5, 7, 2, 3, 5, 7, 8]
        /            \
[1, 2, 5, 7]    [2, 3, 5, 7, 8]

Step 2: Put the arrays on top of each other:
     [1, 2, 5, 7]
[2, 3, 5, 7, 8]

Step 3: Add them together:
[2, 4, 7, 12, 15]

Repeat the above steps n times or until there is only one number left, and then return the array.
*/
package kata

func SplitAndAdd(arr []int, n int) []int {
}