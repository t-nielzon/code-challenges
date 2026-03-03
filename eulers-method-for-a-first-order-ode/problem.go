/*
Euler's Method

We want to calculate the shape of an unknown curve which starts at a given point
with a given slope. This curve satisfies an ordinary differential equation (ODE):

  dy/dx = f(x, y)
  y(x_0) = y_0

For this kata we focus on:

  dy/dx = 2 - e^(-4x) - 2y
  A_0 = (0, 1)

with x in [0, 1]. We take a uniform partition into n sections (n+1 points).
h = 1/n.

Exact solution: z = 1 + 0.5*e^(-4x) - 0.5*e^(-2x)

Return the mean (truncated to 6 decimal places) of the relative errors
between the approximated y_k and exact z_k values.

error in A_k = abs(y_k - z_k) / z_k
mean = sum(errors) / (n + 1)
*/
package kata

func ExEuler(n int) float64 {
}