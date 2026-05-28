package kata

/*
The Walker

The walker starts from point O, walks along OA, AB and BC. When he is in C
(C will be in the upper half-plane), what is the distance CO? What is the
angle tOC in positive degrees, minutes, seconds?

Angle tOA is alpha (here 45 degrees), angle hAB is beta (here 30 degrees),
angle uBC is gamma (here 60 degrees).

Task: function Solve(a, b, c, alpha, beta, gamma) with parameters
- a, b, c: positive integers in units of distance (stand for OA, AB, BC)
- alpha, beta, gamma: positive integers in degrees (positive angles are
  anticlockwise)

Returns a slice:
- first element: distance CO (rounded to the nearest integer)
- then angle tOC with the third following elements:
  - second element: number of degrees in angle tOC (truncated positive integer)
  - third element:  number of minutes in angle tOC (truncated positive integer)
  - fourth element: number of seconds in angle tOC (truncated positive integer)

Example:
Solve(12, 20, 18, 45, 30, 60) -> [15, 135, 49, 18]
*/

func Solve(a, b, c, alpha, beta, gamma int) []int {
	return nil
}