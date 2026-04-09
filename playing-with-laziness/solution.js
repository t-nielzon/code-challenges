function find(matrix) {
  for (let n = 0; ; n++) {
    for (let j = 0; j <= n; j++) {
      if (matrix[n][j]) return [n, j];
    }
    for (let i = 0; i < n; i++) {
      if (matrix[i][n]) return [i, n];
    }
  }
}