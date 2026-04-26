/*
 * Euler's Method for a first-order ODE
 *
 * dy/dx = 2 - e^(-4x) - 2y, with A_0 = (0, 1), x in [0, 1]
 * Exact solution: z = 1 + 0.5*e^(-4x) - 0.5*e^(-2x)
 *
 * Given n steps (h = 1/n), compute the Euler approximation y_k and the
 * exact z_k at each x_k = k/n for k = 0..n. Return the mean of the
 * relative errors |y_k - z_k| / z_k, truncated to 6 decimal places.
 *
 * Note: at k = 0, y_0 = z_0 = 1, so the relative error there is 0.
 */

function ex_euler(n) {
  // your code here
}