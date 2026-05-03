/*
We'll create a function that takes in two parameters:

* a sequence (length and types of items are irrelevant)
* a function (value, index) that will be called on members of the sequence and their index. The function will return either true or false.

Your function will iterate through the members of the sequence in order until the provided function returns true; at which point your function will return that item's **index**.

If the function given returns false for all members of the sequence, your function should return -1.

Example:
	func isEven(v interface{}, index int) bool {
		return int(v.(int))%2 == 0
	}

	FindInArray([]interface{}{1, 3, 5, 6, 7}, isEven) == 3
*/

package kata

func FindInArray(seq []interface{}, fn func(interface{}, int) bool) int {
}