package main

/*
Simpson's Rule - Approximate Integration

An integral:
∫[a,b] f(x)dx

can be approximated by Simpson's rule:
(b-a)/(3n) * (f(a) + f(b) + 4*Σ(i=1 to n/2)[f(a+(2i-1)*h)] + 2*Σ(i=1 to n/2-1)[f(a+2i*h)])

Where:
- h = (b - a) / n
- n is an even integer
- a <= b

We want to apply Simpson's rule with the function:
f(x) = (3/2) * sin(x)^3

Write a function called `simpson` with parameter `n` which returns the value of the 
integral of f on the interval [0, π].

Notes:
- Don't round or truncate your results
- n will always be even
- The exact value of the integral of f on [0, π] is 2
*/

import "math"

func simpson(n int) float64 {
	// Your solution here
	return 0
}