function prefill(n, v) {
  // n must be a non-negative integer, or a string made purely of digits;
  // anything else (booleans, null, floats, negatives, "1.5") is rejected
  const isValid =
    (typeof n === 'number' && Number.isInteger(n) && n >= 0) ||
    (typeof n === 'string' && /^\d+$/.test(n));

  if (!isValid) {
    throw new TypeError(n + ' is invalid');
  }

  return new Array(Number(n)).fill(v);
}