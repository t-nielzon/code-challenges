function isPolydivisible(n) {
  const digits = String(n).split('');
  let prefix = 0;
  for (let i = 0; i < digits.length; i++) {
    prefix = prefix * 10 + Number(digits[i]);
    if (prefix % (i + 1) !== 0) return false;
  }
  return true;
}

function next(n) {
  // polydivisible numbers are finite; the largest has 25 digits, so any
  // candidate beyond that range cannot be polydivisible.
  const MAX = 3608528850368400786036725n;
  let candidate = BigInt(n) + 1n;
  while (candidate <= MAX) {
    if (isPolydivisible(candidate)) return Number(candidate) <= Number.MAX_SAFE_INTEGER ? Number(candidate) : candidate;
    candidate++;
  }
  return null;
}