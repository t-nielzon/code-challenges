package kata

import "math"

/*
The aim of the kata is to try to show how difficult it can be to calculate decimals of an irrational number with a certain precision. We have chosen to get a few decimals of the number "pi" using
the following infinite series (Leibniz 1646–1716):

PI / 4 = 1 - 1/3 + 1/5 - 1/7 + ... which gives an approximation of PI / 4.

http://en.wikipedia.org/wiki/Leibniz_formula_for_%CF%80

To have a measure of the difficulty we will count how many iterations are needed to calculate PI with a given precision of `epsilon`.

There are several ways to determine the precision of the calculus but to keep things easy we will calculate `PI` within epsilon of your language Math::PI constant.

In other words, given as input a precision of `epsilon` we will stop the iterative process when the absolute value of the difference between our calculation using the Leibniz series and the Math::PI constant of your language is less than `epsilon`.

Your function returns an array or a string or a tuple depending on the language (See sample tests) with
- your number of iterations
- your approximation of PI with 10 decimals

Example:
iter_pi(0.001) --> [1000, 3.1405926538]
*/

func IterPi(epsilon float64) [2]float64 {
	_ = math.Pi
	return [2]float64{}
}