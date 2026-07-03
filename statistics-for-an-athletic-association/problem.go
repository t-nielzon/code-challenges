// You are the "computer expert" of a local Athletic Association (C.A.A.).
// Many teams of runners come to compete. Each time you get a string of
// all race results of every team who has run.
// For example here is a string showing the individual results of a team of 5 runners:
//
//   "01|15|59, 1|47|6, 01|17|20, 1|32|34, 2|3|17"
//
// Each part of the string is of the form: h|m|s
// where h, m, s (h for hour, m for minutes, s for seconds) are positive or null
// integer (represented as strings) with one or two digits. Substrings in the
// input string are separated by ", " or ",".
//
// To compare the results of the teams you are asked for giving
// three statistics: range, average and median.
//
// - Range: difference between the lowest and highest values.
//   In {4, 6, 9, 3, 7} the lowest value is 3, and the highest is 9,
//   so the range is 9 - 3 = 6.
//
// - Mean or Average: To calculate mean, add together all of the numbers
//   and then divide the sum by the total count of numbers.
//
// - Median: the number separating the higher half of a data sample from the
//   lower half. For an odd number of observations, pick the middle one.
//   For an even number, the median is the mean of the two middle values.
//
// Your task is to return a string giving these 3 values. For the example above:
//
//   "Range: 00|47|18 Average: 01|35|15 Median: 01|32|34"
//
// of the form: "Range: hh|mm|ss Average: hh|mm|ss Median: hh|mm|ss"
// where hh, mm, ss are integers (represented by strings) with each 2 digits.
//
// Remarks:
//   1. if a result in seconds is ab.xy... it will be given truncated as ab.
//   2. if the given string is "" you will return ""

package kata

func Stat(strg string) string {
	return ""
}