function going(n) {
  // un = sum_{k=1}^{n} (k! / n!); summing from k=n downwards avoids
  // ever building a huge factorial: the largest term is 1 and each
  // step down divides by the next index.
  let sum = 0;
  let term = 1;
  for (let i = n; i >= 1; i--) {
    sum += term;
    term /= i;
  }
  // truncate to 6 decimals as the kata expects
  return Math.floor(sum * 1e6) / 1e6;
}