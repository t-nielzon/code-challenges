function upperTriangular(mat) {
  for (let i = 1; i < mat.length; i++) {
    for (let j = 0; j < i; j++) {
      if (mat[i][j] !== 0) return false;
    }
  }
  return true;
}

function lowerTriangular(mat) {
  for (let i = 0; i < mat.length - 1; i++) {
    for (let j = i + 1; j < mat[i].length; j++) {
      if (mat[i][j] !== 0) return false;
    }
  }
  return true;
}