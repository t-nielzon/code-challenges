function sum(n) {
  if (n === undefined) return 0;
  return function next(m) {
    if (m === undefined) return n;
    return sum(n + m);
  };
}