function digitalRoot(n) {
  return n === 0 ? 0 : 1 + (n - 1) % 9;
}

module.exports = digitalRoot;