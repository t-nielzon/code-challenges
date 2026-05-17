/*
 * Euler's Method for a first-order ODE
 *
 * We want to approximate the solution of:
 *   dy/dx = 2 - e^(-4x) - 2y,  y(0) = 1,  x ∈ [0, 1]
 *
 * Using Euler's method with n steps (h = 1/n):
 *   y_{k+1} = y_k + f(x_k, y_k) * h
 *
 * The exact solution is:
 *   z = 1 + 0.5 * e^(-4x) - 0.5 * e^(-2x)
 *
 * For each x_k compute y_k (Euler approximation) and z_k (exact).
 * Relative error at A_k = |y_k - z_k| / z_k  (0 when z_k == y_k, e.g. at x=0).
 * Return the mean of relative errors, truncated to 6 decimal places.
 *
 * Examples:
 *   ex_euler(10) -> 0.026314
 *   ex_euler(17) -> 0.015193
 */

function ex_euler(n) {
  // your code here
}