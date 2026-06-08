function* paperFold() {
  // for 1-indexed n = 2^k * m with m odd, the term is 1 when m % 4 === 1, else 0.
  for (let n = 1; ; n++) {
    let m = n;
    while ((m & 1) === 0) m >>= 1; // strip trailing zero bits to expose the odd core
    yield (m & 3) === 1 ? 1 : 0;
  }
}