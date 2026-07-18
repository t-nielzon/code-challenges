/*
Consider the sequence S(n, z) = (1 - z)(z + z**2 + z**3 + ... + z**n) where z is a complex number
and n a positive integer (n > 0).

When n goes to infinity and z has a correct value (ie z is in its domain of convergence D), 
S(n, z) goes to a finite limit lim depending on z.

Experiment with S(n, z) to guess the domain of convergence D of S and lim value when z is in D.
Then determine the smallest integer n such that abs(S(n, z) - lim) < eps where eps is a given 
small real number and abs(Z) is the modulus or norm of the complex number Z.

Call f the function f(z, eps) which returns n.
If z is such that S(n, z) has no finite limit (when z is outside of D) f will return -1.
*/
package main

import (
	"math"
	"math/cmplx"
)

func F(z complex128, eps float64) int {
	// Your implementation here
	return 0
}