var isPP = function(n) {
  for (var k = 2; k <= Math.log2(n); k++) {
    var m = Math.round(Math.pow(n, 1 / k));
    for (var candidate = m - 1; candidate <= m + 1; candidate++) {
      if (candidate > 1 && Math.pow(candidate, k) === n) {
        return [candidate, k];
      }
    }
  }
  return null;
};