/*
 * Euler's Method for a first-order ODE
 * 
 * We want to calculate the shape of an unknown curve which starts at a given point
 * with a given slope. This curve satisfies an ordinary differential equation (ODE):
 * 
 * dy/dx = f(x, y); y(x_0) = y_0
 * 
 * For this kata, we focus on:
 * dy/dx = 2 - e^(-4x) - 2y; with starting point A_0 = (0,1), x ∈ [0, 1]
 * 
 * The exact solution is: z = 1 + 0.5e^(-4x) - 0.5e^(-2x)
 * 
 * Using Euler's method with n steps, we calculate:
 * - h = 1/n (step size)
 * - y_{k+1} = y_k + f(x_k, y_k) * h
 * - relative error at A_k = |y_k - z_k| / z_k
 * - return the mean of all relative errors (truncated to 6 decimal places)
 */

function ex_euler(n) {
  // TODO: implement Euler's method
}