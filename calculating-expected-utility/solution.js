function expectedUtility(p0, p1, utilities) {
  let total = 0;
  for (let i = 0; i < p0.length; i++) {
    for (let j = 0; j < p1.length; j++) {
      total += p0[i] * p1[j] * utilities[i][j];
    }
  }
  return Math.round(total * 100) / 100;
}