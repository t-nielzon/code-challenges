/*
The cuckoo bird pops out of the cuckoo clock and chimes once on the quarter hour, half hour,
and three-quarter hour. At the beginning of each hour (1-12), it chimes out the hour.
Given the current time and a positive integer n, determine the time when the cuckoo bird
has chimed n times.

Inputs:
  initial_time - a string in the format "HH:MM", where 1 <= HH <= 12 and 0 <= MM <= 59,
                 with leading 0's if necessary.
  n - an integer representing the target number of chimes, with 1 <= n <= 200.

Output:
  The time when the cuckoo bird has chimed n times - a string in the same format as initial_time.

If the cuckoo bird chimes at initial_time, include those chimes in the count.
If the n'th chime is reached on the hour, report the time at the start of that hour
(i.e. assume the cuckoo finishes chiming before the minute is up).

Example: "03:38", 19 should return "06:00".
*/

package kata

func CuckooClock(initialTime string, n int) string {
	// your code here
	return ""
}