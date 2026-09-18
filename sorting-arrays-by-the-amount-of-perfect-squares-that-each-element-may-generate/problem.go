package main

import (
	"math"
	"sort"
	"strconv"
)

/*
You will be given an array of positive integers. The array should be sorted
by the amount of distinct perfect squares and reversed, that can be generated
from each number permuting its digits.

The array should be sorted by:
1. Count of distinct perfect squares that can be generated (descending)
2. If counts are equal, by the number value itself (ascending)

Example:
sort_by_perfsq([715, 112, 136, 169, 144]) == [169, 144, 112, 136, 715]

Features:
- Number of tests: 80
- Arrays between 4 and 20 elements
- Integers having from 1 to 7 digits included
*/
func SortByPerfsq(arr []int) []int {
	// TODO: implement
	return arr
}