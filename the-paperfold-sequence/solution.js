function* paperFold() {
  // for index n (1-based), write n = 2^k * m with m odd; the term is
  // 1 when m % 4 === 1 and 0 when m % 4 === 3.
  for (let n = 1; ; n++) {
    let m = n;
    while ((m & 1) === 0) m >>>= 1;
    yield (m & 3) === 1 ? 1 : 0;
  }
}