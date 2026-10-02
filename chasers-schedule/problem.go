/*
Kata: Chaser's schedule
Difficulty: 6 kyu

A runner, who runs with base speed s with duration t will cover a distance d: d = s * t
However, this runner can sprint for one unit of time with double speed s * 2
After sprinting, base speed s will permanently be reduced by 1, and for the next one unit of time runner will enter recovery phase and can't sprint again.

Your task, given base speed s and time t, is to find the maximum possible distance d.

Input:
1 <= s < 1000
1 <= t < 1000

Example:
Given s = 2 and t = 4.
Possible sequences include:
- RRRS: 2 + 2 + 2 + 4 = 10
- RRSR: 2 + 2 + 4 + 1 = 9
- RSRS: 2 + 4 + 1 + 2 = 9
- SRRR: 4 + 1 + 1 + 1 = 7

The maximum possible distance d is 10.

Where:
- R: Normal Run / Recovery
- S: Sprint
*/

package main

func chaserScore(s, t int) int {
	return 0
}