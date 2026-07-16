/*
Euler's Method for a First-Order ODE

We want to calculate the shape of an unknown curve which starts at a given point with a given slope.
This curve satisfies an ordinary differential equation (ODE):
  dy/dx = f(x, y); y(x_0) = y_0

The starting point A_0(x_0, y_0) is known as well as the slope to the curve at A_0.
Take a small step along the tangent line up to a point A_1. After several steps,
a polygonal curve A_0, A_1, ..., A_n is computed.

We define points whose x-coordinates are x_0, x_1, ..., x_n and y-coordinates such that:
  y_{k+1} = y_k + f(x_k, y_k) × h
where h is the common step. If T is the length x_n - x_0, then h = T/n.

For this kata, we focus on:
  dy/dx = 2 - e^(-4x) - 2y; A_0 = (0,1)
with x ∈ [0, 1], n steps, and h = 1/n.

The exact solution is:
  z = 1 + 0.5e^(-4x) - 0.5e^(-2x)

Task: Return the mean (truncated to 6 decimal places) of the relative errors between
the approximation y_k and exact solution z_k:
  error in A_k = abs(y_k - z_k) / z_k
  mean = sum(errors) / (n + 1)
*/
package main

func Ex_euler(n int) float64 {
	// Implementation here
}