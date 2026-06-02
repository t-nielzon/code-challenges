function prefill(n, v) {
  if (!/^\d+$/.test(String(n))) {
    throw new TypeError(`${n} is invalid`);
  }

  return Array(Number(n)).fill(v);
}