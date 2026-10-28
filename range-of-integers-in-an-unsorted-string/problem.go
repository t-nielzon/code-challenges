/*
Range of Integers in an Unsorted String
Difficulty: 5 kyu

In this kata, your task is to write a function that returns the smallest and largest 
integers in an unsorted string. In this kata, a range is considered a finite sequence 
of consecutive integers.

Input:
- A string comprised of integers in an unknown range; think of this string as the 
  result when a range of integers is shuffled around in random order then joined 
  together into a string
- An integer value representing the size of the range

Output:
Your function should return the starting (minimum) and ending (maximum) numbers of 
the range in the form of an array comprised of two integers.

Example:
inputString := "1568141291110137"
MysteryRange(inputString, 10) // [6 15]

The 10 numbers in this string are:
15 6 8 14 12 9 11 10 13 7
Therefore the range of numbers is from 6 to 15
*/

package kata

func MysteryRange(inputString string, size int) [2]int {
	// TODO
	return [2]int{}
}