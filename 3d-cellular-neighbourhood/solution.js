function getNeighbourhood(type, mat, coordinates, distance) {
  if (distance === 0) return [];
  if (!mat || mat.length === 0) return [];

  const [i, j, k] = coordinates;
  const M = mat.length;
  const N = mat[0] ? mat[0].length : 0;
  const K = (mat[0] && mat[0][0]) ? mat[0][0].length : 0;

  if (N === 0 || K === 0) return [];
  if (i < 0 || i >= M || j < 0 || j >= N || k < 0 || k >= K) return [];

  const result = [];
  for (let di = -distance; di <= distance; di++) {
    for (let dj = -distance; dj <= distance; dj++) {
      for (let dk = -distance; dk <= distance; dk++) {
        if (di === 0 && dj === 0 && dk === 0) continue;
        const ni = i + di;
        const nj = j + dj;
        const nk = k + dk;
        if (ni < 0 || ni >= M || nj < 0 || nj >= N || nk < 0 || nk >= K) continue;
        if (type === 'moore') {
          result.push(mat[ni][nj][nk]);
        } else if (type === 'von_neumann') {
          if (Math.abs(di) + Math.abs(dj) + Math.abs(dk) <= distance) {
            result.push(mat[ni][nj][nk]);
          }
        }
      }
    }
  }
  return result;
}