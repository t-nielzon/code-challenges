package main

/*
## Euler's Method

We want to calculate the shape of an unknown curve which starts at a given point with a given slope.
This curve satisfies an ordinary differential equation (ODE):

dy/dx = f(x, y); y(x_0) = y_0

The starting point A_0(x_0, y_0) is known as well as the slope to the curve at A_0.

We define points A_0, A_1, A_2, ..., A_n whose x-coordinates are x_0, x_1, ..., x_n and
y-coordinates are such that y_{k+1} = y_k + f(x_k, y_k) × h where h is the common step.
If T is the length x_n - x_0 we have h = T/n.

For this kata we focus on the differential equation:
dy/dx = 2 - e^(-4x) - 2y; A_0 = (0,1) with x ∈ [0, 1]

We split the region into n sections (hence n + 1 points). Since T is always 1, h = 1/n.

The exact solution is: z = 1 + 0.5*e^(-4x) - 0.5*e^(-2x)

Task: For a given number n of steps, return the mean (truncated to 6 decimal places)
of the relative errors between the y_k (approximation) and z_k (exact solution).
Error at A_k = abs(y_k - z_k) / z_k
Mean = sum(errors) / (n + 1)

Examples:
ex_euler(10) should return: 0.026314
ex_euler(17) should return: 0.015193
*/

import "math"

func ExEuler(n int) float64 {
	// TODO: implement
	return 0
}