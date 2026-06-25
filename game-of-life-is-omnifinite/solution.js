/*
 * A diagonal line of cells decays predictably: every step both endpoints die
 * (one live neighbor each) while interior cells survive (two live neighbors
 * each), and no off-diagonal births ever occur. So the line shrinks by one cell
 * on each end per generation.
 *
 * For an odd length L = 2N-1 the lengths go 2N-1, 2N-3, ..., 3, 1 and then the
 * lone cell dies, giving exactly N steps. This also minimizes total live cells
 * across all generations (sum = N^2).
 */

function omnifinite(n) {
  if (n <= 0) return [];

  const length = 2 * n - 1;
  const points = [];
  for (let i = 0; i < length; i++) {
    points.push([i, i]);
  }
  return points;
}

module.exports = omnifinite;