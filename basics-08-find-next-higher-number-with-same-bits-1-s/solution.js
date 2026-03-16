function nextHigher(n) {
  const c = n & -n;
  const r = n + c;
  return (((r ^ n) >> 2) / c) | r;
}

module.exports = nextHigher;