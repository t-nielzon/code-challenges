package kata

/*
## Euler's Method

We want to calculate the shape of an unknown curve which starts at a given
point with a given slope. This curve satisfies an ordinary differential
equation (ODE):

	dy/dx = f(x, y);
	y(x0) = y0

The starting point A0 (x0, y0) is known as well as the slope to the curve
at A0 and then the tangent line at A0.

Take a small step along that tangent line up to a point A1. Along this small
step, the slope does not change too much, so A1 will be close to the curve.
If we suppose that A1 is close enough to the curve, the same reasoning as for
the point A1 above can be used for other points. After several steps, a
polygonal curve A0, A1, ..., An is computed. The error between the two curves
will be small if the step is small.

We define points A0, A1, A2, ..., An whose x-coordinates are x0, x1, ..., xn
and y-coordinates are such that y(k+1) = y(k) + f(x(k), y(k)) * h where h is
the common step. If T is the length xn - x0 we have h = T/n.

Task

For this kata we will focus on the following differential equation:

	dy/dx = 2 - e^(-4x) - 2y;
	A0 = (0, 1)

with x in [0, 1]. We will then take a uniform partition of the region of x
between 0 and 1 and split it into n sections (hence n + 1 points). n will be
the input to the function ExEuler(n) and since T is always 1, h = 1/n.

We know that an exact solution is

	z = 1 + 0.5e^(-4x) - 0.5e^(-2x).

For each x(k) we are able to calculate the y(k) as well as the values z(k) of
the exact solution.

Our task is, for a given number n of steps, to return the mean (truncated to
6 decimal places) of the relative errors between the y(k) (our approximation)
and the z(k) (the exact solution). For that we can use:

	error in A(k) = abs(y(k) - z(k)) / z(k)

and then the mean is sum(errors in A(k)) / (n + 1)

Examples

ExEuler(10) should return: 0.026314 (truncated from 0.026314433214799246)
with
Y = [1.0, 0.9..., 0.85..., 0.83..., 0.83..., 0.85..., 0.86..., 0.88..., 0.90..., 0.91..., 0.93...]
Z = [1.0, 0.9..., 0.88..., 0.87..., 0.87..., 0.88..., 0.89..., 0.90..., 0.91..., 0.93..., 0.94...]
Relative errors = [0.0, 0.02..., 0.04..., 0.04..., 0.04..., 0.03..., 0.03..., 0.02..., 0.01..., 0.01..., 0.01...]

ExEuler(17) should return: 0.015193 (truncated from 0.015193336263370796).
As expected, as n increases, our error reduces.

Links: https://en.wikipedia.org/wiki/Euler_method
*/

func ExEuler(n int) float64 {
	// TODO: implement
	return 0
}